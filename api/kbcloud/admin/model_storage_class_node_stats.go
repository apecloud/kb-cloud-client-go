// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageClassNodeStats storageClassNodeStats is a storage class node stats.
type StorageClassNodeStats struct {
	// the name of the node
	NodeName string `json:"nodeName"`
	// the status of the node
	NodeStatus string `json:"nodeStatus"`
	// the total requests size of all PVCs on the node
	Requests float64 `json:"requests"`
	// the actual disk usage of hostpath on the node
	Usage float64 `json:"usage"`
	// the capacity of hostpath on the node
	Capacity float64 `json:"capacity"`
	// the number of PVCs on the node
	PvcCount int32 `json:"pvcCount"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageClassNodeStats instantiates a new StorageClassNodeStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageClassNodeStats(nodeName string, nodeStatus string, requests float64, usage float64, capacity float64, pvcCount int32) *StorageClassNodeStats {
	this := StorageClassNodeStats{}
	this.NodeName = nodeName
	this.NodeStatus = nodeStatus
	this.Requests = requests
	this.Usage = usage
	this.Capacity = capacity
	this.PvcCount = pvcCount
	return &this
}

// NewStorageClassNodeStatsWithDefaults instantiates a new StorageClassNodeStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageClassNodeStatsWithDefaults() *StorageClassNodeStats {
	this := StorageClassNodeStats{}
	return &this
}

// GetNodeName returns the NodeName field value.
func (o *StorageClassNodeStats) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value.
func (o *StorageClassNodeStats) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeStatus returns the NodeStatus field value.
func (o *StorageClassNodeStats) GetNodeStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NodeStatus
}

// GetNodeStatusOk returns a tuple with the NodeStatus field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetNodeStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeStatus, true
}

// SetNodeStatus sets field value.
func (o *StorageClassNodeStats) SetNodeStatus(v string) {
	o.NodeStatus = v
}

// GetRequests returns the Requests field value.
func (o *StorageClassNodeStats) GetRequests() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetRequestsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *StorageClassNodeStats) SetRequests(v float64) {
	o.Requests = v
}

// GetUsage returns the Usage field value.
func (o *StorageClassNodeStats) GetUsage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value.
func (o *StorageClassNodeStats) SetUsage(v float64) {
	o.Usage = v
}

// GetCapacity returns the Capacity field value.
func (o *StorageClassNodeStats) GetCapacity() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetCapacityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value.
func (o *StorageClassNodeStats) SetCapacity(v float64) {
	o.Capacity = v
}

// GetPvcCount returns the PvcCount field value.
func (o *StorageClassNodeStats) GetPvcCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.PvcCount
}

// GetPvcCountOk returns a tuple with the PvcCount field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetPvcCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PvcCount, true
}

// SetPvcCount sets field value.
func (o *StorageClassNodeStats) SetPvcCount(v int32) {
	o.PvcCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageClassNodeStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["nodeName"] = o.NodeName
	toSerialize["nodeStatus"] = o.NodeStatus
	toSerialize["requests"] = o.Requests
	toSerialize["usage"] = o.Usage
	toSerialize["capacity"] = o.Capacity
	toSerialize["pvcCount"] = o.PvcCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageClassNodeStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeName   *string  `json:"nodeName"`
		NodeStatus *string  `json:"nodeStatus"`
		Requests   *float64 `json:"requests"`
		Usage      *float64 `json:"usage"`
		Capacity   *float64 `json:"capacity"`
		PvcCount   *int32   `json:"pvcCount"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.NodeName == nil {
		return fmt.Errorf("required field nodeName missing")
	}
	if all.NodeStatus == nil {
		return fmt.Errorf("required field nodeStatus missing")
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Usage == nil {
		return fmt.Errorf("required field usage missing")
	}
	if all.Capacity == nil {
		return fmt.Errorf("required field capacity missing")
	}
	if all.PvcCount == nil {
		return fmt.Errorf("required field pvcCount missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeName", "nodeStatus", "requests", "usage", "capacity", "pvcCount"})
	} else {
		return err
	}
	o.NodeName = *all.NodeName
	o.NodeStatus = *all.NodeStatus
	o.Requests = *all.Requests
	o.Usage = *all.Usage
	o.Capacity = *all.Capacity
	o.PvcCount = *all.PvcCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
