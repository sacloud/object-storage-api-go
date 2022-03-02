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

package fake

import (
	"sync"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// Engine Fakeサーバであつかうダミーデータを表す
//
// Serverに渡した後はDataStore内のデータを外部から操作しないこと
// Note: 本来はサイトごとに保持するデータを分離すべきだが、現状だと実質1サイトのみのため分離していない。
type Engine struct {
	// Clusters サイト(クラスター)
	Clusters []*v1.Cluster

	// Buckets バケット
	Buckets []*v1.Bucket

	// Account アカウント
	Account *v1.Account

	// AccountKeys アカウントのアクセスキー
	AccountKeys []*v1.AccountKey

	// Permissions パーミッション
	Permissions []*v1.Permission

	PermissionKeys []*v1.PermissionKey

	// GeneratedID 採番済みの最終ID
	//
	// DataStoreの各フィールドの値との整合性は確認されないため利用者側が管理する必要がある
	GeneratedID int

	mu sync.RWMutex
}

func (engine *Engine) lock() func() {
	engine.mu.Lock()
	return engine.mu.Unlock
}

func (engine *Engine) rLock() func() {
	engine.mu.RLock()
	return engine.mu.RUnlock
}

// nextId GeneratedIDを+1したものを返す
//
// ロックは行わないため呼び出し側で適切に制御すること
func (engine *Engine) nextId() int {
	engine.GeneratedID++
	id := engine.GeneratedID
	return id
}
