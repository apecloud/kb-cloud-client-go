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

// ParameterTemplateApi service type
type ParameterTemplateApi common.Service

// CreateParameterTemplate Create parameter template.
func (a *ParameterTemplateApi) CreateParameterTemplate(ctx _context.Context, orgName string, body ParamTplCreate) (ParamTplListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue ParamTplListItem
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "parameterTemplate",
		OperationID: "createParameterTemplate",
		Path:        "/admin/v1/parameterTemplates",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParameterTemplateApi.CreateParameterTemplate")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/parameterTemplates"

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

// DeleteParameterTemplate Delete parameter template.
func (a *ParameterTemplateApi) DeleteParameterTemplate(ctx _context.Context, orgName string, parameterTemplateName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "parameterTemplate",
		OperationID: "deleteParameterTemplate",
		Path:        "/admin/v1/parameterTemplate/{parameterTemplateName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParameterTemplateApi.DeleteParameterTemplate")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/parameterTemplate/{parameterTemplateName}"
	localVarPath = strings.Replace(localVarPath, "{"+"parameterTemplateName"+"}", _neturl.PathEscape(common.ParameterToString(parameterTemplateName, "")), -1)

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

// ExportParameterTemplateFromCluster Export parameter template from cluster.
func (a *ParameterTemplateApi) ExportParameterTemplateFromCluster(ctx _context.Context, orgName string, clusterName string, body ParamTplCreateFromCluster) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "parameterTemplate",
		OperationID: "exportParameterTemplateFromCluster",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/parameterTemplate",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParameterTemplateApi.ExportParameterTemplateFromCluster")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/parameterTemplate"
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

// GetClusterParameterTemplateOptionalParameters holds optional parameters for GetClusterParameterTemplate.
type GetClusterParameterTemplateOptionalParameters struct {
	Component  *string
	EngineName *string
}

// NewGetClusterParameterTemplateOptionalParameters creates an empty struct for parameters.
func NewGetClusterParameterTemplateOptionalParameters() *GetClusterParameterTemplateOptionalParameters {
	this := GetClusterParameterTemplateOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *GetClusterParameterTemplateOptionalParameters) WithComponent(component string) *GetClusterParameterTemplateOptionalParameters {
	r.Component = &component
	return r
}

// WithEngineName sets the corresponding parameter name and returns the struct.
func (r *GetClusterParameterTemplateOptionalParameters) WithEngineName(engineName string) *GetClusterParameterTemplateOptionalParameters {
	r.EngineName = &engineName
	return r
}

// GetClusterParameterTemplate Get cluster parameter template.
func (a *ParameterTemplateApi) GetClusterParameterTemplate(ctx _context.Context, orgName string, clusterName string, o ...GetClusterParameterTemplateOptionalParameters) (ParamTplApplyToClusterList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ParamTplApplyToClusterList
		optionalParams      GetClusterParameterTemplateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetClusterParameterTemplateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "parameterTemplate",
		OperationID: "getClusterParameterTemplate",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/parameterTemplate",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParameterTemplateApi.GetClusterParameterTemplate")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/parameterTemplate"
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

// ListParameterTemplatesOptionalParameters holds optional parameters for ListParameterTemplates.
type ListParameterTemplatesOptionalParameters struct {
	OrgName    *string
	Partition  *ParameterTemplatePartition
	Version    *string
	Component  *string
	EngineName *string
	EngineMode *string
}

// NewListParameterTemplatesOptionalParameters creates an empty struct for parameters.
func NewListParameterTemplatesOptionalParameters() *ListParameterTemplatesOptionalParameters {
	this := ListParameterTemplatesOptionalParameters{}
	return &this
}

// WithOrgName sets the corresponding parameter name and returns the struct.
func (r *ListParameterTemplatesOptionalParameters) WithOrgName(orgName string) *ListParameterTemplatesOptionalParameters {
	r.OrgName = &orgName
	return r
}

// WithPartition sets the corresponding parameter name and returns the struct.
func (r *ListParameterTemplatesOptionalParameters) WithPartition(partition ParameterTemplatePartition) *ListParameterTemplatesOptionalParameters {
	r.Partition = &partition
	return r
}

// WithVersion sets the corresponding parameter name and returns the struct.
func (r *ListParameterTemplatesOptionalParameters) WithVersion(version string) *ListParameterTemplatesOptionalParameters {
	r.Version = &version
	return r
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *ListParameterTemplatesOptionalParameters) WithComponent(component string) *ListParameterTemplatesOptionalParameters {
	r.Component = &component
	return r
}

// WithEngineName sets the corresponding parameter name and returns the struct.
func (r *ListParameterTemplatesOptionalParameters) WithEngineName(engineName string) *ListParameterTemplatesOptionalParameters {
	r.EngineName = &engineName
	return r
}

// WithEngineMode sets the corresponding parameter name and returns the struct.
func (r *ListParameterTemplatesOptionalParameters) WithEngineMode(engineMode string) *ListParameterTemplatesOptionalParameters {
	r.EngineMode = &engineMode
	return r
}

// ListParameterTemplates List parameter templates.
func (a *ParameterTemplateApi) ListParameterTemplates(ctx _context.Context, o ...ListParameterTemplatesOptionalParameters) (ParamTplList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ParamTplList
		optionalParams      ListParameterTemplatesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListParameterTemplatesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "parameterTemplate",
		OperationID: "listParameterTemplates",
		Path:        "/admin/v1/parameterTemplates",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParameterTemplateApi.ListParameterTemplates")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/parameterTemplates"

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
	if optionalParams.EngineMode != nil {
		localVarQueryParams.Add("engineMode", common.ParameterToString(*optionalParams.EngineMode, ""))
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

// PatchParameterTemplate Update parameter template.
func (a *ParameterTemplateApi) PatchParameterTemplate(ctx _context.Context, orgName string, parameterTemplateName string, body ParamTplUpdate) (ParamTplListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue ParamTplListItem
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "parameterTemplate",
		OperationID: "patchParameterTemplate",
		Path:        "/admin/v1/parameterTemplate/{parameterTemplateName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParameterTemplateApi.PatchParameterTemplate")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/parameterTemplate/{parameterTemplateName}"
	localVarPath = strings.Replace(localVarPath, "{"+"parameterTemplateName"+"}", _neturl.PathEscape(common.ParameterToString(parameterTemplateName, "")), -1)

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

// ReadParameterTemplateOptionalParameters holds optional parameters for ReadParameterTemplate.
type ReadParameterTemplateOptionalParameters struct {
	Partition  *ParameterTemplatePartition
	RawContent *bool
}

// NewReadParameterTemplateOptionalParameters creates an empty struct for parameters.
func NewReadParameterTemplateOptionalParameters() *ReadParameterTemplateOptionalParameters {
	this := ReadParameterTemplateOptionalParameters{}
	return &this
}

// WithPartition sets the corresponding parameter name and returns the struct.
func (r *ReadParameterTemplateOptionalParameters) WithPartition(partition ParameterTemplatePartition) *ReadParameterTemplateOptionalParameters {
	r.Partition = &partition
	return r
}

// WithRawContent sets the corresponding parameter name and returns the struct.
func (r *ReadParameterTemplateOptionalParameters) WithRawContent(rawContent bool) *ReadParameterTemplateOptionalParameters {
	r.RawContent = &rawContent
	return r
}

// ReadParameterTemplate Get parameter template details.
func (a *ParameterTemplateApi) ReadParameterTemplate(ctx _context.Context, orgName string, parameterTemplateName string, o ...ReadParameterTemplateOptionalParameters) (ParameterList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ParameterList
		optionalParams      ReadParameterTemplateOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ReadParameterTemplateOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "parameterTemplate",
		OperationID: "readParameterTemplate",
		Path:        "/admin/v1/parameterTemplate/{parameterTemplateName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParameterTemplateApi.ReadParameterTemplate")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/parameterTemplate/{parameterTemplateName}"
	localVarPath = strings.Replace(localVarPath, "{"+"parameterTemplateName"+"}", _neturl.PathEscape(common.ParameterToString(parameterTemplateName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("orgName", common.ParameterToString(orgName, ""))
	if optionalParams.Partition != nil {
		localVarQueryParams.Add("partition", common.ParameterToString(*optionalParams.Partition, ""))
	}
	if optionalParams.RawContent != nil {
		localVarQueryParams.Add("rawContent", common.ParameterToString(*optionalParams.RawContent, ""))
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

// UpdateClusterParameterTemplate Update cluster parameter template.
func (a *ParameterTemplateApi) UpdateClusterParameterTemplate(ctx _context.Context, orgName string, clusterName string, body []ParamTplsItem) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "parameterTemplate",
		OperationID: "updateClusterParameterTemplate",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/parameterTemplate",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".ParameterTemplateApi.UpdateClusterParameterTemplate")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/parameterTemplate"
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

// NewParameterTemplateApi Returns NewParameterTemplateApi.
func NewParameterTemplateApi(client *common.APIClient) *ParameterTemplateApi {
	return &ParameterTemplateApi{
		Client: client,
	}
}
