# Claude Code Prompt Directory Preset (ccpdp)

## プロジェクト概要
このプロジェクトは、Claude Codeでコーディングする際に必要なプロンプトファイルやメモリーファイルを自動生成するCLIツールです。シンプルなコマンドでディレクトリ構造を作成できます。

## 基本的な使用方法
```bash
# 基本形式
ccpdp "親ディレクトリ > 子ディレクトリ"

# 実例
ccpdp "docs > instructions" "docs > logs"
# → docs/instructions/ と docs/logs/ が作成される

# プリセットを使用
ccpdp --preset claude-basic
```

## 主な機能
- コマンドライン引数による柔軟なディレクトリ指定（`>` で階層を表現）
- よく使う構成をプリセットとして利用
- 必要なファイル（.gitkeep、README.md等）の自動生成
- Claude Code用の設定ファイルの配置

## 技術スタック
- Go 1.21+
- Cobra (CLIフレームワーク)

## プロジェクト構成
```
/
├── cmd/
│   └── claude-preset/     # エントリーポイント
├── internal/
│   ├── cli/              # CLIインターフェース処理
│   ├── generator/        # ディレクトリ・ファイル生成ロジック
│   └── templates/        # プロジェクトテンプレート定義
├── pkg/                  # 公開パッケージ
├── go.mod
└── go.sum
```

## 開発ガイドライン

### コーディング規約
- 標準的なGoの命名規則に従う
- すべての公開関数にはコメントを記載
- エラーハンドリングは適切に行う

### テスト方針
- すべての公開関数および主要な内部関数にユニットテストを実装
- テストカバレッジ80%以上を目標
- 境界値テストとエラーケースのテストを網羅的に実施

### 主要なデータ構造
```go
type ProjectConfig struct {
    Purpose         string              // プロジェクトの目的
    Architecture    string              // 選択されたアーキテクチャ
    Directories     []string            // 生成するディレクトリリスト
    PromptTemplates map[string]string   // プロンプトテンプレート
}

type Template struct {
    Name        string                  // テンプレート名
    Description string                  // テンプレートの説明
    Structure   map[string][]string     // ディレクトリ構造
    Prompts     map[string]string       // プロンプトファイル
}
```

## 実装の流れ
1. **初期化フェーズ**: Cobraでコマンドを初期化
2. **選択フェーズ**: ユーザーにプロジェクトの目的とアーキテクチャを選択させる
3. **生成フェーズ**: 選択に基づいてディレクトリとファイルを生成
4. **完了フェーズ**: 生成結果を表示

## 注意事項
- 既存ファイルの上書きは行わない
- ファイルシステムへの書き込み権限が必要
- クロスプラットフォーム対応を考慮した実装

## 今後の拡張予定
- プロジェクトテンプレートのカスタマイズ機能
- 既存プロジェクトへのClaude Code設定追加機能
- プロンプトテンプレートの管理機能