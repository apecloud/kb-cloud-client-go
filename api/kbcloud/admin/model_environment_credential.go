// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentCredential struct {
	// Name of the environment credential
	Name string `json:"name"`
	// The intended usage of the environment credential.
	Usage EnvironmentCredentialUsage `json:"usage"`
	// Key of the environment credential entry, such as accessKeyId for AWS
	Key string `json:"key"`
	// Secret of the environment credential entry, such as secretAccessKey for AWS
	Secret string `json:"secret"`
	// Description of the environment credential
	Description *string `json:"description,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is represented in RFC3339 form and is in UTC.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt is a timestamp representing the server time when this object was updated. It is represented in RFC3339 form and is in UTC.
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentCredential instantiates a new EnvironmentCredential object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCredential(name string, usage EnvironmentCredentialUsage, key string, secret string, createdAt time.Time, updatedAt time.Time) *EnvironmentCredential {
	this := EnvironmentCredential{}
	this.Name = name
	this.Usage = usage
	this.Key = key
	this.Secret = secret
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewEnvironmentCredentialWithDefaults instantiates a new EnvironmentCredential object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentCredentialWithDefaults() *EnvironmentCredential {
	this := EnvironmentCredential{}
	return &this
}

// GetName returns the Name field value.
func (o *EnvironmentCredential) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EnvironmentCredential) SetName(v string) {
	o.Name = v
}

// GetUsage returns the Usage field value.
func (o *EnvironmentCredential) GetUsage() EnvironmentCredentialUsage {
	if o == nil {
		var ret EnvironmentCredentialUsage
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetUsageOk() (*EnvironmentCredentialUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value.
func (o *EnvironmentCredential) SetUsage(v EnvironmentCredentialUsage) {
	o.Usage = v
}

// GetKey returns the Key field value.
func (o *EnvironmentCredential) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *EnvironmentCredential) SetKey(v string) {
	o.Key = v
}

// GetSecret returns the Secret field value.
func (o *EnvironmentCredential) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value.
func (o *EnvironmentCredential) SetSecret(v string) {
	o.Secret = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentCredential) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentCredential) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentCredential) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *EnvironmentCredential) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *EnvironmentCredential) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *EnvironmentCredential) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *EnvironmentCredential) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["usage"] = o.Usage
	toSerialize["key"] = o.Key
	toSerialize["secret"] = o.Secret
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentCredential) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                     `json:"name"`
		Usage       *EnvironmentCredentialUsage `json:"usage"`
		Key         *string                     `json:"key"`
		Secret      *string                     `json:"secret"`
		Description *string                     `json:"description,omitempty"`
		CreatedAt   *time.Time                  `json:"createdAt"`
		UpdatedAt   *time.Time                  `json:"updatedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Usage == nil {
		return fmt.Errorf("required field usage missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Secret == nil {
		return fmt.Errorf("required field secret missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "usage", "key", "secret", "description", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Usage.IsValid() {
		hasInvalidField = true
	} else {
		o.Usage = *all.Usage
	}
	o.Key = *all.Key
	o.Secret = *all.Secret
	o.Description = all.Description
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
