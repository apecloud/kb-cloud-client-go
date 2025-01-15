// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AutohealingConfig cluster instance autohealing process config
type AutohealingConfig struct {
	// Set to true to enable the auto-healing process, which automatically rebuilds instances when a node fails or is in maintenance mode.
	EnableAutoHealing bool `json:"enableAutoHealing"`
	// The number of seconds to wait before starting a rebuild when a node fails.
	RebuildStartDelaySecondsWhenNodeFail int32 `json:"rebuildStartDelaySecondsWhenNodeFail"`
	// The number of seconds to wait before starting a rebuild when a node is in maintenance mode.
	RebuildStartDelaySecondsWhenInMaintenance int32 `json:"rebuildStartDelaySecondsWhenInMaintenance"`
	// The minimum time in seconds between consecutive rebuild jobs for one cluster.
	MinClusterRebuildIntervalSeconds int32 `json:"minClusterRebuildIntervalSeconds"`
	// The maximum number of rebuild jobs that can run simultaneously for a single node.
	RebuildConcurrencyPerNode int32 `json:"rebuildConcurrencyPerNode"`
	// The maximum number of rebuild jobs that can run simultaneously for the whole k8s cluster.
	RebuildConcurrencyGlobally int32 `json:"rebuildConcurrencyGlobally"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutohealingConfig instantiates a new AutohealingConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutohealingConfig(enableAutoHealing bool, rebuildStartDelaySecondsWhenNodeFail int32, rebuildStartDelaySecondsWhenInMaintenance int32, minClusterRebuildIntervalSeconds int32, rebuildConcurrencyPerNode int32, rebuildConcurrencyGlobally int32) *AutohealingConfig {
	this := AutohealingConfig{}
	this.EnableAutoHealing = enableAutoHealing
	this.RebuildStartDelaySecondsWhenNodeFail = rebuildStartDelaySecondsWhenNodeFail
	this.RebuildStartDelaySecondsWhenInMaintenance = rebuildStartDelaySecondsWhenInMaintenance
	this.MinClusterRebuildIntervalSeconds = minClusterRebuildIntervalSeconds
	this.RebuildConcurrencyPerNode = rebuildConcurrencyPerNode
	this.RebuildConcurrencyGlobally = rebuildConcurrencyGlobally
	return &this
}

// NewAutohealingConfigWithDefaults instantiates a new AutohealingConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutohealingConfigWithDefaults() *AutohealingConfig {
	this := AutohealingConfig{}
	return &this
}

// GetEnableAutoHealing returns the EnableAutoHealing field value.
func (o *AutohealingConfig) GetEnableAutoHealing() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EnableAutoHealing
}

// GetEnableAutoHealingOk returns a tuple with the EnableAutoHealing field value
// and a boolean to check if the value has been set.
func (o *AutohealingConfig) GetEnableAutoHealingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableAutoHealing, true
}

// SetEnableAutoHealing sets field value.
func (o *AutohealingConfig) SetEnableAutoHealing(v bool) {
	o.EnableAutoHealing = v
}

// GetRebuildStartDelaySecondsWhenNodeFail returns the RebuildStartDelaySecondsWhenNodeFail field value.
func (o *AutohealingConfig) GetRebuildStartDelaySecondsWhenNodeFail() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.RebuildStartDelaySecondsWhenNodeFail
}

// GetRebuildStartDelaySecondsWhenNodeFailOk returns a tuple with the RebuildStartDelaySecondsWhenNodeFail field value
// and a boolean to check if the value has been set.
func (o *AutohealingConfig) GetRebuildStartDelaySecondsWhenNodeFailOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RebuildStartDelaySecondsWhenNodeFail, true
}

// SetRebuildStartDelaySecondsWhenNodeFail sets field value.
func (o *AutohealingConfig) SetRebuildStartDelaySecondsWhenNodeFail(v int32) {
	o.RebuildStartDelaySecondsWhenNodeFail = v
}

// GetRebuildStartDelaySecondsWhenInMaintenance returns the RebuildStartDelaySecondsWhenInMaintenance field value.
func (o *AutohealingConfig) GetRebuildStartDelaySecondsWhenInMaintenance() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.RebuildStartDelaySecondsWhenInMaintenance
}

// GetRebuildStartDelaySecondsWhenInMaintenanceOk returns a tuple with the RebuildStartDelaySecondsWhenInMaintenance field value
// and a boolean to check if the value has been set.
func (o *AutohealingConfig) GetRebuildStartDelaySecondsWhenInMaintenanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RebuildStartDelaySecondsWhenInMaintenance, true
}

// SetRebuildStartDelaySecondsWhenInMaintenance sets field value.
func (o *AutohealingConfig) SetRebuildStartDelaySecondsWhenInMaintenance(v int32) {
	o.RebuildStartDelaySecondsWhenInMaintenance = v
}

// GetMinClusterRebuildIntervalSeconds returns the MinClusterRebuildIntervalSeconds field value.
func (o *AutohealingConfig) GetMinClusterRebuildIntervalSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.MinClusterRebuildIntervalSeconds
}

// GetMinClusterRebuildIntervalSecondsOk returns a tuple with the MinClusterRebuildIntervalSeconds field value
// and a boolean to check if the value has been set.
func (o *AutohealingConfig) GetMinClusterRebuildIntervalSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinClusterRebuildIntervalSeconds, true
}

// SetMinClusterRebuildIntervalSeconds sets field value.
func (o *AutohealingConfig) SetMinClusterRebuildIntervalSeconds(v int32) {
	o.MinClusterRebuildIntervalSeconds = v
}

// GetRebuildConcurrencyPerNode returns the RebuildConcurrencyPerNode field value.
func (o *AutohealingConfig) GetRebuildConcurrencyPerNode() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.RebuildConcurrencyPerNode
}

// GetRebuildConcurrencyPerNodeOk returns a tuple with the RebuildConcurrencyPerNode field value
// and a boolean to check if the value has been set.
func (o *AutohealingConfig) GetRebuildConcurrencyPerNodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RebuildConcurrencyPerNode, true
}

// SetRebuildConcurrencyPerNode sets field value.
func (o *AutohealingConfig) SetRebuildConcurrencyPerNode(v int32) {
	o.RebuildConcurrencyPerNode = v
}

// GetRebuildConcurrencyGlobally returns the RebuildConcurrencyGlobally field value.
func (o *AutohealingConfig) GetRebuildConcurrencyGlobally() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.RebuildConcurrencyGlobally
}

// GetRebuildConcurrencyGloballyOk returns a tuple with the RebuildConcurrencyGlobally field value
// and a boolean to check if the value has been set.
func (o *AutohealingConfig) GetRebuildConcurrencyGloballyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RebuildConcurrencyGlobally, true
}

// SetRebuildConcurrencyGlobally sets field value.
func (o *AutohealingConfig) SetRebuildConcurrencyGlobally(v int32) {
	o.RebuildConcurrencyGlobally = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutohealingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enableAutoHealing"] = o.EnableAutoHealing
	toSerialize["rebuildStartDelaySecondsWhenNodeFail"] = o.RebuildStartDelaySecondsWhenNodeFail
	toSerialize["rebuildStartDelaySecondsWhenInMaintenance"] = o.RebuildStartDelaySecondsWhenInMaintenance
	toSerialize["minClusterRebuildIntervalSeconds"] = o.MinClusterRebuildIntervalSeconds
	toSerialize["rebuildConcurrencyPerNode"] = o.RebuildConcurrencyPerNode
	toSerialize["rebuildConcurrencyGlobally"] = o.RebuildConcurrencyGlobally

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutohealingConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnableAutoHealing                         *bool  `json:"enableAutoHealing"`
		RebuildStartDelaySecondsWhenNodeFail      *int32 `json:"rebuildStartDelaySecondsWhenNodeFail"`
		RebuildStartDelaySecondsWhenInMaintenance *int32 `json:"rebuildStartDelaySecondsWhenInMaintenance"`
		MinClusterRebuildIntervalSeconds          *int32 `json:"minClusterRebuildIntervalSeconds"`
		RebuildConcurrencyPerNode                 *int32 `json:"rebuildConcurrencyPerNode"`
		RebuildConcurrencyGlobally                *int32 `json:"rebuildConcurrencyGlobally"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EnableAutoHealing == nil {
		return fmt.Errorf("required field enableAutoHealing missing")
	}
	if all.RebuildStartDelaySecondsWhenNodeFail == nil {
		return fmt.Errorf("required field rebuildStartDelaySecondsWhenNodeFail missing")
	}
	if all.RebuildStartDelaySecondsWhenInMaintenance == nil {
		return fmt.Errorf("required field rebuildStartDelaySecondsWhenInMaintenance missing")
	}
	if all.MinClusterRebuildIntervalSeconds == nil {
		return fmt.Errorf("required field minClusterRebuildIntervalSeconds missing")
	}
	if all.RebuildConcurrencyPerNode == nil {
		return fmt.Errorf("required field rebuildConcurrencyPerNode missing")
	}
	if all.RebuildConcurrencyGlobally == nil {
		return fmt.Errorf("required field rebuildConcurrencyGlobally missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enableAutoHealing", "rebuildStartDelaySecondsWhenNodeFail", "rebuildStartDelaySecondsWhenInMaintenance", "minClusterRebuildIntervalSeconds", "rebuildConcurrencyPerNode", "rebuildConcurrencyGlobally"})
	} else {
		return err
	}
	o.EnableAutoHealing = *all.EnableAutoHealing
	o.RebuildStartDelaySecondsWhenNodeFail = *all.RebuildStartDelaySecondsWhenNodeFail
	o.RebuildStartDelaySecondsWhenInMaintenance = *all.RebuildStartDelaySecondsWhenInMaintenance
	o.MinClusterRebuildIntervalSeconds = *all.MinClusterRebuildIntervalSeconds
	o.RebuildConcurrencyPerNode = *all.RebuildConcurrencyPerNode
	o.RebuildConcurrencyGlobally = *all.RebuildConcurrencyGlobally

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
