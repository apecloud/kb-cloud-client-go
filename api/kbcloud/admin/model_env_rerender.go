// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// EnvRerender Rule telling whether the env of an existing component is re-rendered when the mode transitions.
type EnvRerender struct {
	// Re-render the whole env of the component.
	AllEnv *bool `json:"allEnv,omitempty"`
	// Re-render only these env vars, when allEnv is not set.
	IncludeEnvs []string `json:"includeEnvs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvRerender instantiates a new EnvRerender object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvRerender() *EnvRerender {
	this := EnvRerender{}
	return &this
}

// NewEnvRerenderWithDefaults instantiates a new EnvRerender object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvRerenderWithDefaults() *EnvRerender {
	this := EnvRerender{}
	return &this
}

// GetAllEnv returns the AllEnv field value if set, zero value otherwise.
func (o *EnvRerender) GetAllEnv() bool {
	if o == nil || o.AllEnv == nil {
		var ret bool
		return ret
	}
	return *o.AllEnv
}

// GetAllEnvOk returns a tuple with the AllEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvRerender) GetAllEnvOk() (*bool, bool) {
	if o == nil || o.AllEnv == nil {
		return nil, false
	}
	return o.AllEnv, true
}

// HasAllEnv returns a boolean if a field has been set.
func (o *EnvRerender) HasAllEnv() bool {
	return o != nil && o.AllEnv != nil
}

// SetAllEnv gets a reference to the given bool and assigns it to the AllEnv field.
func (o *EnvRerender) SetAllEnv(v bool) {
	o.AllEnv = &v
}

// GetIncludeEnvs returns the IncludeEnvs field value if set, zero value otherwise.
func (o *EnvRerender) GetIncludeEnvs() []string {
	if o == nil || o.IncludeEnvs == nil {
		var ret []string
		return ret
	}
	return o.IncludeEnvs
}

// GetIncludeEnvsOk returns a tuple with the IncludeEnvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvRerender) GetIncludeEnvsOk() (*[]string, bool) {
	if o == nil || o.IncludeEnvs == nil {
		return nil, false
	}
	return &o.IncludeEnvs, true
}

// HasIncludeEnvs returns a boolean if a field has been set.
func (o *EnvRerender) HasIncludeEnvs() bool {
	return o != nil && o.IncludeEnvs != nil
}

// SetIncludeEnvs gets a reference to the given []string and assigns it to the IncludeEnvs field.
func (o *EnvRerender) SetIncludeEnvs(v []string) {
	o.IncludeEnvs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvRerender) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.AllEnv != nil {
		toSerialize["allEnv"] = o.AllEnv
	}
	if o.IncludeEnvs != nil {
		toSerialize["includeEnvs"] = o.IncludeEnvs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvRerender) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllEnv      *bool    `json:"allEnv,omitempty"`
		IncludeEnvs []string `json:"includeEnvs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"allEnv", "includeEnvs"})
	} else {
		return err
	}
	o.AllEnv = all.AllEnv
	o.IncludeEnvs = all.IncludeEnvs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
