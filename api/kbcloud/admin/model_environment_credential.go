// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentCredential struct {
	// Name of the environment credential
	Name *string `json:"name,omitempty"`
	// Cloud provider of the environment credential
	Provider *string `json:"provider,omitempty"`
	// Type of the environment credential
	CredentialType *EnvironmentCredentialType `json:"credentialType,omitempty"`
	// Description of the environment credential
	Description *string `json:"description,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is represented in RFC3339 form and is in UTC.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UpdatedAt is a timestamp representing the server time when this object was updated. It is represented in RFC3339 form and is in UTC.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Entries of the environment credential
	Entries []EnvironmentCredentialEntry `json:"entries,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentCredential instantiates a new EnvironmentCredential object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCredential() *EnvironmentCredential {
	this := EnvironmentCredential{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentCredential) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentCredential) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentCredential) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *EnvironmentCredential) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *EnvironmentCredential) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *EnvironmentCredential) SetProvider(v string) {
	o.Provider = &v
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise.
func (o *EnvironmentCredential) GetCredentialType() EnvironmentCredentialType {
	if o == nil || o.CredentialType == nil {
		var ret EnvironmentCredentialType
		return ret
	}
	return *o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetCredentialTypeOk() (*EnvironmentCredentialType, bool) {
	if o == nil || o.CredentialType == nil {
		return nil, false
	}
	return o.CredentialType, true
}

// HasCredentialType returns a boolean if a field has been set.
func (o *EnvironmentCredential) HasCredentialType() bool {
	return o != nil && o.CredentialType != nil
}

// SetCredentialType gets a reference to the given EnvironmentCredentialType and assigns it to the CredentialType field.
func (o *EnvironmentCredential) SetCredentialType(v EnvironmentCredentialType) {
	o.CredentialType = &v
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

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EnvironmentCredential) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EnvironmentCredential) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EnvironmentCredential) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentCredential) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentCredential) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentCredential) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *EnvironmentCredential) GetEntries() []EnvironmentCredentialEntry {
	if o == nil || o.Entries == nil {
		var ret []EnvironmentCredentialEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredential) GetEntriesOk() (*[]EnvironmentCredentialEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return &o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *EnvironmentCredential) HasEntries() bool {
	return o != nil && o.Entries != nil
}

// SetEntries gets a reference to the given []EnvironmentCredentialEntry and assigns it to the Entries field.
func (o *EnvironmentCredential) SetEntries(v []EnvironmentCredentialEntry) {
	o.Entries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.CredentialType != nil {
		toSerialize["credentialType"] = o.CredentialType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentCredential) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string                      `json:"name,omitempty"`
		Provider       *string                      `json:"provider,omitempty"`
		CredentialType *EnvironmentCredentialType   `json:"credentialType,omitempty"`
		Description    *string                      `json:"description,omitempty"`
		CreatedAt      *time.Time                   `json:"createdAt,omitempty"`
		UpdatedAt      *time.Time                   `json:"updatedAt,omitempty"`
		Entries        []EnvironmentCredentialEntry `json:"entries,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "provider", "credentialType", "description", "createdAt", "updatedAt", "entries"})
	} else {
		return err
	}
	hasInvalidField := false
	o.Name = all.Name
	o.Provider = all.Provider
	if all.CredentialType != nil && !all.CredentialType.IsValid() {
		hasInvalidField = true
	} else {
		o.CredentialType = all.CredentialType
	}
	o.Description = all.Description
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.Entries = all.Entries
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}
	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
