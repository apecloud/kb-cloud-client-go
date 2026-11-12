// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBTokenCreate struct {
	OrgId       string                     `json:"orgID"`
	Description *string                    `json:"description,omitempty"`
	Permissions []InfluxDBBucketPermission `json:"permissions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInfluxDBTokenCreate instantiates a new InfluxDBTokenCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInfluxDBTokenCreate(orgId string, permissions []InfluxDBBucketPermission) *InfluxDBTokenCreate {
	this := InfluxDBTokenCreate{}
	this.OrgId = orgId
	this.Permissions = permissions
	return &this
}

// NewInfluxDBTokenCreateWithDefaults instantiates a new InfluxDBTokenCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInfluxDBTokenCreateWithDefaults() *InfluxDBTokenCreate {
	this := InfluxDBTokenCreate{}
	return &this
}

// GetOrgId returns the OrgId field value.
func (o *InfluxDBTokenCreate) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreate) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *InfluxDBTokenCreate) SetOrgId(v string) {
	o.OrgId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InfluxDBTokenCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InfluxDBTokenCreate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InfluxDBTokenCreate) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value.
func (o *InfluxDBTokenCreate) GetPermissions() []InfluxDBBucketPermission {
	if o == nil {
		var ret []InfluxDBBucketPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *InfluxDBTokenCreate) GetPermissionsOk() (*[]InfluxDBBucketPermission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value.
func (o *InfluxDBTokenCreate) SetPermissions(v []InfluxDBBucketPermission) {
	o.Permissions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InfluxDBTokenCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["orgID"] = o.OrgId
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["permissions"] = o.Permissions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InfluxDBTokenCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgId       *string                     `json:"orgID"`
		Description *string                     `json:"description,omitempty"`
		Permissions *[]InfluxDBBucketPermission `json:"permissions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgID missing")
	}
	if all.Permissions == nil {
		return fmt.Errorf("required field permissions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgID", "description", "permissions"})
	} else {
		return err
	}
	o.OrgId = *all.OrgId
	o.Description = all.Description
	o.Permissions = *all.Permissions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
