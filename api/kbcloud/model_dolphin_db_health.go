// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBHealth struct {
	Database       DolphinDBHealthSnapshot   `json:"database"`
	Infrastructure DolphinDBHealthCheck      `json:"infrastructure"`
	Instances      []DolphinDBInstanceHealth `json:"instances"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBHealth instantiates a new DolphinDBHealth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBHealth(database DolphinDBHealthSnapshot, infrastructure DolphinDBHealthCheck, instances []DolphinDBInstanceHealth) *DolphinDBHealth {
	this := DolphinDBHealth{}
	this.Database = database
	this.Infrastructure = infrastructure
	this.Instances = instances
	return &this
}

// NewDolphinDBHealthWithDefaults instantiates a new DolphinDBHealth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBHealthWithDefaults() *DolphinDBHealth {
	this := DolphinDBHealth{}
	return &this
}

// GetDatabase returns the Database field value.
func (o *DolphinDBHealth) GetDatabase() DolphinDBHealthSnapshot {
	if o == nil {
		var ret DolphinDBHealthSnapshot
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *DolphinDBHealth) GetDatabaseOk() (*DolphinDBHealthSnapshot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *DolphinDBHealth) SetDatabase(v DolphinDBHealthSnapshot) {
	o.Database = v
}

// GetInfrastructure returns the Infrastructure field value.
func (o *DolphinDBHealth) GetInfrastructure() DolphinDBHealthCheck {
	if o == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return o.Infrastructure
}

// GetInfrastructureOk returns a tuple with the Infrastructure field value
// and a boolean to check if the value has been set.
func (o *DolphinDBHealth) GetInfrastructureOk() (*DolphinDBHealthCheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Infrastructure, true
}

// SetInfrastructure sets field value.
func (o *DolphinDBHealth) SetInfrastructure(v DolphinDBHealthCheck) {
	o.Infrastructure = v
}

// GetInstances returns the Instances field value.
func (o *DolphinDBHealth) GetInstances() []DolphinDBInstanceHealth {
	if o == nil {
		var ret []DolphinDBInstanceHealth
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *DolphinDBHealth) GetInstancesOk() (*[]DolphinDBInstanceHealth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instances, true
}

// SetInstances sets field value.
func (o *DolphinDBHealth) SetInstances(v []DolphinDBInstanceHealth) {
	o.Instances = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["database"] = o.Database
	toSerialize["infrastructure"] = o.Infrastructure
	toSerialize["instances"] = o.Instances

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBHealth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database       *DolphinDBHealthSnapshot   `json:"database"`
		Infrastructure *DolphinDBHealthCheck      `json:"infrastructure"`
		Instances      *[]DolphinDBInstanceHealth `json:"instances"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.Infrastructure == nil {
		return fmt.Errorf("required field infrastructure missing")
	}
	if all.Instances == nil {
		return fmt.Errorf("required field instances missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "infrastructure", "instances"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Database.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Database = *all.Database
	if all.Infrastructure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Infrastructure = *all.Infrastructure
	o.Instances = *all.Instances

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
