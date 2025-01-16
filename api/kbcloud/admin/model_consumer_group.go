// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ConsumerGroup struct {
	GroupId *string  `json:"groupId,omitempty"`
	State   *string  `json:"state,omitempty"`
	Lag     *int64   `json:"lag,omitempty"`
	Topics  []string `json:"topics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConsumerGroup instantiates a new ConsumerGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConsumerGroup() *ConsumerGroup {
	this := ConsumerGroup{}
	return &this
}

// NewConsumerGroupWithDefaults instantiates a new ConsumerGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConsumerGroupWithDefaults() *ConsumerGroup {
	this := ConsumerGroup{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ConsumerGroup) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ConsumerGroup) HasGroupId() bool {
	return o != nil && o.GroupId != nil
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ConsumerGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ConsumerGroup) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroup) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ConsumerGroup) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ConsumerGroup) SetState(v string) {
	o.State = &v
}

// GetLag returns the Lag field value if set, zero value otherwise.
func (o *ConsumerGroup) GetLag() int64 {
	if o == nil || o.Lag == nil {
		var ret int64
		return ret
	}
	return *o.Lag
}

// GetLagOk returns a tuple with the Lag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroup) GetLagOk() (*int64, bool) {
	if o == nil || o.Lag == nil {
		return nil, false
	}
	return o.Lag, true
}

// HasLag returns a boolean if a field has been set.
func (o *ConsumerGroup) HasLag() bool {
	return o != nil && o.Lag != nil
}

// SetLag gets a reference to the given int64 and assigns it to the Lag field.
func (o *ConsumerGroup) SetLag(v int64) {
	o.Lag = &v
}

// GetTopics returns the Topics field value if set, zero value otherwise.
func (o *ConsumerGroup) GetTopics() []string {
	if o == nil || o.Topics == nil {
		var ret []string
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroup) GetTopicsOk() (*[]string, bool) {
	if o == nil || o.Topics == nil {
		return nil, false
	}
	return &o.Topics, true
}

// HasTopics returns a boolean if a field has been set.
func (o *ConsumerGroup) HasTopics() bool {
	return o != nil && o.Topics != nil
}

// SetTopics gets a reference to the given []string and assigns it to the Topics field.
func (o *ConsumerGroup) SetTopics(v []string) {
	o.Topics = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConsumerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Lag != nil {
		toSerialize["lag"] = o.Lag
	}
	if o.Topics != nil {
		toSerialize["topics"] = o.Topics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConsumerGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupId *string  `json:"groupId,omitempty"`
		State   *string  `json:"state,omitempty"`
		Lag     *int64   `json:"lag,omitempty"`
		Topics  []string `json:"topics,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"groupId", "state", "lag", "topics"})
	} else {
		return err
	}
	o.GroupId = all.GroupId
	o.State = all.State
	o.Lag = all.Lag
	o.Topics = all.Topics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
