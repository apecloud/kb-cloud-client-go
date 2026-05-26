// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineVersionDisableServiceVersion EngineVersionRecord disable service version option
type EngineVersionDisableServiceVersion struct {
	// Name of the engine
	EngineName string `json:"engineName"`
	// Version of the engine
	Version string `json:"version"`
	// The service version to disable
	ServiceVersion string `json:"serviceVersion"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineVersionDisableServiceVersion instantiates a new EngineVersionDisableServiceVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineVersionDisableServiceVersion(engineName string, version string, serviceVersion string) *EngineVersionDisableServiceVersion {
	this := EngineVersionDisableServiceVersion{}
	this.EngineName = engineName
	this.Version = version
	this.ServiceVersion = serviceVersion
	return &this
}

// NewEngineVersionDisableServiceVersionWithDefaults instantiates a new EngineVersionDisableServiceVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineVersionDisableServiceVersionWithDefaults() *EngineVersionDisableServiceVersion {
	this := EngineVersionDisableServiceVersion{}
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *EngineVersionDisableServiceVersion) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineVersionDisableServiceVersion) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineVersionDisableServiceVersion) SetEngineName(v string) {
	o.EngineName = v
}

// GetVersion returns the Version field value.
func (o *EngineVersionDisableServiceVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EngineVersionDisableServiceVersion) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *EngineVersionDisableServiceVersion) SetVersion(v string) {
	o.Version = v
}

// GetServiceVersion returns the ServiceVersion field value.
func (o *EngineVersionDisableServiceVersion) GetServiceVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value
// and a boolean to check if the value has been set.
func (o *EngineVersionDisableServiceVersion) GetServiceVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceVersion, true
}

// SetServiceVersion sets field value.
func (o *EngineVersionDisableServiceVersion) SetServiceVersion(v string) {
	o.ServiceVersion = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineVersionDisableServiceVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["version"] = o.Version
	toSerialize["serviceVersion"] = o.ServiceVersion

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineVersionDisableServiceVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName     *string `json:"engineName"`
		Version        *string `json:"version"`
		ServiceVersion *string `json:"serviceVersion"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	if all.ServiceVersion == nil {
		return fmt.Errorf("required field serviceVersion missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "version", "serviceVersion"})
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Version = *all.Version
	o.ServiceVersion = *all.ServiceVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
