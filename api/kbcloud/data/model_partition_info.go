// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package data

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type PartitionInfo struct {
	Id *int32 `json:"id,omitempty"`
	BeginningOffset *int64 `json:"beginningOffset,omitempty"`
	EndOffset *int64 `json:"endOffset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewPartitionInfo instantiates a new PartitionInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPartitionInfo() *PartitionInfo {
	this := PartitionInfo{}
	return &this
}

// NewPartitionInfoWithDefaults instantiates a new PartitionInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPartitionInfoWithDefaults() *PartitionInfo {
	this := PartitionInfo{}
	return &this
}
// GetId returns the Id field value if set, zero value otherwise.
func (o *PartitionInfo) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionInfo) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PartitionInfo) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PartitionInfo) SetId(v int32) {
	o.Id = &v
}


// GetBeginningOffset returns the BeginningOffset field value if set, zero value otherwise.
func (o *PartitionInfo) GetBeginningOffset() int64 {
	if o == nil || o.BeginningOffset == nil {
		var ret int64
		return ret
	}
	return *o.BeginningOffset
}

// GetBeginningOffsetOk returns a tuple with the BeginningOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionInfo) GetBeginningOffsetOk() (*int64, bool) {
	if o == nil || o.BeginningOffset == nil {
		return nil, false
	}
	return o.BeginningOffset, true
}

// HasBeginningOffset returns a boolean if a field has been set.
func (o *PartitionInfo) HasBeginningOffset() bool {
	return o != nil && o.BeginningOffset != nil
}

// SetBeginningOffset gets a reference to the given int64 and assigns it to the BeginningOffset field.
func (o *PartitionInfo) SetBeginningOffset(v int64) {
	o.BeginningOffset = &v
}


// GetEndOffset returns the EndOffset field value if set, zero value otherwise.
func (o *PartitionInfo) GetEndOffset() int64 {
	if o == nil || o.EndOffset == nil {
		var ret int64
		return ret
	}
	return *o.EndOffset
}

// GetEndOffsetOk returns a tuple with the EndOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartitionInfo) GetEndOffsetOk() (*int64, bool) {
	if o == nil || o.EndOffset == nil {
		return nil, false
	}
	return o.EndOffset, true
}

// HasEndOffset returns a boolean if a field has been set.
func (o *PartitionInfo) HasEndOffset() bool {
	return o != nil && o.EndOffset != nil
}

// SetEndOffset gets a reference to the given int64 and assigns it to the EndOffset field.
func (o *PartitionInfo) SetEndOffset(v int64) {
	o.EndOffset = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o PartitionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.BeginningOffset != nil {
		toSerialize["beginningOffset"] = o.BeginningOffset
	}
	if o.EndOffset != nil {
		toSerialize["endOffset"] = o.EndOffset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PartitionInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *int32 `json:"id,omitempty"`
		BeginningOffset *int64 `json:"beginningOffset,omitempty"`
		EndOffset *int64 `json:"endOffset,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "id", "beginningOffset", "endOffset",  })
	} else {
		return err
	}
	o.Id = all.Id
	o.BeginningOffset = all.BeginningOffset
	o.EndOffset = all.EndOffset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
