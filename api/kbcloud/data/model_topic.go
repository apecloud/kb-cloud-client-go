// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package data

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Topic struct {
	Name                string `json:"name"`
	PartitionsCount     *int32 `json:"partitionsCount,omitempty"`
	ConsumerGroupsCount *int32 `json:"consumerGroupsCount,omitempty"`
	ReplicaCount        *int32 `json:"replicaCount,omitempty"`
	TotalLogSize        *int64 `json:"TotalLogSize,omitempty"`
	// show if the topic is internal topic
	Internal *bool `json:"internal,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopic instantiates a new Topic object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopic(name string) *Topic {
	this := Topic{}
	this.Name = name
	return &this
}

// NewTopicWithDefaults instantiates a new Topic object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopicWithDefaults() *Topic {
	this := Topic{}
	return &this
}

// GetName returns the Name field value.
func (o *Topic) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Topic) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Topic) SetName(v string) {
	o.Name = v
}

// GetPartitionsCount returns the PartitionsCount field value if set, zero value otherwise.
func (o *Topic) GetPartitionsCount() int32 {
	if o == nil || o.PartitionsCount == nil {
		var ret int32
		return ret
	}
	return *o.PartitionsCount
}

// GetPartitionsCountOk returns a tuple with the PartitionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetPartitionsCountOk() (*int32, bool) {
	if o == nil || o.PartitionsCount == nil {
		return nil, false
	}
	return o.PartitionsCount, true
}

// HasPartitionsCount returns a boolean if a field has been set.
func (o *Topic) HasPartitionsCount() bool {
	return o != nil && o.PartitionsCount != nil
}

// SetPartitionsCount gets a reference to the given int32 and assigns it to the PartitionsCount field.
func (o *Topic) SetPartitionsCount(v int32) {
	o.PartitionsCount = &v
}

// GetConsumerGroupsCount returns the ConsumerGroupsCount field value if set, zero value otherwise.
func (o *Topic) GetConsumerGroupsCount() int32 {
	if o == nil || o.ConsumerGroupsCount == nil {
		var ret int32
		return ret
	}
	return *o.ConsumerGroupsCount
}

// GetConsumerGroupsCountOk returns a tuple with the ConsumerGroupsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetConsumerGroupsCountOk() (*int32, bool) {
	if o == nil || o.ConsumerGroupsCount == nil {
		return nil, false
	}
	return o.ConsumerGroupsCount, true
}

// HasConsumerGroupsCount returns a boolean if a field has been set.
func (o *Topic) HasConsumerGroupsCount() bool {
	return o != nil && o.ConsumerGroupsCount != nil
}

// SetConsumerGroupsCount gets a reference to the given int32 and assigns it to the ConsumerGroupsCount field.
func (o *Topic) SetConsumerGroupsCount(v int32) {
	o.ConsumerGroupsCount = &v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *Topic) GetReplicaCount() int32 {
	if o == nil || o.ReplicaCount == nil {
		var ret int32
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetReplicaCountOk() (*int32, bool) {
	if o == nil || o.ReplicaCount == nil {
		return nil, false
	}
	return o.ReplicaCount, true
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *Topic) HasReplicaCount() bool {
	return o != nil && o.ReplicaCount != nil
}

// SetReplicaCount gets a reference to the given int32 and assigns it to the ReplicaCount field.
func (o *Topic) SetReplicaCount(v int32) {
	o.ReplicaCount = &v
}

// GetTotalLogSize returns the TotalLogSize field value if set, zero value otherwise.
func (o *Topic) GetTotalLogSize() int64 {
	if o == nil || o.TotalLogSize == nil {
		var ret int64
		return ret
	}
	return *o.TotalLogSize
}

// GetTotalLogSizeOk returns a tuple with the TotalLogSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetTotalLogSizeOk() (*int64, bool) {
	if o == nil || o.TotalLogSize == nil {
		return nil, false
	}
	return o.TotalLogSize, true
}

// HasTotalLogSize returns a boolean if a field has been set.
func (o *Topic) HasTotalLogSize() bool {
	return o != nil && o.TotalLogSize != nil
}

// SetTotalLogSize gets a reference to the given int64 and assigns it to the TotalLogSize field.
func (o *Topic) SetTotalLogSize(v int64) {
	o.TotalLogSize = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *Topic) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Topic) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *Topic) HasInternal() bool {
	return o != nil && o.Internal != nil
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *Topic) SetInternal(v bool) {
	o.Internal = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Topic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.PartitionsCount != nil {
		toSerialize["partitionsCount"] = o.PartitionsCount
	}
	if o.ConsumerGroupsCount != nil {
		toSerialize["consumerGroupsCount"] = o.ConsumerGroupsCount
	}
	if o.ReplicaCount != nil {
		toSerialize["replicaCount"] = o.ReplicaCount
	}
	if o.TotalLogSize != nil {
		toSerialize["TotalLogSize"] = o.TotalLogSize
	}
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Topic) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                *string `json:"name"`
		PartitionsCount     *int32  `json:"partitionsCount,omitempty"`
		ConsumerGroupsCount *int32  `json:"consumerGroupsCount,omitempty"`
		ReplicaCount        *int32  `json:"replicaCount,omitempty"`
		TotalLogSize        *int64  `json:"TotalLogSize,omitempty"`
		Internal            *bool   `json:"internal,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "partitionsCount", "consumerGroupsCount", "replicaCount", "TotalLogSize", "internal"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.PartitionsCount = all.PartitionsCount
	o.ConsumerGroupsCount = all.ConsumerGroupsCount
	o.ReplicaCount = all.ReplicaCount
	o.TotalLogSize = all.TotalLogSize
	o.Internal = all.Internal

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
