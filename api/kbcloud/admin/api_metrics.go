// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"bytes"
	_context "context"
	_fmt "fmt"
	_io "io"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MetricsApi service type
type MetricsApi common.Service

// GetAggregateMetaDataOptionalParameters holds optional parameters for GetAggregateMetaData.
type GetAggregateMetaDataOptionalParameters struct {
	Start *int64
	End *int64
}

// NewGetAggregateMetaDataOptionalParameters creates an empty struct for parameters.
func NewGetAggregateMetaDataOptionalParameters() *GetAggregateMetaDataOptionalParameters {
	this := GetAggregateMetaDataOptionalParameters{}
	return &this
}
// WithStart sets the corresponding parameter name and returns the struct.
func (r *GetAggregateMetaDataOptionalParameters) WithStart(start int64) *GetAggregateMetaDataOptionalParameters {
	r.Start = &start
	return r
}
// WithEnd sets the corresponding parameter name and returns the struct.
func (r *GetAggregateMetaDataOptionalParameters) WithEnd(end int64) *GetAggregateMetaDataOptionalParameters {
	r.End = &end
	return r
}

// GetAggregateMetaData Get aggregate meta data.
// Get aggregate meta data including total count and time series
func (a *MetricsApi) GetAggregateMetaData(ctx _context.Context, metaData AggregateMetaDataType, o ...GetAggregateMetaDataOptionalParameters) (AggregateMetaData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  AggregateMetaData
		optionalParams GetAggregateMetaDataOptionalParameters
	)

    
    if len(o) > 1 {
        return  localVarReturnValue, nil, common.ReportError("only one argument of type GetAggregateMetaDataOptionalParameters is allowed")
    }
    if len(o) == 1 {
        optionalParams = o[0]
    }
    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".MetricsApi.GetAggregateMetaData")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/metrics/metaData/aggregate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("metaData", common.ParameterToString(metaData, ""))
	if optionalParams.Start != nil {
		localVarQueryParams.Add("start", common.ParameterToString(*optionalParams.Start, ""))
	}
	if optionalParams.End != nil {
		localVarQueryParams.Add("end", common.ParameterToString(*optionalParams.End, ""))
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
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403{
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
			ErrorBody:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetEnvironmentStats Get environment stats.
// Get environment current stats
func (a *MetricsApi) GetEnvironmentStats(ctx _context.Context) (EnvironmentStats, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  EnvironmentStats
	)

    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".MetricsApi.GetEnvironmentStats")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/metrics/environment/stats"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403{
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
			ErrorBody:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// QueryClusterMetricsOptionalParameters holds optional parameters for QueryClusterMetrics.
type QueryClusterMetricsOptionalParameters struct {
	Start *int64
	End *int64
}

// NewQueryClusterMetricsOptionalParameters creates an empty struct for parameters.
func NewQueryClusterMetricsOptionalParameters() *QueryClusterMetricsOptionalParameters {
	this := QueryClusterMetricsOptionalParameters{}
	return &this
}
// WithStart sets the corresponding parameter name and returns the struct.
func (r *QueryClusterMetricsOptionalParameters) WithStart(start int64) *QueryClusterMetricsOptionalParameters {
	r.Start = &start
	return r
}
// WithEnd sets the corresponding parameter name and returns the struct.
func (r *QueryClusterMetricsOptionalParameters) WithEnd(end int64) *QueryClusterMetricsOptionalParameters {
	r.End = &end
	return r
}

// QueryClusterMetrics Query cluster metrics.
// Query cluster metrics by specified metric name and instance name, support instant and range query
func (a *MetricsApi) QueryClusterMetrics(ctx _context.Context, orgName string, clusterName string, query string, queryType MetricsQueryType, o ...QueryClusterMetricsOptionalParameters) (ClusterMetrics, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  ClusterMetrics
		optionalParams QueryClusterMetricsOptionalParameters
	)

    
    if len(o) > 1 {
        return  localVarReturnValue, nil, common.ReportError("only one argument of type QueryClusterMetricsOptionalParameters is allowed")
    }
    if len(o) == 1 {
        optionalParams = o[0]
    }
    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".MetricsApi.QueryClusterMetrics")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("query", common.ParameterToString(query, ""))
	localVarQueryParams.Add("queryType", common.ParameterToString(queryType, ""))
	if optionalParams.Start != nil {
		localVarQueryParams.Add("start", common.ParameterToString(*optionalParams.Start, ""))
	}
	if optionalParams.End != nil {
		localVarQueryParams.Add("end", common.ParameterToString(*optionalParams.End, ""))
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
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403||localVarHTTPResponse.StatusCode == 404{
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
			ErrorBody:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewMetricsApi Returns NewMetricsApi.
func NewMetricsApi(client *common.APIClient) *MetricsApi {
	return &MetricsApi{
		Client: client,
	}
}
