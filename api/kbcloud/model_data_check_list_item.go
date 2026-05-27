// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataCheckListItem struct {
	CheckId   *string `json:"checkID,omitempty"`
	CheckName *string `json:"checkName,omitempty"`
	// Type of check to run. A standalone check runs exactly one type.
	CheckType          *DataCheckType        `json:"checkType,omitempty"`
	CheckStatus        *DataCheckStatus      `json:"checkStatus,omitempty"`
	CheckSummary       *DataCheckListSummary `json:"checkSummary,omitempty"`
	EnvironmentId      *string               `json:"environmentID,omitempty"`
	EnvironmentName    *string               `json:"environmentName,omitempty"`
	Project            *string               `json:"project,omitempty"`
	ReplicationObjects *DataChannelObject    `json:"replicationObjects,omitempty"`
	ErrorMessage       common.NullableString `json:"errorMessage,omitempty"`
	CreatedAt          *time.Time            `json:"createdAt,omitempty"`
	UpdatedAt          *time.Time            `json:"updatedAt,omitempty"`
	FinishedAt         common.NullableTime   `json:"finishedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataCheckListItem instantiates a new DataCheckListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataCheckListItem() *DataCheckListItem {
	this := DataCheckListItem{}
	return &this
}

// NewDataCheckListItemWithDefaults instantiates a new DataCheckListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataCheckListItemWithDefaults() *DataCheckListItem {
	this := DataCheckListItem{}
	return &this
}

// GetCheckId returns the CheckId field value if set, zero value otherwise.
func (o *DataCheckListItem) GetCheckId() string {
	if o == nil || o.CheckId == nil {
		var ret string
		return ret
	}
	return *o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetCheckIdOk() (*string, bool) {
	if o == nil || o.CheckId == nil {
		return nil, false
	}
	return o.CheckId, true
}

// HasCheckId returns a boolean if a field has been set.
func (o *DataCheckListItem) HasCheckId() bool {
	return o != nil && o.CheckId != nil
}

// SetCheckId gets a reference to the given string and assigns it to the CheckId field.
func (o *DataCheckListItem) SetCheckId(v string) {
	o.CheckId = &v
}

// GetCheckName returns the CheckName field value if set, zero value otherwise.
func (o *DataCheckListItem) GetCheckName() string {
	if o == nil || o.CheckName == nil {
		var ret string
		return ret
	}
	return *o.CheckName
}

// GetCheckNameOk returns a tuple with the CheckName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetCheckNameOk() (*string, bool) {
	if o == nil || o.CheckName == nil {
		return nil, false
	}
	return o.CheckName, true
}

// HasCheckName returns a boolean if a field has been set.
func (o *DataCheckListItem) HasCheckName() bool {
	return o != nil && o.CheckName != nil
}

// SetCheckName gets a reference to the given string and assigns it to the CheckName field.
func (o *DataCheckListItem) SetCheckName(v string) {
	o.CheckName = &v
}

// GetCheckType returns the CheckType field value if set, zero value otherwise.
func (o *DataCheckListItem) GetCheckType() DataCheckType {
	if o == nil || o.CheckType == nil {
		var ret DataCheckType
		return ret
	}
	return *o.CheckType
}

// GetCheckTypeOk returns a tuple with the CheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetCheckTypeOk() (*DataCheckType, bool) {
	if o == nil || o.CheckType == nil {
		return nil, false
	}
	return o.CheckType, true
}

// HasCheckType returns a boolean if a field has been set.
func (o *DataCheckListItem) HasCheckType() bool {
	return o != nil && o.CheckType != nil
}

// SetCheckType gets a reference to the given DataCheckType and assigns it to the CheckType field.
func (o *DataCheckListItem) SetCheckType(v DataCheckType) {
	o.CheckType = &v
}

// GetCheckStatus returns the CheckStatus field value if set, zero value otherwise.
func (o *DataCheckListItem) GetCheckStatus() DataCheckStatus {
	if o == nil || o.CheckStatus == nil {
		var ret DataCheckStatus
		return ret
	}
	return *o.CheckStatus
}

// GetCheckStatusOk returns a tuple with the CheckStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetCheckStatusOk() (*DataCheckStatus, bool) {
	if o == nil || o.CheckStatus == nil {
		return nil, false
	}
	return o.CheckStatus, true
}

// HasCheckStatus returns a boolean if a field has been set.
func (o *DataCheckListItem) HasCheckStatus() bool {
	return o != nil && o.CheckStatus != nil
}

// SetCheckStatus gets a reference to the given DataCheckStatus and assigns it to the CheckStatus field.
func (o *DataCheckListItem) SetCheckStatus(v DataCheckStatus) {
	o.CheckStatus = &v
}

// GetCheckSummary returns the CheckSummary field value if set, zero value otherwise.
func (o *DataCheckListItem) GetCheckSummary() DataCheckListSummary {
	if o == nil || o.CheckSummary == nil {
		var ret DataCheckListSummary
		return ret
	}
	return *o.CheckSummary
}

// GetCheckSummaryOk returns a tuple with the CheckSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetCheckSummaryOk() (*DataCheckListSummary, bool) {
	if o == nil || o.CheckSummary == nil {
		return nil, false
	}
	return o.CheckSummary, true
}

// HasCheckSummary returns a boolean if a field has been set.
func (o *DataCheckListItem) HasCheckSummary() bool {
	return o != nil && o.CheckSummary != nil
}

// SetCheckSummary gets a reference to the given DataCheckListSummary and assigns it to the CheckSummary field.
func (o *DataCheckListItem) SetCheckSummary(v DataCheckListSummary) {
	o.CheckSummary = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *DataCheckListItem) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *DataCheckListItem) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *DataCheckListItem) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *DataCheckListItem) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *DataCheckListItem) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *DataCheckListItem) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *DataCheckListItem) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *DataCheckListItem) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *DataCheckListItem) SetProject(v string) {
	o.Project = &v
}

// GetReplicationObjects returns the ReplicationObjects field value if set, zero value otherwise.
func (o *DataCheckListItem) GetReplicationObjects() DataChannelObject {
	if o == nil || o.ReplicationObjects == nil {
		var ret DataChannelObject
		return ret
	}
	return *o.ReplicationObjects
}

// GetReplicationObjectsOk returns a tuple with the ReplicationObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetReplicationObjectsOk() (*DataChannelObject, bool) {
	if o == nil || o.ReplicationObjects == nil {
		return nil, false
	}
	return o.ReplicationObjects, true
}

// HasReplicationObjects returns a boolean if a field has been set.
func (o *DataCheckListItem) HasReplicationObjects() bool {
	return o != nil && o.ReplicationObjects != nil
}

// SetReplicationObjects gets a reference to the given DataChannelObject and assigns it to the ReplicationObjects field.
func (o *DataCheckListItem) SetReplicationObjects(v DataChannelObject) {
	o.ReplicationObjects = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckListItem) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckListItem) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *DataCheckListItem) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage.IsSet()
}

// SetErrorMessage gets a reference to the given common.NullableString and assigns it to the ErrorMessage field.
func (o *DataCheckListItem) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil.
func (o *DataCheckListItem) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil.
func (o *DataCheckListItem) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DataCheckListItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DataCheckListItem) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DataCheckListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DataCheckListItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DataCheckListItem) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DataCheckListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckListItem) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckListItem) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DataCheckListItem) HasFinishedAt() bool {
	return o != nil && o.FinishedAt.IsSet()
}

// SetFinishedAt gets a reference to the given common.NullableTime and assigns it to the FinishedAt field.
func (o *DataCheckListItem) SetFinishedAt(v time.Time) {
	o.FinishedAt.Set(&v)
}

// SetFinishedAtNil sets the value for FinishedAt to be an explicit nil.
func (o *DataCheckListItem) SetFinishedAtNil() {
	o.FinishedAt.Set(nil)
}

// UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil.
func (o *DataCheckListItem) UnsetFinishedAt() {
	o.FinishedAt.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DataCheckListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CheckId != nil {
		toSerialize["checkID"] = o.CheckId
	}
	if o.CheckName != nil {
		toSerialize["checkName"] = o.CheckName
	}
	if o.CheckType != nil {
		toSerialize["checkType"] = o.CheckType
	}
	if o.CheckStatus != nil {
		toSerialize["checkStatus"] = o.CheckStatus
	}
	if o.CheckSummary != nil {
		toSerialize["checkSummary"] = o.CheckSummary
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentID"] = o.EnvironmentId
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.ReplicationObjects != nil {
		toSerialize["replicationObjects"] = o.ReplicationObjects
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *DataCheckListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckId            *string               `json:"checkID,omitempty"`
		CheckName          *string               `json:"checkName,omitempty"`
		CheckType          *DataCheckType        `json:"checkType,omitempty"`
		CheckStatus        *DataCheckStatus      `json:"checkStatus,omitempty"`
		CheckSummary       *DataCheckListSummary `json:"checkSummary,omitempty"`
		EnvironmentId      *string               `json:"environmentID,omitempty"`
		EnvironmentName    *string               `json:"environmentName,omitempty"`
		Project            *string               `json:"project,omitempty"`
		ReplicationObjects *DataChannelObject    `json:"replicationObjects,omitempty"`
		ErrorMessage       common.NullableString `json:"errorMessage,omitempty"`
		CreatedAt          *time.Time            `json:"createdAt,omitempty"`
		UpdatedAt          *time.Time            `json:"updatedAt,omitempty"`
		FinishedAt         common.NullableTime   `json:"finishedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"checkID", "checkName", "checkType", "checkStatus", "checkSummary", "environmentID", "environmentName", "project", "replicationObjects", "errorMessage", "createdAt", "updatedAt", "finishedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CheckId = all.CheckId
	o.CheckName = all.CheckName
	if all.CheckType != nil && !all.CheckType.IsValid() {
		hasInvalidField = true
	} else {
		o.CheckType = all.CheckType
	}
	if all.CheckStatus != nil && !all.CheckStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.CheckStatus = all.CheckStatus
	}
	if all.CheckSummary != nil && all.CheckSummary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CheckSummary = all.CheckSummary
	o.EnvironmentId = all.EnvironmentId
	o.EnvironmentName = all.EnvironmentName
	o.Project = all.Project
	if all.ReplicationObjects != nil && all.ReplicationObjects.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReplicationObjects = all.ReplicationObjects
	o.ErrorMessage = all.ErrorMessage
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.FinishedAt = all.FinishedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
