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

// DiagnosticsApi service type
type DiagnosticsApi common.Service

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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sqlAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsDamengSQLAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sqlAnalysis"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions/{sessionId}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsDamengSession")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions/{sessionId}"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions/{sessionId}/lockAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsDamengSessionLockAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions/{sessionId}/lockAnalysis"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/spaceAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsDamengSpaceAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/spaceAnalysis"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/performanceTrends",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlPerformanceTrends")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/performanceTrends"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSQLAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions/{pid}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSession")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions/{pid}"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions/{pid}/lockAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSessionLockAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions/{pid}/lockAnalysis"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/spaceAnalysis",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSpaceAnalysis")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/spaceAnalysis"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ListDiagnosticsDamengSessions")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/dameng/sessions"
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ListDiagnosticsPostgresqlSessions")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sessions"
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
