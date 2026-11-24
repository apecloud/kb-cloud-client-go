// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ComponentOptionVersionBlockedUpgradeFromItem One upgrade restriction while the cluster runs a service version matching from, upgrading to a service version matching to is not allowed.
type ComponentOptionVersionBlockedUpgradeFromItem struct {
	// Pattern of the service version the cluster currently runs. Supports exact version "7.4.7", version prefix "7.4", and comparisons ">=7.4.5", ">7.4.4", "<=7.4.7", "<7.4.8".
	From string `json:"from"`
	// Pattern of the service version the upgrade is blocked to, with the same syntax as from. Omit to block every upgrade target, which makes from the end of the line for the component.
	To *string `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOptionVersionBlockedUpgradeFromItem instantiates a new ComponentOptionVersionBlockedUpgradeFromItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOptionVersionBlockedUpgradeFromItem(from string) *ComponentOptionVersionBlockedUpgradeFromItem {
	this := ComponentOptionVersionBlockedUpgradeFromItem{}
	this.From = from
	return &this
}

// NewComponentOptionVersionBlockedUpgradeFromItemWithDefaults instantiates a new ComponentOptionVersionBlockedUpgradeFromItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOptionVersionBlockedUpgradeFromItemWithDefaults() *ComponentOptionVersionBlockedUpgradeFromItem {
	this := ComponentOptionVersionBlockedUpgradeFromItem{}
	return &this
}

// GetFrom returns the From field value.
func (o *ComponentOptionVersionBlockedUpgradeFromItem) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionBlockedUpgradeFromItem) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *ComponentOptionVersionBlockedUpgradeFromItem) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ComponentOptionVersionBlockedUpgradeFromItem) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionBlockedUpgradeFromItem) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ComponentOptionVersionBlockedUpgradeFromItem) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *ComponentOptionVersionBlockedUpgradeFromItem) SetTo(v string) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOptionVersionBlockedUpgradeFromItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOptionVersionBlockedUpgradeFromItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From *string `json:"from"`
		To   *string `json:"to,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"from", "to"})
	} else {
		return err
	}
	o.From = *all.From
	o.To = all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
