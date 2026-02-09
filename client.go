// Copyright 2022-2026 The object-storage-api-go Authors
// SPDX-License-Identifier: Apache-2.0

package objectstorage

import (
	"context"
	"fmt"
	"runtime"

	v2 "github.com/sacloud/object-storage-api-go/apis/v2"
	"github.com/sacloud/saclient-go"
)

const DefaultAPIRootURLFed = "https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/"
const DefaultAPIRootURLSite = "https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/%s/v2/"

var NewUserAgent = fmt.Sprintf(
	"object-storage-api-go/%s (%s/%s; +https://github.com/sacloud/object-storage-api-go)",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
)

type dummySecuritySource struct{}

func (ss dummySecuritySource) BasicAuth(ctx context.Context, operationName v2.OperationName) (v2.BasicAuth, error) {
	return v2.BasicAuth{Username: "", Password: "", Roles: nil}, nil
}

type FedClient struct {
	client *v2.Client
}

func NewFedClient(client saclient.ClientAPI) (*FedClient, error) {
	return NewFedClientWithAPIRootURL(client, DefaultAPIRootURLFed)
}

func NewFedClientWithAPIRootURL(client saclient.ClientAPI, apiRootURL string) (*FedClient, error) {
	dupable, ok := client.(saclient.ClientOptionAPI)
	if !ok {
		return nil, NewError("client does not implement saclient.ClientOptionAPI", nil)
	}
	augmented, err := dupable.DupWith(
		saclient.WithUserAgent(NewUserAgent),
		saclient.WithBigInt(false),
		saclient.WithForceAutomaticAuthentication(),
	)
	if err != nil {
		return nil, err
	}
	c, err := v2.NewClient(apiRootURL, &dummySecuritySource{}, v2.WithClient(augmented))
	if err != nil {
		return nil, err
	}
	return &FedClient{client: c}, nil
}

type SiteClient struct {
	client *v2.Client
}

func NewSiteClient(client saclient.ClientAPI, siteId string) (*SiteClient, error) {
	return NewSiteClientWithAPIRootURL(client, fmt.Sprintf(DefaultAPIRootURLSite, siteId))
}

func NewSiteClientWithAPIRootURL(client saclient.ClientAPI, apiRootURL string) (*SiteClient, error) {
	dupable, ok := client.(saclient.ClientOptionAPI)
	if !ok {
		return nil, NewError("client does not implement saclient.ClientOptionAPI", nil)
	}
	argumented, err := dupable.DupWith(
		saclient.WithUserAgent(NewUserAgent),
		saclient.WithBigInt(false),
		saclient.WithForceAutomaticAuthentication(),
	)
	if err != nil {
		return nil, err
	}
	c, err := v2.NewClient(apiRootURL, &dummySecuritySource{}, v2.WithClient(argumented))
	if err != nil {
		return nil, err
	}
	return &SiteClient{client: c}, nil
}
