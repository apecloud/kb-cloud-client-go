// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

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
	Filter *string
}

// NewListAvailableCharsetsOptionalParameters creates an empty struct for parameters.
func NewListAvailableCharsetsOptionalParameters() *ListAvailableCharsetsOptionalParameters {
	this := ListAvailableCharsetsOptionalParameters{}
	return &this
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListAvailableCharsetsOptionalParameters) WithFilter(filter string) *ListAvailableCharsetsOptionalParameters {
	r.Filter = &filter
	return r
}

// ListAvailableCharsets List cluster charsets.
// list all available charsets in cluster
func (a *RdbmsApi) ListAvailableCharsets(ctx _context.Context, engineName string, orgName string, clusterName string, o ...ListAvailableCharsetsOptionalParameters) ([]string, *_nethttp.Response, error) {
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
		Path:        "/admin/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/charsets",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RdbmsApi.ListAvailableCharsets")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/charsets"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// ListAvailableCollatesOptionalParameters holds optional parameters for ListAvailableCollates.
type ListAvailableCollatesOptionalParameters struct {
	Charset *string
	Filter  *string
}

// NewListAvailableCollatesOptionalParameters creates an empty struct for parameters.
func NewListAvailableCollatesOptionalParameters() *ListAvailableCollatesOptionalParameters {
	this := ListAvailableCollatesOptionalParameters{}
	return &this
}

// WithCharset sets the corresponding parameter name and returns the struct.
func (r *ListAvailableCollatesOptionalParameters) WithCharset(charset string) *ListAvailableCollatesOptionalParameters {
	r.Charset = &charset
	return r
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListAvailableCollatesOptionalParameters) WithFilter(filter string) *ListAvailableCollatesOptionalParameters {
	r.Filter = &filter
	return r
}

// ListAvailableCollates List cluster collates.
// list all available collates in cluster
func (a *RdbmsApi) ListAvailableCollates(ctx _context.Context, engineName string, orgName string, clusterName string, o ...ListAvailableCollatesOptionalParameters) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []string
		optionalParams      ListAvailableCollatesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListAvailableCollatesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "rdbms",
		OperationID: "listAvailableCollates",
		Path:        "/admin/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/collates",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RdbmsApi.ListAvailableCollates")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/collates"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// ListAvailableCtypesOptionalParameters holds optional parameters for ListAvailableCtypes.
type ListAvailableCtypesOptionalParameters struct {
	Charset *string
	Filter  *string
}

// NewListAvailableCtypesOptionalParameters creates an empty struct for parameters.
func NewListAvailableCtypesOptionalParameters() *ListAvailableCtypesOptionalParameters {
	this := ListAvailableCtypesOptionalParameters{}
	return &this
}

// WithCharset sets the corresponding parameter name and returns the struct.
func (r *ListAvailableCtypesOptionalParameters) WithCharset(charset string) *ListAvailableCtypesOptionalParameters {
	r.Charset = &charset
	return r
}

// WithFilter sets the corresponding parameter name and returns the struct.
func (r *ListAvailableCtypesOptionalParameters) WithFilter(filter string) *ListAvailableCtypesOptionalParameters {
	r.Filter = &filter
	return r
}

// ListAvailableCtypes List cluster ctypes.
// list all available ctypes in cluster
func (a *RdbmsApi) ListAvailableCtypes(ctx _context.Context, engineName string, orgName string, clusterName string, o ...ListAvailableCtypesOptionalParameters) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []string
		optionalParams      ListAvailableCtypesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListAvailableCtypesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "rdbms",
		OperationID: "listAvailableCtypes",
		Path:        "/admin/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/ctypes",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RdbmsApi.ListAvailableCtypes")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/ctypes"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
