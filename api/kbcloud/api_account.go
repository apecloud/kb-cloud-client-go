// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"context"
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AccountApi service type
type AccountApi common.Service

// BindSeaweedFSPolicies bindSeaweedFSPolicies.
// Replaces named policy bindings. An empty policies array removes all direct bindings. Existing credentials, legacy actions and other identity attributes are preserved.
func (a *AccountApi) BindSeaweedFSPolicies(ctx _context.Context, orgName string, clusterName string, accountName string, body interface{}) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPut
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "bindSeaweedFSPolicies",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/policies",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.BindSeaweedFSPolicies")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/policies"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "*/*"

	// body params
	localVarPostBody = &body
	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// ChangeAccountObjectPrivilegesOptionalParameters holds optional parameters for ChangeAccountObjectPrivileges.
type ChangeAccountObjectPrivilegesOptionalParameters struct {
	SubjectType       *ObjectPrivilegeSubjectType
	Host              *string
	SubjectQualifiers *string
}

// NewChangeAccountObjectPrivilegesOptionalParameters creates an empty struct for parameters.
func NewChangeAccountObjectPrivilegesOptionalParameters() *ChangeAccountObjectPrivilegesOptionalParameters {
	this := ChangeAccountObjectPrivilegesOptionalParameters{}
	return &this
}

// WithSubjectType sets the corresponding parameter name and returns the struct.
func (r *ChangeAccountObjectPrivilegesOptionalParameters) WithSubjectType(subjectType ObjectPrivilegeSubjectType) *ChangeAccountObjectPrivilegesOptionalParameters {
	r.SubjectType = &subjectType
	return r
}

// WithHost sets the corresponding parameter name and returns the struct.
func (r *ChangeAccountObjectPrivilegesOptionalParameters) WithHost(host string) *ChangeAccountObjectPrivilegesOptionalParameters {
	r.Host = &host
	return r
}

// WithSubjectQualifiers sets the corresponding parameter name and returns the struct.
func (r *ChangeAccountObjectPrivilegesOptionalParameters) WithSubjectQualifiers(subjectQualifiers string) *ChangeAccountObjectPrivilegesOptionalParameters {
	r.SubjectQualifiers = &subjectQualifiers
	return r
}

// ChangeAccountObjectPrivileges Grant or revoke native object and column privileges.
// Native object GRANT/REVOKE alongside presets. Consult capabilities for engine/version-specific object types and privileges. PostgreSQL RESTRICT is the default; MySQL static GRANT OPTION is scope-wide.
func (a *AccountApi) ChangeAccountObjectPrivileges(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, body ObjectPrivilegeChange, o ...ChangeAccountObjectPrivilegesOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
		optionalParams     ChangeAccountObjectPrivilegesOptionalParameters
	)

	if len(o) > 1 {
		return nil, common.ReportError("only one argument of type ChangeAccountObjectPrivilegesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "changeAccountObjectPrivileges",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ChangeAccountObjectPrivileges")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.SubjectType != nil {
		localVarQueryParams.Add("subjectType", common.ParameterToString(*optionalParams.SubjectType, ""))
	}
	if optionalParams.Host != nil {
		localVarQueryParams.Add("host", common.ParameterToString(*optionalParams.Host, ""))
	}
	if optionalParams.SubjectQualifiers != nil {
		localVarQueryParams.Add("subjectQualifiers", common.ParameterToString(*optionalParams.SubjectQualifiers, ""))
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "*/*"

	// body params
	localVarPostBody = &body
	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 500 {
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

// CreateAccount Create cluster account.
// create an account in cluster
func (a *AccountApi) CreateAccount(ctx _context.Context, engineName string, orgName string, clusterName string, body Account) (Account, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Account
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "createAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.CreateAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
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
		[2]string{"DigestAuth", "Authorization"},
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

// CreateHiveAccount Create Hive account.
// create an hive account
func (a *AccountApi) CreateHiveAccount(ctx _context.Context, orgName string, clusterName string, body Account) (Account, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Account
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "createHiveAccount",
		Path:        "/api/v1/data/hive/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.CreateHiveAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/hive/organizations/{orgName}/clusters/{clusterName}/accounts"
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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
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

// CreateKafkaAccount Create cluster account.
// create an account in cluster
func (a *AccountApi) CreateKafkaAccount(ctx _context.Context, orgName string, clusterName string, body Account) (Account, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Account
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "createKafkaAccount",
		Path:        "/api/v1/data/kafka/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.CreateKafkaAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/kafka/organizations/{orgName}/clusters/{clusterName}/accounts"
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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 || localVarHTTPResponse.StatusCode == 501 {
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

// CreateMilvusRole Create Milvus role.
// Create a Milvus native role and grant its initial privileges.
func (a *AccountApi) CreateMilvusRole(ctx _context.Context, orgName string, clusterName string, body MilvusRoleCreateRequest) (MilvusRole, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue MilvusRole
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "createMilvusRole",
		Path:        "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.CreateMilvusRole")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles"
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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 500 {
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

// CreateMongoDBAccount Create mongodb account.
// create an account in mongodb
func (a *AccountApi) CreateMongoDBAccount(ctx _context.Context, orgName string, clusterName string, body Account) (Account, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Account
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "createMongoDBAccount",
		Path:        "/api/v1/data/mongodb/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.CreateMongoDBAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/mongodb/organizations/{orgName}/clusters/{clusterName}/accounts"
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
		[2]string{"DigestAuth", "Authorization"},
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

// CreateSeaweedFSAccessKey createSeaweedFSAccessKey.
// Creates a credential on the identity. Account.name must match accountName, role must be BASICUSER. Optional extra.accessKey and password select the access key and secret key; omitted or empty values are generated. The returned password is disclosed only in this response.
func (a *AccountApi) CreateSeaweedFSAccessKey(ctx _context.Context, orgName string, clusterName string, accountName string, body Account) (Account, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Account
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "createSeaweedFSAccessKey",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/access-keys",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.CreateSeaweedFSAccessKey")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/access-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// CreateSeaweedFSAccount createSeaweedFSAccount.
// Creates a BASICUSER identity with named policies from extra.policies. Optional extra.accessKey and password select the credential; empty values are generated. The password is disclosed only in this response.
func (a *AccountApi) CreateSeaweedFSAccount(ctx _context.Context, orgName string, clusterName string, body Account) (Account, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue Account
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "createSeaweedFSAccount",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.CreateSeaweedFSAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts"
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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// CreateServiceAccount Create MinIO access key.
// create a MinIO service account access key
func (a *AccountApi) CreateServiceAccount(ctx _context.Context, engineName string, orgName string, clusterName string, body DmsServiceAccountCreateRequest) (DmsServiceAccount, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue DmsServiceAccount
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "createServiceAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.CreateServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
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
		[2]string{"DigestAuth", "Authorization"},
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

// DeleteAccount Delete cluster account.
// delete an account in cluster
func (a *AccountApi) DeleteAccount(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// DeleteHiveAccount Delete hive account.
// delete an hive account in cluster
func (a *AccountApi) DeleteHiveAccount(ctx _context.Context, orgName string, clusterName string, accountName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteHiveAccount",
		Path:        "/api/v1/data/hive/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteHiveAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/hive/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// DeleteKafkaAccount Delete cluster account.
// delete an account in cluster
func (a *AccountApi) DeleteKafkaAccount(ctx _context.Context, orgName string, clusterName string, accountName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteKafkaAccount",
		Path:        "/api/v1/data/kafka/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteKafkaAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/kafka/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 || localVarHTTPResponse.StatusCode == 501 {
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

// DeleteMilvusRole Delete Milvus role.
// Delete a Milvus native role.
func (a *AccountApi) DeleteMilvusRole(ctx _context.Context, orgName string, clusterName string, roleName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteMilvusRole",
		Path:        "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles/{roleName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteMilvusRole")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles/{roleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", _neturl.PathEscape(common.ParameterToString(roleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// DeleteMongoDBAccount Delete mongodb account.
// delete an account in mongodb
func (a *AccountApi) DeleteMongoDBAccount(ctx _context.Context, orgName string, clusterName string, accountName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteMongoDBAccount",
		Path:        "/api/v1/data/mongodb/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteMongoDBAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/mongodb/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// DeleteMssqlAccount delete mssql account compatible with windows account.
func (a *AccountApi) DeleteMssqlAccount(ctx _context.Context, orgName string, clusterName string, account string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteMssqlAccount",
		Path:        "/api/v1/data/mssql/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteMssqlAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/mssql/organizations/{orgName}/clusters/{clusterName}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("account", common.ParameterToString(account, ""))
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// DeleteSeaweedFSAccessKey deleteSeaweedFSAccessKey.
func (a *AccountApi) DeleteSeaweedFSAccessKey(ctx _context.Context, orgName string, clusterName string, accountName string, accessKey string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteSeaweedFSAccessKey",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/access-keys/{accessKey}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteSeaweedFSAccessKey")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/access-keys/{accessKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessKey"+"}", _neturl.PathEscape(common.ParameterToString(accessKey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// DeleteSeaweedFSAccount deleteSeaweedFSAccount.
// delete an account in cluster
func (a *AccountApi) DeleteSeaweedFSAccount(ctx _context.Context, orgName string, clusterName string, accountName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteSeaweedFSAccount",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteSeaweedFSAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// DeleteSeaweedFSPolicy deleteSeaweedFSPolicy.
// Deletes an unused named policy. Attached policies are rejected; no identities or groups are implicitly changed.
func (a *AccountApi) DeleteSeaweedFSPolicy(ctx _context.Context, orgName string, clusterName string, policyName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteSeaweedFSPolicy",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/policies/{policyName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteSeaweedFSPolicy")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/policies/{policyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyName"+"}", _neturl.PathEscape(common.ParameterToString(policyName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// DeleteServiceAccount Delete MinIO access key.
// delete a MinIO service account access key
func (a *AccountApi) DeleteServiceAccount(ctx _context.Context, engineName string, orgName string, clusterName string, accessKey string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "deleteServiceAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys/{accessKey}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.DeleteServiceAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys/{accessKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessKey"+"}", _neturl.PathEscape(common.ParameterToString(accessKey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// GetAccountObjectPrivilegeCapabilities List native object privilege capabilities for the connected server.
// Native object GRANT/REVOKE alongside presets. Consult capabilities for engine/version-specific object types and privileges. PostgreSQL RESTRICT is the default; MySQL static GRANT OPTION is scope-wide.
func (a *AccountApi) GetAccountObjectPrivilegeCapabilities(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string) (ObjectPrivilegeCapabilities, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ObjectPrivilegeCapabilities
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "getAccountObjectPrivilegeCapabilities",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/capabilities",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.GetAccountObjectPrivilegeCapabilities")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/capabilities"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// GetAccountObjectPrivilegesOptionalParameters holds optional parameters for GetAccountObjectPrivileges.
type GetAccountObjectPrivilegesOptionalParameters struct {
	DatabaseName      *string
	SchemaName        *string
	ObjectName        *string
	ObjectId          *string
	SubjectType       *ObjectPrivilegeSubjectType
	Host              *string
	CatalogName       *string
	SubjectQualifiers *string
}

// NewGetAccountObjectPrivilegesOptionalParameters creates an empty struct for parameters.
func NewGetAccountObjectPrivilegesOptionalParameters() *GetAccountObjectPrivilegesOptionalParameters {
	this := GetAccountObjectPrivilegesOptionalParameters{}
	return &this
}

// WithDatabaseName sets the corresponding parameter name and returns the struct.
func (r *GetAccountObjectPrivilegesOptionalParameters) WithDatabaseName(databaseName string) *GetAccountObjectPrivilegesOptionalParameters {
	r.DatabaseName = &databaseName
	return r
}

// WithSchemaName sets the corresponding parameter name and returns the struct.
func (r *GetAccountObjectPrivilegesOptionalParameters) WithSchemaName(schemaName string) *GetAccountObjectPrivilegesOptionalParameters {
	r.SchemaName = &schemaName
	return r
}

// WithObjectName sets the corresponding parameter name and returns the struct.
func (r *GetAccountObjectPrivilegesOptionalParameters) WithObjectName(objectName string) *GetAccountObjectPrivilegesOptionalParameters {
	r.ObjectName = &objectName
	return r
}

// WithObjectId sets the corresponding parameter name and returns the struct.
func (r *GetAccountObjectPrivilegesOptionalParameters) WithObjectId(objectId string) *GetAccountObjectPrivilegesOptionalParameters {
	r.ObjectId = &objectId
	return r
}

// WithSubjectType sets the corresponding parameter name and returns the struct.
func (r *GetAccountObjectPrivilegesOptionalParameters) WithSubjectType(subjectType ObjectPrivilegeSubjectType) *GetAccountObjectPrivilegesOptionalParameters {
	r.SubjectType = &subjectType
	return r
}

// WithHost sets the corresponding parameter name and returns the struct.
func (r *GetAccountObjectPrivilegesOptionalParameters) WithHost(host string) *GetAccountObjectPrivilegesOptionalParameters {
	r.Host = &host
	return r
}

// WithCatalogName sets the corresponding parameter name and returns the struct.
func (r *GetAccountObjectPrivilegesOptionalParameters) WithCatalogName(catalogName string) *GetAccountObjectPrivilegesOptionalParameters {
	r.CatalogName = &catalogName
	return r
}

// WithSubjectQualifiers sets the corresponding parameter name and returns the struct.
func (r *GetAccountObjectPrivilegesOptionalParameters) WithSubjectQualifiers(subjectQualifiers string) *GetAccountObjectPrivilegesOptionalParameters {
	r.SubjectQualifiers = &subjectQualifiers
	return r
}

// GetAccountObjectPrivileges Inspect native object and column privileges.
// Native object GRANT/REVOKE alongside presets. Consult capabilities for engine/version-specific object types and privileges. PostgreSQL RESTRICT is the default; MySQL static GRANT OPTION is scope-wide.
func (a *AccountApi) GetAccountObjectPrivileges(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, objectType string, o ...GetAccountObjectPrivilegesOptionalParameters) (ObjectPrivilegeSnapshot, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ObjectPrivilegeSnapshot
		optionalParams      GetAccountObjectPrivilegesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetAccountObjectPrivilegesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "getAccountObjectPrivileges",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.GetAccountObjectPrivileges")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("objectType", common.ParameterToString(objectType, ""))
	if optionalParams.DatabaseName != nil {
		localVarQueryParams.Add("databaseName", common.ParameterToString(*optionalParams.DatabaseName, ""))
	}
	if optionalParams.SchemaName != nil {
		localVarQueryParams.Add("schemaName", common.ParameterToString(*optionalParams.SchemaName, ""))
	}
	if optionalParams.ObjectName != nil {
		localVarQueryParams.Add("objectName", common.ParameterToString(*optionalParams.ObjectName, ""))
	}
	if optionalParams.ObjectId != nil {
		localVarQueryParams.Add("objectId", common.ParameterToString(*optionalParams.ObjectId, ""))
	}
	if optionalParams.SubjectType != nil {
		localVarQueryParams.Add("subjectType", common.ParameterToString(*optionalParams.SubjectType, ""))
	}
	if optionalParams.Host != nil {
		localVarQueryParams.Add("host", common.ParameterToString(*optionalParams.Host, ""))
	}
	if optionalParams.CatalogName != nil {
		localVarQueryParams.Add("catalogName", common.ParameterToString(*optionalParams.CatalogName, ""))
	}
	if optionalParams.SubjectQualifiers != nil {
		localVarQueryParams.Add("subjectQualifiers", common.ParameterToString(*optionalParams.SubjectQualifiers, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// GetAccountPrivilegeOverviewOptionalParameters holds optional parameters for GetAccountPrivilegeOverview.
type GetAccountPrivilegeOverviewOptionalParameters struct {
	SubjectType       *ObjectPrivilegeSubjectType
	Host              *string
	SubjectQualifiers *string
	DatabaseName      *string
}

// NewGetAccountPrivilegeOverviewOptionalParameters creates an empty struct for parameters.
func NewGetAccountPrivilegeOverviewOptionalParameters() *GetAccountPrivilegeOverviewOptionalParameters {
	this := GetAccountPrivilegeOverviewOptionalParameters{}
	return &this
}

// WithSubjectType sets the corresponding parameter name and returns the struct.
func (r *GetAccountPrivilegeOverviewOptionalParameters) WithSubjectType(subjectType ObjectPrivilegeSubjectType) *GetAccountPrivilegeOverviewOptionalParameters {
	r.SubjectType = &subjectType
	return r
}

// WithHost sets the corresponding parameter name and returns the struct.
func (r *GetAccountPrivilegeOverviewOptionalParameters) WithHost(host string) *GetAccountPrivilegeOverviewOptionalParameters {
	r.Host = &host
	return r
}

// WithSubjectQualifiers sets the corresponding parameter name and returns the struct.
func (r *GetAccountPrivilegeOverviewOptionalParameters) WithSubjectQualifiers(subjectQualifiers string) *GetAccountPrivilegeOverviewOptionalParameters {
	r.SubjectQualifiers = &subjectQualifiers
	return r
}

// WithDatabaseName sets the corresponding parameter name and returns the struct.
func (r *GetAccountPrivilegeOverviewOptionalParameters) WithDatabaseName(databaseName string) *GetAccountPrivilegeOverviewOptionalParameters {
	r.DatabaseName = &databaseName
	return r
}

// GetAccountPrivilegeOverview Read native account privilege summary or selected database evidence.
// Without databaseName, returns role attributes, memberships, database access summaries and shared ACLs without connecting to other databases. PostgreSQL databaseName reads schema, object, ownership and PUBLIC evidence from that database only. Failed reads return an error, never an empty grant list. SQL is read-only evidence, not an exhaustive effective-access decision.
func (a *AccountApi) GetAccountPrivilegeOverview(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, o ...GetAccountPrivilegeOverviewOptionalParameters) ([]ObjectPrivilegeEvidence, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []ObjectPrivilegeEvidence
		optionalParams      GetAccountPrivilegeOverviewOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetAccountPrivilegeOverviewOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "getAccountPrivilegeOverview",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/overview",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.GetAccountPrivilegeOverview")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/overview"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.SubjectType != nil {
		localVarQueryParams.Add("subjectType", common.ParameterToString(*optionalParams.SubjectType, ""))
	}
	if optionalParams.Host != nil {
		localVarQueryParams.Add("host", common.ParameterToString(*optionalParams.Host, ""))
	}
	if optionalParams.SubjectQualifiers != nil {
		localVarQueryParams.Add("subjectQualifiers", common.ParameterToString(*optionalParams.SubjectQualifiers, ""))
	}
	if optionalParams.DatabaseName != nil {
		localVarQueryParams.Add("databaseName", common.ParameterToString(*optionalParams.DatabaseName, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// GetMilvusRole Get Milvus role.
// Get a Milvus native role and its privilege grants.
func (a *AccountApi) GetMilvusRole(ctx _context.Context, orgName string, clusterName string, roleName string) (MilvusRole, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MilvusRole
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "getMilvusRole",
		Path:        "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles/{roleName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.GetMilvusRole")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles/{roleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", _neturl.PathEscape(common.ParameterToString(roleName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// GetRootAccountPasswordOptionalParameters holds optional parameters for GetRootAccountPassword.
type GetRootAccountPasswordOptionalParameters struct {
	Component *string
}

// NewGetRootAccountPasswordOptionalParameters creates an empty struct for parameters.
func NewGetRootAccountPasswordOptionalParameters() *GetRootAccountPasswordOptionalParameters {
	this := GetRootAccountPasswordOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *GetRootAccountPasswordOptionalParameters) WithComponent(component string) *GetRootAccountPasswordOptionalParameters {
	r.Component = &component
	return r
}

// GetRootAccountPassword get root account password.
// get root account password
func (a *AccountApi) GetRootAccountPassword(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, o ...GetRootAccountPasswordOptionalParameters) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue string
		optionalParams      GetRootAccountPasswordOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetRootAccountPasswordOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "getRootAccountPassword",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/root-password",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.GetRootAccountPassword")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/root-password"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("accountName", common.ParameterToString(accountName, ""))
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 500 {
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

// GetRootSeaweedFSAccountPasswordOptionalParameters holds optional parameters for GetRootSeaweedFSAccountPassword.
type GetRootSeaweedFSAccountPasswordOptionalParameters struct {
	Component *string
}

// NewGetRootSeaweedFSAccountPasswordOptionalParameters creates an empty struct for parameters.
func NewGetRootSeaweedFSAccountPasswordOptionalParameters() *GetRootSeaweedFSAccountPasswordOptionalParameters {
	this := GetRootSeaweedFSAccountPasswordOptionalParameters{}
	return &this
}

// WithComponent sets the corresponding parameter name and returns the struct.
func (r *GetRootSeaweedFSAccountPasswordOptionalParameters) WithComponent(component string) *GetRootSeaweedFSAccountPasswordOptionalParameters {
	r.Component = &component
	return r
}

// GetRootSeaweedFSAccountPassword getRootSeaweedFSAccountPassword.
// get root account password
func (a *AccountApi) GetRootSeaweedFSAccountPassword(ctx _context.Context, orgName string, clusterName string, accountName string, o ...GetRootSeaweedFSAccountPasswordOptionalParameters) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue string
		optionalParams      GetRootSeaweedFSAccountPasswordOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetRootSeaweedFSAccountPasswordOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "getRootSeaweedFSAccountPassword",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/root-password",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.GetRootSeaweedFSAccountPassword")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/root-password"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("accountName", common.ParameterToString(accountName, ""))
	if optionalParams.Component != nil {
		localVarQueryParams.Add("component", common.ParameterToString(*optionalParams.Component, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 500 {
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

// GetServiceAccount Get MinIO access key.
// get a MinIO service account access key
func (a *AccountApi) GetServiceAccount(ctx _context.Context, engineName string, orgName string, clusterName string, accessKey string) (DmsServiceAccount, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DmsServiceAccount
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "getServiceAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys/{accessKey}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.GetServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys/{accessKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessKey"+"}", _neturl.PathEscape(common.ParameterToString(accessKey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
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

// ListAccountObjectPrivilegeColumnsOptionalParameters holds optional parameters for ListAccountObjectPrivilegeColumns.
type ListAccountObjectPrivilegeColumnsOptionalParameters struct {
	SchemaName *string
}

// NewListAccountObjectPrivilegeColumnsOptionalParameters creates an empty struct for parameters.
func NewListAccountObjectPrivilegeColumnsOptionalParameters() *ListAccountObjectPrivilegeColumnsOptionalParameters {
	this := ListAccountObjectPrivilegeColumnsOptionalParameters{}
	return &this
}

// WithSchemaName sets the corresponding parameter name and returns the struct.
func (r *ListAccountObjectPrivilegeColumnsOptionalParameters) WithSchemaName(schemaName string) *ListAccountObjectPrivilegeColumnsOptionalParameters {
	r.SchemaName = &schemaName
	return r
}

// ListAccountObjectPrivilegeColumns List exact column names for account object privileges.
// Read MySQL or PostgreSQL relation column names using the account management connection. No workbench datasource is required. Names are returned verbatim for column-level grants.
func (a *AccountApi) ListAccountObjectPrivilegeColumns(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, objectType string, databaseName string, objectName string, o ...ListAccountObjectPrivilegeColumnsOptionalParameters) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []string
		optionalParams      ListAccountObjectPrivilegeColumnsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListAccountObjectPrivilegeColumnsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listAccountObjectPrivilegeColumns",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/columns",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListAccountObjectPrivilegeColumns")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/columns"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("objectType", common.ParameterToString(objectType, ""))
	localVarQueryParams.Add("databaseName", common.ParameterToString(databaseName, ""))
	localVarQueryParams.Add("objectName", common.ParameterToString(objectName, ""))
	if optionalParams.SchemaName != nil {
		localVarQueryParams.Add("schemaName", common.ParameterToString(*optionalParams.SchemaName, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// ListAccountObjectPrivilegeObjectsOptionalParameters holds optional parameters for ListAccountObjectPrivilegeObjects.
type ListAccountObjectPrivilegeObjectsOptionalParameters struct {
	DatabaseName *string
	SchemaName   *string
	Offset       *int64
	Limit        *int64
	CatalogName  *string
}

// NewListAccountObjectPrivilegeObjectsOptionalParameters creates an empty struct for parameters.
func NewListAccountObjectPrivilegeObjectsOptionalParameters() *ListAccountObjectPrivilegeObjectsOptionalParameters {
	this := ListAccountObjectPrivilegeObjectsOptionalParameters{}
	return &this
}

// WithDatabaseName sets the corresponding parameter name and returns the struct.
func (r *ListAccountObjectPrivilegeObjectsOptionalParameters) WithDatabaseName(databaseName string) *ListAccountObjectPrivilegeObjectsOptionalParameters {
	r.DatabaseName = &databaseName
	return r
}

// WithSchemaName sets the corresponding parameter name and returns the struct.
func (r *ListAccountObjectPrivilegeObjectsOptionalParameters) WithSchemaName(schemaName string) *ListAccountObjectPrivilegeObjectsOptionalParameters {
	r.SchemaName = &schemaName
	return r
}

// WithOffset sets the corresponding parameter name and returns the struct.
func (r *ListAccountObjectPrivilegeObjectsOptionalParameters) WithOffset(offset int64) *ListAccountObjectPrivilegeObjectsOptionalParameters {
	r.Offset = &offset
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *ListAccountObjectPrivilegeObjectsOptionalParameters) WithLimit(limit int64) *ListAccountObjectPrivilegeObjectsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithCatalogName sets the corresponding parameter name and returns the struct.
func (r *ListAccountObjectPrivilegeObjectsOptionalParameters) WithCatalogName(catalogName string) *ListAccountObjectPrivilegeObjectsOptionalParameters {
	r.CatalogName = &catalogName
	return r
}

// ListAccountObjectPrivilegeObjects Discover existing native privilege objects with stable routine identities.
// Lists existing objects of one supported kind. Pagination is not a fixed snapshot. PostgreSQL requires a connection database; schema-contained kinds require schemaName. MySQL global/database listing omits databaseName; other kinds require it.
func (a *AccountApi) ListAccountObjectPrivilegeObjects(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, objectType string, o ...ListAccountObjectPrivilegeObjectsOptionalParameters) (ObjectPrivilegeObjectList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ObjectPrivilegeObjectList
		optionalParams      ListAccountObjectPrivilegeObjectsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListAccountObjectPrivilegeObjectsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listAccountObjectPrivilegeObjects",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/objects",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListAccountObjectPrivilegeObjects")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/objects"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("objectType", common.ParameterToString(objectType, ""))
	if optionalParams.DatabaseName != nil {
		localVarQueryParams.Add("databaseName", common.ParameterToString(*optionalParams.DatabaseName, ""))
	}
	if optionalParams.SchemaName != nil {
		localVarQueryParams.Add("schemaName", common.ParameterToString(*optionalParams.SchemaName, ""))
	}
	if optionalParams.Offset != nil {
		localVarQueryParams.Add("offset", common.ParameterToString(*optionalParams.Offset, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.CatalogName != nil {
		localVarQueryParams.Add("catalogName", common.ParameterToString(*optionalParams.CatalogName, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// ListAccounts List cluster accounts.
// list accounts in cluster
func (a *AccountApi) ListAccounts(ctx _context.Context, engineName string, orgName string, clusterName string) ([]AccountListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []AccountListItem
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listAccounts",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListAccounts")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// ListKafkaAccounts List cluster accounts.
// list accounts in cluster
func (a *AccountApi) ListKafkaAccounts(ctx _context.Context, orgName string, clusterName string) ([]AccountListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []AccountListItem
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listKafkaAccounts",
		Path:        "/api/v1/data/kafka/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListKafkaAccounts")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/kafka/organizations/{orgName}/clusters/{clusterName}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 || localVarHTTPResponse.StatusCode == 501 {
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

// ListMilvusRoles List Milvus roles.
// List Milvus native roles and their privilege grants.
func (a *AccountApi) ListMilvusRoles(ctx _context.Context, orgName string, clusterName string) (MilvusRoleList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MilvusRoleList
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listMilvusRoles",
		Path:        "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListMilvusRoles")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// ListMongoDBAccounts List mongodb accounts.
// list accounts in mongodb
func (a *AccountApi) ListMongoDBAccounts(ctx _context.Context, orgName string, clusterName string) ([]AccountListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []AccountListItem
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listMongoDBAccounts",
		Path:        "/api/v1/data/mongodb/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListMongoDBAccounts")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/mongodb/organizations/{orgName}/clusters/{clusterName}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// ListSeaweedFSAccounts listSeaweedFSAccounts.
// Lists SeaweedFS identities and access-key metadata without secret keys.
func (a *AccountApi) ListSeaweedFSAccounts(ctx _context.Context, orgName string, clusterName string) ([]AccountListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []AccountListItem
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listSeaweedFSAccounts",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListSeaweedFSAccounts")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// ListSeaweedFSPolicies listSeaweedFSPolicies.
func (a *AccountApi) ListSeaweedFSPolicies(ctx _context.Context, orgName string, clusterName string) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listSeaweedFSPolicies",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/policies",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListSeaweedFSPolicies")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/policies"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// ListServiceAccountsOptionalParameters holds optional parameters for ListServiceAccounts.
type ListServiceAccountsOptionalParameters struct {
	UserName *string
}

// NewListServiceAccountsOptionalParameters creates an empty struct for parameters.
func NewListServiceAccountsOptionalParameters() *ListServiceAccountsOptionalParameters {
	this := ListServiceAccountsOptionalParameters{}
	return &this
}

// WithUserName sets the corresponding parameter name and returns the struct.
func (r *ListServiceAccountsOptionalParameters) WithUserName(userName string) *ListServiceAccountsOptionalParameters {
	r.UserName = &userName
	return r
}

// ListServiceAccounts List MinIO access keys.
// list MinIO service account access keys
func (a *AccountApi) ListServiceAccounts(ctx _context.Context, engineName string, orgName string, clusterName string, o ...ListServiceAccountsOptionalParameters) ([]DmsServiceAccount, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []DmsServiceAccount
		optionalParams      ListServiceAccountsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListServiceAccountsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "listServiceAccounts",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ListServiceAccounts")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.UserName != nil {
		localVarQueryParams.Add("userName", common.ParameterToString(*optionalParams.UserName, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
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

// LockAccount Lock cluster account.
// lock an account in cluster
func (a *AccountApi) LockAccount(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "lockAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/lock",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.LockAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/lock"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// PreviewAccountPrivilegeResetOptionalParameters holds optional parameters for PreviewAccountPrivilegeReset.
type PreviewAccountPrivilegeResetOptionalParameters struct {
	SubjectType       *ObjectPrivilegeSubjectType
	Host              *string
	SubjectQualifiers *string
}

// NewPreviewAccountPrivilegeResetOptionalParameters creates an empty struct for parameters.
func NewPreviewAccountPrivilegeResetOptionalParameters() *PreviewAccountPrivilegeResetOptionalParameters {
	this := PreviewAccountPrivilegeResetOptionalParameters{}
	return &this
}

// WithSubjectType sets the corresponding parameter name and returns the struct.
func (r *PreviewAccountPrivilegeResetOptionalParameters) WithSubjectType(subjectType ObjectPrivilegeSubjectType) *PreviewAccountPrivilegeResetOptionalParameters {
	r.SubjectType = &subjectType
	return r
}

// WithHost sets the corresponding parameter name and returns the struct.
func (r *PreviewAccountPrivilegeResetOptionalParameters) WithHost(host string) *PreviewAccountPrivilegeResetOptionalParameters {
	r.Host = &host
	return r
}

// WithSubjectQualifiers sets the corresponding parameter name and returns the struct.
func (r *PreviewAccountPrivilegeResetOptionalParameters) WithSubjectQualifiers(subjectQualifiers string) *PreviewAccountPrivilegeResetOptionalParameters {
	r.SubjectQualifiers = &subjectQualifiers
	return r
}

// PreviewAccountPrivilegeReset Preview clearing explicit account privileges.
// MySQL clears direct privileges at every scope (including columns, routines, dynamic privileges and grant options), explicit role memberships and PROXY grants for one exact user/host. Rejects protected accounts and servers with mandatory roles before writing. Preview does not mutate. Execution re-reads current grants; concurrent new grants can be included. Statements are not atomic; on failure refresh and preview before retrying. Existing sessions follow MySQL privilege-cache semantics; this does not lock the account or terminate connections. PostgreSQL clears direct ACLs, default ACLs, memberships and elevated role attributes across databases. PUBLIC access remains. Object ownership prevents reset; RESTRICT protects dependent grants. Each database is transactional, but cross-database reset is not atomic. No objects are dropped or reassigned.
func (a *AccountApi) PreviewAccountPrivilegeReset(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, o ...PreviewAccountPrivilegeResetOptionalParameters) ([]ObjectPrivilegeResetStep, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []ObjectPrivilegeResetStep
		optionalParams      PreviewAccountPrivilegeResetOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type PreviewAccountPrivilegeResetOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "previewAccountPrivilegeReset",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/reset",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.PreviewAccountPrivilegeReset")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/reset"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.SubjectType != nil {
		localVarQueryParams.Add("subjectType", common.ParameterToString(*optionalParams.SubjectType, ""))
	}
	if optionalParams.Host != nil {
		localVarQueryParams.Add("host", common.ParameterToString(*optionalParams.Host, ""))
	}
	if optionalParams.SubjectQualifiers != nil {
		localVarQueryParams.Add("subjectQualifiers", common.ParameterToString(*optionalParams.SubjectQualifiers, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// PutSeaweedFSPolicy putSeaweedFSPolicy.
// Creates or replaces a named AWS IAM JSON policy document. Updates affect every identity and group that uses this policy.
func (a *AccountApi) PutSeaweedFSPolicy(ctx _context.Context, orgName string, clusterName string, policyName string, body map[string]interface{}) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPut
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "putSeaweedFSPolicy",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/policies/{policyName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.PutSeaweedFSPolicy")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/policies/{policyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyName"+"}", _neturl.PathEscape(common.ParameterToString(policyName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "*/*"

	// body params
	localVarPostBody = &body
	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// ResetAccountPrivilegesOptionalParameters holds optional parameters for ResetAccountPrivileges.
type ResetAccountPrivilegesOptionalParameters struct {
	SubjectType       *ObjectPrivilegeSubjectType
	Host              *string
	SubjectQualifiers *string
}

// NewResetAccountPrivilegesOptionalParameters creates an empty struct for parameters.
func NewResetAccountPrivilegesOptionalParameters() *ResetAccountPrivilegesOptionalParameters {
	this := ResetAccountPrivilegesOptionalParameters{}
	return &this
}

// WithSubjectType sets the corresponding parameter name and returns the struct.
func (r *ResetAccountPrivilegesOptionalParameters) WithSubjectType(subjectType ObjectPrivilegeSubjectType) *ResetAccountPrivilegesOptionalParameters {
	r.SubjectType = &subjectType
	return r
}

// WithHost sets the corresponding parameter name and returns the struct.
func (r *ResetAccountPrivilegesOptionalParameters) WithHost(host string) *ResetAccountPrivilegesOptionalParameters {
	r.Host = &host
	return r
}

// WithSubjectQualifiers sets the corresponding parameter name and returns the struct.
func (r *ResetAccountPrivilegesOptionalParameters) WithSubjectQualifiers(subjectQualifiers string) *ResetAccountPrivilegesOptionalParameters {
	r.SubjectQualifiers = &subjectQualifiers
	return r
}

// ResetAccountPrivileges Clear explicit account privileges.
// MySQL clears direct privileges at every scope (including columns, routines, dynamic privileges and grant options), explicit role memberships and PROXY grants for one exact user/host. Rejects protected accounts and servers with mandatory roles before writing. Preview does not mutate. Execution re-reads current grants; concurrent new grants can be included. Statements are not atomic; on failure refresh and preview before retrying. Existing sessions follow MySQL privilege-cache semantics; this does not lock the account or terminate connections. PostgreSQL clears direct ACLs, default ACLs, memberships and elevated role attributes across databases. PUBLIC access remains. Object ownership prevents reset; RESTRICT protects dependent grants. Each database is transactional, but cross-database reset is not atomic. No objects are dropped or reassigned.
func (a *AccountApi) ResetAccountPrivileges(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, o ...ResetAccountPrivilegesOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
		optionalParams     ResetAccountPrivilegesOptionalParameters
	)

	if len(o) > 1 {
		return nil, common.ReportError("only one argument of type ResetAccountPrivilegesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "resetAccountPrivileges",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/reset",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.ResetAccountPrivileges")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/object-privileges/reset"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.SubjectType != nil {
		localVarQueryParams.Add("subjectType", common.ParameterToString(*optionalParams.SubjectType, ""))
	}
	if optionalParams.Host != nil {
		localVarQueryParams.Add("host", common.ParameterToString(*optionalParams.Host, ""))
	}
	if optionalParams.SubjectQualifiers != nil {
		localVarQueryParams.Add("subjectQualifiers", common.ParameterToString(*optionalParams.SubjectQualifiers, ""))
	}
	localVarHeaderParams["Accept"] = "*/*"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// UnlockAccount Unlock cluster account.
// unlock an account in cluster
func (a *AccountApi) UnlockAccount(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "unlockAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/unlock",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UnlockAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/unlock"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"DigestAuth", "Authorization"},
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

// UpdateAccount update cluster account.
// update an account in cluster
func (a *AccountApi) UpdateAccount(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, body Account) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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

// UpdateAccountPrivileges update account privileges.
// update account privileges for rdbms engine
func (a *AccountApi) UpdateAccountPrivileges(ctx _context.Context, engineName string, orgName string, clusterName string, accountName string, body []PrivilegeListItem) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateAccountPrivileges",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/privileges",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateAccountPrivileges")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/privileges"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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

// UpdateHiveAccount update hive account.
// update an hive account in cluster
func (a *AccountApi) UpdateHiveAccount(ctx _context.Context, orgName string, clusterName string, accountName string, body Account) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateHiveAccount",
		Path:        "/api/v1/data/hive/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateHiveAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/hive/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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

// UpdateKafkaAccount update cluster account.
// update an account in cluster
func (a *AccountApi) UpdateKafkaAccount(ctx _context.Context, orgName string, clusterName string, accountName string, body Account) (Account, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue Account
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateKafkaAccount",
		Path:        "/api/v1/data/kafka/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateKafkaAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/kafka/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 || localVarHTTPResponse.StatusCode == 501 {
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

// UpdateMilvusAccountPassword Update Milvus account password.
// Update a Milvus account password.
func (a *AccountApi) UpdateMilvusAccountPassword(ctx _context.Context, orgName string, clusterName string, accountName string, body MilvusAccountPassword) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateMilvusAccountPassword",
		Path:        "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateMilvusAccountPassword")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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

// UpdateMilvusAccountRoles Update Milvus account roles.
// Replace a Milvus account's role assignments.
func (a *AccountApi) UpdateMilvusAccountRoles(ctx _context.Context, orgName string, clusterName string, accountName string, body MilvusAccountRoles) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPut
		localVarPostBody   interface{}
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateMilvusAccountRoles",
		Path:        "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/roles",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateMilvusAccountRoles")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 500 {
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

// UpdateMilvusRole Update Milvus role.
// Update an existing Milvus native role by replacing its privilege grants.
func (a *AccountApi) UpdateMilvusRole(ctx _context.Context, orgName string, clusterName string, roleName string, body MilvusRoleRequest) (MilvusRole, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue MilvusRole
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateMilvusRole",
		Path:        "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles/{roleName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateMilvusRole")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/milvus/organizations/{orgName}/clusters/{clusterName}/accounts/roles/{roleName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleName"+"}", _neturl.PathEscape(common.ParameterToString(roleName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 500 {
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

// UpdateMssqlAccountOptionalParameters holds optional parameters for UpdateMssqlAccount.
type UpdateMssqlAccountOptionalParameters struct {
	Body *Account
}

// NewUpdateMssqlAccountOptionalParameters creates an empty struct for parameters.
func NewUpdateMssqlAccountOptionalParameters() *UpdateMssqlAccountOptionalParameters {
	this := UpdateMssqlAccountOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *UpdateMssqlAccountOptionalParameters) WithBody(body Account) *UpdateMssqlAccountOptionalParameters {
	r.Body = &body
	return r
}

// UpdateMssqlAccount update mssql account compatible with windows account.
func (a *AccountApi) UpdateMssqlAccount(ctx _context.Context, orgName string, clusterName string, o ...UpdateMssqlAccountOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
		optionalParams     UpdateMssqlAccountOptionalParameters
	)

	if len(o) > 1 {
		return nil, common.ReportError("only one argument of type UpdateMssqlAccountOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateMssqlAccount",
		Path:        "/api/v1/data/mssql/organizations/{orgName}/clusters/{clusterName}/accounts",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateMssqlAccount")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/mssql/organizations/{orgName}/clusters/{clusterName}/accounts"
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
		[2]string{"DigestAuth", "Authorization"},
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

// UpdateSeaweedFSAccount updateSeaweedFSAccount.
// Resets the secret for extra.accessKey. Account.name must match accountName. extra.accessKey is required. The returned password is disclosed only in this response.
func (a *AccountApi) UpdateSeaweedFSAccount(ctx _context.Context, orgName string, clusterName string, accountName string, body Account) (Account, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue Account
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateSeaweedFSAccount",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateSeaweedFSAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 503 {
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

// UpdateSeaweedFSAccountPrivileges updateSeaweedFSAccountPrivileges.
// Unsupported for SeaweedFS. Use the named policy binding API instead.
func (a *AccountApi) UpdateSeaweedFSAccountPrivileges(ctx _context.Context, orgName string, clusterName string, accountName string, body []PrivilegeListItem) (APIErrorResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue APIErrorResponse
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateSeaweedFSAccountPrivileges",
		Path:        "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/privileges",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateSeaweedFSAccountPrivileges")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/seaweedfs/organizations/{orgName}/clusters/{clusterName}/accounts/{accountName}/privileges"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", _neturl.PathEscape(common.ParameterToString(accountName, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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
		if localVarHTTPResponse.StatusCode == 400 {
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

// UpdateServiceAccount Update MinIO access key.
// update a MinIO service account access key
func (a *AccountApi) UpdateServiceAccount(ctx _context.Context, engineName string, orgName string, clusterName string, accessKey string, body DmsServiceAccountUpdateRequest) (DmsServiceAccount, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue DmsServiceAccount
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "account",
		OperationID: "updateServiceAccount",
		Path:        "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys/{accessKey}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".AccountApi.UpdateServiceAccount")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/data/{engineName}/organizations/{orgName}/clusters/{clusterName}/access-keys/{accessKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"engineName"+"}", _neturl.PathEscape(common.ParameterToString(engineName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accessKey"+"}", _neturl.PathEscape(common.ParameterToString(accessKey, "")), -1)

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
		[2]string{"DigestAuth", "Authorization"},
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

// NewAccountApi Returns NewAccountApi.
func NewAccountApi(client *common.APIClient) *AccountApi {
	return &AccountApi{
		Client: client,
	}
}
