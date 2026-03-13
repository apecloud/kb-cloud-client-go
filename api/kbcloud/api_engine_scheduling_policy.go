// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"context"
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineSchedulingPolicyApi service type
type EngineSchedulingPolicyApi common.Service

// ListEngineGlobalSchedulingStrategiesOptionalParameters holds optional parameters for ListEngineGlobalSchedulingStrategies.
type ListEngineGlobalSchedulingStrategiesOptionalParameters struct {
	EngineName *string
}

// NewListEngineGlobalSchedulingStrategiesOptionalParameters creates an empty struct for parameters.
func NewListEngineGlobalSchedulingStrategiesOptionalParameters() *ListEngineGlobalSchedulingStrategiesOptionalParameters {
	this := ListEngineGlobalSchedulingStrategiesOptionalParameters{}
	return &this
}

// WithEngineName sets the corresponding parameter name and returns the struct.
func (r *ListEngineGlobalSchedulingStrategiesOptionalParameters) WithEngineName(engineName string) *ListEngineGlobalSchedulingStrategiesOptionalParameters {
	r.EngineName = &engineName
	return r
}

// ListEngineGlobalSchedulingStrategies List engine global scheduling strategies.
func (a *EngineSchedulingPolicyApi) ListEngineGlobalSchedulingStrategies(ctx _context.Context, o ...ListEngineGlobalSchedulingStrategiesOptionalParameters) ([]EngineGlobalSchedulingStrategy, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []EngineGlobalSchedulingStrategy
		optionalParams      ListEngineGlobalSchedulingStrategiesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEngineGlobalSchedulingStrategiesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engineSchedulingPolicy",
		OperationID: "listEngineGlobalSchedulingStrategies",
		Path:        "/api/v1/engineGlobalSchedulingStrategies",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineSchedulingPolicyApi.ListEngineGlobalSchedulingStrategies")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/engineGlobalSchedulingStrategies"

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

// ListEngineSchedulingRulesOptionalParameters holds optional parameters for ListEngineSchedulingRules.
type ListEngineSchedulingRulesOptionalParameters struct {
	EngineName *string
}

// NewListEngineSchedulingRulesOptionalParameters creates an empty struct for parameters.
func NewListEngineSchedulingRulesOptionalParameters() *ListEngineSchedulingRulesOptionalParameters {
	this := ListEngineSchedulingRulesOptionalParameters{}
	return &this
}

// WithEngineName sets the corresponding parameter name and returns the struct.
func (r *ListEngineSchedulingRulesOptionalParameters) WithEngineName(engineName string) *ListEngineSchedulingRulesOptionalParameters {
	r.EngineName = &engineName
	return r
}

// ListEngineSchedulingRules List engine scheduling rules.
func (a *EngineSchedulingPolicyApi) ListEngineSchedulingRules(ctx _context.Context, o ...ListEngineSchedulingRulesOptionalParameters) ([]EngineSchedulingRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []EngineSchedulingRule
		optionalParams      ListEngineSchedulingRulesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEngineSchedulingRulesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "engineSchedulingPolicy",
		OperationID: "listEngineSchedulingRules",
		Path:        "/api/v1/engineSchedulingRules",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineSchedulingPolicyApi.ListEngineSchedulingRules")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/engineSchedulingRules"

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

// NewEngineSchedulingPolicyApi Returns NewEngineSchedulingPolicyApi.
func NewEngineSchedulingPolicyApi(client *common.APIClient) *EngineSchedulingPolicyApi {
	return &EngineSchedulingPolicyApi{
		Client: client,
	}
}
