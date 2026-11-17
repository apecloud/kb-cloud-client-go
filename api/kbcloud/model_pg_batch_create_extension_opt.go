// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PgBatchCreateExtensionOpt struct {
	// Target databases. At most 100 database-extension pairs per request.
	Databases []string `json:"databases"`
	// Extensions to enable using their default versions and schemas.
	Names []string `json:"names"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPgBatchCreateExtensionOpt instantiates a new PgBatchCreateExtensionOpt object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPgBatchCreateExtensionOpt(databases []string, names []string) *PgBatchCreateExtensionOpt {
	this := PgBatchCreateExtensionOpt{}
	this.Databases = databases
	this.Names = names
	return &this
}

// NewPgBatchCreateExtensionOptWithDefaults instantiates a new PgBatchCreateExtensionOpt object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPgBatchCreateExtensionOptWithDefaults() *PgBatchCreateExtensionOpt {
	this := PgBatchCreateExtensionOpt{}
	return &this
}

// GetDatabases returns the Databases field value.
func (o *PgBatchCreateExtensionOpt) GetDatabases() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value
// and a boolean to check if the value has been set.
func (o *PgBatchCreateExtensionOpt) GetDatabasesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Databases, true
}

// SetDatabases sets field value.
func (o *PgBatchCreateExtensionOpt) SetDatabases(v []string) {
	o.Databases = v
}

// GetNames returns the Names field value.
func (o *PgBatchCreateExtensionOpt) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *PgBatchCreateExtensionOpt) GetNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value.
func (o *PgBatchCreateExtensionOpt) SetNames(v []string) {
	o.Names = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PgBatchCreateExtensionOpt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["databases"] = o.Databases
	toSerialize["names"] = o.Names

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PgBatchCreateExtensionOpt) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Databases *[]string `json:"databases"`
		Names     *[]string `json:"names"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Databases == nil {
		return fmt.Errorf("required field databases missing")
	}
	if all.Names == nil {
		return fmt.Errorf("required field names missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databases", "names"})
	} else {
		return err
	}
	o.Databases = *all.Databases
	o.Names = *all.Names

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
