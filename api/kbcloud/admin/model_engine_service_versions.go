// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type EngineServiceVersions struct {
	// component type, refer to componentDef and support NamePrefix
	Component *string `json:"component,omitempty"`
	UpgradeableVersions []string `json:"upgradeableVersions,omitempty"`
	Versions []EngineServiceVersionsVersionsItem `json:"versions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewEngineServiceVersions instantiates a new EngineServiceVersions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineServiceVersions() *EngineServiceVersions {
	this := EngineServiceVersions{}
	return &this
}

// NewEngineServiceVersionsWithDefaults instantiates a new EngineServiceVersions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineServiceVersionsWithDefaults() *EngineServiceVersions {
	this := EngineServiceVersions{}
	return &this
}
// GetComponent returns the Component field value if set, zero value otherwise.
func (o *EngineServiceVersions) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersions) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *EngineServiceVersions) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *EngineServiceVersions) SetComponent(v string) {
	o.Component = &v
}


// GetUpgradeableVersions returns the UpgradeableVersions field value if set, zero value otherwise.
func (o *EngineServiceVersions) GetUpgradeableVersions() []string {
	if o == nil || o.UpgradeableVersions == nil {
		var ret []string
		return ret
	}
	return o.UpgradeableVersions
}

// GetUpgradeableVersionsOk returns a tuple with the UpgradeableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersions) GetUpgradeableVersionsOk() (*[]string, bool) {
	if o == nil || o.UpgradeableVersions == nil {
		return nil, false
	}
	return &o.UpgradeableVersions, true
}

// HasUpgradeableVersions returns a boolean if a field has been set.
func (o *EngineServiceVersions) HasUpgradeableVersions() bool {
	return o != nil && o.UpgradeableVersions != nil
}

// SetUpgradeableVersions gets a reference to the given []string and assigns it to the UpgradeableVersions field.
func (o *EngineServiceVersions) SetUpgradeableVersions(v []string) {
	o.UpgradeableVersions = v
}


// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *EngineServiceVersions) GetVersions() []EngineServiceVersionsVersionsItem {
	if o == nil || o.Versions == nil {
		var ret []EngineServiceVersionsVersionsItem
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersions) GetVersionsOk() (*[]EngineServiceVersionsVersionsItem, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *EngineServiceVersions) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []EngineServiceVersionsVersionsItem and assigns it to the Versions field.
func (o *EngineServiceVersions) SetVersions(v []EngineServiceVersionsVersionsItem) {
	o.Versions = v
}



// MarshalJSON serializes the struct using spec logic.
func (o EngineServiceVersions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.UpgradeableVersions != nil {
		toSerialize["upgradeableVersions"] = o.UpgradeableVersions
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineServiceVersions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string `json:"component,omitempty"`
		UpgradeableVersions []string `json:"upgradeableVersions,omitempty"`
		Versions []EngineServiceVersionsVersionsItem `json:"versions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "component", "upgradeableVersions", "versions",  })
	} else {
		return err
	}
	o.Component = all.Component
	o.UpgradeableVersions = all.UpgradeableVersions
	o.Versions = all.Versions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
