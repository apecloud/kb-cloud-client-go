// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineDefinitionVersion struct {
	Version common.NullableString         `json:"version,omitempty"`
	Query   *EngineDefinitionVersionQuery `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineDefinitionVersion instantiates a new EngineDefinitionVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineDefinitionVersion() *EngineDefinitionVersion {
	this := EngineDefinitionVersion{}
	return &this
}

// NewEngineDefinitionVersionWithDefaults instantiates a new EngineDefinitionVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineDefinitionVersionWithDefaults() *EngineDefinitionVersion {
	this := EngineDefinitionVersion{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineDefinitionVersion) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineDefinitionVersion) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *EngineDefinitionVersion) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given common.NullableString and assigns it to the Version field.
func (o *EngineDefinitionVersion) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *EngineDefinitionVersion) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *EngineDefinitionVersion) UnsetVersion() {
	o.Version.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *EngineDefinitionVersion) GetQuery() EngineDefinitionVersionQuery {
	if o == nil || o.Query == nil {
		var ret EngineDefinitionVersionQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinitionVersion) GetQueryOk() (*EngineDefinitionVersionQuery, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *EngineDefinitionVersion) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given EngineDefinitionVersionQuery and assigns it to the Query field.
func (o *EngineDefinitionVersion) SetQuery(v EngineDefinitionVersionQuery) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineDefinitionVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineDefinitionVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Version common.NullableString         `json:"version,omitempty"`
		Query   *EngineDefinitionVersionQuery `json:"query,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"version", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Version = all.Version
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
