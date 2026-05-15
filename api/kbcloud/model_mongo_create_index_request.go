// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoCreateIndexRequest struct {
	// index key pattern
	Keys    interface{}        `json:"keys,omitempty"`
	Options *MongoIndexOptions `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoCreateIndexRequest instantiates a new MongoCreateIndexRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoCreateIndexRequest() *MongoCreateIndexRequest {
	this := MongoCreateIndexRequest{}
	return &this
}

// NewMongoCreateIndexRequestWithDefaults instantiates a new MongoCreateIndexRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoCreateIndexRequestWithDefaults() *MongoCreateIndexRequest {
	this := MongoCreateIndexRequest{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *MongoCreateIndexRequest) GetKeys() interface{} {
	if o == nil || o.Keys == nil {
		var ret interface{}
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateIndexRequest) GetKeysOk() (*interface{}, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return &o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *MongoCreateIndexRequest) HasKeys() bool {
	return o != nil && o.Keys != nil
}

// SetKeys gets a reference to the given interface{} and assigns it to the Keys field.
func (o *MongoCreateIndexRequest) SetKeys(v interface{}) {
	o.Keys = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *MongoCreateIndexRequest) GetOptions() MongoIndexOptions {
	if o == nil || o.Options == nil {
		var ret MongoIndexOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateIndexRequest) GetOptionsOk() (*MongoIndexOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MongoCreateIndexRequest) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given MongoIndexOptions and assigns it to the Options field.
func (o *MongoCreateIndexRequest) SetOptions(v MongoIndexOptions) {
	o.Options = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoCreateIndexRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
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
func (o *MongoCreateIndexRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Keys    interface{}        `json:"keys,omitempty"`
		Options *MongoIndexOptions `json:"options,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"keys", "options"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Keys = all.Keys
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
