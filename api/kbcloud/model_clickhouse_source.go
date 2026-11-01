// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseSource struct {
	// Source expression, for example clusterAllReplicas('ck', system.processes).
	Name *string `json:"name,omitempty"`
	// Logical source kind (system table or collection step).
	Kind *string `json:"kind,omitempty"`
	// Collection status (success, partial, failed).
	Status *string `json:"status,omitempty"`
	// Short human-readable status detail.
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSource instantiates a new ClickhouseSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSource() *ClickhouseSource {
	this := ClickhouseSource{}
	return &this
}

// NewClickhouseSourceWithDefaults instantiates a new ClickhouseSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSourceWithDefaults() *ClickhouseSource {
	this := ClickhouseSource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClickhouseSource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClickhouseSource) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClickhouseSource) SetName(v string) {
	o.Name = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClickhouseSource) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSource) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClickhouseSource) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClickhouseSource) SetKind(v string) {
	o.Kind = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClickhouseSource) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSource) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClickhouseSource) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClickhouseSource) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClickhouseSource) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSource) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClickhouseSource) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClickhouseSource) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
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
func (o *ClickhouseSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string `json:"name,omitempty"`
		Kind    *string `json:"kind,omitempty"`
		Status  *string `json:"status,omitempty"`
		Message *string `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "kind", "status", "message"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Kind = all.Kind
	o.Status = all.Status
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
