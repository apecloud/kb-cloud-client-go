// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RestoreStatus restore status
// NODESCRIPTION RestoreStatus
//
// Deprecated: This model is deprecated.
type RestoreStatus struct {
	// NODESCRIPTION Actions
	Actions []RestoreStatusActionsItem `json:"actions,omitempty"`
	// completion time
	CompletionTimestamp *string `json:"completionTimestamp,omitempty"`
	// NODESCRIPTION Conditions
	Conditions []RestoreStatusConditionsItem `json:"conditions,omitempty"`
	// restore phase
	Phase *string `json:"phase,omitempty"`
	// start time
	StartTimestamp *string `json:"startTimestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestoreStatus instantiates a new RestoreStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestoreStatus() *RestoreStatus {
	this := RestoreStatus{}
	return &this
}

// NewRestoreStatusWithDefaults instantiates a new RestoreStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestoreStatusWithDefaults() *RestoreStatus {
	this := RestoreStatus{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *RestoreStatus) GetActions() []RestoreStatusActionsItem {
	if o == nil || o.Actions == nil {
		var ret []RestoreStatusActionsItem
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatus) GetActionsOk() (*[]RestoreStatusActionsItem, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *RestoreStatus) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given []RestoreStatusActionsItem and assigns it to the Actions field.
func (o *RestoreStatus) SetActions(v []RestoreStatusActionsItem) {
	o.Actions = v
}

// GetCompletionTimestamp returns the CompletionTimestamp field value if set, zero value otherwise.
func (o *RestoreStatus) GetCompletionTimestamp() string {
	if o == nil || o.CompletionTimestamp == nil {
		var ret string
		return ret
	}
	return *o.CompletionTimestamp
}

// GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatus) GetCompletionTimestampOk() (*string, bool) {
	if o == nil || o.CompletionTimestamp == nil {
		return nil, false
	}
	return o.CompletionTimestamp, true
}

// HasCompletionTimestamp returns a boolean if a field has been set.
func (o *RestoreStatus) HasCompletionTimestamp() bool {
	return o != nil && o.CompletionTimestamp != nil
}

// SetCompletionTimestamp gets a reference to the given string and assigns it to the CompletionTimestamp field.
func (o *RestoreStatus) SetCompletionTimestamp(v string) {
	o.CompletionTimestamp = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *RestoreStatus) GetConditions() []RestoreStatusConditionsItem {
	if o == nil || o.Conditions == nil {
		var ret []RestoreStatusConditionsItem
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatus) GetConditionsOk() (*[]RestoreStatusConditionsItem, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *RestoreStatus) HasConditions() bool {
	return o != nil && o.Conditions != nil
}

// SetConditions gets a reference to the given []RestoreStatusConditionsItem and assigns it to the Conditions field.
func (o *RestoreStatus) SetConditions(v []RestoreStatusConditionsItem) {
	o.Conditions = v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *RestoreStatus) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatus) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *RestoreStatus) HasPhase() bool {
	return o != nil && o.Phase != nil
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *RestoreStatus) SetPhase(v string) {
	o.Phase = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *RestoreStatus) GetStartTimestamp() string {
	if o == nil || o.StartTimestamp == nil {
		var ret string
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatus) GetStartTimestampOk() (*string, bool) {
	if o == nil || o.StartTimestamp == nil {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *RestoreStatus) HasStartTimestamp() bool {
	return o != nil && o.StartTimestamp != nil
}

// SetStartTimestamp gets a reference to the given string and assigns it to the StartTimestamp field.
func (o *RestoreStatus) SetStartTimestamp(v string) {
	o.StartTimestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestoreStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.CompletionTimestamp != nil {
		toSerialize["completionTimestamp"] = o.CompletionTimestamp
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.StartTimestamp != nil {
		toSerialize["startTimestamp"] = o.StartTimestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestoreStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions             []RestoreStatusActionsItem    `json:"actions,omitempty"`
		CompletionTimestamp *string                       `json:"completionTimestamp,omitempty"`
		Conditions          []RestoreStatusConditionsItem `json:"conditions,omitempty"`
		Phase               *string                       `json:"phase,omitempty"`
		StartTimestamp      *string                       `json:"startTimestamp,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"actions", "completionTimestamp", "conditions", "phase", "startTimestamp"})
	} else {
		return err
	}
	o.Actions = all.Actions
	o.CompletionTimestamp = all.CompletionTimestamp
	o.Conditions = all.Conditions
	o.Phase = all.Phase
	o.StartTimestamp = all.StartTimestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
