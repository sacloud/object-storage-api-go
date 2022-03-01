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

import "github.com/gin-gonic/gin"

// ListPermissions パーミッション一覧の取得
// (GET /{site_name}/v2/permissions)
func (s *Server) ListPermissions(c *gin.Context, siteName string) {}

// CreatePermission パーミッションの作成
// (POST /{site_name}/v2/permissions)
func (s *Server) CreatePermission(c *gin.Context, siteName string) {}

// DeletePermission パーミッションの削除
// (DELETE /{site_name}/v2/permissions/{id})
func (s *Server) DeletePermission(c *gin.Context, siteName string, id string) {}

// ReadPermission パーミッションの取得
// (GET /{site_name}/v2/permissions/{id})
func (s *Server) ReadPermission(c *gin.Context, siteName string, id string) {}

// UpdatePermission パーミッションの更新
// (PUT /{site_name}/v2/permissions/{id})
func (s *Server) UpdatePermission(c *gin.Context, siteName string, id string) {}

// ListPermissionAccessKeys パーミッションが保有するアクセスキー一覧の取得
// (GET /{site_name}/v2/permissions/{id}/keys)
func (s *Server) ListPermissionAccessKeys(c *gin.Context, siteName string, id string) {}

// CreatePermissionAccessKey パーミッションのアクセスキーの発行
// (POST /{site_name}/v2/permissions/{id}/keys)
func (s *Server) CreatePermissionAccessKey(c *gin.Context, siteName string, id string) {}

// DeletePermissionAccessKey パーミッションが保有するアクセスキーの削除
// (DELETE /{site_name}/v2/permissions/{id}/keys/{key_id})
func (s *Server) DeletePermissionAccessKey(c *gin.Context, siteName string, id string, keyId string) {
}

// ReadPermissionAccessKey パーミッションが保有するアクセスキーの取得
// (GET /{site_name}/v2/permissions/{id}/keys/{key_id})
func (s *Server) ReadPermissionAccessKey(c *gin.Context, siteName string, id string, keyId string) {}

// ReadSiteStatus サイトのステータスの取得
// (GET /{site_name}/v2/status)
func (s *Server) ReadSiteStatus(c *gin.Context, siteName string) {}
