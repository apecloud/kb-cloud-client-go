// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineApi service type
type EngineApi common.Service

// EngineActionInOrg Manage engine in organization.
func (a *EngineApi) EngineActionInOrg(ctx _context.Context, orgName string, actionInfo interface{}) (bool, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue bool
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineApi.EngineActionInOrg")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/engines"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &actionInfo
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

// ListEnginesInEnvOptionalParameters holds optional parameters for ListEnginesInEnv.
type ListEnginesInEnvOptionalParameters struct {
	Name     *string
	Type     *EngineType
	Version  *string
	Provider *string
	All      *bool
}

// NewListEnginesInEnvOptionalParameters creates an empty struct for parameters.
func NewListEnginesInEnvOptionalParameters() *ListEnginesInEnvOptionalParameters {
	this := ListEnginesInEnvOptionalParameters{}
	return &this
}

// WithName sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInEnvOptionalParameters) WithName(name string) *ListEnginesInEnvOptionalParameters {
	r.Name = &name
	return r
}

// WithType sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInEnvOptionalParameters) WithType(typeVar EngineType) *ListEnginesInEnvOptionalParameters {
	r.Type = &typeVar
	return r
}

// WithVersion sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInEnvOptionalParameters) WithVersion(version string) *ListEnginesInEnvOptionalParameters {
	r.Version = &version
	return r
}

// WithProvider sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInEnvOptionalParameters) WithProvider(provider string) *ListEnginesInEnvOptionalParameters {
	r.Provider = &provider
	return r
}

// WithAll sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInEnvOptionalParameters) WithAll(all bool) *ListEnginesInEnvOptionalParameters {
	r.All = &all
	return r
}

// ListEnginesInEnv List engines in environment.
func (a *EngineApi) ListEnginesInEnv(ctx _context.Context, environmentName string, o ...ListEnginesInEnvOptionalParameters) ([]Engine, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []Engine
		optionalParams      ListEnginesInEnvOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEnginesInEnvOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineApi.ListEnginesInEnv")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/environments/{environmentName}/engines"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentName"+"}", _neturl.PathEscape(common.ParameterToString(environmentName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Name != nil {
		localVarQueryParams.Add("name", common.ParameterToString(*optionalParams.Name, ""))
	}
	if optionalParams.Type != nil {
		localVarQueryParams.Add("type", common.ParameterToString(*optionalParams.Type, ""))
	}
	if optionalParams.Version != nil {
		localVarQueryParams.Add("version", common.ParameterToString(*optionalParams.Version, ""))
	}
	if optionalParams.Provider != nil {
		localVarQueryParams.Add("provider", common.ParameterToString(*optionalParams.Provider, ""))
	}
	if optionalParams.All != nil {
		localVarQueryParams.Add("all", common.ParameterToString(*optionalParams.All, ""))
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

// ListEnginesInOrgOptionalParameters holds optional parameters for ListEnginesInOrg.
type ListEnginesInOrgOptionalParameters struct {
	EnvironmentName *string
	Name            *string
	Type            *EngineType
	Version         *string
	Provider        *string
}

// NewListEnginesInOrgOptionalParameters creates an empty struct for parameters.
func NewListEnginesInOrgOptionalParameters() *ListEnginesInOrgOptionalParameters {
	this := ListEnginesInOrgOptionalParameters{}
	return &this
}

// WithEnvironmentName sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInOrgOptionalParameters) WithEnvironmentName(environmentName string) *ListEnginesInOrgOptionalParameters {
	r.EnvironmentName = &environmentName
	return r
}

// WithName sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInOrgOptionalParameters) WithName(name string) *ListEnginesInOrgOptionalParameters {
	r.Name = &name
	return r
}

// WithType sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInOrgOptionalParameters) WithType(typeVar EngineType) *ListEnginesInOrgOptionalParameters {
	r.Type = &typeVar
	return r
}

// WithVersion sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInOrgOptionalParameters) WithVersion(version string) *ListEnginesInOrgOptionalParameters {
	r.Version = &version
	return r
}

// WithProvider sets the corresponding parameter name and returns the struct.
func (r *ListEnginesInOrgOptionalParameters) WithProvider(provider string) *ListEnginesInOrgOptionalParameters {
	r.Provider = &provider
	return r
}

// ListEnginesInOrg List engines in organization.
func (a *EngineApi) ListEnginesInOrg(ctx _context.Context, orgName string, o ...ListEnginesInOrgOptionalParameters) ([]Engine, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []Engine
		optionalParams      ListEnginesInOrgOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEnginesInOrgOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineApi.ListEnginesInOrg")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/engines"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.EnvironmentName != nil {
		localVarQueryParams.Add("environmentName", common.ParameterToString(*optionalParams.EnvironmentName, ""))
	}
	if optionalParams.Name != nil {
		localVarQueryParams.Add("name", common.ParameterToString(*optionalParams.Name, ""))
	}
	if optionalParams.Type != nil {
		localVarQueryParams.Add("type", common.ParameterToString(*optionalParams.Type, ""))
	}
	if optionalParams.Version != nil {
		localVarQueryParams.Add("version", common.ParameterToString(*optionalParams.Version, ""))
	}
	if optionalParams.Provider != nil {
		localVarQueryParams.Add("provider", common.ParameterToString(*optionalParams.Provider, ""))
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

// NewEngineApi Returns NewEngineApi.
func NewEngineApi(client *common.APIClient) *EngineApi {
	return &EngineApi{
		Client: client,
	}
}
