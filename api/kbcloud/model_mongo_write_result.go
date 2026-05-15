// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoWriteResult struct {
	MatchedCount  *int64 `json:"matchedCount,omitempty"`
	ModifiedCount *int64 `json:"modifiedCount,omitempty"`
	DeletedCount  *int64 `json:"deletedCount,omitempty"`
	InsertedCount *int64 `json:"insertedCount,omitempty"`
	UpsertedCount *int64 `json:"upsertedCount,omitempty"`
	// upserted document id
	UpsertedId interface{} `json:"upsertedId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoWriteResult instantiates a new MongoWriteResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoWriteResult() *MongoWriteResult {
	this := MongoWriteResult{}
	return &this
}

// NewMongoWriteResultWithDefaults instantiates a new MongoWriteResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoWriteResultWithDefaults() *MongoWriteResult {
	this := MongoWriteResult{}
	return &this
}

// GetMatchedCount returns the MatchedCount field value if set, zero value otherwise.
func (o *MongoWriteResult) GetMatchedCount() int64 {
	if o == nil || o.MatchedCount == nil {
		var ret int64
		return ret
	}
	return *o.MatchedCount
}

// GetMatchedCountOk returns a tuple with the MatchedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoWriteResult) GetMatchedCountOk() (*int64, bool) {
	if o == nil || o.MatchedCount == nil {
		return nil, false
	}
	return o.MatchedCount, true
}

// HasMatchedCount returns a boolean if a field has been set.
func (o *MongoWriteResult) HasMatchedCount() bool {
	return o != nil && o.MatchedCount != nil
}

// SetMatchedCount gets a reference to the given int64 and assigns it to the MatchedCount field.
func (o *MongoWriteResult) SetMatchedCount(v int64) {
	o.MatchedCount = &v
}

// GetModifiedCount returns the ModifiedCount field value if set, zero value otherwise.
func (o *MongoWriteResult) GetModifiedCount() int64 {
	if o == nil || o.ModifiedCount == nil {
		var ret int64
		return ret
	}
	return *o.ModifiedCount
}

// GetModifiedCountOk returns a tuple with the ModifiedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoWriteResult) GetModifiedCountOk() (*int64, bool) {
	if o == nil || o.ModifiedCount == nil {
		return nil, false
	}
	return o.ModifiedCount, true
}

// HasModifiedCount returns a boolean if a field has been set.
func (o *MongoWriteResult) HasModifiedCount() bool {
	return o != nil && o.ModifiedCount != nil
}

// SetModifiedCount gets a reference to the given int64 and assigns it to the ModifiedCount field.
func (o *MongoWriteResult) SetModifiedCount(v int64) {
	o.ModifiedCount = &v
}

// GetDeletedCount returns the DeletedCount field value if set, zero value otherwise.
func (o *MongoWriteResult) GetDeletedCount() int64 {
	if o == nil || o.DeletedCount == nil {
		var ret int64
		return ret
	}
	return *o.DeletedCount
}

// GetDeletedCountOk returns a tuple with the DeletedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoWriteResult) GetDeletedCountOk() (*int64, bool) {
	if o == nil || o.DeletedCount == nil {
		return nil, false
	}
	return o.DeletedCount, true
}

// HasDeletedCount returns a boolean if a field has been set.
func (o *MongoWriteResult) HasDeletedCount() bool {
	return o != nil && o.DeletedCount != nil
}

// SetDeletedCount gets a reference to the given int64 and assigns it to the DeletedCount field.
func (o *MongoWriteResult) SetDeletedCount(v int64) {
	o.DeletedCount = &v
}

// GetInsertedCount returns the InsertedCount field value if set, zero value otherwise.
func (o *MongoWriteResult) GetInsertedCount() int64 {
	if o == nil || o.InsertedCount == nil {
		var ret int64
		return ret
	}
	return *o.InsertedCount
}

// GetInsertedCountOk returns a tuple with the InsertedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoWriteResult) GetInsertedCountOk() (*int64, bool) {
	if o == nil || o.InsertedCount == nil {
		return nil, false
	}
	return o.InsertedCount, true
}

// HasInsertedCount returns a boolean if a field has been set.
func (o *MongoWriteResult) HasInsertedCount() bool {
	return o != nil && o.InsertedCount != nil
}

// SetInsertedCount gets a reference to the given int64 and assigns it to the InsertedCount field.
func (o *MongoWriteResult) SetInsertedCount(v int64) {
	o.InsertedCount = &v
}

// GetUpsertedCount returns the UpsertedCount field value if set, zero value otherwise.
func (o *MongoWriteResult) GetUpsertedCount() int64 {
	if o == nil || o.UpsertedCount == nil {
		var ret int64
		return ret
	}
	return *o.UpsertedCount
}

// GetUpsertedCountOk returns a tuple with the UpsertedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoWriteResult) GetUpsertedCountOk() (*int64, bool) {
	if o == nil || o.UpsertedCount == nil {
		return nil, false
	}
	return o.UpsertedCount, true
}

// HasUpsertedCount returns a boolean if a field has been set.
func (o *MongoWriteResult) HasUpsertedCount() bool {
	return o != nil && o.UpsertedCount != nil
}

// SetUpsertedCount gets a reference to the given int64 and assigns it to the UpsertedCount field.
func (o *MongoWriteResult) SetUpsertedCount(v int64) {
	o.UpsertedCount = &v
}

// GetUpsertedId returns the UpsertedId field value if set, zero value otherwise.
func (o *MongoWriteResult) GetUpsertedId() interface{} {
	if o == nil || o.UpsertedId == nil {
		var ret interface{}
		return ret
	}
	return o.UpsertedId
}

// GetUpsertedIdOk returns a tuple with the UpsertedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoWriteResult) GetUpsertedIdOk() (*interface{}, bool) {
	if o == nil || o.UpsertedId == nil {
		return nil, false
	}
	return &o.UpsertedId, true
}

// HasUpsertedId returns a boolean if a field has been set.
func (o *MongoWriteResult) HasUpsertedId() bool {
	return o != nil && o.UpsertedId != nil
}

// SetUpsertedId gets a reference to the given interface{} and assigns it to the UpsertedId field.
func (o *MongoWriteResult) SetUpsertedId(v interface{}) {
	o.UpsertedId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoWriteResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MatchedCount != nil {
		toSerialize["matchedCount"] = o.MatchedCount
	}
	if o.ModifiedCount != nil {
		toSerialize["modifiedCount"] = o.ModifiedCount
	}
	if o.DeletedCount != nil {
		toSerialize["deletedCount"] = o.DeletedCount
	}
	if o.InsertedCount != nil {
		toSerialize["insertedCount"] = o.InsertedCount
	}
	if o.UpsertedCount != nil {
		toSerialize["upsertedCount"] = o.UpsertedCount
	}
	if o.UpsertedId != nil {
		toSerialize["upsertedId"] = o.UpsertedId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoWriteResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MatchedCount  *int64      `json:"matchedCount,omitempty"`
		ModifiedCount *int64      `json:"modifiedCount,omitempty"`
		DeletedCount  *int64      `json:"deletedCount,omitempty"`
		InsertedCount *int64      `json:"insertedCount,omitempty"`
		UpsertedCount *int64      `json:"upsertedCount,omitempty"`
		UpsertedId    interface{} `json:"upsertedId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"matchedCount", "modifiedCount", "deletedCount", "insertedCount", "upsertedCount", "upsertedId"})
	} else {
		return err
	}
	o.MatchedCount = all.MatchedCount
	o.ModifiedCount = all.ModifiedCount
	o.DeletedCount = all.DeletedCount
	o.InsertedCount = all.InsertedCount
	o.UpsertedCount = all.UpsertedCount
	o.UpsertedId = all.UpsertedId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
