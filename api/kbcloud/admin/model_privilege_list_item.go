// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PrivilegeListItem Database and its assigned privileges.
// NODESCRIPTION PrivilegeListItem
//
// Deprecated: This model is deprecated.
type PrivilegeListItem struct {
	// The name of the database.
	DatabaseName *string `json:"databaseName,omitempty"`
	// The type of privilege.
	Privileges string `json:"privileges"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPrivilegeListItem instantiates a new PrivilegeListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPrivilegeListItem(privileges string) *PrivilegeListItem {
	this := PrivilegeListItem{}
	this.Privileges = privileges
	return &this
}

// NewPrivilegeListItemWithDefaults instantiates a new PrivilegeListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPrivilegeListItemWithDefaults() *PrivilegeListItem {
	this := PrivilegeListItem{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *PrivilegeListItem) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegeListItem) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *PrivilegeListItem) HasDatabaseName() bool {
	return o != nil && o.DatabaseName != nil
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *PrivilegeListItem) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetPrivileges returns the Privileges field value.
func (o *PrivilegeListItem) GetPrivileges() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *PrivilegeListItem) GetPrivilegesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *PrivilegeListItem) SetPrivileges(v string) {
	o.Privileges = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PrivilegeListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	toSerialize["privileges"] = o.Privileges

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PrivilegeListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabaseName *string `json:"databaseName,omitempty"`
		Privileges   *string `json:"privileges"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Privileges == nil {
		return fmt.Errorf("required field privileges missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databaseName", "privileges"})
	} else {
		return err
	}
	o.DatabaseName = all.DatabaseName
	o.Privileges = *all.Privileges

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
