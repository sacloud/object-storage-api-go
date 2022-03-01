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

// DeleteSiteAccount サイトアカウントの削除
// (DELETE /{site_name}/v2/account)
func (s *Server) DeleteSiteAccount(c *gin.Context, siteName string) {
	if err := s.Engine.DeleteSiteAccount(siteName); err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// ReadSiteAccount サイトアカウントの取得
// (GET /{site_name}/v2/account)
func (s *Server) ReadSiteAccount(c *gin.Context, siteName string) {
	account, err := s.Engine.ReadSiteAccount(siteName)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &v1.AccountResponseBody{
		Data: *account,
	})
}

// CreateSiteAccount サイトアカウントの作成
// (POST /{site_name}/v2/account)
func (s *Server) CreateSiteAccount(c *gin.Context, siteName string) {
	account, err := s.Engine.CreateSiteAccount(siteName)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, &v1.AccountResponseBody{
		Data: *account,
	})
}
