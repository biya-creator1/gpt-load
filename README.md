# GPT-Load (Gemini Edition)

中文文档 | [English](README_EN.md)

[![Release](https://img.shields.io/github/v/release/tbphp/gpt-load)](https://github.com/tbphp/gpt-load/releases)
![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

一个轻量、高效的 **Google Gemini** 接口透明代理服务。采用 Go 语言开发，具备智能密钥管理、负载均衡和完善的监控功能，专为 Gemini 开发者设计。

详细请查看[官方文档](https://www.gpt-load.com/docs)

<a href="https://hellogithub.com/repository/tbphp/gpt-load" target="_blank"><img src="https://api.hellogithub.com/v1/widgets/recommend.svg?rid=554dc4c46eb14092b9b0c56f1eb9021c&claim_uid=Qlh8vzrWJ0HCneG" alt="Featured｜HelloGitHub" style="width: 250px; height: 54px;" width="250" height="54" /></a>

## 功能特性

- **Gemini 透明代理**: 完全保留 Google Gemini 原生 API 格式，支持 Gemini 1.5 Pro/Flash 等
- **智能密钥管理**: 高性能 Gemini API Key 密钥池，支持分组管理、自动轮换和故障恢复
- **负载均衡**: 支持多上游端点的加权负载均衡，提升 Gemini 服务可用性
- **智能故障处理**: 自动密钥黑名单管理和恢复机制，确保服务连续性
- **轻量化架构**: 默认使用 SQLite 和内存缓存，无需外部数据库和 Redis，极速启动
- **现代化管理**: 基于 Vue 3 的 Web 管理界面，直观易用
- **全面监控**: 实时统计、健康检查、详细请求日志
- **高性能设计**: 零拷贝流式传输、连接池复用、原子操作

## 支持的 AI 服务

GPT-Load 作为透明代理服务，完整保留 Google Gemini 的原生 API 格式：

- **Google Gemini 格式**: Gemini 1.5 Pro、Gemini 1.5 Flash、Gemini Pro Vision 等模型的原生 API

## 快速开始

### 方式一：Docker 快速开始

```bash
docker run -d --name gpt-load \
    -p 3001:3001 \
    -e AUTH_KEY=sk-123456 \
    -v "$(pwd)/data":/app/data \
    ghcr.io/tbphp/gpt-load:latest
```

> 使用 `sk-123456` 登录管理界面：<http://localhost:3001>

### API 调用示例

假设创建了名为 `gemini` 的分组：

**原始调用方式：**

```bash
curl -X POST https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-pro:generateContent?key=your-gemini-key \
  -H "Content-Type: application/json" \
  -d '{"contents": [{"parts": [{"text": "Hello"}]}]}'
```

**代理调用方式：**

```bash
curl -X POST http://localhost:3001/proxy/gemini/v1beta/models/gemini-1.5-pro:generateContent?key=your-proxy-key \
  -H "Content-Type: application/json" \
  -d '{"contents": [{"parts": [{"text": "Hello"}]}]}'
```

**变更说明：**

- 将 `https://generativelanguage.googleapis.com` 替换为 `http://localhost:3001/proxy/gemini`
- 将 URL 参数中的 `key=your-gemini-key` 替换为**代理密钥**

### Python SDK 调用示例

```python
import google.generativeai as genai

# 配置 API 密钥和基础 URL
genai.configure(
    api_key="your-proxy-key",  # 使用代理密钥
    client_options={"api_endpoint": "http://localhost:3001/proxy/gemini"}
)

model = genai.GenerativeModel('gemini-1.5-pro')
response = model.generate_content("Hello")
```

## 许可证

MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。
