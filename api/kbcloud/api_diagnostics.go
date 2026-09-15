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

// DiagnosticsApi service type
type DiagnosticsApi common.Service

// ExplainDiagnosticsElasticsearchAllocationOptionalParameters holds optional parameters for ExplainDiagnosticsElasticsearchAllocation.
type ExplainDiagnosticsElasticsearchAllocationOptionalParameters struct {
	CurrentNode *string
}

// NewExplainDiagnosticsElasticsearchAllocationOptionalParameters creates an empty struct for parameters.
func NewExplainDiagnosticsElasticsearchAllocationOptionalParameters() *ExplainDiagnosticsElasticsearchAllocationOptionalParameters {
	this := ExplainDiagnosticsElasticsearchAllocationOptionalParameters{}
	return &this
}

// WithCurrentNode sets the corresponding parameter name and returns the struct.
func (r *ExplainDiagnosticsElasticsearchAllocationOptionalParameters) WithCurrentNode(currentNode string) *ExplainDiagnosticsElasticsearchAllocationOptionalParameters {
	r.CurrentNode = &currentNode
	return r
}

// ExplainDiagnosticsElasticsearchAllocation Explain Elasticsearch shard allocation.
func (a *DiagnosticsApi) ExplainDiagnosticsElasticsearchAllocation(ctx _context.Context, orgName string, clusterName string, index string, shard int64, primary bool, o ...ExplainDiagnosticsElasticsearchAllocationOptionalParameters) (ElasticsearchAllocationExplanation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ElasticsearchAllocationExplanation
		optionalParams      ExplainDiagnosticsElasticsearchAllocationOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ExplainDiagnosticsElasticsearchAllocationOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "explainDiagnosticsElasticsearchAllocation",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/allocationExplain",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ExplainDiagnosticsElasticsearchAllocation")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/allocationExplain"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if common.Strlen(index) < 1 {
		return localVarReturnValue, nil, common.ReportError("index must have at least 1 elements")
	}
	if shard < 0 {
		return localVarReturnValue, nil, common.ReportError("shard must be greater than 0")
	}
	localVarQueryParams.Add("index", common.ParameterToString(index, ""))
	localVarQueryParams.Add("shard", common.ParameterToString(shard, ""))
	localVarQueryParams.Add("primary", common.ParameterToString(primary, ""))
	if optionalParams.CurrentNode != nil {
		localVarQueryParams.Add("currentNode", common.ParameterToString(*optionalParams.CurrentNode, ""))
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

// ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters holds optional parameters for ExplainDiagnosticsPostgresqlSQLFingerprint.
type ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters struct {
	Fingerprint *string
}

// NewExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters creates an empty struct for parameters.
func NewExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters() *ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters {
	this := ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters{}
	return &this
}

// WithFingerprint sets the corresponding parameter name and returns the struct.
func (r *ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters) WithFingerprint(fingerprint string) *ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters {
	r.Fingerprint = &fingerprint
	return r
}

// ExplainDiagnosticsPostgresqlSQLFingerprint Explain PostgreSQL SQL fingerprint.
// Explicitly trigger a safe PostgreSQL EXPLAIN for one SQL fingerprint. The request does not accept raw SQL. DMS resolves the exact server-side parameterized SELECT statement identity independently of the ranking window, produces an estimated standard or generic plan without parameter sample values, rejects multiple statements, and never runs EXPLAIN ANALYZE or the original SQL. The plan uses current catalog statistics and the DMS connection context; it does not reconstruct the original role, search_path, session settings, parameter values, or historical actual plan.
func (a *DiagnosticsApi) ExplainDiagnosticsPostgresqlSQLFingerprint(ctx _context.Context, orgName string, clusterName string, queryId string, database string, user string, topLevel bool, o ...ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters) (PostgresqlSQLFingerprintExplainResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue PostgresqlSQLFingerprintExplainResponse
		optionalParams      ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "explainDiagnosticsPostgresqlSQLFingerprint",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis/queries/{queryID}/explain",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ExplainDiagnosticsPostgresqlSQLFingerprint")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis/queries/{queryID}/explain"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queryID"+"}", _neturl.PathEscape(common.ParameterToString(queryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("database", common.ParameterToString(database, ""))
	localVarQueryParams.Add("user", common.ParameterToString(user, ""))
	localVarQueryParams.Add("topLevel", common.ParameterToString(topLevel, ""))
	if optionalParams.Fingerprint != nil {
		localVarQueryParams.Add("fingerprint", common.ParameterToString(*optionalParams.Fingerprint, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 422 || localVarHTTPResponse.StatusCode == 500 || localVarHTTPResponse.StatusCode == 503 {
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

// GetDiagnosticsDamengSQLAnalysisOptionalParameters holds optional parameters for GetDiagnosticsDamengSQLAnalysis.
type GetDiagnosticsDamengSQLAnalysisOptionalParameters struct {
	Limit *int64
}

// NewGetDiagnosticsDamengSQLAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsDamengSQLAnalysisOptionalParameters() *GetDiagnosticsDamengSQLAnalysisOptionalParameters {
	this := GetDiagnosticsDamengSQLAnalysisOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsDamengSQLAnalysisOptionalParameters) WithLimit(limit int64) *GetDiagnosticsDamengSQLAnalysisOptionalParameters {
	r.Limit = &limit
	return r
}

// GetDiagnosticsDamengSQLAnalysis Get Dameng SQL analysis.
// Get a read-only Dameng SQL analysis snapshot from V$SYSTEM_LONG_EXEC_SQLS (long-running SQL) and V$SYSTEM_LARGE_MEM_SQLS (high-memory SQL). The response does not expose execution plans or remediation actions.
func (a *DiagnosticsApi) GetDiagnosticsDamengSQLAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsDamengSQLAnalysisOptionalParameters) (DamengSQLAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DamengSQLAnalysis
		optionalParams      GetDiagnosticsDamengSQLAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsDamengSQLAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsDamengSQLAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sqlAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsDamengSQLAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sqlAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
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

// GetDiagnosticsDamengSession Get Dameng session detail.
// Get a single Dameng session detail by session ID from V$SESSIONS, including SQL text, transaction info, and client details.
func (a *DiagnosticsApi) GetDiagnosticsDamengSession(ctx _context.Context, orgName string, clusterName string, sessionId int64) (DamengSessionDetail, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DamengSessionDetail
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsDamengSession",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions/{sessionId}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsDamengSession")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions/{sessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", _neturl.PathEscape(common.ParameterToString(sessionId, "")), -1)

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

// GetDiagnosticsDamengSessionLockAnalysis Get Dameng session lock analysis.
// Get lock and blocking analysis for a Dameng session. Queries V$LOCK and V$SESSIONS to identify blocking/blocked relationships and lock details.
func (a *DiagnosticsApi) GetDiagnosticsDamengSessionLockAnalysis(ctx _context.Context, orgName string, clusterName string, sessionId int64) (DamengLockAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DamengLockAnalysis
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsDamengSessionLockAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions/{sessionId}/lockAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsDamengSessionLockAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions/{sessionId}/lockAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", _neturl.PathEscape(common.ParameterToString(sessionId, "")), -1)

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

// GetDiagnosticsDamengSpaceAnalysisOptionalParameters holds optional parameters for GetDiagnosticsDamengSpaceAnalysis.
type GetDiagnosticsDamengSpaceAnalysisOptionalParameters struct {
	TableLimit *int64
	IndexLimit *int64
	Schema     *string
	SkipBasic  *bool
}

// NewGetDiagnosticsDamengSpaceAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsDamengSpaceAnalysisOptionalParameters() *GetDiagnosticsDamengSpaceAnalysisOptionalParameters {
	this := GetDiagnosticsDamengSpaceAnalysisOptionalParameters{}
	return &this
}

// WithTableLimit sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsDamengSpaceAnalysisOptionalParameters) WithTableLimit(tableLimit int64) *GetDiagnosticsDamengSpaceAnalysisOptionalParameters {
	r.TableLimit = &tableLimit
	return r
}

// WithIndexLimit sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsDamengSpaceAnalysisOptionalParameters) WithIndexLimit(indexLimit int64) *GetDiagnosticsDamengSpaceAnalysisOptionalParameters {
	r.IndexLimit = &indexLimit
	return r
}

// WithSchema sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsDamengSpaceAnalysisOptionalParameters) WithSchema(schema string) *GetDiagnosticsDamengSpaceAnalysisOptionalParameters {
	r.Schema = &schema
	return r
}

// WithSkipBasic sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsDamengSpaceAnalysisOptionalParameters) WithSkipBasic(skipBasic bool) *GetDiagnosticsDamengSpaceAnalysisOptionalParameters {
	r.SkipBasic = &skipBasic
	return r
}

// GetDiagnosticsDamengSpaceAnalysis Get Dameng space analysis.
// Get a read-only Dameng space snapshot including tablespace usage, schema sizes, top tables, and top indexes. Queries dba_free_space, dba_data_files, DBA_SEGMENTS, DBA_TABLES, DBA_INDEXES, and DBA_IND_COLUMNS.
func (a *DiagnosticsApi) GetDiagnosticsDamengSpaceAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsDamengSpaceAnalysisOptionalParameters) (DamengSpaceAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DamengSpaceAnalysis
		optionalParams      GetDiagnosticsDamengSpaceAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsDamengSpaceAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsDamengSpaceAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/spaceAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsDamengSpaceAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/spaceAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.TableLimit != nil {
		localVarQueryParams.Add("tableLimit", common.ParameterToString(*optionalParams.TableLimit, ""))
	}
	if optionalParams.IndexLimit != nil {
		localVarQueryParams.Add("indexLimit", common.ParameterToString(*optionalParams.IndexLimit, ""))
	}
	if optionalParams.Schema != nil {
		localVarQueryParams.Add("schema", common.ParameterToString(*optionalParams.Schema, ""))
	}
	if optionalParams.SkipBasic != nil {
		localVarQueryParams.Add("skipBasic", common.ParameterToString(*optionalParams.SkipBasic, ""))
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

// GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters holds optional parameters for GetDiagnosticsElasticsearchStorageAnalysis.
type GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters struct {
	Node      *string
	Index     *string
	SortBy    *ElasticsearchIndexSortBy
	SortOrder *ElasticsearchSortOrder
}

// NewGetDiagnosticsElasticsearchStorageAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsElasticsearchStorageAnalysisOptionalParameters() *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters {
	this := GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters{}
	return &this
}

// WithNode sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters) WithNode(node string) *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters {
	r.Node = &node
	return r
}

// WithIndex sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters) WithIndex(index string) *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters {
	r.Index = &index
	return r
}

// WithSortBy sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters) WithSortBy(sortBy ElasticsearchIndexSortBy) *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters {
	r.SortBy = &sortBy
	return r
}

// WithSortOrder sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters) WithSortOrder(sortOrder ElasticsearchSortOrder) *GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters {
	r.SortOrder = &sortOrder
	return r
}

// GetDiagnosticsElasticsearchStorageAnalysis Get Elasticsearch storage and index analysis.
func (a *DiagnosticsApi) GetDiagnosticsElasticsearchStorageAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters) (ElasticsearchStorageAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ElasticsearchStorageAnalysis
		optionalParams      GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsElasticsearchStorageAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsElasticsearchStorageAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/storageAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsElasticsearchStorageAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/storageAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Node != nil {
		localVarQueryParams.Add("node", common.ParameterToString(*optionalParams.Node, ""))
	}
	if optionalParams.Index != nil {
		localVarQueryParams.Add("index", common.ParameterToString(*optionalParams.Index, ""))
	}
	if optionalParams.SortBy != nil {
		localVarQueryParams.Add("sortBy", common.ParameterToString(*optionalParams.SortBy, ""))
	}
	if optionalParams.SortOrder != nil {
		localVarQueryParams.Add("sortOrder", common.ParameterToString(*optionalParams.SortOrder, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 500 {
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

// GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters holds optional parameters for GetDiagnosticsElasticsearchTaskAnalysis.
type GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters struct {
	Type      *ElasticsearchHotThreadsType
	Node      *string
	TaskLimit *int32
	Threads   *int32
}

// NewGetDiagnosticsElasticsearchTaskAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsElasticsearchTaskAnalysisOptionalParameters() *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters {
	this := GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters{}
	return &this
}

// WithType sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters) WithType(typeVar ElasticsearchHotThreadsType) *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters {
	r.Type = &typeVar
	return r
}

// WithNode sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters) WithNode(node string) *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters {
	r.Node = &node
	return r
}

// WithTaskLimit sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters) WithTaskLimit(taskLimit int32) *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters {
	r.TaskLimit = &taskLimit
	return r
}

// WithThreads sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters) WithThreads(threads int32) *GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters {
	r.Threads = &threads
	return r
}

// GetDiagnosticsElasticsearchTaskAnalysis Get Elasticsearch task and hot-thread analysis.
func (a *DiagnosticsApi) GetDiagnosticsElasticsearchTaskAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters) (ElasticsearchTaskAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ElasticsearchTaskAnalysis
		optionalParams      GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsElasticsearchTaskAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsElasticsearchTaskAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/taskAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsElasticsearchTaskAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/taskAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Type != nil {
		localVarQueryParams.Add("type", common.ParameterToString(*optionalParams.Type, ""))
	}
	if optionalParams.Node != nil {
		localVarQueryParams.Add("node", common.ParameterToString(*optionalParams.Node, ""))
	}
	if optionalParams.TaskLimit != nil {
		localVarQueryParams.Add("taskLimit", common.ParameterToString(*optionalParams.TaskLimit, ""))
	}
	if optionalParams.Threads != nil {
		localVarQueryParams.Add("threads", common.ParameterToString(*optionalParams.Threads, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 500 {
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

// GetDiagnosticsKingbaseSession Get Kingbase session diagnostics.
func (a *DiagnosticsApi) GetDiagnosticsKingbaseSession(ctx _context.Context, orgName string, clusterName string, pid int64) (KingbaseSessionDetail, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue KingbaseSessionDetail
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsKingbaseSession",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/kingbase/sessions/{pid}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsKingbaseSession")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/kingbase/sessions/{pid}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pid"+"}", _neturl.PathEscape(common.ParameterToString(pid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if pid < 1 {
		return localVarReturnValue, nil, common.ReportError("pid must be greater than 1")
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

// GetDiagnosticsKingbaseSessionLockAnalysis Get Kingbase session lock analysis.
// Returns the selected session and its direct lock rows from the current snapshot.
func (a *DiagnosticsApi) GetDiagnosticsKingbaseSessionLockAnalysis(ctx _context.Context, orgName string, clusterName string, pid int64) (KingbaseLockAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue KingbaseLockAnalysis
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsKingbaseSessionLockAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/kingbase/sessions/{pid}/lockAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsKingbaseSessionLockAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/kingbase/sessions/{pid}/lockAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pid"+"}", _neturl.PathEscape(common.ParameterToString(pid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if pid < 1 {
		return localVarReturnValue, nil, common.ReportError("pid must be greater than 1")
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

// GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters holds optional parameters for GetDiagnosticsMariaDBPerformanceTrends.
type GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters struct {
	Range *string
	Step  *string
}

// NewGetDiagnosticsMariaDBPerformanceTrendsOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsMariaDBPerformanceTrendsOptionalParameters() *GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters {
	this := GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters{}
	return &this
}

// WithRange sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters) WithRange(rangeVar string) *GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters {
	r.Range = &rangeVar
	return r
}

// WithStep sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters) WithStep(step string) *GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters {
	r.Step = &step
	return r
}

// GetDiagnosticsMariaDBPerformanceTrends Get MariaDB performance trends.
// Get read-only MariaDB performance trends from backend-owned metrics and database capability checks. The response does not expose SQL, PromQL, internal endpoints, or credentials.
func (a *DiagnosticsApi) GetDiagnosticsMariaDBPerformanceTrends(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters) (PerformanceTrends, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PerformanceTrends
		optionalParams      GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsMariaDBPerformanceTrendsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMariaDBPerformanceTrends",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mariadb/performanceTrends",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMariaDBPerformanceTrends")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mariadb/performanceTrends"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Range != nil {
		localVarQueryParams.Add("range", common.ParameterToString(*optionalParams.Range, ""))
	}
	if optionalParams.Step != nil {
		localVarQueryParams.Add("step", common.ParameterToString(*optionalParams.Step, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 500 {
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

// GetDiagnosticsMariaDBSQLAnalysisOptionalParameters holds optional parameters for GetDiagnosticsMariaDBSQLAnalysis.
type GetDiagnosticsMariaDBSQLAnalysisOptionalParameters struct {
	Limit   *int64
	OrderBy *string
}

// NewGetDiagnosticsMariaDBSQLAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsMariaDBSQLAnalysisOptionalParameters() *GetDiagnosticsMariaDBSQLAnalysisOptionalParameters {
	this := GetDiagnosticsMariaDBSQLAnalysisOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMariaDBSQLAnalysisOptionalParameters) WithLimit(limit int64) *GetDiagnosticsMariaDBSQLAnalysisOptionalParameters {
	r.Limit = &limit
	return r
}

// WithOrderBy sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMariaDBSQLAnalysisOptionalParameters) WithOrderBy(orderBy string) *GetDiagnosticsMariaDBSQLAnalysisOptionalParameters {
	r.OrderBy = &orderBy
	return r
}

// GetDiagnosticsMariaDBSQLAnalysis Get MariaDB SQL analysis.
// Get a read-only MariaDB SQL fingerprint ranking from Performance Schema statement digests. The response does not expose raw SQL text, time-window aggregation, execution plans, or remediation actions.
func (a *DiagnosticsApi) GetDiagnosticsMariaDBSQLAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsMariaDBSQLAnalysisOptionalParameters) (MysqlSQLAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MysqlSQLAnalysis
		optionalParams      GetDiagnosticsMariaDBSQLAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsMariaDBSQLAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMariaDBSQLAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mariadb/sqlAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMariaDBSQLAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mariadb/sqlAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.OrderBy != nil {
		localVarQueryParams.Add("orderBy", common.ParameterToString(*optionalParams.OrderBy, ""))
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

// GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters holds optional parameters for GetDiagnosticsMariaDBSpaceAnalysis.
type GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters struct {
	DatabaseName *string
}

// NewGetDiagnosticsMariaDBSpaceAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsMariaDBSpaceAnalysisOptionalParameters() *GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters {
	this := GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters{}
	return &this
}

// WithDatabaseName sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters) WithDatabaseName(databaseName string) *GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters {
	r.DatabaseName = &databaseName
	return r
}

// GetDiagnosticsMariaDBSpaceAnalysis Get MariaDB space analysis.
// Get a read-only MariaDB compatible space snapshot and fixed backend-owned storage metrics.
func (a *DiagnosticsApi) GetDiagnosticsMariaDBSpaceAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters) (MysqlSpaceAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MysqlSpaceAnalysis
		optionalParams      GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsMariaDBSpaceAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMariaDBSpaceAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mariadb/spaceAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMariaDBSpaceAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mariadb/spaceAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// GetDiagnosticsMssqlSession Get a SQL Server session.
func (a *DiagnosticsApi) GetDiagnosticsMssqlSession(ctx _context.Context, orgName string, clusterName string, sessionId int64, startedAt string) (MssqlSession, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MssqlSession
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMssqlSession",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mssql/sessions/{sessionId}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMssqlSession")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mssql/sessions/{sessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", _neturl.PathEscape(common.ParameterToString(sessionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if sessionId < 1 {
		return localVarReturnValue, nil, common.ReportError("sessionId must be greater than 1")
	}
	if sessionId > 32767 {
		return localVarReturnValue, nil, common.ReportError("sessionId must be less than 32767")
	}
	localVarQueryParams.Add("startedAt", common.ParameterToString(startedAt, ""))
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

// GetDiagnosticsMssqlSessionCachedPlan Read the cached compiled plan of a live SQL Server request.
func (a *DiagnosticsApi) GetDiagnosticsMssqlSessionCachedPlan(ctx _context.Context, orgName string, clusterName string, sessionId int64, startedAt string, requestId int64, requestStart string) (MssqlCachedPlan, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MssqlCachedPlan
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMssqlSessionCachedPlan",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mssql/sessions/{sessionId}/requests/{requestId}/cachedPlan",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMssqlSessionCachedPlan")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mssql/sessions/{sessionId}/requests/{requestId}/cachedPlan"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", _neturl.PathEscape(common.ParameterToString(sessionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", _neturl.PathEscape(common.ParameterToString(requestId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if sessionId < 1 {
		return localVarReturnValue, nil, common.ReportError("sessionId must be greater than 1")
	}
	if sessionId > 32767 {
		return localVarReturnValue, nil, common.ReportError("sessionId must be less than 32767")
	}
	if requestId < 0 {
		return localVarReturnValue, nil, common.ReportError("requestId must be greater than 0")
	}
	localVarQueryParams.Add("startedAt", common.ParameterToString(startedAt, ""))
	localVarQueryParams.Add("requestStart", common.ParameterToString(requestStart, ""))
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

// GetDiagnosticsMssqlSessionLockAnalysis Get SQL Server blocking and locks.
func (a *DiagnosticsApi) GetDiagnosticsMssqlSessionLockAnalysis(ctx _context.Context, orgName string, clusterName string, sessionId int64, startedAt string) (MssqlLockSnapshot, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MssqlLockSnapshot
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMssqlSessionLockAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mssql/sessions/{sessionId}/lockAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMssqlSessionLockAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mssql/sessions/{sessionId}/lockAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", _neturl.PathEscape(common.ParameterToString(sessionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if sessionId < 1 {
		return localVarReturnValue, nil, common.ReportError("sessionId must be greater than 1")
	}
	if sessionId > 32767 {
		return localVarReturnValue, nil, common.ReportError("sessionId must be less than 32767")
	}
	localVarQueryParams.Add("startedAt", common.ParameterToString(startedAt, ""))
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

// GetDiagnosticsMysqlPerformanceTrendsOptionalParameters holds optional parameters for GetDiagnosticsMysqlPerformanceTrends.
type GetDiagnosticsMysqlPerformanceTrendsOptionalParameters struct {
	Range *string
	Step  *string
}

// NewGetDiagnosticsMysqlPerformanceTrendsOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsMysqlPerformanceTrendsOptionalParameters() *GetDiagnosticsMysqlPerformanceTrendsOptionalParameters {
	this := GetDiagnosticsMysqlPerformanceTrendsOptionalParameters{}
	return &this
}

// WithRange sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMysqlPerformanceTrendsOptionalParameters) WithRange(rangeVar string) *GetDiagnosticsMysqlPerformanceTrendsOptionalParameters {
	r.Range = &rangeVar
	return r
}

// WithStep sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMysqlPerformanceTrendsOptionalParameters) WithStep(step string) *GetDiagnosticsMysqlPerformanceTrendsOptionalParameters {
	r.Step = &step
	return r
}

// GetDiagnosticsMysqlPerformanceTrends Get MySQL performance trends.
// Get read-only MySQL 5.7, 8.0, and 8.4 performance trends from backend-owned metrics and database capability checks. The response does not expose SQL, PromQL, internal endpoints, or credentials.
func (a *DiagnosticsApi) GetDiagnosticsMysqlPerformanceTrends(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsMysqlPerformanceTrendsOptionalParameters) (PerformanceTrends, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PerformanceTrends
		optionalParams      GetDiagnosticsMysqlPerformanceTrendsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsMysqlPerformanceTrendsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMysqlPerformanceTrends",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mysql/performanceTrends",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMysqlPerformanceTrends")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mysql/performanceTrends"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Range != nil {
		localVarQueryParams.Add("range", common.ParameterToString(*optionalParams.Range, ""))
	}
	if optionalParams.Step != nil {
		localVarQueryParams.Add("step", common.ParameterToString(*optionalParams.Step, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 500 {
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

// GetDiagnosticsMysqlSQLAnalysisOptionalParameters holds optional parameters for GetDiagnosticsMysqlSQLAnalysis.
type GetDiagnosticsMysqlSQLAnalysisOptionalParameters struct {
	Limit   *int64
	OrderBy *string
}

// NewGetDiagnosticsMysqlSQLAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsMysqlSQLAnalysisOptionalParameters() *GetDiagnosticsMysqlSQLAnalysisOptionalParameters {
	this := GetDiagnosticsMysqlSQLAnalysisOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMysqlSQLAnalysisOptionalParameters) WithLimit(limit int64) *GetDiagnosticsMysqlSQLAnalysisOptionalParameters {
	r.Limit = &limit
	return r
}

// WithOrderBy sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMysqlSQLAnalysisOptionalParameters) WithOrderBy(orderBy string) *GetDiagnosticsMysqlSQLAnalysisOptionalParameters {
	r.OrderBy = &orderBy
	return r
}

// GetDiagnosticsMysqlSQLAnalysis Get MySQL SQL analysis.
// Get a read-only MySQL SQL fingerprint ranking from Performance Schema statement digests. The response does not expose raw SQL text, time-window aggregation, execution plans, or remediation actions.
func (a *DiagnosticsApi) GetDiagnosticsMysqlSQLAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsMysqlSQLAnalysisOptionalParameters) (MysqlSQLAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MysqlSQLAnalysis
		optionalParams      GetDiagnosticsMysqlSQLAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsMysqlSQLAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMysqlSQLAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mysql/sqlAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMysqlSQLAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mysql/sqlAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.OrderBy != nil {
		localVarQueryParams.Add("orderBy", common.ParameterToString(*optionalParams.OrderBy, ""))
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

// GetDiagnosticsMysqlSpaceAnalysisOptionalParameters holds optional parameters for GetDiagnosticsMysqlSpaceAnalysis.
type GetDiagnosticsMysqlSpaceAnalysisOptionalParameters struct {
	DatabaseName *string
}

// NewGetDiagnosticsMysqlSpaceAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsMysqlSpaceAnalysisOptionalParameters() *GetDiagnosticsMysqlSpaceAnalysisOptionalParameters {
	this := GetDiagnosticsMysqlSpaceAnalysisOptionalParameters{}
	return &this
}

// WithDatabaseName sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsMysqlSpaceAnalysisOptionalParameters) WithDatabaseName(databaseName string) *GetDiagnosticsMysqlSpaceAnalysisOptionalParameters {
	r.DatabaseName = &databaseName
	return r
}

// GetDiagnosticsMysqlSpaceAnalysis Get MySQL space analysis.
// Get a read-only MySQL 5.7, 8.0, or 8.4 compatible space snapshot and fixed backend-owned storage metrics.
func (a *DiagnosticsApi) GetDiagnosticsMysqlSpaceAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsMysqlSpaceAnalysisOptionalParameters) (MysqlSpaceAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MysqlSpaceAnalysis
		optionalParams      GetDiagnosticsMysqlSpaceAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsMysqlSpaceAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsMysqlSpaceAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mysql/spaceAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsMysqlSpaceAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mysql/spaceAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters holds optional parameters for GetDiagnosticsPostgresqlPerformanceTrends.
type GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters struct {
	Range *string
	Step  *string
}

// NewGetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters() *GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters {
	this := GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters{}
	return &this
}

// WithRange sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters) WithRange(rangeVar string) *GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters {
	r.Range = &rangeVar
	return r
}

// WithStep sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters) WithStep(step string) *GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters {
	r.Step = &step
	return r
}

// GetDiagnosticsPostgresqlPerformanceTrends Get PostgreSQL performance trends.
// Get read-only PostgreSQL performance trends from backend-owned Prometheus queries. The response does not expose SQL, PromQL, internal endpoints, or remediation actions.
func (a *DiagnosticsApi) GetDiagnosticsPostgresqlPerformanceTrends(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters) (PerformanceTrends, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PerformanceTrends
		optionalParams      GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsPostgresqlPerformanceTrendsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsPostgresqlPerformanceTrends",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/performanceTrends",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlPerformanceTrends")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/performanceTrends"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Range != nil {
		localVarQueryParams.Add("range", common.ParameterToString(*optionalParams.Range, ""))
	}
	if optionalParams.Step != nil {
		localVarQueryParams.Add("step", common.ParameterToString(*optionalParams.Step, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 500 {
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

// GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters holds optional parameters for GetDiagnosticsPostgresqlSQLAnalysis.
type GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters struct {
	Limit   *int64
	OrderBy *string
}

// NewGetDiagnosticsPostgresqlSQLAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsPostgresqlSQLAnalysisOptionalParameters() *GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters {
	this := GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters) WithLimit(limit int64) *GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters {
	r.Limit = &limit
	return r
}

// WithOrderBy sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters) WithOrderBy(orderBy string) *GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters {
	r.OrderBy = &orderBy
	return r
}

// GetDiagnosticsPostgresqlSQLAnalysis Get PostgreSQL SQL analysis.
// Get a read-only PostgreSQL SQL fingerprint ranking from pg_stat_statements. The response does not expose full SQL text, time-window aggregation, summary cards, slow-log relation, execution plans, or remediation actions.
func (a *DiagnosticsApi) GetDiagnosticsPostgresqlSQLAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters) (PostgresqlSQLAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PostgresqlSQLAnalysis
		optionalParams      GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsPostgresqlSQLAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsPostgresqlSQLAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSQLAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.OrderBy != nil {
		localVarQueryParams.Add("orderBy", common.ParameterToString(*optionalParams.OrderBy, ""))
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

// GetDiagnosticsPostgresqlSession Get PostgreSQL session basic diagnostics.
// Get one PostgreSQL session basic diagnostics record by backend pid.
func (a *DiagnosticsApi) GetDiagnosticsPostgresqlSession(ctx _context.Context, orgName string, clusterName string, pid int64) (PostgresqlSession, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PostgresqlSession
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsPostgresqlSession",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions/{pid}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSession")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions/{pid}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pid"+"}", _neturl.PathEscape(common.ParameterToString(pid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if pid < 1 {
		return localVarReturnValue, nil, common.ReportError("pid must be greater than 1")
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

// GetDiagnosticsPostgresqlSessionLockAnalysis Get PostgreSQL session lock analysis.
// Get read-only lock analysis for one PostgreSQL backend pid from the current DMS PostgreSQL lock snapshot.
func (a *DiagnosticsApi) GetDiagnosticsPostgresqlSessionLockAnalysis(ctx _context.Context, orgName string, clusterName string, pid int64) (PostgresqlLockAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PostgresqlLockAnalysis
	)

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsPostgresqlSessionLockAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions/{pid}/lockAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSessionLockAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions/{pid}/lockAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pid"+"}", _neturl.PathEscape(common.ParameterToString(pid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if pid < 1 {
		return localVarReturnValue, nil, common.ReportError("pid must be greater than 1")
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

// GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters holds optional parameters for GetDiagnosticsPostgresqlSpaceAnalysis.
type GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters struct {
	DatabaseName *string
}

// NewGetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters() *GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters {
	this := GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters{}
	return &this
}

// WithDatabaseName sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters) WithDatabaseName(databaseName string) *GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters {
	r.DatabaseName = &databaseName
	return r
}

// GetDiagnosticsPostgresqlSpaceAnalysis Get PostgreSQL space analysis.
// Get a read-only PostgreSQL space snapshot from DMS and fixed backend-owned storage metrics. The response does not expose SQL, PromQL, storage history, or remediation actions.
func (a *DiagnosticsApi) GetDiagnosticsPostgresqlSpaceAnalysis(ctx _context.Context, orgName string, clusterName string, o ...GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters) (PostgresqlSpaceAnalysis, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PostgresqlSpaceAnalysis
		optionalParams      GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsPostgresqlSpaceAnalysisOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsPostgresqlSpaceAnalysis",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/spaceAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSpaceAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/spaceAnalysis"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// ListDiagnosticsDamengSessionsOptionalParameters holds optional parameters for ListDiagnosticsDamengSessions.
type ListDiagnosticsDamengSessionsOptionalParameters struct {
	Limit *int64
	State *string
}

// NewListDiagnosticsDamengSessionsOptionalParameters creates an empty struct for parameters.
func NewListDiagnosticsDamengSessionsOptionalParameters() *ListDiagnosticsDamengSessionsOptionalParameters {
	this := ListDiagnosticsDamengSessionsOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsDamengSessionsOptionalParameters) WithLimit(limit int64) *ListDiagnosticsDamengSessionsOptionalParameters {
	r.Limit = &limit
	return r
}

// WithState sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsDamengSessionsOptionalParameters) WithState(state string) *ListDiagnosticsDamengSessionsOptionalParameters {
	r.State = &state
	return r
}

// ListDiagnosticsDamengSessions List Dameng sessions.
// List Dameng sessions from V$SESSIONS with lock status from V$LOCK. Returns session ID, user, state, client IP, transaction ID, lock status, SQL text, schema, duration, and more.
func (a *DiagnosticsApi) ListDiagnosticsDamengSessions(ctx _context.Context, orgName string, clusterName string, o ...ListDiagnosticsDamengSessionsOptionalParameters) ([]DamengSessionListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []DamengSessionListItem
		optionalParams      ListDiagnosticsDamengSessionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListDiagnosticsDamengSessionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "listDiagnosticsDamengSessions",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ListDiagnosticsDamengSessions")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
	}
	if optionalParams.State != nil {
		localVarQueryParams.Add("state", common.ParameterToString(*optionalParams.State, ""))
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

// ListDiagnosticsElasticsearchRecoveriesOptionalParameters holds optional parameters for ListDiagnosticsElasticsearchRecoveries.
type ListDiagnosticsElasticsearchRecoveriesOptionalParameters struct {
	Index *string
	Node  *string
}

// NewListDiagnosticsElasticsearchRecoveriesOptionalParameters creates an empty struct for parameters.
func NewListDiagnosticsElasticsearchRecoveriesOptionalParameters() *ListDiagnosticsElasticsearchRecoveriesOptionalParameters {
	this := ListDiagnosticsElasticsearchRecoveriesOptionalParameters{}
	return &this
}

// WithIndex sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsElasticsearchRecoveriesOptionalParameters) WithIndex(index string) *ListDiagnosticsElasticsearchRecoveriesOptionalParameters {
	r.Index = &index
	return r
}

// WithNode sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsElasticsearchRecoveriesOptionalParameters) WithNode(node string) *ListDiagnosticsElasticsearchRecoveriesOptionalParameters {
	r.Node = &node
	return r
}

// ListDiagnosticsElasticsearchRecoveries List active Elasticsearch shard recoveries.
func (a *DiagnosticsApi) ListDiagnosticsElasticsearchRecoveries(ctx _context.Context, orgName string, clusterName string, o ...ListDiagnosticsElasticsearchRecoveriesOptionalParameters) (ElasticsearchRecoveryList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ElasticsearchRecoveryList
		optionalParams      ListDiagnosticsElasticsearchRecoveriesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListDiagnosticsElasticsearchRecoveriesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "listDiagnosticsElasticsearchRecoveries",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/recoveries",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ListDiagnosticsElasticsearchRecoveries")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/recoveries"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Index != nil {
		localVarQueryParams.Add("index", common.ParameterToString(*optionalParams.Index, ""))
	}
	if optionalParams.Node != nil {
		localVarQueryParams.Add("node", common.ParameterToString(*optionalParams.Node, ""))
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

// ListDiagnosticsElasticsearchShardsOptionalParameters holds optional parameters for ListDiagnosticsElasticsearchShards.
type ListDiagnosticsElasticsearchShardsOptionalParameters struct {
	Index *string
	Node  *string
	State *ElasticsearchShardState
}

// NewListDiagnosticsElasticsearchShardsOptionalParameters creates an empty struct for parameters.
func NewListDiagnosticsElasticsearchShardsOptionalParameters() *ListDiagnosticsElasticsearchShardsOptionalParameters {
	this := ListDiagnosticsElasticsearchShardsOptionalParameters{}
	return &this
}

// WithIndex sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsElasticsearchShardsOptionalParameters) WithIndex(index string) *ListDiagnosticsElasticsearchShardsOptionalParameters {
	r.Index = &index
	return r
}

// WithNode sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsElasticsearchShardsOptionalParameters) WithNode(node string) *ListDiagnosticsElasticsearchShardsOptionalParameters {
	r.Node = &node
	return r
}

// WithState sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsElasticsearchShardsOptionalParameters) WithState(state ElasticsearchShardState) *ListDiagnosticsElasticsearchShardsOptionalParameters {
	r.State = &state
	return r
}

// ListDiagnosticsElasticsearchShards List Elasticsearch shards.
func (a *DiagnosticsApi) ListDiagnosticsElasticsearchShards(ctx _context.Context, orgName string, clusterName string, o ...ListDiagnosticsElasticsearchShardsOptionalParameters) (ElasticsearchShardList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ElasticsearchShardList
		optionalParams      ListDiagnosticsElasticsearchShardsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListDiagnosticsElasticsearchShardsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "listDiagnosticsElasticsearchShards",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/shards",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ListDiagnosticsElasticsearchShards")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/elasticsearch/shards"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Index != nil {
		localVarQueryParams.Add("index", common.ParameterToString(*optionalParams.Index, ""))
	}
	if optionalParams.Node != nil {
		localVarQueryParams.Add("node", common.ParameterToString(*optionalParams.Node, ""))
	}
	if optionalParams.State != nil {
		localVarQueryParams.Add("state", common.ParameterToString(*optionalParams.State, ""))
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

// ListDiagnosticsKingbaseSessionsOptionalParameters holds optional parameters for ListDiagnosticsKingbaseSessions.
type ListDiagnosticsKingbaseSessionsOptionalParameters struct {
	Limit *int64
}

// NewListDiagnosticsKingbaseSessionsOptionalParameters creates an empty struct for parameters.
func NewListDiagnosticsKingbaseSessionsOptionalParameters() *ListDiagnosticsKingbaseSessionsOptionalParameters {
	this := ListDiagnosticsKingbaseSessionsOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsKingbaseSessionsOptionalParameters) WithLimit(limit int64) *ListDiagnosticsKingbaseSessionsOptionalParameters {
	r.Limit = &limit
	return r
}

// ListDiagnosticsKingbaseSessions List Kingbase session diagnostics.
func (a *DiagnosticsApi) ListDiagnosticsKingbaseSessions(ctx _context.Context, orgName string, clusterName string, o ...ListDiagnosticsKingbaseSessionsOptionalParameters) (KingbaseSessionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue KingbaseSessionList
		optionalParams      ListDiagnosticsKingbaseSessionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListDiagnosticsKingbaseSessionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "listDiagnosticsKingbaseSessions",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/kingbase/sessions",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ListDiagnosticsKingbaseSessions")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/kingbase/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 500 {
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

// ListDiagnosticsMssqlSessionsOptionalParameters holds optional parameters for ListDiagnosticsMssqlSessions.
type ListDiagnosticsMssqlSessionsOptionalParameters struct {
	Limit *int64
}

// NewListDiagnosticsMssqlSessionsOptionalParameters creates an empty struct for parameters.
func NewListDiagnosticsMssqlSessionsOptionalParameters() *ListDiagnosticsMssqlSessionsOptionalParameters {
	this := ListDiagnosticsMssqlSessionsOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsMssqlSessionsOptionalParameters) WithLimit(limit int64) *ListDiagnosticsMssqlSessionsOptionalParameters {
	r.Limit = &limit
	return r
}

// ListDiagnosticsMssqlSessions List SQL Server sessions.
func (a *DiagnosticsApi) ListDiagnosticsMssqlSessions(ctx _context.Context, orgName string, clusterName string, o ...ListDiagnosticsMssqlSessionsOptionalParameters) (MssqlSessionSnapshot, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MssqlSessionSnapshot
		optionalParams      ListDiagnosticsMssqlSessionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListDiagnosticsMssqlSessionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "listDiagnosticsMssqlSessions",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mssql/sessions",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ListDiagnosticsMssqlSessions")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/mssql/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
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

// ListDiagnosticsPostgresqlSessionsOptionalParameters holds optional parameters for ListDiagnosticsPostgresqlSessions.
type ListDiagnosticsPostgresqlSessionsOptionalParameters struct {
	Limit *int64
}

// NewListDiagnosticsPostgresqlSessionsOptionalParameters creates an empty struct for parameters.
func NewListDiagnosticsPostgresqlSessionsOptionalParameters() *ListDiagnosticsPostgresqlSessionsOptionalParameters {
	this := ListDiagnosticsPostgresqlSessionsOptionalParameters{}
	return &this
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *ListDiagnosticsPostgresqlSessionsOptionalParameters) WithLimit(limit int64) *ListDiagnosticsPostgresqlSessionsOptionalParameters {
	r.Limit = &limit
	return r
}

// ListDiagnosticsPostgresqlSessions List PostgreSQL session basic diagnostics.
// List PostgreSQL session basic diagnostics records. The response includes waitEventType and waitEvent so clients can identify lock-waiting sessions without loading lock rows.
func (a *DiagnosticsApi) ListDiagnosticsPostgresqlSessions(ctx _context.Context, orgName string, clusterName string, o ...ListDiagnosticsPostgresqlSessionsOptionalParameters) ([]PostgresqlSession, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue []PostgresqlSession
		optionalParams      ListDiagnosticsPostgresqlSessionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListDiagnosticsPostgresqlSessionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "listDiagnosticsPostgresqlSessions",
		Path:        "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ListDiagnosticsPostgresqlSessions")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", common.ParameterToString(*optionalParams.Limit, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 500 {
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

// NewDiagnosticsApi Returns NewDiagnosticsApi.
func NewDiagnosticsApi(client *common.APIClient) *DiagnosticsApi {
	return &DiagnosticsApi{
		Client: client,
	}
}
