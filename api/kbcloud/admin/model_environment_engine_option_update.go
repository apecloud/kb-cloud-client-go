// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentEngineOptionUpdate struct {
	// network modes to update
	NetworkModes []string `json:"networkModes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentEngineOptionUpdate instantiates a new EnvironmentEngineOptionUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentEngineOptionUpdate(networkModes []string) *EnvironmentEngineOptionUpdate {
	this := EnvironmentEngineOptionUpdate{}
	this.NetworkModes = networkModes
	return &this
}

// NewEnvironmentEngineOptionUpdateWithDefaults instantiates a new EnvironmentEngineOptionUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentEngineOptionUpdateWithDefaults() *EnvironmentEngineOptionUpdate {
	this := EnvironmentEngineOptionUpdate{}
	return &this
}

// GetNetworkModes returns the NetworkModes field value.
func (o *EnvironmentEngineOptionUpdate) GetNetworkModes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value
// and a boolean to check if the value has been set.
func (o *EnvironmentEngineOptionUpdate) GetNetworkModesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// SetNetworkModes sets field value.
func (o *EnvironmentEngineOptionUpdate) SetNetworkModes(v []string) {
	o.NetworkModes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentEngineOptionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["networkModes"] = o.NetworkModes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentEngineOptionUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NetworkModes *[]string `json:"networkModes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.NetworkModes == nil {
		return fmt.Errorf("required field networkModes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"networkModes"})
	} else {
		return err
	}
	o.NetworkModes = *all.NetworkModes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
