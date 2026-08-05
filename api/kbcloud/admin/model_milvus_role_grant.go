// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MilvusRoleGrant Native Milvus role privilege grant.
type MilvusRoleGrant struct {
	// Milvus database name. Empty means Milvus default/global scope.
	DbName *string `json:"dbName,omitempty"`
	// Milvus collection name. Use '*' for all collections in the database.
	CollectionName *string `json:"collectionName,omitempty"`
	// Native Milvus privilege name, such as DatabaseReadOnly, DatabaseReadWrite, DatabaseAdmin, Search, Insert, or Delete.
	Privilege string `json:"privilege"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMilvusRoleGrant instantiates a new MilvusRoleGrant object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMilvusRoleGrant(privilege string) *MilvusRoleGrant {
	this := MilvusRoleGrant{}
	this.Privilege = privilege
	return &this
}

// NewMilvusRoleGrantWithDefaults instantiates a new MilvusRoleGrant object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMilvusRoleGrantWithDefaults() *MilvusRoleGrant {
	this := MilvusRoleGrant{}
	return &this
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *MilvusRoleGrant) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusRoleGrant) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *MilvusRoleGrant) HasDbName() bool {
	return o != nil && o.DbName != nil
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *MilvusRoleGrant) SetDbName(v string) {
	o.DbName = &v
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise.
func (o *MilvusRoleGrant) GetCollectionName() string {
	if o == nil || o.CollectionName == nil {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusRoleGrant) GetCollectionNameOk() (*string, bool) {
	if o == nil || o.CollectionName == nil {
		return nil, false
	}
	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *MilvusRoleGrant) HasCollectionName() bool {
	return o != nil && o.CollectionName != nil
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *MilvusRoleGrant) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetPrivilege returns the Privilege field value.
func (o *MilvusRoleGrant) GetPrivilege() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value
// and a boolean to check if the value has been set.
func (o *MilvusRoleGrant) GetPrivilegeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privilege, true
}

// SetPrivilege sets field value.
func (o *MilvusRoleGrant) SetPrivilege(v string) {
	o.Privilege = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MilvusRoleGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DbName != nil {
		toSerialize["dbName"] = o.DbName
	}
	if o.CollectionName != nil {
		toSerialize["collectionName"] = o.CollectionName
	}
	toSerialize["privilege"] = o.Privilege

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MilvusRoleGrant) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DbName         *string `json:"dbName,omitempty"`
		CollectionName *string `json:"collectionName,omitempty"`
		Privilege      *string `json:"privilege"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Privilege == nil {
		return fmt.Errorf("required field privilege missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"dbName", "collectionName", "privilege"})
	} else {
		return err
	}
	o.DbName = all.DbName
	o.CollectionName = all.CollectionName
	o.Privilege = *all.Privilege

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
