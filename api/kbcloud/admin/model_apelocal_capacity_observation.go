// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ApelocalCapacityObservation A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
type ApelocalCapacityObservation struct {
	// Known-node sum in GiB or ratio in percent. Omitted when no inputs are known or a ratio denominator is zero. Ratios can exceed 100. Remaining quota is zero when there are selected nodes but none are ready.
	Value *float64 `json:"value,omitempty"`
	// Number of nodes with all inputs for this measurement.
	ObservedNodeCount *int32 `json:"observedNodeCount,omitempty"`
	// Number of selected nodes relevant to this measurement; only ready schedulable nodes for remaining quota.
	ExpectedNodeCount *int32 `json:"expectedNodeCount,omitempty"`
	// Sorted names of relevant nodes with missing or inconsistent inputs for this measurement.
	MissingNodes []string `json:"missingNodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApelocalCapacityObservation instantiates a new ApelocalCapacityObservation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApelocalCapacityObservation() *ApelocalCapacityObservation {
	this := ApelocalCapacityObservation{}
	return &this
}

// NewApelocalCapacityObservationWithDefaults instantiates a new ApelocalCapacityObservation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApelocalCapacityObservationWithDefaults() *ApelocalCapacityObservation {
	this := ApelocalCapacityObservation{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApelocalCapacityObservation) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservation) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApelocalCapacityObservation) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *ApelocalCapacityObservation) SetValue(v float64) {
	o.Value = &v
}

// GetObservedNodeCount returns the ObservedNodeCount field value if set, zero value otherwise.
func (o *ApelocalCapacityObservation) GetObservedNodeCount() int32 {
	if o == nil || o.ObservedNodeCount == nil {
		var ret int32
		return ret
	}
	return *o.ObservedNodeCount
}

// GetObservedNodeCountOk returns a tuple with the ObservedNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservation) GetObservedNodeCountOk() (*int32, bool) {
	if o == nil || o.ObservedNodeCount == nil {
		return nil, false
	}
	return o.ObservedNodeCount, true
}

// HasObservedNodeCount returns a boolean if a field has been set.
func (o *ApelocalCapacityObservation) HasObservedNodeCount() bool {
	return o != nil && o.ObservedNodeCount != nil
}

// SetObservedNodeCount gets a reference to the given int32 and assigns it to the ObservedNodeCount field.
func (o *ApelocalCapacityObservation) SetObservedNodeCount(v int32) {
	o.ObservedNodeCount = &v
}

// GetExpectedNodeCount returns the ExpectedNodeCount field value if set, zero value otherwise.
func (o *ApelocalCapacityObservation) GetExpectedNodeCount() int32 {
	if o == nil || o.ExpectedNodeCount == nil {
		var ret int32
		return ret
	}
	return *o.ExpectedNodeCount
}

// GetExpectedNodeCountOk returns a tuple with the ExpectedNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservation) GetExpectedNodeCountOk() (*int32, bool) {
	if o == nil || o.ExpectedNodeCount == nil {
		return nil, false
	}
	return o.ExpectedNodeCount, true
}

// HasExpectedNodeCount returns a boolean if a field has been set.
func (o *ApelocalCapacityObservation) HasExpectedNodeCount() bool {
	return o != nil && o.ExpectedNodeCount != nil
}

// SetExpectedNodeCount gets a reference to the given int32 and assigns it to the ExpectedNodeCount field.
func (o *ApelocalCapacityObservation) SetExpectedNodeCount(v int32) {
	o.ExpectedNodeCount = &v
}

// GetMissingNodes returns the MissingNodes field value if set, zero value otherwise.
func (o *ApelocalCapacityObservation) GetMissingNodes() []string {
	if o == nil || o.MissingNodes == nil {
		var ret []string
		return ret
	}
	return o.MissingNodes
}

// GetMissingNodesOk returns a tuple with the MissingNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservation) GetMissingNodesOk() (*[]string, bool) {
	if o == nil || o.MissingNodes == nil {
		return nil, false
	}
	return &o.MissingNodes, true
}

// HasMissingNodes returns a boolean if a field has been set.
func (o *ApelocalCapacityObservation) HasMissingNodes() bool {
	return o != nil && o.MissingNodes != nil
}

// SetMissingNodes gets a reference to the given []string and assigns it to the MissingNodes field.
func (o *ApelocalCapacityObservation) SetMissingNodes(v []string) {
	o.MissingNodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApelocalCapacityObservation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ObservedNodeCount != nil {
		toSerialize["observedNodeCount"] = o.ObservedNodeCount
	}
	if o.ExpectedNodeCount != nil {
		toSerialize["expectedNodeCount"] = o.ExpectedNodeCount
	}
	if o.MissingNodes != nil {
		toSerialize["missingNodes"] = o.MissingNodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApelocalCapacityObservation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Value             *float64 `json:"value,omitempty"`
		ObservedNodeCount *int32   `json:"observedNodeCount,omitempty"`
		ExpectedNodeCount *int32   `json:"expectedNodeCount,omitempty"`
		MissingNodes      []string `json:"missingNodes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"value", "observedNodeCount", "expectedNodeCount", "missingNodes"})
	} else {
		return err
	}
	o.Value = all.Value
	o.ObservedNodeCount = all.ObservedNodeCount
	o.ExpectedNodeCount = all.ExpectedNodeCount
	o.MissingNodes = all.MissingNodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
