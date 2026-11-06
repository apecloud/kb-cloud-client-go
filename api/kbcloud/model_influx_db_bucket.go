// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBBucket struct {
	Id    string `json:"id"`
	OrgId string `json:"orgID"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBBucket instantiates a new InfluxDBBucket object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBBucket(id string, orgId string, name string, typeVar string) *InfluxDBBucket {
	this := InfluxDBBucket{}
	this.Id = id
	this.OrgId = orgId
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewInfluxDBBucketWithDefaults instantiates a new InfluxDBBucket object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBBucketWithDefaults() *InfluxDBBucket {
	this := InfluxDBBucket{}
	return &this
}

// GetId returns the Id field value.
func (o *InfluxDBBucket) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucket) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *InfluxDBBucket) SetId(v string) {
	o.Id = v
}

// GetOrgId returns the OrgId field value.
func (o *InfluxDBBucket) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucket) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *InfluxDBBucket) SetOrgId(v string) {
	o.OrgId = v
}

// GetName returns the Name field value.
func (o *InfluxDBBucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InfluxDBBucket) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *InfluxDBBucket) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InfluxDBBucket) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *InfluxDBBucket) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["orgID"] = o.OrgId
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBBucket) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id    *string `json:"id"`
		OrgId *string `json:"orgID"`
		Name  *string `json:"name"`
		Type  *string `json:"type"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgID missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgID", "name", "type"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.OrgId = *all.OrgId
	o.Name = *all.Name
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
