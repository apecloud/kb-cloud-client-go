// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InspectionTaskV2 struct {
	Id      *int32  `json:"id,omitempty"`
	Creator *string `json:"creator,omitempty"`
	Status  *string `json:"status,omitempty"`
	// Specifies the supported engines for the inspection task.
	Engine      *InspectionSupportedEnginesV2 `json:"engine,omitempty"`
	OrgName     *string                       `json:"orgName,omitempty"`
	ClusterId   *string                       `json:"clusterID,omitempty"`
	ClusterName *string                       `json:"clusterName,omitempty"`
	EnvName     *string                       `json:"envName,omitempty"`
	EnvId       *string                       `json:"envID,omitempty"`
	NodeName    *string                       `json:"nodeName,omitempty"`
	IsAuto      *bool                         `json:"isAuto,omitempty"`
	Score       *int32                        `json:"score,omitempty"`
	Result      *string                       `json:"result,omitempty"`
	Items       []InspectionTaskItemV2        `json:"items,omitempty"`
	CreatedAt   *time.Time                    `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time                    `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspectionTaskV2 instantiates a new InspectionTaskV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspectionTaskV2() *InspectionTaskV2 {
	this := InspectionTaskV2{}
	return &this
}

// NewInspectionTaskV2WithDefaults instantiates a new InspectionTaskV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInspectionTaskV2WithDefaults() *InspectionTaskV2 {
	this := InspectionTaskV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InspectionTaskV2) SetId(v int32) {
	o.Id = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetCreator() string {
	if o == nil || o.Creator == nil {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetCreatorOk() (*string, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *InspectionTaskV2) SetCreator(v string) {
	o.Creator = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InspectionTaskV2) SetStatus(v string) {
	o.Status = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetEngine() InspectionSupportedEnginesV2 {
	if o == nil || o.Engine == nil {
		var ret InspectionSupportedEnginesV2
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetEngineOk() (*InspectionSupportedEnginesV2, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given InspectionSupportedEnginesV2 and assigns it to the Engine field.
func (o *InspectionTaskV2) SetEngine(v InspectionSupportedEnginesV2) {
	o.Engine = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *InspectionTaskV2) SetOrgName(v string) {
	o.OrgName = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *InspectionTaskV2) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *InspectionTaskV2) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetEnvName returns the EnvName field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetEnvName() string {
	if o == nil || o.EnvName == nil {
		var ret string
		return ret
	}
	return *o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetEnvNameOk() (*string, bool) {
	if o == nil || o.EnvName == nil {
		return nil, false
	}
	return o.EnvName, true
}

// HasEnvName returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasEnvName() bool {
	return o != nil && o.EnvName != nil
}

// SetEnvName gets a reference to the given string and assigns it to the EnvName field.
func (o *InspectionTaskV2) SetEnvName(v string) {
	o.EnvName = &v
}

// GetEnvId returns the EnvId field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetEnvId() string {
	if o == nil || o.EnvId == nil {
		var ret string
		return ret
	}
	return *o.EnvId
}

// GetEnvIdOk returns a tuple with the EnvId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetEnvIdOk() (*string, bool) {
	if o == nil || o.EnvId == nil {
		return nil, false
	}
	return o.EnvId, true
}

// HasEnvId returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasEnvId() bool {
	return o != nil && o.EnvId != nil
}

// SetEnvId gets a reference to the given string and assigns it to the EnvId field.
func (o *InspectionTaskV2) SetEnvId(v string) {
	o.EnvId = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *InspectionTaskV2) SetNodeName(v string) {
	o.NodeName = &v
}

// GetIsAuto returns the IsAuto field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetIsAuto() bool {
	if o == nil || o.IsAuto == nil {
		var ret bool
		return ret
	}
	return *o.IsAuto
}

// GetIsAutoOk returns a tuple with the IsAuto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetIsAutoOk() (*bool, bool) {
	if o == nil || o.IsAuto == nil {
		return nil, false
	}
	return o.IsAuto, true
}

// HasIsAuto returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasIsAuto() bool {
	return o != nil && o.IsAuto != nil
}

// SetIsAuto gets a reference to the given bool and assigns it to the IsAuto field.
func (o *InspectionTaskV2) SetIsAuto(v bool) {
	o.IsAuto = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetScore() int32 {
	if o == nil || o.Score == nil {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetScoreOk() (*int32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *InspectionTaskV2) SetScore(v int32) {
	o.Score = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *InspectionTaskV2) SetResult(v string) {
	o.Result = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetItems() []InspectionTaskItemV2 {
	if o == nil || o.Items == nil {
		var ret []InspectionTaskItemV2
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetItemsOk() (*[]InspectionTaskItemV2, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []InspectionTaskItemV2 and assigns it to the Items field.
func (o *InspectionTaskV2) SetItems(v []InspectionTaskItemV2) {
	o.Items = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InspectionTaskV2) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InspectionTaskV2) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskV2) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InspectionTaskV2) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InspectionTaskV2) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InspectionTaskV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.ClusterId != nil {
		toSerialize["clusterID"] = o.ClusterId
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.EnvName != nil {
		toSerialize["envName"] = o.EnvName
	}
	if o.EnvId != nil {
		toSerialize["envID"] = o.EnvId
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.IsAuto != nil {
		toSerialize["isAuto"] = o.IsAuto
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InspectionTaskV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *int32                        `json:"id,omitempty"`
		Creator     *string                       `json:"creator,omitempty"`
		Status      *string                       `json:"status,omitempty"`
		Engine      *InspectionSupportedEnginesV2 `json:"engine,omitempty"`
		OrgName     *string                       `json:"orgName,omitempty"`
		ClusterId   *string                       `json:"clusterID,omitempty"`
		ClusterName *string                       `json:"clusterName,omitempty"`
		EnvName     *string                       `json:"envName,omitempty"`
		EnvId       *string                       `json:"envID,omitempty"`
		NodeName    *string                       `json:"nodeName,omitempty"`
		IsAuto      *bool                         `json:"isAuto,omitempty"`
		Score       *int32                        `json:"score,omitempty"`
		Result      *string                       `json:"result,omitempty"`
		Items       []InspectionTaskItemV2        `json:"items,omitempty"`
		CreatedAt   *time.Time                    `json:"createdAt,omitempty"`
		UpdatedAt   *time.Time                    `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "creator", "status", "engine", "orgName", "clusterID", "clusterName", "envName", "envID", "nodeName", "isAuto", "score", "result", "items", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Creator = all.Creator
	o.Status = all.Status
	if all.Engine != nil && !all.Engine.IsValid() {
		hasInvalidField = true
	} else {
		o.Engine = all.Engine
	}
	o.OrgName = all.OrgName
	o.ClusterId = all.ClusterId
	o.ClusterName = all.ClusterName
	o.EnvName = all.EnvName
	o.EnvId = all.EnvId
	o.NodeName = all.NodeName
	o.IsAuto = all.IsAuto
	o.Score = all.Score
	o.Result = all.Result
	o.Items = all.Items
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
