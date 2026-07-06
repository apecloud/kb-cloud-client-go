// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type KoordinatorReservationStatus struct {
	Phase    *string `json:"phase,omitempty"`
	NodeName *string `json:"nodeName,omitempty"`
	// Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
	Allocatable *KoordinatorReservationResources `json:"allocatable,omitempty"`
	// Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
	Allocated *KoordinatorReservationResources `json:"allocated,omitempty"`
	// Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
	Available *KoordinatorReservationResources `json:"available,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKoordinatorReservationStatus instantiates a new KoordinatorReservationStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKoordinatorReservationStatus() *KoordinatorReservationStatus {
	this := KoordinatorReservationStatus{}
	return &this
}

// NewKoordinatorReservationStatusWithDefaults instantiates a new KoordinatorReservationStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKoordinatorReservationStatusWithDefaults() *KoordinatorReservationStatus {
	this := KoordinatorReservationStatus{}
	return &this
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *KoordinatorReservationStatus) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationStatus) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *KoordinatorReservationStatus) HasPhase() bool {
	return o != nil && o.Phase != nil
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *KoordinatorReservationStatus) SetPhase(v string) {
	o.Phase = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *KoordinatorReservationStatus) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationStatus) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *KoordinatorReservationStatus) HasNodeName() bool {
	return o != nil && o.NodeName != nil
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *KoordinatorReservationStatus) SetNodeName(v string) {
	o.NodeName = &v
}

// GetAllocatable returns the Allocatable field value if set, zero value otherwise.
func (o *KoordinatorReservationStatus) GetAllocatable() KoordinatorReservationResources {
	if o == nil || o.Allocatable == nil {
		var ret KoordinatorReservationResources
		return ret
	}
	return *o.Allocatable
}

// GetAllocatableOk returns a tuple with the Allocatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationStatus) GetAllocatableOk() (*KoordinatorReservationResources, bool) {
	if o == nil || o.Allocatable == nil {
		return nil, false
	}
	return o.Allocatable, true
}

// HasAllocatable returns a boolean if a field has been set.
func (o *KoordinatorReservationStatus) HasAllocatable() bool {
	return o != nil && o.Allocatable != nil
}

// SetAllocatable gets a reference to the given KoordinatorReservationResources and assigns it to the Allocatable field.
func (o *KoordinatorReservationStatus) SetAllocatable(v KoordinatorReservationResources) {
	o.Allocatable = &v
}

// GetAllocated returns the Allocated field value if set, zero value otherwise.
func (o *KoordinatorReservationStatus) GetAllocated() KoordinatorReservationResources {
	if o == nil || o.Allocated == nil {
		var ret KoordinatorReservationResources
		return ret
	}
	return *o.Allocated
}

// GetAllocatedOk returns a tuple with the Allocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationStatus) GetAllocatedOk() (*KoordinatorReservationResources, bool) {
	if o == nil || o.Allocated == nil {
		return nil, false
	}
	return o.Allocated, true
}

// HasAllocated returns a boolean if a field has been set.
func (o *KoordinatorReservationStatus) HasAllocated() bool {
	return o != nil && o.Allocated != nil
}

// SetAllocated gets a reference to the given KoordinatorReservationResources and assigns it to the Allocated field.
func (o *KoordinatorReservationStatus) SetAllocated(v KoordinatorReservationResources) {
	o.Allocated = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *KoordinatorReservationStatus) GetAvailable() KoordinatorReservationResources {
	if o == nil || o.Available == nil {
		var ret KoordinatorReservationResources
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationStatus) GetAvailableOk() (*KoordinatorReservationResources, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *KoordinatorReservationStatus) HasAvailable() bool {
	return o != nil && o.Available != nil
}

// SetAvailable gets a reference to the given KoordinatorReservationResources and assigns it to the Available field.
func (o *KoordinatorReservationStatus) SetAvailable(v KoordinatorReservationResources) {
	o.Available = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KoordinatorReservationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.NodeName != nil {
		toSerialize["nodeName"] = o.NodeName
	}
	if o.Allocatable != nil {
		toSerialize["allocatable"] = o.Allocatable
	}
	if o.Allocated != nil {
		toSerialize["allocated"] = o.Allocated
	}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KoordinatorReservationStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Phase       *string                          `json:"phase,omitempty"`
		NodeName    *string                          `json:"nodeName,omitempty"`
		Allocatable *KoordinatorReservationResources `json:"allocatable,omitempty"`
		Allocated   *KoordinatorReservationResources `json:"allocated,omitempty"`
		Available   *KoordinatorReservationResources `json:"available,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"phase", "nodeName", "allocatable", "allocated", "available"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Phase = all.Phase
	o.NodeName = all.NodeName
	if all.Allocatable != nil && all.Allocatable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Allocatable = all.Allocatable
	if all.Allocated != nil && all.Allocated.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Allocated = all.Allocated
	if all.Available != nil && all.Available.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Available = all.Available

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
