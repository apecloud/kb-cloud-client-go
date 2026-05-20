// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoCollectionCapabilities struct {
	CanReadDocuments  *bool `json:"canReadDocuments,omitempty"`
	CanWriteDocuments *bool `json:"canWriteDocuments,omitempty"`
	CanManageIndexes  *bool `json:"canManageIndexes,omitempty"`
	CanSetValidation  *bool `json:"canSetValidation,omitempty"`
	CanDrop           *bool `json:"canDrop,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoCollectionCapabilities instantiates a new MongoCollectionCapabilities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoCollectionCapabilities() *MongoCollectionCapabilities {
	this := MongoCollectionCapabilities{}
	return &this
}

// NewMongoCollectionCapabilitiesWithDefaults instantiates a new MongoCollectionCapabilities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoCollectionCapabilitiesWithDefaults() *MongoCollectionCapabilities {
	this := MongoCollectionCapabilities{}
	return &this
}

// GetCanReadDocuments returns the CanReadDocuments field value if set, zero value otherwise.
func (o *MongoCollectionCapabilities) GetCanReadDocuments() bool {
	if o == nil || o.CanReadDocuments == nil {
		var ret bool
		return ret
	}
	return *o.CanReadDocuments
}

// GetCanReadDocumentsOk returns a tuple with the CanReadDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionCapabilities) GetCanReadDocumentsOk() (*bool, bool) {
	if o == nil || o.CanReadDocuments == nil {
		return nil, false
	}
	return o.CanReadDocuments, true
}

// HasCanReadDocuments returns a boolean if a field has been set.
func (o *MongoCollectionCapabilities) HasCanReadDocuments() bool {
	return o != nil && o.CanReadDocuments != nil
}

// SetCanReadDocuments gets a reference to the given bool and assigns it to the CanReadDocuments field.
func (o *MongoCollectionCapabilities) SetCanReadDocuments(v bool) {
	o.CanReadDocuments = &v
}

// GetCanWriteDocuments returns the CanWriteDocuments field value if set, zero value otherwise.
func (o *MongoCollectionCapabilities) GetCanWriteDocuments() bool {
	if o == nil || o.CanWriteDocuments == nil {
		var ret bool
		return ret
	}
	return *o.CanWriteDocuments
}

// GetCanWriteDocumentsOk returns a tuple with the CanWriteDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionCapabilities) GetCanWriteDocumentsOk() (*bool, bool) {
	if o == nil || o.CanWriteDocuments == nil {
		return nil, false
	}
	return o.CanWriteDocuments, true
}

// HasCanWriteDocuments returns a boolean if a field has been set.
func (o *MongoCollectionCapabilities) HasCanWriteDocuments() bool {
	return o != nil && o.CanWriteDocuments != nil
}

// SetCanWriteDocuments gets a reference to the given bool and assigns it to the CanWriteDocuments field.
func (o *MongoCollectionCapabilities) SetCanWriteDocuments(v bool) {
	o.CanWriteDocuments = &v
}

// GetCanManageIndexes returns the CanManageIndexes field value if set, zero value otherwise.
func (o *MongoCollectionCapabilities) GetCanManageIndexes() bool {
	if o == nil || o.CanManageIndexes == nil {
		var ret bool
		return ret
	}
	return *o.CanManageIndexes
}

// GetCanManageIndexesOk returns a tuple with the CanManageIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionCapabilities) GetCanManageIndexesOk() (*bool, bool) {
	if o == nil || o.CanManageIndexes == nil {
		return nil, false
	}
	return o.CanManageIndexes, true
}

// HasCanManageIndexes returns a boolean if a field has been set.
func (o *MongoCollectionCapabilities) HasCanManageIndexes() bool {
	return o != nil && o.CanManageIndexes != nil
}

// SetCanManageIndexes gets a reference to the given bool and assigns it to the CanManageIndexes field.
func (o *MongoCollectionCapabilities) SetCanManageIndexes(v bool) {
	o.CanManageIndexes = &v
}

// GetCanSetValidation returns the CanSetValidation field value if set, zero value otherwise.
func (o *MongoCollectionCapabilities) GetCanSetValidation() bool {
	if o == nil || o.CanSetValidation == nil {
		var ret bool
		return ret
	}
	return *o.CanSetValidation
}

// GetCanSetValidationOk returns a tuple with the CanSetValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionCapabilities) GetCanSetValidationOk() (*bool, bool) {
	if o == nil || o.CanSetValidation == nil {
		return nil, false
	}
	return o.CanSetValidation, true
}

// HasCanSetValidation returns a boolean if a field has been set.
func (o *MongoCollectionCapabilities) HasCanSetValidation() bool {
	return o != nil && o.CanSetValidation != nil
}

// SetCanSetValidation gets a reference to the given bool and assigns it to the CanSetValidation field.
func (o *MongoCollectionCapabilities) SetCanSetValidation(v bool) {
	o.CanSetValidation = &v
}

// GetCanDrop returns the CanDrop field value if set, zero value otherwise.
func (o *MongoCollectionCapabilities) GetCanDrop() bool {
	if o == nil || o.CanDrop == nil {
		var ret bool
		return ret
	}
	return *o.CanDrop
}

// GetCanDropOk returns a tuple with the CanDrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionCapabilities) GetCanDropOk() (*bool, bool) {
	if o == nil || o.CanDrop == nil {
		return nil, false
	}
	return o.CanDrop, true
}

// HasCanDrop returns a boolean if a field has been set.
func (o *MongoCollectionCapabilities) HasCanDrop() bool {
	return o != nil && o.CanDrop != nil
}

// SetCanDrop gets a reference to the given bool and assigns it to the CanDrop field.
func (o *MongoCollectionCapabilities) SetCanDrop(v bool) {
	o.CanDrop = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoCollectionCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CanReadDocuments != nil {
		toSerialize["canReadDocuments"] = o.CanReadDocuments
	}
	if o.CanWriteDocuments != nil {
		toSerialize["canWriteDocuments"] = o.CanWriteDocuments
	}
	if o.CanManageIndexes != nil {
		toSerialize["canManageIndexes"] = o.CanManageIndexes
	}
	if o.CanSetValidation != nil {
		toSerialize["canSetValidation"] = o.CanSetValidation
	}
	if o.CanDrop != nil {
		toSerialize["canDrop"] = o.CanDrop
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoCollectionCapabilities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CanReadDocuments  *bool `json:"canReadDocuments,omitempty"`
		CanWriteDocuments *bool `json:"canWriteDocuments,omitempty"`
		CanManageIndexes  *bool `json:"canManageIndexes,omitempty"`
		CanSetValidation  *bool `json:"canSetValidation,omitempty"`
		CanDrop           *bool `json:"canDrop,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"canReadDocuments", "canWriteDocuments", "canManageIndexes", "canSetValidation", "canDrop"})
	} else {
		return err
	}
	o.CanReadDocuments = all.CanReadDocuments
	o.CanWriteDocuments = all.CanWriteDocuments
	o.CanManageIndexes = all.CanManageIndexes
	o.CanSetValidation = all.CanSetValidation
	o.CanDrop = all.CanDrop

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
