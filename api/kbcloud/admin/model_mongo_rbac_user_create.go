// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MongoRBACUserCreate struct {
	User     string             `json:"user"`
	Db       string             `json:"db"`
	Password string             `json:"password"`
	Roles    []MongoRBACRoleRef `json:"roles"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACUserCreate instantiates a new MongoRBACUserCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACUserCreate(user string, db string, password string, roles []MongoRBACRoleRef) *MongoRBACUserCreate {
	this := MongoRBACUserCreate{}
	this.User = user
	this.Db = db
	this.Password = password
	this.Roles = roles
	return &this
}

// NewMongoRBACUserCreateWithDefaults instantiates a new MongoRBACUserCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACUserCreateWithDefaults() *MongoRBACUserCreate {
	this := MongoRBACUserCreate{}
	return &this
}

// GetUser returns the User field value.
func (o *MongoRBACUserCreate) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *MongoRBACUserCreate) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *MongoRBACUserCreate) SetUser(v string) {
	o.User = v
}

// GetDb returns the Db field value.
func (o *MongoRBACUserCreate) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *MongoRBACUserCreate) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value.
func (o *MongoRBACUserCreate) SetDb(v string) {
	o.Db = v
}

// GetPassword returns the Password field value.
func (o *MongoRBACUserCreate) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *MongoRBACUserCreate) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *MongoRBACUserCreate) SetPassword(v string) {
	o.Password = v
}

// GetRoles returns the Roles field value.
func (o *MongoRBACUserCreate) GetRoles() []MongoRBACRoleRef {
	if o == nil {
		var ret []MongoRBACRoleRef
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *MongoRBACUserCreate) GetRolesOk() (*[]MongoRBACRoleRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value.
func (o *MongoRBACUserCreate) SetRoles(v []MongoRBACRoleRef) {
	o.Roles = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACUserCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["user"] = o.User
	toSerialize["db"] = o.Db
	toSerialize["password"] = o.Password
	toSerialize["roles"] = o.Roles

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACUserCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		User     *string             `json:"user"`
		Db       *string             `json:"db"`
		Password *string             `json:"password"`
		Roles    *[]MongoRBACRoleRef `json:"roles"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	if all.Db == nil {
		return fmt.Errorf("required field db missing")
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	if all.Roles == nil {
		return fmt.Errorf("required field roles missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"user", "db", "password", "roles"})
	} else {
		return err
	}
	o.User = *all.User
	o.Db = *all.Db
	o.Password = *all.Password
	o.Roles = *all.Roles

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
