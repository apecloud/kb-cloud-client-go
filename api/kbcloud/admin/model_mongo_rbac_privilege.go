// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MongoRBACPrivilege struct {
	// Native MongoDB resource. cluster:true is system operations, not all database data. Empty db/collection are explicit wildcard scopes; system collections need explicit resources.
	Resource MongoRBACResource `json:"resource"`
	Actions  []string          `json:"actions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACPrivilege instantiates a new MongoRBACPrivilege object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACPrivilege(resource MongoRBACResource, actions []string) *MongoRBACPrivilege {
	this := MongoRBACPrivilege{}
	this.Resource = resource
	this.Actions = actions
	return &this
}

// NewMongoRBACPrivilegeWithDefaults instantiates a new MongoRBACPrivilege object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACPrivilegeWithDefaults() *MongoRBACPrivilege {
	this := MongoRBACPrivilege{}
	return &this
}

// GetResource returns the Resource field value.
func (o *MongoRBACPrivilege) GetResource() MongoRBACResource {
	if o == nil {
		var ret MongoRBACResource
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *MongoRBACPrivilege) GetResourceOk() (*MongoRBACResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *MongoRBACPrivilege) SetResource(v MongoRBACResource) {
	o.Resource = v
}

// GetActions returns the Actions field value.
func (o *MongoRBACPrivilege) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *MongoRBACPrivilege) GetActionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value.
func (o *MongoRBACPrivilege) SetActions(v []string) {
	o.Actions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACPrivilege) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["resource"] = o.Resource
	toSerialize["actions"] = o.Actions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACPrivilege) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Resource *MongoRBACResource `json:"resource"`
		Actions  *[]string          `json:"actions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Resource == nil {
		return fmt.Errorf("required field resource missing")
	}
	if all.Actions == nil {
		return fmt.Errorf("required field actions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"resource", "actions"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Resource.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Resource = *all.Resource
	o.Actions = *all.Actions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
