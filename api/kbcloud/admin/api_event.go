// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"context"
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventApi service type
type EventApi common.Service

// GetEvent Query event detail by Event ID.
// Retrieves detailed information about an event based on the provided Event ID.
func (a *EventApi) GetEvent(ctx _context.Context, eventId string) (Event, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue Event
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "event",
		OperationID: "getEvent",
		Path:        "/admin/v1/events/{eventID}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EventApi.GetEvent")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/events/{eventID}"
	localVarPath = strings.Replace(localVarPath, "{"+"eventID"+"}", _neturl.PathEscape(common.ParameterToString(eventId, "")), -1)

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

// GetEventFilterOptionalParameters holds optional parameters for GetEventFilter.
type GetEventFilterOptionalParameters struct {
	ResourceType *string
}

// NewGetEventFilterOptionalParameters creates an empty struct for parameters.
func NewGetEventFilterOptionalParameters() *GetEventFilterOptionalParameters {
	this := GetEventFilterOptionalParameters{}
	return &this
}

// WithResourceType sets the corresponding parameter name and returns the struct.
func (r *GetEventFilterOptionalParameters) WithResourceType(resourceType string) *GetEventFilterOptionalParameters {
	r.ResourceType = &resourceType
	return r
}

// GetEventFilter Query available filters for event listing.
// Query available filters for event listing
func (a *EventApi) GetEventFilter(ctx _context.Context, filterType EventFilterType, source string, start int64, end int64, o ...GetEventFilterOptionalParameters) (EventFilterOptionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EventFilterOptionList
		optionalParams      GetEventFilterOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetEventFilterOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "event",
		OperationID: "getEventFilter",
		Path:        "/admin/v1/eventfilter/{filterType}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EventApi.GetEventFilter")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/eventfilter/{filterType}"
	localVarPath = strings.Replace(localVarPath, "{"+"filterType"+"}", _neturl.PathEscape(common.ParameterToString(filterType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("source", common.ParameterToString(source, ""))
	localVarQueryParams.Add("start", common.ParameterToString(start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(end, ""))
	if optionalParams.ResourceType != nil {
		localVarQueryParams.Add("resourceType", common.ParameterToString(*optionalParams.ResourceType, ""))
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

// ListEventsOptionalParameters holds optional parameters for ListEvents.
type ListEventsOptionalParameters struct {
	Status      *EventResultStatus
	CustomQuery *string
	PageNumber  *int64
	PageSize    *int64
	OrderBy     *string
	SortDesc    *bool
}

// NewListEventsOptionalParameters creates an empty struct for parameters.
func NewListEventsOptionalParameters() *ListEventsOptionalParameters {
	this := ListEventsOptionalParameters{}
	return &this
}

// WithStatus sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithStatus(status EventResultStatus) *ListEventsOptionalParameters {
	r.Status = &status
	return r
}

// WithCustomQuery sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithCustomQuery(customQuery string) *ListEventsOptionalParameters {
	r.CustomQuery = &customQuery
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithPageNumber(pageNumber int64) *ListEventsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithPageSize(pageSize int64) *ListEventsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithOrderBy sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithOrderBy(orderBy string) *ListEventsOptionalParameters {
	r.OrderBy = &orderBy
	return r
}

// WithSortDesc sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithSortDesc(sortDesc bool) *ListEventsOptionalParameters {
	r.SortDesc = &sortDesc
	return r
}

// ListEvents List events.
// List events
func (a *EventApi) ListEvents(ctx _context.Context, start int64, end int64, o ...ListEventsOptionalParameters) (EventList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EventList
		optionalParams      ListEventsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEventsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "event",
		OperationID: "listEvents",
		Path:        "/admin/v1/events",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EventApi.ListEvents")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("start", common.ParameterToString(start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(end, ""))
	if optionalParams.Status != nil {
		localVarQueryParams.Add("status", common.ParameterToString(*optionalParams.Status, ""))
	}
	if optionalParams.CustomQuery != nil {
		localVarQueryParams.Add("customQuery", common.ParameterToString(*optionalParams.CustomQuery, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("pageNumber", common.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("pageSize", common.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.OrderBy != nil {
		localVarQueryParams.Add("orderBy", common.ParameterToString(*optionalParams.OrderBy, ""))
	}
	if optionalParams.SortDesc != nil {
		localVarQueryParams.Add("sortDesc", common.ParameterToString(*optionalParams.SortDesc, ""))
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

// ListOrgEventsOptionalParameters holds optional parameters for ListOrgEvents.
type ListOrgEventsOptionalParameters struct {
	CustomQuery *string
	Status      *EventResultStatus
	PageNumber  *int32
	PageSize    *int32
	OrderBy     *string
	SortDesc    *bool
}

// NewListOrgEventsOptionalParameters creates an empty struct for parameters.
func NewListOrgEventsOptionalParameters() *ListOrgEventsOptionalParameters {
	this := ListOrgEventsOptionalParameters{}
	return &this
}

// WithCustomQuery sets the corresponding parameter name and returns the struct.
func (r *ListOrgEventsOptionalParameters) WithCustomQuery(customQuery string) *ListOrgEventsOptionalParameters {
	r.CustomQuery = &customQuery
	return r
}

// WithStatus sets the corresponding parameter name and returns the struct.
func (r *ListOrgEventsOptionalParameters) WithStatus(status EventResultStatus) *ListOrgEventsOptionalParameters {
	r.Status = &status
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListOrgEventsOptionalParameters) WithPageNumber(pageNumber int32) *ListOrgEventsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListOrgEventsOptionalParameters) WithPageSize(pageSize int32) *ListOrgEventsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithOrderBy sets the corresponding parameter name and returns the struct.
func (r *ListOrgEventsOptionalParameters) WithOrderBy(orderBy string) *ListOrgEventsOptionalParameters {
	r.OrderBy = &orderBy
	return r
}

// WithSortDesc sets the corresponding parameter name and returns the struct.
func (r *ListOrgEventsOptionalParameters) WithSortDesc(sortDesc bool) *ListOrgEventsOptionalParameters {
	r.SortDesc = &sortDesc
	return r
}

// ListOrgEvents List events.
// List events
func (a *EventApi) ListOrgEvents(ctx _context.Context, orgName string, start int64, end int64, o ...ListOrgEventsOptionalParameters) (EventList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EventList
		optionalParams      ListOrgEventsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListOrgEventsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "event",
		OperationID: "listOrgEvents",
		Path:        "/admin/v1/organizations/{orgName}/events",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EventApi.ListOrgEvents")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("start", common.ParameterToString(start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(end, ""))
	if optionalParams.CustomQuery != nil {
		localVarQueryParams.Add("customQuery", common.ParameterToString(*optionalParams.CustomQuery, ""))
	}
	if optionalParams.Status != nil {
		localVarQueryParams.Add("status", common.ParameterToString(*optionalParams.Status, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("pageNumber", common.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("pageSize", common.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.OrderBy != nil {
		localVarQueryParams.Add("orderBy", common.ParameterToString(*optionalParams.OrderBy, ""))
	}
	if optionalParams.SortDesc != nil {
		localVarQueryParams.Add("sortDesc", common.ParameterToString(*optionalParams.SortDesc, ""))
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

// NewEventApi Returns NewEventApi.
func NewEventApi(client *common.APIClient) *EventApi {
	return &EventApi{
		Client: client,
	}
}
