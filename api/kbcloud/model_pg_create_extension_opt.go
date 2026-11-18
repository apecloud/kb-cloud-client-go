// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PgCreateExtensionOpt struct {
	// Target databases. At most 100 database-extension pairs per request.
	Databases []string `json:"databases"`
	// Extensions to enable using their default versions and schemas.
	Names []string `json:"names"`
	// The version of the extension to install.
	Version *string `json:"version,omitempty"`
	// The schema where the extension is installed.
	Schema *string `json:"schema,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPgCreateExtensionOpt instantiates a new PgCreateExtensionOpt object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPgCreateExtensionOpt(databases []string, names []string) *PgCreateExtensionOpt {
	this := PgCreateExtensionOpt{}
	this.Databases = databases
	this.Names = names
	return &this
}

// NewPgCreateExtensionOptWithDefaults instantiates a new PgCreateExtensionOpt object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPgCreateExtensionOptWithDefaults() *PgCreateExtensionOpt {
	this := PgCreateExtensionOpt{}
	return &this
}

// GetDatabases returns the Databases field value.
func (o *PgCreateExtensionOpt) GetDatabases() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value
// and a boolean to check if the value has been set.
func (o *PgCreateExtensionOpt) GetDatabasesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Databases, true
}

// SetDatabases sets field value.
func (o *PgCreateExtensionOpt) SetDatabases(v []string) {
	o.Databases = v
}

// GetNames returns the Names field value.
func (o *PgCreateExtensionOpt) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *PgCreateExtensionOpt) GetNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value.
func (o *PgCreateExtensionOpt) SetNames(v []string) {
	o.Names = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PgCreateExtensionOpt) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgCreateExtensionOpt) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PgCreateExtensionOpt) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PgCreateExtensionOpt) SetVersion(v string) {
	o.Version = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PgCreateExtensionOpt) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgCreateExtensionOpt) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PgCreateExtensionOpt) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PgCreateExtensionOpt) SetSchema(v string) {
	o.Schema = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PgCreateExtensionOpt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["databases"] = o.Databases
	toSerialize["names"] = o.Names
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
func (o *PgCreateExtensionOpt) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Databases *[]string `json:"databases"`
		Names     *[]string `json:"names"`
		Version   *string   `json:"version,omitempty"`
		Schema    *string   `json:"schema,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"databases", "names", "version", "schema"})
	} else {
		return err
	}
	o.Databases = *all.Databases
	o.Names = *all.Names
	o.Version = all.Version
	o.Schema = all.Schema

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
