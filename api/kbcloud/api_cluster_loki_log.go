// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"context"
	_context "context"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterLokiLogApi service type
type ClusterLokiLogApi common.Service

// ExportLokiClusterLogsOptionalParameters holds optional parameters for ExportLokiClusterLogs.
type ExportLokiClusterLogsOptionalParameters struct {
	StartTime     *int64
	EndTime       *int64
	Format        *string
	Query         *string
	Filename      *string
	ComponentName *string
	InstanceName  *string
	MaxLines      *int64
}

// NewExportLokiClusterLogsOptionalParameters creates an empty struct for parameters.
func NewExportLokiClusterLogsOptionalParameters() *ExportLokiClusterLogsOptionalParameters {
	this := ExportLokiClusterLogsOptionalParameters{}
	return &this
}

// WithStartTime sets the corresponding parameter name and returns the struct.
func (r *ExportLokiClusterLogsOptionalParameters) WithStartTime(startTime int64) *ExportLokiClusterLogsOptionalParameters {
	r.StartTime = &startTime
	return r
}

// WithEndTime sets the corresponding parameter name and returns the struct.
func (r *ExportLokiClusterLogsOptionalParameters) WithEndTime(endTime int64) *ExportLokiClusterLogsOptionalParameters {
	r.EndTime = &endTime
	return r
}

// WithFormat sets the corresponding parameter name and returns the struct.
func (r *ExportLokiClusterLogsOptionalParameters) WithFormat(format string) *ExportLokiClusterLogsOptionalParameters {
	r.Format = &format
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *ExportLokiClusterLogsOptionalParameters) WithQuery(query string) *ExportLokiClusterLogsOptionalParameters {
	r.Query = &query
	return r
}

// WithFilename sets the corresponding parameter name and returns the struct.
func (r *ExportLokiClusterLogsOptionalParameters) WithFilename(filename string) *ExportLokiClusterLogsOptionalParameters {
	r.Filename = &filename
	return r
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *ExportLokiClusterLogsOptionalParameters) WithComponentName(componentName string) *ExportLokiClusterLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *ExportLokiClusterLogsOptionalParameters) WithInstanceName(instanceName string) *ExportLokiClusterLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithMaxLines sets the corresponding parameter name and returns the struct.
func (r *ExportLokiClusterLogsOptionalParameters) WithMaxLines(maxLines int64) *ExportLokiClusterLogsOptionalParameters {
	r.MaxLines = &maxLines
	return r
}

// ExportLokiClusterLogs Export Loki logs.
// Export logs through the Loki compatibility API.
func (a *ClusterLokiLogApi) ExportLokiClusterLogs(ctx _context.Context, orgName string, clusterName string, logType string, o ...ExportLokiClusterLogsOptionalParameters) (_io.Reader, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue _io.Reader
		optionalParams      ExportLokiClusterLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ExportLokiClusterLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLokiLog",
		OperationID: "exportLokiClusterLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/export",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLokiLogApi.ExportLokiClusterLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/export"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("logType", common.ParameterToString(logType, ""))
	if optionalParams.StartTime != nil {
		localVarQueryParams.Add("startTime", common.ParameterToString(*optionalParams.StartTime, ""))
	}
	if optionalParams.EndTime != nil {
		localVarQueryParams.Add("endTime", common.ParameterToString(*optionalParams.EndTime, ""))
	}
	if optionalParams.Format != nil {
		localVarQueryParams.Add("format", common.ParameterToString(*optionalParams.Format, ""))
	}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", common.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.Filename != nil {
		localVarQueryParams.Add("filename", common.ParameterToString(*optionalParams.Filename, ""))
	}
	if optionalParams.ComponentName != nil {
		localVarQueryParams.Add("componentName", common.ParameterToString(*optionalParams.ComponentName, ""))
	}
	if optionalParams.InstanceName != nil {
		localVarQueryParams.Add("instanceName", common.ParameterToString(*optionalParams.InstanceName, ""))
	}
	if optionalParams.MaxLines != nil {
		localVarQueryParams.Add("maxLines", common.ParameterToString(*optionalParams.MaxLines, ""))
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

	if localVarHTTPResponse.StatusCode >= 300 {

		localVarBody, err := common.ReadBody(localVarHTTPResponse)
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}
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
	localVarReturnValue = localVarHTTPResponse.Body

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetLokiSlowLogStatsOptionalParameters holds optional parameters for GetLokiSlowLogStats.
type GetLokiSlowLogStatsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
}

// NewGetLokiSlowLogStatsOptionalParameters creates an empty struct for parameters.
func NewGetLokiSlowLogStatsOptionalParameters() *GetLokiSlowLogStatsOptionalParameters {
	this := GetLokiSlowLogStatsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *GetLokiSlowLogStatsOptionalParameters) WithComponentName(componentName string) *GetLokiSlowLogStatsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *GetLokiSlowLogStatsOptionalParameters) WithInstanceName(instanceName string) *GetLokiSlowLogStatsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// GetLokiSlowLogStats Query Loki slow-log statistics.
// Query slow-log statistics through the Loki compatibility API.
func (a *ClusterLokiLogApi) GetLokiSlowLogStats(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...GetLokiSlowLogStatsOptionalParameters) (ClusterSlowLogStats, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterSlowLogStats
		optionalParams      GetLokiSlowLogStatsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetLokiSlowLogStatsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLokiLog",
		OperationID: "getLokiSlowLogStats",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/slow/stats",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLokiLogApi.GetLokiSlowLogStats")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/slow/stats"
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

// QueryLokiAuditLogsOptionalParameters holds optional parameters for QueryLokiAuditLogs.
type QueryLokiAuditLogsOptionalParameters struct {
	Limit         *string
	ComponentName *string
	InstanceName  *string
	SortType      *SortType
}

// NewQueryLokiAuditLogsOptionalParameters creates an empty struct for parameters.
func NewQueryLokiAuditLogsOptionalParameters() *QueryLokiAuditLogsOptionalParameters {
	this := QueryLokiAuditLogsOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryLokiAuditLogsOptionalParameters) WithLimit(limit string) *QueryLokiAuditLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiAuditLogsOptionalParameters) WithComponentName(componentName string) *QueryLokiAuditLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiAuditLogsOptionalParameters) WithInstanceName(instanceName string) *QueryLokiAuditLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryLokiAuditLogsOptionalParameters) WithSortType(sortType SortType) *QueryLokiAuditLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryLokiAuditLogs Query Loki audit logs.
// Query audit logs through the Loki compatibility API.
func (a *ClusterLokiLogApi) QueryLokiAuditLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryLokiAuditLogsOptionalParameters) (ClusterExecutionLog, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterExecutionLog
		optionalParams      QueryLokiAuditLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryLokiAuditLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLokiLog",
		OperationID: "queryLokiAuditLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/audit",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLokiLogApi.QueryLokiAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/audit"
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

// QueryLokiErrorLogsOptionalParameters holds optional parameters for QueryLokiErrorLogs.
type QueryLokiErrorLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Filename      *string
	Query         *string
	Limit         *string
	SortType      *SortType
}

// NewQueryLokiErrorLogsOptionalParameters creates an empty struct for parameters.
func NewQueryLokiErrorLogsOptionalParameters() *QueryLokiErrorLogsOptionalParameters {
	this := QueryLokiErrorLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiErrorLogsOptionalParameters) WithComponentName(componentName string) *QueryLokiErrorLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiErrorLogsOptionalParameters) WithInstanceName(instanceName string) *QueryLokiErrorLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithFilename sets the corresponding parameter name and returns the struct.
func (r *QueryLokiErrorLogsOptionalParameters) WithFilename(filename string) *QueryLokiErrorLogsOptionalParameters {
	r.Filename = &filename
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *QueryLokiErrorLogsOptionalParameters) WithQuery(query string) *QueryLokiErrorLogsOptionalParameters {
	r.Query = &query
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryLokiErrorLogsOptionalParameters) WithLimit(limit string) *QueryLokiErrorLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryLokiErrorLogsOptionalParameters) WithSortType(sortType SortType) *QueryLokiErrorLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryLokiErrorLogs Query Loki error logs.
// Query error logs through the Loki compatibility API.
func (a *ClusterLokiLogApi) QueryLokiErrorLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryLokiErrorLogsOptionalParameters) (ClusterRawLogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterRawLogResponse
		optionalParams      QueryLokiErrorLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryLokiErrorLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLokiLog",
		OperationID: "queryLokiErrorLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/error",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLokiLogApi.QueryLokiErrorLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/error"
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

// QueryLokiPodLogsOptionalParameters holds optional parameters for QueryLokiPodLogs.
type QueryLokiPodLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Filename      *string
	Limit         *string
	SortType      *SortType
}

// NewQueryLokiPodLogsOptionalParameters creates an empty struct for parameters.
func NewQueryLokiPodLogsOptionalParameters() *QueryLokiPodLogsOptionalParameters {
	this := QueryLokiPodLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiPodLogsOptionalParameters) WithComponentName(componentName string) *QueryLokiPodLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiPodLogsOptionalParameters) WithInstanceName(instanceName string) *QueryLokiPodLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithFilename sets the corresponding parameter name and returns the struct.
func (r *QueryLokiPodLogsOptionalParameters) WithFilename(filename string) *QueryLokiPodLogsOptionalParameters {
	r.Filename = &filename
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryLokiPodLogsOptionalParameters) WithLimit(limit string) *QueryLokiPodLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryLokiPodLogsOptionalParameters) WithSortType(sortType SortType) *QueryLokiPodLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryLokiPodLogs Query Loki pod logs.
// Query pod logs through the Loki compatibility API.
func (a *ClusterLokiLogApi) QueryLokiPodLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryLokiPodLogsOptionalParameters) (ClusterRawLogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterRawLogResponse
		optionalParams      QueryLokiPodLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryLokiPodLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLokiLog",
		OperationID: "queryLokiPodLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/pod",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLokiLogApi.QueryLokiPodLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/pod"
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

// QueryLokiRunningLogsOptionalParameters holds optional parameters for QueryLokiRunningLogs.
type QueryLokiRunningLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Filename      *string
	Limit         *string
	Query         *string
	SortType      *SortType
}

// NewQueryLokiRunningLogsOptionalParameters creates an empty struct for parameters.
func NewQueryLokiRunningLogsOptionalParameters() *QueryLokiRunningLogsOptionalParameters {
	this := QueryLokiRunningLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiRunningLogsOptionalParameters) WithComponentName(componentName string) *QueryLokiRunningLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiRunningLogsOptionalParameters) WithInstanceName(instanceName string) *QueryLokiRunningLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithFilename sets the corresponding parameter name and returns the struct.
func (r *QueryLokiRunningLogsOptionalParameters) WithFilename(filename string) *QueryLokiRunningLogsOptionalParameters {
	r.Filename = &filename
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryLokiRunningLogsOptionalParameters) WithLimit(limit string) *QueryLokiRunningLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *QueryLokiRunningLogsOptionalParameters) WithQuery(query string) *QueryLokiRunningLogsOptionalParameters {
	r.Query = &query
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryLokiRunningLogsOptionalParameters) WithSortType(sortType SortType) *QueryLokiRunningLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryLokiRunningLogs Query Loki running logs.
// Query running logs through the Loki compatibility API.
func (a *ClusterLokiLogApi) QueryLokiRunningLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryLokiRunningLogsOptionalParameters) (ClusterRawLogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterRawLogResponse
		optionalParams      QueryLokiRunningLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryLokiRunningLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLokiLog",
		OperationID: "queryLokiRunningLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/running",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLokiLogApi.QueryLokiRunningLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/running"
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

// QueryLokiSlowLogsOptionalParameters holds optional parameters for QueryLokiSlowLogs.
type QueryLokiSlowLogsOptionalParameters struct {
	ComponentName *string
	InstanceName  *string
	Query         *string
	Limit         *string
	SortType      *SortType
}

// NewQueryLokiSlowLogsOptionalParameters creates an empty struct for parameters.
func NewQueryLokiSlowLogsOptionalParameters() *QueryLokiSlowLogsOptionalParameters {
	this := QueryLokiSlowLogsOptionalParameters{}
	return &this
}

// WithComponentName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiSlowLogsOptionalParameters) WithComponentName(componentName string) *QueryLokiSlowLogsOptionalParameters {
	r.ComponentName = &componentName
	return r
}

// WithInstanceName sets the corresponding parameter name and returns the struct.
func (r *QueryLokiSlowLogsOptionalParameters) WithInstanceName(instanceName string) *QueryLokiSlowLogsOptionalParameters {
	r.InstanceName = &instanceName
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *QueryLokiSlowLogsOptionalParameters) WithQuery(query string) *QueryLokiSlowLogsOptionalParameters {
	r.Query = &query
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryLokiSlowLogsOptionalParameters) WithLimit(limit string) *QueryLokiSlowLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryLokiSlowLogsOptionalParameters) WithSortType(sortType SortType) *QueryLokiSlowLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryLokiSlowLogs Query Loki slow logs.
// Query slow logs through the Loki compatibility API.
func (a *ClusterLokiLogApi) QueryLokiSlowLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryLokiSlowLogsOptionalParameters) (ClusterExecutionLog, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterExecutionLog
		optionalParams      QueryLokiSlowLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryLokiSlowLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "clusterLokiLog",
		OperationID: "queryLokiSlowLogs",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/slow",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLokiLogApi.QueryLokiSlowLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/logs/loki/slow"
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

// NewClusterLokiLogApi Returns NewClusterLokiLogApi.
func NewClusterLokiLogApi(client *common.APIClient) *ClusterLokiLogApi {
	return &ClusterLokiLogApi{
		Client: client,
	}
}
