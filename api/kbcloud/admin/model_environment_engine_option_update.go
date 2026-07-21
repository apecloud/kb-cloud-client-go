// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentEngineOptionUpdate struct {
	// environment engine option to update
	NetworkModes []string `json:"networkModes,omitempty"`
	// Number of data volume templates supported by this engine+mode in the environment.
	// At least one configurable field, such as networkModes or dataVolumeCount, must be provided.
	//
	DataVolumeCount common.NullableInt64 `json:"dataVolumeCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentEngineOptionUpdate instantiates a new EnvironmentEngineOptionUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentEngineOptionUpdate() *EnvironmentEngineOptionUpdate {
	this := EnvironmentEngineOptionUpdate{}
	return &this
}

// NewEnvironmentEngineOptionUpdateWithDefaults instantiates a new EnvironmentEngineOptionUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentEngineOptionUpdateWithDefaults() *EnvironmentEngineOptionUpdate {
	this := EnvironmentEngineOptionUpdate{}
	return &this
}

// GetNetworkModes returns the NetworkModes field value if set, zero value otherwise.
func (o *EnvironmentEngineOptionUpdate) GetNetworkModes() []string {
	if o == nil || o.NetworkModes == nil {
		var ret []string
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOptionUpdate) GetNetworkModesOk() (*[]string, bool) {
	if o == nil || o.NetworkModes == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// HasNetworkModes returns a boolean if a field has been set.
func (o *EnvironmentEngineOptionUpdate) HasNetworkModes() bool {
	return o != nil && o.NetworkModes != nil
}

// SetNetworkModes gets a reference to the given []string and assigns it to the NetworkModes field.
func (o *EnvironmentEngineOptionUpdate) SetNetworkModes(v []string) {
	o.NetworkModes = v
}

// GetDataVolumeCount returns the DataVolumeCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEngineOptionUpdate) GetDataVolumeCount() int64 {
	if o == nil || o.DataVolumeCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DataVolumeCount.Get()
}

// GetDataVolumeCountOk returns a tuple with the DataVolumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentEngineOptionUpdate) GetDataVolumeCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataVolumeCount.Get(), o.DataVolumeCount.IsSet()
}

// HasDataVolumeCount returns a boolean if a field has been set.
func (o *EnvironmentEngineOptionUpdate) HasDataVolumeCount() bool {
	return o != nil && o.DataVolumeCount.IsSet()
}

// SetDataVolumeCount gets a reference to the given common.NullableInt64 and assigns it to the DataVolumeCount field.
func (o *EnvironmentEngineOptionUpdate) SetDataVolumeCount(v int64) {
	o.DataVolumeCount.Set(&v)
}

// SetDataVolumeCountNil sets the value for DataVolumeCount to be an explicit nil.
func (o *EnvironmentEngineOptionUpdate) SetDataVolumeCountNil() {
	o.DataVolumeCount.Set(nil)
}

// UnsetDataVolumeCount ensures that no value is present for DataVolumeCount, not even an explicit nil.
func (o *EnvironmentEngineOptionUpdate) UnsetDataVolumeCount() {
	o.DataVolumeCount.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentEngineOptionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NetworkModes != nil {
		toSerialize["networkModes"] = o.NetworkModes
	}
	if o.DataVolumeCount.IsSet() {
		toSerialize["dataVolumeCount"] = o.DataVolumeCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentEngineOptionUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NetworkModes    []string             `json:"networkModes,omitempty"`
		DataVolumeCount common.NullableInt64 `json:"dataVolumeCount,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"networkModes", "dataVolumeCount"})
	} else {
		return err
	}
	o.NetworkModes = all.NetworkModes
	o.DataVolumeCount = all.DataVolumeCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
