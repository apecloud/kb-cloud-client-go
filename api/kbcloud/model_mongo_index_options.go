// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoIndexOptions struct {
	Name   *string `json:"name,omitempty"`
	Unique *bool   `json:"unique,omitempty"`
	Sparse *bool   `json:"sparse,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoIndexOptions instantiates a new MongoIndexOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoIndexOptions() *MongoIndexOptions {
	this := MongoIndexOptions{}
	return &this
}

// NewMongoIndexOptionsWithDefaults instantiates a new MongoIndexOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoIndexOptionsWithDefaults() *MongoIndexOptions {
	this := MongoIndexOptions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MongoIndexOptions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexOptions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MongoIndexOptions) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MongoIndexOptions) SetName(v string) {
	o.Name = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *MongoIndexOptions) GetUnique() bool {
	if o == nil || o.Unique == nil {
		var ret bool
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexOptions) GetUniqueOk() (*bool, bool) {
	if o == nil || o.Unique == nil {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *MongoIndexOptions) HasUnique() bool {
	return o != nil && o.Unique != nil
}

// SetUnique gets a reference to the given bool and assigns it to the Unique field.
func (o *MongoIndexOptions) SetUnique(v bool) {
	o.Unique = &v
}

// GetSparse returns the Sparse field value if set, zero value otherwise.
func (o *MongoIndexOptions) GetSparse() bool {
	if o == nil || o.Sparse == nil {
		var ret bool
		return ret
	}
	return *o.Sparse
}

// GetSparseOk returns a tuple with the Sparse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoIndexOptions) GetSparseOk() (*bool, bool) {
	if o == nil || o.Sparse == nil {
		return nil, false
	}
	return o.Sparse, true
}

// HasSparse returns a boolean if a field has been set.
func (o *MongoIndexOptions) HasSparse() bool {
	return o != nil && o.Sparse != nil
}

// SetSparse gets a reference to the given bool and assigns it to the Sparse field.
func (o *MongoIndexOptions) SetSparse(v bool) {
	o.Sparse = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoIndexOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Unique != nil {
		toSerialize["unique"] = o.Unique
	}
	if o.Sparse != nil {
		toSerialize["sparse"] = o.Sparse
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoIndexOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string `json:"name,omitempty"`
		Unique *bool   `json:"unique,omitempty"`
		Sparse *bool   `json:"sparse,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "unique", "sparse"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Unique = all.Unique
	o.Sparse = all.Sparse

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
