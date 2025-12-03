// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"context"
	_context "context"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnableServiceVersionApi service type
type EnableServiceVersionApi common.Service

// EnableServiceVersionOptionalParameters holds optional parameters for EnableServiceVersion.
type EnableServiceVersionOptionalParameters struct {
	Image *_io.Reader
}

// NewEnableServiceVersionOptionalParameters creates an empty struct for parameters.
func NewEnableServiceVersionOptionalParameters() *EnableServiceVersionOptionalParameters {
	this := EnableServiceVersionOptionalParameters{}
	return &this
}

// WithImage sets the corresponding parameter name and returns the struct.
func (r *EnableServiceVersionOptionalParameters) WithImage(image _io.Reader) *EnableServiceVersionOptionalParameters {
	r.Image = &image
	return r
}

// EnableServiceVersion enable the service version of the engine.
// enable the service version of the engine and create related resources
func (a *EnableServiceVersionApi) EnableServiceVersion(ctx _context.Context, environmentName string, engineName string, body interface{}, o ...EnableServiceVersionOptionalParameters) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue interface{}
		optionalParams      EnableServiceVersionOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type EnableServiceVersionOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "enableServiceVersion",
		OperationID: "EnableServiceVersion",
		Path:        "/admin/v1/environments/{environmentName}/engines/{engineName}/serviceVersion",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".EnableServiceVersionApi.EnableServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/environments/{environmentName}/engines/{engineName}/serviceVersion"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentName"+"}", _neturl.PathEscape(common.ParameterToString(environmentName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	formFile := common.FormFile{}
	formFile.FormFileName = "image"
	var localVarFile _io.Reader
	if optionalParams.Image != nil {
		localVarFile = *optionalParams.Image
	}
	if localVarFile != nil {
		fbs, _ := _io.ReadAll(localVarFile)
		formFile.FileBytes = fbs
	}

	localVarFormParams, err = common.BuildFormParams(body)
	if err != nil {
		return localVarReturnValue, nil, common.ReportError("Failed to build form params: %s", err.Error())
	}
	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"BearerToken", "authorization"},
	)
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, &formFile)
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 {
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

// NewEnableServiceVersionApi Returns NewEnableServiceVersionApi.
func NewEnableServiceVersionApi(client *common.APIClient) *EnableServiceVersionApi {
	return &EnableServiceVersionApi{
		Client: client,
	}
}
