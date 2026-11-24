// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OracleTablespaceFile struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	Status          string  `json:"status"`
	AllocatedSizeMb float64 `json:"allocatedSizeMB"`
	UsedSizeMb      float64 `json:"usedSizeMB"`
	AutoExtend      bool    `json:"autoExtend"`
	NextSizeMb      float64 `json:"nextSizeMB"`
	MaxSizeMb       float64 `json:"maxSizeMB"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOracleTablespaceFile instantiates a new OracleTablespaceFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOracleTablespaceFile(id int64, name string, status string, allocatedSizeMb float64, usedSizeMb float64, autoExtend bool, nextSizeMb float64, maxSizeMb float64) *OracleTablespaceFile {
	this := OracleTablespaceFile{}
	this.Id = id
	this.Name = name
	this.Status = status
	this.AllocatedSizeMb = allocatedSizeMb
	this.UsedSizeMb = usedSizeMb
	this.AutoExtend = autoExtend
	this.NextSizeMb = nextSizeMb
	this.MaxSizeMb = maxSizeMb
	return &this
}

// NewOracleTablespaceFileWithDefaults instantiates a new OracleTablespaceFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOracleTablespaceFileWithDefaults() *OracleTablespaceFile {
	this := OracleTablespaceFile{}
	return &this
}

// GetId returns the Id field value.
func (o *OracleTablespaceFile) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFile) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OracleTablespaceFile) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *OracleTablespaceFile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OracleTablespaceFile) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value.
func (o *OracleTablespaceFile) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFile) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OracleTablespaceFile) SetStatus(v string) {
	o.Status = v
}

// GetAllocatedSizeMb returns the AllocatedSizeMb field value.
func (o *OracleTablespaceFile) GetAllocatedSizeMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AllocatedSizeMb
}

// GetAllocatedSizeMbOk returns a tuple with the AllocatedSizeMb field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFile) GetAllocatedSizeMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllocatedSizeMb, true
}

// SetAllocatedSizeMb sets field value.
func (o *OracleTablespaceFile) SetAllocatedSizeMb(v float64) {
	o.AllocatedSizeMb = v
}

// GetUsedSizeMb returns the UsedSizeMb field value.
func (o *OracleTablespaceFile) GetUsedSizeMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.UsedSizeMb
}

// GetUsedSizeMbOk returns a tuple with the UsedSizeMb field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFile) GetUsedSizeMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedSizeMb, true
}

// SetUsedSizeMb sets field value.
func (o *OracleTablespaceFile) SetUsedSizeMb(v float64) {
	o.UsedSizeMb = v
}

// GetAutoExtend returns the AutoExtend field value.
func (o *OracleTablespaceFile) GetAutoExtend() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AutoExtend
}

// GetAutoExtendOk returns a tuple with the AutoExtend field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFile) GetAutoExtendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoExtend, true
}

// SetAutoExtend sets field value.
func (o *OracleTablespaceFile) SetAutoExtend(v bool) {
	o.AutoExtend = v
}

// GetNextSizeMb returns the NextSizeMb field value.
func (o *OracleTablespaceFile) GetNextSizeMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.NextSizeMb
}

// GetNextSizeMbOk returns a tuple with the NextSizeMb field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFile) GetNextSizeMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextSizeMb, true
}

// SetNextSizeMb sets field value.
func (o *OracleTablespaceFile) SetNextSizeMb(v float64) {
	o.NextSizeMb = v
}

// GetMaxSizeMb returns the MaxSizeMb field value.
func (o *OracleTablespaceFile) GetMaxSizeMb() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxSizeMb
}

// GetMaxSizeMbOk returns a tuple with the MaxSizeMb field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceFile) GetMaxSizeMbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSizeMb, true
}

// SetMaxSizeMb sets field value.
func (o *OracleTablespaceFile) SetMaxSizeMb(v float64) {
	o.MaxSizeMb = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OracleTablespaceFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["allocatedSizeMB"] = o.AllocatedSizeMb
	toSerialize["usedSizeMB"] = o.UsedSizeMb
	toSerialize["autoExtend"] = o.AutoExtend
	toSerialize["nextSizeMB"] = o.NextSizeMb
	toSerialize["maxSizeMB"] = o.MaxSizeMb

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OracleTablespaceFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id              *int64   `json:"id"`
		Name            *string  `json:"name"`
		Status          *string  `json:"status"`
		AllocatedSizeMb *float64 `json:"allocatedSizeMB"`
		UsedSizeMb      *float64 `json:"usedSizeMB"`
		AutoExtend      *bool    `json:"autoExtend"`
		NextSizeMb      *float64 `json:"nextSizeMB"`
		MaxSizeMb       *float64 `json:"maxSizeMB"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.AllocatedSizeMb == nil {
		return fmt.Errorf("required field allocatedSizeMB missing")
	}
	if all.UsedSizeMb == nil {
		return fmt.Errorf("required field usedSizeMB missing")
	}
	if all.AutoExtend == nil {
		return fmt.Errorf("required field autoExtend missing")
	}
	if all.NextSizeMb == nil {
		return fmt.Errorf("required field nextSizeMB missing")
	}
	if all.MaxSizeMb == nil {
		return fmt.Errorf("required field maxSizeMB missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "status", "allocatedSizeMB", "usedSizeMB", "autoExtend", "nextSizeMB", "maxSizeMB"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Name = *all.Name
	o.Status = *all.Status
	o.AllocatedSizeMb = *all.AllocatedSizeMb
	o.UsedSizeMb = *all.UsedSizeMb
	o.AutoExtend = *all.AutoExtend
	o.NextSizeMb = *all.NextSizeMb
	o.MaxSizeMb = *all.MaxSizeMb

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
