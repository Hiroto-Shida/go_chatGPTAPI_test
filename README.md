## golangでchatGPTAPIを叩くプログラム

> 主にこちらの記事を参考に作成
> https://qiita.com/bluewolfali/items/9031cd56f1743aec3e9a

### ファイル役割
- /main.go
  - メインの実行ファイル
- mypkg/type.go
  - 型定義したパッケージファイル

### プログラム内容(ざっくり) 
※基本的にはプログラム内にコメントアウトで書いてあるのでそちらを参照
- chatGPTAPIに対するユーザの質問入力，chatGPTAPIからの返答の取得
- 本プログラムでは質問と返答のサイクルを2回だけ繰り返すようにしている
- system_content 変数に返答するChatGPTの設定
  - 本プログラムでは語尾の指定や，返答する文章の制限などを軽く設定している
  - よりよい参考例
    > https://qiita.com/sakasegawa/items/db2cff79bd14faf2c8e0
- .envファイルにAPI_KEYを保存し，Goの「godotenv」パッケージを使ってAPI_KEYを読み込み．
- あとはgoによるapiを叩いているだけ

commit test
