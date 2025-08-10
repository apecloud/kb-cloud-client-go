// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"context"
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertMetricsApi service type
type AlertMetricsApi common.Service

// ListAlertMetricsOptionalParameters holds optional parameters for ListAlertMetrics.
type ListAlertMetricsOptionalParameters struct {
	Category *string
}

// NewListAlertMetricsOptionalParameters creates an empty struct for parameters.
func NewListAlertMetricsOptionalParameters() *ListAlertMetricsOptionalParameters {
	this := ListAlertMetricsOptionalParameters{}
	return &this
}

// WithCategory sets the corresponding parameter name and returns the struct.
func (r *ListAlertMetricsOptionalParameters) WithCategory(category string) *ListAlertMetricsOptionalParameters {
	r.Category = &category
	return r
}

// ListAlertMetrics List alert metric types.
func (a *AlertMetricsApi) ListAlertMetrics(ctx _context.Context, o ...ListAlertMetricsOptionalParameters) (AlertMetricList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AlertMetricList
		optionalParams      ListAlertMetricsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListAlertMetricsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "alertMetrics",
		OperationID: "listAlertMetrics",
		Path:        "/admin/v1/alerts/metrics",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AlertMetricsApi.ListAlertMetrics")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/alerts/metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Category != nil {
		localVarQueryParams.Add("category", common.ParameterToString(*optionalParams.Category, ""))
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

// NewAlertMetricsApi Returns NewAlertMetricsApi.
func NewAlertMetricsApi(client *common.APIClient) *AlertMetricsApi {
	return &AlertMetricsApi{
		Client: client,
	}
}
