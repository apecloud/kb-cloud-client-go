// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeTarget struct {
	// Native object kind returned by capabilities. Table includes views; view/foreignTable/materializedView also validate the concrete kind.
	ObjectType string `json:"objectType"`
	// Database name used to locate the resource. PostgreSQL also uses this as the connection database, including cluster-wide objects.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Native schema containing the object, or the schema itself for schema objects.
	SchemaName *string `json:"schemaName,omitempty"`
	// Exact object name, never a SQL fragment. Required except for global, database, schema, and largeObject.
	ObjectName *string `json:"objectName,omitempty"`
	// Opaque native identity returned by discovery. Interpretation and validation belong to the engine. PostgreSQL routines and large objects require a nonzero uint32 OID; routine names must also match.
	ObjectId *string `json:"objectId,omitempty"`
	// Native catalog name when required by engine capabilities. Unsupported fields are rejected.
	CatalogName *string `json:"catalogName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeTarget instantiates a new ObjectPrivilegeTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeTarget(objectType string) *ObjectPrivilegeTarget {
	this := ObjectPrivilegeTarget{}
	this.ObjectType = objectType
	return &this
}

// NewObjectPrivilegeTargetWithDefaults instantiates a new ObjectPrivilegeTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeTargetWithDefaults() *ObjectPrivilegeTarget {
	this := ObjectPrivilegeTarget{}
	return &this
}

// GetObjectType returns the ObjectType field value.
func (o *ObjectPrivilegeTarget) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeTarget) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value.
func (o *ObjectPrivilegeTarget) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *ObjectPrivilegeTarget) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeTarget) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *ObjectPrivilegeTarget) HasDatabaseName() bool {
	return o != nil && o.DatabaseName != nil
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *ObjectPrivilegeTarget) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *ObjectPrivilegeTarget) GetSchemaName() string {
	if o == nil || o.SchemaName == nil {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeTarget) GetSchemaNameOk() (*string, bool) {
	if o == nil || o.SchemaName == nil {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *ObjectPrivilegeTarget) HasSchemaName() bool {
	return o != nil && o.SchemaName != nil
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *ObjectPrivilegeTarget) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ObjectPrivilegeTarget) GetObjectName() string {
	if o == nil || o.ObjectName == nil {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeTarget) GetObjectNameOk() (*string, bool) {
	if o == nil || o.ObjectName == nil {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ObjectPrivilegeTarget) HasObjectName() bool {
	return o != nil && o.ObjectName != nil
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ObjectPrivilegeTarget) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *ObjectPrivilegeTarget) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeTarget) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *ObjectPrivilegeTarget) HasObjectId() bool {
	return o != nil && o.ObjectId != nil
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *ObjectPrivilegeTarget) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetCatalogName returns the CatalogName field value if set, zero value otherwise.
func (o *ObjectPrivilegeTarget) GetCatalogName() string {
	if o == nil || o.CatalogName == nil {
		var ret string
		return ret
	}
	return *o.CatalogName
}

// GetCatalogNameOk returns a tuple with the CatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeTarget) GetCatalogNameOk() (*string, bool) {
	if o == nil || o.CatalogName == nil {
		return nil, false
	}
	return o.CatalogName, true
}

// HasCatalogName returns a boolean if a field has been set.
func (o *ObjectPrivilegeTarget) HasCatalogName() bool {
	return o != nil && o.CatalogName != nil
}

// SetCatalogName gets a reference to the given string and assigns it to the CatalogName field.
func (o *ObjectPrivilegeTarget) SetCatalogName(v string) {
	o.CatalogName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["objectType"] = o.ObjectType
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if o.SchemaName != nil {
		toSerialize["schemaName"] = o.SchemaName
	}
	if o.ObjectName != nil {
		toSerialize["objectName"] = o.ObjectName
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.CatalogName != nil {
		toSerialize["catalogName"] = o.CatalogName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ObjectType   *string `json:"objectType"`
		DatabaseName *string `json:"databaseName,omitempty"`
		SchemaName   *string `json:"schemaName,omitempty"`
		ObjectName   *string `json:"objectName,omitempty"`
		ObjectId     *string `json:"objectId,omitempty"`
		CatalogName  *string `json:"catalogName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ObjectType == nil {
		return fmt.Errorf("required field objectType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"objectType", "databaseName", "schemaName", "objectName", "objectId", "catalogName"})
	} else {
		return err
	}
	o.ObjectType = *all.ObjectType
	o.DatabaseName = all.DatabaseName
	o.SchemaName = all.SchemaName
	o.ObjectName = all.ObjectName
	o.ObjectId = all.ObjectId
	o.CatalogName = all.CatalogName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
