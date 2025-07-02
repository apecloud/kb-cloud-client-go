// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"context"
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceApi service type
type InstanceApi common.Service

// ListInstancesOptionalParameters holds optional parameters for ListInstances.
type ListInstancesOptionalParameters struct {
	EnvName     *string
	NodeName    *string
	OrgName     *string
	ClusterName *string
}

// NewListInstancesOptionalParameters creates an empty struct for parameters.
func NewListInstancesOptionalParameters() *ListInstancesOptionalParameters {
	this := ListInstancesOptionalParameters{}
	return &this
}

// WithEnvName sets the corresponding parameter name and returns the struct.
func (r *ListInstancesOptionalParameters) WithEnvName(envName string) *ListInstancesOptionalParameters {
	r.EnvName = &envName
	return r
}

// WithNodeName sets the corresponding parameter name and returns the struct.
func (r *ListInstancesOptionalParameters) WithNodeName(nodeName string) *ListInstancesOptionalParameters {
	r.NodeName = &nodeName
	return r
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *ListInstancesOptionalParameters) WithOrgName(orgName string) *ListInstancesOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithClusterName sets the corresponding parameter name and returns the struct.
func (r *ListInstancesOptionalParameters) WithClusterName(clusterName string) *ListInstancesOptionalParameters {
	r.ClusterName = &clusterName
	return r
}

// ListInstances List cluster instances.
// List instances based on query parameters with constraints.
// Deprecated: This API is deprecated.
func (a *InstanceApi) ListInstances(ctx _context.Context, o ...ListInstancesOptionalParameters) (InstanceList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue InstanceList
		optionalParams      ListInstancesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListInstancesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "instance",
		OperationID: "listInstances",
		Path:        "/admin/v1/instances",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".InstanceApi.ListInstances")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/instances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.EnvName != nil {
		localVarQueryParams.Add("envName", common.ParameterToString(*optionalParams.EnvName, ""))
	}
	if optionalParams.NodeName != nil {
		localVarQueryParams.Add("nodeName", common.ParameterToString(*optionalParams.NodeName, ""))
	}
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.ClusterName != nil {
		localVarQueryParams.Add("clusterName", common.ParameterToString(*optionalParams.ClusterName, ""))
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

// NewInstanceApi Returns NewInstanceApi.
func NewInstanceApi(client *common.APIClient) *InstanceApi {
	return &InstanceApi{
		Client: client,
	}
}
