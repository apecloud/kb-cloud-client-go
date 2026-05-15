// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogTemplate Aggregated slow log query template
type ClusterSlowLogTemplate struct {
	// Query template identifier
	TemplateId *string `json:"templateId,omitempty"`
	// Normalized SQL or command template
	NormalizedQuery *string `json:"normalizedQuery,omitempty"`
	// Representative SQL sample for the template
	SampleSql  *string `json:"sampleSQL,omitempty"`
	TotalCount *int64  `json:"totalCount,omitempty"`
	// Total execution time in seconds
	TotalExecutionTime *float64 `json:"totalExecutionTime,omitempty"`
	// Average execution time in seconds
	AvgExecutionTime *float64 `json:"avgExecutionTime,omitempty"`
	// Maximum execution time in seconds
	MaxExecutionTime *float64 `json:"maxExecutionTime,omitempty"`
	// P95 execution time in seconds
	P95ExecutionTime *float64 `json:"p95ExecutionTime,omitempty"`
	// Number of unique database names in the time range
	DbCount *int64 `json:"dbCount,omitempty"`
	// Number of unique users in the time range
	UserCount *int64 `json:"userCount,omitempty"`
	// Number of unique client IPs in the time range
	ClientIpCount *int64 `json:"clientIPCount,omitempty"`
	// Number of unique application names in the time range
	AppCount *int64 `json:"appCount,omitempty"`
	// Average rows examined by the template
	AvgRowsExamined *float64 `json:"avgRowsExamined,omitempty"`
	// Average rows returned or sent by the template
	AvgRowsSent *float64 `json:"avgRowsSent,omitempty"`
	// Average lock wait time in seconds
	AvgLockTime *float64 `json:"avgLockTime,omitempty"`
	// Sampled database names for the template
	DbNames []string `json:"dbNames,omitempty"`
	// Sampled users for the template
	Users []string `json:"users,omitempty"`
	// Sampled client IPs for the template
	ClientIPs []string `json:"clientIPs,omitempty"`
	// Sampled application names for the template
	Apps []string `json:"apps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogTemplate instantiates a new ClusterSlowLogTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogTemplate() *ClusterSlowLogTemplate {
	this := ClusterSlowLogTemplate{}
	return &this
}

// NewClusterSlowLogTemplateWithDefaults instantiates a new ClusterSlowLogTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogTemplateWithDefaults() *ClusterSlowLogTemplate {
	this := ClusterSlowLogTemplate{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasTemplateId() bool {
	return o != nil && o.TemplateId != nil
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *ClusterSlowLogTemplate) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetNormalizedQuery returns the NormalizedQuery field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetNormalizedQuery() string {
	if o == nil || o.NormalizedQuery == nil {
		var ret string
		return ret
	}
	return *o.NormalizedQuery
}

// GetNormalizedQueryOk returns a tuple with the NormalizedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetNormalizedQueryOk() (*string, bool) {
	if o == nil || o.NormalizedQuery == nil {
		return nil, false
	}
	return o.NormalizedQuery, true
}

// HasNormalizedQuery returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasNormalizedQuery() bool {
	return o != nil && o.NormalizedQuery != nil
}

// SetNormalizedQuery gets a reference to the given string and assigns it to the NormalizedQuery field.
func (o *ClusterSlowLogTemplate) SetNormalizedQuery(v string) {
	o.NormalizedQuery = &v
}

// GetSampleSql returns the SampleSql field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetSampleSql() string {
	if o == nil || o.SampleSql == nil {
		var ret string
		return ret
	}
	return *o.SampleSql
}

// GetSampleSqlOk returns a tuple with the SampleSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetSampleSqlOk() (*string, bool) {
	if o == nil || o.SampleSql == nil {
		return nil, false
	}
	return o.SampleSql, true
}

// HasSampleSql returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasSampleSql() bool {
	return o != nil && o.SampleSql != nil
}

// SetSampleSql gets a reference to the given string and assigns it to the SampleSql field.
func (o *ClusterSlowLogTemplate) SetSampleSql(v string) {
	o.SampleSql = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ClusterSlowLogTemplate) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetTotalExecutionTime returns the TotalExecutionTime field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetTotalExecutionTime() float64 {
	if o == nil || o.TotalExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.TotalExecutionTime
}

// GetTotalExecutionTimeOk returns a tuple with the TotalExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetTotalExecutionTimeOk() (*float64, bool) {
	if o == nil || o.TotalExecutionTime == nil {
		return nil, false
	}
	return o.TotalExecutionTime, true
}

// HasTotalExecutionTime returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasTotalExecutionTime() bool {
	return o != nil && o.TotalExecutionTime != nil
}

// SetTotalExecutionTime gets a reference to the given float64 and assigns it to the TotalExecutionTime field.
func (o *ClusterSlowLogTemplate) SetTotalExecutionTime(v float64) {
	o.TotalExecutionTime = &v
}

// GetAvgExecutionTime returns the AvgExecutionTime field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetAvgExecutionTime() float64 {
	if o == nil || o.AvgExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.AvgExecutionTime
}

// GetAvgExecutionTimeOk returns a tuple with the AvgExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetAvgExecutionTimeOk() (*float64, bool) {
	if o == nil || o.AvgExecutionTime == nil {
		return nil, false
	}
	return o.AvgExecutionTime, true
}

// HasAvgExecutionTime returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasAvgExecutionTime() bool {
	return o != nil && o.AvgExecutionTime != nil
}

// SetAvgExecutionTime gets a reference to the given float64 and assigns it to the AvgExecutionTime field.
func (o *ClusterSlowLogTemplate) SetAvgExecutionTime(v float64) {
	o.AvgExecutionTime = &v
}

// GetMaxExecutionTime returns the MaxExecutionTime field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetMaxExecutionTime() float64 {
	if o == nil || o.MaxExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.MaxExecutionTime
}

// GetMaxExecutionTimeOk returns a tuple with the MaxExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetMaxExecutionTimeOk() (*float64, bool) {
	if o == nil || o.MaxExecutionTime == nil {
		return nil, false
	}
	return o.MaxExecutionTime, true
}

// HasMaxExecutionTime returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasMaxExecutionTime() bool {
	return o != nil && o.MaxExecutionTime != nil
}

// SetMaxExecutionTime gets a reference to the given float64 and assigns it to the MaxExecutionTime field.
func (o *ClusterSlowLogTemplate) SetMaxExecutionTime(v float64) {
	o.MaxExecutionTime = &v
}

// GetP95ExecutionTime returns the P95ExecutionTime field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetP95ExecutionTime() float64 {
	if o == nil || o.P95ExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.P95ExecutionTime
}

// GetP95ExecutionTimeOk returns a tuple with the P95ExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetP95ExecutionTimeOk() (*float64, bool) {
	if o == nil || o.P95ExecutionTime == nil {
		return nil, false
	}
	return o.P95ExecutionTime, true
}

// HasP95ExecutionTime returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasP95ExecutionTime() bool {
	return o != nil && o.P95ExecutionTime != nil
}

// SetP95ExecutionTime gets a reference to the given float64 and assigns it to the P95ExecutionTime field.
func (o *ClusterSlowLogTemplate) SetP95ExecutionTime(v float64) {
	o.P95ExecutionTime = &v
}

// GetDbCount returns the DbCount field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetDbCount() int64 {
	if o == nil || o.DbCount == nil {
		var ret int64
		return ret
	}
	return *o.DbCount
}

// GetDbCountOk returns a tuple with the DbCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetDbCountOk() (*int64, bool) {
	if o == nil || o.DbCount == nil {
		return nil, false
	}
	return o.DbCount, true
}

// HasDbCount returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasDbCount() bool {
	return o != nil && o.DbCount != nil
}

// SetDbCount gets a reference to the given int64 and assigns it to the DbCount field.
func (o *ClusterSlowLogTemplate) SetDbCount(v int64) {
	o.DbCount = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetUserCount() int64 {
	if o == nil || o.UserCount == nil {
		var ret int64
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetUserCountOk() (*int64, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasUserCount() bool {
	return o != nil && o.UserCount != nil
}

// SetUserCount gets a reference to the given int64 and assigns it to the UserCount field.
func (o *ClusterSlowLogTemplate) SetUserCount(v int64) {
	o.UserCount = &v
}

// GetClientIpCount returns the ClientIpCount field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetClientIpCount() int64 {
	if o == nil || o.ClientIpCount == nil {
		var ret int64
		return ret
	}
	return *o.ClientIpCount
}

// GetClientIpCountOk returns a tuple with the ClientIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetClientIpCountOk() (*int64, bool) {
	if o == nil || o.ClientIpCount == nil {
		return nil, false
	}
	return o.ClientIpCount, true
}

// HasClientIpCount returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasClientIpCount() bool {
	return o != nil && o.ClientIpCount != nil
}

// SetClientIpCount gets a reference to the given int64 and assigns it to the ClientIpCount field.
func (o *ClusterSlowLogTemplate) SetClientIpCount(v int64) {
	o.ClientIpCount = &v
}

// GetAppCount returns the AppCount field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetAppCount() int64 {
	if o == nil || o.AppCount == nil {
		var ret int64
		return ret
	}
	return *o.AppCount
}

// GetAppCountOk returns a tuple with the AppCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetAppCountOk() (*int64, bool) {
	if o == nil || o.AppCount == nil {
		return nil, false
	}
	return o.AppCount, true
}

// HasAppCount returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasAppCount() bool {
	return o != nil && o.AppCount != nil
}

// SetAppCount gets a reference to the given int64 and assigns it to the AppCount field.
func (o *ClusterSlowLogTemplate) SetAppCount(v int64) {
	o.AppCount = &v
}

// GetAvgRowsExamined returns the AvgRowsExamined field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetAvgRowsExamined() float64 {
	if o == nil || o.AvgRowsExamined == nil {
		var ret float64
		return ret
	}
	return *o.AvgRowsExamined
}

// GetAvgRowsExaminedOk returns a tuple with the AvgRowsExamined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetAvgRowsExaminedOk() (*float64, bool) {
	if o == nil || o.AvgRowsExamined == nil {
		return nil, false
	}
	return o.AvgRowsExamined, true
}

// HasAvgRowsExamined returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasAvgRowsExamined() bool {
	return o != nil && o.AvgRowsExamined != nil
}

// SetAvgRowsExamined gets a reference to the given float64 and assigns it to the AvgRowsExamined field.
func (o *ClusterSlowLogTemplate) SetAvgRowsExamined(v float64) {
	o.AvgRowsExamined = &v
}

// GetAvgRowsSent returns the AvgRowsSent field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetAvgRowsSent() float64 {
	if o == nil || o.AvgRowsSent == nil {
		var ret float64
		return ret
	}
	return *o.AvgRowsSent
}

// GetAvgRowsSentOk returns a tuple with the AvgRowsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetAvgRowsSentOk() (*float64, bool) {
	if o == nil || o.AvgRowsSent == nil {
		return nil, false
	}
	return o.AvgRowsSent, true
}

// HasAvgRowsSent returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasAvgRowsSent() bool {
	return o != nil && o.AvgRowsSent != nil
}

// SetAvgRowsSent gets a reference to the given float64 and assigns it to the AvgRowsSent field.
func (o *ClusterSlowLogTemplate) SetAvgRowsSent(v float64) {
	o.AvgRowsSent = &v
}

// GetAvgLockTime returns the AvgLockTime field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetAvgLockTime() float64 {
	if o == nil || o.AvgLockTime == nil {
		var ret float64
		return ret
	}
	return *o.AvgLockTime
}

// GetAvgLockTimeOk returns a tuple with the AvgLockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetAvgLockTimeOk() (*float64, bool) {
	if o == nil || o.AvgLockTime == nil {
		return nil, false
	}
	return o.AvgLockTime, true
}

// HasAvgLockTime returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasAvgLockTime() bool {
	return o != nil && o.AvgLockTime != nil
}

// SetAvgLockTime gets a reference to the given float64 and assigns it to the AvgLockTime field.
func (o *ClusterSlowLogTemplate) SetAvgLockTime(v float64) {
	o.AvgLockTime = &v
}

// GetDbNames returns the DbNames field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetDbNames() []string {
	if o == nil || o.DbNames == nil {
		var ret []string
		return ret
	}
	return o.DbNames
}

// GetDbNamesOk returns a tuple with the DbNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetDbNamesOk() (*[]string, bool) {
	if o == nil || o.DbNames == nil {
		return nil, false
	}
	return &o.DbNames, true
}

// HasDbNames returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasDbNames() bool {
	return o != nil && o.DbNames != nil
}

// SetDbNames gets a reference to the given []string and assigns it to the DbNames field.
func (o *ClusterSlowLogTemplate) SetDbNames(v []string) {
	o.DbNames = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetUsersOk() (*[]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *ClusterSlowLogTemplate) SetUsers(v []string) {
	o.Users = v
}

// GetClientIPs returns the ClientIPs field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetClientIPs() []string {
	if o == nil || o.ClientIPs == nil {
		var ret []string
		return ret
	}
	return o.ClientIPs
}

// GetClientIPsOk returns a tuple with the ClientIPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetClientIPsOk() (*[]string, bool) {
	if o == nil || o.ClientIPs == nil {
		return nil, false
	}
	return &o.ClientIPs, true
}

// HasClientIPs returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasClientIPs() bool {
	return o != nil && o.ClientIPs != nil
}

// SetClientIPs gets a reference to the given []string and assigns it to the ClientIPs field.
func (o *ClusterSlowLogTemplate) SetClientIPs(v []string) {
	o.ClientIPs = v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *ClusterSlowLogTemplate) GetApps() []string {
	if o == nil || o.Apps == nil {
		var ret []string
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogTemplate) GetAppsOk() (*[]string, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return &o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ClusterSlowLogTemplate) HasApps() bool {
	return o != nil && o.Apps != nil
}

// SetApps gets a reference to the given []string and assigns it to the Apps field.
func (o *ClusterSlowLogTemplate) SetApps(v []string) {
	o.Apps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.NormalizedQuery != nil {
		toSerialize["normalizedQuery"] = o.NormalizedQuery
	}
	if o.SampleSql != nil {
		toSerialize["sampleSQL"] = o.SampleSql
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.TotalExecutionTime != nil {
		toSerialize["totalExecutionTime"] = o.TotalExecutionTime
	}
	if o.AvgExecutionTime != nil {
		toSerialize["avgExecutionTime"] = o.AvgExecutionTime
	}
	if o.MaxExecutionTime != nil {
		toSerialize["maxExecutionTime"] = o.MaxExecutionTime
	}
	if o.P95ExecutionTime != nil {
		toSerialize["p95ExecutionTime"] = o.P95ExecutionTime
	}
	if o.DbCount != nil {
		toSerialize["dbCount"] = o.DbCount
	}
	if o.UserCount != nil {
		toSerialize["userCount"] = o.UserCount
	}
	if o.ClientIpCount != nil {
		toSerialize["clientIPCount"] = o.ClientIpCount
	}
	if o.AppCount != nil {
		toSerialize["appCount"] = o.AppCount
	}
	if o.AvgRowsExamined != nil {
		toSerialize["avgRowsExamined"] = o.AvgRowsExamined
	}
	if o.AvgRowsSent != nil {
		toSerialize["avgRowsSent"] = o.AvgRowsSent
	}
	if o.AvgLockTime != nil {
		toSerialize["avgLockTime"] = o.AvgLockTime
	}
	if o.DbNames != nil {
		toSerialize["dbNames"] = o.DbNames
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.ClientIPs != nil {
		toSerialize["clientIPs"] = o.ClientIPs
	}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TemplateId         *string  `json:"templateId,omitempty"`
		NormalizedQuery    *string  `json:"normalizedQuery,omitempty"`
		SampleSql          *string  `json:"sampleSQL,omitempty"`
		TotalCount         *int64   `json:"totalCount,omitempty"`
		TotalExecutionTime *float64 `json:"totalExecutionTime,omitempty"`
		AvgExecutionTime   *float64 `json:"avgExecutionTime,omitempty"`
		MaxExecutionTime   *float64 `json:"maxExecutionTime,omitempty"`
		P95ExecutionTime   *float64 `json:"p95ExecutionTime,omitempty"`
		DbCount            *int64   `json:"dbCount,omitempty"`
		UserCount          *int64   `json:"userCount,omitempty"`
		ClientIpCount      *int64   `json:"clientIPCount,omitempty"`
		AppCount           *int64   `json:"appCount,omitempty"`
		AvgRowsExamined    *float64 `json:"avgRowsExamined,omitempty"`
		AvgRowsSent        *float64 `json:"avgRowsSent,omitempty"`
		AvgLockTime        *float64 `json:"avgLockTime,omitempty"`
		DbNames            []string `json:"dbNames,omitempty"`
		Users              []string `json:"users,omitempty"`
		ClientIPs          []string `json:"clientIPs,omitempty"`
		Apps               []string `json:"apps,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"templateId", "normalizedQuery", "sampleSQL", "totalCount", "totalExecutionTime", "avgExecutionTime", "maxExecutionTime", "p95ExecutionTime", "dbCount", "userCount", "clientIPCount", "appCount", "avgRowsExamined", "avgRowsSent", "avgLockTime", "dbNames", "users", "clientIPs", "apps"})
	} else {
		return err
	}
	o.TemplateId = all.TemplateId
	o.NormalizedQuery = all.NormalizedQuery
	o.SampleSql = all.SampleSql
	o.TotalCount = all.TotalCount
	o.TotalExecutionTime = all.TotalExecutionTime
	o.AvgExecutionTime = all.AvgExecutionTime
	o.MaxExecutionTime = all.MaxExecutionTime
	o.P95ExecutionTime = all.P95ExecutionTime
	o.DbCount = all.DbCount
	o.UserCount = all.UserCount
	o.ClientIpCount = all.ClientIpCount
	o.AppCount = all.AppCount
	o.AvgRowsExamined = all.AvgRowsExamined
	o.AvgRowsSent = all.AvgRowsSent
	o.AvgLockTime = all.AvgLockTime
	o.DbNames = all.DbNames
	o.Users = all.Users
	o.ClientIPs = all.ClientIPs
	o.Apps = all.Apps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
