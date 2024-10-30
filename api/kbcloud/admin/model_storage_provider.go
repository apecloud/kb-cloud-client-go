// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// StorageProvider Storage Provider comprises specifications that provide guidance accessing remote storage. 
type StorageProvider struct {
	// Unique identifier for the storage provider
	Id *string `json:"id,omitempty"`
	// Name of the storage provider
	Name *string `json:"name,omitempty"`
	// defines which parameters are credential fields
	Credential []string `json:"credential,omitempty"`
	// the schema for the storage provider for creating a storage
	Schema map[string]StorageProviderSchemaProps `json:"schema,omitempty"`
	// defines which parameters are required
	Required []string `json:"required,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewStorageProvider instantiates a new StorageProvider object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageProvider() *StorageProvider {
	this := StorageProvider{}
	return &this
}

// NewStorageProviderWithDefaults instantiates a new StorageProvider object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageProviderWithDefaults() *StorageProvider {
	this := StorageProvider{}
	return &this
}
// GetId returns the Id field value if set, zero value otherwise.
func (o *StorageProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StorageProvider) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StorageProvider) SetId(v string) {
	o.Id = &v
}


// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageProvider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageProvider) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageProvider) SetName(v string) {
	o.Name = &v
}


// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *StorageProvider) GetCredential() []string {
	if o == nil || o.Credential == nil {
		var ret []string
		return ret
	}
	return o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetCredentialOk() (*[]string, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return &o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *StorageProvider) HasCredential() bool {
	return o != nil && o.Credential != nil
}

// SetCredential gets a reference to the given []string and assigns it to the Credential field.
func (o *StorageProvider) SetCredential(v []string) {
	o.Credential = v
}


// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *StorageProvider) GetSchema() map[string]StorageProviderSchemaProps {
	if o == nil || o.Schema == nil {
		var ret map[string]StorageProviderSchemaProps
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetSchemaOk() (*map[string]StorageProviderSchemaProps, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return &o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *StorageProvider) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given map[string]StorageProviderSchemaProps and assigns it to the Schema field.
func (o *StorageProvider) SetSchema(v map[string]StorageProviderSchemaProps) {
	o.Schema = v
}


// GetRequired returns the Required field value if set, zero value otherwise.
func (o *StorageProvider) GetRequired() []string {
	if o == nil || o.Required == nil {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetRequiredOk() (*[]string, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return &o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *StorageProvider) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *StorageProvider) SetRequired(v []string) {
	o.Required = v
}



// MarshalJSON serializes the struct using spec logic.
func (o StorageProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageProvider) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
		Credential []string `json:"credential,omitempty"`
		Schema map[string]StorageProviderSchemaProps `json:"schema,omitempty"`
		Required []string `json:"required,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "id", "name", "credential", "schema", "required",  })
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.Credential = all.Credential
	o.Schema = all.Schema
	o.Required = all.Required

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
