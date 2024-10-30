// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"bytes"
	_context "context"
	_fmt "fmt"
	_io "io"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ZoneApi service type
type ZoneApi common.Service

// GetZones Get zone.
// Get zone
func (a *ZoneApi) GetZones(ctx _context.Context, providerName string, regionName string, zoneName string) (Zone, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  Zone
	)

    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ZoneApi.GetZones")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/providers/{providerName}/regions/{regionName}/zones/{zoneName}"
	localVarPath = strings.Replace(localVarPath, "{"+"providerName"+"}", _neturl.PathEscape(common.ParameterToString(providerName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"regionName"+"}", _neturl.PathEscape(common.ParameterToString(regionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"zoneName"+"}", _neturl.PathEscape(common.ParameterToString(zoneName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	
        common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"BearerToken", "authorization"},
	)
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403||localVarHTTPResponse.StatusCode == 404{
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListZones Get zone list.
// Get zone list
func (a *ZoneApi) ListZones(ctx _context.Context, providerName string, regionName string) (ZoneList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  ZoneList
	)

    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ZoneApi.ListZones")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/providers/{providerName}/regions/{regionName}/zones"
	localVarPath = strings.Replace(localVarPath, "{"+"providerName"+"}", _neturl.PathEscape(common.ParameterToString(providerName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"regionName"+"}", _neturl.PathEscape(common.ParameterToString(regionName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	
        common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"BearerToken", "authorization"},
	)
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403||localVarHTTPResponse.StatusCode == 404{
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewZoneApi Returns NewZoneApi.
func NewZoneApi(client *common.APIClient) *ZoneApi {
	return &ZoneApi{
		Client: client,
	}
}
