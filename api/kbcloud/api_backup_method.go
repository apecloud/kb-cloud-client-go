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

// BackupMethodApi service type
type BackupMethodApi common.Service

// GetBackupMethodOptionalParameters holds optional parameters for GetBackupMethod.
type GetBackupMethodOptionalParameters struct {
	ClusterId           *string
	EnablePitr          *bool
	WithRebuildInstance *bool
	Component           *string
}

// NewGetBackupMethodOptionalParameters creates an empty struct for parameters.
func NewGetBackupMethodOptionalParameters() *GetBackupMethodOptionalParameters {
	this := GetBackupMethodOptionalParameters{}
	return &this
}

// WithClusterId sets the corresponding parameter name and returns the struct.
func (r *GetBackupMethodOptionalParameters) WithClusterId(clusterId string) *GetBackupMethodOptionalParameters {
	r.ClusterId = &clusterId
	return r
}

// WithEnablePitr sets the corresponding parameter name and returns the struct.
func (r *GetBackupMethodOptionalParameters) WithEnablePitr(enablePitr bool) *GetBackupMethodOptionalParameters {
	r.EnablePitr = &enablePitr
	return r
}

// WithWithRebuildInstance sets the corresponding parameter name and returns the struct.
func (r *GetBackupMethodOptionalParameters) WithWithRebuildInstance(withRebuildInstance bool) *GetBackupMethodOptionalParameters {
	r.WithRebuildInstance = &withRebuildInstance
	return r
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *GetBackupMethodOptionalParameters) WithComponent(component string) *GetBackupMethodOptionalParameters {
	r.Component = &component
	return r
}

// GetBackupMethod get backup method.
func (a *BackupMethodApi) GetBackupMethod(ctx _context.Context, orgName string, engineName string, o ...GetBackupMethodOptionalParameters) (ClusterBackupMethod, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ClusterBackupMethod
		optionalParams      GetBackupMethodOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetBackupMethodOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".BackupMethodApi.GetBackupMethod")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/{engineName}/clusterBackupMethod"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.ClusterId != nil {
		localVarQueryParams.Add("clusterID", common.ParameterToString(*optionalParams.ClusterId, ""))
	}
	if optionalParams.EnablePitr != nil {
		localVarQueryParams.Add("enablePITR", common.ParameterToString(*optionalParams.EnablePitr, ""))
	}
	if optionalParams.WithRebuildInstance != nil {
		localVarQueryParams.Add("withRebuildInstance", common.ParameterToString(*optionalParams.WithRebuildInstance, ""))
	}
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

// NewBackupMethodApi Returns NewBackupMethodApi.
func NewBackupMethodApi(client *common.APIClient) *BackupMethodApi {
	return &BackupMethodApi{
		Client: client,
	}
}
