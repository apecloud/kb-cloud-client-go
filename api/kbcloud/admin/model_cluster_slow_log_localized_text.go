// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogLocalizedText Localized slow log diagnosis text with dynamic arguments
type ClusterSlowLogLocalizedText struct {
	Message *LocalizedDescription `json:"message,omitempty"`
	// Dynamic values used by the message
	Args map[string]string `json:"args,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogLocalizedText instantiates a new ClusterSlowLogLocalizedText object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogLocalizedText() *ClusterSlowLogLocalizedText {
	this := ClusterSlowLogLocalizedText{}
	return &this
}

// NewClusterSlowLogLocalizedTextWithDefaults instantiates a new ClusterSlowLogLocalizedText object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogLocalizedTextWithDefaults() *ClusterSlowLogLocalizedText {
	this := ClusterSlowLogLocalizedText{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterSlowLogLocalizedText) GetMessage() LocalizedDescription {
	if o == nil || o.Message == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogLocalizedText) GetMessageOk() (*LocalizedDescription, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterSlowLogLocalizedText) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given LocalizedDescription and assigns it to the Message field.
func (o *ClusterSlowLogLocalizedText) SetMessage(v LocalizedDescription) {
	o.Message = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *ClusterSlowLogLocalizedText) GetArgs() map[string]string {
	if o == nil || o.Args == nil {
		var ret map[string]string
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogLocalizedText) GetArgsOk() (*map[string]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return &o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *ClusterSlowLogLocalizedText) HasArgs() bool {
	return o != nil && o.Args != nil
}

// SetArgs gets a reference to the given map[string]string and assigns it to the Args field.
func (o *ClusterSlowLogLocalizedText) SetArgs(v map[string]string) {
	o.Args = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogLocalizedText) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogLocalizedText) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *LocalizedDescription `json:"message,omitempty"`
		Args    map[string]string     `json:"args,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "args"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Message != nil && all.Message.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Message = all.Message
	o.Args = all.Args

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
