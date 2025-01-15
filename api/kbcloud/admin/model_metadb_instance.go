// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Metadb_instance Instance is the information of metadb cluster instances
type Metadb_instance struct {
	// Cluster name
	Cluster string `json:"cluster"`
	// cpu with uint cores.
	Cpu string `json:"cpu"`
	// Memory with uint Gi.
	Memory string `json:"memory"`
	// Instance name
	Name string `json:"name"`
	// node name
	Node string `json:"node"`
	// Role for instance
	Role string `json:"role"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetadb_instance instantiates a new Metadb_instance object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetadb_instance(cluster string, cpu string, memory string, name string, node string, role string) *Metadb_instance {
	this := Metadb_instance{}
	this.Cluster = cluster
	this.Cpu = cpu
	this.Memory = memory
	this.Name = name
	this.Node = node
	this.Role = role
	return &this
}

// NewMetadb_instanceWithDefaults instantiates a new Metadb_instance object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetadb_instanceWithDefaults() *Metadb_instance {
	this := Metadb_instance{}
	return &this
}

// GetCluster returns the Cluster field value.
func (o *Metadb_instance) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Metadb_instance) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *Metadb_instance) SetCluster(v string) {
	o.Cluster = v
}

// GetCpu returns the Cpu field value.
func (o *Metadb_instance) GetCpu() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Metadb_instance) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value.
func (o *Metadb_instance) SetCpu(v string) {
	o.Cpu = v
}

// GetMemory returns the Memory field value.
func (o *Metadb_instance) GetMemory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Metadb_instance) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value.
func (o *Metadb_instance) SetMemory(v string) {
	o.Memory = v
}

// GetName returns the Name field value.
func (o *Metadb_instance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metadb_instance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Metadb_instance) SetName(v string) {
	o.Name = v
}

// GetNode returns the Node field value.
func (o *Metadb_instance) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *Metadb_instance) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value.
func (o *Metadb_instance) SetNode(v string) {
	o.Node = v
}

// GetRole returns the Role field value.
func (o *Metadb_instance) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Metadb_instance) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *Metadb_instance) SetRole(v string) {
	o.Role = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Metadb_instance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cluster"] = o.Cluster
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory
	toSerialize["name"] = o.Name
	toSerialize["node"] = o.Node
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Metadb_instance) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster *string `json:"cluster"`
		Cpu     *string `json:"cpu"`
		Memory  *string `json:"memory"`
		Name    *string `json:"name"`
		Node    *string `json:"node"`
		Role    *string `json:"role"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	if all.Cpu == nil {
		return fmt.Errorf("required field cpu missing")
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
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster", "cpu", "memory", "name", "node", "role"})
	} else {
		return err
	}
	o.Cluster = *all.Cluster
	o.Cpu = *all.Cpu
	o.Memory = *all.Memory
	o.Name = *all.Name
	o.Node = *all.Node
	o.Role = *all.Role

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
