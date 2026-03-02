// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterExecutionLogAggregateResponse Aggregate response for cluster execution logs
type ClusterExecutionLogAggregateResponse struct {
	Items []ClusterExecutionLogAggregateItem `json:"items,omitempty"`
	// Aggregate summary for execution logs
	Summary *ClusterExecutionLogAggregateSummary `json:"summary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterExecutionLogAggregateResponse instantiates a new ClusterExecutionLogAggregateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterExecutionLogAggregateResponse() *ClusterExecutionLogAggregateResponse {
	this := ClusterExecutionLogAggregateResponse{}
	return &this
}

// NewClusterExecutionLogAggregateResponseWithDefaults instantiates a new ClusterExecutionLogAggregateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterExecutionLogAggregateResponseWithDefaults() *ClusterExecutionLogAggregateResponse {
	this := ClusterExecutionLogAggregateResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterExecutionLogAggregateResponse) GetItems() []ClusterExecutionLogAggregateItem {
	if o == nil || o.Items == nil {
		var ret []ClusterExecutionLogAggregateItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogAggregateResponse) GetItemsOk() (*[]ClusterExecutionLogAggregateItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterExecutionLogAggregateResponse) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []ClusterExecutionLogAggregateItem and assigns it to the Items field.
func (o *ClusterExecutionLogAggregateResponse) SetItems(v []ClusterExecutionLogAggregateItem) {
	o.Items = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ClusterExecutionLogAggregateResponse) GetSummary() ClusterExecutionLogAggregateSummary {
	if o == nil || o.Summary == nil {
		var ret ClusterExecutionLogAggregateSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogAggregateResponse) GetSummaryOk() (*ClusterExecutionLogAggregateSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ClusterExecutionLogAggregateResponse) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given ClusterExecutionLogAggregateSummary and assigns it to the Summary field.
func (o *ClusterExecutionLogAggregateResponse) SetSummary(v ClusterExecutionLogAggregateSummary) {
	o.Summary = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterExecutionLogAggregateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterExecutionLogAggregateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items   []ClusterExecutionLogAggregateItem   `json:"items,omitempty"`
		Summary *ClusterExecutionLogAggregateSummary `json:"summary,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "summary"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = all.Items
	if all.Summary != nil && all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
