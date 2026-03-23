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

// BillingApi service type
type BillingApi common.Service

// ListOrgBillsOptionalParameters holds optional parameters for ListOrgBills.
type ListOrgBillsOptionalParameters struct {
	BillId           *string
	ClusterId        *string
	ProjectName      *string
	AggregationTime  *AggregationTimeType
	AggregationGroup *OrgAggregationGroupType
}

// NewListOrgBillsOptionalParameters creates an empty struct for parameters.
func NewListOrgBillsOptionalParameters() *ListOrgBillsOptionalParameters {
	this := ListOrgBillsOptionalParameters{}
	return &this
}

// WithBillId sets the corresponding parameter name and returns the struct.
func (r *ListOrgBillsOptionalParameters) WithBillId(billId string) *ListOrgBillsOptionalParameters {
	r.BillId = &billId
	return r
}

// WithClusterId sets the corresponding parameter name and returns the struct.
func (r *ListOrgBillsOptionalParameters) WithClusterId(clusterId string) *ListOrgBillsOptionalParameters {
	r.ClusterId = &clusterId
	return r
}

// WithProjectName sets the corresponding parameter name and returns the struct.
func (r *ListOrgBillsOptionalParameters) WithProjectName(projectName string) *ListOrgBillsOptionalParameters {
	r.ProjectName = &projectName
	return r
}

// WithAggregationTime sets the corresponding parameter name and returns the struct.
func (r *ListOrgBillsOptionalParameters) WithAggregationTime(aggregationTime AggregationTimeType) *ListOrgBillsOptionalParameters {
	r.AggregationTime = &aggregationTime
	return r
}

// WithAggregationGroup sets the corresponding parameter name and returns the struct.
func (r *ListOrgBillsOptionalParameters) WithAggregationGroup(aggregationGroup OrgAggregationGroupType) *ListOrgBillsOptionalParameters {
	r.AggregationGroup = &aggregationGroup
	return r
}

// ListOrgBills List bills in the organization.
func (a *BillingApi) ListOrgBills(ctx _context.Context, orgName string, start int64, end int64, o ...ListOrgBillsOptionalParameters) (BillList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue BillList
		optionalParams      ListOrgBillsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListOrgBillsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "billing",
		OperationID: "listOrgBills",
		Path:        "/api/v1/organizations/{orgName}/bills",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".BillingApi.ListOrgBills")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/bills"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

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
