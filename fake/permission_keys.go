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
	"fmt"
	"time"

	"github.com/getlantern/deepcopy"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// ListPermissionAccessKeys パーミッションが保有するアクセスキー一覧の取得
// (GET /{site_name}/v2/permissions/{id}/keys)
func (engine *Engine) ListPermissionAccessKeys(siteId string, permissionId int64) ([]v1.PermissionKey, error) {
	defer engine.rLock()()

	if err := engine.siteAndPermissionExist(siteId, permissionId); err != nil {
		return nil, err
	}
	return engine.permissionKeys(), nil
}

// CreatePermissionAccessKey パーミッションのアクセスキーの発行
// (POST /{site_name}/v2/permissions/{id}/keys)
func (engine *Engine) CreatePermissionAccessKey(siteId string, permissionId int64) (*v1.PermissionKey, error) {
	defer engine.lock()()

	if err := engine.siteAndPermissionExist(siteId, permissionId); err != nil {
		return nil, err
	}

	// Note: 本来はパーミッションに紐づくキーが存在する場合はエラーにすべきだが、fakeでは未チェック

	key := &v1.PermissionKey{
		CreatedAt: v1.CreatedAt(time.Now()),
		Id:        v1.AccessKeyID(fmt.Sprintf("%d", engine.nextId())),
		Secret:    "secret", // fakeでは固定値を返す
	}

	engine.PermissionKeys = append(engine.PermissionKeys, key)
	return engine.copyPermissionKey(key)
}

// DeletePermissionAccessKey パーミッションが保有するアクセスキーの削除
// (DELETE /{site_name}/v2/permissions/{id}/keys/{key_id})
func (engine *Engine) DeletePermissionAccessKey(siteId string, permissionId int64, permissionKeyId string) error {
	defer engine.lock()()

	if err := engine.siteAndPermissionExist(siteId, permissionId); err != nil {
		return err
	}

	var key *v1.PermissionKey
	if key = engine.getPermissionKeyById(permissionKeyId); key == nil {
		return NewError(ErrorTypeNotFound, "permission_key", permissionKeyId,
			"パーミッションキーが存在しません site_id: %s, permission_id: %d, permission_key_id: %s",
			siteId, permissionId, permissionKeyId)
	}

	var deleted []*v1.PermissionKey
	for _, k := range engine.PermissionKeys {
		if k.Id.String() != key.Id.String() {
			deleted = append(deleted, k)
		}
	}
	engine.PermissionKeys = deleted
	return nil
}

// ReadPermissionAccessKey パーミッションが保有するアクセスキーの取得
// (GET /{site_name}/v2/permissions/{id}/keys/{key_id})
func (engine *Engine) ReadPermissionAccessKey(siteId string, permissionId int64, permissionKeyId string) (*v1.PermissionKey, error) {
	defer engine.rLock()()

	if err := engine.siteAndPermissionExist(siteId, permissionId); err != nil {
		return nil, err
	}

	var key *v1.PermissionKey
	if key = engine.getPermissionKeyById(permissionKeyId); key == nil {
		return nil, NewError(ErrorTypeNotFound, "permission_key", permissionKeyId,
			"パーミッションキーが存在しません site_id: %s, permission_id: %d, permission_key_id: %s",
			siteId, permissionId, permissionKeyId)
	}
	k, err := engine.copyPermissionKey(key)
	if err != nil {
		return nil, err
	}
	k.Secret = "" // 新規作成時のみ参照できる項目
	return k, nil
}

// permissionKeys engine.PermissionKeysを非ポインタ型にして返す
func (engine *Engine) permissionKeys() []v1.PermissionKey {
	var keys []v1.PermissionKey
	for _, k := range engine.PermissionKeys {
		key := *k
		key.Secret = "" // 新規作成時のみ参照できる項目
		keys = append(keys, key)
	}
	return keys
}

func (engine *Engine) copyPermissionKey(source *v1.PermissionKey) (*v1.PermissionKey, error) {
	if source == nil {
		return nil, fmt.Errorf("source is nil")
	}
	var key v1.PermissionKey
	if err := deepcopy.Copy(&key, source); err != nil {
		return nil, err
	}
	return &key, nil
}

func (engine *Engine) getPermissionKeyById(permissionKeyId string) *v1.PermissionKey {
	if permissionKeyId == "" {
		return nil
	}
	for _, k := range engine.PermissionKeys {
		if k.Id.String() == permissionKeyId {
			return k
		}
	}
	return nil
}
