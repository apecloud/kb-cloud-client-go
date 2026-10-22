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

// EngineOptionApi service type
type EngineOptionApi common.Service

// GetEngineOptionOptionalParameters holds optional parameters for GetEngineOption.
type GetEngineOptionOptionalParameters struct {
	OrgName         *string
	EnvironmentName *string
	EngineMode      *string
	ServiceVersion  *string
}

// NewGetEngineOptionOptionalParameters creates an empty struct for parameters.
func NewGetEngineOptionOptionalParameters() *GetEngineOptionOptionalParameters {
	this := GetEngineOptionOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *GetEngineOptionOptionalParameters) WithOrgName(orgName string) *GetEngineOptionOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithEnvironmentName sets the corresponding parameter name and returns the struct.
func (r *GetEngineOptionOptionalParameters) WithEnvironmentName(environmentName string) *GetEngineOptionOptionalParameters {
	r.EnvironmentName = &environmentName
	return r
}

// WithEngineMode sets the corresponding parameter name and returns the struct.
func (r *GetEngineOptionOptionalParameters) WithEngineMode(engineMode string) *GetEngineOptionOptionalParameters {
	r.EngineMode = &engineMode
	return r
}

// WithServiceVersion sets the corresponding parameter name and returns the struct.
func (r *GetEngineOptionOptionalParameters) WithServiceVersion(serviceVersion string) *GetEngineOptionOptionalParameters {
	r.ServiceVersion = &serviceVersion
	return r
}

// GetEngineOption Get engineOption.
// Get a engineOption detail
func (a *EngineOptionApi) GetEngineOption(ctx _context.Context, engineName string, o ...GetEngineOptionOptionalParameters) (EngineOption, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EngineOption
		optionalParams      GetEngineOptionOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetEngineOptionOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engineOption",
		OperationID: "getEngineOption",
		Path:        "/api/v1/engineOptions/{engineName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineOptionApi.GetEngineOption")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/engineOptions/{engineName}"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.EnvironmentName != nil {
		localVarQueryParams.Add("environmentName", common.ParameterToString(*optionalParams.EnvironmentName, ""))
	}
	if optionalParams.EngineMode != nil {
		localVarQueryParams.Add("engineMode", common.ParameterToString(*optionalParams.EngineMode, ""))
	}
	if optionalParams.ServiceVersion != nil {
		localVarQueryParams.Add("serviceVersion", common.ParameterToString(*optionalParams.ServiceVersion, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// ListEngineOptionsOptionalParameters holds optional parameters for ListEngineOptions.
type ListEngineOptionsOptionalParameters struct {
	Version                             *EngineOptionVersion
	FilterLoadBackupFromOfflineInstance *bool
}

// NewListEngineOptionsOptionalParameters creates an empty struct for parameters.
func NewListEngineOptionsOptionalParameters() *ListEngineOptionsOptionalParameters {
	this := ListEngineOptionsOptionalParameters{}
	return &this
}

// WithVersion sets the corresponding parameter name and returns the struct.
func (r *ListEngineOptionsOptionalParameters) WithVersion(version EngineOptionVersion) *ListEngineOptionsOptionalParameters {
	r.Version = &version
	return r
}

// WithFilterLoadBackupFromOfflineInstance sets the corresponding parameter name and returns the struct.
func (r *ListEngineOptionsOptionalParameters) WithFilterLoadBackupFromOfflineInstance(filterLoadBackupFromOfflineInstance bool) *ListEngineOptionsOptionalParameters {
	r.FilterLoadBackupFromOfflineInstance = &filterLoadBackupFromOfflineInstance
	return r
}

// ListEngineOptions List all engineOptions.
// list all engineOptions
func (a *EngineOptionApi) ListEngineOptions(ctx _context.Context, o ...ListEngineOptionsOptionalParameters) (EngineOptionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EngineOptionList
		optionalParams      ListEngineOptionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEngineOptionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engineOption",
		OperationID: "listEngineOptions",
		Path:        "/api/v1/engineOptions",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineOptionApi.ListEngineOptions")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/engineOptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Version != nil {
		localVarQueryParams.Add("version", common.ParameterToString(*optionalParams.Version, ""))
	}
	if optionalParams.FilterLoadBackupFromOfflineInstance != nil {
		localVarQueryParams.Add("filterLoadBackupFromOfflineInstance", common.ParameterToString(*optionalParams.FilterLoadBackupFromOfflineInstance, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// NewEngineOptionApi Returns NewEngineOptionApi.
func NewEngineOptionApi(client *common.APIClient) *EngineOptionApi {
	return &EngineOptionApi{
		Client: client,
	}
}
