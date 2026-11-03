// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClickhouseSessionList struct {
	// Collection time in RFC3339 UTC.
	CollectedAt string `json:"collectedAt"`
	// Detected deployment mode (standalone or cluster). When clusterAllReplicas fails and the local node is used instead, this is reported as standalone.
	ClusterMode *string `json:"clusterMode,omitempty"`
	// Active queries collected from system.processes, ordered by the requested metric.
	Items []ClickhouseSession `json:"items"`
	// Per-source collection status. In cluster mode one entry is returned per clusterAllReplicas expression.
	Sources []ClickhouseSource `json:"sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseSessionList instantiates a new ClickhouseSessionList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseSessionList(collectedAt string, items []ClickhouseSession, sources []ClickhouseSource) *ClickhouseSessionList {
	this := ClickhouseSessionList{}
	this.CollectedAt = collectedAt
	this.Items = items
	this.Sources = sources
	return &this
}

// NewClickhouseSessionListWithDefaults instantiates a new ClickhouseSessionList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseSessionListWithDefaults() *ClickhouseSessionList {
	this := ClickhouseSessionList{}
	return &this
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ClickhouseSessionList) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSessionList) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ClickhouseSessionList) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetClusterMode returns the ClusterMode field value if set, zero value otherwise.
func (o *ClickhouseSessionList) GetClusterMode() string {
	if o == nil || o.ClusterMode == nil {
		var ret string
		return ret
	}
	return *o.ClusterMode
}

// GetClusterModeOk returns a tuple with the ClusterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseSessionList) GetClusterModeOk() (*string, bool) {
	if o == nil || o.ClusterMode == nil {
		return nil, false
	}
	return o.ClusterMode, true
}

// HasClusterMode returns a boolean if a field has been set.
func (o *ClickhouseSessionList) HasClusterMode() bool {
	return o != nil && o.ClusterMode != nil
}

// SetClusterMode gets a reference to the given string and assigns it to the ClusterMode field.
func (o *ClickhouseSessionList) SetClusterMode(v string) {
	o.ClusterMode = &v
}

// GetItems returns the Items field value.
func (o *ClickhouseSessionList) GetItems() []ClickhouseSession {
	if o == nil {
		var ret []ClickhouseSession
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSessionList) GetItemsOk() (*[]ClickhouseSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *ClickhouseSessionList) SetItems(v []ClickhouseSession) {
	o.Items = v
}

// GetSources returns the Sources field value.
func (o *ClickhouseSessionList) GetSources() []ClickhouseSource {
	if o == nil {
		var ret []ClickhouseSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ClickhouseSessionList) GetSourcesOk() (*[]ClickhouseSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *ClickhouseSessionList) SetSources(v []ClickhouseSource) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseSessionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["collectedAt"] = o.CollectedAt
	if o.ClusterMode != nil {
		toSerialize["clusterMode"] = o.ClusterMode
	}
	toSerialize["items"] = o.Items
	toSerialize["sources"] = o.Sources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseSessionList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CollectedAt *string              `json:"collectedAt"`
		ClusterMode *string              `json:"clusterMode,omitempty"`
		Items       *[]ClickhouseSession `json:"items"`
		Sources     *[]ClickhouseSource  `json:"sources"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"collectedAt", "clusterMode", "items", "sources"})
	} else {
		return err
	}
	o.CollectedAt = *all.CollectedAt
	o.ClusterMode = all.ClusterMode
	o.Items = *all.Items
	o.Sources = *all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
