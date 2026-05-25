// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentStartRunResponse struct {
	RunId          *string                   `json:"runId,omitempty"`
	HermesRunId    *string                   `json:"hermesRunId,omitempty"`
	ConversationId *string                   `json:"conversationId,omitempty"`
	OrgName        *string                   `json:"orgName,omitempty"`
	ClusterName    *string                   `json:"clusterName,omitempty"`
	Model          *string                   `json:"model,omitempty"`
	Status         *HermesAgentRunStatus     `json:"status,omitempty"`
	StatusReason   *string                   `json:"statusReason,omitempty"`
	CreatedAt      *time.Time                `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time                `json:"updatedAt,omitempty"`
	Runtime        *HermesAgentRuntimeStatus `json:"runtime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentStartRunResponse instantiates a new HermesAgentStartRunResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentStartRunResponse() *HermesAgentStartRunResponse {
	this := HermesAgentStartRunResponse{}
	return &this
}

// GetRunId returns the RunId field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetRunId() string {
	if o == nil || o.RunId == nil {
		var ret string
		return ret
	}
	return *o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetRunIdOk() (*string, bool) {
	if o == nil || o.RunId == nil {
		return nil, false
	}
	return o.RunId, true
}

// HasRunId returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasRunId() bool {
	return o != nil && o.RunId != nil
}

// SetRunId gets a reference to the given string and assigns it to the RunId field.
func (o *HermesAgentStartRunResponse) SetRunId(v string) {
	o.RunId = &v
}

// GetHermesRunId returns the HermesRunId field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetHermesRunId() string {
	if o == nil || o.HermesRunId == nil {
		var ret string
		return ret
	}
	return *o.HermesRunId
}

// GetHermesRunIdOk returns a tuple with the HermesRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetHermesRunIdOk() (*string, bool) {
	if o == nil || o.HermesRunId == nil {
		return nil, false
	}
	return o.HermesRunId, true
}

// HasHermesRunId returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasHermesRunId() bool {
	return o != nil && o.HermesRunId != nil
}

// SetHermesRunId gets a reference to the given string and assigns it to the HermesRunId field.
func (o *HermesAgentStartRunResponse) SetHermesRunId(v string) {
	o.HermesRunId = &v
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetConversationId() string {
	if o == nil || o.ConversationId == nil {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetConversationIdOk() (*string, bool) {
	if o == nil || o.ConversationId == nil {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasConversationId() bool {
	return o != nil && o.ConversationId != nil
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *HermesAgentStartRunResponse) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *HermesAgentStartRunResponse) SetOrgName(v string) {
	o.OrgName = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HermesAgentStartRunResponse) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *HermesAgentStartRunResponse) SetModel(v string) {
	o.Model = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetStatus() HermesAgentRunStatus {
	if o == nil || o.Status == nil {
		var ret HermesAgentRunStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetStatusOk() (*HermesAgentRunStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given HermesAgentRunStatus and assigns it to the Status field.
func (o *HermesAgentStartRunResponse) SetStatus(v HermesAgentRunStatus) {
	o.Status = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetStatusReason() string {
	if o == nil || o.StatusReason == nil {
		var ret string
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetStatusReasonOk() (*string, bool) {
	if o == nil || o.StatusReason == nil {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasStatusReason() bool {
	return o != nil && o.StatusReason != nil
}

// SetStatusReason gets a reference to the given string and assigns it to the StatusReason field.
func (o *HermesAgentStartRunResponse) SetStatusReason(v string) {
	o.StatusReason = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *HermesAgentStartRunResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *HermesAgentStartRunResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *HermesAgentStartRunResponse) GetRuntime() HermesAgentRuntimeStatus {
	if o == nil || o.Runtime == nil {
		var ret HermesAgentRuntimeStatus
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentStartRunResponse) GetRuntimeOk() (*HermesAgentRuntimeStatus, bool) {
	if o == nil || o.Runtime == nil {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *HermesAgentStartRunResponse) HasRuntime() bool {
	return o != nil && o.Runtime != nil
}

// SetRuntime gets a reference to the given HermesAgentRuntimeStatus and assigns it to the Runtime field.
func (o *HermesAgentStartRunResponse) SetRuntime(v HermesAgentRuntimeStatus) {
	o.Runtime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentStartRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.RunId != nil {
		toSerialize["runId"] = o.RunId
	}
	if o.HermesRunId != nil {
		toSerialize["hermesRunId"] = o.HermesRunId
	}
	if o.ConversationId != nil {
		toSerialize["conversationId"] = o.ConversationId
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusReason != nil {
		toSerialize["statusReason"] = o.StatusReason
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
	if o.Runtime != nil {
		toSerialize["runtime"] = o.Runtime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HermesAgentStartRunResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RunId          *string                   `json:"runId,omitempty"`
		HermesRunId    *string                   `json:"hermesRunId,omitempty"`
		ConversationId *string                   `json:"conversationId,omitempty"`
		OrgName        *string                   `json:"orgName,omitempty"`
		ClusterName    *string                   `json:"clusterName,omitempty"`
		Model          *string                   `json:"model,omitempty"`
		Status         *HermesAgentRunStatus     `json:"status,omitempty"`
		StatusReason   *string                   `json:"statusReason,omitempty"`
		CreatedAt      *time.Time                `json:"createdAt,omitempty"`
		UpdatedAt      *time.Time                `json:"updatedAt,omitempty"`
		Runtime        *HermesAgentRuntimeStatus `json:"runtime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runId", "hermesRunId", "conversationId", "orgName", "clusterName", "model", "status", "statusReason", "createdAt", "updatedAt", "runtime"})
	} else {
		return err
	}
	hasInvalidField := false
	o.RunId = all.RunId
	o.HermesRunId = all.HermesRunId
	o.ConversationId = all.ConversationId
	o.OrgName = all.OrgName
	o.ClusterName = all.ClusterName
	o.Model = all.Model
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.StatusReason = all.StatusReason
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	if all.Runtime != nil && all.Runtime.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Runtime = all.Runtime
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}
	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
