// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Event event is the information of operation event
type Event struct {
	// ID of the event
	Id *string `json:"id,omitempty"`
	// ID of the resource
	ResourceId *string `json:"resourceId,omitempty"`
	// Type of the resource
	ResourceType *EventResourceType `json:"resourceType,omitempty"`
	// Name of the resource
	ResourceName *string `json:"resourceName,omitempty"`
	// operator of the event, if source is user, operator is user name; if source is system, operator is system name
	Operator *string `json:"operator,omitempty"`
	// The user ID of the operator
	OperatorId *string `json:"operatorId,omitempty"`
	// Details will include the extra event info, such as update cluster which field, OpsRequest content etc
	Details *string `json:"details,omitempty"`
	// hasTask is true if the event has a task
	HasTask *bool `json:"hasTask,omitempty"`
	// result of the operation event
	Result *string `json:"result,omitempty"`
	// Event name is OpsRequest name or cluster operation name
	EventName *string `json:"eventName,omitempty"`
	// Display name is custom operation name
	DisplayName *string `json:"displayName,omitempty"`
	// The org of the event
	OrgName *string `json:"orgName,omitempty"`
	// result status of the operation event
	ResultStatus *EventResultStatus `json:"resultStatus,omitempty"`
	// event source
	Source *EventSource `json:"source,omitempty"`
	// event end time
	End *time.Time `json:"end,omitempty"`
	// event start time
	Start *time.Time `json:"start,omitempty"`
	// event created time
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEvent instantiates a new Event object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEvent() *Event {
	this := Event{}
	return &this
}

// NewEventWithDefaults instantiates a new Event object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Event) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Event) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Event) SetId(v string) {
	o.Id = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *Event) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *Event) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *Event) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *Event) GetResourceType() EventResourceType {
	if o == nil || o.ResourceType == nil {
		var ret EventResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetResourceTypeOk() (*EventResourceType, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Event) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given EventResourceType and assigns it to the ResourceType field.
func (o *Event) SetResourceType(v EventResourceType) {
	o.ResourceType = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *Event) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *Event) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *Event) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *Event) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Event) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *Event) SetOperator(v string) {
	o.Operator = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *Event) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *Event) HasOperatorId() bool {
	return o != nil && o.OperatorId != nil
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *Event) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Event) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Event) HasDetails() bool {
	return o != nil && o.Details != nil
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *Event) SetDetails(v string) {
	o.Details = &v
}

// GetHasTask returns the HasTask field value if set, zero value otherwise.
func (o *Event) GetHasTask() bool {
	if o == nil || o.HasTask == nil {
		var ret bool
		return ret
	}
	return *o.HasTask
}

// GetHasTaskOk returns a tuple with the HasTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetHasTaskOk() (*bool, bool) {
	if o == nil || o.HasTask == nil {
		return nil, false
	}
	return o.HasTask, true
}

// HasHasTask returns a boolean if a field has been set.
func (o *Event) HasHasTask() bool {
	return o != nil && o.HasTask != nil
}

// SetHasTask gets a reference to the given bool and assigns it to the HasTask field.
func (o *Event) SetHasTask(v bool) {
	o.HasTask = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *Event) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *Event) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *Event) SetResult(v string) {
	o.Result = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *Event) GetEventName() string {
	if o == nil || o.EventName == nil {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEventNameOk() (*string, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *Event) HasEventName() bool {
	return o != nil && o.EventName != nil
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *Event) SetEventName(v string) {
	o.EventName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Event) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Event) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Event) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Event) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Event) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Event) SetOrgName(v string) {
	o.OrgName = &v
}

// GetResultStatus returns the ResultStatus field value if set, zero value otherwise.
func (o *Event) GetResultStatus() EventResultStatus {
	if o == nil || o.ResultStatus == nil {
		var ret EventResultStatus
		return ret
	}
	return *o.ResultStatus
}

// GetResultStatusOk returns a tuple with the ResultStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetResultStatusOk() (*EventResultStatus, bool) {
	if o == nil || o.ResultStatus == nil {
		return nil, false
	}
	return o.ResultStatus, true
}

// HasResultStatus returns a boolean if a field has been set.
func (o *Event) HasResultStatus() bool {
	return o != nil && o.ResultStatus != nil
}

// SetResultStatus gets a reference to the given EventResultStatus and assigns it to the ResultStatus field.
func (o *Event) SetResultStatus(v EventResultStatus) {
	o.ResultStatus = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Event) GetSource() EventSource {
	if o == nil || o.Source == nil {
		var ret EventSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSourceOk() (*EventSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Event) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given EventSource and assigns it to the Source field.
func (o *Event) SetSource(v EventSource) {
	o.Source = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Event) GetEnd() time.Time {
	if o == nil || o.End == nil {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEndOk() (*time.Time, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Event) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *Event) SetEnd(v time.Time) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Event) GetStart() time.Time {
	if o == nil || o.Start == nil {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStartOk() (*time.Time, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Event) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *Event) SetStart(v time.Time) {
	o.Start = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Event) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Event) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Event) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.ResourceName != nil {
		toSerialize["resourceName"] = o.ResourceName
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.HasTask != nil {
		toSerialize["hasTask"] = o.HasTask
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.EventName != nil {
		toSerialize["eventName"] = o.EventName
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.ResultStatus != nil {
		toSerialize["resultStatus"] = o.ResultStatus
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.End != nil {
		if o.End.Nanosecond() == 0 {
			toSerialize["end"] = o.End.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end"] = o.End.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Start != nil {
		if o.Start.Nanosecond() == 0 {
			toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Event) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           *string            `json:"id,omitempty"`
		ResourceId   *string            `json:"resourceId,omitempty"`
		ResourceType *EventResourceType `json:"resourceType,omitempty"`
		ResourceName *string            `json:"resourceName,omitempty"`
		Operator     *string            `json:"operator,omitempty"`
		OperatorId   *string            `json:"operatorId,omitempty"`
		Details      *string            `json:"details,omitempty"`
		HasTask      *bool              `json:"hasTask,omitempty"`
		Result       *string            `json:"result,omitempty"`
		EventName    *string            `json:"eventName,omitempty"`
		DisplayName  *string            `json:"displayName,omitempty"`
		OrgName      *string            `json:"orgName,omitempty"`
		ResultStatus *EventResultStatus `json:"resultStatus,omitempty"`
		Source       *EventSource       `json:"source,omitempty"`
		End          *time.Time         `json:"end,omitempty"`
		Start        *time.Time         `json:"start,omitempty"`
		CreatedAt    *time.Time         `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "resourceId", "resourceType", "resourceName", "operator", "operatorId", "details", "hasTask", "result", "eventName", "displayName", "orgName", "resultStatus", "source", "end", "start", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.ResourceId = all.ResourceId
	if all.ResourceType != nil && !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = all.ResourceType
	}
	o.ResourceName = all.ResourceName
	o.Operator = all.Operator
	o.OperatorId = all.OperatorId
	o.Details = all.Details
	o.HasTask = all.HasTask
	o.Result = all.Result
	o.EventName = all.EventName
	o.DisplayName = all.DisplayName
	o.OrgName = all.OrgName
	if all.ResultStatus != nil && !all.ResultStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ResultStatus = all.ResultStatus
	}
	if all.Source != nil && !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = all.Source
	}
	o.End = all.End
	o.Start = all.Start
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
