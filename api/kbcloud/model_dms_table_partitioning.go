// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsTablePartitioning struct {
	// The partitioning statement for the table
	Statement *string `json:"statement,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsTablePartitioning instantiates a new DmsTablePartitioning object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsTablePartitioning() *DmsTablePartitioning {
	this := DmsTablePartitioning{}
	return &this
}

// NewDmsTablePartitioningWithDefaults instantiates a new DmsTablePartitioning object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsTablePartitioningWithDefaults() *DmsTablePartitioning {
	this := DmsTablePartitioning{}
	return &this
}

// GetStatement returns the Statement field value if set, zero value otherwise.
func (o *DmsTablePartitioning) GetStatement() string {
	if o == nil || o.Statement == nil {
		var ret string
		return ret
	}
	return *o.Statement
}

// GetStatementOk returns a tuple with the Statement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTablePartitioning) GetStatementOk() (*string, bool) {
	if o == nil || o.Statement == nil {
		return nil, false
	}
	return o.Statement, true
}

// HasStatement returns a boolean if a field has been set.
func (o *DmsTablePartitioning) HasStatement() bool {
	return o != nil && o.Statement != nil
}

// SetStatement gets a reference to the given string and assigns it to the Statement field.
func (o *DmsTablePartitioning) SetStatement(v string) {
	o.Statement = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsTablePartitioning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Statement != nil {
		toSerialize["statement"] = o.Statement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsTablePartitioning) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Statement *string `json:"statement,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"statement"})
	} else {
		return err
	}
	o.Statement = all.Statement

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
