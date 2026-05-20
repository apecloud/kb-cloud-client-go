// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoCreateCollectionRequest struct {
	Options *MongoCreateCollectionOptions `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoCreateCollectionRequest instantiates a new MongoCreateCollectionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoCreateCollectionRequest() *MongoCreateCollectionRequest {
	this := MongoCreateCollectionRequest{}
	return &this
}

// NewMongoCreateCollectionRequestWithDefaults instantiates a new MongoCreateCollectionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoCreateCollectionRequestWithDefaults() *MongoCreateCollectionRequest {
	this := MongoCreateCollectionRequest{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *MongoCreateCollectionRequest) GetOptions() MongoCreateCollectionOptions {
	if o == nil || o.Options == nil {
		var ret MongoCreateCollectionOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateCollectionRequest) GetOptionsOk() (*MongoCreateCollectionOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MongoCreateCollectionRequest) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given MongoCreateCollectionOptions and assigns it to the Options field.
func (o *MongoCreateCollectionRequest) SetOptions(v MongoCreateCollectionOptions) {
	o.Options = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoCreateCollectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoCreateCollectionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Options *MongoCreateCollectionOptions `json:"options,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"options"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
