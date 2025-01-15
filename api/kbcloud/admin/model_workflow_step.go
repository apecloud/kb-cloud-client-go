// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// WorkflowStep workflow step
type WorkflowStep struct {
	Name       *string             `json:"name,omitempty"`
	Phase      *string             `json:"phase,omitempty"`
	Message    *string             `json:"message,omitempty"`
	StartedAt  common.NullableTime `json:"startedAt,omitempty"`
	FinishedAt common.NullableTime `json:"finishedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowStep instantiates a new WorkflowStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowStep() *WorkflowStep {
	this := WorkflowStep{}
	return &this
}

// NewWorkflowStepWithDefaults instantiates a new WorkflowStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowStepWithDefaults() *WorkflowStep {
	this := WorkflowStep{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowStep) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowStep) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowStep) SetName(v string) {
	o.Name = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *WorkflowStep) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStep) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *WorkflowStep) HasPhase() bool {
	return o != nil && o.Phase != nil
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *WorkflowStep) SetPhase(v string) {
	o.Phase = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WorkflowStep) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStep) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WorkflowStep) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WorkflowStep) SetMessage(v string) {
	o.Message = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowStep) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WorkflowStep) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *WorkflowStep) HasStartedAt() bool {
	return o != nil && o.StartedAt.IsSet()
}

// SetStartedAt gets a reference to the given common.NullableTime and assigns it to the StartedAt field.
func (o *WorkflowStep) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// SetStartedAtNil sets the value for StartedAt to be an explicit nil.
func (o *WorkflowStep) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil.
func (o *WorkflowStep) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowStep) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *WorkflowStep) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *WorkflowStep) HasFinishedAt() bool {
	return o != nil && o.FinishedAt.IsSet()
}

// SetFinishedAt gets a reference to the given common.NullableTime and assigns it to the FinishedAt field.
func (o *WorkflowStep) SetFinishedAt(v time.Time) {
	o.FinishedAt.Set(&v)
}

// SetFinishedAtNil sets the value for FinishedAt to be an explicit nil.
func (o *WorkflowStep) SetFinishedAtNil() {
	o.FinishedAt.Set(nil)
}

// UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil.
func (o *WorkflowStep) UnsetFinishedAt() {
	o.FinishedAt.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.StartedAt.IsSet() {
		toSerialize["startedAt"] = o.StartedAt.Get()
	}
	if o.FinishedAt.IsSet() {
		toSerialize["finishedAt"] = o.FinishedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string             `json:"name,omitempty"`
		Phase      *string             `json:"phase,omitempty"`
		Message    *string             `json:"message,omitempty"`
		StartedAt  common.NullableTime `json:"startedAt,omitempty"`
		FinishedAt common.NullableTime `json:"finishedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "phase", "message", "startedAt", "finishedAt"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Phase = all.Phase
	o.Message = all.Message
	o.StartedAt = all.StartedAt
	o.FinishedAt = all.FinishedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
