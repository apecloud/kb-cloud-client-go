// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RbmqVHost struct {
	// Virtual host name
	Name *string `json:"name,omitempty"`
	// Virtual host description
	Description *string  `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	// Type of queue to create in virtual host when unspecified
	DefaultQueueType *string `json:"default_queue_type,omitempty"`
	// True if tracing is enabled for this virtual host
	Tracing *bool `json:"tracing,omitempty"`
	// Total number of messages in queues of this virtual host
	Messages *int32 `json:"messages,omitempty"`
	// Rate of messages in queues of this virtual host per second
	MessagesRate *float64 `json:"messages_rate,omitempty"`
	// Total number of messages ready to be delivered in queues of this virtual host
	MessagesReady *int32 `json:"messages_ready,omitempty"`
	// Rate of messages ready to be delivered in queues of this virtual host per second
	MessagesReadyRate *float64 `json:"messages_ready_rate,omitempty"`
	// Total number of messages pending acknowledgement from consumers in this virtual host
	MessagesUnacknowledged *int32 `json:"messages_unacknowledged,omitempty"`
	// Rate of messages pending acknowledgement from consumers in this virtual host per second
	MessagesUnacknowledgedRate *float64 `json:"messages_unacknowledged_rate,omitempty"`
	// Octets received
	RecvOct *int64 `json:"recv_oct,omitempty"`
	// Octets sent
	SendOct *int64 `json:"send_oct,omitempty"`
	// Count of received messages
	RecvCnt *int64 `json:"recv_cnt,omitempty"`
	// Count of sent messages
	SendCnt *int64 `json:"send_cnt,omitempty"`
	// Count of pending messages to be sent
	SendPend *int64 `json:"send_pend,omitempty"`
	// Rate of octets received per second
	RecvOctRate *float64 `json:"recv_oct_rate,omitempty"`
	// Rate of octets sent per second
	SendOctRate *float64 `json:"send_oct_rate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRbmqVHost instantiates a new RbmqVHost object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRbmqVHost() *RbmqVHost {
	this := RbmqVHost{}
	return &this
}

// NewRbmqVHostWithDefaults instantiates a new RbmqVHost object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRbmqVHostWithDefaults() *RbmqVHost {
	this := RbmqVHost{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RbmqVHost) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RbmqVHost) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RbmqVHost) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RbmqVHost) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RbmqVHost) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RbmqVHost) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RbmqVHost) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RbmqVHost) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RbmqVHost) SetTags(v []string) {
	o.Tags = v
}

// GetDefaultQueueType returns the DefaultQueueType field value if set, zero value otherwise.
func (o *RbmqVHost) GetDefaultQueueType() string {
	if o == nil || o.DefaultQueueType == nil {
		var ret string
		return ret
	}
	return *o.DefaultQueueType
}

// GetDefaultQueueTypeOk returns a tuple with the DefaultQueueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetDefaultQueueTypeOk() (*string, bool) {
	if o == nil || o.DefaultQueueType == nil {
		return nil, false
	}
	return o.DefaultQueueType, true
}

// HasDefaultQueueType returns a boolean if a field has been set.
func (o *RbmqVHost) HasDefaultQueueType() bool {
	return o != nil && o.DefaultQueueType != nil
}

// SetDefaultQueueType gets a reference to the given string and assigns it to the DefaultQueueType field.
func (o *RbmqVHost) SetDefaultQueueType(v string) {
	o.DefaultQueueType = &v
}

// GetTracing returns the Tracing field value if set, zero value otherwise.
func (o *RbmqVHost) GetTracing() bool {
	if o == nil || o.Tracing == nil {
		var ret bool
		return ret
	}
	return *o.Tracing
}

// GetTracingOk returns a tuple with the Tracing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetTracingOk() (*bool, bool) {
	if o == nil || o.Tracing == nil {
		return nil, false
	}
	return o.Tracing, true
}

// HasTracing returns a boolean if a field has been set.
func (o *RbmqVHost) HasTracing() bool {
	return o != nil && o.Tracing != nil
}

// SetTracing gets a reference to the given bool and assigns it to the Tracing field.
func (o *RbmqVHost) SetTracing(v bool) {
	o.Tracing = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *RbmqVHost) GetMessages() int32 {
	if o == nil || o.Messages == nil {
		var ret int32
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetMessagesOk() (*int32, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *RbmqVHost) HasMessages() bool {
	return o != nil && o.Messages != nil
}

// SetMessages gets a reference to the given int32 and assigns it to the Messages field.
func (o *RbmqVHost) SetMessages(v int32) {
	o.Messages = &v
}

// GetMessagesRate returns the MessagesRate field value if set, zero value otherwise.
func (o *RbmqVHost) GetMessagesRate() float64 {
	if o == nil || o.MessagesRate == nil {
		var ret float64
		return ret
	}
	return *o.MessagesRate
}

// GetMessagesRateOk returns a tuple with the MessagesRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetMessagesRateOk() (*float64, bool) {
	if o == nil || o.MessagesRate == nil {
		return nil, false
	}
	return o.MessagesRate, true
}

// HasMessagesRate returns a boolean if a field has been set.
func (o *RbmqVHost) HasMessagesRate() bool {
	return o != nil && o.MessagesRate != nil
}

// SetMessagesRate gets a reference to the given float64 and assigns it to the MessagesRate field.
func (o *RbmqVHost) SetMessagesRate(v float64) {
	o.MessagesRate = &v
}

// GetMessagesReady returns the MessagesReady field value if set, zero value otherwise.
func (o *RbmqVHost) GetMessagesReady() int32 {
	if o == nil || o.MessagesReady == nil {
		var ret int32
		return ret
	}
	return *o.MessagesReady
}

// GetMessagesReadyOk returns a tuple with the MessagesReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetMessagesReadyOk() (*int32, bool) {
	if o == nil || o.MessagesReady == nil {
		return nil, false
	}
	return o.MessagesReady, true
}

// HasMessagesReady returns a boolean if a field has been set.
func (o *RbmqVHost) HasMessagesReady() bool {
	return o != nil && o.MessagesReady != nil
}

// SetMessagesReady gets a reference to the given int32 and assigns it to the MessagesReady field.
func (o *RbmqVHost) SetMessagesReady(v int32) {
	o.MessagesReady = &v
}

// GetMessagesReadyRate returns the MessagesReadyRate field value if set, zero value otherwise.
func (o *RbmqVHost) GetMessagesReadyRate() float64 {
	if o == nil || o.MessagesReadyRate == nil {
		var ret float64
		return ret
	}
	return *o.MessagesReadyRate
}

// GetMessagesReadyRateOk returns a tuple with the MessagesReadyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetMessagesReadyRateOk() (*float64, bool) {
	if o == nil || o.MessagesReadyRate == nil {
		return nil, false
	}
	return o.MessagesReadyRate, true
}

// HasMessagesReadyRate returns a boolean if a field has been set.
func (o *RbmqVHost) HasMessagesReadyRate() bool {
	return o != nil && o.MessagesReadyRate != nil
}

// SetMessagesReadyRate gets a reference to the given float64 and assigns it to the MessagesReadyRate field.
func (o *RbmqVHost) SetMessagesReadyRate(v float64) {
	o.MessagesReadyRate = &v
}

// GetMessagesUnacknowledged returns the MessagesUnacknowledged field value if set, zero value otherwise.
func (o *RbmqVHost) GetMessagesUnacknowledged() int32 {
	if o == nil || o.MessagesUnacknowledged == nil {
		var ret int32
		return ret
	}
	return *o.MessagesUnacknowledged
}

// GetMessagesUnacknowledgedOk returns a tuple with the MessagesUnacknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetMessagesUnacknowledgedOk() (*int32, bool) {
	if o == nil || o.MessagesUnacknowledged == nil {
		return nil, false
	}
	return o.MessagesUnacknowledged, true
}

// HasMessagesUnacknowledged returns a boolean if a field has been set.
func (o *RbmqVHost) HasMessagesUnacknowledged() bool {
	return o != nil && o.MessagesUnacknowledged != nil
}

// SetMessagesUnacknowledged gets a reference to the given int32 and assigns it to the MessagesUnacknowledged field.
func (o *RbmqVHost) SetMessagesUnacknowledged(v int32) {
	o.MessagesUnacknowledged = &v
}

// GetMessagesUnacknowledgedRate returns the MessagesUnacknowledgedRate field value if set, zero value otherwise.
func (o *RbmqVHost) GetMessagesUnacknowledgedRate() float64 {
	if o == nil || o.MessagesUnacknowledgedRate == nil {
		var ret float64
		return ret
	}
	return *o.MessagesUnacknowledgedRate
}

// GetMessagesUnacknowledgedRateOk returns a tuple with the MessagesUnacknowledgedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetMessagesUnacknowledgedRateOk() (*float64, bool) {
	if o == nil || o.MessagesUnacknowledgedRate == nil {
		return nil, false
	}
	return o.MessagesUnacknowledgedRate, true
}

// HasMessagesUnacknowledgedRate returns a boolean if a field has been set.
func (o *RbmqVHost) HasMessagesUnacknowledgedRate() bool {
	return o != nil && o.MessagesUnacknowledgedRate != nil
}

// SetMessagesUnacknowledgedRate gets a reference to the given float64 and assigns it to the MessagesUnacknowledgedRate field.
func (o *RbmqVHost) SetMessagesUnacknowledgedRate(v float64) {
	o.MessagesUnacknowledgedRate = &v
}

// GetRecvOct returns the RecvOct field value if set, zero value otherwise.
func (o *RbmqVHost) GetRecvOct() int64 {
	if o == nil || o.RecvOct == nil {
		var ret int64
		return ret
	}
	return *o.RecvOct
}

// GetRecvOctOk returns a tuple with the RecvOct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetRecvOctOk() (*int64, bool) {
	if o == nil || o.RecvOct == nil {
		return nil, false
	}
	return o.RecvOct, true
}

// HasRecvOct returns a boolean if a field has been set.
func (o *RbmqVHost) HasRecvOct() bool {
	return o != nil && o.RecvOct != nil
}

// SetRecvOct gets a reference to the given int64 and assigns it to the RecvOct field.
func (o *RbmqVHost) SetRecvOct(v int64) {
	o.RecvOct = &v
}

// GetSendOct returns the SendOct field value if set, zero value otherwise.
func (o *RbmqVHost) GetSendOct() int64 {
	if o == nil || o.SendOct == nil {
		var ret int64
		return ret
	}
	return *o.SendOct
}

// GetSendOctOk returns a tuple with the SendOct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetSendOctOk() (*int64, bool) {
	if o == nil || o.SendOct == nil {
		return nil, false
	}
	return o.SendOct, true
}

// HasSendOct returns a boolean if a field has been set.
func (o *RbmqVHost) HasSendOct() bool {
	return o != nil && o.SendOct != nil
}

// SetSendOct gets a reference to the given int64 and assigns it to the SendOct field.
func (o *RbmqVHost) SetSendOct(v int64) {
	o.SendOct = &v
}

// GetRecvCnt returns the RecvCnt field value if set, zero value otherwise.
func (o *RbmqVHost) GetRecvCnt() int64 {
	if o == nil || o.RecvCnt == nil {
		var ret int64
		return ret
	}
	return *o.RecvCnt
}

// GetRecvCntOk returns a tuple with the RecvCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetRecvCntOk() (*int64, bool) {
	if o == nil || o.RecvCnt == nil {
		return nil, false
	}
	return o.RecvCnt, true
}

// HasRecvCnt returns a boolean if a field has been set.
func (o *RbmqVHost) HasRecvCnt() bool {
	return o != nil && o.RecvCnt != nil
}

// SetRecvCnt gets a reference to the given int64 and assigns it to the RecvCnt field.
func (o *RbmqVHost) SetRecvCnt(v int64) {
	o.RecvCnt = &v
}

// GetSendCnt returns the SendCnt field value if set, zero value otherwise.
func (o *RbmqVHost) GetSendCnt() int64 {
	if o == nil || o.SendCnt == nil {
		var ret int64
		return ret
	}
	return *o.SendCnt
}

// GetSendCntOk returns a tuple with the SendCnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetSendCntOk() (*int64, bool) {
	if o == nil || o.SendCnt == nil {
		return nil, false
	}
	return o.SendCnt, true
}

// HasSendCnt returns a boolean if a field has been set.
func (o *RbmqVHost) HasSendCnt() bool {
	return o != nil && o.SendCnt != nil
}

// SetSendCnt gets a reference to the given int64 and assigns it to the SendCnt field.
func (o *RbmqVHost) SetSendCnt(v int64) {
	o.SendCnt = &v
}

// GetSendPend returns the SendPend field value if set, zero value otherwise.
func (o *RbmqVHost) GetSendPend() int64 {
	if o == nil || o.SendPend == nil {
		var ret int64
		return ret
	}
	return *o.SendPend
}

// GetSendPendOk returns a tuple with the SendPend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetSendPendOk() (*int64, bool) {
	if o == nil || o.SendPend == nil {
		return nil, false
	}
	return o.SendPend, true
}

// HasSendPend returns a boolean if a field has been set.
func (o *RbmqVHost) HasSendPend() bool {
	return o != nil && o.SendPend != nil
}

// SetSendPend gets a reference to the given int64 and assigns it to the SendPend field.
func (o *RbmqVHost) SetSendPend(v int64) {
	o.SendPend = &v
}

// GetRecvOctRate returns the RecvOctRate field value if set, zero value otherwise.
func (o *RbmqVHost) GetRecvOctRate() float64 {
	if o == nil || o.RecvOctRate == nil {
		var ret float64
		return ret
	}
	return *o.RecvOctRate
}

// GetRecvOctRateOk returns a tuple with the RecvOctRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetRecvOctRateOk() (*float64, bool) {
	if o == nil || o.RecvOctRate == nil {
		return nil, false
	}
	return o.RecvOctRate, true
}

// HasRecvOctRate returns a boolean if a field has been set.
func (o *RbmqVHost) HasRecvOctRate() bool {
	return o != nil && o.RecvOctRate != nil
}

// SetRecvOctRate gets a reference to the given float64 and assigns it to the RecvOctRate field.
func (o *RbmqVHost) SetRecvOctRate(v float64) {
	o.RecvOctRate = &v
}

// GetSendOctRate returns the SendOctRate field value if set, zero value otherwise.
func (o *RbmqVHost) GetSendOctRate() float64 {
	if o == nil || o.SendOctRate == nil {
		var ret float64
		return ret
	}
	return *o.SendOctRate
}

// GetSendOctRateOk returns a tuple with the SendOctRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVHost) GetSendOctRateOk() (*float64, bool) {
	if o == nil || o.SendOctRate == nil {
		return nil, false
	}
	return o.SendOctRate, true
}

// HasSendOctRate returns a boolean if a field has been set.
func (o *RbmqVHost) HasSendOctRate() bool {
	return o != nil && o.SendOctRate != nil
}

// SetSendOctRate gets a reference to the given float64 and assigns it to the SendOctRate field.
func (o *RbmqVHost) SetSendOctRate(v float64) {
	o.SendOctRate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RbmqVHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.DefaultQueueType != nil {
		toSerialize["default_queue_type"] = o.DefaultQueueType
	}
	if o.Tracing != nil {
		toSerialize["tracing"] = o.Tracing
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.MessagesRate != nil {
		toSerialize["messages_rate"] = o.MessagesRate
	}
	if o.MessagesReady != nil {
		toSerialize["messages_ready"] = o.MessagesReady
	}
	if o.MessagesReadyRate != nil {
		toSerialize["messages_ready_rate"] = o.MessagesReadyRate
	}
	if o.MessagesUnacknowledged != nil {
		toSerialize["messages_unacknowledged"] = o.MessagesUnacknowledged
	}
	if o.MessagesUnacknowledgedRate != nil {
		toSerialize["messages_unacknowledged_rate"] = o.MessagesUnacknowledgedRate
	}
	if o.RecvOct != nil {
		toSerialize["recv_oct"] = o.RecvOct
	}
	if o.SendOct != nil {
		toSerialize["send_oct"] = o.SendOct
	}
	if o.RecvCnt != nil {
		toSerialize["recv_cnt"] = o.RecvCnt
	}
	if o.SendCnt != nil {
		toSerialize["send_cnt"] = o.SendCnt
	}
	if o.SendPend != nil {
		toSerialize["send_pend"] = o.SendPend
	}
	if o.RecvOctRate != nil {
		toSerialize["recv_oct_rate"] = o.RecvOctRate
	}
	if o.SendOctRate != nil {
		toSerialize["send_oct_rate"] = o.SendOctRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RbmqVHost) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                       *string  `json:"name,omitempty"`
		Description                *string  `json:"description,omitempty"`
		Tags                       []string `json:"tags,omitempty"`
		DefaultQueueType           *string  `json:"default_queue_type,omitempty"`
		Tracing                    *bool    `json:"tracing,omitempty"`
		Messages                   *int32   `json:"messages,omitempty"`
		MessagesRate               *float64 `json:"messages_rate,omitempty"`
		MessagesReady              *int32   `json:"messages_ready,omitempty"`
		MessagesReadyRate          *float64 `json:"messages_ready_rate,omitempty"`
		MessagesUnacknowledged     *int32   `json:"messages_unacknowledged,omitempty"`
		MessagesUnacknowledgedRate *float64 `json:"messages_unacknowledged_rate,omitempty"`
		RecvOct                    *int64   `json:"recv_oct,omitempty"`
		SendOct                    *int64   `json:"send_oct,omitempty"`
		RecvCnt                    *int64   `json:"recv_cnt,omitempty"`
		SendCnt                    *int64   `json:"send_cnt,omitempty"`
		SendPend                   *int64   `json:"send_pend,omitempty"`
		RecvOctRate                *float64 `json:"recv_oct_rate,omitempty"`
		SendOctRate                *float64 `json:"send_oct_rate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "tags", "default_queue_type", "tracing", "messages", "messages_rate", "messages_ready", "messages_ready_rate", "messages_unacknowledged", "messages_unacknowledged_rate", "recv_oct", "send_oct", "recv_cnt", "send_cnt", "send_pend", "recv_oct_rate", "send_oct_rate"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Description = all.Description
	o.Tags = all.Tags
	o.DefaultQueueType = all.DefaultQueueType
	o.Tracing = all.Tracing
	o.Messages = all.Messages
	o.MessagesRate = all.MessagesRate
	o.MessagesReady = all.MessagesReady
	o.MessagesReadyRate = all.MessagesReadyRate
	o.MessagesUnacknowledged = all.MessagesUnacknowledged
	o.MessagesUnacknowledgedRate = all.MessagesUnacknowledgedRate
	o.RecvOct = all.RecvOct
	o.SendOct = all.SendOct
	o.RecvCnt = all.RecvCnt
	o.SendCnt = all.SendCnt
	o.SendPend = all.SendPend
	o.RecvOctRate = all.RecvOctRate
	o.SendOctRate = all.SendOctRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
