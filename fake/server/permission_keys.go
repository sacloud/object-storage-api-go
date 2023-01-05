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

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// ListPermissionAccessKeys パーミッションが保有するアクセスキー一覧の取得
// (GET /{site_name}/v2/permissions/{id}/keys)
func (s *Server) GetPermissionKeys(c *gin.Context, siteId string, permissionId v1.PermissionID) {
	keys, err := s.Engine.ListPermissionAccessKeys(siteId, permissionId.Int64())
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &v1.PermissionKeysResponseBody{
		Data: keys,
	})
}

// CreatePermissionAccessKey パーミッションのアクセスキーの発行
// (POST /{site_name}/v2/permissions/{id}/keys)
func (s *Server) CreatePermissionKey(c *gin.Context, siteId string, permissionId v1.PermissionID) {
	key, err := s.Engine.CreatePermissionAccessKey(siteId, permissionId.Int64())
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, &v1.PermissionKeyResponseBody{
		Data: *key,
	})
}

// DeletePermissionAccessKey パーミッションが保有するアクセスキーの削除
// (DELETE /{site_name}/v2/permissions/{id}/keys/{key_id})
func (s *Server) DeletePermissionKey(c *gin.Context, siteId string, permissionId v1.PermissionID, permissionKeyId v1.AccessKeyID) {
	if err := s.Engine.DeletePermissionAccessKey(siteId, permissionId.Int64(), permissionKeyId.String()); err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// ReadPermissionAccessKey パーミッションが保有するアクセスキーの取得
// (GET /{site_name}/v2/permissions/{id}/keys/{key_id})
func (s *Server) GetPermissionKey(c *gin.Context, siteId string, permissionId v1.PermissionID, permissionKeyId v1.AccessKeyID) {
	key, err := s.Engine.ReadPermissionAccessKey(siteId, permissionId.Int64(), permissionKeyId.String())
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &v1.PermissionKeyResponseBody{
		Data: *key,
	})
}
