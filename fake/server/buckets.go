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

// DeleteBucket バケットの削除
// (DELETE /fed/v1/buckets/{name})
func (s *Server) DeleteBucket(c *gin.Context, bucketName v1.BucketName) {
	var paramJSON v1.DeleteBucketJSONRequestBody
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.Engine.DeleteBucket(paramJSON.ClusterId, bucketName.String()); err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateBucket バケットの作成
// (PUT /fed/v1/buckets/{name})
func (s *Server) CreateBucket(c *gin.Context, bucketName v1.BucketName) {
	var paramJSON v1.CreateBucketJSONRequestBody
	if err := c.ShouldBindJSON(&paramJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bucket, err := s.Engine.CreateBucket(paramJSON.ClusterId, bucketName.String())
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, &v1.CreateBucketResponseBody{
		Data: *bucket,
	})
}
