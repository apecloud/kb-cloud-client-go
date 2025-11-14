// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// StorageProviderSchemaProps the schema properties for storage provider parameters
type StorageProviderSchemaProps struct {
	// default value for the parameter
	Default interface{} `json:"default,omitempty"`
	// description of the parameter
	Description *string `json:"description,omitempty"`
	// type of the parameter
	Type *string `json:"type,omitempty"`
	// whether to hide this parameter from frontend display
	Hidden *bool `json:"hidden,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageProviderSchemaProps instantiates a new StorageProviderSchemaProps object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageProviderSchemaProps() *StorageProviderSchemaProps {
	this := StorageProviderSchemaProps{}
	var hidden bool = false
	this.Hidden = &hidden
	return &this
}

// NewStorageProviderSchemaPropsWithDefaults instantiates a new StorageProviderSchemaProps object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageProviderSchemaPropsWithDefaults() *StorageProviderSchemaProps {
	this := StorageProviderSchemaProps{}
	var hidden bool = false
	this.Hidden = &hidden
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *StorageProviderSchemaProps) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProviderSchemaProps) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *StorageProviderSchemaProps) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *StorageProviderSchemaProps) SetDefault(v interface{}) {
	o.Default = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageProviderSchemaProps) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProviderSchemaProps) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageProviderSchemaProps) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageProviderSchemaProps) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageProviderSchemaProps) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProviderSchemaProps) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageProviderSchemaProps) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageProviderSchemaProps) SetType(v string) {
	o.Type = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *StorageProviderSchemaProps) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProviderSchemaProps) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *StorageProviderSchemaProps) HasHidden() bool {
	return o != nil && o.Hidden != nil
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *StorageProviderSchemaProps) SetHidden(v bool) {
	o.Hidden = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageProviderSchemaProps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageProviderSchemaProps) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default     interface{} `json:"default,omitempty"`
		Description *string     `json:"description,omitempty"`
		Type        *string     `json:"type,omitempty"`
		Hidden      *bool       `json:"hidden,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"default", "description", "type", "hidden"})
	} else {
		return err
	}
	o.Default = all.Default
	o.Description = all.Description
	o.Type = all.Type
	o.Hidden = all.Hidden

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
