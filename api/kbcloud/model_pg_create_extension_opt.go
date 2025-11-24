// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PGCreateExtensionOpt struct {
	// The name of the PostgreSQL extension.
	Name string `json:"name"`
	// The database where the extension is installed.
	Database string `json:"database"`
	// The version of the extension to install.
	Version *string `json:"version,omitempty"`
	// The schema where the extension is installed.
	Schema *string `json:"schema,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPGCreateExtensionOpt instantiates a new PGCreateExtensionOpt object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPGCreateExtensionOpt(name string, database string) *PGCreateExtensionOpt {
	this := PGCreateExtensionOpt{}
	this.Name = name
	this.Database = database
	return &this
}

// NewPGCreateExtensionOptWithDefaults instantiates a new PGCreateExtensionOpt object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPGCreateExtensionOptWithDefaults() *PGCreateExtensionOpt {
	this := PGCreateExtensionOpt{}
	return &this
}

// GetName returns the Name field value.
func (o *PGCreateExtensionOpt) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PGCreateExtensionOpt) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PGCreateExtensionOpt) SetName(v string) {
	o.Name = v
}

// GetDatabase returns the Database field value.
func (o *PGCreateExtensionOpt) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *PGCreateExtensionOpt) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *PGCreateExtensionOpt) SetDatabase(v string) {
	o.Database = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PGCreateExtensionOpt) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PGCreateExtensionOpt) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PGCreateExtensionOpt) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PGCreateExtensionOpt) SetVersion(v string) {
	o.Version = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PGCreateExtensionOpt) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PGCreateExtensionOpt) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PGCreateExtensionOpt) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PGCreateExtensionOpt) SetSchema(v string) {
	o.Schema = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PGCreateExtensionOpt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["database"] = o.Database
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PGCreateExtensionOpt) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name"`
		Database *string `json:"database"`
		Version  *string `json:"version,omitempty"`
		Schema   *string `json:"schema,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "database", "version", "schema"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Database = *all.Database
	o.Version = all.Version
	o.Schema = all.Schema

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
