// Copyright 2022-2025 The sacloud/object-storage-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package objectstorage_test

import (
	"context"
	"fmt"

	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

const defaultServerURL = "https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0"

var serverURL = defaultServerURL

// Example_clusterAPI サイト(クラスタ)APIの利用例
func Example_clusterAPI() {
	client := &objectstorage.Client{
		APIRootURL: serverURL, // 省略可能
	}

	// サイト一覧の取得
	siteOp := objectstorage.NewSiteOp(client)
	sites, err := siteOp.List(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(sites[0].DisplayName)
	// output:
	// 石狩第1サイト
}

// Example_bucketAPI バケットAPIの利用例
func Example_bucketAPI() {
	client := &objectstorage.Client{
		APIRootURL: serverURL, // 省略可能
	}
	ctx := context.Background()

	// サイトIDが必要なため取得しておく
	siteOp := objectstorage.NewSiteOp(client)
	sites, err := siteOp.List(ctx)
	if err != nil {
		panic(err)
	}
	siteId := sites[0].Id

	// バケットの作成
	bucketName := "your-bucket-name"
	bucketAPI := objectstorage.NewBucketOp(client)
	bucket, err := bucketAPI.Create(ctx, siteId, bucketName)
	if err != nil {
		panic(err)
	}

	// バケットの削除
	defer func() {
		if err := bucketAPI.Delete(ctx, siteId, bucketName); err != nil {
			panic(err)
		}
	}()

	fmt.Println(bucket.Name)
	// output:
	// your-bucket-name
}

// Example_accountAPI アカウントAPIの利用例
func Example_accountAPI() {
	client := &objectstorage.Client{
		APIRootURL: serverURL, // 省略可能
	}
	ctx := context.Background()

	// サイトIDが必要なため取得しておく
	siteOp := objectstorage.NewSiteOp(client)
	sites, err := siteOp.List(ctx)
	if err != nil {
		panic(err)
	}
	siteId := sites[0].Id

	// アカウントの参照
	accountOp := objectstorage.NewAccountOp(client)
	account, err := accountOp.Read(ctx, siteId)
	if err != nil {
		panic(err)
	}

	// アクセスキーの作成
	accessKey, err := accountOp.CreateAccessKey(ctx, siteId)
	if err != nil {
		panic(err)
	}

	// アクセスキーの一覧
	accessKeys, err := accountOp.ListAccessKeys(ctx, siteId)
	if err != nil {
		panic(err)
	}
	if len(accessKeys) == 0 {
		panic("ListAccessKeys failed")
	}

	// アクセスキーの削除
	defer func() {
		if err := accountOp.DeleteAccessKey(ctx, siteId, accessKey.Id.String()); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("AccountCode: %s, Secret: %s", account.Code, accessKey.Secret)
	// output:
	// AccountCode: member@account@isk01, Secret: secret
}

// Example_permissionAPI パーミッションAPIの利用例
func Example_permissionAPI() {
	client := &objectstorage.Client{
		APIRootURL: serverURL, // 省略可能
	}
	ctx := context.Background()

	// サイトIDが必要なため取得しておく
	siteOp := objectstorage.NewSiteOp(client)
	sites, err := siteOp.List(ctx)
	if err != nil {
		panic(err)
	}
	siteId := sites[0].Id

	// パーミッションの作成
	permissionOp := objectstorage.NewPermissionOp(client)
	permission, err := permissionOp.Create(ctx, siteId, &v1.CreatePermissionParams{
		BucketControls: v1.BucketControls{
			{
				BucketName: "bucket1",
				CanRead:    true,
				CanWrite:   true,
			},
		},
		DisplayName: "foobar",
	})
	if err != nil {
		panic(err)
	}

	// アクセスキーの作成
	accessKey, err := permissionOp.CreateAccessKey(ctx, siteId, permission.Id.Int64())
	if err != nil {
		panic(err)
	}

	// パーミッション/アクセスキーの削除
	defer func() {
		if err := permissionOp.DeleteAccessKey(ctx, siteId, permission.Id.Int64(), accessKey.Id.String()); err != nil {
			panic(err)
		}
		if err := permissionOp.Delete(ctx, siteId, permission.Id.Int64()); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Permission: %s, Secret: %s", permission.DisplayName, accessKey.Secret)
	// output:
	// Permission: foobar, Secret: secret
}

// Example_siteStatusAPI サイトステータスAPIの利用例
func Example_siteStatusAPI() {
	client := &objectstorage.Client{
		APIRootURL: serverURL, // 省略可能
	}
	ctx := context.Background()

	// サイトIDが必要なため取得しておく
	siteOp := objectstorage.NewSiteOp(client)
	sites, err := siteOp.List(ctx)
	if err != nil {
		panic(err)
	}
	siteId := sites[0].Id

	// サイトステータスの参照
	statusOp := objectstorage.NewSiteStatusOp(client)
	status, err := statusOp.Read(ctx, siteId)
	if err != nil {
		panic(err)
	}

	fmt.Println(status.StatusCode.Status)
	// output:
	// ok
}
