// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InspectionTaskCreate struct {
	Creator     *string `json:"creator,omitempty"`
	Engine      *string `json:"engine,omitempty"`
	OrgName     *string `json:"orgName,omitempty"`
	ClusterId   *string `json:"clusterID,omitempty"`
	ClusterName *string `json:"clusterName,omitempty"`
	EnvName     *string `json:"envName,omitempty"`
	EnvId       *string `json:"envID,omitempty"`
	// Node name(s) for inspection. Multiple nodes can be specified as a comma-separated string (e.g. "node1,node2,node3").
	NodeName       *string    `json:"nodeName,omitempty"`
	IsAuto         *bool      `json:"isAuto,omitempty"`
	SavedDays      *int32     `json:"savedDays,omitempty"`
	TimeRangeStart *time.Time `json:"timeRangeStart,omitempty"`
	TimeRangeEnd   *time.Time `json:"timeRangeEnd,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspectionTaskCreate instantiates a new InspectionTaskCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspectionTaskCreate() *InspectionTaskCreate {
	this := InspectionTaskCreate{}
	return &this
}

// NewInspectionTaskCreateWithDefaults instantiates a new InspectionTaskCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInspectionTaskCreateWithDefaults() *InspectionTaskCreate {
	this := InspectionTaskCreate{}
	return &this
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetCreator() string {
	if o == nil || o.Creator == nil {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetCreatorOk() (*string, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *InspectionTaskCreate) SetCreator(v string) {
	o.Creator = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *InspectionTaskCreate) SetEngine(v string) {
	o.Engine = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *InspectionTaskCreate) SetOrgName(v string) {
	o.OrgName = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *InspectionTaskCreate) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *InspectionTaskCreate) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetEnvName returns the EnvName field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetEnvName() string {
	if o == nil || o.EnvName == nil {
		var ret string
		return ret
	}
	return *o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetEnvNameOk() (*string, bool) {
	if o == nil || o.EnvName == nil {
		return nil, false
	}
	return o.EnvName, true
}

// HasEnvName returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasEnvName() bool {
	return o != nil && o.EnvName != nil
}

// SetEnvName gets a reference to the given string and assigns it to the EnvName field.
func (o *InspectionTaskCreate) SetEnvName(v string) {
	o.EnvName = &v
}

// GetEnvId returns the EnvId field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetEnvId() string {
	if o == nil || o.EnvId == nil {
		var ret string
		return ret
	}
	return *o.EnvId
}

// GetEnvIdOk returns a tuple with the EnvId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetEnvIdOk() (*string, bool) {
	if o == nil || o.EnvId == nil {
		return nil, false
	}
	return o.EnvId, true
}

// HasEnvId returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasEnvId() bool {
	return o != nil && o.EnvId != nil
}

// SetEnvId gets a reference to the given string and assigns it to the EnvId field.
func (o *InspectionTaskCreate) SetEnvId(v string) {
	o.EnvId = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *InspectionTaskCreate) SetNodeName(v string) {
	o.NodeName = &v
}

// GetIsAuto returns the IsAuto field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetIsAuto() bool {
	if o == nil || o.IsAuto == nil {
		var ret bool
		return ret
	}
	return *o.IsAuto
}

// GetIsAutoOk returns a tuple with the IsAuto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetIsAutoOk() (*bool, bool) {
	if o == nil || o.IsAuto == nil {
		return nil, false
	}
	return o.IsAuto, true
}

// HasIsAuto returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasIsAuto() bool {
	return o != nil && o.IsAuto != nil
}

// SetIsAuto gets a reference to the given bool and assigns it to the IsAuto field.
func (o *InspectionTaskCreate) SetIsAuto(v bool) {
	o.IsAuto = &v
}

// GetSavedDays returns the SavedDays field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetSavedDays() int32 {
	if o == nil || o.SavedDays == nil {
		var ret int32
		return ret
	}
	return *o.SavedDays
}

// GetSavedDaysOk returns a tuple with the SavedDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetSavedDaysOk() (*int32, bool) {
	if o == nil || o.SavedDays == nil {
		return nil, false
	}
	return o.SavedDays, true
}

// HasSavedDays returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasSavedDays() bool {
	return o != nil && o.SavedDays != nil
}

// SetSavedDays gets a reference to the given int32 and assigns it to the SavedDays field.
func (o *InspectionTaskCreate) SetSavedDays(v int32) {
	o.SavedDays = &v
}

// GetTimeRangeStart returns the TimeRangeStart field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetTimeRangeStart() time.Time {
	if o == nil || o.TimeRangeStart == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeRangeStart
}

// GetTimeRangeStartOk returns a tuple with the TimeRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetTimeRangeStartOk() (*time.Time, bool) {
	if o == nil || o.TimeRangeStart == nil {
		return nil, false
	}
	return o.TimeRangeStart, true
}

// HasTimeRangeStart returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasTimeRangeStart() bool {
	return o != nil && o.TimeRangeStart != nil
}

// SetTimeRangeStart gets a reference to the given time.Time and assigns it to the TimeRangeStart field.
func (o *InspectionTaskCreate) SetTimeRangeStart(v time.Time) {
	o.TimeRangeStart = &v
}

// GetTimeRangeEnd returns the TimeRangeEnd field value if set, zero value otherwise.
func (o *InspectionTaskCreate) GetTimeRangeEnd() time.Time {
	if o == nil || o.TimeRangeEnd == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeRangeEnd
}

// GetTimeRangeEndOk returns a tuple with the TimeRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskCreate) GetTimeRangeEndOk() (*time.Time, bool) {
	if o == nil || o.TimeRangeEnd == nil {
		return nil, false
	}
	return o.TimeRangeEnd, true
}

// HasTimeRangeEnd returns a boolean if a field has been set.
func (o *InspectionTaskCreate) HasTimeRangeEnd() bool {
	return o != nil && o.TimeRangeEnd != nil
}

// SetTimeRangeEnd gets a reference to the given time.Time and assigns it to the TimeRangeEnd field.
func (o *InspectionTaskCreate) SetTimeRangeEnd(v time.Time) {
	o.TimeRangeEnd = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InspectionTaskCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
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
	if o.SavedDays != nil {
		toSerialize["savedDays"] = o.SavedDays
	}
	if o.TimeRangeStart != nil {
		if o.TimeRangeStart.Nanosecond() == 0 {
			toSerialize["timeRangeStart"] = o.TimeRangeStart.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timeRangeStart"] = o.TimeRangeStart.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.TimeRangeEnd != nil {
		if o.TimeRangeEnd.Nanosecond() == 0 {
			toSerialize["timeRangeEnd"] = o.TimeRangeEnd.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timeRangeEnd"] = o.TimeRangeEnd.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InspectionTaskCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Creator        *string    `json:"creator,omitempty"`
		Engine         *string    `json:"engine,omitempty"`
		OrgName        *string    `json:"orgName,omitempty"`
		ClusterId      *string    `json:"clusterID,omitempty"`
		ClusterName    *string    `json:"clusterName,omitempty"`
		EnvName        *string    `json:"envName,omitempty"`
		EnvId          *string    `json:"envID,omitempty"`
		NodeName       *string    `json:"nodeName,omitempty"`
		IsAuto         *bool      `json:"isAuto,omitempty"`
		SavedDays      *int32     `json:"savedDays,omitempty"`
		TimeRangeStart *time.Time `json:"timeRangeStart,omitempty"`
		TimeRangeEnd   *time.Time `json:"timeRangeEnd,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"creator", "engine", "orgName", "clusterID", "clusterName", "envName", "envID", "nodeName", "isAuto", "savedDays", "timeRangeStart", "timeRangeEnd"})
	} else {
		return err
	}
	o.Creator = all.Creator
	o.Engine = all.Engine
	o.OrgName = all.OrgName
	o.ClusterId = all.ClusterId
	o.ClusterName = all.ClusterName
	o.EnvName = all.EnvName
	o.EnvId = all.EnvId
	o.NodeName = all.NodeName
	o.IsAuto = all.IsAuto
	o.SavedDays = all.SavedDays
	o.TimeRangeStart = all.TimeRangeStart
	o.TimeRangeEnd = all.TimeRangeEnd

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
