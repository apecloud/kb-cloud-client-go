// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineNetworkModeOptionsItem struct {
	// engine name
	EngineName string `json:"engineName"`
	// engine mode
	Mode string `json:"mode"`
	// available network modes for this engine+mode
	NetworkModes []NetworkMode `json:"networkModes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineNetworkModeOptionsItem instantiates a new EngineNetworkModeOptionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineNetworkModeOptionsItem(engineName string, mode string, networkModes []NetworkMode) *EngineNetworkModeOptionsItem {
	this := EngineNetworkModeOptionsItem{}
	this.EngineName = engineName
	this.Mode = mode
	this.NetworkModes = networkModes
	return &this
}

// NewEngineNetworkModeOptionsItemWithDefaults instantiates a new EngineNetworkModeOptionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineNetworkModeOptionsItemWithDefaults() *EngineNetworkModeOptionsItem {
	this := EngineNetworkModeOptionsItem{}
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *EngineNetworkModeOptionsItem) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkModeOptionsItem) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineNetworkModeOptionsItem) SetEngineName(v string) {
	o.EngineName = v
}

// GetMode returns the Mode field value.
func (o *EngineNetworkModeOptionsItem) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkModeOptionsItem) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *EngineNetworkModeOptionsItem) SetMode(v string) {
	o.Mode = v
}

// GetNetworkModes returns the NetworkModes field value.
func (o *EngineNetworkModeOptionsItem) GetNetworkModes() []NetworkMode {
	if o == nil {
		var ret []NetworkMode
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkModeOptionsItem) GetNetworkModesOk() (*[]NetworkMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// SetNetworkModes sets field value.
func (o *EngineNetworkModeOptionsItem) SetNetworkModes(v []NetworkMode) {
	o.NetworkModes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineNetworkModeOptionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["mode"] = o.Mode
	toSerialize["networkModes"] = o.NetworkModes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineNetworkModeOptionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName   *string        `json:"engineName"`
		Mode         *string        `json:"mode"`
		NetworkModes *[]NetworkMode `json:"networkModes"`
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
	if all.NetworkModes == nil {
		return fmt.Errorf("required field networkModes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "mode", "networkModes"})
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Mode = *all.Mode
	o.NetworkModes = *all.NetworkModes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
