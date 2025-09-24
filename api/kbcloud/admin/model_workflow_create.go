// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type WorkflowCreate struct {
	// workflow type
	Type         *WorkflowType `json:"type,omitempty"`
	Version      *string       `json:"version,omitempty"`
	OteldVersion *string       `json:"oteldVersion,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowCreate instantiates a new WorkflowCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowCreate() *WorkflowCreate {
	this := WorkflowCreate{}
	return &this
}

// NewWorkflowCreateWithDefaults instantiates a new WorkflowCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowCreateWithDefaults() *WorkflowCreate {
	this := WorkflowCreate{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowCreate) GetType() WorkflowType {
	if o == nil || o.Type == nil {
		var ret WorkflowType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCreate) GetTypeOk() (*WorkflowType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowCreate) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given WorkflowType and assigns it to the Type field.
func (o *WorkflowCreate) SetType(v WorkflowType) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowCreate) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCreate) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowCreate) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *WorkflowCreate) SetVersion(v string) {
	o.Version = &v
}

// GetOteldVersion returns the OteldVersion field value if set, zero value otherwise.
func (o *WorkflowCreate) GetOteldVersion() string {
	if o == nil || o.OteldVersion == nil {
		var ret string
		return ret
	}
	return *o.OteldVersion
}

// GetOteldVersionOk returns a tuple with the OteldVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCreate) GetOteldVersionOk() (*string, bool) {
	if o == nil || o.OteldVersion == nil {
		return nil, false
	}
	return o.OteldVersion, true
}

// HasOteldVersion returns a boolean if a field has been set.
func (o *WorkflowCreate) HasOteldVersion() bool {
	return o != nil && o.OteldVersion != nil
}

// SetOteldVersion gets a reference to the given string and assigns it to the OteldVersion field.
func (o *WorkflowCreate) SetOteldVersion(v string) {
	o.OteldVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.OteldVersion != nil {
		toSerialize["oteldVersion"] = o.OteldVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type         *WorkflowType `json:"type,omitempty"`
		Version      *string       `json:"version,omitempty"`
		OteldVersion *string       `json:"oteldVersion,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "version", "oteldVersion"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Version = all.Version
	o.OteldVersion = all.OteldVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
