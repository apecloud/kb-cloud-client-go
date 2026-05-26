// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoWriteResponse struct {
	InsertedCount *int64        `json:"insertedCount,omitempty"`
	InsertedIds   []interface{} `json:"insertedIds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoWriteResponse instantiates a new MongoWriteResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoWriteResponse() *MongoWriteResponse {
	this := MongoWriteResponse{}
	return &this
}

// NewMongoWriteResponseWithDefaults instantiates a new MongoWriteResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoWriteResponseWithDefaults() *MongoWriteResponse {
	this := MongoWriteResponse{}
	return &this
}

// GetInsertedCount returns the InsertedCount field value if set, zero value otherwise.
func (o *MongoWriteResponse) GetInsertedCount() int64 {
	if o == nil || o.InsertedCount == nil {
		var ret int64
		return ret
	}
	return *o.InsertedCount
}

// GetInsertedCountOk returns a tuple with the InsertedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoWriteResponse) GetInsertedCountOk() (*int64, bool) {
	if o == nil || o.InsertedCount == nil {
		return nil, false
	}
	return o.InsertedCount, true
}

// HasInsertedCount returns a boolean if a field has been set.
func (o *MongoWriteResponse) HasInsertedCount() bool {
	return o != nil && o.InsertedCount != nil
}

// SetInsertedCount gets a reference to the given int64 and assigns it to the InsertedCount field.
func (o *MongoWriteResponse) SetInsertedCount(v int64) {
	o.InsertedCount = &v
}

// GetInsertedIds returns the InsertedIds field value if set, zero value otherwise.
func (o *MongoWriteResponse) GetInsertedIds() []interface{} {
	if o == nil || o.InsertedIds == nil {
		var ret []interface{}
		return ret
	}
	return o.InsertedIds
}

// GetInsertedIdsOk returns a tuple with the InsertedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoWriteResponse) GetInsertedIdsOk() (*[]interface{}, bool) {
	if o == nil || o.InsertedIds == nil {
		return nil, false
	}
	return &o.InsertedIds, true
}

// HasInsertedIds returns a boolean if a field has been set.
func (o *MongoWriteResponse) HasInsertedIds() bool {
	return o != nil && o.InsertedIds != nil
}

// SetInsertedIds gets a reference to the given []interface{} and assigns it to the InsertedIds field.
func (o *MongoWriteResponse) SetInsertedIds(v []interface{}) {
	o.InsertedIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoWriteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.InsertedCount != nil {
		toSerialize["insertedCount"] = o.InsertedCount
	}
	if o.InsertedIds != nil {
		toSerialize["insertedIds"] = o.InsertedIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoWriteResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InsertedCount *int64        `json:"insertedCount,omitempty"`
		InsertedIds   []interface{} `json:"insertedIds,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"insertedCount", "insertedIds"})
	} else {
		return err
	}
	o.InsertedCount = all.InsertedCount
	o.InsertedIds = all.InsertedIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
