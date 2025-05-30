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

// AnalyzeApi service type
type AnalyzeApi common.Service

// AnalyzeBackupOptionalParameters holds optional parameters for AnalyzeBackup.
type AnalyzeBackupOptionalParameters struct {
	Model *string
}

// NewAnalyzeBackupOptionalParameters creates an empty struct for parameters.
func NewAnalyzeBackupOptionalParameters() *AnalyzeBackupOptionalParameters {
	this := AnalyzeBackupOptionalParameters{}
	return &this
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeBackupOptionalParameters) WithModel(model string) *AnalyzeBackupOptionalParameters {
	r.Model = &model
	return r
}

// AnalyzeBackup Analyze backup.
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeBackup(ctx _context.Context, orgName string, backupId string, o ...AnalyzeBackupOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeBackupOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeBackupOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeBackup",
		Path:        "/admin/v1/organizations/{orgName}/backups/{backupId}/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeBackup")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/backups/{backupId}/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"backupId"+"}", _neturl.PathEscape(common.ParameterToString(backupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
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

// AnalyzeClusterParamOptionalParameters holds optional parameters for AnalyzeClusterParam.
type AnalyzeClusterParamOptionalParameters struct {
	Model          *string
	ParameterValue *string
}

// NewAnalyzeClusterParamOptionalParameters creates an empty struct for parameters.
func NewAnalyzeClusterParamOptionalParameters() *AnalyzeClusterParamOptionalParameters {
	this := AnalyzeClusterParamOptionalParameters{}
	return &this
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeClusterParamOptionalParameters) WithModel(model string) *AnalyzeClusterParamOptionalParameters {
	r.Model = &model
	return r
}

// WithParameterValue sets the corresponding parameter name and returns the struct.
func (r *AnalyzeClusterParamOptionalParameters) WithParameterValue(parameterValue string) *AnalyzeClusterParamOptionalParameters {
	r.ParameterValue = &parameterValue
	return r
}

// AnalyzeClusterParam Analyze cluster parameter.
// analyze cluster parameter, deprecated, instead use analyzeClusterParameter
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeClusterParam(ctx _context.Context, orgName string, clusterName string, parameterName string, o ...AnalyzeClusterParamOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeClusterParamOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeClusterParamOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeClusterParam",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/params/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeClusterParam")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/params/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("parameterName", common.ParameterToString(parameterName, ""))
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
	}
	if optionalParams.ParameterValue != nil {
		localVarQueryParams.Add("parameterValue", common.ParameterToString(*optionalParams.ParameterValue, ""))
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

// AnalyzeClusterParameterOptionalParameters holds optional parameters for AnalyzeClusterParameter.
type AnalyzeClusterParameterOptionalParameters struct {
	Model          *string
	ParameterValue *string
}

// NewAnalyzeClusterParameterOptionalParameters creates an empty struct for parameters.
func NewAnalyzeClusterParameterOptionalParameters() *AnalyzeClusterParameterOptionalParameters {
	this := AnalyzeClusterParameterOptionalParameters{}
	return &this
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeClusterParameterOptionalParameters) WithModel(model string) *AnalyzeClusterParameterOptionalParameters {
	r.Model = &model
	return r
}

// WithParameterValue sets the corresponding parameter name and returns the struct.
func (r *AnalyzeClusterParameterOptionalParameters) WithParameterValue(parameterValue string) *AnalyzeClusterParameterOptionalParameters {
	r.ParameterValue = &parameterValue
	return r
}

// AnalyzeClusterParameter Analyze cluster parameter.
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeClusterParameter(ctx _context.Context, orgName string, clusterName string, parameterName string, o ...AnalyzeClusterParameterOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeClusterParameterOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeClusterParameterOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeClusterParameter",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/parameter/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeClusterParameter")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/parameter/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("parameterName", common.ParameterToString(parameterName, ""))
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
	}
	if optionalParams.ParameterValue != nil {
		localVarQueryParams.Add("parameterValue", common.ParameterToString(*optionalParams.ParameterValue, ""))
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

// AnalyzeClusterRestoreOptionalParameters holds optional parameters for AnalyzeClusterRestore.
type AnalyzeClusterRestoreOptionalParameters struct {
	Model *string
}

// NewAnalyzeClusterRestoreOptionalParameters creates an empty struct for parameters.
func NewAnalyzeClusterRestoreOptionalParameters() *AnalyzeClusterRestoreOptionalParameters {
	this := AnalyzeClusterRestoreOptionalParameters{}
	return &this
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeClusterRestoreOptionalParameters) WithModel(model string) *AnalyzeClusterRestoreOptionalParameters {
	r.Model = &model
	return r
}

// AnalyzeClusterRestore Analyze cluster restore tasks.
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeClusterRestore(ctx _context.Context, orgName string, clusterName string, o ...AnalyzeClusterRestoreOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeClusterRestoreOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeClusterRestoreOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeClusterRestore",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/restore/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeClusterRestore")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/restore/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
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

// AnalyzeLogsOptionalParameters holds optional parameters for AnalyzeLogs.
type AnalyzeLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Limit         *string
	SortType      *SortType
	Model         *string
}

// NewAnalyzeLogsOptionalParameters creates an empty struct for parameters.
func NewAnalyzeLogsOptionalParameters() *AnalyzeLogsOptionalParameters {
	this := AnalyzeLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *AnalyzeLogsOptionalParameters) WithComponentName(componentName string) *AnalyzeLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *AnalyzeLogsOptionalParameters) WithInstanceName(instanceName string) *AnalyzeLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *AnalyzeLogsOptionalParameters) WithLimit(limit string) *AnalyzeLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *AnalyzeLogsOptionalParameters) WithSortType(sortType SortType) *AnalyzeLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeLogsOptionalParameters) WithModel(model string) *AnalyzeLogsOptionalParameters {
	r.Model = &model
	return r
}

// AnalyzeLogs Analyze cluster error logs.
// Analyze error logs of a cluster
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...AnalyzeLogsOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeLogs",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/logs/error/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/logs/error/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("startTime", common.ParameterToString(startTime, ""))
	localVarQueryParams.Add("endTime", common.ParameterToString(endTime, ""))
	if optionalParams.ComponentName != nil {
		localVarQueryParams.Add("componentName", common.ParameterToString(*optionalParams.ComponentName, ""))
	}
	if optionalParams.InstanceName != nil {
		localVarQueryParams.Add("instanceName", common.ParameterToString(*optionalParams.InstanceName, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.SortType != nil {
		localVarQueryParams.Add("sortType", common.ParameterToString(*optionalParams.SortType, ""))
	}
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
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

// AnalyzeOpsOptionalParameters holds optional parameters for AnalyzeOps.
type AnalyzeOpsOptionalParameters struct {
	Model *string
}

// NewAnalyzeOpsOptionalParameters creates an empty struct for parameters.
func NewAnalyzeOpsOptionalParameters() *AnalyzeOpsOptionalParameters {
	this := AnalyzeOpsOptionalParameters{}
	return &this
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeOpsOptionalParameters) WithModel(model string) *AnalyzeOpsOptionalParameters {
	r.Model = &model
	return r
}

// AnalyzeOps Analyze OpsRequest.
// analyze a OpsRequest
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeOps(ctx _context.Context, orgName string, opsName string, clusterName string, opsType string, o ...AnalyzeOpsOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeOpsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeOpsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeOps",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/opsrequests/{opsName}/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeOps")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/opsrequests/{opsName}/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"opsName"+"}", _neturl.PathEscape(common.ParameterToString(opsName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("opsType", common.ParameterToString(opsType, ""))
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
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

// AnalyzeParamOptionalParameters holds optional parameters for AnalyzeParam.
type AnalyzeParamOptionalParameters struct {
	Mode      *string
	OrgName   *string
	Partition *ParamTplPartition
}

// NewAnalyzeParamOptionalParameters creates an empty struct for parameters.
func NewAnalyzeParamOptionalParameters() *AnalyzeParamOptionalParameters {
	this := AnalyzeParamOptionalParameters{}
	return &this
}

// WithMode sets the corresponding parameter name and returns the struct.
func (r *AnalyzeParamOptionalParameters) WithMode(mode string) *AnalyzeParamOptionalParameters {
	r.Mode = &mode
	return r
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *AnalyzeParamOptionalParameters) WithOrgName(orgName string) *AnalyzeParamOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithPartition sets the corresponding parameter name and returns the struct.
func (r *AnalyzeParamOptionalParameters) WithPartition(partition ParamTplPartition) *AnalyzeParamOptionalParameters {
	r.Partition = &partition
	return r
}

// AnalyzeParam Analyze parameter.
// analyze parameter, deprecated, instead use analyzeParameter
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeParam(ctx _context.Context, paramTplName string, parameterName string, o ...AnalyzeParamOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeParamOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeParamOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeParam",
		Path:        "/admin/v1/organizations/paramTpls/{paramTplName}/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeParam")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/paramTpls/{paramTplName}/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"paramTplName"+"}", _neturl.PathEscape(common.ParameterToString(paramTplName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("parameterName", common.ParameterToString(parameterName, ""))
	if optionalParams.Mode != nil {
		localVarQueryParams.Add("mode", common.ParameterToString(*optionalParams.Mode, ""))
	}
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.Partition != nil {
		localVarQueryParams.Add("partition", common.ParameterToString(*optionalParams.Partition, ""))
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

// AnalyzeParameterOptionalParameters holds optional parameters for AnalyzeParameter.
type AnalyzeParameterOptionalParameters struct {
	Model     *string
	OrgName   *string
	Partition *ParamTplPartition
}

// NewAnalyzeParameterOptionalParameters creates an empty struct for parameters.
func NewAnalyzeParameterOptionalParameters() *AnalyzeParameterOptionalParameters {
	this := AnalyzeParameterOptionalParameters{}
	return &this
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeParameterOptionalParameters) WithModel(model string) *AnalyzeParameterOptionalParameters {
	r.Model = &model
	return r
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *AnalyzeParameterOptionalParameters) WithOrgName(orgName string) *AnalyzeParameterOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithPartition sets the corresponding parameter name and returns the struct.
func (r *AnalyzeParameterOptionalParameters) WithPartition(partition ParamTplPartition) *AnalyzeParameterOptionalParameters {
	r.Partition = &partition
	return r
}

// AnalyzeParameter Analyze parameter.
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeParameter(ctx _context.Context, parameterTemplateName string, parameterName string, o ...AnalyzeParameterOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeParameterOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeParameterOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeParameter",
		Path:        "/admin/v1/organizations/parameterTemplate/{parameterTemplateName}/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeParameter")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/parameterTemplate/{parameterTemplateName}/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"parameterTemplateName"+"}", _neturl.PathEscape(common.ParameterToString(parameterTemplateName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("parameterName", common.ParameterToString(parameterName, ""))
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
	}
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.Partition != nil {
		localVarQueryParams.Add("partition", common.ParameterToString(*optionalParams.Partition, ""))
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

// AnalyzeServiceOptionalParameters holds optional parameters for AnalyzeService.
type AnalyzeServiceOptionalParameters struct {
	PortName  *string
	Endpoints *string
	Model     *string
}

// NewAnalyzeServiceOptionalParameters creates an empty struct for parameters.
func NewAnalyzeServiceOptionalParameters() *AnalyzeServiceOptionalParameters {
	this := AnalyzeServiceOptionalParameters{}
	return &this
}

// WithPortName sets the corresponding parameter name and returns the struct.
func (r *AnalyzeServiceOptionalParameters) WithPortName(portName string) *AnalyzeServiceOptionalParameters {
	r.PortName = &portName
	return r
}

// WithEndpoints sets the corresponding parameter name and returns the struct.
func (r *AnalyzeServiceOptionalParameters) WithEndpoints(endpoints string) *AnalyzeServiceOptionalParameters {
	r.Endpoints = &endpoints
	return r
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeServiceOptionalParameters) WithModel(model string) *AnalyzeServiceOptionalParameters {
	r.Model = &model
	return r
}

// AnalyzeService Analyze service.
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeService(ctx _context.Context, orgName string, clusterName string, serviceName string, o ...AnalyzeServiceOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeServiceOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeServiceOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeService",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/services/{serviceName}/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeService")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/services/{serviceName}/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceName"+"}", _neturl.PathEscape(common.ParameterToString(serviceName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PortName != nil {
		localVarQueryParams.Add("portName", common.ParameterToString(*optionalParams.PortName, ""))
	}
	if optionalParams.Endpoints != nil {
		localVarQueryParams.Add("endpoints", common.ParameterToString(*optionalParams.Endpoints, ""))
	}
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 {
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

// AnalyzeSlowLogsOptionalParameters holds optional parameters for AnalyzeSlowLogs.
type AnalyzeSlowLogsOptionalParameters struct {
	Body  *ClusterExecutionLog
	Model *string
}

// NewAnalyzeSlowLogsOptionalParameters creates an empty struct for parameters.
func NewAnalyzeSlowLogsOptionalParameters() *AnalyzeSlowLogsOptionalParameters {
	this := AnalyzeSlowLogsOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *AnalyzeSlowLogsOptionalParameters) WithBody(body ClusterExecutionLog) *AnalyzeSlowLogsOptionalParameters {
	r.Body = &body
	return r
}

// WithModel sets the corresponding parameter name and returns the struct.
func (r *AnalyzeSlowLogsOptionalParameters) WithModel(model string) *AnalyzeSlowLogsOptionalParameters {
	r.Model = &model
	return r
}

// AnalyzeSlowLogs Analyze cluster slow logs.
// Analyze slow logs of a cluster
// Deprecated: This API is deprecated.
func (a *AnalyzeApi) AnalyzeSlowLogs(ctx _context.Context, orgName string, clusterName string, o ...AnalyzeSlowLogsOptionalParameters) (AnalysisResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue AnalysisResult
		optionalParams      AnalyzeSlowLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type AnalyzeSlowLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "analyze",
		OperationID: "analyzeSlowLogs",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/logs/slow/analyze",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AnalyzeApi.AnalyzeSlowLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/logs/slow/analyze"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Model != nil {
		localVarQueryParams.Add("model", common.ParameterToString(*optionalParams.Model, ""))
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	if optionalParams.Body != nil {
		localVarPostBody = &optionalParams.Body
	}
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

// NewAnalyzeApi Returns NewAnalyzeApi.
func NewAnalyzeApi(client *common.APIClient) *AnalyzeApi {
	return &AnalyzeApi{
		Client: client,
	}
}
