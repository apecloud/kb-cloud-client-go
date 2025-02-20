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

// ListEventsOptionalParameters holds optional parameters for ListEvents.
type ListEventsOptionalParameters struct {
	OrgName      *string
	ResourceId   *int64
	ResourceType *string
	EventName    *string
	OperatorId   *int64
	PageNumber   *int64
	PageSize     *int64
	OrderBy      *string
}

// NewListEventsOptionalParameters creates an empty struct for parameters.
func NewListEventsOptionalParameters() *ListEventsOptionalParameters {
	this := ListEventsOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithOrgName(orgName string) *ListEventsOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithResourceId sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithResourceId(resourceId int64) *ListEventsOptionalParameters {
	r.ResourceId = &resourceId
	return r
}

// WithResourceType sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithResourceType(resourceType string) *ListEventsOptionalParameters {
	r.ResourceType = &resourceType
	return r
}

// WithEventName sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithEventName(eventName string) *ListEventsOptionalParameters {
	r.EventName = &eventName
	return r
}

// WithOperatorId sets the corresponding parameter name and returns the struct.
func (r *ListEventsOptionalParameters) WithOperatorId(operatorId int64) *ListEventsOptionalParameters {
	r.OperatorId = &operatorId
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

// NewEventApi Returns NewEventApi.
func NewEventApi(client *common.APIClient) *EventApi {
	return &EventApi{
		Client: client,
	}
}
