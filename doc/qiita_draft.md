# 楽天市場がごちゃごちゃで探しにくいのでAIエージェントに検索させるCLIツールを作った

**タグ:** `Go` `CLI` `個人開発` `AIエージェント` `ポイ活`

---

楽天市場、好きですか？

僕は好きです。ポイントはたまるし、品揃えは豊富だし。

でも**サイトがごちゃごちゃしすぎて探しにくい**。広告は多い、バナーは派手、気づいたら全然関係ないページにいる。

そこで、楽天ウェブサービスAPIをCLIから叩けるツール「**raku-cli**」を作りました。

出力はすべてJSON。AIエージェントに渡して「いい感じに探しておいて」が実現できます。

## リポジトリ

https://github.com/dugabot0/raku-cli

## インストール

[Releases](https://github.com/dugabot0/raku-cli/releases/tag/v0.1.0) からバイナリを落とすだけです。

```bash
# macOS (Apple Silicon)
curl -L https://github.com/dugabot0/raku-cli/releases/download/v0.1.0/raku-cli_darwin_arm64 -o raku-cli
chmod +x raku-cli
```

ソースからビルドする場合:

```bash
git clone https://github.com/dugabot0/raku-cli.git
cd raku-cli
go build -o raku-cli .
```

## 設定

`~/.config/raku-cli/config.yaml` を作成します。

```yaml
rakuten:
  app_id: "your-app-id"
  affiliate_id: "your-affiliate-id"   # アフィリエイトIDがあれば
  access_key: "pk_..."                # Ichiba OpenAPI用（任意）
  origin: "https://www.example.com"   # 許可されたWebサイト
```

[楽天ウェブサービス](https://webservice.rakuten.co.jp/)でアプリIDを取得してください。無料です。

環境変数でも設定できます:

```bash
export RAKUTEN_APP_ID="your-app-id"
export RAKUTEN_AFFILIATE_ID="your-affiliate-id"
export RAKUTEN_ACCESS_KEY="pk_..."
export RAKUTEN_ORIGIN="https://www.example.com"
```

## 対応API

| コマンド | 内容 |
|---|---|
| `ichiba items` | 市場アイテム検索 |
| `ichiba ranking` | ランキング |
| `ichiba genre` | ジャンル一覧 |
| `books search` | 書籍・CD・DVD・ゲーム横断検索 |
| `books book/cd/dvd/game/magazine` | メディア別検索 |
| `travel hotels` | ホテル検索 |
| `travel vacant` | 空室検索 |
| `travel ranking` | ホテルランキング |
| `misc kobo` | Kobo電子書籍検索 |
| `misc gora` | GORAゴルフ場検索 |
| `misc recipe` | レシピカテゴリ |

## 使い方

### 基本

```bash
# キーワード検索
raku-cli ichiba items --keyword "ノートパソコン" --pretty

# ランキングを見る
raku-cli ichiba ranking --pretty

# ホテルを探す
raku-cli travel hotels --large-area japan --middle-area tokyo --pretty
```

### AIエージェントに渡す

JSON出力なのでそのままAIエージェントに投げられます。

楽天のサイトで探すより「CLIで取得してAIに判断させる」ほうが圧倒的に楽です。

```
「Nintendo Switchのソフトで評価が高くて安いものを探して」
→ raku-cli books game --keyword "Nintendo Switch" --sort reviewAverage の結果をAIに渡す
```

サイトを何ページもスクロールする必要がなく、AIが整理して返してくれます。

### ポイ活に使う（おすすめ）

個人的に一番好きな使い方がこれです。

楽天ポイントの倍率が高い商品だけに絞る `--point-rate-flag` フラグがあります。

```bash
# ポイント倍率が設定されている商品のみ
raku-cli ichiba items --keyword "洗剤" --point-rate-flag 1 --pretty

# さらにポイント5倍以上に絞る
raku-cli ichiba items --keyword "洗剤" --point-rate-flag 1 --point-rate 5 --pretty
```

楽天お買い物マラソンやスーパーSALEのとき、ポイント倍率が高い商品を自動で拾ってきてAIに「この中で一番コスパいいのどれ？」と聞けばポイ活が捗ります。

送料無料と組み合わせることも:

```bash
raku-cli ichiba items --keyword "コーヒー" --postage-flag 1 --point-rate-flag 1 --sort "-itemPrice" --pretty
```

## フラグ一覧（ichiba items）

主なフラグを抜粋します。APIのパラメータをほぼ全部実装しています。

```
--keyword           検索キーワード
--sort              並び順（+itemPrice/-itemPrice/reviewCount/reviewAverage 等）
--min-price         最低価格
--max-price         最高価格
--point-rate-flag   1=ポイント倍率商品のみ
--point-rate        最低ポイント倍率（2〜10）
--postage-flag      1=送料無料のみ
--has-review-flag   1=レビューあり商品のみ
--image-flag        1=画像あり商品のみ
--or-flag           1=OR検索
--ng-keyword        除外キーワード
--hits              件数（最大30）
--page              ページ番号
--pretty            整形して出力
```

`books`、`travel`、`kobo`、`gora` も同様に詳細なフラグに対応しています。

## 実装について

コードはClaudeに書いてもらいました（前回のav-cliと同じパターン）。

- **言語**: Go
- **フレームワーク**: Cobra + Viper
- **出力**: すべてJSONでstdout、エラーはstderr
- **終了コード**: 0=正常、3=認証エラー、4=ネットワークエラー

APIレスポンスはそのままJSON出力するので、`jq` やAIエージェントと相性がいいです。

```bash
# jqと組み合わせてアイテム名と価格だけ抽出
raku-cli ichiba items --keyword "ヘッドフォン" --hits 5 | jq '.Items[] | {name: .itemName, price: .itemPrice}'
```

## まとめ

- 楽天のサイトはごちゃごちゃしてて探しにくい
- CLIで取ってきてAIに任せると快適
- `--point-rate-flag 1` でポイ活が捗る

ぜひ使ってみてください。
