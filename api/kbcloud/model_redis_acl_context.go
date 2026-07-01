// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisACLContext struct {
	// Redis ACL user bound to the datasource.
	User *string `json:"user,omitempty"`
	// Key patterns the datasource ACL user is allowed to access.
	KeyPatterns []string `json:"keyPatterns,omitempty"`
	// Pub/Sub channel patterns the datasource ACL user is allowed to access.
	ChannelPatterns []string `json:"channelPatterns,omitempty"`
	// Command allow/deny rules for the datasource ACL user.
	CommandRules []string `json:"commandRules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisACLContext instantiates a new RedisACLContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisACLContext() *RedisACLContext {
	this := RedisACLContext{}
	return &this
}

// NewRedisACLContextWithDefaults instantiates a new RedisACLContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisACLContextWithDefaults() *RedisACLContext {
	this := RedisACLContext{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RedisACLContext) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisACLContext) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RedisACLContext) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *RedisACLContext) SetUser(v string) {
	o.User = &v
}

// GetKeyPatterns returns the KeyPatterns field value if set, zero value otherwise.
func (o *RedisACLContext) GetKeyPatterns() []string {
	if o == nil || o.KeyPatterns == nil {
		var ret []string
		return ret
	}
	return o.KeyPatterns
}

// GetKeyPatternsOk returns a tuple with the KeyPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisACLContext) GetKeyPatternsOk() (*[]string, bool) {
	if o == nil || o.KeyPatterns == nil {
		return nil, false
	}
	return &o.KeyPatterns, true
}

// HasKeyPatterns returns a boolean if a field has been set.
func (o *RedisACLContext) HasKeyPatterns() bool {
	return o != nil && o.KeyPatterns != nil
}

// SetKeyPatterns gets a reference to the given []string and assigns it to the KeyPatterns field.
func (o *RedisACLContext) SetKeyPatterns(v []string) {
	o.KeyPatterns = v
}

// GetChannelPatterns returns the ChannelPatterns field value if set, zero value otherwise.
func (o *RedisACLContext) GetChannelPatterns() []string {
	if o == nil || o.ChannelPatterns == nil {
		var ret []string
		return ret
	}
	return o.ChannelPatterns
}

// GetChannelPatternsOk returns a tuple with the ChannelPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisACLContext) GetChannelPatternsOk() (*[]string, bool) {
	if o == nil || o.ChannelPatterns == nil {
		return nil, false
	}
	return &o.ChannelPatterns, true
}

// HasChannelPatterns returns a boolean if a field has been set.
func (o *RedisACLContext) HasChannelPatterns() bool {
	return o != nil && o.ChannelPatterns != nil
}

// SetChannelPatterns gets a reference to the given []string and assigns it to the ChannelPatterns field.
func (o *RedisACLContext) SetChannelPatterns(v []string) {
	o.ChannelPatterns = v
}

// GetCommandRules returns the CommandRules field value if set, zero value otherwise.
func (o *RedisACLContext) GetCommandRules() []string {
	if o == nil || o.CommandRules == nil {
		var ret []string
		return ret
	}
	return o.CommandRules
}

// GetCommandRulesOk returns a tuple with the CommandRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisACLContext) GetCommandRulesOk() (*[]string, bool) {
	if o == nil || o.CommandRules == nil {
		return nil, false
	}
	return &o.CommandRules, true
}

// HasCommandRules returns a boolean if a field has been set.
func (o *RedisACLContext) HasCommandRules() bool {
	return o != nil && o.CommandRules != nil
}

// SetCommandRules gets a reference to the given []string and assigns it to the CommandRules field.
func (o *RedisACLContext) SetCommandRules(v []string) {
	o.CommandRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisACLContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.KeyPatterns != nil {
		toSerialize["keyPatterns"] = o.KeyPatterns
	}
	if o.ChannelPatterns != nil {
		toSerialize["channelPatterns"] = o.ChannelPatterns
	}
	if o.CommandRules != nil {
		toSerialize["commandRules"] = o.CommandRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisACLContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		User            *string  `json:"user,omitempty"`
		KeyPatterns     []string `json:"keyPatterns,omitempty"`
		ChannelPatterns []string `json:"channelPatterns,omitempty"`
		CommandRules    []string `json:"commandRules,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"user", "keyPatterns", "channelPatterns", "commandRules"})
	} else {
		return err
	}
	o.User = all.User
	o.KeyPatterns = all.KeyPatterns
	o.ChannelPatterns = all.ChannelPatterns
	o.CommandRules = all.CommandRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
