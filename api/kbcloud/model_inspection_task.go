// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InspectionTask struct {
	Id        *int32     `json:"id,omitempty"`
	OrgName   *string    `json:"orgName,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Creator   *string    `json:"creator,omitempty"`
	// Specifies the type of inspection task.
	ResourceType  *InspectionTaskType `json:"resourceType,omitempty"`
	ClusterId     *string             `json:"clusterID,omitempty"`
	ClusterName   *string             `json:"clusterName,omitempty"`
	ClusterEngine *string             `json:"clusterEngine,omitempty"`
	NodeName      *string             `json:"nodeName,omitempty"`
	Severity      *string             `json:"severity,omitempty"`
	Scenario      *string             `json:"scenario,omitempty"`
	Status        *string             `json:"status,omitempty"`
	Score         *string             `json:"score,omitempty"`
	Result        *string             `json:"result,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspectionTask instantiates a new InspectionTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspectionTask() *InspectionTask {
	this := InspectionTask{}
	var resourceType InspectionTaskType = InspectionTaskTypeCluster
	this.ResourceType = &resourceType
	return &this
}

// NewInspectionTaskWithDefaults instantiates a new InspectionTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInspectionTaskWithDefaults() *InspectionTask {
	this := InspectionTask{}
	var resourceType InspectionTaskType = InspectionTaskTypeCluster
	this.ResourceType = &resourceType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InspectionTask) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InspectionTask) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InspectionTask) SetId(v int32) {
	o.Id = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *InspectionTask) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *InspectionTask) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *InspectionTask) SetOrgName(v string) {
	o.OrgName = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InspectionTask) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InspectionTask) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InspectionTask) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InspectionTask) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InspectionTask) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InspectionTask) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *InspectionTask) GetCreator() string {
	if o == nil || o.Creator == nil {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetCreatorOk() (*string, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *InspectionTask) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *InspectionTask) SetCreator(v string) {
	o.Creator = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *InspectionTask) GetResourceType() InspectionTaskType {
	if o == nil || o.ResourceType == nil {
		var ret InspectionTaskType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetResourceTypeOk() (*InspectionTaskType, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *InspectionTask) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given InspectionTaskType and assigns it to the ResourceType field.
func (o *InspectionTask) SetResourceType(v InspectionTaskType) {
	o.ResourceType = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *InspectionTask) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *InspectionTask) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *InspectionTask) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *InspectionTask) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *InspectionTask) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *InspectionTask) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterEngine returns the ClusterEngine field value if set, zero value otherwise.
func (o *InspectionTask) GetClusterEngine() string {
	if o == nil || o.ClusterEngine == nil {
		var ret string
		return ret
	}
	return *o.ClusterEngine
}

// GetClusterEngineOk returns a tuple with the ClusterEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetClusterEngineOk() (*string, bool) {
	if o == nil || o.ClusterEngine == nil {
		return nil, false
	}
	return o.ClusterEngine, true
}

// HasClusterEngine returns a boolean if a field has been set.
func (o *InspectionTask) HasClusterEngine() bool {
	return o != nil && o.ClusterEngine != nil
}

// SetClusterEngine gets a reference to the given string and assigns it to the ClusterEngine field.
func (o *InspectionTask) SetClusterEngine(v string) {
	o.ClusterEngine = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *InspectionTask) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *InspectionTask) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *InspectionTask) SetNodeName(v string) {
	o.NodeName = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InspectionTask) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InspectionTask) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *InspectionTask) SetSeverity(v string) {
	o.Severity = &v
}

// GetScenario returns the Scenario field value if set, zero value otherwise.
func (o *InspectionTask) GetScenario() string {
	if o == nil || o.Scenario == nil {
		var ret string
		return ret
	}
	return *o.Scenario
}

// GetScenarioOk returns a tuple with the Scenario field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetScenarioOk() (*string, bool) {
	if o == nil || o.Scenario == nil {
		return nil, false
	}
	return o.Scenario, true
}

// HasScenario returns a boolean if a field has been set.
func (o *InspectionTask) HasScenario() bool {
	return o != nil && o.Scenario != nil
}

// SetScenario gets a reference to the given string and assigns it to the Scenario field.
func (o *InspectionTask) SetScenario(v string) {
	o.Scenario = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InspectionTask) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InspectionTask) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InspectionTask) SetStatus(v string) {
	o.Status = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *InspectionTask) GetScore() string {
	if o == nil || o.Score == nil {
		var ret string
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetScoreOk() (*string, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *InspectionTask) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given string and assigns it to the Score field.
func (o *InspectionTask) SetScore(v string) {
	o.Score = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InspectionTask) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTask) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InspectionTask) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *InspectionTask) SetResult(v string) {
	o.Result = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InspectionTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.ClusterId != nil {
		toSerialize["clusterID"] = o.ClusterId
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.ClusterEngine != nil {
		toSerialize["clusterEngine"] = o.ClusterEngine
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Scenario != nil {
		toSerialize["scenario"] = o.Scenario
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InspectionTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *int32              `json:"id,omitempty"`
		OrgName       *string             `json:"orgName,omitempty"`
		UpdatedAt     *time.Time          `json:"updatedAt,omitempty"`
		CreatedAt     *time.Time          `json:"createdAt,omitempty"`
		Creator       *string             `json:"creator,omitempty"`
		ResourceType  *InspectionTaskType `json:"resourceType,omitempty"`
		ClusterId     *string             `json:"clusterID,omitempty"`
		ClusterName   *string             `json:"clusterName,omitempty"`
		ClusterEngine *string             `json:"clusterEngine,omitempty"`
		NodeName      *string             `json:"nodeName,omitempty"`
		Severity      *string             `json:"severity,omitempty"`
		Scenario      *string             `json:"scenario,omitempty"`
		Status        *string             `json:"status,omitempty"`
		Score         *string             `json:"score,omitempty"`
		Result        *string             `json:"result,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgName", "updatedAt", "createdAt", "creator", "resourceType", "clusterID", "clusterName", "clusterEngine", "nodeName", "severity", "scenario", "status", "score", "result"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.OrgName = all.OrgName
	o.UpdatedAt = all.UpdatedAt
	o.CreatedAt = all.CreatedAt
	o.Creator = all.Creator
	if all.ResourceType != nil && !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = all.ResourceType
	}
	o.ClusterId = all.ClusterId
	o.ClusterName = all.ClusterName
	o.ClusterEngine = all.ClusterEngine
	o.NodeName = all.NodeName
	o.Severity = all.Severity
	o.Scenario = all.Scenario
	o.Status = all.Status
	o.Score = all.Score
	o.Result = all.Result

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
