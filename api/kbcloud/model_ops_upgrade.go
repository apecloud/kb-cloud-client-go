// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// OpsUpgrade OpsUpgrade is the payload to upgrade a KubeBlocks cluster 
type OpsUpgrade struct {
	Version string `json:"version"`
	Component *string `json:"component,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewOpsUpgrade instantiates a new OpsUpgrade object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsUpgrade(version string) *OpsUpgrade {
	this := OpsUpgrade{}
	this.Version = version
	return &this
}

// NewOpsUpgradeWithDefaults instantiates a new OpsUpgrade object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsUpgradeWithDefaults() *OpsUpgrade {
	this := OpsUpgrade{}
	return &this
}
// GetVersion returns the Version field value.
func (o *OpsUpgrade) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OpsUpgrade) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *OpsUpgrade) SetVersion(v string) {
	o.Version = v
}


// GetComponent returns the Component field value if set, zero value otherwise.
func (o *OpsUpgrade) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsUpgrade) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *OpsUpgrade) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *OpsUpgrade) SetComponent(v string) {
	o.Component = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o OpsUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["version"] = o.Version
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsUpgrade) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Version *string `json:"version"`
		Component *string `json:"component,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "version", "component",  })
	} else {
		return err
	}
	o.Version = *all.Version
	o.Component = all.Component

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
