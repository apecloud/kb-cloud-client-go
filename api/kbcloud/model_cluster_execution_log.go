// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterExecutionLog Cluster execution log is the execution log of the cluster
type ClusterExecutionLog struct {
	Items []ClusterExecutionLogItem `json:"items,omitempty"`
	// Cluster log pagination is the pagination of the cluster log
	Pagination *ClusterLogPagination `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterExecutionLog instantiates a new ClusterExecutionLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterExecutionLog() *ClusterExecutionLog {
	this := ClusterExecutionLog{}
	return &this
}

// NewClusterExecutionLogWithDefaults instantiates a new ClusterExecutionLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterExecutionLogWithDefaults() *ClusterExecutionLog {
	this := ClusterExecutionLog{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterExecutionLog) GetItems() []ClusterExecutionLogItem {
	if o == nil || o.Items == nil {
		var ret []ClusterExecutionLogItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetItemsOk() (*[]ClusterExecutionLogItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterExecutionLog) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []ClusterExecutionLogItem and assigns it to the Items field.
func (o *ClusterExecutionLog) SetItems(v []ClusterExecutionLogItem) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ClusterExecutionLog) GetPagination() ClusterLogPagination {
	if o == nil || o.Pagination == nil {
		var ret ClusterLogPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLog) GetPaginationOk() (*ClusterLogPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ClusterExecutionLog) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given ClusterLogPagination and assigns it to the Pagination field.
func (o *ClusterExecutionLog) SetPagination(v ClusterLogPagination) {
	o.Pagination = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterExecutionLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterExecutionLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      []ClusterExecutionLogItem `json:"items,omitempty"`
		Pagination *ClusterLogPagination     `json:"pagination,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "pagination"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = all.Items
	if all.Pagination != nil && all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = all.Pagination

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
