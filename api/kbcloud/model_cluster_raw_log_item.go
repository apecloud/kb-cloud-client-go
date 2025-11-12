// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterRawLogItem Cluster raw log item represents a single log entry
type ClusterRawLogItem struct {
	Message   *string `json:"message,omitempty"`
	Timestamp *int32  `json:"timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterRawLogItem instantiates a new ClusterRawLogItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterRawLogItem() *ClusterRawLogItem {
	this := ClusterRawLogItem{}
	return &this
}

// NewClusterRawLogItemWithDefaults instantiates a new ClusterRawLogItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterRawLogItemWithDefaults() *ClusterRawLogItem {
	this := ClusterRawLogItem{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterRawLogItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRawLogItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterRawLogItem) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClusterRawLogItem) SetMessage(v string) {
	o.Message = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ClusterRawLogItem) GetTimestamp() int32 {
	if o == nil || o.Timestamp == nil {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRawLogItem) GetTimestampOk() (*int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ClusterRawLogItem) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *ClusterRawLogItem) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterRawLogItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterRawLogItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message   *string `json:"message,omitempty"`
		Timestamp *int32  `json:"timestamp,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "timestamp"})
	} else {
		return err
	}
	o.Message = all.Message
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
