// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ApikeyCreate APIKeyCreate is the payload for creating an APIKey

type ApikeyCreate struct {
	// The description of the APIKey
	Description string `json:"description"`
	// The expired time of the APIKey
	ExpiredAt time.Time `json:"expiredAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApikeyCreate instantiates a new ApikeyCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApikeyCreate(description string, expiredAt time.Time) *ApikeyCreate {
	this := ApikeyCreate{}
	this.Description = description
	this.ExpiredAt = expiredAt
	return &this
}

// NewApikeyCreateWithDefaults instantiates a new ApikeyCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApikeyCreateWithDefaults() *ApikeyCreate {
	this := ApikeyCreate{}
	return &this
}

// GetDescription returns the Description field value.
func (o *ApikeyCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApikeyCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ApikeyCreate) SetDescription(v string) {
	o.Description = v
}

// GetExpiredAt returns the ExpiredAt field value.
func (o *ApikeyCreate) GetExpiredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
func (o *ApikeyCreate) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiredAt, true
}

// SetExpiredAt sets field value.
func (o *ApikeyCreate) SetExpiredAt(v time.Time) {
	o.ExpiredAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApikeyCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	if o.ExpiredAt.Nanosecond() == 0 {
		toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApikeyCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string    `json:"description"`
		ExpiredAt   *time.Time `json:"expiredAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.ExpiredAt == nil {
		return fmt.Errorf("required field expiredAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "expiredAt"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.ExpiredAt = *all.ExpiredAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
