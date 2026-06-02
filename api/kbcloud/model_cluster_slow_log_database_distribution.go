// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogDatabaseDistribution Slow log count for a database
type ClusterSlowLogDatabaseDistribution struct {
	DbName *string `json:"dbName,omitempty"`
	Count  *int64  `json:"count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogDatabaseDistribution instantiates a new ClusterSlowLogDatabaseDistribution object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogDatabaseDistribution() *ClusterSlowLogDatabaseDistribution {
	this := ClusterSlowLogDatabaseDistribution{}
	return &this
}

// NewClusterSlowLogDatabaseDistributionWithDefaults instantiates a new ClusterSlowLogDatabaseDistribution object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogDatabaseDistributionWithDefaults() *ClusterSlowLogDatabaseDistribution {
	this := ClusterSlowLogDatabaseDistribution{}
	return &this
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *ClusterSlowLogDatabaseDistribution) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDatabaseDistribution) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *ClusterSlowLogDatabaseDistribution) HasDbName() bool {
	return o != nil && o.DbName != nil
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *ClusterSlowLogDatabaseDistribution) SetDbName(v string) {
	o.DbName = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ClusterSlowLogDatabaseDistribution) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDatabaseDistribution) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ClusterSlowLogDatabaseDistribution) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ClusterSlowLogDatabaseDistribution) SetCount(v int64) {
	o.Count = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogDatabaseDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DbName != nil {
		toSerialize["dbName"] = o.DbName
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogDatabaseDistribution) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DbName *string `json:"dbName,omitempty"`
		Count  *int64  `json:"count,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"dbName", "count"})
	} else {
		return err
	}
	o.DbName = all.DbName
	o.Count = all.Count

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
