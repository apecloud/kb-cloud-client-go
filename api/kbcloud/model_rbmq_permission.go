// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RbmqPermission struct {
	// List of VHost permissions
	VhostPermission []RbmqVhostPermission `json:"vhostPermission,omitempty"`
	// List of topic permissions
	TopicPermission []RbmqTopicPermission `json:"topicPermission,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRbmqPermission instantiates a new RbmqPermission object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRbmqPermission() *RbmqPermission {
	this := RbmqPermission{}
	return &this
}

// NewRbmqPermissionWithDefaults instantiates a new RbmqPermission object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRbmqPermissionWithDefaults() *RbmqPermission {
	this := RbmqPermission{}
	return &this
}

// GetVhostPermission returns the VhostPermission field value if set, zero value otherwise.
func (o *RbmqPermission) GetVhostPermission() []RbmqVhostPermission {
	if o == nil || o.VhostPermission == nil {
		var ret []RbmqVhostPermission
		return ret
	}
	return o.VhostPermission
}

// GetVhostPermissionOk returns a tuple with the VhostPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqPermission) GetVhostPermissionOk() (*[]RbmqVhostPermission, bool) {
	if o == nil || o.VhostPermission == nil {
		return nil, false
	}
	return &o.VhostPermission, true
}

// HasVhostPermission returns a boolean if a field has been set.
func (o *RbmqPermission) HasVhostPermission() bool {
	return o != nil && o.VhostPermission != nil
}

// SetVhostPermission gets a reference to the given []RbmqVhostPermission and assigns it to the VhostPermission field.
func (o *RbmqPermission) SetVhostPermission(v []RbmqVhostPermission) {
	o.VhostPermission = v
}

// GetTopicPermission returns the TopicPermission field value if set, zero value otherwise.
func (o *RbmqPermission) GetTopicPermission() []RbmqTopicPermission {
	if o == nil || o.TopicPermission == nil {
		var ret []RbmqTopicPermission
		return ret
	}
	return o.TopicPermission
}

// GetTopicPermissionOk returns a tuple with the TopicPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqPermission) GetTopicPermissionOk() (*[]RbmqTopicPermission, bool) {
	if o == nil || o.TopicPermission == nil {
		return nil, false
	}
	return &o.TopicPermission, true
}

// HasTopicPermission returns a boolean if a field has been set.
func (o *RbmqPermission) HasTopicPermission() bool {
	return o != nil && o.TopicPermission != nil
}

// SetTopicPermission gets a reference to the given []RbmqTopicPermission and assigns it to the TopicPermission field.
func (o *RbmqPermission) SetTopicPermission(v []RbmqTopicPermission) {
	o.TopicPermission = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RbmqPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.VhostPermission != nil {
		toSerialize["vhostPermission"] = o.VhostPermission
	}
	if o.TopicPermission != nil {
		toSerialize["topicPermission"] = o.TopicPermission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RbmqPermission) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		VhostPermission []RbmqVhostPermission `json:"vhostPermission,omitempty"`
		TopicPermission []RbmqTopicPermission `json:"topicPermission,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"vhostPermission", "topicPermission"})
	} else {
		return err
	}
	o.VhostPermission = all.VhostPermission
	o.TopicPermission = all.TopicPermission

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
