// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBUser struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	// Management or privileged users cannot be modified here.
	Protected bool `json:"protected"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBUser instantiates a new InfluxDBUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBUser(id string, name string, status string, protected bool) *InfluxDBUser {
	this := InfluxDBUser{}
	this.Id = id
	this.Name = name
	this.Status = status
	this.Protected = protected
	return &this
}

// NewInfluxDBUserWithDefaults instantiates a new InfluxDBUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBUserWithDefaults() *InfluxDBUser {
	this := InfluxDBUser{}
	return &this
}

// GetId returns the Id field value.
func (o *InfluxDBUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InfluxDBUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *InfluxDBUser) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *InfluxDBUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InfluxDBUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InfluxDBUser) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value.
func (o *InfluxDBUser) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InfluxDBUser) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *InfluxDBUser) SetStatus(v string) {
	o.Status = v
}

// GetProtected returns the Protected field value.
func (o *InfluxDBUser) GetProtected() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value
// and a boolean to check if the value has been set.
func (o *InfluxDBUser) GetProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protected, true
}

// SetProtected sets field value.
func (o *InfluxDBUser) SetProtected(v bool) {
	o.Protected = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["protected"] = o.Protected

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string `json:"id"`
		Name      *string `json:"name"`
		Status    *string `json:"status"`
		Protected *bool   `json:"protected"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Protected == nil {
		return fmt.Errorf("required field protected missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "status", "protected"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Name = *all.Name
	o.Status = *all.Status
	o.Protected = *all.Protected

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
