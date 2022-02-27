// Copyright 2021-2022 The phy-go authors
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

package v1

import (
	"fmt"
	"net/http"
	"strings"
)

// errorResponser APIレスポンスが返すエラー型が実装すべきインターフェース
//
// Note: 本来は各エラー型(ErrorNNN)でerrorインターフェースを実装させたかったが、
//       各エラー型にはswagger.yamlでErrorというフィールドが定義されているため実装できない。
//       このため別途errorを返すためのインターフェースとしてerrorResponserを用いている。
type errorResponser interface {
	ActualError() error
}

var (
	_ errorResponser = (*Error400)(nil)
	_ errorResponser = (*Error401)(nil)
	_ errorResponser = (*Error403)(nil)
	_ errorResponser = (*Error404)(nil)
	_ errorResponser = (*Error409)(nil)
	_ errorResponser = (*ErrorDefault)(nil)
)

var commonErrorFormat = "status: %d, message: %s, trace: %s, inner_error: %s"

func (e Error400) ActualError() error {
	return fmt.Errorf(commonErrorFormat, http.StatusBadRequest, e.Error.Message, e.Error.TraceId, e.Error.Errors)
}

func (e Error401) ActualError() error {
	return fmt.Errorf(commonErrorFormat, http.StatusUnauthorized, e.Error.Message, e.Error.TraceId, e.Error.Errors)
}

func (e Error403) ActualError() error {
	return fmt.Errorf(commonErrorFormat, http.StatusForbidden, e.Error.Message, e.Error.TraceId, e.Error.Errors)
}

func (e Error404) ActualError() error {
	return fmt.Errorf(commonErrorFormat, http.StatusNotFound, e.Error.Message, e.Error.TraceId, e.Error.Errors)
}

func (e Error409) ActualError() error {
	return fmt.Errorf(commonErrorFormat, http.StatusConflict, e.Error.Message, e.Error.TraceId, e.Error.Errors)
}

func (e ErrorDefault) ActualError() error {
	status := e.Error.Code
	if status == 0 {
		// この段階までに既知のステータスコード判定はされている。ここで不明な場合は500にしておく
		status = http.StatusInternalServerError
	}
	return fmt.Errorf(commonErrorFormat, status, e.Error.Message, e.Error.TraceId, e.Error.Errors)
}

func (e Error) String() string {
	return fmt.Sprintf("domain: %s, location: %s, location_type: %s, message: %s, reason: %s",
		e.Domain,
		e.Location,
		e.LocationType,
		e.Message,
		e.Reason,
	)
}

func (e Errors) String() string {
	if len(e) == 0 {
		return "(empty)"
	}
	var errorStrings []string
	for _, err := range e {
		errorStrings = append(errorStrings, fmt.Sprintf("{%s}", err.String()))
	}

	return strings.Join(errorStrings, ", ")
}
