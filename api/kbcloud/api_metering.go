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

// MeteringApi service type
type MeteringApi common.Service

// ListOrgMeteringTracksOptionalParameters holds optional parameters for ListOrgMeteringTracks.
type ListOrgMeteringTracksOptionalParameters struct {
	ClusterId       *string
	ProjectName     *string
	EnvironmentName *string
	PageNumber      *int64
	PageSize        *int64
}

// NewListOrgMeteringTracksOptionalParameters creates an empty struct for parameters.
func NewListOrgMeteringTracksOptionalParameters() *ListOrgMeteringTracksOptionalParameters {
	this := ListOrgMeteringTracksOptionalParameters{}
	return &this
}

// WithClusterId sets the corresponding parameter name and returns the struct.
func (r *ListOrgMeteringTracksOptionalParameters) WithClusterId(clusterId string) *ListOrgMeteringTracksOptionalParameters {
	r.ClusterId = &clusterId
	return r
}

// WithProjectName sets the corresponding parameter name and returns the struct.
func (r *ListOrgMeteringTracksOptionalParameters) WithProjectName(projectName string) *ListOrgMeteringTracksOptionalParameters {
	r.ProjectName = &projectName
	return r
}

// WithEnvironmentName sets the corresponding parameter name and returns the struct.
func (r *ListOrgMeteringTracksOptionalParameters) WithEnvironmentName(environmentName string) *ListOrgMeteringTracksOptionalParameters {
	r.EnvironmentName = &environmentName
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListOrgMeteringTracksOptionalParameters) WithPageNumber(pageNumber int64) *ListOrgMeteringTracksOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListOrgMeteringTracksOptionalParameters) WithPageSize(pageSize int64) *ListOrgMeteringTracksOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// ListOrgMeteringTracks List metering tracks in the organization.
func (a *MeteringApi) ListOrgMeteringTracks(ctx _context.Context, orgName string, start int64, end int64, o ...ListOrgMeteringTracksOptionalParameters) (MeteringTrackList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue MeteringTrackList
		optionalParams      ListOrgMeteringTracksOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, common.ReportError("only one argument of type ListOrgMeteringTracksOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	// Add api info to context
	apiInfo := common.APIInfo{
		Tag:         "metering",
		OperationID: "listOrgMeteringTracks",
		Path:        "/api/v1/organizations/{orgName}/metering/tracks",
		Version:     "",
	}
	ctx = context.WithValue(ctx, common.APIInfoCtxKey, apiInfo)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".MeteringApi.ListOrgMeteringTracks")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/metering/tracks"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("start", common.ParameterToString(start, ""))
	localVarQueryParams.Add("end", common.ParameterToString(end, ""))
	if optionalParams.ClusterId != nil {
		localVarQueryParams.Add("clusterID", common.ParameterToString(*optionalParams.ClusterId, ""))
	}
	if optionalParams.ProjectName != nil {
		localVarQueryParams.Add("projectName", common.ParameterToString(*optionalParams.ProjectName, ""))
	}
	if optionalParams.EnvironmentName != nil {
		localVarQueryParams.Add("environmentName", common.ParameterToString(*optionalParams.EnvironmentName, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("pageNumber", common.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("pageSize", common.ParameterToString(*optionalParams.PageSize, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// NewMeteringApi Returns NewMeteringApi.
func NewMeteringApi(client *common.APIClient) *MeteringApi {
	return &MeteringApi{
		Client: client,
	}
}
