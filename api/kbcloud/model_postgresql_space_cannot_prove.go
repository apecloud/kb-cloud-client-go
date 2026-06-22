// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSpaceCannotProve struct {
	Growth        bool     `json:"growth"`
	Pvc           bool     `json:"pvc"`
	ObjectHistory bool     `json:"objectHistory"`
	Reasons       []string `json:"reasons"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSpaceCannotProve instantiates a new PostgresqlSpaceCannotProve object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSpaceCannotProve(growth bool, pvc bool, objectHistory bool, reasons []string) *PostgresqlSpaceCannotProve {
	this := PostgresqlSpaceCannotProve{}
	this.Growth = growth
	this.Pvc = pvc
	this.ObjectHistory = objectHistory
	this.Reasons = reasons
	return &this
}

// NewPostgresqlSpaceCannotProveWithDefaults instantiates a new PostgresqlSpaceCannotProve object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSpaceCannotProveWithDefaults() *PostgresqlSpaceCannotProve {
	this := PostgresqlSpaceCannotProve{}
	return &this
}

// GetGrowth returns the Growth field value.
func (o *PostgresqlSpaceCannotProve) GetGrowth() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Growth
}

// GetGrowthOk returns a tuple with the Growth field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceCannotProve) GetGrowthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Growth, true
}

// SetGrowth sets field value.
func (o *PostgresqlSpaceCannotProve) SetGrowth(v bool) {
	o.Growth = v
}

// GetPvc returns the Pvc field value.
func (o *PostgresqlSpaceCannotProve) GetPvc() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Pvc
}

// GetPvcOk returns a tuple with the Pvc field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceCannotProve) GetPvcOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pvc, true
}

// SetPvc sets field value.
func (o *PostgresqlSpaceCannotProve) SetPvc(v bool) {
	o.Pvc = v
}

// GetObjectHistory returns the ObjectHistory field value.
func (o *PostgresqlSpaceCannotProve) GetObjectHistory() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ObjectHistory
}

// GetObjectHistoryOk returns a tuple with the ObjectHistory field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceCannotProve) GetObjectHistoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectHistory, true
}

// SetObjectHistory sets field value.
func (o *PostgresqlSpaceCannotProve) SetObjectHistory(v bool) {
	o.ObjectHistory = v
}

// GetReasons returns the Reasons field value.
func (o *PostgresqlSpaceCannotProve) GetReasons() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSpaceCannotProve) GetReasonsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reasons, true
}

// SetReasons sets field value.
func (o *PostgresqlSpaceCannotProve) SetReasons(v []string) {
	o.Reasons = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSpaceCannotProve) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["growth"] = o.Growth
	toSerialize["pvc"] = o.Pvc
	toSerialize["objectHistory"] = o.ObjectHistory
	toSerialize["reasons"] = o.Reasons

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSpaceCannotProve) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Growth        *bool     `json:"growth"`
		Pvc           *bool     `json:"pvc"`
		ObjectHistory *bool     `json:"objectHistory"`
		Reasons       *[]string `json:"reasons"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Growth == nil {
		return fmt.Errorf("required field growth missing")
	}
	if all.Pvc == nil {
		return fmt.Errorf("required field pvc missing")
	}
	if all.ObjectHistory == nil {
		return fmt.Errorf("required field objectHistory missing")
	}
	if all.Reasons == nil {
		return fmt.Errorf("required field reasons missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"growth", "pvc", "objectHistory", "reasons"})
	} else {
		return err
	}
	o.Growth = *all.Growth
	o.Pvc = *all.Pvc
	o.ObjectHistory = *all.ObjectHistory
	o.Reasons = *all.Reasons

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
