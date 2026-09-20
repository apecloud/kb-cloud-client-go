// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeComponentInstanceTemplate Declares instance-template support for this mode component.
// Create requires instanceTemplates with name and replicas.
// HScale replica changes require instanceTemplates with name and replicas;
// shard-count changes use shards; they may be set together with
// instanceTemplates (KB ops fill both).
// VScale requires instanceTemplates with name and classCode;
// only those names are updated.
// The platform does not even-split or fill in missing templates.
// Instance online/offline can be set together with instanceTemplates.
// When absent, the component uses component-level operations.
type ModeComponentInstanceTemplate struct {
	// Declared instance template names. Request payloads may only use
	// these names. Names must be unique.
	//
	Names []string `json:"names"`
	// Operations supported via instance templates.
	// Valid values: "vscale" (vertical scaling), "hscale" (horizontal scaling).
	//
	Ops []InstanceTemplateOp `json:"ops"`
	// Overlay fields the frontend should render on each instance template.
	// Omitted or empty means only name and replicas. Request fields not
	// listed here are rejected.
	//
	Fields []InstanceTemplateOverlayField `json:"fields,omitempty"`
	// Env vars the frontend may render when fields includes env.
	// Each item is one fillable env name. Request env names must be in
	// this list.
	//
	Env []InstanceTemplateEnv `json:"env,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeComponentInstanceTemplate instantiates a new ModeComponentInstanceTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeComponentInstanceTemplate(names []string, ops []InstanceTemplateOp) *ModeComponentInstanceTemplate {
	this := ModeComponentInstanceTemplate{}
	this.Names = names
	this.Ops = ops
	return &this
}

// NewModeComponentInstanceTemplateWithDefaults instantiates a new ModeComponentInstanceTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeComponentInstanceTemplateWithDefaults() *ModeComponentInstanceTemplate {
	this := ModeComponentInstanceTemplate{}
	return &this
}

// GetNames returns the Names field value.
func (o *ModeComponentInstanceTemplate) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *ModeComponentInstanceTemplate) GetNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value.
func (o *ModeComponentInstanceTemplate) SetNames(v []string) {
	o.Names = v
}

// GetOps returns the Ops field value.
func (o *ModeComponentInstanceTemplate) GetOps() []InstanceTemplateOp {
	if o == nil {
		var ret []InstanceTemplateOp
		return ret
	}
	return o.Ops
}

// GetOpsOk returns a tuple with the Ops field value
// and a boolean to check if the value has been set.
func (o *ModeComponentInstanceTemplate) GetOpsOk() (*[]InstanceTemplateOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ops, true
}

// SetOps sets field value.
func (o *ModeComponentInstanceTemplate) SetOps(v []InstanceTemplateOp) {
	o.Ops = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ModeComponentInstanceTemplate) GetFields() []InstanceTemplateOverlayField {
	if o == nil || o.Fields == nil {
		var ret []InstanceTemplateOverlayField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeComponentInstanceTemplate) GetFieldsOk() (*[]InstanceTemplateOverlayField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ModeComponentInstanceTemplate) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given []InstanceTemplateOverlayField and assigns it to the Fields field.
func (o *ModeComponentInstanceTemplate) SetFields(v []InstanceTemplateOverlayField) {
	o.Fields = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ModeComponentInstanceTemplate) GetEnv() []InstanceTemplateEnv {
	if o == nil || o.Env == nil {
		var ret []InstanceTemplateEnv
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeComponentInstanceTemplate) GetEnvOk() (*[]InstanceTemplateEnv, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return &o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ModeComponentInstanceTemplate) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given []InstanceTemplateEnv and assigns it to the Env field.
func (o *ModeComponentInstanceTemplate) SetEnv(v []InstanceTemplateEnv) {
	o.Env = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeComponentInstanceTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["names"] = o.Names
	toSerialize["ops"] = o.Ops
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeComponentInstanceTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Names  *[]string                      `json:"names"`
		Ops    *[]InstanceTemplateOp          `json:"ops"`
		Fields []InstanceTemplateOverlayField `json:"fields,omitempty"`
		Env    []InstanceTemplateEnv          `json:"env,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Names == nil {
		return fmt.Errorf("required field names missing")
	}
	if all.Ops == nil {
		return fmt.Errorf("required field ops missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"names", "ops", "fields", "env"})
	} else {
		return err
	}
	o.Names = *all.Names
	o.Ops = *all.Ops
	o.Fields = all.Fields
	o.Env = all.Env

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
