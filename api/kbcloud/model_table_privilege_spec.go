// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TablePrivilegeSpec struct {
	// Native privilege keyword returned by supportedTablePrivileges or supportedColumnPrivileges.
	Name string `json:"name"`
	// Exact column names. Empty or omitted means the whole table.
	Columns []string `json:"columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTablePrivilegeSpec instantiates a new TablePrivilegeSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTablePrivilegeSpec(name string) *TablePrivilegeSpec {
	this := TablePrivilegeSpec{}
	this.Name = name
	return &this
}

// NewTablePrivilegeSpecWithDefaults instantiates a new TablePrivilegeSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTablePrivilegeSpecWithDefaults() *TablePrivilegeSpec {
	this := TablePrivilegeSpec{}
	return &this
}

// GetName returns the Name field value.
func (o *TablePrivilegeSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TablePrivilegeSpec) SetName(v string) {
	o.Name = v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TablePrivilegeSpec) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeSpec) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TablePrivilegeSpec) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *TablePrivilegeSpec) SetColumns(v []string) {
	o.Columns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TablePrivilegeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TablePrivilegeSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string  `json:"name"`
		Columns []string `json:"columns,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "columns"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Columns = all.Columns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
