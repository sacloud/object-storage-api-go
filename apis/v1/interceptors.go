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

package v1

import (
	"context"
	"net/http"
)

// OjsAuthInterceptor オブジェクトストレージAPIリクエストに認証情報の注入を行う
func OjsAuthInterceptor(token, secret string) func(context.Context, *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(token, secret)
		return nil
	}
}
