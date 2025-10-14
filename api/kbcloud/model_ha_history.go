// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// HaHistory HA history is the records of HA operations, include `switchover` and `failover`
type HaHistory struct {
	ComponentName       string    `json:"componentName"`
	Type                string    `json:"type"`
	OldPrimary          string    `json:"oldPrimary"`
	NewPrimary          string    `json:"newPrimary"`
	Status              string    `json:"status"`
	Message             string    `json:"message"`
	OperatorId          *string   `json:"operatorID,omitempty"`
	OperatorName        *string   `json:"operatorName,omitempty"`
	OpsrequestName      *string   `json:"opsrequestName,omitempty"`
	OpsrequestNamespace *string   `json:"opsrequestNamespace,omitempty"`
	StartAt             time.Time `json:"startAt"`
	EndAt               time.Time `json:"endAt"`
	CreatedAt           time.Time `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHaHistory instantiates a new HaHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHaHistory(componentName string, typeVar string, oldPrimary string, newPrimary string, status string, message string, startAt time.Time, endAt time.Time, createdAt time.Time) *HaHistory {
	this := HaHistory{}
	this.ComponentName = componentName
	this.Type = typeVar
	this.OldPrimary = oldPrimary
	this.NewPrimary = newPrimary
	this.Status = status
	this.Message = message
	this.StartAt = startAt
	this.EndAt = endAt
	this.CreatedAt = createdAt
	return &this
}

// NewHaHistoryWithDefaults instantiates a new HaHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHaHistoryWithDefaults() *HaHistory {
	this := HaHistory{}
	return &this
}

// GetComponentName returns the ComponentName field value.
func (o *HaHistory) GetComponentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetComponentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentName, true
}

// SetComponentName sets field value.
func (o *HaHistory) SetComponentName(v string) {
	o.ComponentName = v
}

// GetType returns the Type field value.
func (o *HaHistory) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HaHistory) SetType(v string) {
	o.Type = v
}

// GetOldPrimary returns the OldPrimary field value.
func (o *HaHistory) GetOldPrimary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OldPrimary
}

// GetOldPrimaryOk returns a tuple with the OldPrimary field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetOldPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldPrimary, true
}

// SetOldPrimary sets field value.
func (o *HaHistory) SetOldPrimary(v string) {
	o.OldPrimary = v
}

// GetNewPrimary returns the NewPrimary field value.
func (o *HaHistory) GetNewPrimary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NewPrimary
}

// GetNewPrimaryOk returns a tuple with the NewPrimary field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetNewPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPrimary, true
}

// SetNewPrimary sets field value.
func (o *HaHistory) SetNewPrimary(v string) {
	o.NewPrimary = v
}

// GetStatus returns the Status field value.
func (o *HaHistory) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *HaHistory) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value.
func (o *HaHistory) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *HaHistory) SetMessage(v string) {
	o.Message = v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *HaHistory) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistory) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *HaHistory) HasOperatorId() bool {
	return o != nil && o.OperatorId != nil
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *HaHistory) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetOperatorName returns the OperatorName field value if set, zero value otherwise.
func (o *HaHistory) GetOperatorName() string {
	if o == nil || o.OperatorName == nil {
		var ret string
		return ret
	}
	return *o.OperatorName
}

// GetOperatorNameOk returns a tuple with the OperatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistory) GetOperatorNameOk() (*string, bool) {
	if o == nil || o.OperatorName == nil {
		return nil, false
	}
	return o.OperatorName, true
}

// HasOperatorName returns a boolean if a field has been set.
func (o *HaHistory) HasOperatorName() bool {
	return o != nil && o.OperatorName != nil
}

// SetOperatorName gets a reference to the given string and assigns it to the OperatorName field.
func (o *HaHistory) SetOperatorName(v string) {
	o.OperatorName = &v
}

// GetOpsrequestName returns the OpsrequestName field value if set, zero value otherwise.
func (o *HaHistory) GetOpsrequestName() string {
	if o == nil || o.OpsrequestName == nil {
		var ret string
		return ret
	}
	return *o.OpsrequestName
}

// GetOpsrequestNameOk returns a tuple with the OpsrequestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistory) GetOpsrequestNameOk() (*string, bool) {
	if o == nil || o.OpsrequestName == nil {
		return nil, false
	}
	return o.OpsrequestName, true
}

// HasOpsrequestName returns a boolean if a field has been set.
func (o *HaHistory) HasOpsrequestName() bool {
	return o != nil && o.OpsrequestName != nil
}

// SetOpsrequestName gets a reference to the given string and assigns it to the OpsrequestName field.
func (o *HaHistory) SetOpsrequestName(v string) {
	o.OpsrequestName = &v
}

// GetOpsrequestNamespace returns the OpsrequestNamespace field value if set, zero value otherwise.
func (o *HaHistory) GetOpsrequestNamespace() string {
	if o == nil || o.OpsrequestNamespace == nil {
		var ret string
		return ret
	}
	return *o.OpsrequestNamespace
}

// GetOpsrequestNamespaceOk returns a tuple with the OpsrequestNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaHistory) GetOpsrequestNamespaceOk() (*string, bool) {
	if o == nil || o.OpsrequestNamespace == nil {
		return nil, false
	}
	return o.OpsrequestNamespace, true
}

// HasOpsrequestNamespace returns a boolean if a field has been set.
func (o *HaHistory) HasOpsrequestNamespace() bool {
	return o != nil && o.OpsrequestNamespace != nil
}

// SetOpsrequestNamespace gets a reference to the given string and assigns it to the OpsrequestNamespace field.
func (o *HaHistory) SetOpsrequestNamespace(v string) {
	o.OpsrequestNamespace = &v
}

// GetStartAt returns the StartAt field value.
func (o *HaHistory) GetStartAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetStartAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value.
func (o *HaHistory) SetStartAt(v time.Time) {
	o.StartAt = v
}

// GetEndAt returns the EndAt field value.
func (o *HaHistory) GetEndAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetEndAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndAt, true
}

// SetEndAt sets field value.
func (o *HaHistory) SetEndAt(v time.Time) {
	o.EndAt = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *HaHistory) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *HaHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *HaHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HaHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["componentName"] = o.ComponentName
	toSerialize["type"] = o.Type
	toSerialize["oldPrimary"] = o.OldPrimary
	toSerialize["newPrimary"] = o.NewPrimary
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	if o.OperatorId != nil {
		toSerialize["operatorID"] = o.OperatorId
	}
	if o.OperatorName != nil {
		toSerialize["operatorName"] = o.OperatorName
	}
	if o.OpsrequestName != nil {
		toSerialize["opsrequestName"] = o.OpsrequestName
	}
	if o.OpsrequestNamespace != nil {
		toSerialize["opsrequestNamespace"] = o.OpsrequestNamespace
	}
	if o.StartAt.Nanosecond() == 0 {
		toSerialize["startAt"] = o.StartAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["startAt"] = o.StartAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.EndAt.Nanosecond() == 0 {
		toSerialize["endAt"] = o.EndAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["endAt"] = o.EndAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HaHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentName       *string    `json:"componentName"`
		Type                *string    `json:"type"`
		OldPrimary          *string    `json:"oldPrimary"`
		NewPrimary          *string    `json:"newPrimary"`
		Status              *string    `json:"status"`
		Message             *string    `json:"message"`
		OperatorId          *string    `json:"operatorID,omitempty"`
		OperatorName        *string    `json:"operatorName,omitempty"`
		OpsrequestName      *string    `json:"opsrequestName,omitempty"`
		OpsrequestNamespace *string    `json:"opsrequestNamespace,omitempty"`
		StartAt             *time.Time `json:"startAt"`
		EndAt               *time.Time `json:"endAt"`
		CreatedAt           *time.Time `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ComponentName == nil {
		return fmt.Errorf("required field componentName missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.OldPrimary == nil {
		return fmt.Errorf("required field oldPrimary missing")
	}
	if all.NewPrimary == nil {
		return fmt.Errorf("required field newPrimary missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.StartAt == nil {
		return fmt.Errorf("required field startAt missing")
	}
	if all.EndAt == nil {
		return fmt.Errorf("required field endAt missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"componentName", "type", "oldPrimary", "newPrimary", "status", "message", "operatorID", "operatorName", "opsrequestName", "opsrequestNamespace", "startAt", "endAt", "createdAt"})
	} else {
		return err
	}
	o.ComponentName = *all.ComponentName
	o.Type = *all.Type
	o.OldPrimary = *all.OldPrimary
	o.NewPrimary = *all.NewPrimary
	o.Status = *all.Status
	o.Message = *all.Message
	o.OperatorId = all.OperatorId
	o.OperatorName = all.OperatorName
	o.OpsrequestName = all.OpsrequestName
	o.OpsrequestNamespace = all.OpsrequestNamespace
	o.StartAt = *all.StartAt
	o.EndAt = *all.EndAt
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
