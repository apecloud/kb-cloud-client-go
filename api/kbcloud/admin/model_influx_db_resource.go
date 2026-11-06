// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBResource struct {
	Type  string  `json:"type"`
	Id    *string `json:"id,omitempty"`
	OrgId *string `json:"orgID,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBResource instantiates a new InfluxDBResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBResource(typeVar string) *InfluxDBResource {
	this := InfluxDBResource{}
	this.Type = typeVar
	return &this
}

// NewInfluxDBResourceWithDefaults instantiates a new InfluxDBResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBResourceWithDefaults() *InfluxDBResource {
	this := InfluxDBResource{}
	return &this
}

// GetType returns the Type field value.
func (o *InfluxDBResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InfluxDBResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *InfluxDBResource) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InfluxDBResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InfluxDBResource) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InfluxDBResource) SetId(v string) {
	o.Id = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *InfluxDBResource) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBResource) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *InfluxDBResource) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *InfluxDBResource) SetOrgId(v string) {
	o.OrgId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrgId != nil {
		toSerialize["orgID"] = o.OrgId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *string `json:"type"`
		Id    *string `json:"id,omitempty"`
		OrgId *string `json:"orgID,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "id", "orgID"})
	} else {
		return err
	}
	o.Type = *all.Type
	o.Id = all.Id
	o.OrgId = all.OrgId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
