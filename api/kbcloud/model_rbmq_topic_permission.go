// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RbmqTopicPermission struct {
	// VHost name
	Vhost *string `json:"vhost,omitempty"`
	// Exchange name
	Exchange *string `json:"exchange,omitempty"`
	// Read permission regexp
	Read *string `json:"read,omitempty"`
	// Write permission regexp
	Write *string `json:"write,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRbmqTopicPermission instantiates a new RbmqTopicPermission object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRbmqTopicPermission() *RbmqTopicPermission {
	this := RbmqTopicPermission{}
	return &this
}

// NewRbmqTopicPermissionWithDefaults instantiates a new RbmqTopicPermission object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRbmqTopicPermissionWithDefaults() *RbmqTopicPermission {
	this := RbmqTopicPermission{}
	return &this
}

// GetVhost returns the Vhost field value if set, zero value otherwise.
func (o *RbmqTopicPermission) GetVhost() string {
	if o == nil || o.Vhost == nil {
		var ret string
		return ret
	}
	return *o.Vhost
}

// GetVhostOk returns a tuple with the Vhost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqTopicPermission) GetVhostOk() (*string, bool) {
	if o == nil || o.Vhost == nil {
		return nil, false
	}
	return o.Vhost, true
}

// HasVhost returns a boolean if a field has been set.
func (o *RbmqTopicPermission) HasVhost() bool {
	return o != nil && o.Vhost != nil
}

// SetVhost gets a reference to the given string and assigns it to the Vhost field.
func (o *RbmqTopicPermission) SetVhost(v string) {
	o.Vhost = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *RbmqTopicPermission) GetExchange() string {
	if o == nil || o.Exchange == nil {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqTopicPermission) GetExchangeOk() (*string, bool) {
	if o == nil || o.Exchange == nil {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *RbmqTopicPermission) HasExchange() bool {
	return o != nil && o.Exchange != nil
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *RbmqTopicPermission) SetExchange(v string) {
	o.Exchange = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *RbmqTopicPermission) GetRead() string {
	if o == nil || o.Read == nil {
		var ret string
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqTopicPermission) GetReadOk() (*string, bool) {
	if o == nil || o.Read == nil {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *RbmqTopicPermission) HasRead() bool {
	return o != nil && o.Read != nil
}

// SetRead gets a reference to the given string and assigns it to the Read field.
func (o *RbmqTopicPermission) SetRead(v string) {
	o.Read = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *RbmqTopicPermission) GetWrite() string {
	if o == nil || o.Write == nil {
		var ret string
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqTopicPermission) GetWriteOk() (*string, bool) {
	if o == nil || o.Write == nil {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *RbmqTopicPermission) HasWrite() bool {
	return o != nil && o.Write != nil
}

// SetWrite gets a reference to the given string and assigns it to the Write field.
func (o *RbmqTopicPermission) SetWrite(v string) {
	o.Write = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RbmqTopicPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Vhost != nil {
		toSerialize["vhost"] = o.Vhost
	}
	if o.Exchange != nil {
		toSerialize["exchange"] = o.Exchange
	}
	if o.Read != nil {
		toSerialize["read"] = o.Read
	}
	if o.Write != nil {
		toSerialize["write"] = o.Write
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RbmqTopicPermission) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Vhost    *string `json:"vhost,omitempty"`
		Exchange *string `json:"exchange,omitempty"`
		Read     *string `json:"read,omitempty"`
		Write    *string `json:"write,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"vhost", "exchange", "read", "write"})
	} else {
		return err
	}
	o.Vhost = all.Vhost
	o.Exchange = all.Exchange
	o.Read = all.Read
	o.Write = all.Write

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
