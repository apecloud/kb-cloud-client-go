// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DolphinDBVolumeUsage struct {
	Node       *string `json:"node,omitempty"`
	VolumePath *string `json:"volumePath,omitempty"`
	Usage      *int64  `json:"usage,omitempty"`
	Capacity   *int64  `json:"capacity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBVolumeUsage instantiates a new DolphinDBVolumeUsage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBVolumeUsage() *DolphinDBVolumeUsage {
	this := DolphinDBVolumeUsage{}
	return &this
}

// NewDolphinDBVolumeUsageWithDefaults instantiates a new DolphinDBVolumeUsage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBVolumeUsageWithDefaults() *DolphinDBVolumeUsage {
	this := DolphinDBVolumeUsage{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *DolphinDBVolumeUsage) GetNode() string {
	if o == nil || o.Node == nil {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBVolumeUsage) GetNodeOk() (*string, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *DolphinDBVolumeUsage) HasNode() bool {
	return o != nil && o.Node != nil
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *DolphinDBVolumeUsage) SetNode(v string) {
	o.Node = &v
}

// GetVolumePath returns the VolumePath field value if set, zero value otherwise.
func (o *DolphinDBVolumeUsage) GetVolumePath() string {
	if o == nil || o.VolumePath == nil {
		var ret string
		return ret
	}
	return *o.VolumePath
}

// GetVolumePathOk returns a tuple with the VolumePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBVolumeUsage) GetVolumePathOk() (*string, bool) {
	if o == nil || o.VolumePath == nil {
		return nil, false
	}
	return o.VolumePath, true
}

// HasVolumePath returns a boolean if a field has been set.
func (o *DolphinDBVolumeUsage) HasVolumePath() bool {
	return o != nil && o.VolumePath != nil
}

// SetVolumePath gets a reference to the given string and assigns it to the VolumePath field.
func (o *DolphinDBVolumeUsage) SetVolumePath(v string) {
	o.VolumePath = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *DolphinDBVolumeUsage) GetUsage() int64 {
	if o == nil || o.Usage == nil {
		var ret int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBVolumeUsage) GetUsageOk() (*int64, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *DolphinDBVolumeUsage) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given int64 and assigns it to the Usage field.
func (o *DolphinDBVolumeUsage) SetUsage(v int64) {
	o.Usage = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *DolphinDBVolumeUsage) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBVolumeUsage) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *DolphinDBVolumeUsage) HasCapacity() bool {
	return o != nil && o.Capacity != nil
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *DolphinDBVolumeUsage) SetCapacity(v int64) {
	o.Capacity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBVolumeUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}
	if o.VolumePath != nil {
		toSerialize["volumePath"] = o.VolumePath
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBVolumeUsage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Node       *string `json:"node,omitempty"`
		VolumePath *string `json:"volumePath,omitempty"`
		Usage      *int64  `json:"usage,omitempty"`
		Capacity   *int64  `json:"capacity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"node", "volumePath", "usage", "capacity"})
	} else {
		return err
	}
	o.Node = all.Node
	o.VolumePath = all.VolumePath
	o.Usage = all.Usage
	o.Capacity = all.Capacity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
