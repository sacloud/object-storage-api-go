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
	"encoding/json"
	"net/http"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
)

// DeleteBucket バケットの削除
// (DELETE /fed/v1/buckets/{name})
func (s *Server) DeleteBucket(w http.ResponseWriter, r *http.Request, bucketName v1.BucketName) {
	var paramJSON v1.DeleteBucketJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&paramJSON); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.Engine.DeleteBucket(paramJSON.ClusterId, bucketName.String()); err != nil {
		// let the global error handler in HandlerWithOptions handle it by returning the error
		// but since generated wrapper calls our method directly, write appropriate response here
		switch e := err.(type) {
		case *fake.Error:
			if e.Type == fake.ErrorTypeNotFound {
				http.Error(w, e.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateBucket バケットの作成
// (PUT /fed/v1/buckets/{name})
func (s *Server) CreateBucket(w http.ResponseWriter, r *http.Request, bucketName v1.BucketName) {
	var paramJSON v1.CreateBucketJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&paramJSON); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bucket, err := s.Engine.CreateBucket(paramJSON.ClusterId, bucketName.String())
	if err != nil {
		switch e := err.(type) {
		case *fake.Error:
			if e.Type == fake.ErrorTypeConflict {
				http.Error(w, e.Error(), http.StatusConflict)
				return
			}
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&v1.CreateBucketResponseBody{Data: *bucket})
}
