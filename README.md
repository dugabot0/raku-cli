# raku-cli

楽天ウェブサービス API の CLI ツールです。市場・書籍・トラベル・Kobo・GORA・レシピをコマンドラインから検索できます。

## インストール

```bash
git clone https://github.com/dugabot0/raku-cli.git
cd raku-cli
go build -o raku-cli .
```

## 設定

`~/.config/raku-cli/config.yaml` を作成してください。

```yaml
rakuten:
  app_id: "your-app-id"
  affiliate_id: "your-affiliate-id"   # 省略可
  access_key: "pk_..."                # 省略可（Ichiba OpenAPI 用）
  origin: "https://www.example.com"   # 省略可（許可されたWebサイト）
```

環境変数でも設定できます。

```bash
export RAKUTEN_APP_ID="your-app-id"
export RAKUTEN_AFFILIATE_ID="your-affiliate-id"
export RAKUTEN_ACCESS_KEY="pk_..."
export RAKUTEN_ORIGIN="https://www.example.com"
```

## 使い方

```bash
# グローバルフラグ
--pretty    # JSON を整形して出力
--quiet     # エラーログを非表示
--timeout   # タイムアウト（デフォルト: 30s）
```

### 市場（Ichiba）

```bash
raku-cli ichiba items --keyword "ノートパソコン" --pretty
raku-cli ichiba items --keyword "本" --min-price 1000 --max-price 5000 --hits 10
raku-cli ichiba genre --genre-id 0 --pretty
raku-cli ichiba ranking --genre-id 555086 --pretty
```

### 書籍（Books）

```bash
raku-cli books search --keyword "golang" --pretty
raku-cli books book --keyword "村上春樹" --pretty
raku-cli books cd --keyword "米津玄師" --pretty
raku-cli books dvd --keyword "ジブリ" --pretty
raku-cli books magazine --keyword "週刊" --pretty
raku-cli books game --keyword "ポケモン" --pretty
raku-cli books genre --genre-id 001 --pretty
```

### トラベル（Travel）

```bash
raku-cli travel hotels --large-area japan --middle-area hokkaido --small-area sapporo --detail-area A --pretty
raku-cli travel hotel --hotel-no 901 --pretty
raku-cli travel vacant --large-area japan --middle-area tokyo --small-area tokyo --detail-area A --checkin-date 2026-04-01 --checkout-date 2026-04-02 --adult-num 2 --pretty
raku-cli travel area --pretty
raku-cli travel ranking --genre onsen --pretty
```

### その他（Misc）

```bash
raku-cli misc recipe --pretty
raku-cli misc kobo --keyword "漫画" --pretty
raku-cli misc gora --keyword "箱根" --pretty
```

## 終了コード

| コード | 意味 |
|--------|------|
| 0 | 正常終了 |
| 1 | 一般エラー |
| 2 | 入力エラー |
| 3 | 認証エラー |
| 4 | ネットワークエラー |

## ライセンス

MIT
