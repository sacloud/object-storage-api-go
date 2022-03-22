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

package objectstorage

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
	sacloudhttp "github.com/sacloud/go-http"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// DefaultAPIRootURL デフォルトのAPIルートURL
const DefaultAPIRootURL = "https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0"

// UserAgent APIリクエスト時のユーザーエージェント
var UserAgent = fmt.Sprintf(
	"object-storage-api-go/%s (%s/%s; +https://github.com/sacloud/object-storage-api-go) %s",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
	sacloudhttp.DefaultUserAgent,
)

// Client APIクライアント
type Client struct {
	// Token APIキー: トークン
	Token string
	// Token APIキー: シークレット
	Secret string

	// AcceptLanguage APIリクエスト時のAccept-Languageヘッダーの値
	AcceptLanguage string

	// Gzip APIリクエストでgzipを有効にするかのフラグ
	Gzip bool

	// APIRootURL APIのリクエスト先URLプレフィックス、省略可能
	APIRootURL string

	// Trace トレースログ出力フラグ
	Trace bool

	// HTTPClient APIリクエストで使用されるHTTPクライアント
	//
	// 省略した場合はhttp.DefaultClientが使用される
	HTTPClient *http.Client

	initOnce sync.Once
}

func (c *Client) serverURL() string {
	v := DefaultAPIRootURL
	if c.APIRootURL != "" {
		v = c.APIRootURL
	}
	if !strings.HasSuffix(v, "/") {
		v += "/"
	}
	return v
}

func (c *Client) httpClient() *http.Client {
	client := http.DefaultClient
	if c.HTTPClient != nil {
		client = c.HTTPClient
	}

	c.initOnce.Do(func() {
		if c.Trace {
			client.Transport = &sacloudhttp.TracingRoundTripper{
				Transport: client.Transport,
			}
		}
	})
	return client
}

func (c *Client) apiClient() *v1.ClientWithResponses {
	httpClient := &sacloudhttp.Client{
		AccessToken:       c.Token,
		AccessTokenSecret: c.Secret,
		UserAgent:         UserAgent,
		AcceptLanguage:    c.AcceptLanguage,
		Gzip:              c.Gzip,
		CheckRetryFunc: func(ctx context.Context, resp *http.Response, err error) (bool, error) {
			if ctx.Err() != nil {
				return false, ctx.Err()
			}
			if err != nil {
				return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
			}
			if resp.StatusCode == 0 { // ステータスコードに応じてリトライしたい場合はここで対応
				return true, nil
			}
			return false, nil
		},
		HTTPClient: c.httpClient(),
	}
	return &v1.ClientWithResponses{
		ClientInterface: &v1.Client{
			Server: c.serverURL(),
			Client: httpClient,
		},
	}
}
