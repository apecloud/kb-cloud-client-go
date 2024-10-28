// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineLicenseApi service type
type EngineLicenseApi common.Service

// ListEngineLicensesOptionalParameters holds optional parameters for ListEngineLicenses.
type ListEngineLicensesOptionalParameters struct {
	EngineName *string
}

// NewListEngineLicensesOptionalParameters creates an empty struct for parameters.
func NewListEngineLicensesOptionalParameters() *ListEngineLicensesOptionalParameters {
	this := ListEngineLicensesOptionalParameters{}
	return &this
}

// WithEngineName sets the corresponding parameter name and returns the struct.
func (r *ListEngineLicensesOptionalParameters) WithEngineName(engineName string) *ListEngineLicensesOptionalParameters {
	r.EngineName = &engineName
	return r
}

// ListEngineLicenses List all engineLicenses.
// list all engineLicenses
// NODESCRIPTION ListEngineLicenses
// Deprecated: This API is deprecated.
func (a *EngineLicenseApi) ListEngineLicenses(ctx _context.Context, o ...ListEngineLicensesOptionalParameters) (EngineLicenseList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue EngineLicenseList
		optionalParams      ListEngineLicensesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListEngineLicensesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EngineLicenseApi.ListEngineLicenses")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/engineLicenses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.EngineName != nil {
		localVarQueryParams.Add("engineName", common.ParameterToString(*optionalParams.EngineName, ""))
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

// NewEngineLicenseApi Returns NewEngineLicenseApi.
func NewEngineLicenseApi(client *common.APIClient) *EngineLicenseApi {
	return &EngineLicenseApi{
		Client: client,
	}
}
