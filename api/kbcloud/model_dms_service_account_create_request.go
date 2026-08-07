// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DmsServiceAccountCreateRequest MinIO service account access-key create request
type DmsServiceAccountCreateRequest struct {
	// MinIO user that owns the service account.
	UserName string `json:"userName"`
	// Optional fixed access key. If omitted, MinIO generates it.
	AccessKey *string `json:"accessKey,omitempty"`
	// Optional fixed secret key. If omitted, MinIO generates it.
	SecretKey *string `json:"secretKey,omitempty"`
	// Display name of the service account.
	Name *string `json:"name,omitempty"`
	// Description of the service account.
	Description *string `json:"description,omitempty"`
	// Optional service account expiration time.
	Expiration *time.Time `json:"expiration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsServiceAccountCreateRequest instantiates a new DmsServiceAccountCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsServiceAccountCreateRequest(userName string) *DmsServiceAccountCreateRequest {
	this := DmsServiceAccountCreateRequest{}
	this.UserName = userName
	return &this
}

// NewDmsServiceAccountCreateRequestWithDefaults instantiates a new DmsServiceAccountCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsServiceAccountCreateRequestWithDefaults() *DmsServiceAccountCreateRequest {
	this := DmsServiceAccountCreateRequest{}
	return &this
}

// GetUserName returns the UserName field value.
func (o *DmsServiceAccountCreateRequest) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *DmsServiceAccountCreateRequest) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *DmsServiceAccountCreateRequest) SetUserName(v string) {
	o.UserName = v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *DmsServiceAccountCreateRequest) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccountCreateRequest) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *DmsServiceAccountCreateRequest) HasAccessKey() bool {
	return o != nil && o.AccessKey != nil
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *DmsServiceAccountCreateRequest) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *DmsServiceAccountCreateRequest) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccountCreateRequest) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *DmsServiceAccountCreateRequest) HasSecretKey() bool {
	return o != nil && o.SecretKey != nil
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *DmsServiceAccountCreateRequest) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DmsServiceAccountCreateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccountCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DmsServiceAccountCreateRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DmsServiceAccountCreateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DmsServiceAccountCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccountCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DmsServiceAccountCreateRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DmsServiceAccountCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *DmsServiceAccountCreateRequest) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccountCreateRequest) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *DmsServiceAccountCreateRequest) HasExpiration() bool {
	return o != nil && o.Expiration != nil
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *DmsServiceAccountCreateRequest) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsServiceAccountCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["userName"] = o.UserName
	if o.AccessKey != nil {
		toSerialize["accessKey"] = o.AccessKey
	}
	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Expiration != nil {
		if o.Expiration.Nanosecond() == 0 {
			toSerialize["expiration"] = o.Expiration.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiration"] = o.Expiration.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsServiceAccountCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserName    *string    `json:"userName"`
		AccessKey   *string    `json:"accessKey,omitempty"`
		SecretKey   *string    `json:"secretKey,omitempty"`
		Name        *string    `json:"name,omitempty"`
		Description *string    `json:"description,omitempty"`
		Expiration  *time.Time `json:"expiration,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userName", "accessKey", "secretKey", "name", "description", "expiration"})
	} else {
		return err
	}
	o.UserName = *all.UserName
	o.AccessKey = all.AccessKey
	o.SecretKey = all.SecretKey
	o.Name = all.Name
	o.Description = all.Description
	o.Expiration = all.Expiration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
