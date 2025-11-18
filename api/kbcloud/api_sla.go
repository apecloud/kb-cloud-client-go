// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"context"
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
	"strings"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SLAApi service type
type SLAApi common.Service

// CalculateSLAInOrgOptionalParameters holds optional parameters for CalculateSLAInOrg.
type CalculateSLAInOrgOptionalParameters struct {
	EnvironmentName *[]string
	ClusterId       *[]string
	Engine          *[]string
	StartTime       *int64
}

// NewCalculateSLAInOrgOptionalParameters creates an empty struct for parameters.
func NewCalculateSLAInOrgOptionalParameters() *CalculateSLAInOrgOptionalParameters {
	this := CalculateSLAInOrgOptionalParameters{}
	return &this
}

// WithEnvironmentName sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAInOrgOptionalParameters) WithEnvironmentName(environmentName []string) *CalculateSLAInOrgOptionalParameters {
	r.EnvironmentName = &environmentName
	return r
}

// WithClusterId sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAInOrgOptionalParameters) WithClusterId(clusterId []string) *CalculateSLAInOrgOptionalParameters {
	r.ClusterId = &clusterId
	return r
}

// WithEngine sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAInOrgOptionalParameters) WithEngine(engine []string) *CalculateSLAInOrgOptionalParameters {
	r.Engine = &engine
	return r
}

// WithStartTime sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAInOrgOptionalParameters) WithStartTime(startTime int64) *CalculateSLAInOrgOptionalParameters {
	r.StartTime = &startTime
	return r
}

// CalculateSLAInOrg Calculate SLA for a organization.
// Calculate SLA for a organization
func (a *SLAApi) CalculateSLAInOrg(ctx _context.Context, orgName string, o ...CalculateSLAInOrgOptionalParameters) ([]SLA, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []SLA
		optionalParams      CalculateSLAInOrgOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type CalculateSLAInOrgOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "SLA",
		OperationID: "CalculateSLAInOrg",
		Path:        "/api/v1/organizations/{orgName}/sla",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SLAApi.CalculateSLAInOrg")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/sla"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.EnvironmentName != nil {
		t := *optionalParams.EnvironmentName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("environmentName", common.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("environmentName", common.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.ClusterId != nil {
		t := *optionalParams.ClusterId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("clusterID", common.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("clusterID", common.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.Engine != nil {
		t := *optionalParams.Engine
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("engine", common.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("engine", common.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.StartTime != nil {
		localVarQueryParams.Add("startTime", common.ParameterToString(*optionalParams.StartTime, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
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

// NewSLAApi Returns NewSLAApi.
func NewSLAApi(client *common.APIClient) *SLAApi {
	return &SLAApi{
		Client: client,
	}
}
