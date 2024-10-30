// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type EngineLicense struct {
	EngineName string `json:"engineName"`
	Name string `json:"name"`
	Key *string `json:"key,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IdString *string `json:"idString,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewEngineLicense instantiates a new EngineLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineLicense(engineName string, name string) *EngineLicense {
	this := EngineLicense{}
	this.EngineName = engineName
	this.Name = name
	return &this
}

// NewEngineLicenseWithDefaults instantiates a new EngineLicense object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineLicenseWithDefaults() *EngineLicense {
	this := EngineLicense{}
	return &this
}
// GetEngineName returns the EngineName field value.
func (o *EngineLicense) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineLicense) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineLicense) SetEngineName(v string) {
	o.EngineName = v
}


// GetName returns the Name field value.
func (o *EngineLicense) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EngineLicense) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EngineLicense) SetName(v string) {
	o.Name = v
}


// GetKey returns the Key field value if set, zero value otherwise.
func (o *EngineLicense) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicense) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EngineLicense) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EngineLicense) SetKey(v string) {
	o.Key = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EngineLicense) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicense) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EngineLicense) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EngineLicense) SetDescription(v string) {
	o.Description = &v
}


// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *EngineLicense) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicense) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *EngineLicense) HasExpiredAt() bool {
	return o != nil && o.ExpiredAt != nil
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *EngineLicense) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}


// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EngineLicense) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicense) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EngineLicense) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EngineLicense) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *EngineLicense) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicense) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EngineLicense) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EngineLicense) SetId(v int32) {
	o.Id = &v
}


// GetIdString returns the IdString field value if set, zero value otherwise.
func (o *EngineLicense) GetIdString() string {
	if o == nil || o.IdString == nil {
		var ret string
		return ret
	}
	return *o.IdString
}

// GetIdStringOk returns a tuple with the IdString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicense) GetIdStringOk() (*string, bool) {
	if o == nil || o.IdString == nil {
		return nil, false
	}
	return o.IdString, true
}

// HasIdString returns a boolean if a field has been set.
func (o *EngineLicense) HasIdString() bool {
	return o != nil && o.IdString != nil
}

// SetIdString gets a reference to the given string and assigns it to the IdString field.
func (o *EngineLicense) SetIdString(v string) {
	o.IdString = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o EngineLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["name"] = o.Name
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExpiredAt != nil {
		if o.ExpiredAt.Nanosecond() == 0 {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdString != nil {
		toSerialize["idString"] = o.IdString
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineLicense) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName *string `json:"engineName"`
		Name *string `json:"name"`
		Key *string `json:"key,omitempty"`
		Description *string `json:"description,omitempty"`
		ExpiredAt *time.Time `json:"expiredAt,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		Id *int32 `json:"id,omitempty"`
		IdString *string `json:"idString,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "engineName", "name", "key", "description", "expiredAt", "createdAt", "id", "idString",  })
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Name = *all.Name
	o.Key = all.Key
	o.Description = all.Description
	o.ExpiredAt = all.ExpiredAt
	o.CreatedAt = all.CreatedAt
	o.Id = all.Id
	o.IdString = all.IdString

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
