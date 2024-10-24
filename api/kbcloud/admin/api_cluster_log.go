// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
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

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryAuditLogsOptionalParameters) WithSortType(sortType SortType) *QueryAuditLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryAuditLogs Query cluster audit logs.
// Query audit logs of a cluster
// NODESCRIPTION QueryAuditLogs
// Deprecated: This API is deprecated.
func (a *ClusterLogApi) QueryAuditLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryAuditLogsOptionalParameters) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue interface{}
		optionalParams      QueryAuditLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryAuditLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QueryAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/logs/audit"
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
// NODESCRIPTION QueryErrorLogs
// Deprecated: This API is deprecated.
func (a *ClusterLogApi) QueryErrorLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryErrorLogsOptionalParameters) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue interface{}
		optionalParams      QueryErrorLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryErrorLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QueryErrorLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/logs/error"
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
	Limit         *string
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

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *QueryRunningLogsOptionalParameters) WithLimit(limit string) *QueryRunningLogsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithSortType sets the corresponding parameter name and returns the struct.
func (r *QueryRunningLogsOptionalParameters) WithSortType(sortType SortType) *QueryRunningLogsOptionalParameters {
	r.SortType = &sortType
	return r
}

// QueryRunningLogs Query cluster running logs.
// Query running logs of a cluster
// NODESCRIPTION QueryRunningLogs
// Deprecated: This API is deprecated.
func (a *ClusterLogApi) QueryRunningLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QueryRunningLogsOptionalParameters) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue interface{}
		optionalParams      QueryRunningLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryRunningLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QueryRunningLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/logs/running"
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
// NODESCRIPTION QuerySlowLogs
// Deprecated: This API is deprecated.
func (a *ClusterLogApi) QuerySlowLogs(ctx _context.Context, orgName string, clusterName string, startTime string, endTime string, o ...QuerySlowLogsOptionalParameters) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue interface{}
		optionalParams      QuerySlowLogsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QuerySlowLogsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ClusterLogApi.QuerySlowLogs")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/logs/slow"
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
