// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlDeadlockEvidence struct {
	Detected  bool    `json:"detected"`
	CyclePids []int64 `json:"cyclePids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlDeadlockEvidence instantiates a new PostgresqlDeadlockEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlDeadlockEvidence(detected bool, cyclePids []int64) *PostgresqlDeadlockEvidence {
	this := PostgresqlDeadlockEvidence{}
	this.Detected = detected
	this.CyclePids = cyclePids
	return &this
}

// NewPostgresqlDeadlockEvidenceWithDefaults instantiates a new PostgresqlDeadlockEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlDeadlockEvidenceWithDefaults() *PostgresqlDeadlockEvidence {
	this := PostgresqlDeadlockEvidence{}
	return &this
}

// GetDetected returns the Detected field value.
func (o *PostgresqlDeadlockEvidence) GetDetected() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Detected
}

// GetDetectedOk returns a tuple with the Detected field value
// and a boolean to check if the value has been set.
func (o *PostgresqlDeadlockEvidence) GetDetectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detected, true
}

// SetDetected sets field value.
func (o *PostgresqlDeadlockEvidence) SetDetected(v bool) {
	o.Detected = v
}

// GetCyclePids returns the CyclePids field value.
func (o *PostgresqlDeadlockEvidence) GetCyclePids() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.CyclePids
}

// GetCyclePidsOk returns a tuple with the CyclePids field value
// and a boolean to check if the value has been set.
func (o *PostgresqlDeadlockEvidence) GetCyclePidsOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CyclePids, true
}

// SetCyclePids sets field value.
func (o *PostgresqlDeadlockEvidence) SetCyclePids(v []int64) {
	o.CyclePids = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlDeadlockEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["detected"] = o.Detected
	toSerialize["cyclePids"] = o.CyclePids

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlDeadlockEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Detected  *bool    `json:"detected"`
		CyclePids *[]int64 `json:"cyclePids"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Detected == nil {
		return fmt.Errorf("required field detected missing")
	}
	if all.CyclePids == nil {
		return fmt.Errorf("required field cyclePids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"detected", "cyclePids"})
	} else {
		return err
	}
	o.Detected = *all.Detected
	o.CyclePids = *all.CyclePids

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
