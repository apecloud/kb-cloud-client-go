// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisPubSubSession struct {
	SessionId *string    `json:"sessionId,omitempty"`
	Channels  []string   `json:"channels,omitempty"`
	Patterns  []string   `json:"patterns,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisPubSubSession instantiates a new RedisPubSubSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisPubSubSession() *RedisPubSubSession {
	this := RedisPubSubSession{}
	return &this
}

// NewRedisPubSubSessionWithDefaults instantiates a new RedisPubSubSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisPubSubSessionWithDefaults() *RedisPubSubSession {
	this := RedisPubSubSession{}
	return &this
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *RedisPubSubSession) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubSession) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *RedisPubSubSession) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *RedisPubSubSession) SetSessionId(v string) {
	o.SessionId = &v
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *RedisPubSubSession) GetChannels() []string {
	if o == nil || o.Channels == nil {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubSession) GetChannelsOk() (*[]string, bool) {
	if o == nil || o.Channels == nil {
		return nil, false
	}
	return &o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *RedisPubSubSession) HasChannels() bool {
	return o != nil && o.Channels != nil
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *RedisPubSubSession) SetChannels(v []string) {
	o.Channels = v
}

// GetPatterns returns the Patterns field value if set, zero value otherwise.
func (o *RedisPubSubSession) GetPatterns() []string {
	if o == nil || o.Patterns == nil {
		var ret []string
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubSession) GetPatternsOk() (*[]string, bool) {
	if o == nil || o.Patterns == nil {
		return nil, false
	}
	return &o.Patterns, true
}

// HasPatterns returns a boolean if a field has been set.
func (o *RedisPubSubSession) HasPatterns() bool {
	return o != nil && o.Patterns != nil
}

// SetPatterns gets a reference to the given []string and assigns it to the Patterns field.
func (o *RedisPubSubSession) SetPatterns(v []string) {
	o.Patterns = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RedisPubSubSession) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubSession) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RedisPubSubSession) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RedisPubSubSession) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisPubSubSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if o.Patterns != nil {
		toSerialize["patterns"] = o.Patterns
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisPubSubSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionId *string    `json:"sessionId,omitempty"`
		Channels  []string   `json:"channels,omitempty"`
		Patterns  []string   `json:"patterns,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sessionId", "channels", "patterns", "createdAt"})
	} else {
		return err
	}
	o.SessionId = all.SessionId
	o.Channels = all.Channels
	o.Patterns = all.Patterns
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
