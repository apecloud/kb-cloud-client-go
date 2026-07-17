// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengTablespaceItem struct {
	// Tablespace name.
	Name string `json:"name"`
	// Total size in megabytes.
	TotalMb float64 `json:"totalMb"`
	// Free size in megabytes.
	FreeMb float64 `json:"freeMb"`
	// Used size in megabytes.
	UsedMb float64 `json:"usedMb"`
	// Usage ratio as percentage (0-100).
	UsageRatio float64 `json:"usageRatio"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengTablespaceItem instantiates a new DamengTablespaceItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengTablespaceItem(name string, totalMb float64, freeMb float64, usedMb float64, usageRatio float64) *DamengTablespaceItem {
	this := DamengTablespaceItem{}
	this.Name = name
	this.TotalMb = totalMb
	this.FreeMb = freeMb
	this.UsedMb = usedMb
	this.UsageRatio = usageRatio
	return &this
}

// NewDamengTablespaceItemWithDefaults instantiates a new DamengTablespaceItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengTablespaceItemWithDefaults() *DamengTablespaceItem {
	this := DamengTablespaceItem{}
	return &this
}

// GetName returns the Name field value.
func (o *DamengTablespaceItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DamengTablespaceItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DamengTablespaceItem) SetName(v string) {
	o.Name = v
}

// GetTotalMb returns the TotalMb field value.
func (o *DamengTablespaceItem) GetTotalMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TotalMb
}

// GetTotalMbOk returns a tuple with the TotalMb field value
// and a boolean to check if the value has been set.
func (o *DamengTablespaceItem) GetTotalMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMb, true
}

// SetTotalMb sets field value.
func (o *DamengTablespaceItem) SetTotalMb(v float64) {
	o.TotalMb = v
}

// GetFreeMb returns the FreeMb field value.
func (o *DamengTablespaceItem) GetFreeMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.FreeMb
}

// GetFreeMbOk returns a tuple with the FreeMb field value
// and a boolean to check if the value has been set.
func (o *DamengTablespaceItem) GetFreeMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeMb, true
}

// SetFreeMb sets field value.
func (o *DamengTablespaceItem) SetFreeMb(v float64) {
	o.FreeMb = v
}

// GetUsedMb returns the UsedMb field value.
func (o *DamengTablespaceItem) GetUsedMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.UsedMb
}

// GetUsedMbOk returns a tuple with the UsedMb field value
// and a boolean to check if the value has been set.
func (o *DamengTablespaceItem) GetUsedMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedMb, true
}

// SetUsedMb sets field value.
func (o *DamengTablespaceItem) SetUsedMb(v float64) {
	o.UsedMb = v
}

// GetUsageRatio returns the UsageRatio field value.
func (o *DamengTablespaceItem) GetUsageRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.UsageRatio
}

// GetUsageRatioOk returns a tuple with the UsageRatio field value
// and a boolean to check if the value has been set.
func (o *DamengTablespaceItem) GetUsageRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageRatio, true
}

// SetUsageRatio sets field value.
func (o *DamengTablespaceItem) SetUsageRatio(v float64) {
	o.UsageRatio = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengTablespaceItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["totalMb"] = o.TotalMb
	toSerialize["freeMb"] = o.FreeMb
	toSerialize["usedMb"] = o.UsedMb
	toSerialize["usageRatio"] = o.UsageRatio

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengTablespaceItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string  `json:"name"`
		TotalMb    *float64 `json:"totalMb"`
		FreeMb     *float64 `json:"freeMb"`
		UsedMb     *float64 `json:"usedMb"`
		UsageRatio *float64 `json:"usageRatio"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TotalMb == nil {
		return fmt.Errorf("required field totalMb missing")
	}
	if all.FreeMb == nil {
		return fmt.Errorf("required field freeMb missing")
	}
	if all.UsedMb == nil {
		return fmt.Errorf("required field usedMb missing")
	}
	if all.UsageRatio == nil {
		return fmt.Errorf("required field usageRatio missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "totalMb", "freeMb", "usedMb", "usageRatio"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.TotalMb = *all.TotalMb
	o.FreeMb = *all.FreeMb
	o.UsedMb = *all.UsedMb
	o.UsageRatio = *all.UsageRatio

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
