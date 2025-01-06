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

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

// ListClusters サイト一覧の取得
// (GET /fed/v1/clusters)
func (s *Server) GetClusters(c *gin.Context) {
	clusters, err := s.Engine.ListClusters()
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &v1.ListClustersResponseBody{
		Data: clusters,
	})
}

// ReadCluster サイトの取得
// (GET /fed/v1/clusters/{id})
func (s *Server) GetCluster(c *gin.Context, siteId string) {
	cluster, err := s.Engine.ReadCluster(siteId)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, &v1.ReadClusterResponseBody{
		Data: cluster,
	})
}
