package store

import (
	"gpt-load/internal/types"

	"github.com/sirupsen/logrus"
)

// NewStore creates a new store based on the application configuration.
func NewStore(cfg types.ConfigManager) (Store, error) {
	logrus.Info("Using in-memory store.")
	return NewMemoryStore(), nil
}
