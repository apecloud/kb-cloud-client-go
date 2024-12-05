// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InstanceTypes struct {
	// The unique identifier for the instance type.
	Id *string `json:"id,omitempty"`
	// Creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Any characters, no more than 255 characters long.
	Description *string `json:"description,omitempty"`
	// Any characters, no more than 63 characters long.
	DisplayName    *string         `json:"displayName,omitempty"`
	OccupationType *OccupationType `json:"occupationType,omitempty"`
	// Consists of lowercase English letters and numbers, 1-8 characters long. Used to add to the specification code.
	Name                *string              `json:"name,omitempty"`
	ServerConfiguration *ServerConfiguration `json:"serverConfiguration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceTypes instantiates a new InstanceTypes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceTypes() *InstanceTypes {
	this := InstanceTypes{}
	return &this
}

// NewInstanceTypesWithDefaults instantiates a new InstanceTypes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceTypesWithDefaults() *InstanceTypes {
	this := InstanceTypes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceTypes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceTypes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceTypes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InstanceTypes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InstanceTypes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InstanceTypes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceTypes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InstanceTypes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstanceTypes) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InstanceTypes) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypes) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InstanceTypes) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InstanceTypes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetOccupationType returns the OccupationType field value if set, zero value otherwise.
func (o *InstanceTypes) GetOccupationType() OccupationType {
	if o == nil || o.OccupationType == nil {
		var ret OccupationType
		return ret
	}
	return *o.OccupationType
}

// GetOccupationTypeOk returns a tuple with the OccupationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypes) GetOccupationTypeOk() (*OccupationType, bool) {
	if o == nil || o.OccupationType == nil {
		return nil, false
	}
	return o.OccupationType, true
}

// HasOccupationType returns a boolean if a field has been set.
func (o *InstanceTypes) HasOccupationType() bool {
	return o != nil && o.OccupationType != nil
}

// SetOccupationType gets a reference to the given OccupationType and assigns it to the OccupationType field.
func (o *InstanceTypes) SetOccupationType(v OccupationType) {
	o.OccupationType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceTypes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceTypes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceTypes) SetName(v string) {
	o.Name = &v
}

// GetServerConfiguration returns the ServerConfiguration field value if set, zero value otherwise.
func (o *InstanceTypes) GetServerConfiguration() ServerConfiguration {
	if o == nil || o.ServerConfiguration == nil {
		var ret ServerConfiguration
		return ret
	}
	return *o.ServerConfiguration
}

// GetServerConfigurationOk returns a tuple with the ServerConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypes) GetServerConfigurationOk() (*ServerConfiguration, bool) {
	if o == nil || o.ServerConfiguration == nil {
		return nil, false
	}
	return o.ServerConfiguration, true
}

// HasServerConfiguration returns a boolean if a field has been set.
func (o *InstanceTypes) HasServerConfiguration() bool {
	return o != nil && o.ServerConfiguration != nil
}

// SetServerConfiguration gets a reference to the given ServerConfiguration and assigns it to the ServerConfiguration field.
func (o *InstanceTypes) SetServerConfiguration(v ServerConfiguration) {
	o.ServerConfiguration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.OccupationType != nil {
		toSerialize["occupationType"] = o.OccupationType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ServerConfiguration != nil {
		toSerialize["serverConfiguration"] = o.ServerConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceTypes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                  *string              `json:"id,omitempty"`
		CreatedAt           *time.Time           `json:"createdAt,omitempty"`
		Description         *string              `json:"description,omitempty"`
		DisplayName         *string              `json:"displayName,omitempty"`
		OccupationType      *OccupationType      `json:"occupationType,omitempty"`
		Name                *string              `json:"name,omitempty"`
		ServerConfiguration *ServerConfiguration `json:"serverConfiguration,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "createdAt", "description", "displayName", "occupationType", "name", "serverConfiguration"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	if all.OccupationType != nil && !all.OccupationType.IsValid() {
		hasInvalidField = true
	} else {
		o.OccupationType = all.OccupationType
	}
	o.Name = all.Name
	if all.ServerConfiguration != nil && all.ServerConfiguration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ServerConfiguration = all.ServerConfiguration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
