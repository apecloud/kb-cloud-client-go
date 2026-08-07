// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DmsServiceAccount MinIO service account access-key
type DmsServiceAccount struct {
	// MinIO user that owns the service account.
	UserName *string `json:"userName,omitempty"`
	// Service account access key.
	AccessKey *string `json:"accessKey,omitempty"`
	// Secret key returned only by create.
	SecretKey *string `json:"secretKey,omitempty"`
	// Service account status.
	Status *string `json:"status,omitempty"`
	// Display name of the service account.
	Name *string `json:"name,omitempty"`
	// Description of the service account.
	Description *string `json:"description,omitempty"`
	// Service account expiration time.
	Expiration *time.Time `json:"expiration,omitempty"`
	// Service account policy text.
	Policy *string `json:"policy,omitempty"`
	// Whether the policy is implied from the parent user.
	ImpliedPolicy *bool `json:"impliedPolicy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsServiceAccount instantiates a new DmsServiceAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsServiceAccount() *DmsServiceAccount {
	this := DmsServiceAccount{}
	return &this
}

// NewDmsServiceAccountWithDefaults instantiates a new DmsServiceAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsServiceAccountWithDefaults() *DmsServiceAccount {
	this := DmsServiceAccount{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DmsServiceAccount) SetUserName(v string) {
	o.UserName = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasAccessKey() bool {
	return o != nil && o.AccessKey != nil
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *DmsServiceAccount) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasSecretKey() bool {
	return o != nil && o.SecretKey != nil
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *DmsServiceAccount) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DmsServiceAccount) SetStatus(v string) {
	o.Status = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DmsServiceAccount) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DmsServiceAccount) SetDescription(v string) {
	o.Description = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasExpiration() bool {
	return o != nil && o.Expiration != nil
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *DmsServiceAccount) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetPolicy() string {
	if o == nil || o.Policy == nil {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetPolicyOk() (*string, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasPolicy() bool {
	return o != nil && o.Policy != nil
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *DmsServiceAccount) SetPolicy(v string) {
	o.Policy = &v
}

// GetImpliedPolicy returns the ImpliedPolicy field value if set, zero value otherwise.
func (o *DmsServiceAccount) GetImpliedPolicy() bool {
	if o == nil || o.ImpliedPolicy == nil {
		var ret bool
		return ret
	}
	return *o.ImpliedPolicy
}

// GetImpliedPolicyOk returns a tuple with the ImpliedPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsServiceAccount) GetImpliedPolicyOk() (*bool, bool) {
	if o == nil || o.ImpliedPolicy == nil {
		return nil, false
	}
	return o.ImpliedPolicy, true
}

// HasImpliedPolicy returns a boolean if a field has been set.
func (o *DmsServiceAccount) HasImpliedPolicy() bool {
	return o != nil && o.ImpliedPolicy != nil
}

// SetImpliedPolicy gets a reference to the given bool and assigns it to the ImpliedPolicy field.
func (o *DmsServiceAccount) SetImpliedPolicy(v bool) {
	o.ImpliedPolicy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.AccessKey != nil {
		toSerialize["accessKey"] = o.AccessKey
	}
	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
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
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	if o.ImpliedPolicy != nil {
		toSerialize["impliedPolicy"] = o.ImpliedPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsServiceAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserName      *string    `json:"userName,omitempty"`
		AccessKey     *string    `json:"accessKey,omitempty"`
		SecretKey     *string    `json:"secretKey,omitempty"`
		Status        *string    `json:"status,omitempty"`
		Name          *string    `json:"name,omitempty"`
		Description   *string    `json:"description,omitempty"`
		Expiration    *time.Time `json:"expiration,omitempty"`
		Policy        *string    `json:"policy,omitempty"`
		ImpliedPolicy *bool      `json:"impliedPolicy,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userName", "accessKey", "secretKey", "status", "name", "description", "expiration", "policy", "impliedPolicy"})
	} else {
		return err
	}
	o.UserName = all.UserName
	o.AccessKey = all.AccessKey
	o.SecretKey = all.SecretKey
	o.Status = all.Status
	o.Name = all.Name
	o.Description = all.Description
	o.Expiration = all.Expiration
	o.Policy = all.Policy
	o.ImpliedPolicy = all.ImpliedPolicy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
