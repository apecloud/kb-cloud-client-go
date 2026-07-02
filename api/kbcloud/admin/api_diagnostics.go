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
// Explicitly trigger a safe PostgreSQL EXPLAIN for one SQL fingerprint. The request does not accept raw SQL. The backend uses only a deterministic server-side SELECT sample when available, normalizes it with the same SELECT-only guard as slow-log EXPLAIN, and never runs EXPLAIN ANALYZE or original SQL.
func (a *DiagnosticsApi) ExplainDiagnosticsPostgresqlSQLFingerprint(ctx _context.Context, orgName string, clusterName string, queryId string, database string, user string, o ...ExplainDiagnosticsPostgresqlSQLFingerprintOptionalParameters) (PostgresqlSQLFingerprintExplainResponse, *_nethttp.Response, error) {
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
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis/queries/{queryID}/explain",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.ExplainDiagnosticsPostgresqlSQLFingerprint")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis/queries/{queryID}/explain"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queryID"+"}", _neturl.PathEscape(common.ParameterToString(queryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("database", common.ParameterToString(database, ""))
	localVarQueryParams.Add("user", common.ParameterToString(user, ""))
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

// GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters holds optional parameters for GetDiagnosticsPostgresqlSQLFingerprintDetail.
type GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters struct {
	Fingerprint *string
}

// NewGetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters creates an empty struct for parameters.
func NewGetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters() *GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters {
	this := GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters{}
	return &this
}

// WithFingerprint sets the corresponding parameter name and returns the struct.
func (r *GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters) WithFingerprint(fingerprint string) *GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters {
	r.Fingerprint = &fingerprint
	return r
}

// GetDiagnosticsPostgresqlSQLFingerprintDetail Get PostgreSQL SQL fingerprint detail.
// Get one PostgreSQL SQL fingerprint detail from the current pg_stat_statements ranking scope. This endpoint is keyed by queryID, database, and user. It returns redacted SQL summary and current statistics only; it does not expose raw SQL, trends, slow-log relation, I/O analysis, suggestions, or execution plans.
func (a *DiagnosticsApi) GetDiagnosticsPostgresqlSQLFingerprintDetail(ctx _context.Context, orgName string, clusterName string, queryId string, database string, user string, o ...GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters) (PostgresqlSQLFingerprintDetail, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PostgresqlSQLFingerprintDetail
		optionalParams      GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type GetDiagnosticsPostgresqlSQLFingerprintDetailOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "diagnostics",
		OperationID: "getDiagnosticsPostgresqlSQLFingerprintDetail",
		Path:        "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis/queries/{queryID}",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".DiagnosticsApi.GetDiagnosticsPostgresqlSQLFingerprintDetail")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/admin/v1/organizations/{orgName}/clusters/{clusterName}/diagnostics/postgresql/sqlAnalysis/queries/{queryID}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queryID"+"}", _neturl.PathEscape(common.ParameterToString(queryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("database", common.ParameterToString(database, ""))
	localVarQueryParams.Add("user", common.ParameterToString(user, ""))
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
