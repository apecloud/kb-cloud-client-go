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

// ParamTplApi service type
type ParamTplApi common.Service

// CreateParamTpl Create parameter template.
func (a *ParamTplApi) CreateParamTpl(ctx _context.Context, orgName string, body ParamTplCreate) (ParamTplListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue ParamTplListItem
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParamTplApi.CreateParamTpl")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/paramTpls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("orgName", common.ParameterToString(orgName, ""))
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

// CreateParamTplFromCluster Export configuration template from cluster.
func (a *ParamTplApi) CreateParamTplFromCluster(ctx _context.Context, orgName string, clusterName string, body ParamTplCreateFromCluster) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParamTplApi.CreateParamTplFromCluster")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/paramTpls"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// DeleteParamTpl Delete configuration template.
func (a *ParamTplApi) DeleteParamTpl(ctx _context.Context, orgName string, paramTplName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParamTplApi.DeleteParamTpl")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/paramTpls/{paramTplName}"
	localVarPath = strings.Replace(localVarPath, "{"+"paramTplName"+"}", _neturl.PathEscape(common.ParameterToString(paramTplName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("orgName", common.ParameterToString(orgName, ""))
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// GetClusterParamTplsOptionalParameters holds optional parameters for GetClusterParamTpls.
type GetClusterParamTplsOptionalParameters struct {
	Component  *string
	EngineName *string
}

// NewGetClusterParamTplsOptionalParameters creates an empty struct for parameters.
func NewGetClusterParamTplsOptionalParameters() *GetClusterParamTplsOptionalParameters {
	this := GetClusterParamTplsOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *GetClusterParamTplsOptionalParameters) WithComponent(component string) *GetClusterParamTplsOptionalParameters {
	r.Component = &component
	return r
}

// WithEngineName sets the corresponding parameter name and returns the struct.
func (r *GetClusterParamTplsOptionalParameters) WithEngineName(engineName string) *GetClusterParamTplsOptionalParameters {
	r.EngineName = &engineName
	return r
}

// GetClusterParamTpls Get cluster configuration templates.
func (a *ParamTplApi) GetClusterParamTpls(ctx _context.Context, orgName string, clusterName string, o ...GetClusterParamTplsOptionalParameters) (ParamTplApplToClusterList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ParamTplApplToClusterList
		optionalParams      GetClusterParamTplsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetClusterParamTplsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParamTplApi.GetClusterParamTpls")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/paramTpls"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
	}
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

// ListParamTplsOptionalParameters holds optional parameters for ListParamTpls.
type ListParamTplsOptionalParameters struct {
	OrgName    *string
	Partition  *ParamTplPartition
	Version    *string
	Component  *string
	EngineName *string
}

// NewListParamTplsOptionalParameters creates an empty struct for parameters.
func NewListParamTplsOptionalParameters() *ListParamTplsOptionalParameters {
	this := ListParamTplsOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *ListParamTplsOptionalParameters) WithOrgName(orgName string) *ListParamTplsOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithPartition sets the corresponding parameter name and returns the struct.
func (r *ListParamTplsOptionalParameters) WithPartition(partition ParamTplPartition) *ListParamTplsOptionalParameters {
	r.Partition = &partition
	return r
}

// WithVersion sets the corresponding parameter name and returns the struct.
func (r *ListParamTplsOptionalParameters) WithVersion(version string) *ListParamTplsOptionalParameters {
	r.Version = &version
	return r
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *ListParamTplsOptionalParameters) WithComponent(component string) *ListParamTplsOptionalParameters {
	r.Component = &component
	return r
}

// WithEngineName sets the corresponding parameter name and returns the struct.
func (r *ListParamTplsOptionalParameters) WithEngineName(engineName string) *ListParamTplsOptionalParameters {
	r.EngineName = &engineName
	return r
}

// ListParamTpls List parameter templates.
func (a *ParamTplApi) ListParamTpls(ctx _context.Context, o ...ListParamTplsOptionalParameters) (ParamTplList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ParamTplList
		optionalParams      ListParamTplsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListParamTplsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParamTplApi.ListParamTpls")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/paramTpls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.OrgName != nil {
		localVarQueryParams.Add("orgName", common.ParameterToString(*optionalParams.OrgName, ""))
	}
	if optionalParams.Partition != nil {
		localVarQueryParams.Add("partition", common.ParameterToString(*optionalParams.Partition, ""))
	}
	if optionalParams.Version != nil {
		localVarQueryParams.Add("version", common.ParameterToString(*optionalParams.Version, ""))
	}
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
	}
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

// PatchParamTpl Update configuration template.
// partially update the parameter template
func (a *ParamTplApi) PatchParamTpl(ctx _context.Context, orgName string, paramTplName string, body ParamTplUpdate) (ParamTplListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue ParamTplListItem
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParamTplApi.PatchParamTpl")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/paramTpls/{paramTplName}"
	localVarPath = strings.Replace(localVarPath, "{"+"paramTplName"+"}", _neturl.PathEscape(common.ParameterToString(paramTplName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("orgName", common.ParameterToString(orgName, ""))
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

// ReadParamTplOptionalParameters holds optional parameters for ReadParamTpl.
type ReadParamTplOptionalParameters struct {
	Partition *ParamTplPartition
}

// NewReadParamTplOptionalParameters creates an empty struct for parameters.
func NewReadParamTplOptionalParameters() *ReadParamTplOptionalParameters {
	this := ReadParamTplOptionalParameters{}
	return &this
}

// WithPartition sets the corresponding parameter name and returns the struct.
func (r *ReadParamTplOptionalParameters) WithPartition(partition ParamTplPartition) *ReadParamTplOptionalParameters {
	r.Partition = &partition
	return r
}

// ReadParamTpl Get parameter template details.
func (a *ParamTplApi) ReadParamTpl(ctx _context.Context, orgName string, paramTplName string, o ...ReadParamTplOptionalParameters) (ParamTplGet, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ParamTplGet
		optionalParams      ReadParamTplOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ReadParamTplOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParamTplApi.ReadParamTpl")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/paramTpls/{paramTplName}"
	localVarPath = strings.Replace(localVarPath, "{"+"paramTplName"+"}", _neturl.PathEscape(common.ParameterToString(paramTplName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("orgName", common.ParameterToString(orgName, ""))
	if optionalParams.Partition != nil {
		localVarQueryParams.Add("partition", common.ParameterToString(*optionalParams.Partition, ""))
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

// NewParamTplApi Returns NewParamTplApi.
func NewParamTplApi(client *common.APIClient) *ParamTplApi {
	return &ParamTplApi{
		Client: client,
	}
}
