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

// ClusterLogApi service type
type ClusterLogApi common.Service

// QueryAuditLogsOptionalParameters holds optional parameters for QueryAuditLogs.
type QueryAuditLogsOptionalParameters struct {
	Limit         *string
	ComponentName *string
	InstanceName  *string
	SortType      *SortType
}

// NewQueryAuditLogsOptionalParameters creates an empty struct for parameters.
func NewQueryAuditLogsOptionalParameters() *QueryAuditLogsOptionalParameters {
	this := QueryAuditLogsOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryAuditLogsOptionalParameters) WithLimit(limit string) *QueryAuditLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryAuditLogsOptionalParameters) WithComponentName(componentName string) *QueryAuditLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryAuditLogsOptionalParameters) WithInstanceName(instanceName string) *QueryAuditLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryAuditLogsOptionalParameters) WithSortType(sortType SortType) *QueryAuditLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryAuditLogs Query cluster audit logs.
// Query audit logs of a cluster
func (a *ClusterLogApi) QueryAuditLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryAuditLogsOptionalParameters) (ClusterExecutionLog, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterExecutionLog
		optionalParams      QueryAuditLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryAuditLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLog",
		OperationID: "queryAuditLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/audit",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QueryAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/audit"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("startTime", common.ParameterToString(startTime, ""))
	localVarQueryParams.Add("endTime", common.ParameterToString(endTime, ""))
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.ComponentName != nil {
		localVarQueryParams.Add("componentName", common.ParameterToString(*optionalParams.ComponentName, ""))
	}
	if optionalParams.InstanceName != nil {
		localVarQueryParams.Add("instanceName", common.ParameterToString(*optionalParams.InstanceName, ""))
	}
	if optionalParams.SortType != nil {
		localVarQueryParams.Add("sortType", common.ParameterToString(*optionalParams.SortType, ""))
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

// QueryErrorLogsOptionalParameters holds optional parameters for QueryErrorLogs.
type QueryErrorLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Filename      *string
	Query         *string
	Limit         *string
	SortType      *SortType
}

// NewQueryErrorLogsOptionalParameters creates an empty struct for parameters.
func NewQueryErrorLogsOptionalParameters() *QueryErrorLogsOptionalParameters {
	this := QueryErrorLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryErrorLogsOptionalParameters) WithComponentName(componentName string) *QueryErrorLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryErrorLogsOptionalParameters) WithInstanceName(instanceName string) *QueryErrorLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithFilename sets the corresponding parameter name and returns the struct.
func (r *QueryErrorLogsOptionalParameters) WithFilename(filename string) *QueryErrorLogsOptionalParameters {
	r.Filename = &filename
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *QueryErrorLogsOptionalParameters) WithQuery(query string) *QueryErrorLogsOptionalParameters {
	r.Query = &query
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryErrorLogsOptionalParameters) WithLimit(limit string) *QueryErrorLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryErrorLogsOptionalParameters) WithSortType(sortType SortType) *QueryErrorLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryErrorLogs Query cluster error logs.
// Query error logs of a cluster
func (a *ClusterLogApi) QueryErrorLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryErrorLogsOptionalParameters) (ClusterRawLogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterRawLogResponse
		optionalParams      QueryErrorLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryErrorLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLog",
		OperationID: "queryErrorLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/error",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QueryErrorLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/error"
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
	if optionalParams.Filename != nil {
		localVarQueryParams.Add("filename", common.ParameterToString(*optionalParams.Filename, ""))
	}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", common.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.SortType != nil {
		localVarQueryParams.Add("sortType", common.ParameterToString(*optionalParams.SortType, ""))
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

// QueryPodLogsOptionalParameters holds optional parameters for QueryPodLogs.
type QueryPodLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Filename      *string
	Limit         *string
	SortType      *SortType
}

// NewQueryPodLogsOptionalParameters creates an empty struct for parameters.
func NewQueryPodLogsOptionalParameters() *QueryPodLogsOptionalParameters {
	this := QueryPodLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryPodLogsOptionalParameters) WithComponentName(componentName string) *QueryPodLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryPodLogsOptionalParameters) WithInstanceName(instanceName string) *QueryPodLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithFilename sets the corresponding parameter name and returns the struct.
func (r *QueryPodLogsOptionalParameters) WithFilename(filename string) *QueryPodLogsOptionalParameters {
	r.Filename = &filename
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryPodLogsOptionalParameters) WithLimit(limit string) *QueryPodLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryPodLogsOptionalParameters) WithSortType(sortType SortType) *QueryPodLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryPodLogs Query cluster pod logs.
// Query pod logs of a cluster
func (a *ClusterLogApi) QueryPodLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryPodLogsOptionalParameters) (ClusterRawLogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterRawLogResponse
		optionalParams      QueryPodLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryPodLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLog",
		OperationID: "queryPodLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/pod",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QueryPodLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/pod"
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
	if optionalParams.Filename != nil {
		localVarQueryParams.Add("filename", common.ParameterToString(*optionalParams.Filename, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.SortType != nil {
		localVarQueryParams.Add("sortType", common.ParameterToString(*optionalParams.SortType, ""))
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

// QueryRunningLogsOptionalParameters holds optional parameters for QueryRunningLogs.
type QueryRunningLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Filename      *string
	Limit         *string
	Query         *string
	SortType      *SortType
}

// NewQueryRunningLogsOptionalParameters creates an empty struct for parameters.
func NewQueryRunningLogsOptionalParameters() *QueryRunningLogsOptionalParameters {
	this := QueryRunningLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryRunningLogsOptionalParameters) WithComponentName(componentName string) *QueryRunningLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryRunningLogsOptionalParameters) WithInstanceName(instanceName string) *QueryRunningLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithFilename sets the corresponding parameter name and returns the struct.
func (r *QueryRunningLogsOptionalParameters) WithFilename(filename string) *QueryRunningLogsOptionalParameters {
	r.Filename = &filename
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryRunningLogsOptionalParameters) WithLimit(limit string) *QueryRunningLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *QueryRunningLogsOptionalParameters) WithQuery(query string) *QueryRunningLogsOptionalParameters {
	r.Query = &query
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryRunningLogsOptionalParameters) WithSortType(sortType SortType) *QueryRunningLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryRunningLogs Query cluster running logs.
// Query running logs of a cluster
func (a *ClusterLogApi) QueryRunningLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryRunningLogsOptionalParameters) (ClusterRawLogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterRawLogResponse
		optionalParams      QueryRunningLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryRunningLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLog",
		OperationID: "queryRunningLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/running",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QueryRunningLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/running"
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
	if optionalParams.Filename != nil {
		localVarQueryParams.Add("filename", common.ParameterToString(*optionalParams.Filename, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", common.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.SortType != nil {
		localVarQueryParams.Add("sortType", common.ParameterToString(*optionalParams.SortType, ""))
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

// QuerySlowLogsOptionalParameters holds optional parameters for QuerySlowLogs.
type QuerySlowLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Query         *string
	Limit         *string
	SortType      *SortType
}

// NewQuerySlowLogsOptionalParameters creates an empty struct for parameters.
func NewQuerySlowLogsOptionalParameters() *QuerySlowLogsOptionalParameters {
	this := QuerySlowLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QuerySlowLogsOptionalParameters) WithComponentName(componentName string) *QuerySlowLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QuerySlowLogsOptionalParameters) WithInstanceName(instanceName string) *QuerySlowLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *QuerySlowLogsOptionalParameters) WithQuery(query string) *QuerySlowLogsOptionalParameters {
	r.Query = &query
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QuerySlowLogsOptionalParameters) WithLimit(limit string) *QuerySlowLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QuerySlowLogsOptionalParameters) WithSortType(sortType SortType) *QuerySlowLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QuerySlowLogs Query cluster slow logs.
// Query slow logs of a cluster
func (a *ClusterLogApi) QuerySlowLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QuerySlowLogsOptionalParameters) (ClusterExecutionLog, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterExecutionLog
		optionalParams      QuerySlowLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QuerySlowLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLog",
		OperationID: "querySlowLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/slow",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QuerySlowLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/slow"
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
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", common.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.SortType != nil {
		localVarQueryParams.Add("sortType", common.ParameterToString(*optionalParams.SortType, ""))
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

// NewClusterLogApi Returns NewClusterLogApi.
func NewClusterLogApi(client *common.APIClient) *ClusterLogApi {
	return &ClusterLogApi{
		Client: client,
	}
}
