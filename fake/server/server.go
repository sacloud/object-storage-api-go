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
	"fmt"
	"net/http"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
)

var _ v1.ServerInterface = (*Server)(nil)

// Server オブジェクトストレージAPIのためのFake HTTPサーバ
type Server struct {
	// Engine fakeサーバの処理を担当するエンジンのインスタンス
	Engine *fake.Engine
}

// Handler http.Handlerを返す
func (s *Server) Handler() http.Handler {
	// Build a std http handler using generated HandlerWithOptions
	// Provide middleware and error handler if needed via options.
	// Add a simple ping handler and then register generated handlers.

	// Create base mux
	mux := http.NewServeMux()
	// ping endpoint
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	// Use generated handler with base router
	h := v1.HandlerWithOptions(s, v1.StdHTTPServerOptions{
		BaseRouter: mux,
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			// Map fake.Error to proper responses similar to previous implementation
			if engineErr, ok := err.(*fake.Error); ok {
				switch engineErr.Type {
				case fake.ErrorTypeInvalidRequest:
					w.WriteHeader(http.StatusBadRequest)
					_, _ = w.Write([]byte(fmt.Sprintf("invalid request: %s", engineErr.Error())))
					return
				case fake.ErrorTypeNotFound:
					w.WriteHeader(http.StatusNotFound)
					_, _ = w.Write([]byte(fmt.Sprintf("not found: %s", engineErr.Error())))
					return
				case fake.ErrorTypeConflict:
					w.WriteHeader(http.StatusConflict)
					_, _ = w.Write([]byte(fmt.Sprintf("conflict: %s", engineErr.Error())))
					return
				}
			}
			http.Error(w, fmt.Sprintf("unknown error: %s", err), http.StatusInternalServerError)
		},
	})

	return h
}
