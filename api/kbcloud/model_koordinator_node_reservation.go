// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KoordinatorNodeReservation struct {
	ResourceClass   KoordinatorReservationResourceClass `json:"resourceClass"`
	ReservationName *string                             `json:"reservationName,omitempty"`
	// Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
	Resources KoordinatorReservationResources `json:"resources"`
	// Kubernetes duration string. Koordinator defaults omitted ttl to 24h.
	Ttl            *string                               `json:"ttl,omitempty"`
	AllocateOnce   *bool                                 `json:"allocateOnce,omitempty"`
	PreAllocation  *bool                                 `json:"preAllocation,omitempty"`
	AllocatePolicy *KoordinatorReservationAllocatePolicy `json:"allocatePolicy,omitempty"`
	Status         *KoordinatorReservationStatus         `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKoordinatorNodeReservation instantiates a new KoordinatorNodeReservation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKoordinatorNodeReservation(resourceClass KoordinatorReservationResourceClass, resources KoordinatorReservationResources) *KoordinatorNodeReservation {
	this := KoordinatorNodeReservation{}
	this.ResourceClass = resourceClass
	this.Resources = resources
	var ttl string = "24h"
	this.Ttl = &ttl
	return &this
}

// NewKoordinatorNodeReservationWithDefaults instantiates a new KoordinatorNodeReservation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKoordinatorNodeReservationWithDefaults() *KoordinatorNodeReservation {
	this := KoordinatorNodeReservation{}
	var ttl string = "24h"
	this.Ttl = &ttl
	return &this
}

// GetResourceClass returns the ResourceClass field value.
func (o *KoordinatorNodeReservation) GetResourceClass() KoordinatorReservationResourceClass {
	if o == nil {
		var ret KoordinatorReservationResourceClass
		return ret
	}
	return o.ResourceClass
}

// GetResourceClassOk returns a tuple with the ResourceClass field value
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservation) GetResourceClassOk() (*KoordinatorReservationResourceClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceClass, true
}

// SetResourceClass sets field value.
func (o *KoordinatorNodeReservation) SetResourceClass(v KoordinatorReservationResourceClass) {
	o.ResourceClass = v
}

// GetReservationName returns the ReservationName field value if set, zero value otherwise.
func (o *KoordinatorNodeReservation) GetReservationName() string {
	if o == nil || o.ReservationName == nil {
		var ret string
		return ret
	}
	return *o.ReservationName
}

// GetReservationNameOk returns a tuple with the ReservationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservation) GetReservationNameOk() (*string, bool) {
	if o == nil || o.ReservationName == nil {
		return nil, false
	}
	return o.ReservationName, true
}

// HasReservationName returns a boolean if a field has been set.
func (o *KoordinatorNodeReservation) HasReservationName() bool {
	return o != nil && o.ReservationName != nil
}

// SetReservationName gets a reference to the given string and assigns it to the ReservationName field.
func (o *KoordinatorNodeReservation) SetReservationName(v string) {
	o.ReservationName = &v
}

// GetResources returns the Resources field value.
func (o *KoordinatorNodeReservation) GetResources() KoordinatorReservationResources {
	if o == nil {
		var ret KoordinatorReservationResources
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservation) GetResourcesOk() (*KoordinatorReservationResources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value.
func (o *KoordinatorNodeReservation) SetResources(v KoordinatorReservationResources) {
	o.Resources = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *KoordinatorNodeReservation) GetTtl() string {
	if o == nil || o.Ttl == nil {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservation) GetTtlOk() (*string, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *KoordinatorNodeReservation) HasTtl() bool {
	return o != nil && o.Ttl != nil
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *KoordinatorNodeReservation) SetTtl(v string) {
	o.Ttl = &v
}

// GetAllocateOnce returns the AllocateOnce field value if set, zero value otherwise.
func (o *KoordinatorNodeReservation) GetAllocateOnce() bool {
	if o == nil || o.AllocateOnce == nil {
		var ret bool
		return ret
	}
	return *o.AllocateOnce
}

// GetAllocateOnceOk returns a tuple with the AllocateOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservation) GetAllocateOnceOk() (*bool, bool) {
	if o == nil || o.AllocateOnce == nil {
		return nil, false
	}
	return o.AllocateOnce, true
}

// HasAllocateOnce returns a boolean if a field has been set.
func (o *KoordinatorNodeReservation) HasAllocateOnce() bool {
	return o != nil && o.AllocateOnce != nil
}

// SetAllocateOnce gets a reference to the given bool and assigns it to the AllocateOnce field.
func (o *KoordinatorNodeReservation) SetAllocateOnce(v bool) {
	o.AllocateOnce = &v
}

// GetPreAllocation returns the PreAllocation field value if set, zero value otherwise.
func (o *KoordinatorNodeReservation) GetPreAllocation() bool {
	if o == nil || o.PreAllocation == nil {
		var ret bool
		return ret
	}
	return *o.PreAllocation
}

// GetPreAllocationOk returns a tuple with the PreAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservation) GetPreAllocationOk() (*bool, bool) {
	if o == nil || o.PreAllocation == nil {
		return nil, false
	}
	return o.PreAllocation, true
}

// HasPreAllocation returns a boolean if a field has been set.
func (o *KoordinatorNodeReservation) HasPreAllocation() bool {
	return o != nil && o.PreAllocation != nil
}

// SetPreAllocation gets a reference to the given bool and assigns it to the PreAllocation field.
func (o *KoordinatorNodeReservation) SetPreAllocation(v bool) {
	o.PreAllocation = &v
}

// GetAllocatePolicy returns the AllocatePolicy field value if set, zero value otherwise.
func (o *KoordinatorNodeReservation) GetAllocatePolicy() KoordinatorReservationAllocatePolicy {
	if o == nil || o.AllocatePolicy == nil {
		var ret KoordinatorReservationAllocatePolicy
		return ret
	}
	return *o.AllocatePolicy
}

// GetAllocatePolicyOk returns a tuple with the AllocatePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservation) GetAllocatePolicyOk() (*KoordinatorReservationAllocatePolicy, bool) {
	if o == nil || o.AllocatePolicy == nil {
		return nil, false
	}
	return o.AllocatePolicy, true
}

// HasAllocatePolicy returns a boolean if a field has been set.
func (o *KoordinatorNodeReservation) HasAllocatePolicy() bool {
	return o != nil && o.AllocatePolicy != nil
}

// SetAllocatePolicy gets a reference to the given KoordinatorReservationAllocatePolicy and assigns it to the AllocatePolicy field.
func (o *KoordinatorNodeReservation) SetAllocatePolicy(v KoordinatorReservationAllocatePolicy) {
	o.AllocatePolicy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KoordinatorNodeReservation) GetStatus() KoordinatorReservationStatus {
	if o == nil || o.Status == nil {
		var ret KoordinatorReservationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservation) GetStatusOk() (*KoordinatorReservationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KoordinatorNodeReservation) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given KoordinatorReservationStatus and assigns it to the Status field.
func (o *KoordinatorNodeReservation) SetStatus(v KoordinatorReservationStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KoordinatorNodeReservation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["resourceClass"] = o.ResourceClass
	if o.ReservationName != nil {
		toSerialize["reservationName"] = o.ReservationName
	}
	toSerialize["resources"] = o.Resources
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.AllocateOnce != nil {
		toSerialize["allocateOnce"] = o.AllocateOnce
	}
	if o.PreAllocation != nil {
		toSerialize["preAllocation"] = o.PreAllocation
	}
	if o.AllocatePolicy != nil {
		toSerialize["allocatePolicy"] = o.AllocatePolicy
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KoordinatorNodeReservation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceClass   *KoordinatorReservationResourceClass  `json:"resourceClass"`
		ReservationName *string                               `json:"reservationName,omitempty"`
		Resources       *KoordinatorReservationResources      `json:"resources"`
		Ttl             *string                               `json:"ttl,omitempty"`
		AllocateOnce    *bool                                 `json:"allocateOnce,omitempty"`
		PreAllocation   *bool                                 `json:"preAllocation,omitempty"`
		AllocatePolicy  *KoordinatorReservationAllocatePolicy `json:"allocatePolicy,omitempty"`
		Status          *KoordinatorReservationStatus         `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ResourceClass == nil {
		return fmt.Errorf("required field resourceClass missing")
	}
	if all.Resources == nil {
		return fmt.Errorf("required field resources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"resourceClass", "reservationName", "resources", "ttl", "allocateOnce", "preAllocation", "allocatePolicy", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.ResourceClass.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceClass = *all.ResourceClass
	}
	o.ReservationName = all.ReservationName
	if all.Resources.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Resources = *all.Resources
	o.Ttl = all.Ttl
	o.AllocateOnce = all.AllocateOnce
	o.PreAllocation = all.PreAllocation
	if all.AllocatePolicy != nil && !all.AllocatePolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.AllocatePolicy = all.AllocatePolicy
	}
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
