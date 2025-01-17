// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParamTplApplyToClusterListItem parameter template applicable to the cluster information
type ParamTplApplyToClusterListItem struct {
	// The number of parameters in the parameter template
	Count int32 `json:"count"`
	// Name of parameter template. Name must be unique within an Org
	Name string `json:"name"`
	// id of parameter template
	Id *string `json:"id,omitempty"`
	// whether to restart after applying this parameter template or not
	NeedRestart bool `json:"needRestart"`
	// the template partition
	Partition string `json:"partition"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplApplyToClusterListItem instantiates a new ParamTplApplyToClusterListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplApplyToClusterListItem(count int32, name string, needRestart bool, partition string) *ParamTplApplyToClusterListItem {
	this := ParamTplApplyToClusterListItem{}
	this.Count = count
	this.Name = name
	this.NeedRestart = needRestart
	this.Partition = partition
	return &this
}

// NewParamTplApplyToClusterListItemWithDefaults instantiates a new ParamTplApplyToClusterListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplApplyToClusterListItemWithDefaults() *ParamTplApplyToClusterListItem {
	this := ParamTplApplyToClusterListItem{}
	return &this
}

// GetCount returns the Count field value.
func (o *ParamTplApplyToClusterListItem) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ParamTplApplyToClusterListItem) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *ParamTplApplyToClusterListItem) SetCount(v int32) {
	o.Count = v
}

// GetName returns the Name field value.
func (o *ParamTplApplyToClusterListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParamTplApplyToClusterListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParamTplApplyToClusterListItem) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ParamTplApplyToClusterListItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplApplyToClusterListItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ParamTplApplyToClusterListItem) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ParamTplApplyToClusterListItem) SetId(v string) {
	o.Id = &v
}

// GetNeedRestart returns the NeedRestart field value.
func (o *ParamTplApplyToClusterListItem) GetNeedRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NeedRestart
}

// GetNeedRestartOk returns a tuple with the NeedRestart field value
// and a boolean to check if the value has been set.
func (o *ParamTplApplyToClusterListItem) GetNeedRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedRestart, true
}

// SetNeedRestart sets field value.
func (o *ParamTplApplyToClusterListItem) SetNeedRestart(v bool) {
	o.NeedRestart = v
}

// GetPartition returns the Partition field value.
func (o *ParamTplApplyToClusterListItem) GetPartition() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value
// and a boolean to check if the value has been set.
func (o *ParamTplApplyToClusterListItem) GetPartitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partition, true
}

// SetPartition sets field value.
func (o *ParamTplApplyToClusterListItem) SetPartition(v string) {
	o.Partition = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplApplyToClusterListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	toSerialize["name"] = o.Name
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["needRestart"] = o.NeedRestart
	toSerialize["partition"] = o.Partition

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParamTplApplyToClusterListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count       *int32  `json:"count"`
		Name        *string `json:"name"`
		Id          *string `json:"id,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"count", "name", "id", "needRestart", "partition"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.Name = *all.Name
	o.Id = all.Id
	o.NeedRestart = *all.NeedRestart
	o.Partition = *all.Partition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
