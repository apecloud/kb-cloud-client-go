// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlStorageOverview struct {
	Source string `json:"source"`
	// Physical PostgreSQL replica storage usage split by pod/PVC/role.
	Instances []PostgresqlStorageInstanceUsage `json:"instances,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlStorageOverview instantiates a new PostgresqlStorageOverview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlStorageOverview(source string) *PostgresqlStorageOverview {
	this := PostgresqlStorageOverview{}
	this.Source = source
	return &this
}

// NewPostgresqlStorageOverviewWithDefaults instantiates a new PostgresqlStorageOverview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlStorageOverviewWithDefaults() *PostgresqlStorageOverview {
	this := PostgresqlStorageOverview{}
	return &this
}

// GetSource returns the Source field value.
func (o *PostgresqlStorageOverview) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageOverview) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *PostgresqlStorageOverview) SetSource(v string) {
	o.Source = v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *PostgresqlStorageOverview) GetInstances() []PostgresqlStorageInstanceUsage {
	if o == nil || o.Instances == nil {
		var ret []PostgresqlStorageInstanceUsage
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageOverview) GetInstancesOk() (*[]PostgresqlStorageInstanceUsage, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return &o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *PostgresqlStorageOverview) HasInstances() bool {
	return o != nil && o.Instances != nil
}

// SetInstances gets a reference to the given []PostgresqlStorageInstanceUsage and assigns it to the Instances field.
func (o *PostgresqlStorageOverview) SetInstances(v []PostgresqlStorageInstanceUsage) {
	o.Instances = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlStorageOverview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlStorageOverview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source    *string                          `json:"source"`
		Instances []PostgresqlStorageInstanceUsage `json:"instances,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "instances"})
	} else {
		return err
	}
	o.Source = *all.Source
	o.Instances = all.Instances

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
