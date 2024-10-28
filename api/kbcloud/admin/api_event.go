// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
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
	ResourceId   *int32
	ResourceType *string
	EventName    *string
	OperatorId   *int32
	Start        *int64
	End          *int64
}

// NewQueryClusterEventsOptionalParameters creates an empty struct for parameters.
func NewQueryClusterEventsOptionalParameters() *QueryClusterEventsOptionalParameters {
	this := QueryClusterEventsOptionalParameters{}
	return &this
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

// WithStart sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithStart(start int64) *QueryClusterEventsOptionalParameters {
	r.Start = &start
	return r
}

// WithEnd sets the corresponding parameter name and returns the struct.
func (r *QueryClusterEventsOptionalParameters) WithEnd(end int64) *QueryClusterEventsOptionalParameters {
	r.End = &end
	return r
}

// QueryClusterEvents Query operation events.
// Query events of clusters
// NODESCRIPTION QueryClusterEvents
// Deprecated: This API is deprecated.
func (a *EventApi) QueryClusterEvents(ctx _context.Context, orgName string, o ...QueryClusterEventsOptionalParameters) (EventList, *_nethttp.Response, error) {
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

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EventApi.QueryClusterEvents")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
