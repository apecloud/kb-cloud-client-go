// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MetricsConfig Configuration for metrics-related components (e.g., Victoriametrics)
type MetricsConfig struct {
	Cluster StaticCluster `json:"cluster"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricsConfig instantiates a new MetricsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricsConfig(cluster StaticCluster) *MetricsConfig {
	this := MetricsConfig{}
	this.Cluster = cluster
	return &this
}

// NewMetricsConfigWithDefaults instantiates a new MetricsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricsConfigWithDefaults() *MetricsConfig {
	this := MetricsConfig{}
	return &this
}

// GetCluster returns the Cluster field value.
func (o *MetricsConfig) GetCluster() StaticCluster {
	if o == nil {
		var ret StaticCluster
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *MetricsConfig) GetClusterOk() (*StaticCluster, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *MetricsConfig) SetCluster(v StaticCluster) {
	o.Cluster = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cluster"] = o.Cluster

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster *StaticCluster `json:"cluster"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = *all.Cluster

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
