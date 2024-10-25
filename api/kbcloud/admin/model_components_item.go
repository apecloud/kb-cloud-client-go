// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION ComponentsItem
type ComponentsItem struct {
	// component name
	Name *string `json:"name,omitempty"`
	// component type, refer to componentDef and support NamePrefix
	Component *string `json:"component,omitempty"`
	// componentDefinition name
	ComponentDef *string `json:"componentDef,omitempty"`
	// number of components
	CompNum *int32 `json:"compNum,omitempty"`
	// enable monitor
	Monitor *bool `json:"monitor,omitempty"`
	// The number of replicas, for standalone mode, the replicas is 1, for raftGroup mode, the default replicas is 3.
	Replicas *int32 `json:"replicas,omitempty"`
	// NODESCRIPTION ClassCode
	ClassCode *string `json:"classCode,omitempty"`
	// NODESCRIPTION ClassSeries
	ClassSeries *string `json:"classSeries,omitempty"`
	// CPU cores.
	Cpu *float64 `json:"cpu,omitempty"`
	// Memory, the unit is Gi.
	Memory *float64 `json:"memory,omitempty"`
	// Storage
	Storage *string `json:"storage,omitempty"`
	// StorageClass name
	StorageClass *string `json:"storageClass,omitempty"`
	// NODESCRIPTION Volumes
	Volumes []ComponentsItemVolumesItem `json:"volumes,omitempty"`
	// Cluster main component codeShort
	CodeShort *string `json:"codeShort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentsItem instantiates a new ComponentsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentsItem() *ComponentsItem {
	this := ComponentsItem{}
	return &this
}

// NewComponentsItemWithDefaults instantiates a new ComponentsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentsItemWithDefaults() *ComponentsItem {
	this := ComponentsItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentsItem) SetName(v string) {
	o.Name = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ComponentsItem) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ComponentsItem) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ComponentsItem) SetComponent(v string) {
	o.Component = &v
}

// GetComponentDef returns the ComponentDef field value if set, zero value otherwise.
func (o *ComponentsItem) GetComponentDef() string {
	if o == nil || o.ComponentDef == nil {
		var ret string
		return ret
	}
	return *o.ComponentDef
}

// GetComponentDefOk returns a tuple with the ComponentDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetComponentDefOk() (*string, bool) {
	if o == nil || o.ComponentDef == nil {
		return nil, false
	}
	return o.ComponentDef, true
}

// HasComponentDef returns a boolean if a field has been set.
func (o *ComponentsItem) HasComponentDef() bool {
	return o != nil && o.ComponentDef != nil
}

// SetComponentDef gets a reference to the given string and assigns it to the ComponentDef field.
func (o *ComponentsItem) SetComponentDef(v string) {
	o.ComponentDef = &v
}

// GetCompNum returns the CompNum field value if set, zero value otherwise.
func (o *ComponentsItem) GetCompNum() int32 {
	if o == nil || o.CompNum == nil {
		var ret int32
		return ret
	}
	return *o.CompNum
}

// GetCompNumOk returns a tuple with the CompNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetCompNumOk() (*int32, bool) {
	if o == nil || o.CompNum == nil {
		return nil, false
	}
	return o.CompNum, true
}

// HasCompNum returns a boolean if a field has been set.
func (o *ComponentsItem) HasCompNum() bool {
	return o != nil && o.CompNum != nil
}

// SetCompNum gets a reference to the given int32 and assigns it to the CompNum field.
func (o *ComponentsItem) SetCompNum(v int32) {
	o.CompNum = &v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *ComponentsItem) GetMonitor() bool {
	if o == nil || o.Monitor == nil {
		var ret bool
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetMonitorOk() (*bool, bool) {
	if o == nil || o.Monitor == nil {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *ComponentsItem) HasMonitor() bool {
	return o != nil && o.Monitor != nil
}

// SetMonitor gets a reference to the given bool and assigns it to the Monitor field.
func (o *ComponentsItem) SetMonitor(v bool) {
	o.Monitor = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ComponentsItem) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ComponentsItem) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *ComponentsItem) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *ComponentsItem) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *ComponentsItem) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *ComponentsItem) SetClassCode(v string) {
	o.ClassCode = &v
}

// GetClassSeries returns the ClassSeries field value if set, zero value otherwise.
func (o *ComponentsItem) GetClassSeries() string {
	if o == nil || o.ClassSeries == nil {
		var ret string
		return ret
	}
	return *o.ClassSeries
}

// GetClassSeriesOk returns a tuple with the ClassSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetClassSeriesOk() (*string, bool) {
	if o == nil || o.ClassSeries == nil {
		return nil, false
	}
	return o.ClassSeries, true
}

// HasClassSeries returns a boolean if a field has been set.
func (o *ComponentsItem) HasClassSeries() bool {
	return o != nil && o.ClassSeries != nil
}

// SetClassSeries gets a reference to the given string and assigns it to the ClassSeries field.
func (o *ComponentsItem) SetClassSeries(v string) {
	o.ClassSeries = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ComponentsItem) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ComponentsItem) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *ComponentsItem) SetCpu(v float64) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ComponentsItem) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ComponentsItem) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *ComponentsItem) SetMemory(v float64) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ComponentsItem) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ComponentsItem) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *ComponentsItem) SetStorage(v string) {
	o.Storage = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *ComponentsItem) GetStorageClass() string {
	if o == nil || o.StorageClass == nil {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetStorageClassOk() (*string, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *ComponentsItem) HasStorageClass() bool {
	return o != nil && o.StorageClass != nil
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *ComponentsItem) SetStorageClass(v string) {
	o.StorageClass = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ComponentsItem) GetVolumes() []ComponentsItemVolumesItem {
	if o == nil || o.Volumes == nil {
		var ret []ComponentsItemVolumesItem
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetVolumesOk() (*[]ComponentsItemVolumesItem, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ComponentsItem) HasVolumes() bool {
	return o != nil && o.Volumes != nil
}

// SetVolumes gets a reference to the given []ComponentsItemVolumesItem and assigns it to the Volumes field.
func (o *ComponentsItem) SetVolumes(v []ComponentsItemVolumesItem) {
	o.Volumes = v
}

// GetCodeShort returns the CodeShort field value if set, zero value otherwise.
func (o *ComponentsItem) GetCodeShort() string {
	if o == nil || o.CodeShort == nil {
		var ret string
		return ret
	}
	return *o.CodeShort
}

// GetCodeShortOk returns a tuple with the CodeShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentsItem) GetCodeShortOk() (*string, bool) {
	if o == nil || o.CodeShort == nil {
		return nil, false
	}
	return o.CodeShort, true
}

// HasCodeShort returns a boolean if a field has been set.
func (o *ComponentsItem) HasCodeShort() bool {
	return o != nil && o.CodeShort != nil
}

// SetCodeShort gets a reference to the given string and assigns it to the CodeShort field.
func (o *ComponentsItem) SetCodeShort(v string) {
	o.CodeShort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ComponentDef != nil {
		toSerialize["componentDef"] = o.ComponentDef
	}
	if o.CompNum != nil {
		toSerialize["compNum"] = o.CompNum
	}
	if o.Monitor != nil {
		toSerialize["monitor"] = o.Monitor
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}
	if o.ClassSeries != nil {
		toSerialize["classSeries"] = o.ClassSeries
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.StorageClass != nil {
		toSerialize["storageClass"] = o.StorageClass
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.CodeShort != nil {
		toSerialize["codeShort"] = o.CodeShort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string                     `json:"name,omitempty"`
		Component    *string                     `json:"component,omitempty"`
		ComponentDef *string                     `json:"componentDef,omitempty"`
		CompNum      *int32                      `json:"compNum,omitempty"`
		Monitor      *bool                       `json:"monitor,omitempty"`
		Replicas     *int32                      `json:"replicas,omitempty"`
		ClassCode    *string                     `json:"classCode,omitempty"`
		ClassSeries  *string                     `json:"classSeries,omitempty"`
		Cpu          *float64                    `json:"cpu,omitempty"`
		Memory       *float64                    `json:"memory,omitempty"`
		Storage      *string                     `json:"storage,omitempty"`
		StorageClass *string                     `json:"storageClass,omitempty"`
		Volumes      []ComponentsItemVolumesItem `json:"volumes,omitempty"`
		CodeShort    *string                     `json:"codeShort,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "component", "componentDef", "compNum", "monitor", "replicas", "classCode", "classSeries", "cpu", "memory", "storage", "storageClass", "volumes", "codeShort"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Component = all.Component
	o.ComponentDef = all.ComponentDef
	o.CompNum = all.CompNum
	o.Monitor = all.Monitor
	o.Replicas = all.Replicas
	o.ClassCode = all.ClassCode
	o.ClassSeries = all.ClassSeries
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.Storage = all.Storage
	o.StorageClass = all.StorageClass
	o.Volumes = all.Volumes
	o.CodeShort = all.CodeShort

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
