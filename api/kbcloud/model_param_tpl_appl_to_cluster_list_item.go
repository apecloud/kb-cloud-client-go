// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// ParamTplApplToClusterListItem parameter template applicable to the cluster information
type ParamTplApplToClusterListItem struct {
	// The number of parameters in the parameter template
	Count int32 `json:"count"`
	// Name of parameter template. Name must be unique within an Org
	Name string `json:"name"`
	// whether to restart after applying this parameter template or not
	NeedRestart bool `json:"needRestart"`
	// the template partition
	Partition string `json:"partition"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplApplToClusterListItem instantiates a new ParamTplApplToClusterListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplApplToClusterListItem(count int32, name string, needRestart bool, partition string) *ParamTplApplToClusterListItem {
	this := ParamTplApplToClusterListItem{}
	this.Count = count
	this.Name = name
	this.NeedRestart = needRestart
	this.Partition = partition
	return &this
}

// NewParamTplApplToClusterListItemWithDefaults instantiates a new ParamTplApplToClusterListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplApplToClusterListItemWithDefaults() *ParamTplApplToClusterListItem {
	this := ParamTplApplToClusterListItem{}
	return &this
}

// GetCount returns the Count field value.
func (o *ParamTplApplToClusterListItem) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ParamTplApplToClusterListItem) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *ParamTplApplToClusterListItem) SetCount(v int32) {
	o.Count = v
}

// GetName returns the Name field value.
func (o *ParamTplApplToClusterListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParamTplApplToClusterListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParamTplApplToClusterListItem) SetName(v string) {
	o.Name = v
}

// GetNeedRestart returns the NeedRestart field value.
func (o *ParamTplApplToClusterListItem) GetNeedRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NeedRestart
}

// GetNeedRestartOk returns a tuple with the NeedRestart field value
// and a boolean to check if the value has been set.
func (o *ParamTplApplToClusterListItem) GetNeedRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedRestart, true
}

// SetNeedRestart sets field value.
func (o *ParamTplApplToClusterListItem) SetNeedRestart(v bool) {
	o.NeedRestart = v
}

// GetPartition returns the Partition field value.
func (o *ParamTplApplToClusterListItem) GetPartition() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value
// and a boolean to check if the value has been set.
func (o *ParamTplApplToClusterListItem) GetPartitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partition, true
}

// SetPartition sets field value.
func (o *ParamTplApplToClusterListItem) SetPartition(v string) {
	o.Partition = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplApplToClusterListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	toSerialize["name"] = o.Name
	toSerialize["needRestart"] = o.NeedRestart
	toSerialize["partition"] = o.Partition

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParamTplApplToClusterListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count       *int32  `json:"count"`
		Name        *string `json:"name"`
		NeedRestart *bool   `json:"needRestart"`
		Partition   *string `json:"partition"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.NeedRestart == nil {
		return fmt.Errorf("required field needRestart missing")
	}
	if all.Partition == nil {
		return fmt.Errorf("required field partition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"count", "name", "needRestart", "partition"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.Name = *all.Name
	o.NeedRestart = *all.NeedRestart
	o.Partition = *all.Partition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
