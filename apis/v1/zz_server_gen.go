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

// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package v1

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// サイトアカウントの削除
	// (DELETE /account)
	DeleteAccount(c *gin.Context)
	// サイトアカウントの取得
	// (GET /account)
	GetAccount(c *gin.Context)
	// サイトアカウントの作成
	// (POST /account)
	PostAccount(c *gin.Context)
	// サイトアカウントのアクセスキーの取得
	// (GET /account/keys)
	GetAccountKeys(c *gin.Context)
	// サイトアカウントのアクセスキーの発行
	// (POST /account/keys)
	PostAccountKeys(c *gin.Context)
	// サイトアカウントのアクセスキーの削除
	// (DELETE /account/keys/{id})
	DeleteAccountKeysId(c *gin.Context, id string)
	// サイトアカウントのアクセスキーの取得
	// (GET /account/keys/{id})
	GetAccountKeysId(c *gin.Context, id string)
	// バケットの削除
	// (DELETE /buckets/{name})
	DeleteBucketsName(c *gin.Context, name string)
	// バケットの作成
	// (PUT /buckets/{name})
	PutBucketsName(c *gin.Context, name string)
	// サイト一覧の取得
	// (GET /clusters)
	GetClusters(c *gin.Context)
	// サイトの取得
	// (GET /clusters/{id})
	GetClustersId(c *gin.Context, id string)
	// パーミッション一覧の取得
	// (GET /permissions)
	GetPermissions(c *gin.Context)
	// パーミッションの作成
	// (POST /permissions)
	PostPermissions(c *gin.Context)
	// パーミッションの削除
	// (DELETE /permissions/{id})
	DeletePermissionsId(c *gin.Context, id string)
	// パーミッションの取得
	// (GET /permissions/{id})
	GetPermissionsId(c *gin.Context, id string)
	// パーミッションの更新
	// (PUT /permissions/{id})
	PutPermissionsId(c *gin.Context, id string)
	// パーミッションが保有するアクセスキー一覧の取得
	// (GET /permissions/{id}/keys)
	GetPermissionsIdKeys(c *gin.Context, id string)
	// パーミッションのアクセスキーの発行
	// (POST /permissions/{id}/keys)
	PostPermissionsIdKeys(c *gin.Context, id string)
	// パーミッションが保有するアクセスキーの削除
	// (DELETE /permissions/{id}/keys/{key_id})
	DeletePermissionsIdKeysKeyId(c *gin.Context, id string, keyId string)
	// パーミッションが保有するアクセスキーの取得
	// (GET /permissions/{id}/keys/{key_id})
	GetPermissionsIdKeysKeyId(c *gin.Context, id string, keyId string)
	// サイトのステータスの取得
	// (GET /status)
	GetStatus(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// DeleteAccount operation middleware
func (siw *ServerInterfaceWrapper) DeleteAccount(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteAccount(c)
}

// GetAccount operation middleware
func (siw *ServerInterfaceWrapper) GetAccount(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAccount(c)
}

// PostAccount operation middleware
func (siw *ServerInterfaceWrapper) PostAccount(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostAccount(c)
}

// GetAccountKeys operation middleware
func (siw *ServerInterfaceWrapper) GetAccountKeys(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAccountKeys(c)
}

// PostAccountKeys operation middleware
func (siw *ServerInterfaceWrapper) PostAccountKeys(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostAccountKeys(c)
}

// DeleteAccountKeysId operation middleware
func (siw *ServerInterfaceWrapper) DeleteAccountKeysId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteAccountKeysId(c, id)
}

// GetAccountKeysId operation middleware
func (siw *ServerInterfaceWrapper) GetAccountKeysId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAccountKeysId(c, id)
}

// DeleteBucketsName operation middleware
func (siw *ServerInterfaceWrapper) DeleteBucketsName(c *gin.Context) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", c.Param("name"), &name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter name: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeleteBucketsName(c, name)
}

// PutBucketsName operation middleware
func (siw *ServerInterfaceWrapper) PutBucketsName(c *gin.Context) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", c.Param("name"), &name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter name: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PutBucketsName(c, name)
}

// GetClusters operation middleware
func (siw *ServerInterfaceWrapper) GetClusters(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetClusters(c)
}

// GetClustersId operation middleware
func (siw *ServerInterfaceWrapper) GetClustersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetClustersId(c, id)
}

// GetPermissions operation middleware
func (siw *ServerInterfaceWrapper) GetPermissions(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPermissions(c)
}

// PostPermissions operation middleware
func (siw *ServerInterfaceWrapper) PostPermissions(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostPermissions(c)
}

// DeletePermissionsId operation middleware
func (siw *ServerInterfaceWrapper) DeletePermissionsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeletePermissionsId(c, id)
}

// GetPermissionsId operation middleware
func (siw *ServerInterfaceWrapper) GetPermissionsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPermissionsId(c, id)
}

// PutPermissionsId operation middleware
func (siw *ServerInterfaceWrapper) PutPermissionsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PutPermissionsId(c, id)
}

// GetPermissionsIdKeys operation middleware
func (siw *ServerInterfaceWrapper) GetPermissionsIdKeys(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPermissionsIdKeys(c, id)
}

// PostPermissionsIdKeys operation middleware
func (siw *ServerInterfaceWrapper) PostPermissionsIdKeys(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostPermissionsIdKeys(c, id)
}

// DeletePermissionsIdKeysKeyId operation middleware
func (siw *ServerInterfaceWrapper) DeletePermissionsIdKeysKeyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	// ------------- Path parameter "key_id" -------------
	var keyId string

	err = runtime.BindStyledParameter("simple", false, "key_id", c.Param("key_id"), &keyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter key_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.DeletePermissionsIdKeysKeyId(c, id, keyId)
}

// GetPermissionsIdKeysKeyId operation middleware
func (siw *ServerInterfaceWrapper) GetPermissionsIdKeysKeyId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	// ------------- Path parameter "key_id" -------------
	var keyId string

	err = runtime.BindStyledParameter("simple", false, "key_id", c.Param("key_id"), &keyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter key_id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetPermissionsIdKeysKeyId(c, id, keyId)
}

// GetStatus operation middleware
func (siw *ServerInterfaceWrapper) GetStatus(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetStatus(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.DELETE(options.BaseURL+"/account", wrapper.DeleteAccount)

	router.GET(options.BaseURL+"/account", wrapper.GetAccount)

	router.POST(options.BaseURL+"/account", wrapper.PostAccount)

	router.GET(options.BaseURL+"/account/keys", wrapper.GetAccountKeys)

	router.POST(options.BaseURL+"/account/keys", wrapper.PostAccountKeys)

	router.DELETE(options.BaseURL+"/account/keys/:id", wrapper.DeleteAccountKeysId)

	router.GET(options.BaseURL+"/account/keys/:id", wrapper.GetAccountKeysId)

	router.DELETE(options.BaseURL+"/buckets/:name", wrapper.DeleteBucketsName)

	router.PUT(options.BaseURL+"/buckets/:name", wrapper.PutBucketsName)

	router.GET(options.BaseURL+"/clusters", wrapper.GetClusters)

	router.GET(options.BaseURL+"/clusters/:id", wrapper.GetClustersId)

	router.GET(options.BaseURL+"/permissions", wrapper.GetPermissions)

	router.POST(options.BaseURL+"/permissions", wrapper.PostPermissions)

	router.DELETE(options.BaseURL+"/permissions/:id", wrapper.DeletePermissionsId)

	router.GET(options.BaseURL+"/permissions/:id", wrapper.GetPermissionsId)

	router.PUT(options.BaseURL+"/permissions/:id", wrapper.PutPermissionsId)

	router.GET(options.BaseURL+"/permissions/:id/keys", wrapper.GetPermissionsIdKeys)

	router.POST(options.BaseURL+"/permissions/:id/keys", wrapper.PostPermissionsIdKeys)

	router.DELETE(options.BaseURL+"/permissions/:id/keys/:key_id", wrapper.DeletePermissionsIdKeysKeyId)

	router.GET(options.BaseURL+"/permissions/:id/keys/:key_id", wrapper.GetPermissionsIdKeysKeyId)

	router.GET(options.BaseURL+"/status", wrapper.GetStatus)

	return router
}
