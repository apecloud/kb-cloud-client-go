// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Register Register manifest for environment
type Register struct {
	// Configuration for privisoining environment resources
	Type EnvironmentRegisterType `json:"type"`
	// The base64 encoded kubeconfig file contents to connect to this environment
	Kubeconfig *string `json:"kubeconfig,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRegister instantiates a new Register object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRegister(typeVar EnvironmentRegisterType) *Register {
	this := Register{}
	this.Type = typeVar
	return &this
}

// NewRegisterWithDefaults instantiates a new Register object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRegisterWithDefaults() *Register {
	this := Register{}
	return &this
}

// GetType returns the Type field value.
func (o *Register) GetType() EnvironmentRegisterType {
	if o == nil {
		var ret EnvironmentRegisterType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Register) GetTypeOk() (*EnvironmentRegisterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Register) SetType(v EnvironmentRegisterType) {
	o.Type = v
}

// GetKubeconfig returns the Kubeconfig field value if set, zero value otherwise.
func (o *Register) GetKubeconfig() string {
	if o == nil || o.Kubeconfig == nil {
		var ret string
		return ret
	}
	return *o.Kubeconfig
}

// GetKubeconfigOk returns a tuple with the Kubeconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Register) GetKubeconfigOk() (*string, bool) {
	if o == nil || o.Kubeconfig == nil {
		return nil, false
	}
	return o.Kubeconfig, true
}

// HasKubeconfig returns a boolean if a field has been set.
func (o *Register) HasKubeconfig() bool {
	return o != nil && o.Kubeconfig != nil
}

// SetKubeconfig gets a reference to the given string and assigns it to the Kubeconfig field.
func (o *Register) SetKubeconfig(v string) {
	o.Kubeconfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Register) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	if o.Kubeconfig != nil {
		toSerialize["kubeconfig"] = o.Kubeconfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Register) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type       *EnvironmentRegisterType `json:"type"`
		Kubeconfig *string                  `json:"kubeconfig,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "kubeconfig"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Kubeconfig = all.Kubeconfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
