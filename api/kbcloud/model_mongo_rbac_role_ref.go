// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MongoRBACRoleRef struct {
	Role string `json:"role"`
	Db   string `json:"db"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACRoleRef instantiates a new MongoRBACRoleRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACRoleRef(role string, db string) *MongoRBACRoleRef {
	this := MongoRBACRoleRef{}
	this.Role = role
	this.Db = db
	return &this
}

// NewMongoRBACRoleRefWithDefaults instantiates a new MongoRBACRoleRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACRoleRefWithDefaults() *MongoRBACRoleRef {
	this := MongoRBACRoleRef{}
	return &this
}

// GetRole returns the Role field value.
func (o *MongoRBACRoleRef) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *MongoRBACRoleRef) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *MongoRBACRoleRef) SetRole(v string) {
	o.Role = v
}

// GetDb returns the Db field value.
func (o *MongoRBACRoleRef) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *MongoRBACRoleRef) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value.
func (o *MongoRBACRoleRef) SetDb(v string) {
	o.Db = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACRoleRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role
	toSerialize["db"] = o.Db

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACRoleRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role *string `json:"role"`
		Db   *string `json:"db"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.Db == nil {
		return fmt.Errorf("required field db missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "db"})
	} else {
		return err
	}
	o.Role = *all.Role
	o.Db = *all.Db

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
