// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterLogPagination Cluster log pagination is the pagination of the cluster log
type ClusterLogPagination struct {
	EndTime       *string `json:"endTime,omitempty"`
	Limit         *string `json:"limit,omitempty"`
	NextEndTime   *string `json:"nextEndTime,omitempty"`
	NextStartTime *string `json:"nextStartTime,omitempty"`
	StartTime     *string `json:"startTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterLogPagination instantiates a new ClusterLogPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterLogPagination() *ClusterLogPagination {
	this := ClusterLogPagination{}
	return &this
}

// NewClusterLogPaginationWithDefaults instantiates a new ClusterLogPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterLogPaginationWithDefaults() *ClusterLogPagination {
	this := ClusterLogPagination{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ClusterLogPagination) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogPagination) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ClusterLogPagination) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *ClusterLogPagination) SetEndTime(v string) {
	o.EndTime = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ClusterLogPagination) GetLimit() string {
	if o == nil || o.Limit == nil {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogPagination) GetLimitOk() (*string, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ClusterLogPagination) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *ClusterLogPagination) SetLimit(v string) {
	o.Limit = &v
}

// GetNextEndTime returns the NextEndTime field value if set, zero value otherwise.
func (o *ClusterLogPagination) GetNextEndTime() string {
	if o == nil || o.NextEndTime == nil {
		var ret string
		return ret
	}
	return *o.NextEndTime
}

// GetNextEndTimeOk returns a tuple with the NextEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogPagination) GetNextEndTimeOk() (*string, bool) {
	if o == nil || o.NextEndTime == nil {
		return nil, false
	}
	return o.NextEndTime, true
}

// HasNextEndTime returns a boolean if a field has been set.
func (o *ClusterLogPagination) HasNextEndTime() bool {
	return o != nil && o.NextEndTime != nil
}

// SetNextEndTime gets a reference to the given string and assigns it to the NextEndTime field.
func (o *ClusterLogPagination) SetNextEndTime(v string) {
	o.NextEndTime = &v
}

// GetNextStartTime returns the NextStartTime field value if set, zero value otherwise.
func (o *ClusterLogPagination) GetNextStartTime() string {
	if o == nil || o.NextStartTime == nil {
		var ret string
		return ret
	}
	return *o.NextStartTime
}

// GetNextStartTimeOk returns a tuple with the NextStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogPagination) GetNextStartTimeOk() (*string, bool) {
	if o == nil || o.NextStartTime == nil {
		return nil, false
	}
	return o.NextStartTime, true
}

// HasNextStartTime returns a boolean if a field has been set.
func (o *ClusterLogPagination) HasNextStartTime() bool {
	return o != nil && o.NextStartTime != nil
}

// SetNextStartTime gets a reference to the given string and assigns it to the NextStartTime field.
func (o *ClusterLogPagination) SetNextStartTime(v string) {
	o.NextStartTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ClusterLogPagination) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogPagination) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ClusterLogPagination) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ClusterLogPagination) SetStartTime(v string) {
	o.StartTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterLogPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextEndTime != nil {
		toSerialize["nextEndTime"] = o.NextEndTime
	}
	if o.NextStartTime != nil {
		toSerialize["nextStartTime"] = o.NextStartTime
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterLogPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndTime       *string `json:"endTime,omitempty"`
		Limit         *string `json:"limit,omitempty"`
		NextEndTime   *string `json:"nextEndTime,omitempty"`
		NextStartTime *string `json:"nextStartTime,omitempty"`
		StartTime     *string `json:"startTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"endTime", "limit", "nextEndTime", "nextStartTime", "startTime"})
	} else {
		return err
	}
	o.EndTime = all.EndTime
	o.Limit = all.Limit
	o.NextEndTime = all.NextEndTime
	o.NextStartTime = all.NextStartTime
	o.StartTime = all.StartTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
