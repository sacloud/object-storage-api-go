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

// DeleteBucket バケットの削除
// (DELETE /fed/v1/buckets/{name})
func (s *Server) DeleteBucket(c *gin.Context, name string) {

}

// CreateBucket バケットの作成
// (PUT /fed/v1/buckets/{name})
func (s *Server) CreateBucket(c *gin.Context, name string) {

}
