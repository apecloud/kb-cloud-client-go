// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

type ServerConfiguration struct {
	// Server architecture type, can be empty.
	ArchitectureType *string `json:"architectureType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServerConfiguration instantiates a new ServerConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServerConfiguration() *ServerConfiguration {
	this := ServerConfiguration{}
	return &this
}

// NewServerConfigurationWithDefaults instantiates a new ServerConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServerConfigurationWithDefaults() *ServerConfiguration {
	this := ServerConfiguration{}
	return &this
}

// GetArchitectureType returns the ArchitectureType field value if set, zero value otherwise.
func (o *ServerConfiguration) GetArchitectureType() string {
	if o == nil || o.ArchitectureType == nil {
		var ret string
		return ret
	}
	return *o.ArchitectureType
}

// GetArchitectureTypeOk returns a tuple with the ArchitectureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfiguration) GetArchitectureTypeOk() (*string, bool) {
	if o == nil || o.ArchitectureType == nil {
		return nil, false
	}
	return o.ArchitectureType, true
}

// HasArchitectureType returns a boolean if a field has been set.
func (o *ServerConfiguration) HasArchitectureType() bool {
	return o != nil && o.ArchitectureType != nil
}

// SetArchitectureType gets a reference to the given string and assigns it to the ArchitectureType field.
func (o *ServerConfiguration) SetArchitectureType(v string) {
	o.ArchitectureType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServerConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ArchitectureType != nil {
		toSerialize["architectureType"] = o.ArchitectureType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServerConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArchitectureType *string `json:"architectureType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"architectureType"})
	} else {
		return err
	}
	o.ArchitectureType = all.ArchitectureType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
