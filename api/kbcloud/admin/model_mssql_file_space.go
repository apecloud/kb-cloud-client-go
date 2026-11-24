// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlFileSpace struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	SizeBytes int64  `json:"sizeBytes"`
	UsedBytes *int64 `json:"usedBytes,omitempty"`
	// -1 means grow until the volume is full; zero disables growth.
	MaxBytes int64 `json:"maxBytes"`
	// Percent when percentGrowth is true; otherwise number of 8 KiB pages. Zero disables autogrowth.
	Growth        int64 `json:"growth"`
	PercentGrowth bool  `json:"percentGrowth"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlFileSpace instantiates a new MssqlFileSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlFileSpace(name string, typeVar string, sizeBytes int64, maxBytes int64, growth int64, percentGrowth bool) *MssqlFileSpace {
	this := MssqlFileSpace{}
	this.Name = name
	this.Type = typeVar
	this.SizeBytes = sizeBytes
	this.MaxBytes = maxBytes
	this.Growth = growth
	this.PercentGrowth = percentGrowth
	return &this
}

// NewMssqlFileSpaceWithDefaults instantiates a new MssqlFileSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlFileSpaceWithDefaults() *MssqlFileSpace {
	this := MssqlFileSpace{}
	return &this
}

// GetName returns the Name field value.
func (o *MssqlFileSpace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MssqlFileSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MssqlFileSpace) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *MssqlFileSpace) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MssqlFileSpace) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *MssqlFileSpace) SetType(v string) {
	o.Type = v
}

// GetSizeBytes returns the SizeBytes field value.
func (o *MssqlFileSpace) GetSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlFileSpace) GetSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeBytes, true
}

// SetSizeBytes sets field value.
func (o *MssqlFileSpace) SetSizeBytes(v int64) {
	o.SizeBytes = v
}

// GetUsedBytes returns the UsedBytes field value if set, zero value otherwise.
func (o *MssqlFileSpace) GetUsedBytes() int64 {
	if o == nil || o.UsedBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsedBytes
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MssqlFileSpace) GetUsedBytesOk() (*int64, bool) {
	if o == nil || o.UsedBytes == nil {
		return nil, false
	}
	return o.UsedBytes, true
}

// HasUsedBytes returns a boolean if a field has been set.
func (o *MssqlFileSpace) HasUsedBytes() bool {
	return o != nil && o.UsedBytes != nil
}

// SetUsedBytes gets a reference to the given int64 and assigns it to the UsedBytes field.
func (o *MssqlFileSpace) SetUsedBytes(v int64) {
	o.UsedBytes = &v
}

// GetMaxBytes returns the MaxBytes field value.
func (o *MssqlFileSpace) GetMaxBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxBytes
}

// GetMaxBytesOk returns a tuple with the MaxBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlFileSpace) GetMaxBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxBytes, true
}

// SetMaxBytes sets field value.
func (o *MssqlFileSpace) SetMaxBytes(v int64) {
	o.MaxBytes = v
}

// GetGrowth returns the Growth field value.
func (o *MssqlFileSpace) GetGrowth() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Growth
}

// GetGrowthOk returns a tuple with the Growth field value
// and a boolean to check if the value has been set.
func (o *MssqlFileSpace) GetGrowthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Growth, true
}

// SetGrowth sets field value.
func (o *MssqlFileSpace) SetGrowth(v int64) {
	o.Growth = v
}

// GetPercentGrowth returns the PercentGrowth field value.
func (o *MssqlFileSpace) GetPercentGrowth() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PercentGrowth
}

// GetPercentGrowthOk returns a tuple with the PercentGrowth field value
// and a boolean to check if the value has been set.
func (o *MssqlFileSpace) GetPercentGrowthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentGrowth, true
}

// SetPercentGrowth sets field value.
func (o *MssqlFileSpace) SetPercentGrowth(v bool) {
	o.PercentGrowth = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlFileSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["sizeBytes"] = o.SizeBytes
	if o.UsedBytes != nil {
		toSerialize["usedBytes"] = o.UsedBytes
	}
	toSerialize["maxBytes"] = o.MaxBytes
	toSerialize["growth"] = o.Growth
	toSerialize["percentGrowth"] = o.PercentGrowth

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlFileSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name          *string `json:"name"`
		Type          *string `json:"type"`
		SizeBytes     *int64  `json:"sizeBytes"`
		UsedBytes     *int64  `json:"usedBytes,omitempty"`
		MaxBytes      *int64  `json:"maxBytes"`
		Growth        *int64  `json:"growth"`
		PercentGrowth *bool   `json:"percentGrowth"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.SizeBytes == nil {
		return fmt.Errorf("required field sizeBytes missing")
	}
	if all.MaxBytes == nil {
		return fmt.Errorf("required field maxBytes missing")
	}
	if all.Growth == nil {
		return fmt.Errorf("required field growth missing")
	}
	if all.PercentGrowth == nil {
		return fmt.Errorf("required field percentGrowth missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "sizeBytes", "usedBytes", "maxBytes", "growth", "percentGrowth"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Type = *all.Type
	o.SizeBytes = *all.SizeBytes
	o.UsedBytes = all.UsedBytes
	o.MaxBytes = *all.MaxBytes
	o.Growth = *all.Growth
	o.PercentGrowth = *all.PercentGrowth

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
