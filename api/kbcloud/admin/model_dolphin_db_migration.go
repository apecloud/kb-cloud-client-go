// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBMigration struct {
	ChunkId    string  `json:"chunkId"`
	SrcNode    string  `json:"srcNode"`
	DestNode   string  `json:"destNode"`
	DestVolume *string `json:"destVolume,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBMigration instantiates a new DolphinDBMigration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBMigration(chunkId string, srcNode string, destNode string) *DolphinDBMigration {
	this := DolphinDBMigration{}
	this.ChunkId = chunkId
	this.SrcNode = srcNode
	this.DestNode = destNode
	return &this
}

// NewDolphinDBMigrationWithDefaults instantiates a new DolphinDBMigration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBMigrationWithDefaults() *DolphinDBMigration {
	this := DolphinDBMigration{}
	return &this
}

// GetChunkId returns the ChunkId field value.
func (o *DolphinDBMigration) GetChunkId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChunkId
}

// GetChunkIdOk returns a tuple with the ChunkId field value
// and a boolean to check if the value has been set.
func (o *DolphinDBMigration) GetChunkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChunkId, true
}

// SetChunkId sets field value.
func (o *DolphinDBMigration) SetChunkId(v string) {
	o.ChunkId = v
}

// GetSrcNode returns the SrcNode field value.
func (o *DolphinDBMigration) GetSrcNode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SrcNode
}

// GetSrcNodeOk returns a tuple with the SrcNode field value
// and a boolean to check if the value has been set.
func (o *DolphinDBMigration) GetSrcNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SrcNode, true
}

// SetSrcNode sets field value.
func (o *DolphinDBMigration) SetSrcNode(v string) {
	o.SrcNode = v
}

// GetDestNode returns the DestNode field value.
func (o *DolphinDBMigration) GetDestNode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DestNode
}

// GetDestNodeOk returns a tuple with the DestNode field value
// and a boolean to check if the value has been set.
func (o *DolphinDBMigration) GetDestNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestNode, true
}

// SetDestNode sets field value.
func (o *DolphinDBMigration) SetDestNode(v string) {
	o.DestNode = v
}

// GetDestVolume returns the DestVolume field value if set, zero value otherwise.
func (o *DolphinDBMigration) GetDestVolume() string {
	if o == nil || o.DestVolume == nil {
		var ret string
		return ret
	}
	return *o.DestVolume
}

// GetDestVolumeOk returns a tuple with the DestVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBMigration) GetDestVolumeOk() (*string, bool) {
	if o == nil || o.DestVolume == nil {
		return nil, false
	}
	return o.DestVolume, true
}

// HasDestVolume returns a boolean if a field has been set.
func (o *DolphinDBMigration) HasDestVolume() bool {
	return o != nil && o.DestVolume != nil
}

// SetDestVolume gets a reference to the given string and assigns it to the DestVolume field.
func (o *DolphinDBMigration) SetDestVolume(v string) {
	o.DestVolume = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBMigration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["chunkId"] = o.ChunkId
	toSerialize["srcNode"] = o.SrcNode
	toSerialize["destNode"] = o.DestNode
	if o.DestVolume != nil {
		toSerialize["destVolume"] = o.DestVolume
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBMigration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChunkId    *string `json:"chunkId"`
		SrcNode    *string `json:"srcNode"`
		DestNode   *string `json:"destNode"`
		DestVolume *string `json:"destVolume,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"chunkId", "srcNode", "destNode", "destVolume"})
	} else {
		return err
	}
	o.ChunkId = *all.ChunkId
	o.SrcNode = *all.SrcNode
	o.DestNode = *all.DestNode
	o.DestVolume = all.DestVolume

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
