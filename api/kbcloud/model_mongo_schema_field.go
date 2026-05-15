// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoSchemaField struct {
	Name       *string  `json:"name,omitempty"`
	Types      []string `json:"types,omitempty"`
	Occurrence *float64 `json:"occurrence,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoSchemaField instantiates a new MongoSchemaField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoSchemaField() *MongoSchemaField {
	this := MongoSchemaField{}
	return &this
}

// NewMongoSchemaFieldWithDefaults instantiates a new MongoSchemaField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoSchemaFieldWithDefaults() *MongoSchemaField {
	this := MongoSchemaField{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MongoSchemaField) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSchemaField) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MongoSchemaField) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MongoSchemaField) SetName(v string) {
	o.Name = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *MongoSchemaField) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSchemaField) GetTypesOk() (*[]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return &o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *MongoSchemaField) HasTypes() bool {
	return o != nil && o.Types != nil
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *MongoSchemaField) SetTypes(v []string) {
	o.Types = v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *MongoSchemaField) GetOccurrence() float64 {
	if o == nil || o.Occurrence == nil {
		var ret float64
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSchemaField) GetOccurrenceOk() (*float64, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *MongoSchemaField) HasOccurrence() bool {
	return o != nil && o.Occurrence != nil
}

// SetOccurrence gets a reference to the given float64 and assigns it to the Occurrence field.
func (o *MongoSchemaField) SetOccurrence(v float64) {
	o.Occurrence = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoSchemaField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoSchemaField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string  `json:"name,omitempty"`
		Types      []string `json:"types,omitempty"`
		Occurrence *float64 `json:"occurrence,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "types", "occurrence"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Types = all.Types
	o.Occurrence = all.Occurrence

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
