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

//go:build acctest
// +build acctest

package objectstorage_test

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

const (
	siteId  = "isk01"                               // Note: 本来はサイトAPI経由で取得すべきだが現状ではサイトが1つしかないため定数として保持/利用する
	charSet = "abcdefghijklmnopqrstuvwxyz012346789" // ランダム名生成で利用する文字
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// TestAccSiteAndStatusAPI サイト/ステータス関連APIの疎通確認
func TestAccSiteAndStatusAPI(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx := context.Background()
	siteOp := objectstorage.NewSiteOp(accTestClient)

	// List
	sites, err := siteOp.List(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, sites)

	// Read
	site, err := siteOp.Read(ctx, sites[0].Id)
	require.NoError(t, err)
	require.NotEmpty(t, site)

	// Site Status
	statusOp := objectstorage.NewSiteStatusOp(accTestClient)
	status, err := statusOp.Read(ctx, site.Id)

	require.NoError(t, err)
	require.NotEmpty(t, status)
}

// TestAccBucketHandling バケット周りの一連の操作のテスト
//
// 一連の操作は以下のステップで行う
//   - Step1: バケットを作成
//   - Step2: アクセスキーを用いてAWS SDK経由でバケットが作成されたかを確認
//   - Step3: 各リソースのクリーンアップ
//
// Note: バケット一覧の参照のためにサイトアカウントのアクセスキーが必要
func TestAccBucketHandling(t *testing.T) {
	skipIfNoAPIKey(t)
	skipIfNoEnv(t, "SACLOUD_OJS_ACCESS_KEY_ID", "SACLOUD_OJS_SECRET_ACCESS_KEY")

	ctx := context.Background()
	bucketName := randomName(28)

	// Step1: バケット作成
	bucketOp := objectstorage.NewBucketOp(accTestClient)
	{
		created, err := bucketOp.Create(ctx, siteId, bucketName)
		require.NoError(t, err)
		require.NotEmpty(t, created)
	}

	// Step2: バケットにアクセス
	{
		cred := credentials.NewStaticCredentialsProvider(
			os.Getenv("SACLOUD_OJS_ACCESS_KEY_ID"),
			os.Getenv("SACLOUD_OJS_SECRET_ACCESS_KEY"),
			"",
		)
		endpoint := aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:               "https://s3.isk01.sakurastorage.jp/",
					HostnameImmutable: true,
					SigningRegion:     region,
				}, nil
			},
		)

		cfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion("jp-north-1"),
			config.WithCredentialsProvider(cred),
			config.WithEndpointResolverWithOptions(endpoint),
		)
		if err != nil {
			t.Fatal(err)
		}
		s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		})

		output, err := s3Client.ListBuckets(ctx, nil)
		require.NoError(t, err)
		require.NotEmpty(t, output)

		exist := false
		for _, b := range output.Buckets {
			if b.Name != nil && *b.Name == bucketName {
				exist = true
				break
			}
		}
		require.True(t, exist, "bucket %q is not exist", bucketName)
	}

	// Step3: クリーンアップ
	require.NoError(t, bucketOp.Delete(ctx, siteId, bucketName))
}

// TestAccAccessToBucketWithPermissionKey パーミッションキーによるオブジェクトへのアクセス
func TestAccAccessToBucketObjectWithPermissionKey(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx := context.Background()
	bucketName := randomName(28)

	// Step1: バケット作成
	bucketOp := objectstorage.NewBucketOp(accTestClient)
	{
		created, err := bucketOp.Create(ctx, siteId, bucketName)
		require.NoError(t, err)
		require.NotEmpty(t, created)
	}

	// Step2: バケットにアクセスできるパーミッション/アクセスキーの作成
	permissionOp := objectstorage.NewPermissionOp(accTestClient)
	var permission *v1.Permission
	var key *v1.PermissionKey
	{
		created, err := permissionOp.Create(ctx, siteId, &v1.CreatePermissionParams{
			BucketControls: v1.BucketControls{
				{
					BucketName: v1.BucketName(bucketName),
					CanRead:    true,
					CanWrite:   true,
				},
			},
			DisplayName: v1.DisplayName(bucketName),
		})
		require.NoError(t, err)
		require.NotEmpty(t, created)
		permission = created

		createdKey, err := permissionOp.CreateAccessKey(ctx, siteId, permission.Id.Int64())
		require.NoError(t, err)
		require.NotEmpty(t, createdKey)
		key = createdKey
	}

	// Step3: 作成したアクセスキーでバケットにアクセス
	{
		cred := credentials.NewStaticCredentialsProvider(
			key.Id.String(), key.Secret.String(),
			"",
		)
		endpoint := aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:               "https://s3.isk01.sakurastorage.jp/",
					HostnameImmutable: true,
					SigningRegion:     region,
				}, nil
			},
		)

		cfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion("jp-north-1"),
			config.WithCredentialsProvider(cred),
			config.WithEndpointResolverWithOptions(endpoint),
		)
		if err != nil {
			t.Fatal(err)
		}
		s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		})

		objectKey := "foobar"
		objectBodyText := "body of s3://[bucket_name]/foobar"
		objectBody := bytes.NewBufferString(objectBodyText)

		// オブジェクトの作成
		created, err := s3Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: &bucketName,
			Key:    &objectKey,
			Body:   objectBody,
		})
		require.NoError(t, err)
		require.NotEmpty(t, created)

		// オブジェクトの読み込み
		read, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
			Bucket: &bucketName,
			Key:    &objectKey,
		})
		require.NoError(t, err)
		require.NotEmpty(t, read)

		readText, err := io.ReadAll(read.Body)
		require.NoError(t, err)
		require.Equal(t, objectBodyText, string(readText))

		// オブジェクトの削除
		deleted, err := s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
			Bucket: &bucketName,
			Key:    &objectKey,
		})
		require.NoError(t, err)
		require.NotEmpty(t, deleted)
	}

	// Step4: クリーンアップ
	require.NoError(t, permissionOp.DeleteAccessKey(ctx, siteId, permission.Id.Int64(), key.Id.String()))
	require.NoError(t, permissionOp.Delete(ctx, siteId, permission.Id.Int64()))
	require.NoError(t, bucketOp.Delete(ctx, siteId, bucketName))
}

var accTestClient = func() *objectstorage.Client {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")
	rootURL := os.Getenv("SAKURACLOUD_OJS_ROOT_URL")

	if rootURL == "" {
		rootURL = defaultServerURL
	}

	httpClient := &http.Client{}

	client := &objectstorage.Client{
		Token:      token,
		Secret:     secret,
		APIRootURL: serverURL,
		Trace:      os.Getenv("SAKURACLOUD_TRACE") != "",
		HTTPClient: httpClient,
	}

	return client
}()

// skipIfNoEnv 指定の環境変数のいずれかが空の場合はt.SkipNow()する
func skipIfNoEnv(t *testing.T, envs ...string) {
	var emptyEnvs []string
	for _, env := range envs {
		if os.Getenv(env) == "" {
			emptyEnvs = append(emptyEnvs, env)
		}
	}
	if len(emptyEnvs) > 0 {
		for _, env := range emptyEnvs {
			t.Logf("environment variable %q is not set", env)
		}
		t.SkipNow()
	}
}

func skipIfNoAPIKey(t *testing.T) {
	skipIfNoEnv(t, "SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")
}

func randomName(strlen int) string {
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}
