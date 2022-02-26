# object-storage-api-go

- URL: https://github.com/sacloud/object-storage-api-go/pull/1
- Parent: https://github.com/sacloud/sacloud-go/pull/1
- Author: @yamamoto-febc

## 概要

[さくらのクラウド オブジェクトストレージAPI](https://manual.sakura.ad.jp/cloud/objectstorage/api.html)をGo言語から利用するためのライブラリを提供する。

さくらのクラウドオブジェクトストレージAPIは専用サーバPHYと同じくOpenAPI v3.0でのAPI定義が公開されている。

- [さくらのクラウド オブジェクトストレージAPI OpenAPIでの定義](https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.html)

このAPI定義を用いてGo言語向けのコードを生成する。  
またより簡単に使うためのラップしたコードやテスト用モックサーバも提供する。

## やること/やらないこと

### やること

- OpenAPIでのAPI定義から生成したコードの提供
- 生成したコードをラップするコードの提供
- テスト用モックサーバの提供
- テスト用モックサーバを操作するためのCLIの提供

### やらないこと

- Amazon S3互換APIで出来る範囲の機能の実装
  作成済みバケット/オブジェクトの操作といったS3互換APIで出来る機能は対象外とする。

## 実装

- 基本的には[sacloud/phy-go](https://github.com/sacloud/phy-go)の実装を踏襲する
  - コード生成には[oapi-codegen](https://github.com/deepmap/oapi-codegen)を用いる
  - テスト用モックサーバには[gin](https://github.com/gin-gonic/gin)を用いる

## 改訂履歴

- 2022/2/26: 初版作成