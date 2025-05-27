# 仕様書

## 1. 概要
### 1.1 プロジェクト名
Claude Code Prompt Directory Preset

### 1.2 目的
Claude Code でコーディングするときに、プロンプトやClaude Code が記録するメモリー的なテキストを入力しておく保管所を自動で作成するCLIツールです。

### 1.3 スコープ
- プロジェクトの種類に応じたディレクトリ構成の自動生成
- Claude Code用のプロンプトファイルの配置
- プロジェクトテンプレートの選択と適用

## 2. 要件定義
### 2.1 機能要件
#### 2.1.1 必須機能

**基本的な使用方法：**
```bash
# コマンド形式
ccpdp "親ディレクトリ > 子ディレクトリ1" "親ディレクトリ > 子ディレクトリ2"

# 使用例
ccpdp "docs > instructions" "docs > logs"
# → docs/instructions/ と docs/logs/ が作成される

# 複数階層の例
ccpdp "src > components > ui" "src > utils" "tests > unit"
# → src/components/ui/, src/utils/, tests/unit/ が作成される
```

**主な機能：**
1. **コマンドライン引数によるディレクトリ指定**
   - `>` 記号で階層を表現
   - 複数のパスを同時に指定可能
   - 親ディレクトリが存在しない場合は自動作成

2. **プリセットモード（オプション）**
   - よく使うディレクトリ構成をプリセットとして登録
   - `ccpdp --preset web` のような形で呼び出し

3. **必要なファイルの自動生成**
   - 各ディレクトリに適切な `.gitkeep` や `README.md` を配置
   - Claude Code用の設定ファイルを適切な場所に配置

#### 2.1.2 オプション機能
- プロジェクトテンプレートのカスタマイズ機能
- 既存プロジェクトへのClaude Code設定追加機能
- プロンプトテンプレートの管理機能

### 2.2 非機能要件
#### 2.2.1 パフォーマンス
- ディレクトリ生成は1秒以内に完了すること

#### 2.2.2 セキュリティ
- ファイルシステムへの書き込み権限の適切な確認

#### 2.2.3 可用性
- クロスプラットフォーム対応（Windows, macOS, Linux）

#### 2.2.4 保守性
- モジュール化された設計
- 十分なテストカバレッジ

## 3. システム構成
### 3.1 アーキテクチャ
標準的なGoプロジェクト構成を採用：
```
/
├── cmd/
│   └── claude-preset/
│       └── main.go
├── internal/
│   ├── cli/
│   ├── generator/
│   └── templates/
├── pkg/
├── go.mod
└── go.sum
```

### 3.2 技術スタック
- **言語**: Go 1.21+
- **CLIフレームワーク**: Cobra
- **依存関係管理**: Go Modules

### 3.3 インフラ構成
特になし（スタンドアロンCLIツール）

## 4. プリセット一覧

### 4.1 実装予定のプリセット

#### 4.1.1 基本プリセット
```bash
# Claude Code基本構成
ccpdp --preset claude-basic
# 作成されるディレクトリ:
# → .claude/instructions/
# → .claude/logs/
# → docs/specifications/
# → docs/references/

# ドキュメント管理用
ccpdp --preset docs
# → docs/instructions/
# → docs/logs/
# → docs/specifications/
# → docs/references/
# → docs/decisions/

# プロンプト管理用
ccpdp --preset prompts
# → prompts/system/
# → prompts/user/
# → prompts/examples/
# → prompts/templates/

# プロジェクト情報管理用
ccpdp --preset project-info
# → project/requirements/
# → project/architecture/
# → project/decisions/
# → project/meetings/
```

#### 4.1.2 カスタムプリセット定義例
```yaml
# ~/.ccpdp/presets/my-project.yaml
name: my-project
description: "My custom project structure"
directories:
  - "src > components > common"
  - "src > components > features"
  - "src > services"
  - "src > utils"
  - "tests > unit"
  - "tests > e2e"
  - "docs > architecture"
  - "docs > api"
  - ".claude > instructions"
files:
  - path: ".claude/instructions/coding-style.md"
    template: "coding-style"
  - path: "docs/README.md"
    template: "project-readme"
```

### 4.2 追加希望のプリセット
<!-- ここに新しいプリセットのアイデアを追加してください -->
- 
- 
- 

### 4.3 各プリセットに含まれる内容
各プリセットには以下の要素が含まれます：
- ディレクトリ構成
- 各ディレクトリの `.gitkeep` ファイル
- 必要に応じた `README.md` ファイル

## 5. データ設計
### 5.1 データモデル
```go
type ProjectConfig struct {
    Purpose         string
    Architecture    string
    Directories     []string
    PromptTemplates map[string]string
}

type Template struct {
    Name        string
    Description string
    Structure   map[string][]string
    Prompts     map[string]string
}
```

### 5.2 データベース設計
該当なし（設定はJSONファイルで管理）

### 5.3 API仕様
該当なし（CLIツール）

## 6. 画面設計
### 6.1 画面一覧
CLIインターフェース：
- 初期メニュー画面
- プロジェクト目的選択画面
- アーキテクチャ選択画面
- 確認画面
- 実行結果表示画面

### 6.2 画面遷移図
```
開始 → 目的選択 → アーキテクチャ選択 → 確認 → 生成 → 完了
```

### 6.3 UI/UXガイドライン
- 明確で簡潔なプロンプト表示
- 矢印キーでの選択操作
- 進捗状況の表示

## 7. 実装詳細
### 7.1 モジュール構成
- **cmd/claude-preset**: エントリーポイント
- **internal/cli**: CLIインターフェース処理
- **internal/generator**: ディレクトリ・ファイル生成ロジック
- **internal/templates**: プロジェクトテンプレート定義

### 7.2 主要な関数
```go
// CLIの初期化と実行
func Execute() error

// プロジェクト生成
func GenerateProject(config ProjectConfig) error

// テンプレート適用
func ApplyTemplate(template Template, path string) error
```

## 8. テスト
### 8.1 テスト方針
すべての公開関数および主要な内部関数に対してユニットテストを実装します。

### 8.2 テストケース
- 各プロジェクトタイプの生成テスト
- ディレクトリ作成の成功・失敗ケース
- テンプレート適用のテスト
- エラーハンドリングのテスト

### 8.3 受け入れ基準
- テストカバレッジ80%以上
- 境界値テストの実施
- エラーケースの網羅的なテスト

## 9. 運用
### 9.1 デプロイ手順
1. `go build` でバイナリをビルド
2. GitHubリリースにバイナリをアップロード
3. Homebrewフォーミュラの更新（オプション）

### 9.2 監視項目
該当なし（スタンドアロンツール）

### 9.3 バックアップ方針
該当なし

## 10. 制約事項
- Go 1.21以上が必要
- ファイルシステムへの書き込み権限が必要
- 既存ファイルの上書きは行わない

## 11. 用語集
- **Claude Code**: Anthropic社のAIアシスタントツール
- **プロンプト**: Claude Codeに与える指示や文脈情報
- **CLAUDE.md**: Claude Codeがプロジェクトの文脈を記憶するためのファイル

## 12. 参考資料
- [Cobra Documentation](https://cobra.dev/)
- [Claude Code Documentation](https://docs.anthropic.com/claude-code)
- [Go Project Layout](https://github.com/golang-standards/project-layout)

## 13. 更新履歴
| 日付 | バージョン | 更新内容 | 更新者 |
|------|------------|----------|--------|
| 2025/01/27 | 1.0 | 初版作成 | |