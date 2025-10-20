# AI Agent Platform

AIエージェントを作成・管理・実行できるプラットフォーム

## 🎯 プロジェクト概要

このプラットフォームでは以下のことができます：
- 複数のAIエージェントを作成
- エージェントごとにカスタマイズ可能
- チャット形式でエージェントと対話
- 実行履歴の管理

## 🛠️ 技術スタック

- **フロントエンド**: React + TypeScript + Vite
- **バックエンド**: Go + Gin Framework
- **認証**: Keycloak (OAuth 2.0 / OIDC)
- **データベース**: PostgreSQL
- **AI**: OpenAI API / Google Gemini
- **インフラ**: Docker + Kubernetes + GCP
- **CI/CD**: GitHub Actions

## 📁 プロジェクト構造
```
ai-agent-platform/
├── src/
│   ├── backend/          # Go API サーバー
│   └── frontend/         # React アプリケーション
├── docs/                 # ドキュメント
├── tests/                # テストコード
├── scripts/              # 便利スクリプト
├── config/               # 設定ファイル
├── _PROMPT/              # プロンプト管理
└── logs/                 # ログファイル
```

## 🚀 開発環境セットアップ

### 必要なもの
- GitHub Codespaces（推奨）
- Docker & Docker Compose
- Go 1.21+
- Node.js 20+

### 起動方法
```bash
# Docker サービスを起動
docker-compose up -d

# バックエンド起動
cd src/backend
go run main.go

# フロントエンド起動（別ターミナル）
cd src/frontend
npm install
npm run dev
```

## 📚 ドキュメント

詳細なドキュメントは `docs/` フォルダを参照してください。

## 🌿 ブランチ戦略

- `main`: 本番環境用の安定版
- `v1.0.0_feature`: 機能開発用ブランチ

## 📄 ライセンス

MIT License

## 👤 作成者

[Your Name]
