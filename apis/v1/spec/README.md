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
  - 修正前:`^[\\w\\d-_]+$`
  - 修正後:`^[\\w\\d_-]+$`
