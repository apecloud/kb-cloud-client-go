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

// CalculateDailySLAOptionalParameters holds optional parameters for CalculateDailySLA.
type CalculateDailySLAOptionalParameters struct {
	EnvironmentName *[]string
	ClusterId       *[]string
	Engine          *[]string
	OrgName         *[]string
	TimezoneOffset  *int64
}

// NewCalculateDailySLAOptionalParameters creates an empty struct for parameters.
func NewCalculateDailySLAOptionalParameters() *CalculateDailySLAOptionalParameters {
	this := CalculateDailySLAOptionalParameters{}
	return &this
}

// WithEnvironmentName sets the corresponding parameter name and returns the struct.
func (r *CalculateDailySLAOptionalParameters) WithEnvironmentName(environmentName []string) *CalculateDailySLAOptionalParameters {
	r.EnvironmentName = &environmentName
	return r
}

// WithClusterId sets the corresponding parameter name and returns the struct.
func (r *CalculateDailySLAOptionalParameters) WithClusterId(clusterId []string) *CalculateDailySLAOptionalParameters {
	r.ClusterId = &clusterId
	return r
}

// WithEngine sets the corresponding parameter name and returns the struct.
func (r *CalculateDailySLAOptionalParameters) WithEngine(engine []string) *CalculateDailySLAOptionalParameters {
	r.Engine = &engine
	return r
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *CalculateDailySLAOptionalParameters) WithOrgName(orgName []string) *CalculateDailySLAOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithTimezoneOffset sets the corresponding parameter name and returns the struct.
func (r *CalculateDailySLAOptionalParameters) WithTimezoneOffset(timezoneOffset int64) *CalculateDailySLAOptionalParameters {
	r.TimezoneOffset = &timezoneOffset
	return r
}

// CalculateDailySLA Calculate daily SLA for a environment or a cluster since a specific date.
// Calculate daily SLA for a environment or a cluster since a specific date
func (a *SLAApi) CalculateDailySLA(ctx _context.Context, rangeVar int32, o ...CalculateDailySLAOptionalParameters) ([]DailySLA, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []DailySLA
		optionalParams      CalculateDailySLAOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type CalculateDailySLAOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "SLA",
		OperationID: "CalculateDailySLA",
		Path:        "/admin/v1/sla/daily",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SLAApi.CalculateDailySLA")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/sla/daily"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("range", common.ParameterToString(rangeVar, ""))
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
	if optionalParams.TimezoneOffset != nil {
		localVarQueryParams.Add("timezoneOffset", common.ParameterToString(*optionalParams.TimezoneOffset, ""))
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

// CalculateSLAOptionalParameters holds optional parameters for CalculateSLA.
type CalculateSLAOptionalParameters struct {
	EnvironmentName   *[]string
	ClusterId         *[]string
	Engine            *[]string
	OrgName           *[]string
	ClusterNamePrefix *string
	Page              *int32
	PageSize          *int32
	OrderBy           *string
	Desc              *bool
	Limit             *int32
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

// WithClusterNamePrefix sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithClusterNamePrefix(clusterNamePrefix string) *CalculateSLAOptionalParameters {
	r.ClusterNamePrefix = &clusterNamePrefix
	return r
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithPage(page int32) *CalculateSLAOptionalParameters {
	r.Page = &page
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithPageSize(pageSize int32) *CalculateSLAOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithOrderBy sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithOrderBy(orderBy string) *CalculateSLAOptionalParameters {
	r.OrderBy = &orderBy
	return r
}

// WithDesc sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithDesc(desc bool) *CalculateSLAOptionalParameters {
	r.Desc = &desc
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *CalculateSLAOptionalParameters) WithLimit(limit int32) *CalculateSLAOptionalParameters {
	r.Limit = &limit
	return r
}

// CalculateSLA Calculate SLA for a environment.
// Calculate SLA for a environment
func (a *SLAApi) CalculateSLA(ctx _context.Context, rangeVar int32, o ...CalculateSLAOptionalParameters) (ClustersSLA, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClustersSLA
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
	localVarQueryParams.Add("range", common.ParameterToString(rangeVar, ""))
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
	if optionalParams.ClusterNamePrefix != nil {
		localVarQueryParams.Add("clusterNamePrefix", common.ParameterToString(*optionalParams.ClusterNamePrefix, ""))
	}
	if optionalParams.Page != nil {
		localVarQueryParams.Add("page", common.ParameterToString(*optionalParams.Page, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("pageSize", common.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.OrderBy != nil {
		localVarQueryParams.Add("orderBy", common.ParameterToString(*optionalParams.OrderBy, ""))
	}
	if optionalParams.Desc != nil {
		localVarQueryParams.Add("desc", common.ParameterToString(*optionalParams.Desc, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
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

// GetSLASettingsInEnv Get SLA settings.
// Get SLA settings in a environment
func (a *SLAApi) GetSLASettingsInEnv(ctx _context.Context, environmentName string) (SLASettings, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SLASettings
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "SLA",
		OperationID: "GetSLASettingsInEnv",
		Path:        "/admin/v1/sla/settings",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SLAApi.GetSLASettingsInEnv")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/sla/settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("environmentName", common.ParameterToString(environmentName, ""))
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

// GetSLASupportEngines Get SLA supported engines.
// Get SLA supported engines in a environment
func (a *SLAApi) GetSLASupportEngines(ctx _context.Context) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []string
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "SLA",
		OperationID: "GetSLASupportEngines",
		Path:        "/admin/v1/sla/engines",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SLAApi.GetSLASupportEngines")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/sla/engines"

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

// ListUncompliantClustersOptionalParameters holds optional parameters for ListUncompliantClusters.
type ListUncompliantClustersOptionalParameters struct {
	EnvironmentName *[]string
	Engine          *[]string
	OrgName         *[]string
	Threshold       *float64
}

// NewListUncompliantClustersOptionalParameters creates an empty struct for parameters.
func NewListUncompliantClustersOptionalParameters() *ListUncompliantClustersOptionalParameters {
	this := ListUncompliantClustersOptionalParameters{}
	return &this
}

// WithEnvironmentName sets the corresponding parameter name and returns the struct.
func (r *ListUncompliantClustersOptionalParameters) WithEnvironmentName(environmentName []string) *ListUncompliantClustersOptionalParameters {
	r.EnvironmentName = &environmentName
	return r
}

// WithEngine sets the corresponding parameter name and returns the struct.
func (r *ListUncompliantClustersOptionalParameters) WithEngine(engine []string) *ListUncompliantClustersOptionalParameters {
	r.Engine = &engine
	return r
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *ListUncompliantClustersOptionalParameters) WithOrgName(orgName []string) *ListUncompliantClustersOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithThreshold sets the corresponding parameter name and returns the struct.
func (r *ListUncompliantClustersOptionalParameters) WithThreshold(threshold float64) *ListUncompliantClustersOptionalParameters {
	r.Threshold = &threshold
	return r
}

// ListUncompliantClusters List uncompliant clusters.
// List uncompliant clusters that sla lower than the threshold
func (a *SLAApi) ListUncompliantClusters(ctx _context.Context, rangeVar int32, o ...ListUncompliantClustersOptionalParameters) ([]SLA, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []SLA
		optionalParams      ListUncompliantClustersOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListUncompliantClustersOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "SLA",
		OperationID: "ListUncompliantClusters",
		Path:        "/admin/v1/sla/uncompliant",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SLAApi.ListUncompliantClusters")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/sla/uncompliant"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("range", common.ParameterToString(rangeVar, ""))
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
	if optionalParams.Threshold != nil {
		localVarQueryParams.Add("threshold", common.ParameterToString(*optionalParams.Threshold, ""))
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

// UpdateSLASettingsInEnv Update SLA settings.
// Update SLA settings in a environment
func (a *SLAApi) UpdateSLASettingsInEnv(ctx _context.Context, environmentName string, body SLASettings) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "SLA",
		OperationID: "UpdateSLASettingsInEnv",
		Path:        "/admin/v1/sla/settings",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SLAApi.UpdateSLASettingsInEnv")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/sla/settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("environmentName", common.ParameterToString(environmentName, ""))
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"BearerToken", "authorization"},
	)
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// NewSLAApi Returns NewSLAApi.
func NewSLAApi(client *common.APIClient) *SLAApi {
	return &SLAApi{
		Client: client,
	}
}
