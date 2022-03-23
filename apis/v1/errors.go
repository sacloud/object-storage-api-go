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

package v1

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	_ error = (*Error400)(nil)
	_ error = (*Error401)(nil)
	_ error = (*Error403)(nil)
	_ error = (*Error404)(nil)
	_ error = (*Error409)(nil)
	_ error = (*ErrorDefault)(nil)
)

var commonErrorFormat = "status: %d, message: %s, trace: %s, inner_error: %s"

// ActualError ErrorNNNが示すerrorを組み立てて返す
func (e Error400) Error() string {
	return fmt.Sprintf(commonErrorFormat, http.StatusBadRequest, e.Detail.Message, e.Detail.TraceId, e.Detail.Errors)
}

// ActualError ErrorNNNが示すerrorを組み立てて返す
func (e Error401) Error() string {
	return fmt.Sprintf(commonErrorFormat, http.StatusUnauthorized, e.Detail.Message, e.Detail.TraceId, e.Detail.Errors)
}

// ActualError ErrorNNNが示すerrorを組み立てて返す
func (e Error403) Error() string {
	return fmt.Sprintf(commonErrorFormat, http.StatusForbidden, e.Detail.Message, e.Detail.TraceId, e.Detail.Errors)
}

// ActualError ErrorNNNが示すerrorを組み立てて返す
func (e Error404) Error() string {
	return fmt.Sprintf(commonErrorFormat, http.StatusNotFound, e.Detail.Message, e.Detail.TraceId, e.Detail.Errors)
}

// ActualError ErrorNNNが示すerrorを組み立てて返す
func (e Error409) Error() string {
	return fmt.Sprintf(commonErrorFormat, http.StatusConflict, e.Detail.Message, e.Detail.TraceId, e.Detail.Errors)
}

// ActualError ErrorNNNが示すerrorを組み立てて返す
func (e ErrorDefault) Error() string {
	status := e.Detail.Code
	if status == 0 {
		// この段階までに既知のステータスコード判定はされている。ここで不明な場合は500にしておく
		status = http.StatusInternalServerError
	}
	return fmt.Sprintf(commonErrorFormat, status, e.Detail.Message, e.Detail.TraceId, e.Detail.Errors)
}

// String Stringer実装
func (e Error) String() string {
	return fmt.Sprintf("domain: %s, location: %s, location_type: %s, message: %s, reason: %s",
		e.Domain,
		e.Location,
		e.LocationType,
		e.Message,
		e.Reason,
	)
}

// String Stringer実装
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
