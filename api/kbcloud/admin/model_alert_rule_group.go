// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AlertRuleGroup struct {
	Name  *string     `json:"name,omitempty"`
	Rules []AlertRule `json:"rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertRuleGroup instantiates a new AlertRuleGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertRuleGroup() *AlertRuleGroup {
	this := AlertRuleGroup{}
	return &this
}

// NewAlertRuleGroupWithDefaults instantiates a new AlertRuleGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertRuleGroupWithDefaults() *AlertRuleGroup {
	this := AlertRuleGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertRuleGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertRuleGroup) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertRuleGroup) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *AlertRuleGroup) GetRules() []AlertRule {
	if o == nil || o.Rules == nil {
		var ret []AlertRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleGroup) GetRulesOk() (*[]AlertRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *AlertRuleGroup) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []AlertRule and assigns it to the Rules field.
func (o *AlertRuleGroup) SetRules(v []AlertRule) {
	o.Rules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertRuleGroup) MarshalJSON() ([]byte, error) {
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
func (o *AlertRuleGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string     `json:"name,omitempty"`
		Rules []AlertRule `json:"rules,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
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
