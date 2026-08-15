// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ModeServiceRefVersionCompatibility Defines allowed source/ref service version combinations for a mode serviceRef.
// Version patterns match semantic version segments, so "3", "3.0", and
// "3.0.10" can match a whole major, minor, or patch family respectively.
type ModeServiceRefVersionCompatibility struct {
	// Source cluster version patterns. Empty means all source versions.
	//
	Version []string `json:"version,omitempty"`
	// Referenced cluster version patterns. Empty means all referenced versions.
	//
	RefVersion []string              `json:"refVersion,omitempty"`
	Message    *LocalizedDescription `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeServiceRefVersionCompatibility instantiates a new ModeServiceRefVersionCompatibility object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeServiceRefVersionCompatibility() *ModeServiceRefVersionCompatibility {
	this := ModeServiceRefVersionCompatibility{}
	return &this
}

// NewModeServiceRefVersionCompatibilityWithDefaults instantiates a new ModeServiceRefVersionCompatibility object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeServiceRefVersionCompatibilityWithDefaults() *ModeServiceRefVersionCompatibility {
	this := ModeServiceRefVersionCompatibility{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModeServiceRefVersionCompatibility) GetVersion() []string {
	if o == nil || o.Version == nil {
		var ret []string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRefVersionCompatibility) GetVersionOk() (*[]string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return &o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModeServiceRefVersionCompatibility) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given []string and assigns it to the Version field.
func (o *ModeServiceRefVersionCompatibility) SetVersion(v []string) {
	o.Version = v
}

// GetRefVersion returns the RefVersion field value if set, zero value otherwise.
func (o *ModeServiceRefVersionCompatibility) GetRefVersion() []string {
	if o == nil || o.RefVersion == nil {
		var ret []string
		return ret
	}
	return o.RefVersion
}

// GetRefVersionOk returns a tuple with the RefVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRefVersionCompatibility) GetRefVersionOk() (*[]string, bool) {
	if o == nil || o.RefVersion == nil {
		return nil, false
	}
	return &o.RefVersion, true
}

// HasRefVersion returns a boolean if a field has been set.
func (o *ModeServiceRefVersionCompatibility) HasRefVersion() bool {
	return o != nil && o.RefVersion != nil
}

// SetRefVersion gets a reference to the given []string and assigns it to the RefVersion field.
func (o *ModeServiceRefVersionCompatibility) SetRefVersion(v []string) {
	o.RefVersion = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ModeServiceRefVersionCompatibility) GetMessage() LocalizedDescription {
	if o == nil || o.Message == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRefVersionCompatibility) GetMessageOk() (*LocalizedDescription, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ModeServiceRefVersionCompatibility) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given LocalizedDescription and assigns it to the Message field.
func (o *ModeServiceRefVersionCompatibility) SetMessage(v LocalizedDescription) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeServiceRefVersionCompatibility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.RefVersion != nil {
		toSerialize["refVersion"] = o.RefVersion
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeServiceRefVersionCompatibility) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Version    []string              `json:"version,omitempty"`
		RefVersion []string              `json:"refVersion,omitempty"`
		Message    *LocalizedDescription `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"version", "refVersion", "message"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Version = all.Version
	o.RefVersion = all.RefVersion
	if all.Message != nil && all.Message.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
