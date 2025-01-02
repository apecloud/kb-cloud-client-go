// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type Alertadmin_alertRuleGroup struct {
	Name  *string                `json:"name,omitempty"`
	Rules []Alertadmin_alertRule `json:"rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertadmin_alertRuleGroup instantiates a new Alertadmin_alertRuleGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertadmin_alertRuleGroup() *Alertadmin_alertRuleGroup {
	this := Alertadmin_alertRuleGroup{}
	return &this
}

// NewAlertadmin_alertRuleGroupWithDefaults instantiates a new Alertadmin_alertRuleGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertadmin_alertRuleGroupWithDefaults() *Alertadmin_alertRuleGroup {
	this := Alertadmin_alertRuleGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Alertadmin_alertRuleGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertRuleGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Alertadmin_alertRuleGroup) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Alertadmin_alertRuleGroup) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *Alertadmin_alertRuleGroup) GetRules() []Alertadmin_alertRule {
	if o == nil || o.Rules == nil {
		var ret []Alertadmin_alertRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alertadmin_alertRuleGroup) GetRulesOk() (*[]Alertadmin_alertRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *Alertadmin_alertRuleGroup) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []Alertadmin_alertRule and assigns it to the Rules field.
func (o *Alertadmin_alertRuleGroup) SetRules(v []Alertadmin_alertRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Alertadmin_alertRuleGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Alertadmin_alertRuleGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string                `json:"name,omitempty"`
		Rules []Alertadmin_alertRule `json:"rules,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "rules"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Rules = all.Rules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
