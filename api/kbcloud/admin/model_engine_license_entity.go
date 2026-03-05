// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineLicenseEntity struct {
	Id *string `json:"id,omitempty"`
	// ID of the engine license
	LicenseId string `json:"licenseId"`
	// Name of the node
	NodeName    string     `json:"nodeName"`
	Key         *string    `json:"key,omitempty"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineLicenseEntity instantiates a new EngineLicenseEntity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineLicenseEntity(licenseId string, nodeName string) *EngineLicenseEntity {
	this := EngineLicenseEntity{}
	this.LicenseId = licenseId
	this.NodeName = nodeName
	return &this
}

// NewEngineLicenseEntityWithDefaults instantiates a new EngineLicenseEntity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineLicenseEntityWithDefaults() *EngineLicenseEntity {
	this := EngineLicenseEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EngineLicenseEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EngineLicenseEntity) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EngineLicenseEntity) SetId(v string) {
	o.Id = &v
}

// GetLicenseId returns the LicenseId field value.
func (o *EngineLicenseEntity) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntity) GetLicenseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value.
func (o *EngineLicenseEntity) SetLicenseId(v string) {
	o.LicenseId = v
}

// GetNodeName returns the NodeName field value.
func (o *EngineLicenseEntity) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntity) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value.
func (o *EngineLicenseEntity) SetNodeName(v string) {
	o.NodeName = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EngineLicenseEntity) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntity) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EngineLicenseEntity) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EngineLicenseEntity) SetKey(v string) {
	o.Key = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EngineLicenseEntity) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntity) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EngineLicenseEntity) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EngineLicenseEntity) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EngineLicenseEntity) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicenseEntity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EngineLicenseEntity) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EngineLicenseEntity) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineLicenseEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["licenseId"] = o.LicenseId
	toSerialize["nodeName"] = o.NodeName
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineLicenseEntity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string    `json:"id,omitempty"`
		LicenseId   *string    `json:"licenseId"`
		NodeName    *string    `json:"nodeName"`
		Key         *string    `json:"key,omitempty"`
		Description *string    `json:"description,omitempty"`
		CreatedAt   *time.Time `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.LicenseId == nil {
		return fmt.Errorf("required field licenseId missing")
	}
	if all.NodeName == nil {
		return fmt.Errorf("required field nodeName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "licenseId", "nodeName", "key", "description", "createdAt"})
	} else {
		return err
	}
	o.Id = all.Id
	o.LicenseId = *all.LicenseId
	o.NodeName = *all.NodeName
	o.Key = all.Key
	o.Description = all.Description
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
