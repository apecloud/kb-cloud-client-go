// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AiAgentResourceScope Organization-scoped single-cluster Agent execution scope.
type AiAgentResourceScope struct {
	OrgName     *string           `json:"orgName,omitempty"`
	ProjectId   *string           `json:"projectId,omitempty"`
	ClusterName string            `json:"clusterName"`
	Namespace   *string           `json:"namespace,omitempty"`
	Engine      *string           `json:"engine,omitempty"`
	Component   *string           `json:"component,omitempty"`
	Pod         *string           `json:"pod,omitempty"`
	TimeRange   *AiAgentTimeRange `json:"timeRange,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentResourceScope instantiates a new AiAgentResourceScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentResourceScope(clusterName string) *AiAgentResourceScope {
	this := AiAgentResourceScope{}
	this.ClusterName = clusterName
	return &this
}

// NewAiAgentResourceScopeWithDefaults instantiates a new AiAgentResourceScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentResourceScopeWithDefaults() *AiAgentResourceScope {
	this := AiAgentResourceScope{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiAgentResourceScope) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentResourceScope) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiAgentResourceScope) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiAgentResourceScope) SetOrgName(v string) {
	o.OrgName = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AiAgentResourceScope) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentResourceScope) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AiAgentResourceScope) HasProjectId() bool {
	return o != nil && o.ProjectId != nil
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AiAgentResourceScope) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetClusterName returns the ClusterName field value.
func (o *AiAgentResourceScope) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *AiAgentResourceScope) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *AiAgentResourceScope) SetClusterName(v string) {
	o.ClusterName = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AiAgentResourceScope) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentResourceScope) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AiAgentResourceScope) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *AiAgentResourceScope) SetNamespace(v string) {
	o.Namespace = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *AiAgentResourceScope) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentResourceScope) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *AiAgentResourceScope) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *AiAgentResourceScope) SetEngine(v string) {
	o.Engine = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AiAgentResourceScope) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentResourceScope) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AiAgentResourceScope) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AiAgentResourceScope) SetComponent(v string) {
	o.Component = &v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *AiAgentResourceScope) GetPod() string {
	if o == nil || o.Pod == nil {
		var ret string
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentResourceScope) GetPodOk() (*string, bool) {
	if o == nil || o.Pod == nil {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *AiAgentResourceScope) HasPod() bool {
	return o != nil && o.Pod != nil
}

// SetPod gets a reference to the given string and assigns it to the Pod field.
func (o *AiAgentResourceScope) SetPod(v string) {
	o.Pod = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *AiAgentResourceScope) GetTimeRange() AiAgentTimeRange {
	if o == nil || o.TimeRange == nil {
		var ret AiAgentTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentResourceScope) GetTimeRangeOk() (*AiAgentTimeRange, bool) {
	if o == nil || o.TimeRange == nil {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *AiAgentResourceScope) HasTimeRange() bool {
	return o != nil && o.TimeRange != nil
}

// SetTimeRange gets a reference to the given AiAgentTimeRange and assigns it to the TimeRange field.
func (o *AiAgentResourceScope) SetTimeRange(v AiAgentTimeRange) {
	o.TimeRange = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentResourceScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	toSerialize["clusterName"] = o.ClusterName
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Pod != nil {
		toSerialize["pod"] = o.Pod
	}
	if o.TimeRange != nil {
		toSerialize["timeRange"] = o.TimeRange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentResourceScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName     *string           `json:"orgName,omitempty"`
		ProjectId   *string           `json:"projectId,omitempty"`
		ClusterName *string           `json:"clusterName"`
		Namespace   *string           `json:"namespace,omitempty"`
		Engine      *string           `json:"engine,omitempty"`
		Component   *string           `json:"component,omitempty"`
		Pod         *string           `json:"pod,omitempty"`
		TimeRange   *AiAgentTimeRange `json:"timeRange,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "projectId", "clusterName", "namespace", "engine", "component", "pod", "timeRange"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OrgName = all.OrgName
	o.ProjectId = all.ProjectId
	o.ClusterName = *all.ClusterName
	o.Namespace = all.Namespace
	o.Engine = all.Engine
	o.Component = all.Component
	o.Pod = all.Pod
	if all.TimeRange != nil && all.TimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeRange = all.TimeRange

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
