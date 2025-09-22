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
	"os"

	"github.com/gin-gonic/gin"
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
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("OJS_SERVER_DEBUG") != "" {
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	if os.Getenv("OJS_SERVER_LOGGING") != "" {
		engine.Use(gin.Logger())
	}

	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	v1.RegisterHandlers(engine, s)
	return engine
}

func (s *Server) handleError(c *gin.Context, err error) {
	if c == nil || err == nil {
		panic("invalid arguments")
	}

	if engineErr, ok := err.(*fake.Error); ok {
		switch engineErr.Type {
		case fake.ErrorTypeInvalidRequest:
			c.JSON(http.StatusBadRequest, &v1.Error400{
				Detail: v1.ErrorDetail{
					Code:    http.StatusBadRequest,
					Message: v1.ErrorMessage(engineErr.Error()),
				},
			})
			return
		case fake.ErrorTypeNotFound:
			c.JSON(http.StatusNotFound, &v1.Error404{
				Detail: v1.ErrorDetail{
					Code:    http.StatusNotFound,
					Message: v1.ErrorMessage(engineErr.Error()),
				},
			})
			return
		case fake.ErrorTypeConflict:
			c.JSON(http.StatusConflict, &v1.Error409{
				Detail: v1.ErrorDetail{
					Code:    http.StatusConflict,
					Message: v1.ErrorMessage(engineErr.Error()),
				},
			})
			return
		}
	}

	// Note: この実装ではfake.Errorで未知のステータスコードについて全てInternalServerErrorとして扱う
	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("unknown error: %s", err)})
}
