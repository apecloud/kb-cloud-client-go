// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

type BenchOption struct {
	Pgbench  bool `json:"pgbench"`
	Sysbench bool `json:"sysbench"`
	Tpcc     bool `json:"tpcc"`
	Ycsb     bool `json:"ycsb"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBenchOption instantiates a new BenchOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBenchOption(pgbench bool, sysbench bool, tpcc bool, ycsb bool) *BenchOption {
	this := BenchOption{}
	this.Pgbench = pgbench
	this.Sysbench = sysbench
	this.Tpcc = tpcc
	this.Ycsb = ycsb
	return &this
}

// NewBenchOptionWithDefaults instantiates a new BenchOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBenchOptionWithDefaults() *BenchOption {
	this := BenchOption{}
	return &this
}

// GetPgbench returns the Pgbench field value.
func (o *BenchOption) GetPgbench() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Pgbench
}

// GetPgbenchOk returns a tuple with the Pgbench field value
// and a boolean to check if the value has been set.
func (o *BenchOption) GetPgbenchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pgbench, true
}

// SetPgbench sets field value.
func (o *BenchOption) SetPgbench(v bool) {
	o.Pgbench = v
}

// GetSysbench returns the Sysbench field value.
func (o *BenchOption) GetSysbench() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Sysbench
}

// GetSysbenchOk returns a tuple with the Sysbench field value
// and a boolean to check if the value has been set.
func (o *BenchOption) GetSysbenchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sysbench, true
}

// SetSysbench sets field value.
func (o *BenchOption) SetSysbench(v bool) {
	o.Sysbench = v
}

// GetTpcc returns the Tpcc field value.
func (o *BenchOption) GetTpcc() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Tpcc
}

// GetTpccOk returns a tuple with the Tpcc field value
// and a boolean to check if the value has been set.
func (o *BenchOption) GetTpccOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tpcc, true
}

// SetTpcc sets field value.
func (o *BenchOption) SetTpcc(v bool) {
	o.Tpcc = v
}

// GetYcsb returns the Ycsb field value.
func (o *BenchOption) GetYcsb() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Ycsb
}

// GetYcsbOk returns a tuple with the Ycsb field value
// and a boolean to check if the value has been set.
func (o *BenchOption) GetYcsbOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ycsb, true
}

// SetYcsb sets field value.
func (o *BenchOption) SetYcsb(v bool) {
	o.Ycsb = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BenchOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pgbench"] = o.Pgbench
	toSerialize["sysbench"] = o.Sysbench
	toSerialize["tpcc"] = o.Tpcc
	toSerialize["ycsb"] = o.Ycsb

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BenchOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pgbench  *bool `json:"pgbench"`
		Sysbench *bool `json:"sysbench"`
		Tpcc     *bool `json:"tpcc"`
		Ycsb     *bool `json:"ycsb"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Pgbench == nil {
		return fmt.Errorf("required field pgbench missing")
	}
	if all.Sysbench == nil {
		return fmt.Errorf("required field sysbench missing")
	}
	if all.Tpcc == nil {
		return fmt.Errorf("required field tpcc missing")
	}
	if all.Ycsb == nil {
		return fmt.Errorf("required field ycsb missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pgbench", "sysbench", "tpcc", "ycsb"})
	} else {
		return err
	}
	o.Pgbench = *all.Pgbench
	o.Sysbench = *all.Sysbench
	o.Tpcc = *all.Tpcc
	o.Ycsb = *all.Ycsb

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
