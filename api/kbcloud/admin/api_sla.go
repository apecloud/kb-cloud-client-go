// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"context"
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SLAApi service type
type SLAApi common.Service

// CalculateSLAOptionalParameters holds optional parameters for CalculateSLA.
type CalculateSLAOptionalParameters struct {
	EnvironmentName *[]string
	ClusterId       *[]string
	Engine          *[]string
	OrgName         *[]string
	StartTime       *int64
}

// NewCalculateSLAOptionalParameters creates an empty struct for parameters.
func NewCalculateSLAOptionalParameters() *CalculateSLAOptionalParameters {
	this := CalculateSLAOptionalParameters{}
	return &this
}

// WithEnvironmentName sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithEnvironmentName(environmentName []string) *CalculateSLAOptionalParameters {
	r.EnvironmentName = &environmentName
	return r
}

// WithClusterId sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithClusterId(clusterId []string) *CalculateSLAOptionalParameters {
	r.ClusterId = &clusterId
	return r
}

// WithEngine sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithEngine(engine []string) *CalculateSLAOptionalParameters {
	r.Engine = &engine
	return r
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithOrgName(orgName []string) *CalculateSLAOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithStartTime sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithStartTime(startTime int64) *CalculateSLAOptionalParameters {
	r.StartTime = &startTime
	return r
}

// CalculateSLA Calculate SLA for a environment.
// Calculate SLA for a environment
func (a *SLAApi) CalculateSLA(ctx _context.Context, o ...CalculateSLAOptionalParameters) ([]SLA, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []SLA
		optionalParams      CalculateSLAOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type CalculateSLAOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "SLA",
		OperationID: "CalculateSLA",
		Path:        "/admin/v1/sla",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SLAApi.CalculateSLA")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/sla"

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
	if optionalParams.OrgName != nil {
		t := *optionalParams.OrgName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("orgName", common.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("orgName", common.ParameterToString(t, "multi"))
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

// ListEnvironmentOutageRecordOptionalParameters holds optional parameters for ListEnvironmentOutageRecord.
type ListEnvironmentOutageRecordOptionalParameters struct {
	EnvironmentName *string
	OrgName         *string
	ClusterId       *string
	ActiveOnly      *bool
}

// NewListEnvironmentOutageRecordOptionalParameters creates an empty struct for parameters.
func NewListEnvironmentOutageRecordOptionalParameters() *ListEnvironmentOutageRecordOptionalParameters {
	this := ListEnvironmentOutageRecordOptionalParameters{}
	return &this
}

// WithEnvironmentName sets the corresponding parameter name and returns the struct.
func (r *ListEnvironmentOutageRecordOptionalParameters) WithEnvironmentName(environmentName string) *ListEnvironmentOutageRecordOptionalParameters {
	r.EnvironmentName = &environmentName
	return r
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *ListEnvironmentOutageRecordOptionalParameters) WithOrgName(orgName string) *ListEnvironmentOutageRecordOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithClusterId sets the corresponding parameter name and returns the struct.
func (r *ListEnvironmentOutageRecordOptionalParameters) WithClusterId(clusterId string) *ListEnvironmentOutageRecordOptionalParameters {
	r.ClusterId = &clusterId
	return r
}

// WithActiveOnly sets the corresponding parameter name and returns the struct.
func (r *ListEnvironmentOutageRecordOptionalParameters) WithActiveOnly(activeOnly bool) *ListEnvironmentOutageRecordOptionalParameters {
	r.ActiveOnly = &activeOnly
	return r
}

// ListEnvironmentOutageRecord List outage record for environments.
// List outage record for environments
func (a *SLAApi) ListEnvironmentOutageRecord(ctx _context.Context, o ...ListEnvironmentOutageRecordOptionalParameters) ([]OutageRecord, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []OutageRecord
		optionalParams      ListEnvironmentOutageRecordOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEnvironmentOutageRecordOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "SLA",
		OperationID: "ListEnvironmentOutageRecord",
		Path:        "/admin/v1/sla/outages",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SLAApi.ListEnvironmentOutageRecord")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/sla/outages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.EnvironmentName != nil {
		localVarQueryParams.Add("environmentName", common.ParameterToString(*optionalParams.EnvironmentName, ""))
	}
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.ClusterId != nil {
		localVarQueryParams.Add("clusterID", common.ParameterToString(*optionalParams.ClusterId, ""))
	}
	if optionalParams.ActiveOnly != nil {
		localVarQueryParams.Add("activeOnly", common.ParameterToString(*optionalParams.ActiveOnly, ""))
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
