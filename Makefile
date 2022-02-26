#
# Copyright 2022 The sacloud/object-storage-api-go Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

AUTHOR          ?="The sacloud/object-storage-api-go authors"
COPYRIGHT_YEAR  ?="2022"
COPYRIGHT_FILES ?=$$(find . -name "*.go" -print | grep -v "/vendor/")

default: gen fmt set-license goimports lint test

# TODO テスト用モックサーバを追加したらここを修正する
# .PHONY: all
# all: dist/phy-go-fake-server
#
# dist/phy-go-fake-server: *.go
# 	go build -o dist/phy-go-fake-server cmd/phy-go-fake-server/*.go

.PHONY: test
test:
	TESTACC= go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 -race;

.PHONY: testacc
testacc:
	TESTACC=1 go test ./... $(TESTARGS) -v -timeout=120m -parallel=8 ;

.PHONY: tools
tools:
	npm install -g @apidevtools/swagger-cli
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/sacloud/addlicense@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.44.2/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.44.2

.PHONY: clean
clean:
	find . -type f -name "*_gen.go" -delete
	rm apis/v1/spec/original-swagger.yaml
	rm apis/v1/spec/swagger.json

.PHONY: gen
gen: _gen fmt goimports set-license

.PHONY: _gen
_gen: apis/v1/spec/original-swagger.yaml apis/v1/spec/swagger.json apis/v1/zz_types_gen.go apis/v1/zz_client_gen.go apis/v1/zz_server_gen.go
	go generate ./...

apis/v1/spec/original-swagger.yaml: apis/v1/spec/original-swagger.json
	swagger-cli bundle apis/v1/spec/original-swagger.json -o apis/v1/spec/original-swagger.yaml --type yaml

apis/v1/spec/swagger.json: apis/v1/spec/swagger.yaml
	swagger-cli bundle apis/v1/spec/swagger.yaml -o apis/v1/spec/swagger.json --type json

apis/v1/zz_types_gen.go: apis/v1/spec/swagger.yaml apis/v1/spec/codegen/types.yaml
	oapi-codegen -config apis/v1/spec/codegen/types.yaml apis/v1/spec/swagger.yaml

apis/v1/zz_client_gen.go: apis/v1/spec/swagger.yaml apis/v1/spec/codegen/client.yaml
	oapi-codegen -config apis/v1/spec/codegen/client.yaml apis/v1/spec/swagger.yaml

apis/v1/zz_server_gen.go: apis/v1/spec/swagger.yaml apis/v1/spec/codegen/gin.yaml
	oapi-codegen -config apis/v1/spec/codegen/gin.yaml apis/v1/spec/swagger.yaml

.PHONY: goimports
goimports: fmt
	goimports -l -w .

.PHONY: fmt
fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

.PHONY: set-license
set-license:
	@addlicense -c $(AUTHOR) -y $(COPYRIGHT_YEAR) $(COPYRIGHT_FILES)

.PHONY: lint
lint: lint-go lint-def textlint

.PHONY: lint-go
lint-go:
	golangci-lint run ./...

.PHONY: lint-def
lint-def:
	docker run --rm -v $$PWD:$$PWD -w $$PWD stoplight/spectral:latest lint -F warn apis/v1/spec/swagger.yaml

.PHONY: textlint
textlint:
	@docker run --rm -v $$PWD:/work -w /work ghcr.io/sacloud/textlint-action:v0.0.1 .

.PHONY: godoc
godoc:
	echo "URL: http://localhost:6060/pkg/github.com/sacloud/object-storage-api-go/"
	godoc -http=localhost:6060

