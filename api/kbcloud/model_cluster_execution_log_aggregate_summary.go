// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterExecutionLogAggregateSummary Aggregate summary for execution logs
type ClusterExecutionLogAggregateSummary struct {
	TotalCount *int64 `json:"totalCount,omitempty"`
	GroupCount *int64 `json:"groupCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterExecutionLogAggregateSummary instantiates a new ClusterExecutionLogAggregateSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterExecutionLogAggregateSummary() *ClusterExecutionLogAggregateSummary {
	this := ClusterExecutionLogAggregateSummary{}
	return &this
}

// NewClusterExecutionLogAggregateSummaryWithDefaults instantiates a new ClusterExecutionLogAggregateSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterExecutionLogAggregateSummaryWithDefaults() *ClusterExecutionLogAggregateSummary {
	this := ClusterExecutionLogAggregateSummary{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ClusterExecutionLogAggregateSummary) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogAggregateSummary) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ClusterExecutionLogAggregateSummary) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ClusterExecutionLogAggregateSummary) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetGroupCount returns the GroupCount field value if set, zero value otherwise.
func (o *ClusterExecutionLogAggregateSummary) GetGroupCount() int64 {
	if o == nil || o.GroupCount == nil {
		var ret int64
		return ret
	}
	return *o.GroupCount
}

// GetGroupCountOk returns a tuple with the GroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogAggregateSummary) GetGroupCountOk() (*int64, bool) {
	if o == nil || o.GroupCount == nil {
		return nil, false
	}
	return o.GroupCount, true
}

// HasGroupCount returns a boolean if a field has been set.
func (o *ClusterExecutionLogAggregateSummary) HasGroupCount() bool {
	return o != nil && o.GroupCount != nil
}

// SetGroupCount gets a reference to the given int64 and assigns it to the GroupCount field.
func (o *ClusterExecutionLogAggregateSummary) SetGroupCount(v int64) {
	o.GroupCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterExecutionLogAggregateSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.GroupCount != nil {
		toSerialize["groupCount"] = o.GroupCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterExecutionLogAggregateSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCount *int64 `json:"totalCount,omitempty"`
		GroupCount *int64 `json:"groupCount,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalCount", "groupCount"})
	} else {
		return err
	}
	o.TotalCount = all.TotalCount
	o.GroupCount = all.GroupCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
