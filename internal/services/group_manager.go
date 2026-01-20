package services

import (
	"context"
	"fmt"
	"gpt-load/internal/config"
	"gpt-load/internal/models"
	"gpt-load/internal/store"
	"gpt-load/internal/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const GroupUpdateChannel = "groups:updated"

// GroupManager manages the caching of group data.
type GroupManager struct {
	groups          map[string]*models.Group
	db              *gorm.DB
	store           store.Store
	settingsManager *config.SystemSettingsManager
}

// NewGroupManager creates a new, uninitialized GroupManager.
func NewGroupManager(
	db *gorm.DB,
	store store.Store,
	settingsManager *config.SystemSettingsManager,
) *GroupManager {
	return &GroupManager{
		db:              db,
		store:           store,
		settingsManager: settingsManager,
	}
}

// Initialize sets up the cache.
func (gm *GroupManager) Initialize() error {
	return gm.Reload()
}

// Reload fetches the latest data and updates the cache.
func (gm *GroupManager) Reload() error {
	var groups []*models.Group
	if err := gm.db.Find(&groups).Error; err != nil {
		return fmt.Errorf("failed to load groups from db: %w", err)
	}

	groupMap := make(map[string]*models.Group, len(groups))
	for _, group := range groups {
		g := *group
		g.EffectiveConfig = gm.settingsManager.GetEffectiveConfig(g.Config)
		g.ProxyKeysMap = utils.StringToSet(g.ProxyKeys, ",")
		groupMap[g.Name] = &g
		logrus.WithFields(logrus.Fields{
			"group_name":       g.Name,
			"effective_config": g.EffectiveConfig,
		}).Debug("Loaded group with effective config")
	}

	gm.groups = groupMap
	return nil
}

// GetGroupByName retrieves a single group by its name from the cache.
func (gm *GroupManager) GetGroupByName(name string) (*models.Group, error) {
	group, ok := gm.groups[name]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return group, nil
}

// Invalidate triggers a cache reload.
func (gm *GroupManager) Invalidate() error {
	return gm.Reload()
}

// Stop gracefully stops the GroupManager.
func (gm *GroupManager) Stop(ctx context.Context) {
}
