// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AutohealingConfig cluster instance autohealing process config

type AutohealingConfig struct {
	// Set to true to pause the auto-healing process, preventing automatic rebuilding of instances when a node fails or is in maintenance mode.
	PauseAutoHealingOnNodeUnavailable bool `json:"pauseAutoHealingOnNodeUnavailable"`
	// The number of seconds to wait before starting a rebuild when a node fails.
	RebuildStartDelaySeconds int32 `json:"rebuildStartDelaySeconds"`
	// The minimum time in seconds between consecutive rebuild jobs for one cluster.
	MinClusterRebuildIntervalSeconds int32 `json:"minClusterRebuildIntervalSeconds"`
	// The maximum number of rebuild jobs that can run simultaneously for a single node.
	RebuildConcurrencyPerNode int32 `json:"rebuildConcurrencyPerNode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutohealingConfig instantiates a new AutohealingConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutohealingConfig(pauseAutoHealingOnNodeUnavailable bool, rebuildStartDelaySeconds int32, minClusterRebuildIntervalSeconds int32, rebuildConcurrencyPerNode int32) *AutohealingConfig {
	this := AutohealingConfig{}
	this.PauseAutoHealingOnNodeUnavailable = pauseAutoHealingOnNodeUnavailable
	this.RebuildStartDelaySeconds = rebuildStartDelaySeconds
	this.MinClusterRebuildIntervalSeconds = minClusterRebuildIntervalSeconds
	this.RebuildConcurrencyPerNode = rebuildConcurrencyPerNode
	return &this
}

// NewAutohealingConfigWithDefaults instantiates a new AutohealingConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutohealingConfigWithDefaults() *AutohealingConfig {
	this := AutohealingConfig{}
	return &this
}

// GetPauseAutoHealingOnNodeUnavailable returns the PauseAutoHealingOnNodeUnavailable field value.
func (o *AutohealingConfig) GetPauseAutoHealingOnNodeUnavailable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.PauseAutoHealingOnNodeUnavailable
}

// GetPauseAutoHealingOnNodeUnavailableOk returns a tuple with the PauseAutoHealingOnNodeUnavailable field value
// and a boolean to check if the value has been set.
func (o *AutohealingConfig) GetPauseAutoHealingOnNodeUnavailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PauseAutoHealingOnNodeUnavailable, true
}

// SetPauseAutoHealingOnNodeUnavailable sets field value.
func (o *AutohealingConfig) SetPauseAutoHealingOnNodeUnavailable(v bool) {
	o.PauseAutoHealingOnNodeUnavailable = v
}

// GetRebuildStartDelaySeconds returns the RebuildStartDelaySeconds field value.
func (o *AutohealingConfig) GetRebuildStartDelaySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.RebuildStartDelaySeconds
}

// GetRebuildStartDelaySecondsOk returns a tuple with the RebuildStartDelaySeconds field value
// and a boolean to check if the value has been set.
func (o *AutohealingConfig) GetRebuildStartDelaySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RebuildStartDelaySeconds, true
}

// SetRebuildStartDelaySeconds sets field value.
func (o *AutohealingConfig) SetRebuildStartDelaySeconds(v int32) {
	o.RebuildStartDelaySeconds = v
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

// MarshalJSON serializes the struct using spec logic.
func (o AutohealingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pauseAutoHealingOnNodeUnavailable"] = o.PauseAutoHealingOnNodeUnavailable
	toSerialize["rebuildStartDelaySeconds"] = o.RebuildStartDelaySeconds
	toSerialize["minClusterRebuildIntervalSeconds"] = o.MinClusterRebuildIntervalSeconds
	toSerialize["rebuildConcurrencyPerNode"] = o.RebuildConcurrencyPerNode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutohealingConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PauseAutoHealingOnNodeUnavailable *bool  `json:"pauseAutoHealingOnNodeUnavailable"`
		RebuildStartDelaySeconds          *int32 `json:"rebuildStartDelaySeconds"`
		MinClusterRebuildIntervalSeconds  *int32 `json:"minClusterRebuildIntervalSeconds"`
		RebuildConcurrencyPerNode         *int32 `json:"rebuildConcurrencyPerNode"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PauseAutoHealingOnNodeUnavailable == nil {
		return fmt.Errorf("required field pauseAutoHealingOnNodeUnavailable missing")
	}
	if all.RebuildStartDelaySeconds == nil {
		return fmt.Errorf("required field rebuildStartDelaySeconds missing")
	}
	if all.MinClusterRebuildIntervalSeconds == nil {
		return fmt.Errorf("required field minClusterRebuildIntervalSeconds missing")
	}
	if all.RebuildConcurrencyPerNode == nil {
		return fmt.Errorf("required field rebuildConcurrencyPerNode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pauseAutoHealingOnNodeUnavailable", "rebuildStartDelaySeconds", "minClusterRebuildIntervalSeconds", "rebuildConcurrencyPerNode"})
	} else {
		return err
	}
	o.PauseAutoHealingOnNodeUnavailable = *all.PauseAutoHealingOnNodeUnavailable
	o.RebuildStartDelaySeconds = *all.RebuildStartDelaySeconds
	o.MinClusterRebuildIntervalSeconds = *all.MinClusterRebuildIntervalSeconds
	o.RebuildConcurrencyPerNode = *all.RebuildConcurrencyPerNode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
