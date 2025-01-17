// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TopicDetails struct {
	Name     string `json:"name"`
	Internal *bool  `json:"internal,omitempty"`
	// the replica count of the topic
	ReplicaCount *int32          `json:"replicaCount,omitempty"`
	Partitions   []PartitionInfo `json:"partitions,omitempty"`
	TotalLogSize *int64          `json:"totalLogSize,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTopicDetails instantiates a new TopicDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopicDetails(name string) *TopicDetails {
	this := TopicDetails{}
	this.Name = name
	return &this
}

// NewTopicDetailsWithDefaults instantiates a new TopicDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopicDetailsWithDefaults() *TopicDetails {
	this := TopicDetails{}
	return &this
}

// GetName returns the Name field value.
func (o *TopicDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TopicDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TopicDetails) SetName(v string) {
	o.Name = v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *TopicDetails) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicDetails) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *TopicDetails) HasInternal() bool {
	return o != nil && o.Internal != nil
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *TopicDetails) SetInternal(v bool) {
	o.Internal = &v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *TopicDetails) GetReplicaCount() int32 {
	if o == nil || o.ReplicaCount == nil {
		var ret int32
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicDetails) GetReplicaCountOk() (*int32, bool) {
	if o == nil || o.ReplicaCount == nil {
		return nil, false
	}
	return o.ReplicaCount, true
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *TopicDetails) HasReplicaCount() bool {
	return o != nil && o.ReplicaCount != nil
}

// SetReplicaCount gets a reference to the given int32 and assigns it to the ReplicaCount field.
func (o *TopicDetails) SetReplicaCount(v int32) {
	o.ReplicaCount = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *TopicDetails) GetPartitions() []PartitionInfo {
	if o == nil || o.Partitions == nil {
		var ret []PartitionInfo
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicDetails) GetPartitionsOk() (*[]PartitionInfo, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return &o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *TopicDetails) HasPartitions() bool {
	return o != nil && o.Partitions != nil
}

// SetPartitions gets a reference to the given []PartitionInfo and assigns it to the Partitions field.
func (o *TopicDetails) SetPartitions(v []PartitionInfo) {
	o.Partitions = v
}

// GetTotalLogSize returns the TotalLogSize field value if set, zero value otherwise.
func (o *TopicDetails) GetTotalLogSize() int64 {
	if o == nil || o.TotalLogSize == nil {
		var ret int64
		return ret
	}
	return *o.TotalLogSize
}

// GetTotalLogSizeOk returns a tuple with the TotalLogSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicDetails) GetTotalLogSizeOk() (*int64, bool) {
	if o == nil || o.TotalLogSize == nil {
		return nil, false
	}
	return o.TotalLogSize, true
}

// HasTotalLogSize returns a boolean if a field has been set.
func (o *TopicDetails) HasTotalLogSize() bool {
	return o != nil && o.TotalLogSize != nil
}

// SetTotalLogSize gets a reference to the given int64 and assigns it to the TotalLogSize field.
func (o *TopicDetails) SetTotalLogSize(v int64) {
	o.TotalLogSize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopicDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}
	if o.ReplicaCount != nil {
		toSerialize["replicaCount"] = o.ReplicaCount
	}
	if o.Partitions != nil {
		toSerialize["partitions"] = o.Partitions
	}
	if o.TotalLogSize != nil {
		toSerialize["totalLogSize"] = o.TotalLogSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopicDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string         `json:"name"`
		Internal     *bool           `json:"internal,omitempty"`
		ReplicaCount *int32          `json:"replicaCount,omitempty"`
		Partitions   []PartitionInfo `json:"partitions,omitempty"`
		TotalLogSize *int64          `json:"totalLogSize,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "internal", "replicaCount", "partitions", "totalLogSize"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Internal = all.Internal
	o.ReplicaCount = all.ReplicaCount
	o.Partitions = all.Partitions
	o.TotalLogSize = all.TotalLogSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
