// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentEngineOptionCreate struct {
	// engine name
	EngineName string `json:"engineName"`
	// engine mode
	Mode string `json:"mode"`
	// network modes
	NetworkModes []string `json:"networkModes,omitempty"`
	// Number of data volume templates supported by this engine+mode in the environment.
	// At least one configurable field, such as networkModes or dataVolumeCount, must be provided.
	//
	DataVolumeCount common.NullableInt64 `json:"dataVolumeCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentEngineOptionCreate instantiates a new EnvironmentEngineOptionCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentEngineOptionCreate(engineName string, mode string) *EnvironmentEngineOptionCreate {
	this := EnvironmentEngineOptionCreate{}
	this.EngineName = engineName
	this.Mode = mode
	return &this
}

// NewEnvironmentEngineOptionCreateWithDefaults instantiates a new EnvironmentEngineOptionCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentEngineOptionCreateWithDefaults() *EnvironmentEngineOptionCreate {
	this := EnvironmentEngineOptionCreate{}
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *EnvironmentEngineOptionCreate) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOptionCreate) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EnvironmentEngineOptionCreate) SetEngineName(v string) {
	o.EngineName = v
}

// GetMode returns the Mode field value.
func (o *EnvironmentEngineOptionCreate) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOptionCreate) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *EnvironmentEngineOptionCreate) SetMode(v string) {
	o.Mode = v
}

// GetNetworkModes returns the NetworkModes field value if set, zero value otherwise.
func (o *EnvironmentEngineOptionCreate) GetNetworkModes() []string {
	if o == nil || o.NetworkModes == nil {
		var ret []string
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOptionCreate) GetNetworkModesOk() (*[]string, bool) {
	if o == nil || o.NetworkModes == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// HasNetworkModes returns a boolean if a field has been set.
func (o *EnvironmentEngineOptionCreate) HasNetworkModes() bool {
	return o != nil && o.NetworkModes != nil
}

// SetNetworkModes gets a reference to the given []string and assigns it to the NetworkModes field.
func (o *EnvironmentEngineOptionCreate) SetNetworkModes(v []string) {
	o.NetworkModes = v
}

// GetDataVolumeCount returns the DataVolumeCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEngineOptionCreate) GetDataVolumeCount() int64 {
	if o == nil || o.DataVolumeCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DataVolumeCount.Get()
}

// GetDataVolumeCountOk returns a tuple with the DataVolumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentEngineOptionCreate) GetDataVolumeCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataVolumeCount.Get(), o.DataVolumeCount.IsSet()
}

// HasDataVolumeCount returns a boolean if a field has been set.
func (o *EnvironmentEngineOptionCreate) HasDataVolumeCount() bool {
	return o != nil && o.DataVolumeCount.IsSet()
}

// SetDataVolumeCount gets a reference to the given common.NullableInt64 and assigns it to the DataVolumeCount field.
func (o *EnvironmentEngineOptionCreate) SetDataVolumeCount(v int64) {
	o.DataVolumeCount.Set(&v)
}

// SetDataVolumeCountNil sets the value for DataVolumeCount to be an explicit nil.
func (o *EnvironmentEngineOptionCreate) SetDataVolumeCountNil() {
	o.DataVolumeCount.Set(nil)
}

// UnsetDataVolumeCount ensures that no value is present for DataVolumeCount, not even an explicit nil.
func (o *EnvironmentEngineOptionCreate) UnsetDataVolumeCount() {
	o.DataVolumeCount.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentEngineOptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["mode"] = o.Mode
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
func (o *EnvironmentEngineOptionCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName      *string              `json:"engineName"`
		Mode            *string              `json:"mode"`
		NetworkModes    []string             `json:"networkModes,omitempty"`
		DataVolumeCount common.NullableInt64 `json:"dataVolumeCount,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "mode", "networkModes", "dataVolumeCount"})
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Mode = *all.Mode
	o.NetworkModes = all.NetworkModes
	o.DataVolumeCount = all.DataVolumeCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
