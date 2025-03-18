// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Instance Instance is the information of KubeBlocks cluster instances
type Instance struct {
	// Access mode for instance
	AccessMode string `json:"accessMode"`
	// Org name
	OrgName *string `json:"orgName,omitempty"`
	// Cloud for instance
	Cloud string `json:"cloud"`
	// Cluster name
	Cluster string `json:"cluster"`
	// Component name
	ComponentName *string `json:"componentName,omitempty"`
	// ComponentDefinition name
	ComponentDef *string `json:"componentDef,omitempty"`
	// component type, refer to componentDef and support NamePrefix
	Component string `json:"component"`
	// cpu with uint cores.
	Cpu string `json:"cpu"`
	// created at
	CreatedAt string `json:"createdAt"`
	// Memory with uint Gi.
	Memory string `json:"memory"`
	// Instance name
	Name string `json:"name"`
	// node name
	Node string `json:"node"`
	// Region for instance
	Region string `json:"region"`
	// Role for instance
	Role string `json:"role"`
	// Status for instance
	Status InstanceStatus `json:"status"`
	// storage sets the storage size value mapping key
	Storage []InstanceStorageItem `json:"storage,omitempty"`
	// Available zone for instance
	Zone string `json:"zone"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstance instantiates a new Instance object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstance(accessMode string, cloud string, cluster string, component string, cpu string, createdAt string, memory string, name string, node string, region string, role string, status InstanceStatus, zone string) *Instance {
	this := Instance{}
	this.AccessMode = accessMode
	this.Cloud = cloud
	this.Cluster = cluster
	this.Component = component
	this.Cpu = cpu
	this.CreatedAt = createdAt
	this.Memory = memory
	this.Name = name
	this.Node = node
	this.Region = region
	this.Role = role
	this.Status = status
	this.Zone = zone
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	return &this
}

// GetAccessMode returns the AccessMode field value.
func (o *Instance) GetAccessMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value
// and a boolean to check if the value has been set.
func (o *Instance) GetAccessModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessMode, true
}

// SetAccessMode sets field value.
func (o *Instance) SetAccessMode(v string) {
	o.AccessMode = v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Instance) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Instance) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Instance) SetOrgName(v string) {
	o.OrgName = &v
}

// GetCloud returns the Cloud field value.
func (o *Instance) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value.
func (o *Instance) SetCloud(v string) {
	o.Cloud = v
}

// GetCluster returns the Cluster field value.
func (o *Instance) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Instance) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *Instance) SetCluster(v string) {
	o.Cluster = v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *Instance) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *Instance) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *Instance) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetComponentDef returns the ComponentDef field value if set, zero value otherwise.
func (o *Instance) GetComponentDef() string {
	if o == nil || o.ComponentDef == nil {
		var ret string
		return ret
	}
	return *o.ComponentDef
}

// GetComponentDefOk returns a tuple with the ComponentDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetComponentDefOk() (*string, bool) {
	if o == nil || o.ComponentDef == nil {
		return nil, false
	}
	return o.ComponentDef, true
}

// HasComponentDef returns a boolean if a field has been set.
func (o *Instance) HasComponentDef() bool {
	return o != nil && o.ComponentDef != nil
}

// SetComponentDef gets a reference to the given string and assigns it to the ComponentDef field.
func (o *Instance) SetComponentDef(v string) {
	o.ComponentDef = &v
}

// GetComponent returns the Component field value.
func (o *Instance) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *Instance) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *Instance) SetComponent(v string) {
	o.Component = v
}

// GetCpu returns the Cpu field value.
func (o *Instance) GetCpu() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value.
func (o *Instance) SetCpu(v string) {
	o.Cpu = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Instance) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Instance) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetMemory returns the Memory field value.
func (o *Instance) GetMemory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Instance) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value.
func (o *Instance) SetMemory(v string) {
	o.Memory = v
}

// GetName returns the Name field value.
func (o *Instance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Instance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Instance) SetName(v string) {
	o.Name = v
}

// GetNode returns the Node field value.
func (o *Instance) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *Instance) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value.
func (o *Instance) SetNode(v string) {
	o.Node = v
}

// GetRegion returns the Region field value.
func (o *Instance) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Instance) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *Instance) SetRegion(v string) {
	o.Region = v
}

// GetRole returns the Role field value.
func (o *Instance) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Instance) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *Instance) SetRole(v string) {
	o.Role = v
}

// GetStatus returns the Status field value.
func (o *Instance) GetStatus() InstanceStatus {
	if o == nil {
		var ret InstanceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Instance) GetStatusOk() (*InstanceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *Instance) SetStatus(v InstanceStatus) {
	o.Status = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *Instance) GetStorage() []InstanceStorageItem {
	if o == nil || o.Storage == nil {
		var ret []InstanceStorageItem
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetStorageOk() (*[]InstanceStorageItem, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return &o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *Instance) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given []InstanceStorageItem and assigns it to the Storage field.
func (o *Instance) SetStorage(v []InstanceStorageItem) {
	o.Storage = v
}

// GetZone returns the Zone field value.
func (o *Instance) GetZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *Instance) GetZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value.
func (o *Instance) SetZone(v string) {
	o.Zone = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Instance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["accessMode"] = o.AccessMode
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	toSerialize["cloud"] = o.Cloud
	toSerialize["cluster"] = o.Cluster
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	if o.ComponentDef != nil {
		toSerialize["componentDef"] = o.ComponentDef
	}
	toSerialize["component"] = o.Component
	toSerialize["cpu"] = o.Cpu
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["memory"] = o.Memory
	toSerialize["name"] = o.Name
	toSerialize["node"] = o.Node
	toSerialize["region"] = o.Region
	toSerialize["role"] = o.Role
	toSerialize["status"] = o.Status
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	toSerialize["zone"] = o.Zone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Instance) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessMode    *string               `json:"accessMode"`
		OrgName       *string               `json:"orgName,omitempty"`
		Cloud         *string               `json:"cloud"`
		Cluster       *string               `json:"cluster"`
		ComponentName *string               `json:"componentName,omitempty"`
		ComponentDef  *string               `json:"componentDef,omitempty"`
		Component     *string               `json:"component"`
		Cpu           *string               `json:"cpu"`
		CreatedAt     *string               `json:"createdAt"`
		Memory        *string               `json:"memory"`
		Name          *string               `json:"name"`
		Node          *string               `json:"node"`
		Region        *string               `json:"region"`
		Role          *string               `json:"role"`
		Status        *InstanceStatus       `json:"status"`
		Storage       []InstanceStorageItem `json:"storage,omitempty"`
		Zone          *string               `json:"zone"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.AccessMode == nil {
		return fmt.Errorf("required field accessMode missing")
	}
	if all.Cloud == nil {
		return fmt.Errorf("required field cloud missing")
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Cpu == nil {
		return fmt.Errorf("required field cpu missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.Memory == nil {
		return fmt.Errorf("required field memory missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Node == nil {
		return fmt.Errorf("required field node missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Zone == nil {
		return fmt.Errorf("required field zone missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"accessMode", "orgName", "cloud", "cluster", "componentName", "componentDef", "component", "cpu", "createdAt", "memory", "name", "node", "region", "role", "status", "storage", "zone"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccessMode = *all.AccessMode
	o.OrgName = all.OrgName
	o.Cloud = *all.Cloud
	o.Cluster = *all.Cluster
	o.ComponentName = all.ComponentName
	o.ComponentDef = all.ComponentDef
	o.Component = *all.Component
	o.Cpu = *all.Cpu
	o.CreatedAt = *all.CreatedAt
	o.Memory = *all.Memory
	o.Name = *all.Name
	o.Node = *all.Node
	o.Region = *all.Region
	o.Role = *all.Role
	if all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = *all.Status
	o.Storage = all.Storage
	o.Zone = *all.Zone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
