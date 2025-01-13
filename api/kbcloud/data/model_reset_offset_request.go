// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package data

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type ResetOffsetRequest struct {
	// the partition to reset
	Partition *int32 `json:"partition,omitempty"`
	Seek *OffsetResetStrategy `json:"seek,omitempty"`
	// the offset to reset to
	Offset *int64 `json:"offset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewResetOffsetRequest instantiates a new ResetOffsetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResetOffsetRequest() *ResetOffsetRequest {
	this := ResetOffsetRequest{}
	return &this
}

// NewResetOffsetRequestWithDefaults instantiates a new ResetOffsetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResetOffsetRequestWithDefaults() *ResetOffsetRequest {
	this := ResetOffsetRequest{}
	return &this
}
// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *ResetOffsetRequest) GetPartition() int32 {
	if o == nil || o.Partition == nil {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetOffsetRequest) GetPartitionOk() (*int32, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *ResetOffsetRequest) HasPartition() bool {
	return o != nil && o.Partition != nil
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *ResetOffsetRequest) SetPartition(v int32) {
	o.Partition = &v
}


// GetSeek returns the Seek field value if set, zero value otherwise.
func (o *ResetOffsetRequest) GetSeek() OffsetResetStrategy {
	if o == nil || o.Seek == nil {
		var ret OffsetResetStrategy
		return ret
	}
	return *o.Seek
}

// GetSeekOk returns a tuple with the Seek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetOffsetRequest) GetSeekOk() (*OffsetResetStrategy, bool) {
	if o == nil || o.Seek == nil {
		return nil, false
	}
	return o.Seek, true
}

// HasSeek returns a boolean if a field has been set.
func (o *ResetOffsetRequest) HasSeek() bool {
	return o != nil && o.Seek != nil
}

// SetSeek gets a reference to the given OffsetResetStrategy and assigns it to the Seek field.
func (o *ResetOffsetRequest) SetSeek(v OffsetResetStrategy) {
	o.Seek = &v
}


// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ResetOffsetRequest) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetOffsetRequest) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ResetOffsetRequest) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ResetOffsetRequest) SetOffset(v int64) {
	o.Offset = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o ResetOffsetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
	}
	if o.Seek != nil {
		toSerialize["seek"] = o.Seek
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResetOffsetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Partition *int32 `json:"partition,omitempty"`
		Seek *OffsetResetStrategy `json:"seek,omitempty"`
		Offset *int64 `json:"offset,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "partition", "seek", "offset",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.Partition = all.Partition
	if all.Seek != nil &&!all.Seek.IsValid() {
		hasInvalidField = true
	} else {
		o.Seek = all.Seek
	}
	o.Offset = all.Offset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
