// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Apikey APIKey is the key for API access
// NODESCRIPTION Apikey
//
// Deprecated: This model is deprecated.
type Apikey struct {
	// The name of the APIKey
	AccessKey string `json:"accessKey"`
	// The description for APIKey
	Description string `json:"description"`
	// The expired time of APIKey, return empty without setting duration
	ExpiredAt time.Time `json:"expiredAt"`
	// The create time of APIKey
	CreateAt time.Time `json:"createAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApikey instantiates a new Apikey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApikey(accessKey string, description string, expiredAt time.Time, createAt time.Time) *Apikey {
	this := Apikey{}
	this.AccessKey = accessKey
	this.Description = description
	this.ExpiredAt = expiredAt
	this.CreateAt = createAt
	return &this
}

// NewApikeyWithDefaults instantiates a new Apikey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApikeyWithDefaults() *Apikey {
	this := Apikey{}
	return &this
}

// GetAccessKey returns the AccessKey field value.
func (o *Apikey) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *Apikey) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value.
func (o *Apikey) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetDescription returns the Description field value.
func (o *Apikey) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Apikey) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *Apikey) SetDescription(v string) {
	o.Description = v
}

// GetExpiredAt returns the ExpiredAt field value.
func (o *Apikey) GetExpiredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
func (o *Apikey) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiredAt, true
}

// SetExpiredAt sets field value.
func (o *Apikey) SetExpiredAt(v time.Time) {
	o.ExpiredAt = v
}

// GetCreateAt returns the CreateAt field value.
func (o *Apikey) GetCreateAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreateAt
}

// GetCreateAtOk returns a tuple with the CreateAt field value
// and a boolean to check if the value has been set.
func (o *Apikey) GetCreateAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateAt, true
}

// SetCreateAt sets field value.
func (o *Apikey) SetCreateAt(v time.Time) {
	o.CreateAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Apikey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["accessKey"] = o.AccessKey
	toSerialize["description"] = o.Description
	if o.ExpiredAt.Nanosecond() == 0 {
		toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.CreateAt.Nanosecond() == 0 {
		toSerialize["createAt"] = o.CreateAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createAt"] = o.CreateAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Apikey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKey   *string    `json:"accessKey"`
		Description *string    `json:"description"`
		ExpiredAt   *time.Time `json:"expiredAt"`
		CreateAt    *time.Time `json:"createAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessKey == nil {
		return fmt.Errorf("required field accessKey missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.ExpiredAt == nil {
		return fmt.Errorf("required field expiredAt missing")
	}
	if all.CreateAt == nil {
		return fmt.Errorf("required field createAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"accessKey", "description", "expiredAt", "createAt"})
	} else {
		return err
	}
	o.AccessKey = *all.AccessKey
	o.Description = *all.Description
	o.ExpiredAt = *all.ExpiredAt
	o.CreateAt = *all.CreateAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
