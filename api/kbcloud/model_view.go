// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// View View is the schema for the views API

type View struct {
	// A label selector is a label query over a set of resources.
	InvolvedObjectsSelector *ViewInvolvedObjectsSelector `json:"involvedObjectsSelector,omitempty"`
	// targetResource specifies the type of resource the view is targeting.
	TargetResource *string `json:"targetResource,omitempty"`
	// viewDefRef is a reference to the definition of the view.
	ViewDefRef *string `json:"viewDefRef,omitempty"`
	// Events are a list of events related to the view.
	Events []View_event `json:"events,omitempty"`
	// Progress describes the overall progress of the view.
	Progress *string `json:"progress,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewView instantiates a new View object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewView() *View {
	this := View{}
	return &this
}

// NewViewWithDefaults instantiates a new View object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewViewWithDefaults() *View {
	this := View{}
	return &this
}

// GetInvolvedObjectsSelector returns the InvolvedObjectsSelector field value if set, zero value otherwise.
func (o *View) GetInvolvedObjectsSelector() ViewInvolvedObjectsSelector {
	if o == nil || o.InvolvedObjectsSelector == nil {
		var ret ViewInvolvedObjectsSelector
		return ret
	}
	return *o.InvolvedObjectsSelector
}

// GetInvolvedObjectsSelectorOk returns a tuple with the InvolvedObjectsSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetInvolvedObjectsSelectorOk() (*ViewInvolvedObjectsSelector, bool) {
	if o == nil || o.InvolvedObjectsSelector == nil {
		return nil, false
	}
	return o.InvolvedObjectsSelector, true
}

// HasInvolvedObjectsSelector returns a boolean if a field has been set.
func (o *View) HasInvolvedObjectsSelector() bool {
	return o != nil && o.InvolvedObjectsSelector != nil
}

// SetInvolvedObjectsSelector gets a reference to the given ViewInvolvedObjectsSelector and assigns it to the InvolvedObjectsSelector field.
func (o *View) SetInvolvedObjectsSelector(v ViewInvolvedObjectsSelector) {
	o.InvolvedObjectsSelector = &v
}

// GetTargetResource returns the TargetResource field value if set, zero value otherwise.
func (o *View) GetTargetResource() string {
	if o == nil || o.TargetResource == nil {
		var ret string
		return ret
	}
	return *o.TargetResource
}

// GetTargetResourceOk returns a tuple with the TargetResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetTargetResourceOk() (*string, bool) {
	if o == nil || o.TargetResource == nil {
		return nil, false
	}
	return o.TargetResource, true
}

// HasTargetResource returns a boolean if a field has been set.
func (o *View) HasTargetResource() bool {
	return o != nil && o.TargetResource != nil
}

// SetTargetResource gets a reference to the given string and assigns it to the TargetResource field.
func (o *View) SetTargetResource(v string) {
	o.TargetResource = &v
}

// GetViewDefRef returns the ViewDefRef field value if set, zero value otherwise.
func (o *View) GetViewDefRef() string {
	if o == nil || o.ViewDefRef == nil {
		var ret string
		return ret
	}
	return *o.ViewDefRef
}

// GetViewDefRefOk returns a tuple with the ViewDefRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetViewDefRefOk() (*string, bool) {
	if o == nil || o.ViewDefRef == nil {
		return nil, false
	}
	return o.ViewDefRef, true
}

// HasViewDefRef returns a boolean if a field has been set.
func (o *View) HasViewDefRef() bool {
	return o != nil && o.ViewDefRef != nil
}

// SetViewDefRef gets a reference to the given string and assigns it to the ViewDefRef field.
func (o *View) SetViewDefRef(v string) {
	o.ViewDefRef = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *View) GetEvents() []View_event {
	if o == nil || o.Events == nil {
		var ret []View_event
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetEventsOk() (*[]View_event, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *View) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []View_event and assigns it to the Events field.
func (o *View) SetEvents(v []View_event) {
	o.Events = v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *View) GetProgress() string {
	if o == nil || o.Progress == nil {
		var ret string
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetProgressOk() (*string, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *View) HasProgress() bool {
	return o != nil && o.Progress != nil
}

// SetProgress gets a reference to the given string and assigns it to the Progress field.
func (o *View) SetProgress(v string) {
	o.Progress = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o View) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.InvolvedObjectsSelector != nil {
		toSerialize["involvedObjectsSelector"] = o.InvolvedObjectsSelector
	}
	if o.TargetResource != nil {
		toSerialize["targetResource"] = o.TargetResource
	}
	if o.ViewDefRef != nil {
		toSerialize["viewDefRef"] = o.ViewDefRef
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *View) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InvolvedObjectsSelector *ViewInvolvedObjectsSelector `json:"involvedObjectsSelector,omitempty"`
		TargetResource          *string                      `json:"targetResource,omitempty"`
		ViewDefRef              *string                      `json:"viewDefRef,omitempty"`
		Events                  []View_event                 `json:"events,omitempty"`
		Progress                *string                      `json:"progress,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"involvedObjectsSelector", "targetResource", "viewDefRef", "events", "progress"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.InvolvedObjectsSelector != nil && all.InvolvedObjectsSelector.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InvolvedObjectsSelector = all.InvolvedObjectsSelector
	o.TargetResource = all.TargetResource
	o.ViewDefRef = all.ViewDefRef
	o.Events = all.Events
	o.Progress = all.Progress

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
