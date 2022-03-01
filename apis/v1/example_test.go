// Copyright 2021-2022 The phy-go authors
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
