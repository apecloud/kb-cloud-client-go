// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RbmqVhostPermission struct {
	// VHost name
	Vhost *string `json:"vhost,omitempty"`
	// Read permission regexp
	Read *string `json:"read,omitempty"`
	// Write permission regexp
	Write *string `json:"write,omitempty"`
	// Config permission regexp
	Config *string `json:"config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRbmqVhostPermission instantiates a new RbmqVhostPermission object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRbmqVhostPermission() *RbmqVhostPermission {
	this := RbmqVhostPermission{}
	return &this
}

// NewRbmqVhostPermissionWithDefaults instantiates a new RbmqVhostPermission object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRbmqVhostPermissionWithDefaults() *RbmqVhostPermission {
	this := RbmqVhostPermission{}
	return &this
}

// GetVhost returns the Vhost field value if set, zero value otherwise.
func (o *RbmqVhostPermission) GetVhost() string {
	if o == nil || o.Vhost == nil {
		var ret string
		return ret
	}
	return *o.Vhost
}

// GetVhostOk returns a tuple with the Vhost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVhostPermission) GetVhostOk() (*string, bool) {
	if o == nil || o.Vhost == nil {
		return nil, false
	}
	return o.Vhost, true
}

// HasVhost returns a boolean if a field has been set.
func (o *RbmqVhostPermission) HasVhost() bool {
	return o != nil && o.Vhost != nil
}

// SetVhost gets a reference to the given string and assigns it to the Vhost field.
func (o *RbmqVhostPermission) SetVhost(v string) {
	o.Vhost = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *RbmqVhostPermission) GetRead() string {
	if o == nil || o.Read == nil {
		var ret string
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVhostPermission) GetReadOk() (*string, bool) {
	if o == nil || o.Read == nil {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *RbmqVhostPermission) HasRead() bool {
	return o != nil && o.Read != nil
}

// SetRead gets a reference to the given string and assigns it to the Read field.
func (o *RbmqVhostPermission) SetRead(v string) {
	o.Read = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *RbmqVhostPermission) GetWrite() string {
	if o == nil || o.Write == nil {
		var ret string
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVhostPermission) GetWriteOk() (*string, bool) {
	if o == nil || o.Write == nil {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *RbmqVhostPermission) HasWrite() bool {
	return o != nil && o.Write != nil
}

// SetWrite gets a reference to the given string and assigns it to the Write field.
func (o *RbmqVhostPermission) SetWrite(v string) {
	o.Write = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *RbmqVhostPermission) GetConfig() string {
	if o == nil || o.Config == nil {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqVhostPermission) GetConfigOk() (*string, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *RbmqVhostPermission) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *RbmqVhostPermission) SetConfig(v string) {
	o.Config = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RbmqVhostPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Vhost != nil {
		toSerialize["vhost"] = o.Vhost
	}
	if o.Read != nil {
		toSerialize["read"] = o.Read
	}
	if o.Write != nil {
		toSerialize["write"] = o.Write
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RbmqVhostPermission) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Vhost  *string `json:"vhost,omitempty"`
		Read   *string `json:"read,omitempty"`
		Write  *string `json:"write,omitempty"`
		Config *string `json:"config,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"vhost", "read", "write", "config"})
	} else {
		return err
	}
	o.Vhost = all.Vhost
	o.Read = all.Read
	o.Write = all.Write
	o.Config = all.Config

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
