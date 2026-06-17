// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlWaitChainEdge struct {
	// PID that blocks another session.
	BlockingPid int64 `json:"blockingPid"`
	// PID that is blocked by another session.
	BlockedPid int64 `json:"blockedPid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlWaitChainEdge instantiates a new PostgresqlWaitChainEdge object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlWaitChainEdge(blockingPid int64, blockedPid int64) *PostgresqlWaitChainEdge {
	this := PostgresqlWaitChainEdge{}
	this.BlockingPid = blockingPid
	this.BlockedPid = blockedPid
	return &this
}

// NewPostgresqlWaitChainEdgeWithDefaults instantiates a new PostgresqlWaitChainEdge object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlWaitChainEdgeWithDefaults() *PostgresqlWaitChainEdge {
	this := PostgresqlWaitChainEdge{}
	return &this
}

// GetBlockingPid returns the BlockingPid field value.
func (o *PostgresqlWaitChainEdge) GetBlockingPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BlockingPid
}

// GetBlockingPidOk returns a tuple with the BlockingPid field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitChainEdge) GetBlockingPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockingPid, true
}

// SetBlockingPid sets field value.
func (o *PostgresqlWaitChainEdge) SetBlockingPid(v int64) {
	o.BlockingPid = v
}

// GetBlockedPid returns the BlockedPid field value.
func (o *PostgresqlWaitChainEdge) GetBlockedPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BlockedPid
}

// GetBlockedPidOk returns a tuple with the BlockedPid field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitChainEdge) GetBlockedPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockedPid, true
}

// SetBlockedPid sets field value.
func (o *PostgresqlWaitChainEdge) SetBlockedPid(v int64) {
	o.BlockedPid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlWaitChainEdge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["blockingPid"] = o.BlockingPid
	toSerialize["blockedPid"] = o.BlockedPid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlWaitChainEdge) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BlockingPid *int64 `json:"blockingPid"`
		BlockedPid  *int64 `json:"blockedPid"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.BlockingPid == nil {
		return fmt.Errorf("required field blockingPid missing")
	}
	if all.BlockedPid == nil {
		return fmt.Errorf("required field blockedPid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"blockingPid", "blockedPid"})
	} else {
		return err
	}
	o.BlockingPid = *all.BlockingPid
	o.BlockedPid = *all.BlockedPid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
