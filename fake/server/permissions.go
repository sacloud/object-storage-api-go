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

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// ListPermissions パーミッション一覧の取得
// (GET /{site_name}/v2/permissions)
func (s *Server) GetPermissions(c *gin.Context, siteId string) {
	permissions, err := s.Engine.ListPermissions(siteId)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &v1.PermissionsResponseBody{
		Data: permissions,
	})
}

// CreatePermission パーミッションの作成
// (POST /{site_name}/v2/permissions)
func (s *Server) CreatePermission(c *gin.Context, siteId string) {
	var paramJSON v1.PermissionRequestBody
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission, err := s.Engine.CreatePermission(siteId, &paramJSON)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, &v1.PermissionResponseBody{
		Data: *permission,
	})
}

// DeletePermission パーミッションの削除
// (DELETE /{site_name}/v2/permissions/{id})
func (s *Server) DeletePermission(c *gin.Context, siteId string, permissionId v1.PermissionID) {
	if err := s.Engine.DeletePermission(siteId, permissionId.Int64()); err != nil {
		s.handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GetPermission パーミッションの取得
// (GET /{site_name}/v2/permissions/{id})
func (s *Server) GetPermission(c *gin.Context, siteId string, permissionId v1.PermissionID) {
	permission, err := s.Engine.ReadPermission(siteId, permissionId.Int64())
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.PermissionResponseBody{
		Data: *permission,
	})
}

// UpdatePermission パーミッションの更新
// (PUT /{site_name}/v2/permissions/{id})
func (s *Server) UpdatePermission(c *gin.Context, siteId string, permissionId v1.PermissionID) {
	var paramJSON v1.PermissionRequestBody
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission, err := s.Engine.UpdatePermission(siteId, permissionId.Int64(), &paramJSON)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.PermissionResponseBody{
		Data: *permission,
	})
}
