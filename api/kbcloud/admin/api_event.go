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

// QueryClusterEventsOptionalParameters holds optional parameters for QueryClusterEvents.
type QueryClusterEventsOptionalParameters struct {
	OrgName      *string
	ResourceId   *int32
	ResourceType *string
	EventName    *string
	OperatorId   *int32
	Status       *EventResultStatus
	PageNumber   *int32
	PageSize     *int32
	OrderBy      *string
}

// NewQueryClusterEventsOptionalParameters creates an empty struct for parameters.
func NewQueryClusterEventsOptionalParameters() *QueryClusterEventsOptionalParameters {
	this := QueryClusterEventsOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithOrgName(orgName string) *QueryClusterEventsOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithResourceId sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithResourceId(resourceId int32) *QueryClusterEventsOptionalParameters {
	r.ResourceId = &resourceId
	return r
}

// WithResourceType sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithResourceType(resourceType string) *QueryClusterEventsOptionalParameters {
	r.ResourceType = &resourceType
	return r
}

// WithEventName sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithEventName(eventName string) *QueryClusterEventsOptionalParameters {
	r.EventName = &eventName
	return r
}

// WithOperatorId sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithOperatorId(operatorId int32) *QueryClusterEventsOptionalParameters {
	r.OperatorId = &operatorId
	return r
}

// WithStatus sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithStatus(status EventResultStatus) *QueryClusterEventsOptionalParameters {
	r.Status = &status
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithPageNumber(pageNumber int32) *QueryClusterEventsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithPageSize(pageSize int32) *QueryClusterEventsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithOrderBy sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithOrderBy(orderBy string) *QueryClusterEventsOptionalParameters {
	r.OrderBy = &orderBy
	return r
}

// QueryClusterEvents Query operation events.
// Query events of clusters
func (a *EventApi) QueryClusterEvents(ctx _context.Context, start int64, end int64, o ...QueryClusterEventsOptionalParameters) (EventList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EventList
		optionalParams      QueryClusterEventsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type QueryClusterEventsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "event",
		OperationID: "queryClusterEvents",
		Path:        "/admin/v1/events",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EventApi.QueryClusterEvents")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("start", common.ParameterToString(start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(end, ""))
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.ResourceId != nil {
		localVarQueryParams.Add("resourceId", common.ParameterToString(*optionalParams.ResourceId, ""))
	}
	if optionalParams.ResourceType != nil {
		localVarQueryParams.Add("resourceType", common.ParameterToString(*optionalParams.ResourceType, ""))
	}
	if optionalParams.EventName != nil {
		localVarQueryParams.Add("eventName", common.ParameterToString(*optionalParams.EventName, ""))
	}
	if optionalParams.OperatorId != nil {
		localVarQueryParams.Add("operatorId", common.ParameterToString(*optionalParams.OperatorId, ""))
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

// QueryEventDetail Query event detail by Event ID.
// Retrieves detailed information about an event based on the provided Event ID.
func (a *EventApi) QueryEventDetail(ctx _context.Context, eventId string) (Event, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue Event
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "event",
		OperationID: "queryEventDetail",
		Path:        "/admin/v1/events/{eventID}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EventApi.QueryEventDetail")
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

// QueryEventFilter Query available filters for event listing.
func (a *EventApi) QueryEventFilter(ctx _context.Context, filterType EventQueryFilterType, start int64, end int64) (EventQueryFilterOptionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EventQueryFilterOptionList
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "event",
		OperationID: "queryEventFilter",
		Path:        "/admin/v1/eventqueryfilter",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EventApi.QueryEventFilter")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/eventqueryfilter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("filterType", common.ParameterToString(filterType, ""))
	localVarQueryParams.Add("start", common.ParameterToString(start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(end, ""))
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
