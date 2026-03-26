// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentCredentialSummary struct {
	// Name of the environment credential
	Name string `json:"name"`
	// Cloud provider of the environment credential
	Provider string `json:"provider"`
	// Type of the environment credential
	CredentialType EnvironmentCredentialType `json:"credentialType"`
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

// NewEnvironmentCredentialSummary instantiates a new EnvironmentCredentialSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCredentialSummary(name string, provider string, credentialType EnvironmentCredentialType, createdAt time.Time, updatedAt time.Time) *EnvironmentCredentialSummary {
	this := EnvironmentCredentialSummary{}
	this.Name = name
	this.Provider = provider
	this.CredentialType = credentialType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewEnvironmentCredentialSummaryWithDefaults instantiates a new EnvironmentCredentialSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentCredentialSummaryWithDefaults() *EnvironmentCredentialSummary {
	this := EnvironmentCredentialSummary{}
	return &this
}

// GetName returns the Name field value.
func (o *EnvironmentCredentialSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EnvironmentCredentialSummary) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value.
func (o *EnvironmentCredentialSummary) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialSummary) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *EnvironmentCredentialSummary) SetProvider(v string) {
	o.Provider = v
}

// GetCredentialType returns the CredentialType field value.
func (o *EnvironmentCredentialSummary) GetCredentialType() EnvironmentCredentialType {
	if o == nil {
		var ret EnvironmentCredentialType
		return ret
	}
	return o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialSummary) GetCredentialTypeOk() (*EnvironmentCredentialType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialType, true
}

// SetCredentialType sets field value.
func (o *EnvironmentCredentialSummary) SetCredentialType(v EnvironmentCredentialType) {
	o.CredentialType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentCredentialSummary) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialSummary) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentCredentialSummary) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentCredentialSummary) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *EnvironmentCredentialSummary) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *EnvironmentCredentialSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *EnvironmentCredentialSummary) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialSummary) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *EnvironmentCredentialSummary) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCredentialSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["provider"] = o.Provider
	toSerialize["credentialType"] = o.CredentialType
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
func (o *EnvironmentCredentialSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string                    `json:"name"`
		Provider       *string                    `json:"provider"`
		CredentialType *EnvironmentCredentialType `json:"credentialType"`
		Description    *string                    `json:"description,omitempty"`
		CreatedAt      *time.Time                 `json:"createdAt"`
		UpdatedAt      *time.Time                 `json:"updatedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.CredentialType == nil {
		return fmt.Errorf("required field credentialType missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "provider", "credentialType", "description", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Provider = *all.Provider
	if !all.CredentialType.IsValid() {
		hasInvalidField = true
	} else {
		o.CredentialType = *all.CredentialType
	}
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
