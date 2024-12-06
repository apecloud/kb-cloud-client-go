// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"
)

// ApikeyWithSK APIKeyWithSK is the response for creating an APIKey
type ApikeyWithSK struct {
	// The accessKey of the APIKey
	AccessKey string `json:"accessKey"`
	// The secretKey of the APIKey
	SecretKey string `json:"secretKey"`
	// The description of the APIKey
	Description string `json:"description"`
	// The expired time of APIKey
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	// The create time of APIKey
	CreateAt time.Time `json:"createAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApikeyWithSK instantiates a new ApikeyWithSK object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApikeyWithSK(accessKey string, secretKey string, description string, createAt time.Time) *ApikeyWithSK {
	this := ApikeyWithSK{}
	this.AccessKey = accessKey
	this.SecretKey = secretKey
	this.Description = description
	this.CreateAt = createAt
	return &this
}

// NewApikeyWithSKWithDefaults instantiates a new ApikeyWithSK object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApikeyWithSKWithDefaults() *ApikeyWithSK {
	this := ApikeyWithSK{}
	return &this
}

// GetAccessKey returns the AccessKey field value.
func (o *ApikeyWithSK) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *ApikeyWithSK) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value.
func (o *ApikeyWithSK) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetSecretKey returns the SecretKey field value.
func (o *ApikeyWithSK) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *ApikeyWithSK) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value.
func (o *ApikeyWithSK) SetSecretKey(v string) {
	o.SecretKey = v
}

// GetDescription returns the Description field value.
func (o *ApikeyWithSK) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApikeyWithSK) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ApikeyWithSK) SetDescription(v string) {
	o.Description = v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *ApikeyWithSK) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApikeyWithSK) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *ApikeyWithSK) HasExpiredAt() bool {
	return o != nil && o.ExpiredAt != nil
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *ApikeyWithSK) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// GetCreateAt returns the CreateAt field value.
func (o *ApikeyWithSK) GetCreateAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreateAt
}

// GetCreateAtOk returns a tuple with the CreateAt field value
// and a boolean to check if the value has been set.
func (o *ApikeyWithSK) GetCreateAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateAt, true
}

// SetCreateAt sets field value.
func (o *ApikeyWithSK) SetCreateAt(v time.Time) {
	o.CreateAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApikeyWithSK) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["accessKey"] = o.AccessKey
	toSerialize["secretKey"] = o.SecretKey
	toSerialize["description"] = o.Description
	if o.ExpiredAt != nil {
		if o.ExpiredAt.Nanosecond() == 0 {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *ApikeyWithSK) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKey   *string    `json:"accessKey"`
		SecretKey   *string    `json:"secretKey"`
		Description *string    `json:"description"`
		ExpiredAt   *time.Time `json:"expiredAt,omitempty"`
		CreateAt    *time.Time `json:"createAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessKey == nil {
		return fmt.Errorf("required field accessKey missing")
	}
	if all.SecretKey == nil {
		return fmt.Errorf("required field secretKey missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.CreateAt == nil {
		return fmt.Errorf("required field createAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"accessKey", "secretKey", "description", "expiredAt", "createAt"})
	} else {
		return err
	}
	o.AccessKey = *all.AccessKey
	o.SecretKey = *all.SecretKey
	o.Description = *all.Description
	o.ExpiredAt = all.ExpiredAt
	o.CreateAt = *all.CreateAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
