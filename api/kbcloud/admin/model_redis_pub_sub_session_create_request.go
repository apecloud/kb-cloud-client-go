// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisPubSubSessionCreateRequest struct {
	Channels []string `json:"channels,omitempty"`
	Patterns []string `json:"patterns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisPubSubSessionCreateRequest instantiates a new RedisPubSubSessionCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisPubSubSessionCreateRequest() *RedisPubSubSessionCreateRequest {
	this := RedisPubSubSessionCreateRequest{}
	return &this
}

// NewRedisPubSubSessionCreateRequestWithDefaults instantiates a new RedisPubSubSessionCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisPubSubSessionCreateRequestWithDefaults() *RedisPubSubSessionCreateRequest {
	this := RedisPubSubSessionCreateRequest{}
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *RedisPubSubSessionCreateRequest) GetChannels() []string {
	if o == nil || o.Channels == nil {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubSessionCreateRequest) GetChannelsOk() (*[]string, bool) {
	if o == nil || o.Channels == nil {
		return nil, false
	}
	return &o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *RedisPubSubSessionCreateRequest) HasChannels() bool {
	return o != nil && o.Channels != nil
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *RedisPubSubSessionCreateRequest) SetChannels(v []string) {
	o.Channels = v
}

// GetPatterns returns the Patterns field value if set, zero value otherwise.
func (o *RedisPubSubSessionCreateRequest) GetPatterns() []string {
	if o == nil || o.Patterns == nil {
		var ret []string
		return ret
	}
	return o.Patterns
}

// GetPatternsOk returns a tuple with the Patterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubSessionCreateRequest) GetPatternsOk() (*[]string, bool) {
	if o == nil || o.Patterns == nil {
		return nil, false
	}
	return &o.Patterns, true
}

// HasPatterns returns a boolean if a field has been set.
func (o *RedisPubSubSessionCreateRequest) HasPatterns() bool {
	return o != nil && o.Patterns != nil
}

// SetPatterns gets a reference to the given []string and assigns it to the Patterns field.
func (o *RedisPubSubSessionCreateRequest) SetPatterns(v []string) {
	o.Patterns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisPubSubSessionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if o.Patterns != nil {
		toSerialize["patterns"] = o.Patterns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisPubSubSessionCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channels []string `json:"channels,omitempty"`
		Patterns []string `json:"patterns,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"channels", "patterns"})
	} else {
		return err
	}
	o.Channels = all.Channels
	o.Patterns = all.Patterns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
