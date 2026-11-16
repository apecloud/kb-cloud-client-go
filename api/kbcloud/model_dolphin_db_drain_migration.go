// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBDrainMigration struct {
	ChunkId string `json:"chunkId"`
	SrcNode string `json:"srcNode"`
	// Empty for a redundant replica deletion.
	DestNode string `json:"destNode"`
	// moveReplica or dropRedundantReplica. Preview does not execute either action.
	Action string `json:"action"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBDrainMigration instantiates a new DolphinDBDrainMigration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBDrainMigration(chunkId string, srcNode string, destNode string, action string) *DolphinDBDrainMigration {
	this := DolphinDBDrainMigration{}
	this.ChunkId = chunkId
	this.SrcNode = srcNode
	this.DestNode = destNode
	this.Action = action
	return &this
}

// NewDolphinDBDrainMigrationWithDefaults instantiates a new DolphinDBDrainMigration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBDrainMigrationWithDefaults() *DolphinDBDrainMigration {
	this := DolphinDBDrainMigration{}
	return &this
}

// GetChunkId returns the ChunkId field value.
func (o *DolphinDBDrainMigration) GetChunkId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChunkId
}

// GetChunkIdOk returns a tuple with the ChunkId field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainMigration) GetChunkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChunkId, true
}

// SetChunkId sets field value.
func (o *DolphinDBDrainMigration) SetChunkId(v string) {
	o.ChunkId = v
}

// GetSrcNode returns the SrcNode field value.
func (o *DolphinDBDrainMigration) GetSrcNode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SrcNode
}

// GetSrcNodeOk returns a tuple with the SrcNode field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainMigration) GetSrcNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SrcNode, true
}

// SetSrcNode sets field value.
func (o *DolphinDBDrainMigration) SetSrcNode(v string) {
	o.SrcNode = v
}

// GetDestNode returns the DestNode field value.
func (o *DolphinDBDrainMigration) GetDestNode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DestNode
}

// GetDestNodeOk returns a tuple with the DestNode field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainMigration) GetDestNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestNode, true
}

// SetDestNode sets field value.
func (o *DolphinDBDrainMigration) SetDestNode(v string) {
	o.DestNode = v
}

// GetAction returns the Action field value.
func (o *DolphinDBDrainMigration) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DolphinDBDrainMigration) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *DolphinDBDrainMigration) SetAction(v string) {
	o.Action = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBDrainMigration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["chunkId"] = o.ChunkId
	toSerialize["srcNode"] = o.SrcNode
	toSerialize["destNode"] = o.DestNode
	toSerialize["action"] = o.Action

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBDrainMigration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChunkId  *string `json:"chunkId"`
		SrcNode  *string `json:"srcNode"`
		DestNode *string `json:"destNode"`
		Action   *string `json:"action"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ChunkId == nil {
		return fmt.Errorf("required field chunkId missing")
	}
	if all.SrcNode == nil {
		return fmt.Errorf("required field srcNode missing")
	}
	if all.DestNode == nil {
		return fmt.Errorf("required field destNode missing")
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"chunkId", "srcNode", "destNode", "action"})
	} else {
		return err
	}
	o.ChunkId = *all.ChunkId
	o.SrcNode = *all.SrcNode
	o.DestNode = *all.DestNode
	o.Action = *all.Action

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
