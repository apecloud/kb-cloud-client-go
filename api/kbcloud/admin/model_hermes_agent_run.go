// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentRun struct {
	RunId          string               `json:"runId"`
	HermesRunId    *string              `json:"hermesRunId,omitempty"`
	ConversationId string               `json:"conversationId"`
	OrgName        string               `json:"orgName"`
	ClusterName    *string              `json:"clusterName,omitempty"`
	Model          *string              `json:"model,omitempty"`
	Status         HermesAgentRunStatus `json:"status"`
	StatusReason   *string              `json:"statusReason,omitempty"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentRun instantiates a new HermesAgentRun object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentRun(runId string, conversationId string, orgName string, status HermesAgentRunStatus, createdAt time.Time, updatedAt time.Time) *HermesAgentRun {
	this := HermesAgentRun{}
	this.RunId = runId
	this.ConversationId = conversationId
	this.OrgName = orgName
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewHermesAgentRunWithDefaults instantiates a new HermesAgentRun object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHermesAgentRunWithDefaults() *HermesAgentRun {
	this := HermesAgentRun{}
	return &this
}

// GetRunId returns the RunId field value.
func (o *HermesAgentRun) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *HermesAgentRun) SetRunId(v string) {
	o.RunId = v
}

// GetHermesRunId returns the HermesRunId field value if set, zero value otherwise.
func (o *HermesAgentRun) GetHermesRunId() string {
	if o == nil || o.HermesRunId == nil {
		var ret string
		return ret
	}
	return *o.HermesRunId
}

// GetHermesRunIdOk returns a tuple with the HermesRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetHermesRunIdOk() (*string, bool) {
	if o == nil || o.HermesRunId == nil {
		return nil, false
	}
	return o.HermesRunId, true
}

// HasHermesRunId returns a boolean if a field has been set.
func (o *HermesAgentRun) HasHermesRunId() bool {
	return o != nil && o.HermesRunId != nil
}

// SetHermesRunId gets a reference to the given string and assigns it to the HermesRunId field.
func (o *HermesAgentRun) SetHermesRunId(v string) {
	o.HermesRunId = &v
}

// GetConversationId returns the ConversationId field value.
func (o *HermesAgentRun) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *HermesAgentRun) SetConversationId(v string) {
	o.ConversationId = v
}

// GetOrgName returns the OrgName field value.
func (o *HermesAgentRun) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *HermesAgentRun) SetOrgName(v string) {
	o.OrgName = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *HermesAgentRun) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *HermesAgentRun) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *HermesAgentRun) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *HermesAgentRun) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *HermesAgentRun) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *HermesAgentRun) SetModel(v string) {
	o.Model = &v
}

// GetStatus returns the Status field value.
func (o *HermesAgentRun) GetStatus() HermesAgentRunStatus {
	if o == nil {
		var ret HermesAgentRunStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetStatusOk() (*HermesAgentRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *HermesAgentRun) SetStatus(v HermesAgentRunStatus) {
	o.Status = v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *HermesAgentRun) GetStatusReason() string {
	if o == nil || o.StatusReason == nil {
		var ret string
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetStatusReasonOk() (*string, bool) {
	if o == nil || o.StatusReason == nil {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *HermesAgentRun) HasStatusReason() bool {
	return o != nil && o.StatusReason != nil
}

// SetStatusReason gets a reference to the given string and assigns it to the StatusReason field.
func (o *HermesAgentRun) SetStatusReason(v string) {
	o.StatusReason = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *HermesAgentRun) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *HermesAgentRun) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *HermesAgentRun) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *HermesAgentRun) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *HermesAgentRun) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["runId"] = o.RunId
	if o.HermesRunId != nil {
		toSerialize["hermesRunId"] = o.HermesRunId
	}
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["orgName"] = o.OrgName
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	toSerialize["status"] = o.Status
	if o.StatusReason != nil {
		toSerialize["statusReason"] = o.StatusReason
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HermesAgentRun) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RunId          *string               `json:"runId"`
		HermesRunId    *string               `json:"hermesRunId,omitempty"`
		ConversationId *string               `json:"conversationId"`
		OrgName        *string               `json:"orgName"`
		ClusterName    *string               `json:"clusterName,omitempty"`
		Model          *string               `json:"model,omitempty"`
		Status         *HermesAgentRunStatus `json:"status"`
		StatusReason   *string               `json:"statusReason,omitempty"`
		CreatedAt      *time.Time            `json:"createdAt"`
		UpdatedAt      *time.Time            `json:"updatedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runId", "hermesRunId", "conversationId", "orgName", "clusterName", "model", "status", "statusReason", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RunId = *all.RunId
	o.HermesRunId = all.HermesRunId
	o.ConversationId = *all.ConversationId
	o.OrgName = *all.OrgName
	o.ClusterName = all.ClusterName
	o.Model = all.Model
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.StatusReason = all.StatusReason
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
