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

// AlertReceiverApi service type
type AlertReceiverApi common.Service

// CreateAlertReceiverOptionalParameters holds optional parameters for CreateAlertReceiver.
type CreateAlertReceiverOptionalParameters struct {
	OrgName *string
}

// NewCreateAlertReceiverOptionalParameters creates an empty struct for parameters.
func NewCreateAlertReceiverOptionalParameters() *CreateAlertReceiverOptionalParameters {
	this := CreateAlertReceiverOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *CreateAlertReceiverOptionalParameters) WithOrgName(orgName string) *CreateAlertReceiverOptionalParameters {
	r.OrgName = &orgName
	return r
}

// CreateAlertReceiver Create alert receiver.
func (a *AlertReceiverApi) CreateAlertReceiver(ctx _context.Context, level AlertLevel, category AlertReceiverCategory, body AlertReceiver, o ...CreateAlertReceiverOptionalParameters) (AlertReceiver, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue AlertReceiver
		optionalParams      CreateAlertReceiverOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type CreateAlertReceiverOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "alertReceiver",
		OperationID: "createAlertReceiver",
		Path:        "/admin/v1/alerts/receivers",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AlertReceiverApi.CreateAlertReceiver")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/alerts/receivers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("level", common.ParameterToString(level, ""))
	localVarQueryParams.Add("category", common.ParameterToString(category, ""))
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
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

// DeleteAlertReceiverOptionalParameters holds optional parameters for DeleteAlertReceiver.
type DeleteAlertReceiverOptionalParameters struct {
	OrgName *string
}

// NewDeleteAlertReceiverOptionalParameters creates an empty struct for parameters.
func NewDeleteAlertReceiverOptionalParameters() *DeleteAlertReceiverOptionalParameters {
	this := DeleteAlertReceiverOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *DeleteAlertReceiverOptionalParameters) WithOrgName(orgName string) *DeleteAlertReceiverOptionalParameters {
	r.OrgName = &orgName
	return r
}

// DeleteAlertReceiver Delete alert receiver.
func (a *AlertReceiverApi) DeleteAlertReceiver(ctx _context.Context, receiverId string, level AlertLevel, o ...DeleteAlertReceiverOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
		optionalParams     DeleteAlertReceiverOptionalParameters
	)

	if len(o) > 1 {
		return nil, common.ReportError("only one argument of type DeleteAlertReceiverOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "alertReceiver",
		OperationID: "deleteAlertReceiver",
		Path:        "/admin/v1/alerts/receivers/{receiverId}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AlertReceiverApi.DeleteAlertReceiver")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/alerts/receivers/{receiverId}"
	localVarPath = strings.Replace(localVarPath, "{"+"receiverId"+"}", _neturl.PathEscape(common.ParameterToString(receiverId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("level", common.ParameterToString(level, ""))
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// ListAlertReceiversOptionalParameters holds optional parameters for ListAlertReceivers.
type ListAlertReceiversOptionalParameters struct {
	OrgName  *string
	Category *AlertReceiverCategory
}

// NewListAlertReceiversOptionalParameters creates an empty struct for parameters.
func NewListAlertReceiversOptionalParameters() *ListAlertReceiversOptionalParameters {
	this := ListAlertReceiversOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *ListAlertReceiversOptionalParameters) WithOrgName(orgName string) *ListAlertReceiversOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithCategory sets the corresponding parameter name and returns the struct.
func (r *ListAlertReceiversOptionalParameters) WithCategory(category AlertReceiverCategory) *ListAlertReceiversOptionalParameters {
	r.Category = &category
	return r
}

// ListAlertReceivers List alert receivers.
func (a *AlertReceiverApi) ListAlertReceivers(ctx _context.Context, level AlertLevel, o ...ListAlertReceiversOptionalParameters) (AlertReceiverList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue AlertReceiverList
		optionalParams      ListAlertReceiversOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListAlertReceiversOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "alertReceiver",
		OperationID: "listAlertReceivers",
		Path:        "/admin/v1/alerts/receivers",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AlertReceiverApi.ListAlertReceivers")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/alerts/receivers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("level", common.ParameterToString(level, ""))
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
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

// UpdateAlertReceiverOptionalParameters holds optional parameters for UpdateAlertReceiver.
type UpdateAlertReceiverOptionalParameters struct {
	OrgName *string
}

// NewUpdateAlertReceiverOptionalParameters creates an empty struct for parameters.
func NewUpdateAlertReceiverOptionalParameters() *UpdateAlertReceiverOptionalParameters {
	this := UpdateAlertReceiverOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *UpdateAlertReceiverOptionalParameters) WithOrgName(orgName string) *UpdateAlertReceiverOptionalParameters {
	r.OrgName = &orgName
	return r
}

// UpdateAlertReceiver Update alert receiver.
func (a *AlertReceiverApi) UpdateAlertReceiver(ctx _context.Context, receiverId string, level AlertLevel, body AlertReceiver, o ...UpdateAlertReceiverOptionalParameters) (AlertReceiver, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue AlertReceiver
		optionalParams      UpdateAlertReceiverOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type UpdateAlertReceiverOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "alertReceiver",
		OperationID: "updateAlertReceiver",
		Path:        "/admin/v1/alerts/receivers/{receiverId}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AlertReceiverApi.UpdateAlertReceiver")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/alerts/receivers/{receiverId}"
	localVarPath = strings.Replace(localVarPath, "{"+"receiverId"+"}", _neturl.PathEscape(common.ParameterToString(receiverId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("level", common.ParameterToString(level, ""))
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
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

// NewAlertReceiverApi Returns NewAlertReceiverApi.
func NewAlertReceiverApi(client *common.APIClient) *AlertReceiverApi {
	return &AlertReceiverApi{
		Client: client,
	}
}
