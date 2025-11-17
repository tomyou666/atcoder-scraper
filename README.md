# atcs

AtCoderの問題を取得してローカルに保存するツールです。

## 機能

- AtCoderの問題ページから問題文、制約、入力形式、画像を取得
- JSON形式で保存
- 画像の自動ダウンロード

## インストール方法

### 方法1: GitHub Releasesからダウンロード（推奨）

1. [Releases](https://github.com/tomyou666/atcoder-scraper/releases)ページにアクセス
2. お使いのOSとアーキテクチャに合ったバイナリをダウンロード：
   - **Windows (amd64)**: `atcs-windows-amd64.exe`
   - **Windows (arm64)**: `atcs-windows-arm64.exe`
   - **Linux (amd64)**: `atcs-linux-amd64`
   - **Linux (arm64)**: `atcs-linux-arm64`
   - **macOS (amd64)**: `atcs-darwin-amd64`
   - **macOS (arm64)**: `atcs-darwin-arm64`

3. ダウンロードしたファイルを実行可能にして、PATHに追加

#### Windowsの場合
```powershell
# ダウンロードしたファイルを適切な場所に移動（例：C:\tools\）
# 環境変数PATHに追加
```

#### Linuxの場合（curlでインストール）

```bash
# 最新バージョンをダウンロード（amd64の場合）
curl -L https://github.com/tomyou666/atcoder-scraper/releases/latest/download/atcs-linux-amd64 -o atcs

# arm64の場合
# curl -L https://github.com/tomyou666/atcoder-scraper/releases/latest/download/atcs-linux-arm64 -o atcs

# 実行可能にする
chmod +x atcs

# PATHに追加（例：/usr/local/bin/）
sudo mv atcs /usr/local/bin/

# インストール確認
atcs --help
```

### 方法2: ソースからビルド

#### 前提条件
- Go 1.24以上がインストールされていること

#### ビルド手順

```bash
# リポジトリをクローン
git clone git@github.com:tomyou666/atcoder-scraper.git atcs
cd atcs

# 依存関係をインストール
go mod download

# ビルド
go build -o atcs .

# Windowsの場合
go build -o atcs.exe .
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

このプロジェクトは [MIT License](LICENSE) の下で公開されています。

MIT Licenseにより、以下のことが許可されています：
- ✅ 商用利用
- ✅ 改変
- ✅ 配布
- ✅ 私的使用
- ✅ 特許利用

条件として、ライセンス表示と著作権表示を含める必要があります。

詳細は [LICENSE](LICENSE) ファイルを参照してください。

## 貢献

プルリクエストやイシューの報告を歓迎します！

