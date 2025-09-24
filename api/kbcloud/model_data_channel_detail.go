// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataChannelDetail struct {
	Channel  *DataChannelItem     `json:"channel,omitempty"`
	Progress *DataChannelProgress `json:"progress,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataChannelDetail instantiates a new DataChannelDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelDetail() *DataChannelDetail {
	this := DataChannelDetail{}
	return &this
}

// NewDataChannelDetailWithDefaults instantiates a new DataChannelDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelDetailWithDefaults() *DataChannelDetail {
	this := DataChannelDetail{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *DataChannelDetail) GetChannel() DataChannelItem {
	if o == nil || o.Channel == nil {
		var ret DataChannelItem
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelDetail) GetChannelOk() (*DataChannelItem, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *DataChannelDetail) HasChannel() bool {
	return o != nil && o.Channel != nil
}

// SetChannel gets a reference to the given DataChannelItem and assigns it to the Channel field.
func (o *DataChannelDetail) SetChannel(v DataChannelItem) {
	o.Channel = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *DataChannelDetail) GetProgress() DataChannelProgress {
	if o == nil || o.Progress == nil {
		var ret DataChannelProgress
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelDetail) GetProgressOk() (*DataChannelProgress, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *DataChannelDetail) HasProgress() bool {
	return o != nil && o.Progress != nil
}

// SetProgress gets a reference to the given DataChannelProgress and assigns it to the Progress field.
func (o *DataChannelDetail) SetProgress(v DataChannelProgress) {
	o.Progress = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataChannelDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel  *DataChannelItem     `json:"channel,omitempty"`
		Progress *DataChannelProgress `json:"progress,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"channel", "progress"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Channel != nil && all.Channel.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Channel = all.Channel
	if all.Progress != nil && all.Progress.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Progress = all.Progress

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
