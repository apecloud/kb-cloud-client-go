// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "time"

// Fault response of the fault
type Fault struct {
	// id of the fault
	Id *string `json:"id,omitempty"`
	// name of the fault
	Name *string `json:"name,omitempty"`
	// namespace of the fault
	Namespace *string `json:"namespace,omitempty"`
	// type of the fault
	Type *string `json:"type,omitempty"`
	// action of the fault
	Action *string `json:"action,omitempty"`
	// status of the fault
	Status *string `json:"status,omitempty"`
	// duration of the fault
	Duration *string `json:"duration,omitempty"`
	// environment id of the fault
	EnvironmentId *string `json:"environmentID,omitempty"`
	// environment name of the fault
	EnvironmentName *string `json:"environmentName,omitempty"`
	// org name of the fault
	OrgName *string `json:"orgName,omitempty"`
	// cluster id of the fault
	ClusterId *string `json:"clusterID,omitempty"`
	// cluster name of the fault
	ClusterName *string `json:"clusterName,omitempty"`
	// create time of the fault
	CreateTime *time.Time `json:"createTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFault instantiates a new Fault object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFault() *Fault {
	this := Fault{}
	return &this
}

// NewFaultWithDefaults instantiates a new Fault object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFaultWithDefaults() *Fault {
	this := Fault{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Fault) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Fault) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Fault) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Fault) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Fault) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Fault) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Fault) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Fault) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Fault) SetNamespace(v string) {
	o.Namespace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Fault) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Fault) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Fault) SetType(v string) {
	o.Type = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Fault) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Fault) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Fault) SetAction(v string) {
	o.Action = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Fault) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Fault) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Fault) SetStatus(v string) {
	o.Status = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Fault) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Fault) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *Fault) SetDuration(v string) {
	o.Duration = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *Fault) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *Fault) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *Fault) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *Fault) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *Fault) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *Fault) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Fault) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Fault) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Fault) SetOrgName(v string) {
	o.OrgName = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *Fault) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *Fault) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *Fault) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *Fault) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *Fault) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *Fault) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Fault) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Fault) HasCreateTime() bool {
	return o != nil && o.CreateTime != nil
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Fault) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Fault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentID"] = o.EnvironmentId
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
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
	if o.CreateTime != nil {
		if o.CreateTime.Nanosecond() == 0 {
			toSerialize["createTime"] = o.CreateTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createTime"] = o.CreateTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Fault) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id              *string    `json:"id,omitempty"`
		Name            *string    `json:"name,omitempty"`
		Namespace       *string    `json:"namespace,omitempty"`
		Type            *string    `json:"type,omitempty"`
		Action          *string    `json:"action,omitempty"`
		Status          *string    `json:"status,omitempty"`
		Duration        *string    `json:"duration,omitempty"`
		EnvironmentId   *string    `json:"environmentID,omitempty"`
		EnvironmentName *string    `json:"environmentName,omitempty"`
		OrgName         *string    `json:"orgName,omitempty"`
		ClusterId       *string    `json:"clusterID,omitempty"`
		ClusterName     *string    `json:"clusterName,omitempty"`
		CreateTime      *time.Time `json:"createTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "namespace", "type", "action", "status", "duration", "environmentID", "environmentName", "orgName", "clusterID", "clusterName", "createTime"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.Namespace = all.Namespace
	o.Type = all.Type
	o.Action = all.Action
	o.Status = all.Status
	o.Duration = all.Duration
	o.EnvironmentId = all.EnvironmentId
	o.EnvironmentName = all.EnvironmentName
	o.OrgName = all.OrgName
	o.ClusterId = all.ClusterId
	o.ClusterName = all.ClusterName
	o.CreateTime = all.CreateTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
