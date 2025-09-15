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

// BillingApi service type
type BillingApi common.Service

// QueryBillDetailOptionalParameters holds optional parameters for QueryBillDetail.
type QueryBillDetailOptionalParameters struct {
	BillId           *string
	ClusterId        *string
	OrgName          *string
	ProjectName      *string
	AggregationTime  *AggregationTimeType
	AggregationGroup *AggregationGroupType
}

// NewQueryBillDetailOptionalParameters creates an empty struct for parameters.
func NewQueryBillDetailOptionalParameters() *QueryBillDetailOptionalParameters {
	this := QueryBillDetailOptionalParameters{}
	return &this
}

// WithBillId sets the corresponding parameter name and returns the struct.
func (r *QueryBillDetailOptionalParameters) WithBillId(billId string) *QueryBillDetailOptionalParameters {
	r.BillId = &billId
	return r
}

// WithClusterId sets the corresponding parameter name and returns the struct.
func (r *QueryBillDetailOptionalParameters) WithClusterId(clusterId string) *QueryBillDetailOptionalParameters {
	r.ClusterId = &clusterId
	return r
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *QueryBillDetailOptionalParameters) WithOrgName(orgName string) *QueryBillDetailOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithProjectName sets the corresponding parameter name and returns the struct.
func (r *QueryBillDetailOptionalParameters) WithProjectName(projectName string) *QueryBillDetailOptionalParameters {
	r.ProjectName = &projectName
	return r
}

// WithAggregationTime sets the corresponding parameter name and returns the struct.
func (r *QueryBillDetailOptionalParameters) WithAggregationTime(aggregationTime AggregationTimeType) *QueryBillDetailOptionalParameters {
	r.AggregationTime = &aggregationTime
	return r
}

// WithAggregationGroup sets the corresponding parameter name and returns the struct.
func (r *QueryBillDetailOptionalParameters) WithAggregationGroup(aggregationGroup AggregationGroupType) *QueryBillDetailOptionalParameters {
	r.AggregationGroup = &aggregationGroup
	return r
}

// QueryBillDetail Query bill details.
// Deprecated: This API is deprecated.
func (a *BillingApi) QueryBillDetail(ctx _context.Context, start int64, end int64, o ...QueryBillDetailOptionalParameters) (BillDetailList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue BillDetailList
		optionalParams      QueryBillDetailOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryBillDetailOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "billing",
		OperationID: "queryBillDetail",
		Path:        "/api/v1/billing/detail",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".BillingApi.QueryBillDetail")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/billing/detail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("start", common.ParameterToString(start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(end, ""))
	if optionalParams.BillId != nil {
		localVarQueryParams.Add("billID", common.ParameterToString(*optionalParams.BillId, ""))
	}
	if optionalParams.ClusterId != nil {
		localVarQueryParams.Add("clusterID", common.ParameterToString(*optionalParams.ClusterId, ""))
	}
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.ProjectName != nil {
		localVarQueryParams.Add("projectName", common.ParameterToString(*optionalParams.ProjectName, ""))
	}
	if optionalParams.AggregationTime != nil {
		localVarQueryParams.Add("aggregationTime", common.ParameterToString(*optionalParams.AggregationTime, ""))
	}
	if optionalParams.AggregationGroup != nil {
		localVarQueryParams.Add("aggregationGroup", common.ParameterToString(*optionalParams.AggregationGroup, ""))
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

// QueryBillOverviewOptionalParameters holds optional parameters for QueryBillOverview.
type QueryBillOverviewOptionalParameters struct {
	AggregationTime  *AggregationTimeType
	AggregationGroup *AggregationGroupType
}

// NewQueryBillOverviewOptionalParameters creates an empty struct for parameters.
func NewQueryBillOverviewOptionalParameters() *QueryBillOverviewOptionalParameters {
	this := QueryBillOverviewOptionalParameters{}
	return &this
}

// WithAggregationTime sets the corresponding parameter name and returns the struct.
func (r *QueryBillOverviewOptionalParameters) WithAggregationTime(aggregationTime AggregationTimeType) *QueryBillOverviewOptionalParameters {
	r.AggregationTime = &aggregationTime
	return r
}

// WithAggregationGroup sets the corresponding parameter name and returns the struct.
func (r *QueryBillOverviewOptionalParameters) WithAggregationGroup(aggregationGroup AggregationGroupType) *QueryBillOverviewOptionalParameters {
	r.AggregationGroup = &aggregationGroup
	return r
}

// QueryBillOverview Query bills of all environments and organizations.
// Deprecated: This API is deprecated.
func (a *BillingApi) QueryBillOverview(ctx _context.Context, start int64, end int64, o ...QueryBillOverviewOptionalParameters) (BasicBillList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue BasicBillList
		optionalParams      QueryBillOverviewOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryBillOverviewOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "billing",
		OperationID: "queryBillOverview",
		Path:        "/api/v1/billing/overview",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".BillingApi.QueryBillOverview")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/billing/overview"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("start", common.ParameterToString(start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(end, ""))
	if optionalParams.AggregationTime != nil {
		localVarQueryParams.Add("aggregationTime", common.ParameterToString(*optionalParams.AggregationTime, ""))
	}
	if optionalParams.AggregationGroup != nil {
		localVarQueryParams.Add("aggregationGroup", common.ParameterToString(*optionalParams.AggregationGroup, ""))
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

// NewBillingApi Returns NewBillingApi.
func NewBillingApi(client *common.APIClient) *BillingApi {
	return &BillingApi{
		Client: client,
	}
}
