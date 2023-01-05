// Copyright 2022-2023 The sacloud/object-storage-api-go authors
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

// ListPermissions パーミッション一覧の取得
// (GET /{site_name}/v2/permissions)
func (engine *Engine) ListPermissions(siteId string) ([]v1.Permission, error) {
	defer engine.rLock()()

	if err := engine.siteExist(siteId); err != nil {
		return nil, err
	}
	return engine.permissions(), nil
}

// CreatePermission パーミッションの作成
// (POST /{site_name}/v2/permissions)
func (engine *Engine) CreatePermission(siteId string, params *v1.PermissionRequestBody) (*v1.Permission, error) {
	defer engine.lock()()

	if err := engine.siteExist(siteId); err != nil {
		return nil, err
	}

	permission := &v1.Permission{
		BucketControls: params.BucketControls,
		CreatedAt:      v1.CreatedAt(time.Now()),
		DisplayName:    params.DisplayName,
		Id:             v1.PermissionID(engine.nextId()),
	}

	engine.Permissions = append(engine.Permissions, permission)
	return engine.copyPermission(permission)
}

// DeletePermission パーミッションの削除
// (DELETE /{site_name}/v2/permissions/{id})
func (engine *Engine) DeletePermission(siteId string, permissionId int64) error {
	defer engine.lock()()

	if err := engine.siteAndPermissionExist(siteId, permissionId); err != nil {
		return err
	}

	var deleted []*v1.Permission
	for _, p := range engine.Permissions {
		if p.Id.Int64() != permissionId {
			deleted = append(deleted, p)
		}
	}
	engine.Permissions = deleted
	return nil
}

// ReadPermission パーミッションの取得
// (GET /{site_name}/v2/permissions/{id})
func (engine *Engine) ReadPermission(siteId string, permissionId int64) (*v1.Permission, error) {
	defer engine.rLock()()

	if err := engine.siteAndPermissionExist(siteId, permissionId); err != nil {
		return nil, err
	}

	return engine.copyPermission(engine.getPermissionById(permissionId))
}

// UpdatePermission パーミッションの更新
// (PUT /{site_name}/v2/permissions/{id})
func (engine *Engine) UpdatePermission(siteId string, permissionId int64, params *v1.PermissionRequestBody) (*v1.Permission, error) {
	defer engine.lock()()

	if err := engine.siteAndPermissionExist(siteId, permissionId); err != nil {
		return nil, err
	}

	var updated *v1.Permission
	for _, p := range engine.Permissions {
		if p.Id.Int64() == permissionId {
			p.BucketControls = params.BucketControls
			p.DisplayName = params.DisplayName
			updated = p
		}
	}

	return engine.copyPermission(updated) // チェック済みなのでupdatedはnilになり得ない
}

// permissions engine.Permissionsを非ポインタ型にして返す
func (engine *Engine) permissions() []v1.Permission {
	var permissions []v1.Permission
	for _, p := range engine.Permissions {
		permissions = append(permissions, *p)
	}
	return permissions
}

func (engine *Engine) getPermissionById(permissionId int64) *v1.Permission {
	if permissionId == 0 {
		return nil
	}
	for _, p := range engine.Permissions {
		if p.Id.Int64() == permissionId {
			return p
		}
	}
	return nil
}

func (engine *Engine) copyPermission(source *v1.Permission) (*v1.Permission, error) {
	if source == nil {
		return nil, fmt.Errorf("source is nil")
	}
	var permission v1.Permission
	if err := deepcopy.Copy(&permission, source); err != nil {
		return nil, err
	}
	return &permission, nil
}

func (engine *Engine) siteAndPermissionExist(siteId string, permissionId int64) error {
	if err := engine.siteExist(siteId); err != nil {
		return err
	}
	// Note: API定義上は定義されていないがサイトがないケースでは404が返される
	if permission := engine.getPermissionById(permissionId); permission == nil {
		return newError(ErrorTypeNotFound, "permission", "",
			"指定のパーミッションは存在しません。site_id: %s, id: %d", siteId, permissionId)
	}
	return nil
}
