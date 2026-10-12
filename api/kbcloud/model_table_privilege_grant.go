// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TablePrivilegeGrant struct {
	Name string `json:"name"`
	// Empty means a table-level privilege.
	Column    *string `json:"column,omitempty"`
	Grantable bool    `json:"grantable"`
	// PostgreSQL ACL grantor. Omitted for MySQL and effective privilege results.
	Grantor *string `json:"grantor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTablePrivilegeGrant instantiates a new TablePrivilegeGrant object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTablePrivilegeGrant(name string, grantable bool) *TablePrivilegeGrant {
	this := TablePrivilegeGrant{}
	this.Name = name
	this.Grantable = grantable
	return &this
}

// NewTablePrivilegeGrantWithDefaults instantiates a new TablePrivilegeGrant object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTablePrivilegeGrantWithDefaults() *TablePrivilegeGrant {
	this := TablePrivilegeGrant{}
	return &this
}

// GetName returns the Name field value.
func (o *TablePrivilegeGrant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeGrant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TablePrivilegeGrant) SetName(v string) {
	o.Name = v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *TablePrivilegeGrant) GetColumn() string {
	if o == nil || o.Column == nil {
		var ret string
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeGrant) GetColumnOk() (*string, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *TablePrivilegeGrant) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given string and assigns it to the Column field.
func (o *TablePrivilegeGrant) SetColumn(v string) {
	o.Column = &v
}

// GetGrantable returns the Grantable field value.
func (o *TablePrivilegeGrant) GetGrantable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Grantable
}

// GetGrantableOk returns a tuple with the Grantable field value
// and a boolean to check if the value has been set.
func (o *TablePrivilegeGrant) GetGrantableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grantable, true
}

// SetGrantable sets field value.
func (o *TablePrivilegeGrant) SetGrantable(v bool) {
	o.Grantable = v
}

// GetGrantor returns the Grantor field value if set, zero value otherwise.
func (o *TablePrivilegeGrant) GetGrantor() string {
	if o == nil || o.Grantor == nil {
		var ret string
		return ret
	}
	return *o.Grantor
}

// GetGrantorOk returns a tuple with the Grantor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablePrivilegeGrant) GetGrantorOk() (*string, bool) {
	if o == nil || o.Grantor == nil {
		return nil, false
	}
	return o.Grantor, true
}

// HasGrantor returns a boolean if a field has been set.
func (o *TablePrivilegeGrant) HasGrantor() bool {
	return o != nil && o.Grantor != nil
}

// SetGrantor gets a reference to the given string and assigns it to the Grantor field.
func (o *TablePrivilegeGrant) SetGrantor(v string) {
	o.Grantor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TablePrivilegeGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	toSerialize["grantable"] = o.Grantable
	if o.Grantor != nil {
		toSerialize["grantor"] = o.Grantor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TablePrivilegeGrant) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name"`
		Column    *string `json:"column,omitempty"`
		Grantable *bool   `json:"grantable"`
		Grantor   *string `json:"grantor,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Grantable == nil {
		return fmt.Errorf("required field grantable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "column", "grantable", "grantor"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Column = all.Column
	o.Grantable = *all.Grantable
	o.Grantor = all.Grantor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
