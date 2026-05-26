// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogDetail Slow log fields emitted by oteld
type ClusterSlowLogDetail struct {
	// Query template identifier
	TemplateId *string `json:"templateId,omitempty"`
	// Normalized SQL or command template
	NormalizedQuery *string `json:"normalizedQuery,omitempty"`
	// Application or client name reported by the database
	AppName *string `json:"appName,omitempty"`
	// Rows examined by the query
	RowsExamined *int64 `json:"rowsExamined,omitempty"`
	// Rows returned or sent by the query
	RowsSent *int64 `json:"rowsSent,omitempty"`
	// Lock wait time in seconds
	LockTime *float64 `json:"lockTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogDetail instantiates a new ClusterSlowLogDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogDetail() *ClusterSlowLogDetail {
	this := ClusterSlowLogDetail{}
	return &this
}

// NewClusterSlowLogDetailWithDefaults instantiates a new ClusterSlowLogDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogDetailWithDefaults() *ClusterSlowLogDetail {
	this := ClusterSlowLogDetail{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ClusterSlowLogDetail) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDetail) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ClusterSlowLogDetail) HasTemplateId() bool {
	return o != nil && o.TemplateId != nil
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *ClusterSlowLogDetail) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetNormalizedQuery returns the NormalizedQuery field value if set, zero value otherwise.
func (o *ClusterSlowLogDetail) GetNormalizedQuery() string {
	if o == nil || o.NormalizedQuery == nil {
		var ret string
		return ret
	}
	return *o.NormalizedQuery
}

// GetNormalizedQueryOk returns a tuple with the NormalizedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDetail) GetNormalizedQueryOk() (*string, bool) {
	if o == nil || o.NormalizedQuery == nil {
		return nil, false
	}
	return o.NormalizedQuery, true
}

// HasNormalizedQuery returns a boolean if a field has been set.
func (o *ClusterSlowLogDetail) HasNormalizedQuery() bool {
	return o != nil && o.NormalizedQuery != nil
}

// SetNormalizedQuery gets a reference to the given string and assigns it to the NormalizedQuery field.
func (o *ClusterSlowLogDetail) SetNormalizedQuery(v string) {
	o.NormalizedQuery = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *ClusterSlowLogDetail) GetAppName() string {
	if o == nil || o.AppName == nil {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDetail) GetAppNameOk() (*string, bool) {
	if o == nil || o.AppName == nil {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *ClusterSlowLogDetail) HasAppName() bool {
	return o != nil && o.AppName != nil
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *ClusterSlowLogDetail) SetAppName(v string) {
	o.AppName = &v
}

// GetRowsExamined returns the RowsExamined field value if set, zero value otherwise.
func (o *ClusterSlowLogDetail) GetRowsExamined() int64 {
	if o == nil || o.RowsExamined == nil {
		var ret int64
		return ret
	}
	return *o.RowsExamined
}

// GetRowsExaminedOk returns a tuple with the RowsExamined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDetail) GetRowsExaminedOk() (*int64, bool) {
	if o == nil || o.RowsExamined == nil {
		return nil, false
	}
	return o.RowsExamined, true
}

// HasRowsExamined returns a boolean if a field has been set.
func (o *ClusterSlowLogDetail) HasRowsExamined() bool {
	return o != nil && o.RowsExamined != nil
}

// SetRowsExamined gets a reference to the given int64 and assigns it to the RowsExamined field.
func (o *ClusterSlowLogDetail) SetRowsExamined(v int64) {
	o.RowsExamined = &v
}

// GetRowsSent returns the RowsSent field value if set, zero value otherwise.
func (o *ClusterSlowLogDetail) GetRowsSent() int64 {
	if o == nil || o.RowsSent == nil {
		var ret int64
		return ret
	}
	return *o.RowsSent
}

// GetRowsSentOk returns a tuple with the RowsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDetail) GetRowsSentOk() (*int64, bool) {
	if o == nil || o.RowsSent == nil {
		return nil, false
	}
	return o.RowsSent, true
}

// HasRowsSent returns a boolean if a field has been set.
func (o *ClusterSlowLogDetail) HasRowsSent() bool {
	return o != nil && o.RowsSent != nil
}

// SetRowsSent gets a reference to the given int64 and assigns it to the RowsSent field.
func (o *ClusterSlowLogDetail) SetRowsSent(v int64) {
	o.RowsSent = &v
}

// GetLockTime returns the LockTime field value if set, zero value otherwise.
func (o *ClusterSlowLogDetail) GetLockTime() float64 {
	if o == nil || o.LockTime == nil {
		var ret float64
		return ret
	}
	return *o.LockTime
}

// GetLockTimeOk returns a tuple with the LockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDetail) GetLockTimeOk() (*float64, bool) {
	if o == nil || o.LockTime == nil {
		return nil, false
	}
	return o.LockTime, true
}

// HasLockTime returns a boolean if a field has been set.
func (o *ClusterSlowLogDetail) HasLockTime() bool {
	return o != nil && o.LockTime != nil
}

// SetLockTime gets a reference to the given float64 and assigns it to the LockTime field.
func (o *ClusterSlowLogDetail) SetLockTime(v float64) {
	o.LockTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogDetail) MarshalJSON() ([]byte, error) {
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
	if o.AppName != nil {
		toSerialize["appName"] = o.AppName
	}
	if o.RowsExamined != nil {
		toSerialize["rowsExamined"] = o.RowsExamined
	}
	if o.RowsSent != nil {
		toSerialize["rowsSent"] = o.RowsSent
	}
	if o.LockTime != nil {
		toSerialize["lockTime"] = o.LockTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TemplateId      *string  `json:"templateId,omitempty"`
		NormalizedQuery *string  `json:"normalizedQuery,omitempty"`
		AppName         *string  `json:"appName,omitempty"`
		RowsExamined    *int64   `json:"rowsExamined,omitempty"`
		RowsSent        *int64   `json:"rowsSent,omitempty"`
		LockTime        *float64 `json:"lockTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"templateId", "normalizedQuery", "appName", "rowsExamined", "rowsSent", "lockTime"})
	} else {
		return err
	}
	o.TemplateId = all.TemplateId
	o.NormalizedQuery = all.NormalizedQuery
	o.AppName = all.AppName
	o.RowsExamined = all.RowsExamined
	o.RowsSent = all.RowsSent
	o.LockTime = all.LockTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
