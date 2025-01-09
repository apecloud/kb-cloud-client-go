// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SessionApi service type
type SessionApi common.Service

// KillSessionOptionalParameters holds optional parameters for KillSession.
type KillSessionOptionalParameters struct {
	Keep *bool
}

// NewKillSessionOptionalParameters creates an empty struct for parameters.
func NewKillSessionOptionalParameters() *KillSessionOptionalParameters {
	this := KillSessionOptionalParameters{}
	return &this
}

// WithKeep sets the corresponding parameter name and returns the struct.
func (r *KillSessionOptionalParameters) WithKeep(keep bool) *KillSessionOptionalParameters {
	r.Keep = &keep
	return r
}

// KillSession Kill cluster session.
// kill a session in cluster
func (a *SessionApi) KillSession(ctx _context.Context, engineName string, orgName string, clusterName string, session string, o ...KillSessionOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
		optionalParams     KillSessionOptionalParameters
	)

	if len(o) > 1 {
		return nil, common.ReportError("only one argument of type KillSessionOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SessionApi.KillSession")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/{engineName}/organizations/{orgName}/clusters/{clusterName}/sessions/{session}"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"session"+"}", _neturl.PathEscape(common.ParameterToString(session, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Keep != nil {
		localVarQueryParams.Add("keep", common.ParameterToString(*optionalParams.Keep, ""))
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

// ListSessionsOptionalParameters holds optional parameters for ListSessions.
type ListSessionsOptionalParameters struct {
	All *bool
}

// NewListSessionsOptionalParameters creates an empty struct for parameters.
func NewListSessionsOptionalParameters() *ListSessionsOptionalParameters {
	this := ListSessionsOptionalParameters{}
	return &this
}

// WithAll sets the corresponding parameter name and returns the struct.
func (r *ListSessionsOptionalParameters) WithAll(all bool) *ListSessionsOptionalParameters {
	r.All = &all
	return r
}

// ListSessions List cluster sessions.
// list sessions in cluster
func (a *SessionApi) ListSessions(ctx _context.Context, engineName string, orgName string, clusterName string, o ...ListSessionsOptionalParameters) ([]DmsSession, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []DmsSession
		optionalParams      ListSessionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListSessionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".SessionApi.ListSessions")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/data/v1/{engineName}/organizations/{orgName}/clusters/{clusterName}/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.All != nil {
		localVarQueryParams.Add("all", common.ParameterToString(*optionalParams.All, ""))
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 500 {
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

// NewSessionApi Returns NewSessionApi.
func NewSessionApi(client *common.APIClient) *SessionApi {
	return &SessionApi{
		Client: client,
	}
}
