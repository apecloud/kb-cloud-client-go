// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MysqlStorageOverview struct {
	Source    string                      `json:"source"`
	Instances []MysqlStorageInstanceUsage `json:"instances"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMysqlStorageOverview instantiates a new MysqlStorageOverview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMysqlStorageOverview(source string, instances []MysqlStorageInstanceUsage) *MysqlStorageOverview {
	this := MysqlStorageOverview{}
	this.Source = source
	this.Instances = instances
	return &this
}

// NewMysqlStorageOverviewWithDefaults instantiates a new MysqlStorageOverview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMysqlStorageOverviewWithDefaults() *MysqlStorageOverview {
	this := MysqlStorageOverview{}
	return &this
}

// GetSource returns the Source field value.
func (o *MysqlStorageOverview) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *MysqlStorageOverview) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *MysqlStorageOverview) SetSource(v string) {
	o.Source = v
}

// GetInstances returns the Instances field value.
func (o *MysqlStorageOverview) GetInstances() []MysqlStorageInstanceUsage {
	if o == nil {
		var ret []MysqlStorageInstanceUsage
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *MysqlStorageOverview) GetInstancesOk() (*[]MysqlStorageInstanceUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instances, true
}

// SetInstances sets field value.
func (o *MysqlStorageOverview) SetInstances(v []MysqlStorageInstanceUsage) {
	o.Instances = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MysqlStorageOverview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	toSerialize["instances"] = o.Instances

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MysqlStorageOverview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source    *string                      `json:"source"`
		Instances *[]MysqlStorageInstanceUsage `json:"instances"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Instances == nil {
		return fmt.Errorf("required field instances missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "instances"})
	} else {
		return err
	}
	o.Source = *all.Source
	o.Instances = *all.Instances

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
