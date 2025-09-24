// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvModuleVersion environment module version
type EnvModuleVersion struct {
	Kubeblocks *string `json:"kubeblocks,omitempty"`
	Gemini     *string `json:"gemini,omitempty"`
	Oteld      *string `json:"oteld,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvModuleVersion instantiates a new EnvModuleVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvModuleVersion() *EnvModuleVersion {
	this := EnvModuleVersion{}
	return &this
}

// NewEnvModuleVersionWithDefaults instantiates a new EnvModuleVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvModuleVersionWithDefaults() *EnvModuleVersion {
	this := EnvModuleVersion{}
	return &this
}

// GetKubeblocks returns the Kubeblocks field value if set, zero value otherwise.
func (o *EnvModuleVersion) GetKubeblocks() string {
	if o == nil || o.Kubeblocks == nil {
		var ret string
		return ret
	}
	return *o.Kubeblocks
}

// GetKubeblocksOk returns a tuple with the Kubeblocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvModuleVersion) GetKubeblocksOk() (*string, bool) {
	if o == nil || o.Kubeblocks == nil {
		return nil, false
	}
	return o.Kubeblocks, true
}

// HasKubeblocks returns a boolean if a field has been set.
func (o *EnvModuleVersion) HasKubeblocks() bool {
	return o != nil && o.Kubeblocks != nil
}

// SetKubeblocks gets a reference to the given string and assigns it to the Kubeblocks field.
func (o *EnvModuleVersion) SetKubeblocks(v string) {
	o.Kubeblocks = &v
}

// GetGemini returns the Gemini field value if set, zero value otherwise.
func (o *EnvModuleVersion) GetGemini() string {
	if o == nil || o.Gemini == nil {
		var ret string
		return ret
	}
	return *o.Gemini
}

// GetGeminiOk returns a tuple with the Gemini field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvModuleVersion) GetGeminiOk() (*string, bool) {
	if o == nil || o.Gemini == nil {
		return nil, false
	}
	return o.Gemini, true
}

// HasGemini returns a boolean if a field has been set.
func (o *EnvModuleVersion) HasGemini() bool {
	return o != nil && o.Gemini != nil
}

// SetGemini gets a reference to the given string and assigns it to the Gemini field.
func (o *EnvModuleVersion) SetGemini(v string) {
	o.Gemini = &v
}

// GetOteld returns the Oteld field value if set, zero value otherwise.
func (o *EnvModuleVersion) GetOteld() string {
	if o == nil || o.Oteld == nil {
		var ret string
		return ret
	}
	return *o.Oteld
}

// GetOteldOk returns a tuple with the Oteld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvModuleVersion) GetOteldOk() (*string, bool) {
	if o == nil || o.Oteld == nil {
		return nil, false
	}
	return o.Oteld, true
}

// HasOteld returns a boolean if a field has been set.
func (o *EnvModuleVersion) HasOteld() bool {
	return o != nil && o.Oteld != nil
}

// SetOteld gets a reference to the given string and assigns it to the Oteld field.
func (o *EnvModuleVersion) SetOteld(v string) {
	o.Oteld = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvModuleVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Kubeblocks != nil {
		toSerialize["kubeblocks"] = o.Kubeblocks
	}
	if o.Gemini != nil {
		toSerialize["gemini"] = o.Gemini
	}
	if o.Oteld != nil {
		toSerialize["oteld"] = o.Oteld
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvModuleVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kubeblocks *string `json:"kubeblocks,omitempty"`
		Gemini     *string `json:"gemini,omitempty"`
		Oteld      *string `json:"oteld,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kubeblocks", "gemini", "oteld"})
	} else {
		return err
	}
	o.Kubeblocks = all.Kubeblocks
	o.Gemini = all.Gemini
	o.Oteld = all.Oteld

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
