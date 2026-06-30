// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlWaitGraphEdgeEvidence struct {
	// Lock rows the backend can associate with this wait edge.
	LockRows []PostgresqlLockRow `json:"lockRows"`
	// Non-empty when the backend cannot associate precise lock rows with this edge.
	Limitation string `json:"limitation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlWaitGraphEdgeEvidence instantiates a new PostgresqlWaitGraphEdgeEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlWaitGraphEdgeEvidence(lockRows []PostgresqlLockRow, limitation string) *PostgresqlWaitGraphEdgeEvidence {
	this := PostgresqlWaitGraphEdgeEvidence{}
	this.LockRows = lockRows
	this.Limitation = limitation
	return &this
}

// NewPostgresqlWaitGraphEdgeEvidenceWithDefaults instantiates a new PostgresqlWaitGraphEdgeEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlWaitGraphEdgeEvidenceWithDefaults() *PostgresqlWaitGraphEdgeEvidence {
	this := PostgresqlWaitGraphEdgeEvidence{}
	return &this
}

// GetLockRows returns the LockRows field value.
func (o *PostgresqlWaitGraphEdgeEvidence) GetLockRows() []PostgresqlLockRow {
	if o == nil {
		var ret []PostgresqlLockRow
		return ret
	}
	return o.LockRows
}

// GetLockRowsOk returns a tuple with the LockRows field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitGraphEdgeEvidence) GetLockRowsOk() (*[]PostgresqlLockRow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockRows, true
}

// SetLockRows sets field value.
func (o *PostgresqlWaitGraphEdgeEvidence) SetLockRows(v []PostgresqlLockRow) {
	o.LockRows = v
}

// GetLimitation returns the Limitation field value.
func (o *PostgresqlWaitGraphEdgeEvidence) GetLimitation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Limitation
}

// GetLimitationOk returns a tuple with the Limitation field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitGraphEdgeEvidence) GetLimitationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limitation, true
}

// SetLimitation sets field value.
func (o *PostgresqlWaitGraphEdgeEvidence) SetLimitation(v string) {
	o.Limitation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlWaitGraphEdgeEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["lockRows"] = o.LockRows
	toSerialize["limitation"] = o.Limitation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlWaitGraphEdgeEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LockRows   *[]PostgresqlLockRow `json:"lockRows"`
		Limitation *string              `json:"limitation"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.LockRows == nil {
		return fmt.Errorf("required field lockRows missing")
	}
	if all.Limitation == nil {
		return fmt.Errorf("required field limitation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"lockRows", "limitation"})
	} else {
		return err
	}
	o.LockRows = *all.LockRows
	o.Limitation = *all.Limitation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
