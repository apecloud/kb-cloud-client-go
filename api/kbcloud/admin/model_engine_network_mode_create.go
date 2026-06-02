// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineNetworkModeCreate struct {
	// engine name
	EngineName string `json:"engineName"`
	// engine mode
	Mode string `json:"mode"`
	// network modes
	NetworkModes []string `json:"networkModes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineNetworkModeCreate instantiates a new EngineNetworkModeCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineNetworkModeCreate(engineName string, mode string, networkModes []string) *EngineNetworkModeCreate {
	this := EngineNetworkModeCreate{}
	this.EngineName = engineName
	this.Mode = mode
	this.NetworkModes = networkModes
	return &this
}

// NewEngineNetworkModeCreateWithDefaults instantiates a new EngineNetworkModeCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineNetworkModeCreateWithDefaults() *EngineNetworkModeCreate {
	this := EngineNetworkModeCreate{}
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *EngineNetworkModeCreate) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkModeCreate) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineNetworkModeCreate) SetEngineName(v string) {
	o.EngineName = v
}

// GetMode returns the Mode field value.
func (o *EngineNetworkModeCreate) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkModeCreate) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *EngineNetworkModeCreate) SetMode(v string) {
	o.Mode = v
}

// GetNetworkModes returns the NetworkModes field value.
func (o *EngineNetworkModeCreate) GetNetworkModes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value
// and a boolean to check if the value has been set.
func (o *EngineNetworkModeCreate) GetNetworkModesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// SetNetworkModes sets field value.
func (o *EngineNetworkModeCreate) SetNetworkModes(v []string) {
	o.NetworkModes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineNetworkModeCreate) MarshalJSON() ([]byte, error) {
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
func (o *EngineNetworkModeCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName   *string   `json:"engineName"`
		Mode         *string   `json:"mode"`
		NetworkModes *[]string `json:"networkModes"`
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
