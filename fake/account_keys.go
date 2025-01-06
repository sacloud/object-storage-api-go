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

package fake

import (
	"fmt"
	"time"

	"github.com/getlantern/deepcopy"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// ListAccountAccessKeys サイトアカウントのアクセスキーの取得
// (GET /{site_name}/v2/account/keys)
func (engine *Engine) ListAccountAccessKeys(siteId string) ([]v1.AccountKey, error) {
	defer engine.rLock()()

	if err := engine.siteAndAccountExist(siteId); err != nil {
		return nil, err
	}
	return engine.accountKeys(), nil
}

// CreateAccountAccessKey サイトアカウントのアクセスキーの発行
// (POST /{site_name}/v2/account/keys)
func (engine *Engine) CreateAccountAccessKey(siteId string) (*v1.AccountKey, error) {
	defer engine.lock()()
	if err := engine.siteAndAccountExist(siteId); err != nil {
		return nil, err
	}

	// Note: 本来はサイトに紐づくアカウントキーが存在する場合はエラーにすべきだが、
	//       fakeではサイトごとにデータが分離されていないため未チェックとなっている。
	//       サイトが実質1つなので問題はないと思われる。今後サイトが増えるようであれば実装を検討する。

	key := &v1.AccountKey{
		CreatedAt: v1.CreatedAt(time.Now()),
		Id:        v1.AccessKeyID(fmt.Sprintf("%d", engine.nextId())),
		Secret:    "secret", // fakeでは固定値を返す
	}

	engine.AccountKeys = append(engine.AccountKeys, key)
	return engine.copyAccountKey(key)
}

// DeleteAccountAccessKey サイトアカウントのアクセスキーの削除
// (DELETE /{site_name}/v2/account/keys/{id})
func (engine *Engine) DeleteAccountAccessKey(siteId string, id string) error {
	defer engine.lock()()

	if err := engine.siteAndAccountExist(siteId); err != nil {
		return err
	}

	key := engine.getAccountKeyById(id)
	if key != nil {
		var keys []*v1.AccountKey
		for _, k := range engine.AccountKeys {
			if k.Id.String() != key.Id.String() {
				keys = append(keys, k)
			}
		}
		engine.AccountKeys = keys
		return nil
	}
	return newError(ErrorTypeNotFound, "account_key", id, "アカウントキーが存在しません。id: %s", id)
}

// ReadAccountAccessKey サイトアカウントのアクセスキーの取得
// (GET /{site_name}/v2/account/keys/{id})
func (engine *Engine) ReadAccountAccessKey(siteId string, id string) (*v1.AccountKey, error) {
	defer engine.rLock()()

	if err := engine.siteAndAccountExist(siteId); err != nil {
		return nil, err
	}

	key := engine.getAccountKeyById(id)
	if key != nil {
		k, err := engine.copyAccountKey(key)
		if err != nil {
			return nil, err
		}
		k.Secret = "" // 新規作成時のみ参照できる項目
		return k, nil
	}
	return nil, newError(ErrorTypeNotFound, "account_key", id, "アカウントキーが存在しません。id: %s", id)
}

// accountKeys engine.AccountKeysを非ポインタ型にして返す
func (engine *Engine) accountKeys() []v1.AccountKey {
	var keys []v1.AccountKey
	for _, k := range engine.AccountKeys {
		key := *k
		key.Secret = "" // 新規作成時のみ参照できる項目
		keys = append(keys, key)
	}
	return keys
}

func (engine *Engine) copyAccountKey(source *v1.AccountKey) (*v1.AccountKey, error) {
	if source == nil {
		return nil, fmt.Errorf("source is nil")
	}
	var key v1.AccountKey
	if err := deepcopy.Copy(&key, source); err != nil {
		return nil, err
	}
	return &key, nil
}

func (engine *Engine) getAccountKeyById(id string) *v1.AccountKey {
	if id == "" {
		return nil
	}
	for _, k := range engine.AccountKeys {
		if k.Id.String() == id {
			return k
		}
	}
	return nil
}
