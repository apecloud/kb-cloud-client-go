// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AlterRuleRef struct {
	GroupName *string `json:"groupName,omitempty"`
	AlertName *string `json:"alertName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlterRuleRef instantiates a new AlterRuleRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlterRuleRef() *AlterRuleRef {
	this := AlterRuleRef{}
	return &this
}

// NewAlterRuleRefWithDefaults instantiates a new AlterRuleRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlterRuleRefWithDefaults() *AlterRuleRef {
	this := AlterRuleRef{}
	return &this
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *AlterRuleRef) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlterRuleRef) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *AlterRuleRef) HasGroupName() bool {
	return o != nil && o.GroupName != nil
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *AlterRuleRef) SetGroupName(v string) {
	o.GroupName = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *AlterRuleRef) GetAlertName() string {
	if o == nil || o.AlertName == nil {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlterRuleRef) GetAlertNameOk() (*string, bool) {
	if o == nil || o.AlertName == nil {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *AlterRuleRef) HasAlertName() bool {
	return o != nil && o.AlertName != nil
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *AlterRuleRef) SetAlertName(v string) {
	o.AlertName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlterRuleRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.GroupName != nil {
		toSerialize["groupName"] = o.GroupName
	}
	if o.AlertName != nil {
		toSerialize["alertName"] = o.AlertName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlterRuleRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupName *string `json:"groupName,omitempty"`
		AlertName *string `json:"alertName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"groupName", "alertName"})
	} else {
		return err
	}
	o.GroupName = all.GroupName
	o.AlertName = all.AlertName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
