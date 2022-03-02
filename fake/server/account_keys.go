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

// ListAccountAccessKeys サイトアカウントのアクセスキーの取得
// (GET /{site_name}/v2/account/keys)
func (s *Server) ListAccountAccessKeys(c *gin.Context, siteName string) {
	keys, err := s.Engine.ListAccountAccessKeys(siteName)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &v1.AccountKeysResponseBody{
		Data: keys,
	})
}

// CreateAccountAccessKey サイトアカウントのアクセスキーの発行
// (POST /{site_name}/v2/account/keys)
func (s *Server) CreateAccountAccessKey(c *gin.Context, siteName string) {
	key, err := s.Engine.CreateAccountAccessKey(siteName)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, &v1.AccountKeyResponseBody{
		Data: *key,
	})
}

// DeleteAccountAccessKey サイトアカウントのアクセスキーの削除
// (DELETE /{site_name}/v2/account/keys/{id})
func (s *Server) DeleteAccountAccessKey(c *gin.Context, siteName string, id string) {
	if err := s.Engine.DeleteAccountAccessKey(siteName, id); err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// ReadAccountAccessKey サイトアカウントのアクセスキーの取得
// (GET /{site_name}/v2/account/keys/{id})
func (s *Server) ReadAccountAccessKey(c *gin.Context, siteName string, id string) {
	key, err := s.Engine.ReadAccountAccessKey(siteName, id)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &v1.AccountKeyResponseBody{
		Data: *key,
	})
}