// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsObAlterParameter struct {
	// parameter to alter
	Parameter string `json:"parameter"`
	// old value of parameter
	OldValue string `json:"oldValue"`
	// old value of parameter
	NewValue string `json:"newValue"`
	// tenant mode, need if isVariable is true
	Mode *string `json:"mode,omitempty"`
	// tenant name
	TenantName string `json:"tenantName"`
	// whether the param is tenant variable or not
	IsVariable bool `json:"isVariable"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsObAlterParameter instantiates a new DmsObAlterParameter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsObAlterParameter(parameter string, oldValue string, newValue string, tenantName string, isVariable bool) *DmsObAlterParameter {
	this := DmsObAlterParameter{}
	this.Parameter = parameter
	this.OldValue = oldValue
	this.NewValue = newValue
	this.TenantName = tenantName
	this.IsVariable = isVariable
	return &this
}

// NewDmsObAlterParameterWithDefaults instantiates a new DmsObAlterParameter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsObAlterParameterWithDefaults() *DmsObAlterParameter {
	this := DmsObAlterParameter{}
	return &this
}

// GetParameter returns the Parameter field value.
func (o *DmsObAlterParameter) GetParameter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *DmsObAlterParameter) GetParameterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value.
func (o *DmsObAlterParameter) SetParameter(v string) {
	o.Parameter = v
}

// GetOldValue returns the OldValue field value.
func (o *DmsObAlterParameter) GetOldValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value
// and a boolean to check if the value has been set.
func (o *DmsObAlterParameter) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldValue, true
}

// SetOldValue sets field value.
func (o *DmsObAlterParameter) SetOldValue(v string) {
	o.OldValue = v
}

// GetNewValue returns the NewValue field value.
func (o *DmsObAlterParameter) GetNewValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
func (o *DmsObAlterParameter) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// SetNewValue sets field value.
func (o *DmsObAlterParameter) SetNewValue(v string) {
	o.NewValue = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DmsObAlterParameter) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObAlterParameter) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DmsObAlterParameter) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DmsObAlterParameter) SetMode(v string) {
	o.Mode = &v
}

// GetTenantName returns the TenantName field value.
func (o *DmsObAlterParameter) GetTenantName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value
// and a boolean to check if the value has been set.
func (o *DmsObAlterParameter) GetTenantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantName, true
}

// SetTenantName sets field value.
func (o *DmsObAlterParameter) SetTenantName(v string) {
	o.TenantName = v
}

// GetIsVariable returns the IsVariable field value.
func (o *DmsObAlterParameter) GetIsVariable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsVariable
}

// GetIsVariableOk returns a tuple with the IsVariable field value
// and a boolean to check if the value has been set.
func (o *DmsObAlterParameter) GetIsVariableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVariable, true
}

// SetIsVariable sets field value.
func (o *DmsObAlterParameter) SetIsVariable(v bool) {
	o.IsVariable = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsObAlterParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["parameter"] = o.Parameter
	toSerialize["oldValue"] = o.OldValue
	toSerialize["newValue"] = o.NewValue
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["tenantName"] = o.TenantName
	toSerialize["isVariable"] = o.IsVariable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsObAlterParameter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Parameter  *string `json:"parameter"`
		OldValue   *string `json:"oldValue"`
		NewValue   *string `json:"newValue"`
		Mode       *string `json:"mode,omitempty"`
		TenantName *string `json:"tenantName"`
		IsVariable *bool   `json:"isVariable"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Parameter == nil {
		return fmt.Errorf("required field parameter missing")
	}
	if all.OldValue == nil {
		return fmt.Errorf("required field oldValue missing")
	}
	if all.NewValue == nil {
		return fmt.Errorf("required field newValue missing")
	}
	if all.TenantName == nil {
		return fmt.Errorf("required field tenantName missing")
	}
	if all.IsVariable == nil {
		return fmt.Errorf("required field isVariable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"parameter", "oldValue", "newValue", "mode", "tenantName", "isVariable"})
	} else {
		return err
	}
	o.Parameter = *all.Parameter
	o.OldValue = *all.OldValue
	o.NewValue = *all.NewValue
	o.Mode = all.Mode
	o.TenantName = *all.TenantName
	o.IsVariable = *all.IsVariable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
