// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ACLUserResponse struct {
	Mode     *string   `json:"mode,omitempty"`
	Master   []ACLUser `json:"master,omitempty"`
	Sentinel []ACLUser `json:"sentinel,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewACLUserResponse instantiates a new ACLUserResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewACLUserResponse() *ACLUserResponse {
	this := ACLUserResponse{}
	return &this
}

// NewACLUserResponseWithDefaults instantiates a new ACLUserResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewACLUserResponseWithDefaults() *ACLUserResponse {
	this := ACLUserResponse{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ACLUserResponse) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUserResponse) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ACLUserResponse) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ACLUserResponse) SetMode(v string) {
	o.Mode = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *ACLUserResponse) GetMaster() []ACLUser {
	if o == nil || o.Master == nil {
		var ret []ACLUser
		return ret
	}
	return o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUserResponse) GetMasterOk() (*[]ACLUser, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return &o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *ACLUserResponse) HasMaster() bool {
	return o != nil && o.Master != nil
}

// SetMaster gets a reference to the given []ACLUser and assigns it to the Master field.
func (o *ACLUserResponse) SetMaster(v []ACLUser) {
	o.Master = v
}

// GetSentinel returns the Sentinel field value if set, zero value otherwise.
func (o *ACLUserResponse) GetSentinel() []ACLUser {
	if o == nil || o.Sentinel == nil {
		var ret []ACLUser
		return ret
	}
	return o.Sentinel
}

// GetSentinelOk returns a tuple with the Sentinel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUserResponse) GetSentinelOk() (*[]ACLUser, bool) {
	if o == nil || o.Sentinel == nil {
		return nil, false
	}
	return &o.Sentinel, true
}

// HasSentinel returns a boolean if a field has been set.
func (o *ACLUserResponse) HasSentinel() bool {
	return o != nil && o.Sentinel != nil
}

// SetSentinel gets a reference to the given []ACLUser and assigns it to the Sentinel field.
func (o *ACLUserResponse) SetSentinel(v []ACLUser) {
	o.Sentinel = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ACLUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
	if o.Sentinel != nil {
		toSerialize["sentinel"] = o.Sentinel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ACLUserResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode     *string   `json:"mode,omitempty"`
		Master   []ACLUser `json:"master,omitempty"`
		Sentinel []ACLUser `json:"sentinel,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"mode", "master", "sentinel"})
	} else {
		return err
	}
	o.Mode = all.Mode
	o.Master = all.Master
	o.Sentinel = all.Sentinel

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
