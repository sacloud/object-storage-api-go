## API定義(swagger.yaml)について

オリジナルの定義ファイルは以下のサイトで公開されています。
[https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.html](https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.html)

公開されている定義ファイルのままではlintでエラーになる箇所があるため、手作業で修正しています。
修正は以下のように行っています。

- オリジナルの定義ファイルをダウンロード、`original-swagger.json`として保存
- `make gen`を実行することで`original-swagger.json`から`original-swagger.yaml`へ変換
- `original-swagger.yaml`をコピー/編集し`swagger.yaml`を作成

`original-swagger.yaml`については生成される対象なため`.gitignore`に登録されています。
今後オリジナルの定義ファイルが更新された場合は`original-swagger.yaml`と`swagger.yaml`のdiffを取り、適宜`swagger.yaml`へ反映するようにします。

### オリジナルの定義ファイルにおける既知の問題点

- `components.securitySchemes.Account_api_key`でのタイプミス 
    - 修正前: `schema`
    - 修正後: `scheme`
- `components.securitySchemes.Account_api_key.type`でのタイプミス
    - 修正前: `HTTP`
    - 修正後: `http`
- 正規表現パターンの指定誤り:
  - 修正前:`^[\w\d-_]+$`
  - 修正後:`^[\w\d_-]+$`
  :bulb: `-`を指定する時は最初か最後に書く必要がある
- `example`が`pettern`で指定した正規表現にマッチしない
  - `components.schemas.Code`: pattern=`^\w+$`, examples=`abc01234@foo@isk01` (patternの誤り)
    - 修正前:`^\w+$`
    - 修正後:`^[\w@]+$`
  - `components.schemas.DisplayName`: pattern=`^\w+$`, examples=`abc012345-` (patternの誤り)
    - 修正前:`^\w+$`
    - 修正後:`^.+$`
    :warning: 入力パターンが不明なため`.`にしているが、API側で制御しているのであれば適切に指定するのが望ましい
  - `components.schemas.AccessKeyID`: pattern=`^[\w\d\/]{40}$`, examples=`abcdefABCDEF0123456789` (patternの誤り)
    - 修正前:`^[\w\d\/]{40}$`
    - 修正後:`^[\w\d\/]{1,40}$`
    :warning: 数量詞誤りと思われるが、出現パターンが不明なため広く設定している
  - `components.schemas.SecretAccessKey`: pattern=`^[\w\d\/]{40}$`, examples=`NOTICE: EXISTS ONLY WHEN JUST CREATED` 
    - 修正前: pattern=`^[\w\d\/]{40}$`, `examples=`NOTICE: EXISTS ONLY WHEN JUST CREATED`
    - 修正後: pattern=`^[\w\d\/=]{40}$`, `examples=`==NOTICE==EXISTS/ONLY/WHEN/JUST/CREATED=`
      :warning: パターン不明なため実際の値を数パターンみて判定
  - `components.schemas.PermissionSecret`: pattern=`^[\w\d\/]{40}$`, examples=`NOTICE: EXISTS ONLY WHEN JUST CREATED`
    - 修正前: pattern=`^[\w\d\/]{40}$`, `examples=`NOTICE: EXISTS ONLY WHEN JUST CREATED`
    - 修正後: pattern=`^[\w\d\/=]{40}$`, `examples=`==NOTICE==EXISTS/ONLY/WHEN/JUST/CREATED=`
      `components.schemas.SecretAccessKey`と同じ
  - `components.schemas.Session.access_key_secret`: pattern=`^[\w\d\/]{40}$`, examples=`abcdefABCDEF0123456789` (patternの誤り)
    - 修正前:`^[\w\d\/]{40}$`
    - 修正後:`^[\w\d\/]{1,40}$`
      `components.schemas.AccessKeyID`と同じ
