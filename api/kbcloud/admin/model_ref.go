// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Ref struct {
	// name of config spec, included in configSpecs
	ConfigSpecName string `json:"configSpecName"`
	// name of the referenced configuration template file name, all files in parameter template fold under stores
	TemplateRef string `json:"templateRef"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRef instantiates a new Ref object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRef(configSpecName string, templateRef string) *Ref {
	this := Ref{}
	this.ConfigSpecName = configSpecName
	this.TemplateRef = templateRef
	return &this
}

// NewRefWithDefaults instantiates a new Ref object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRefWithDefaults() *Ref {
	this := Ref{}
	return &this
}

// GetConfigSpecName returns the ConfigSpecName field value.
func (o *Ref) GetConfigSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpecName
}

// GetConfigSpecNameOk returns a tuple with the ConfigSpecName field value
// and a boolean to check if the value has been set.
func (o *Ref) GetConfigSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecName, true
}

// SetConfigSpecName sets field value.
func (o *Ref) SetConfigSpecName(v string) {
	o.ConfigSpecName = v
}

// GetTemplateRef returns the TemplateRef field value.
func (o *Ref) GetTemplateRef() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TemplateRef
}

// GetTemplateRefOk returns a tuple with the TemplateRef field value
// and a boolean to check if the value has been set.
func (o *Ref) GetTemplateRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateRef, true
}

// SetTemplateRef sets field value.
func (o *Ref) SetTemplateRef(v string) {
	o.TemplateRef = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Ref) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["configSpecName"] = o.ConfigSpecName
	toSerialize["templateRef"] = o.TemplateRef

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Ref) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigSpecName *string `json:"configSpecName"`
		TemplateRef    *string `json:"templateRef"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConfigSpecName == nil {
		return fmt.Errorf("required field configSpecName missing")
	}
	if all.TemplateRef == nil {
		return fmt.Errorf("required field templateRef missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"configSpecName", "templateRef"})
	} else {
		return err
	}
	o.ConfigSpecName = *all.ConfigSpecName
	o.TemplateRef = *all.TemplateRef

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
