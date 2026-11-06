// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBPermission struct {
	Action   InfluxDBPermissionAction `json:"action"`
	Resource InfluxDBResource         `json:"resource"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBPermission instantiates a new InfluxDBPermission object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBPermission(action InfluxDBPermissionAction, resource InfluxDBResource) *InfluxDBPermission {
	this := InfluxDBPermission{}
	this.Action = action
	this.Resource = resource
	return &this
}

// NewInfluxDBPermissionWithDefaults instantiates a new InfluxDBPermission object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBPermissionWithDefaults() *InfluxDBPermission {
	this := InfluxDBPermission{}
	return &this
}

// GetAction returns the Action field value.
func (o *InfluxDBPermission) GetAction() InfluxDBPermissionAction {
	if o == nil {
		var ret InfluxDBPermissionAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *InfluxDBPermission) GetActionOk() (*InfluxDBPermissionAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *InfluxDBPermission) SetAction(v InfluxDBPermissionAction) {
	o.Action = v
}

// GetResource returns the Resource field value.
func (o *InfluxDBPermission) GetResource() InfluxDBResource {
	if o == nil {
		var ret InfluxDBResource
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *InfluxDBPermission) GetResourceOk() (*InfluxDBResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value.
func (o *InfluxDBPermission) SetResource(v InfluxDBResource) {
	o.Resource = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["resource"] = o.Resource

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBPermission) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action   *InfluxDBPermissionAction `json:"action"`
		Resource *InfluxDBResource         `json:"resource"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Resource == nil {
		return fmt.Errorf("required field resource missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"action", "resource"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	if all.Resource.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Resource = *all.Resource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
