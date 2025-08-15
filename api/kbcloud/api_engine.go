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

// EngineApi service type
type EngineApi common.Service

// ListServiceVersionOptionalParameters holds optional parameters for ListServiceVersion.
type ListServiceVersionOptionalParameters struct {
	Component *string
}

// NewListServiceVersionOptionalParameters creates an empty struct for parameters.
func NewListServiceVersionOptionalParameters() *ListServiceVersionOptionalParameters {
	this := ListServiceVersionOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *ListServiceVersionOptionalParameters) WithComponent(component string) *ListServiceVersionOptionalParameters {
	r.Component = &component
	return r
}

// ListServiceVersion list the service version of the engine.
// list the service version of the engine
func (a *EngineApi) ListServiceVersion(ctx _context.Context, environmentName string, engineName string, engineMode string, o ...ListServiceVersionOptionalParameters) (EngineServiceVersions, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EngineServiceVersions
		optionalParams      ListServiceVersionOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListServiceVersionOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engine",
		OperationID: "ListServiceVersion",
		Path:        "/api/v1/environments/{environmentName}/engines/{engineName}/serviceVersion",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineApi.ListServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/environments/{environmentName}/engines/{engineName}/serviceVersion"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentName"+"}", _neturl.PathEscape(common.ParameterToString(environmentName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("engineMode", common.ParameterToString(engineMode, ""))
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
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

// ListUpgradeableServiceVersion list upgraded service version of the component.
// list upgraded service version of the component
func (a *EngineApi) ListUpgradeableServiceVersion(ctx _context.Context, clusterName string, orgName string, component string) (EngineServiceVersions, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EngineServiceVersions
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engine",
		OperationID: "ListUpgradeableServiceVersion",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/upgradeableServiceVersion",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineApi.ListUpgradeableServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/upgradeableServiceVersion"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("component", common.ParameterToString(component, ""))
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 {
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

// ListEngineLicensesOptionalParameters holds optional parameters for ListEngineLicenses.
type ListEngineLicensesOptionalParameters struct {
	EngineName *string
}

// NewListEngineLicensesOptionalParameters creates an empty struct for parameters.
func NewListEngineLicensesOptionalParameters() *ListEngineLicensesOptionalParameters {
	this := ListEngineLicensesOptionalParameters{}
	return &this
}

// WithEngineName sets the corresponding parameter name and returns the struct.
func (r *ListEngineLicensesOptionalParameters) WithEngineName(engineName string) *ListEngineLicensesOptionalParameters {
	r.EngineName = &engineName
	return r
}

// ListEngineLicenses List all engineLicenses.
// list all engineLicenses
func (a *EngineApi) ListEngineLicenses(ctx _context.Context, o ...ListEngineLicensesOptionalParameters) (EngineLicenseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EngineLicenseList
		optionalParams      ListEngineLicensesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEngineLicensesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engine",
		OperationID: "listEngineLicenses",
		Path:        "/api/v1/engineLicenses",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineApi.ListEngineLicenses")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/engineLicenses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.EngineName != nil {
		localVarQueryParams.Add("engineName", common.ParameterToString(*optionalParams.EngineName, ""))
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

// ListEngineResourceConstraintsOptionalParameters holds optional parameters for ListEngineResourceConstraints.
type ListEngineResourceConstraintsOptionalParameters struct {
	Engine    *string
	Mode      *string
	Component *string
}

// NewListEngineResourceConstraintsOptionalParameters creates an empty struct for parameters.
func NewListEngineResourceConstraintsOptionalParameters() *ListEngineResourceConstraintsOptionalParameters {
	this := ListEngineResourceConstraintsOptionalParameters{}
	return &this
}

// WithEngine sets the corresponding parameter name and returns the struct.
func (r *ListEngineResourceConstraintsOptionalParameters) WithEngine(engine string) *ListEngineResourceConstraintsOptionalParameters {
	r.Engine = &engine
	return r
}

// WithMode sets the corresponding parameter name and returns the struct.
func (r *ListEngineResourceConstraintsOptionalParameters) WithMode(mode string) *ListEngineResourceConstraintsOptionalParameters {
	r.Mode = &mode
	return r
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *ListEngineResourceConstraintsOptionalParameters) WithComponent(component string) *ListEngineResourceConstraintsOptionalParameters {
	r.Component = &component
	return r
}

// ListEngineResourceConstraints List engine resource constraints.
func (a *EngineApi) ListEngineResourceConstraints(ctx _context.Context, o ...ListEngineResourceConstraintsOptionalParameters) (ResourceConstraintList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ResourceConstraintList
		optionalParams      ListEngineResourceConstraintsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEngineResourceConstraintsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engine",
		OperationID: "listEngineResourceConstraints",
		Path:        "/api/v1/engines/resourceConstraints",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineApi.ListEngineResourceConstraints")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/engines/resourceConstraints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Engine != nil {
		localVarQueryParams.Add("engine", common.ParameterToString(*optionalParams.Engine, ""))
	}
	if optionalParams.Mode != nil {
		localVarQueryParams.Add("mode", common.ParameterToString(*optionalParams.Mode, ""))
	}
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
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

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engine",
		OperationID: "listEnginesInEnv",
		Path:        "/api/v1/environments/{environmentName}/engines",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

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

// NewEngineApi Returns NewEngineApi.
func NewEngineApi(client *common.APIClient) *EngineApi {
	return &EngineApi{
		Client: client,
	}
}
