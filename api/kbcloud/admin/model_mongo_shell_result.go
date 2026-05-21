// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoShellResult struct {
	// mongosh result type
	Type interface{} `json:"type,omitempty"`
	// printable representation
	Printable interface{} `json:"printable,omitempty"`
	// original source value when available
	Source interface{} `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoShellResult instantiates a new MongoShellResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoShellResult() *MongoShellResult {
	this := MongoShellResult{}
	return &this
}

// NewMongoShellResultWithDefaults instantiates a new MongoShellResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoShellResultWithDefaults() *MongoShellResult {
	this := MongoShellResult{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MongoShellResult) GetType() interface{} {
	if o == nil || o.Type == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellResult) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MongoShellResult) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *MongoShellResult) SetType(v interface{}) {
	o.Type = v
}

// GetPrintable returns the Printable field value if set, zero value otherwise.
func (o *MongoShellResult) GetPrintable() interface{} {
	if o == nil || o.Printable == nil {
		var ret interface{}
		return ret
	}
	return o.Printable
}

// GetPrintableOk returns a tuple with the Printable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellResult) GetPrintableOk() (*interface{}, bool) {
	if o == nil || o.Printable == nil {
		return nil, false
	}
	return &o.Printable, true
}

// HasPrintable returns a boolean if a field has been set.
func (o *MongoShellResult) HasPrintable() bool {
	return o != nil && o.Printable != nil
}

// SetPrintable gets a reference to the given interface{} and assigns it to the Printable field.
func (o *MongoShellResult) SetPrintable(v interface{}) {
	o.Printable = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MongoShellResult) GetSource() interface{} {
	if o == nil || o.Source == nil {
		var ret interface{}
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellResult) GetSourceOk() (*interface{}, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MongoShellResult) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given interface{} and assigns it to the Source field.
func (o *MongoShellResult) SetSource(v interface{}) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoShellResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Printable != nil {
		toSerialize["printable"] = o.Printable
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoShellResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type      interface{} `json:"type,omitempty"`
		Printable interface{} `json:"printable,omitempty"`
		Source    interface{} `json:"source,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "printable", "source"})
	} else {
		return err
	}
	o.Type = all.Type
	o.Printable = all.Printable
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
