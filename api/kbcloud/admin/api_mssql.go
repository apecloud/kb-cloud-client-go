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

// MssqlApi service type
type MssqlApi common.Service

// ManageMssqlTDEDatabaseOptionalParameters holds optional parameters for ManageMssqlTDEDatabase.
type ManageMssqlTDEDatabaseOptionalParameters struct {
	Body *DbTDERequest
}

// NewManageMssqlTDEDatabaseOptionalParameters creates an empty struct for parameters.
func NewManageMssqlTDEDatabaseOptionalParameters() *ManageMssqlTDEDatabaseOptionalParameters {
	this := ManageMssqlTDEDatabaseOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *ManageMssqlTDEDatabaseOptionalParameters) WithBody(body DbTDERequest) *ManageMssqlTDEDatabaseOptionalParameters {
	r.Body = &body
	return r
}

// ManageMssqlTDEDatabase batch modify database tde status.
func (a *MssqlApi) ManageMssqlTDEDatabase(ctx _context.Context, orgName string, clusterName string, o ...ManageMssqlTDEDatabaseOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
		optionalParams     ManageMssqlTDEDatabaseOptionalParameters
	)

	if len(o) > 1 {
		return nil, common.ReportError("only one argument of type ManageMssqlTDEDatabaseOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "mssql",
		OperationID: "manageMssqlTDEDatabase",
		Path:        "/admin/v1/data/mssql/organizations/{orgName}/clusters/{clusterName}/tde",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".MssqlApi.ManageMssqlTDEDatabase")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/data/mssql/organizations/{orgName}/clusters/{clusterName}/tde"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	if optionalParams.Body != nil {
		localVarPostBody = &optionalParams.Body
	}
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
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

// NewMssqlApi Returns NewMssqlApi.
func NewMssqlApi(client *common.APIClient) *MssqlApi {
	return &MssqlApi{
		Client: client,
	}
}
