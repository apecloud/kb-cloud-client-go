// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentCredentialCreate struct {
	// Name of the environment credential
	Name string `json:"name"`
	// Cloud provider of the environment credential
	Provider string `json:"provider"`
	// Type of the environment credential
	CredentialType EnvironmentCredentialType `json:"credentialType"`
	// Description of the environment credential
	Description *string                      `json:"description,omitempty"`
	Entries     []EnvironmentCredentialEntry `json:"entries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentCredentialCreate instantiates a new EnvironmentCredentialCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCredentialCreate(name string, provider string, credentialType EnvironmentCredentialType, entries []EnvironmentCredentialEntry) *EnvironmentCredentialCreate {
	this := EnvironmentCredentialCreate{}
	this.Name = name
	this.Provider = provider
	this.CredentialType = credentialType
	this.Entries = entries
	return &this
}

// NewEnvironmentCredentialCreateWithDefaults instantiates a new EnvironmentCredentialCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentCredentialCreateWithDefaults() *EnvironmentCredentialCreate {
	this := EnvironmentCredentialCreate{}
	return &this
}

// GetName returns the Name field value.
func (o *EnvironmentCredentialCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EnvironmentCredentialCreate) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value.
func (o *EnvironmentCredentialCreate) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialCreate) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *EnvironmentCredentialCreate) SetProvider(v string) {
	o.Provider = v
}

// GetCredentialType returns the CredentialType field value.
func (o *EnvironmentCredentialCreate) GetCredentialType() EnvironmentCredentialType {
	if o == nil {
		var ret EnvironmentCredentialType
		return ret
	}
	return o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialCreate) GetCredentialTypeOk() (*EnvironmentCredentialType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredentialType, true
}

// SetCredentialType sets field value.
func (o *EnvironmentCredentialCreate) SetCredentialType(v EnvironmentCredentialType) {
	o.CredentialType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EnvironmentCredentialCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentCredentialCreate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EnvironmentCredentialCreate) SetDescription(v string) {
	o.Description = &v
}

// GetEntries returns the Entries field value.
func (o *EnvironmentCredentialCreate) GetEntries() []EnvironmentCredentialEntry {
	if o == nil {
		var ret []EnvironmentCredentialEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialCreate) GetEntriesOk() (*[]EnvironmentCredentialEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entries, true
}

// SetEntries sets field value.
func (o *EnvironmentCredentialCreate) SetEntries(v []EnvironmentCredentialEntry) {
	o.Entries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCredentialCreate) MarshalJSON() ([]byte, error) {
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
	toSerialize["entries"] = o.Entries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentCredentialCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string                       `json:"name"`
		Provider       *string                       `json:"provider"`
		CredentialType *EnvironmentCredentialType    `json:"credentialType"`
		Description    *string                       `json:"description,omitempty"`
		Entries        *[]EnvironmentCredentialEntry `json:"entries"`
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
	if all.Entries == nil {
		return fmt.Errorf("required field entries missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "provider", "credentialType", "description", "entries"})
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
	o.Entries = *all.Entries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
