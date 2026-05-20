// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoCreateCollectionOptions struct {
	Capped *bool `json:"capped,omitempty"`
	// capped collection size in bytes
	Size *int64 `json:"size,omitempty"`
	// capped collection max document count
	Max *int64 `json:"max,omitempty"`
	// collection validator document
	Validator interface{} `json:"validator,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoCreateCollectionOptions instantiates a new MongoCreateCollectionOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoCreateCollectionOptions() *MongoCreateCollectionOptions {
	this := MongoCreateCollectionOptions{}
	return &this
}

// NewMongoCreateCollectionOptionsWithDefaults instantiates a new MongoCreateCollectionOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoCreateCollectionOptionsWithDefaults() *MongoCreateCollectionOptions {
	this := MongoCreateCollectionOptions{}
	return &this
}

// GetCapped returns the Capped field value if set, zero value otherwise.
func (o *MongoCreateCollectionOptions) GetCapped() bool {
	if o == nil || o.Capped == nil {
		var ret bool
		return ret
	}
	return *o.Capped
}

// GetCappedOk returns a tuple with the Capped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateCollectionOptions) GetCappedOk() (*bool, bool) {
	if o == nil || o.Capped == nil {
		return nil, false
	}
	return o.Capped, true
}

// HasCapped returns a boolean if a field has been set.
func (o *MongoCreateCollectionOptions) HasCapped() bool {
	return o != nil && o.Capped != nil
}

// SetCapped gets a reference to the given bool and assigns it to the Capped field.
func (o *MongoCreateCollectionOptions) SetCapped(v bool) {
	o.Capped = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MongoCreateCollectionOptions) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateCollectionOptions) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MongoCreateCollectionOptions) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *MongoCreateCollectionOptions) SetSize(v int64) {
	o.Size = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *MongoCreateCollectionOptions) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateCollectionOptions) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *MongoCreateCollectionOptions) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *MongoCreateCollectionOptions) SetMax(v int64) {
	o.Max = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *MongoCreateCollectionOptions) GetValidator() interface{} {
	if o == nil || o.Validator == nil {
		var ret interface{}
		return ret
	}
	return o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateCollectionOptions) GetValidatorOk() (*interface{}, bool) {
	if o == nil || o.Validator == nil {
		return nil, false
	}
	return &o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *MongoCreateCollectionOptions) HasValidator() bool {
	return o != nil && o.Validator != nil
}

// SetValidator gets a reference to the given interface{} and assigns it to the Validator field.
func (o *MongoCreateCollectionOptions) SetValidator(v interface{}) {
	o.Validator = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoCreateCollectionOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Capped != nil {
		toSerialize["capped"] = o.Capped
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Validator != nil {
		toSerialize["validator"] = o.Validator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoCreateCollectionOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Capped    *bool       `json:"capped,omitempty"`
		Size      *int64      `json:"size,omitempty"`
		Max       *int64      `json:"max,omitempty"`
		Validator interface{} `json:"validator,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"capped", "size", "max", "validator"})
	} else {
		return err
	}
	o.Capped = all.Capped
	o.Size = all.Size
	o.Max = all.Max
	o.Validator = all.Validator

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
