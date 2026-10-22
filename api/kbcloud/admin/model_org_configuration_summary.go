// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OrgConfigurationSummary struct {
	Type         string                 `json:"type"`
	Count        int64                  `json:"count"`
	Items        []OrgConfigurationItem `json:"items"`
	Truncated    bool                   `json:"truncated"`
	ActionEffect string                 `json:"actionEffect"`
	PurgeEffect  string                 `json:"purgeEffect"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConfigurationSummary instantiates a new OrgConfigurationSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConfigurationSummary(typeVar string, count int64, items []OrgConfigurationItem, truncated bool, actionEffect string, purgeEffect string) *OrgConfigurationSummary {
	this := OrgConfigurationSummary{}
	this.Type = typeVar
	this.Count = count
	this.Items = items
	this.Truncated = truncated
	this.ActionEffect = actionEffect
	this.PurgeEffect = purgeEffect
	return &this
}

// NewOrgConfigurationSummaryWithDefaults instantiates a new OrgConfigurationSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConfigurationSummaryWithDefaults() *OrgConfigurationSummary {
	this := OrgConfigurationSummary{}
	return &this
}

// GetType returns the Type field value.
func (o *OrgConfigurationSummary) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgConfigurationSummary) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OrgConfigurationSummary) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value.
func (o *OrgConfigurationSummary) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *OrgConfigurationSummary) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *OrgConfigurationSummary) SetCount(v int64) {
	o.Count = v
}

// GetItems returns the Items field value.
func (o *OrgConfigurationSummary) GetItems() []OrgConfigurationItem {
	if o == nil {
		var ret []OrgConfigurationItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OrgConfigurationSummary) GetItemsOk() (*[]OrgConfigurationItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *OrgConfigurationSummary) SetItems(v []OrgConfigurationItem) {
	o.Items = v
}

// GetTruncated returns the Truncated field value.
func (o *OrgConfigurationSummary) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *OrgConfigurationSummary) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *OrgConfigurationSummary) SetTruncated(v bool) {
	o.Truncated = v
}

// GetActionEffect returns the ActionEffect field value.
func (o *OrgConfigurationSummary) GetActionEffect() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionEffect
}

// GetActionEffectOk returns a tuple with the ActionEffect field value
// and a boolean to check if the value has been set.
func (o *OrgConfigurationSummary) GetActionEffectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionEffect, true
}

// SetActionEffect sets field value.
func (o *OrgConfigurationSummary) SetActionEffect(v string) {
	o.ActionEffect = v
}

// GetPurgeEffect returns the PurgeEffect field value.
func (o *OrgConfigurationSummary) GetPurgeEffect() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PurgeEffect
}

// GetPurgeEffectOk returns a tuple with the PurgeEffect field value
// and a boolean to check if the value has been set.
func (o *OrgConfigurationSummary) GetPurgeEffectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurgeEffect, true
}

// SetPurgeEffect sets field value.
func (o *OrgConfigurationSummary) SetPurgeEffect(v string) {
	o.PurgeEffect = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConfigurationSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["count"] = o.Count
	toSerialize["items"] = o.Items
	toSerialize["truncated"] = o.Truncated
	toSerialize["actionEffect"] = o.ActionEffect
	toSerialize["purgeEffect"] = o.PurgeEffect

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgConfigurationSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type         *string                 `json:"type"`
		Count        *int64                  `json:"count"`
		Items        *[]OrgConfigurationItem `json:"items"`
		Truncated    *bool                   `json:"truncated"`
		ActionEffect *string                 `json:"actionEffect"`
		PurgeEffect  *string                 `json:"purgeEffect"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Truncated == nil {
		return fmt.Errorf("required field truncated missing")
	}
	if all.ActionEffect == nil {
		return fmt.Errorf("required field actionEffect missing")
	}
	if all.PurgeEffect == nil {
		return fmt.Errorf("required field purgeEffect missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "count", "items", "truncated", "actionEffect", "purgeEffect"})
	} else {
		return err
	}
	o.Type = *all.Type
	o.Count = *all.Count
	o.Items = *all.Items
	o.Truncated = *all.Truncated
	o.ActionEffect = *all.ActionEffect
	o.PurgeEffect = *all.PurgeEffect

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
