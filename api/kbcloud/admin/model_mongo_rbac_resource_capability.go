// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MongoRBACResourceCapability struct {
	// cluster, database, collection, allDatabases, collectionAcrossDatabases
	Type    string   `json:"type"`
	Actions []string `json:"actions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACResourceCapability instantiates a new MongoRBACResourceCapability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACResourceCapability(typeVar string, actions []string) *MongoRBACResourceCapability {
	this := MongoRBACResourceCapability{}
	this.Type = typeVar
	this.Actions = actions
	return &this
}

// NewMongoRBACResourceCapabilityWithDefaults instantiates a new MongoRBACResourceCapability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACResourceCapabilityWithDefaults() *MongoRBACResourceCapability {
	this := MongoRBACResourceCapability{}
	return &this
}

// GetType returns the Type field value.
func (o *MongoRBACResourceCapability) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MongoRBACResourceCapability) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *MongoRBACResourceCapability) SetType(v string) {
	o.Type = v
}

// GetActions returns the Actions field value.
func (o *MongoRBACResourceCapability) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *MongoRBACResourceCapability) GetActionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value.
func (o *MongoRBACResourceCapability) SetActions(v []string) {
	o.Actions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACResourceCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["actions"] = o.Actions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACResourceCapability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type    *string   `json:"type"`
		Actions *[]string `json:"actions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Actions == nil {
		return fmt.Errorf("required field actions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "actions"})
	} else {
		return err
	}
	o.Type = *all.Type
	o.Actions = *all.Actions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
