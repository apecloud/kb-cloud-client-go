// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBHealthSnapshot struct {
	CheckedAt        *time.Time            `json:"checkedAt,omitempty"`
	Version          *string               `json:"version,omitempty"`
	Leader           *string               `json:"leader,omitempty"`
	Nodes            []DolphinDBNodeHealth `json:"nodes,omitempty"`
	EndpointPort     *DolphinDBHealthCheck `json:"endpointPort,omitempty"`
	EndpointQuery    *DolphinDBHealthCheck `json:"endpointQuery,omitempty"`
	Inventory        *DolphinDBHealthCheck `json:"inventory,omitempty"`
	ControllerLeader *DolphinDBHealthCheck `json:"controllerLeader,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBHealthSnapshot instantiates a new DolphinDBHealthSnapshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBHealthSnapshot() *DolphinDBHealthSnapshot {
	this := DolphinDBHealthSnapshot{}
	return &this
}

// NewDolphinDBHealthSnapshotWithDefaults instantiates a new DolphinDBHealthSnapshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBHealthSnapshotWithDefaults() *DolphinDBHealthSnapshot {
	this := DolphinDBHealthSnapshot{}
	return &this
}

// GetCheckedAt returns the CheckedAt field value if set, zero value otherwise.
func (o *DolphinDBHealthSnapshot) GetCheckedAt() time.Time {
	if o == nil || o.CheckedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CheckedAt
}

// GetCheckedAtOk returns a tuple with the CheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthSnapshot) GetCheckedAtOk() (*time.Time, bool) {
	if o == nil || o.CheckedAt == nil {
		return nil, false
	}
	return o.CheckedAt, true
}

// HasCheckedAt returns a boolean if a field has been set.
func (o *DolphinDBHealthSnapshot) HasCheckedAt() bool {
	return o != nil && o.CheckedAt != nil
}

// SetCheckedAt gets a reference to the given time.Time and assigns it to the CheckedAt field.
func (o *DolphinDBHealthSnapshot) SetCheckedAt(v time.Time) {
	o.CheckedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DolphinDBHealthSnapshot) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthSnapshot) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DolphinDBHealthSnapshot) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DolphinDBHealthSnapshot) SetVersion(v string) {
	o.Version = &v
}

// GetLeader returns the Leader field value if set, zero value otherwise.
func (o *DolphinDBHealthSnapshot) GetLeader() string {
	if o == nil || o.Leader == nil {
		var ret string
		return ret
	}
	return *o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthSnapshot) GetLeaderOk() (*string, bool) {
	if o == nil || o.Leader == nil {
		return nil, false
	}
	return o.Leader, true
}

// HasLeader returns a boolean if a field has been set.
func (o *DolphinDBHealthSnapshot) HasLeader() bool {
	return o != nil && o.Leader != nil
}

// SetLeader gets a reference to the given string and assigns it to the Leader field.
func (o *DolphinDBHealthSnapshot) SetLeader(v string) {
	o.Leader = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *DolphinDBHealthSnapshot) GetNodes() []DolphinDBNodeHealth {
	if o == nil || o.Nodes == nil {
		var ret []DolphinDBNodeHealth
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthSnapshot) GetNodesOk() (*[]DolphinDBNodeHealth, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *DolphinDBHealthSnapshot) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []DolphinDBNodeHealth and assigns it to the Nodes field.
func (o *DolphinDBHealthSnapshot) SetNodes(v []DolphinDBNodeHealth) {
	o.Nodes = v
}

// GetEndpointPort returns the EndpointPort field value if set, zero value otherwise.
func (o *DolphinDBHealthSnapshot) GetEndpointPort() DolphinDBHealthCheck {
	if o == nil || o.EndpointPort == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.EndpointPort
}

// GetEndpointPortOk returns a tuple with the EndpointPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthSnapshot) GetEndpointPortOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.EndpointPort == nil {
		return nil, false
	}
	return o.EndpointPort, true
}

// HasEndpointPort returns a boolean if a field has been set.
func (o *DolphinDBHealthSnapshot) HasEndpointPort() bool {
	return o != nil && o.EndpointPort != nil
}

// SetEndpointPort gets a reference to the given DolphinDBHealthCheck and assigns it to the EndpointPort field.
func (o *DolphinDBHealthSnapshot) SetEndpointPort(v DolphinDBHealthCheck) {
	o.EndpointPort = &v
}

// GetEndpointQuery returns the EndpointQuery field value if set, zero value otherwise.
func (o *DolphinDBHealthSnapshot) GetEndpointQuery() DolphinDBHealthCheck {
	if o == nil || o.EndpointQuery == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.EndpointQuery
}

// GetEndpointQueryOk returns a tuple with the EndpointQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthSnapshot) GetEndpointQueryOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.EndpointQuery == nil {
		return nil, false
	}
	return o.EndpointQuery, true
}

// HasEndpointQuery returns a boolean if a field has been set.
func (o *DolphinDBHealthSnapshot) HasEndpointQuery() bool {
	return o != nil && o.EndpointQuery != nil
}

// SetEndpointQuery gets a reference to the given DolphinDBHealthCheck and assigns it to the EndpointQuery field.
func (o *DolphinDBHealthSnapshot) SetEndpointQuery(v DolphinDBHealthCheck) {
	o.EndpointQuery = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *DolphinDBHealthSnapshot) GetInventory() DolphinDBHealthCheck {
	if o == nil || o.Inventory == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthSnapshot) GetInventoryOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *DolphinDBHealthSnapshot) HasInventory() bool {
	return o != nil && o.Inventory != nil
}

// SetInventory gets a reference to the given DolphinDBHealthCheck and assigns it to the Inventory field.
func (o *DolphinDBHealthSnapshot) SetInventory(v DolphinDBHealthCheck) {
	o.Inventory = &v
}

// GetControllerLeader returns the ControllerLeader field value if set, zero value otherwise.
func (o *DolphinDBHealthSnapshot) GetControllerLeader() DolphinDBHealthCheck {
	if o == nil || o.ControllerLeader == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.ControllerLeader
}

// GetControllerLeaderOk returns a tuple with the ControllerLeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBHealthSnapshot) GetControllerLeaderOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.ControllerLeader == nil {
		return nil, false
	}
	return o.ControllerLeader, true
}

// HasControllerLeader returns a boolean if a field has been set.
func (o *DolphinDBHealthSnapshot) HasControllerLeader() bool {
	return o != nil && o.ControllerLeader != nil
}

// SetControllerLeader gets a reference to the given DolphinDBHealthCheck and assigns it to the ControllerLeader field.
func (o *DolphinDBHealthSnapshot) SetControllerLeader(v DolphinDBHealthCheck) {
	o.ControllerLeader = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBHealthSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CheckedAt != nil {
		if o.CheckedAt.Nanosecond() == 0 {
			toSerialize["checkedAt"] = o.CheckedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["checkedAt"] = o.CheckedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Leader != nil {
		toSerialize["leader"] = o.Leader
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if o.EndpointPort != nil {
		toSerialize["endpointPort"] = o.EndpointPort
	}
	if o.EndpointQuery != nil {
		toSerialize["endpointQuery"] = o.EndpointQuery
	}
	if o.Inventory != nil {
		toSerialize["inventory"] = o.Inventory
	}
	if o.ControllerLeader != nil {
		toSerialize["controllerLeader"] = o.ControllerLeader
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBHealthSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckedAt        *time.Time            `json:"checkedAt,omitempty"`
		Version          *string               `json:"version,omitempty"`
		Leader           *string               `json:"leader,omitempty"`
		Nodes            []DolphinDBNodeHealth `json:"nodes,omitempty"`
		EndpointPort     *DolphinDBHealthCheck `json:"endpointPort,omitempty"`
		EndpointQuery    *DolphinDBHealthCheck `json:"endpointQuery,omitempty"`
		Inventory        *DolphinDBHealthCheck `json:"inventory,omitempty"`
		ControllerLeader *DolphinDBHealthCheck `json:"controllerLeader,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"checkedAt", "version", "leader", "nodes", "endpointPort", "endpointQuery", "inventory", "controllerLeader"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CheckedAt = all.CheckedAt
	o.Version = all.Version
	o.Leader = all.Leader
	o.Nodes = all.Nodes
	if all.EndpointPort != nil && all.EndpointPort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EndpointPort = all.EndpointPort
	if all.EndpointQuery != nil && all.EndpointQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EndpointQuery = all.EndpointQuery
	if all.Inventory != nil && all.Inventory.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Inventory = all.Inventory
	if all.ControllerLeader != nil && all.ControllerLeader.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ControllerLeader = all.ControllerLeader

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
