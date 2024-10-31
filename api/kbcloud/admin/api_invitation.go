// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InvitationApi service type
type InvitationApi common.Service

// ListInvitationOptionalParameters holds optional parameters for ListInvitation.
type ListInvitationOptionalParameters struct {
	OrgName *string
	UserId  *string
}

// NewListInvitationOptionalParameters creates an empty struct for parameters.
func NewListInvitationOptionalParameters() *ListInvitationOptionalParameters {
	this := ListInvitationOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *ListInvitationOptionalParameters) WithOrgName(orgName string) *ListInvitationOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithUserId sets the corresponding parameter name and returns the struct.
func (r *ListInvitationOptionalParameters) WithUserId(userId string) *ListInvitationOptionalParameters {
	r.UserId = &userId
	return r
}

// ListInvitation List invitations.
// list the Invitation of specified Org or User
func (a *InvitationApi) ListInvitation(ctx _context.Context, o ...ListInvitationOptionalParameters) (InvitationList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue InvitationList
		optionalParams      ListInvitationOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListInvitationOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".InvitationApi.ListInvitation")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/invitations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.UserId != nil {
		localVarQueryParams.Add("userId", common.ParameterToString(*optionalParams.UserId, ""))
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 {
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

// NewInvitationApi Returns NewInvitationApi.
func NewInvitationApi(client *common.APIClient) *InvitationApi {
	return &InvitationApi{
		Client: client,
	}
}
