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

// RedisApi service type
type RedisApi common.Service

// CreateRedisAccountOptionalParameters holds optional parameters for CreateRedisAccount.
type CreateRedisAccountOptionalParameters struct {
	Component *string
}

// NewCreateRedisAccountOptionalParameters creates an empty struct for parameters.
func NewCreateRedisAccountOptionalParameters() *CreateRedisAccountOptionalParameters {
	this := CreateRedisAccountOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *CreateRedisAccountOptionalParameters) WithComponent(component string) *CreateRedisAccountOptionalParameters {
	r.Component = &component
	return r
}

// CreateRedisAccount create redis account.
func (a *RedisApi) CreateRedisAccount(ctx _context.Context, orgName string, clusterName string, body ACLUser, o ...CreateRedisAccountOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
		optionalParams     CreateRedisAccountOptionalParameters
	)

	if len(o) > 1 {
		return nil, common.ReportError("only one argument of type CreateRedisAccountOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RedisApi.CreateRedisAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/data/redis/organizations/{orgName}/clusters/{clusterName}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "*/*"

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

// DeleteRedisAccountOptionalParameters holds optional parameters for DeleteRedisAccount.
type DeleteRedisAccountOptionalParameters struct {
	Component *string
}

// NewDeleteRedisAccountOptionalParameters creates an empty struct for parameters.
func NewDeleteRedisAccountOptionalParameters() *DeleteRedisAccountOptionalParameters {
	this := DeleteRedisAccountOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *DeleteRedisAccountOptionalParameters) WithComponent(component string) *DeleteRedisAccountOptionalParameters {
	r.Component = &component
	return r
}

// DeleteRedisAccount delete redis account.
func (a *RedisApi) DeleteRedisAccount(ctx _context.Context, orgName string, clusterName string, accountName string, o ...DeleteRedisAccountOptionalParameters) (ClusterTask, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		localVarReturnValue ClusterTask
		optionalParams      DeleteRedisAccountOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type DeleteRedisAccountOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RedisApi.DeleteRedisAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/data/redis/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
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

// ListRedisAccountsOptionalParameters holds optional parameters for ListRedisAccounts.
type ListRedisAccountsOptionalParameters struct {
	Component *string
}

// NewListRedisAccountsOptionalParameters creates an empty struct for parameters.
func NewListRedisAccountsOptionalParameters() *ListRedisAccountsOptionalParameters {
	this := ListRedisAccountsOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *ListRedisAccountsOptionalParameters) WithComponent(component string) *ListRedisAccountsOptionalParameters {
	r.Component = &component
	return r
}

// ListRedisAccounts list redis accounts.
func (a *RedisApi) ListRedisAccounts(ctx _context.Context, orgName string, clusterName string, o ...ListRedisAccountsOptionalParameters) (ACLUserResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ACLUserResponse
		optionalParams      ListRedisAccountsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListRedisAccountsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RedisApi.ListRedisAccounts")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/data/redis/organizations/{orgName}/clusters/{clusterName}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
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

// UpdateRedisAccountOptionalParameters holds optional parameters for UpdateRedisAccount.
type UpdateRedisAccountOptionalParameters struct {
	Component *string
}

// NewUpdateRedisAccountOptionalParameters creates an empty struct for parameters.
func NewUpdateRedisAccountOptionalParameters() *UpdateRedisAccountOptionalParameters {
	this := UpdateRedisAccountOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *UpdateRedisAccountOptionalParameters) WithComponent(component string) *UpdateRedisAccountOptionalParameters {
	r.Component = &component
	return r
}

// UpdateRedisAccount update redis account.
func (a *RedisApi) UpdateRedisAccount(ctx _context.Context, orgName string, clusterName string, accountName string, body ClusterTask, o ...UpdateRedisAccountOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
		optionalParams     UpdateRedisAccountOptionalParameters
	)

	if len(o) > 1 {
		return nil, common.ReportError("only one argument of type UpdateRedisAccountOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RedisApi.UpdateRedisAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/data/redis/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "*/*"

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

// NewRedisApi Returns NewRedisApi.
func NewRedisApi(client *common.APIClient) *RedisApi {
	return &RedisApi{
		Client: client,
	}
}
