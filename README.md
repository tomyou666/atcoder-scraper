# atc

AtCoderの問題を取得してローカルに保存するツールです。

## 機能

- AtCoderの問題ページから問題文、制約、入力形式、画像を取得
- JSON形式で保存
- 画像の自動ダウンロード

## インストール方法

### 方法1: GitHub Releasesからダウンロード（推奨）

1. [Releases](https://github.com/your-username/atc/releases)ページにアクセス
2. お使いのOSとアーキテクチャに合ったバイナリをダウンロード：
   - **Windows (amd64)**: `atc-windows-amd64.exe`
   - **Windows (arm64)**: `atc-windows-arm64.exe`
   - **Linux (amd64)**: `atc-linux-amd64`
   - **Linux (arm64)**: `atc-linux-arm64`
   - **macOS (amd64)**: `atc-darwin-amd64`
   - **macOS (arm64)**: `atc-darwin-arm64`

3. ダウンロードしたファイルを実行可能にして、PATHに追加

#### Windowsの場合
```powershell
# ダウンロードしたファイルを適切な場所に移動（例：C:\tools\）
# 環境変数PATHに追加
```

#### Linux/macOSの場合
```bash
# ダウンロードしたファイルを実行可能にする
chmod +x atc-linux-amd64

# 適切な場所に移動（例：/usr/local/bin/）
sudo mv atc-linux-amd64 /usr/local/bin/atc
```

### 方法2: ソースからビルド

#### 前提条件
- Go 1.24以上がインストールされていること

#### ビルド手順

```bash
# リポジトリをクローン
git clone https://github.com/your-username/atc.git
cd atc

# 依存関係をインストール
go mod download

# ビルド
go build -o atc .

# Windowsの場合
go build -o atc.exe .
```

## 使用方法

```bash
atc <AtCoderの問題URL> [出力ディレクトリ名/ファイル名]
```

### 例

```bash
# 基本的な使用方法
atc https://atcoder.jp/contests/abc123/tasks/abc123_a

# 出力先を指定
atc https://atcoder.jp/contests/abc123/tasks/abc123_a problem_data

# ファイル名を指定
atc https://atcoder.jp/contests/abc123/tasks/abc123_a output.json
```

## 出力形式

問題データはJSON形式で保存されます：

```json
{
  "problem": "問題文の内容",
  "constraints": "制約の内容",
  "input": "入力形式の説明",
  "images": ["画像URL1", "画像URL2"]
}
```

画像は自動的にダウンロードされ、JSONファイルと同じディレクトリに保存されます。

## 開発

### リリースの作成

新しいバージョンをリリースするには：

```bash
# タグを作成
git tag v1.0.0

# タグをプッシュ（これでGitHub Actionsが自動的にビルドとリリースを作成）
git push origin v1.0.0
```

GitHub Actionsが自動的に複数プラットフォーム用のバイナリをビルドし、GitHub Releasesにアップロードします。

## ライセンス

（ライセンスを追加してください）

## 貢献

プルリクエストやイシューの報告を歓迎します！

