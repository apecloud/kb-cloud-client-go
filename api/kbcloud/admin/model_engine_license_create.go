// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineLicenseCreate struct {
	// Name of the engine license
	Name string `json:"name"`
	// Name of the engine
	EngineName string `json:"engineName"`
	// Description of the engine license
	Description *string `json:"description,omitempty"`
	// Expiration date and time of the license (optional)
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineLicenseCreate instantiates a new EngineLicenseCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineLicenseCreate(name string, engineName string) *EngineLicenseCreate {
	this := EngineLicenseCreate{}
	this.Name = name
	this.EngineName = engineName
	return &this
}

// NewEngineLicenseCreateWithDefaults instantiates a new EngineLicenseCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineLicenseCreateWithDefaults() *EngineLicenseCreate {
	this := EngineLicenseCreate{}
	return &this
}

// GetName returns the Name field value.
func (o *EngineLicenseCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EngineLicenseCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EngineLicenseCreate) SetName(v string) {
	o.Name = v
}

// GetEngineName returns the EngineName field value.
func (o *EngineLicenseCreate) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineLicenseCreate) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineLicenseCreate) SetEngineName(v string) {
	o.EngineName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EngineLicenseCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicenseCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EngineLicenseCreate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EngineLicenseCreate) SetDescription(v string) {
	o.Description = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *EngineLicenseCreate) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineLicenseCreate) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *EngineLicenseCreate) HasExpiredAt() bool {
	return o != nil && o.ExpiredAt != nil
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *EngineLicenseCreate) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineLicenseCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["engineName"] = o.EngineName
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineLicenseCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string    `json:"name"`
		EngineName  *string    `json:"engineName"`
		Description *string    `json:"description,omitempty"`
		ExpiredAt   *time.Time `json:"expiredAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "engineName", "description", "expiredAt"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.EngineName = *all.EngineName
	o.Description = all.Description
	o.ExpiredAt = all.ExpiredAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
