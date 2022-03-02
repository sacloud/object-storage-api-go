// Copyright 2022 The sacloud/object-storage-api-go authors
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

package v1_test

import (
	"context"
	"fmt"
	"os"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

var serverURL = "https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0"

// Example API定義から生成されたコードを直接利用する例
func Example_basic() {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	client, err := v1.NewClientWithResponses(serverURL, func(c *v1.Client) error {
		c.RequestEditors = []v1.RequestEditorFn{
			v1.OjsAuthInterceptor(token, secret),
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	resp, err := client.ListClustersWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	sites, err := resp.Result()
	if err != nil {
		panic(err)
	}

	site := sites.Data[0]
	fmt.Println(site.DisplayName)
	// output:
	// 石狩第1サイト
}

// Example_readSiteAccount サイトアカウントの参照の例
func Example_readSiteAccount() {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	client, err := v1.NewClientWithResponses(serverURL, func(c *v1.Client) error {
		c.RequestEditors = []v1.RequestEditorFn{
			v1.OjsAuthInterceptor(token, secret),
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// サイトIDが必要になるためまずサイト一覧を取得
	sitesResp, err := client.ListClustersWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	sites, err := sitesResp.Result()
	if err != nil {
		panic(err)
	}
	siteId := sites.Data[0].Id

	// サイトアカウントの参照
	accountResp, err := client.ReadSiteAccountWithResponse(context.Background(), siteId)
	if err != nil {
		panic(err)
	}

	account, err := accountResp.Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(account.Data.Code)
	// output:
	// member@account@isk01
}

// Example_siteAccountKeys サイトアカウントのキー操作の例
func Example_siteAccountKeys() {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	client, err := v1.NewClientWithResponses(serverURL, func(c *v1.Client) error {
		c.RequestEditors = []v1.RequestEditorFn{
			v1.OjsAuthInterceptor(token, secret),
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// サイトIDが必要になるためまずサイト一覧を取得
	sitesResp, err := client.ListClustersWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	sites, err := sitesResp.Result()
	if err != nil {
		panic(err)
	}
	siteId := sites.Data[0].Id

	// サイトアカウントのキーを作成
	keyResp, err := client.CreateAccountAccessKeyWithResponse(context.Background(), siteId)
	if err != nil {
		panic(err)
	}

	key, err := keyResp.Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(key.Data.Secret)
	// output:
	// secret
}

// Example_bucket バケット操作
func Example_bucket() {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	client, err := v1.NewClientWithResponses(serverURL, func(c *v1.Client) error {
		c.RequestEditors = []v1.RequestEditorFn{
			v1.OjsAuthInterceptor(token, secret),
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// サイトIDが必要になるためまずサイト一覧を取得
	sitesResp, err := client.ListClustersWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	sites, err := sitesResp.Result()
	if err != nil {
		panic(err)
	}
	siteId := sites.Data[0].Id

	// バケット作成
	createParams := v1.CreateBucketJSONRequestBody{
		ClusterId: siteId,
	}

	bucketResp, err := client.CreateBucketWithResponse(context.Background(), "bucket-name", createParams)
	if err != nil {
		panic(err)
	}

	bucket, err := bucketResp.Result()
	if err != nil {
		panic(err)
	}

	defer func() {
		deleteParams := v1.DeleteBucketJSONRequestBody{
			ClusterId: siteId,
		}
		resp, err := client.DeleteBucketWithResponse(context.Background(), "bucket-name", deleteParams)
		if err != nil {
			panic(err)
		}
		if err := resp.Result(); err != nil {
			panic(err)
		}
	}()

	fmt.Println(bucket.Data.Name)
	// output:
	// bucket-name
}

// Example_permissions パーミッション操作の例
func Example_permissions() {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	client, err := v1.NewClientWithResponses(serverURL, func(c *v1.Client) error {
		c.RequestEditors = []v1.RequestEditorFn{
			v1.OjsAuthInterceptor(token, secret),
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// サイトIDが必要になるためまずサイト一覧を取得
	sitesResp, err := client.ListClustersWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	sites, err := sitesResp.Result()
	if err != nil {
		panic(err)
	}
	siteId := sites.Data[0].Id

	// パーミッション作成
	permissionResp, err := client.CreatePermissionWithResponse(context.Background(), siteId, v1.CreatePermissionJSONRequestBody{
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

	permission, err := permissionResp.Result()
	if err != nil {
		panic(err)
	}

	defer func() {
		resp, err := client.DeletePermissionWithResponse(context.Background(), siteId, permission.Data.Id)
		if err != nil {
			panic(err)
		}
		if err := resp.Result(); err != nil {
			panic(err)
		}
	}()

	fmt.Println(permission.Data.DisplayName)
	// output:
	// foobar
}

// Example_permissionKeys パーミッションのキー操作の例
func Example_permissionKeys() {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	client, err := v1.NewClientWithResponses(serverURL, func(c *v1.Client) error {
		c.RequestEditors = []v1.RequestEditorFn{
			v1.OjsAuthInterceptor(token, secret),
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// サイトIDが必要になるためまずサイト一覧を取得
	sitesResp, err := client.ListClustersWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	sites, err := sitesResp.Result()
	if err != nil {
		panic(err)
	}
	siteId := sites.Data[0].Id

	// パーミッション作成
	permissionResp, err := client.CreatePermissionWithResponse(context.Background(), siteId, v1.CreatePermissionJSONRequestBody{
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

	permission, err := permissionResp.Result()
	if err != nil {
		panic(err)
	}

	// パーミッションのキーを作成
	keyResp, err := client.CreatePermissionAccessKeyWithResponse(context.Background(), siteId, permission.Data.Id)
	if err != nil {
		panic(err)
	}

	key, err := keyResp.Result()
	if err != nil {
		panic(err)
	}

	defer func() {
		keyDeleteResp, err := client.DeletePermissionAccessKeyWithResponse(context.Background(), siteId, permission.Data.Id, key.Data.Id)
		if err != nil {
			panic(err)
		}
		if err := keyDeleteResp.Result(); err != nil {
			panic(err)
		}

		permDeleteResp, err := client.DeletePermissionWithResponse(context.Background(), siteId, permission.Data.Id)
		if err != nil {
			panic(err)
		}
		if err := permDeleteResp.Result(); err != nil {
			panic(err)
		}
	}()

	fmt.Println(key.Data.Secret)
	// output:
	// secret
}

// Example_siteStatus サイトステータス確認の例
func Example_siteStatus() {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	client, err := v1.NewClientWithResponses(serverURL, func(c *v1.Client) error {
		c.RequestEditors = []v1.RequestEditorFn{
			v1.OjsAuthInterceptor(token, secret),
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// サイトIDが必要になるためまずサイト一覧を取得
	sitesResp, err := client.ListClustersWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	sites, err := sitesResp.Result()
	if err != nil {
		panic(err)
	}
	siteId := sites.Data[0].Id

	statusResp, err := client.ReadSiteStatusWithResponse(context.Background(), siteId)
	if err != nil {
		panic(err)
	}
	status, err := statusResp.Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(status.Data.StatusCode.Status)
	// output:
	// ok
}
