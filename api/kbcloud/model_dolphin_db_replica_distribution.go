// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DolphinDBReplicaDistribution struct {
	Node         *string `json:"node,omitempty"`
	ReplicaCount *int64  `json:"replicaCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBReplicaDistribution instantiates a new DolphinDBReplicaDistribution object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBReplicaDistribution() *DolphinDBReplicaDistribution {
	this := DolphinDBReplicaDistribution{}
	return &this
}

// NewDolphinDBReplicaDistributionWithDefaults instantiates a new DolphinDBReplicaDistribution object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBReplicaDistributionWithDefaults() *DolphinDBReplicaDistribution {
	this := DolphinDBReplicaDistribution{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *DolphinDBReplicaDistribution) GetNode() string {
	if o == nil || o.Node == nil {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBReplicaDistribution) GetNodeOk() (*string, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *DolphinDBReplicaDistribution) HasNode() bool {
	return o != nil && o.Node != nil
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *DolphinDBReplicaDistribution) SetNode(v string) {
	o.Node = &v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *DolphinDBReplicaDistribution) GetReplicaCount() int64 {
	if o == nil || o.ReplicaCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBReplicaDistribution) GetReplicaCountOk() (*int64, bool) {
	if o == nil || o.ReplicaCount == nil {
		return nil, false
	}
	return o.ReplicaCount, true
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *DolphinDBReplicaDistribution) HasReplicaCount() bool {
	return o != nil && o.ReplicaCount != nil
}

// SetReplicaCount gets a reference to the given int64 and assigns it to the ReplicaCount field.
func (o *DolphinDBReplicaDistribution) SetReplicaCount(v int64) {
	o.ReplicaCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBReplicaDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}
	if o.ReplicaCount != nil {
		toSerialize["replicaCount"] = o.ReplicaCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBReplicaDistribution) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Node         *string `json:"node,omitempty"`
		ReplicaCount *int64  `json:"replicaCount,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"node", "replicaCount"})
	} else {
		return err
	}
	o.Node = all.Node
	o.ReplicaCount = all.ReplicaCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
