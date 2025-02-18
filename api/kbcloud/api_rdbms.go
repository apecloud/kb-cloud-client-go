// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"context"
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RdbmsApi service type
type RdbmsApi common.Service

// ListAvailableCharsetsOptionalParameters holds optional parameters for ListAvailableCharsets.
type ListAvailableCharsetsOptionalParameters struct {
	Charset *string
	Filter  *string
}

// NewListAvailableCharsetsOptionalParameters creates an empty struct for parameters.
func NewListAvailableCharsetsOptionalParameters() *ListAvailableCharsetsOptionalParameters {
	this := ListAvailableCharsetsOptionalParameters{}
	return &this
}

// WithCharset sets the corresponding parameter name and returns the struct.
func (r *ListAvailableCharsetsOptionalParameters) WithCharset(charset string) *ListAvailableCharsetsOptionalParameters {
	r.Charset = &charset
	return r
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListAvailableCharsetsOptionalParameters) WithFilter(filter string) *ListAvailableCharsetsOptionalParameters {
	r.Filter = &filter
	return r
}

// ListAvailableCharsets List cluster charsets,collates or ctypes.
// list all available charsets,collates or ctypes in cluster
func (a *RdbmsApi) ListAvailableCharsets(ctx _context.Context, engineName string, orgName string, clusterName string, typeVar string, o ...ListAvailableCharsetsOptionalParameters) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []string
		optionalParams      ListAvailableCharsetsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListAvailableCharsetsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "rdbms",
		OperationID: "listAvailableCharsets",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/charsets",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RdbmsApi.ListAvailableCharsets")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/charsets"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("type", common.ParameterToString(typeVar, ""))
	if optionalParams.Charset != nil {
		localVarQueryParams.Add("charset", common.ParameterToString(*optionalParams.Charset, ""))
	}
	if optionalParams.Filter != nil {
		localVarQueryParams.Add("filter", common.ParameterToString(*optionalParams.Filter, ""))
	}
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
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
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
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewRdbmsApi Returns NewRdbmsApi.
func NewRdbmsApi(client *common.APIClient) *RdbmsApi {
	return &RdbmsApi{
		Client: client,
	}
}
