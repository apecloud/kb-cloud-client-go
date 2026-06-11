// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KoordinatorNodeReservationUpdate struct {
	// Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
	Resources KoordinatorReservationResources `json:"resources"`
	// Kubernetes duration string. Use "0" for no expiration.
	Ttl            *string                               `json:"ttl,omitempty"`
	AllocateOnce   *bool                                 `json:"allocateOnce,omitempty"`
	PreAllocation  *bool                                 `json:"preAllocation,omitempty"`
	AllocatePolicy *KoordinatorReservationAllocatePolicy `json:"allocatePolicy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKoordinatorNodeReservationUpdate instantiates a new KoordinatorNodeReservationUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKoordinatorNodeReservationUpdate(resources KoordinatorReservationResources) *KoordinatorNodeReservationUpdate {
	this := KoordinatorNodeReservationUpdate{}
	this.Resources = resources
	var ttl string = "0"
	this.Ttl = &ttl
	var allocateOnce bool = false
	this.AllocateOnce = &allocateOnce
	var preAllocation bool = false
	this.PreAllocation = &preAllocation
	var allocatePolicy KoordinatorReservationAllocatePolicy = KoordinatorReservationAllocatePolicyRestricted
	this.AllocatePolicy = &allocatePolicy
	return &this
}

// NewKoordinatorNodeReservationUpdateWithDefaults instantiates a new KoordinatorNodeReservationUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKoordinatorNodeReservationUpdateWithDefaults() *KoordinatorNodeReservationUpdate {
	this := KoordinatorNodeReservationUpdate{}
	var ttl string = "0"
	this.Ttl = &ttl
	var allocateOnce bool = false
	this.AllocateOnce = &allocateOnce
	var preAllocation bool = false
	this.PreAllocation = &preAllocation
	var allocatePolicy KoordinatorReservationAllocatePolicy = KoordinatorReservationAllocatePolicyRestricted
	this.AllocatePolicy = &allocatePolicy
	return &this
}

// GetResources returns the Resources field value.
func (o *KoordinatorNodeReservationUpdate) GetResources() KoordinatorReservationResources {
	if o == nil {
		var ret KoordinatorReservationResources
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservationUpdate) GetResourcesOk() (*KoordinatorReservationResources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value.
func (o *KoordinatorNodeReservationUpdate) SetResources(v KoordinatorReservationResources) {
	o.Resources = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *KoordinatorNodeReservationUpdate) GetTtl() string {
	if o == nil || o.Ttl == nil {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservationUpdate) GetTtlOk() (*string, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *KoordinatorNodeReservationUpdate) HasTtl() bool {
	return o != nil && o.Ttl != nil
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *KoordinatorNodeReservationUpdate) SetTtl(v string) {
	o.Ttl = &v
}

// GetAllocateOnce returns the AllocateOnce field value if set, zero value otherwise.
func (o *KoordinatorNodeReservationUpdate) GetAllocateOnce() bool {
	if o == nil || o.AllocateOnce == nil {
		var ret bool
		return ret
	}
	return *o.AllocateOnce
}

// GetAllocateOnceOk returns a tuple with the AllocateOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservationUpdate) GetAllocateOnceOk() (*bool, bool) {
	if o == nil || o.AllocateOnce == nil {
		return nil, false
	}
	return o.AllocateOnce, true
}

// HasAllocateOnce returns a boolean if a field has been set.
func (o *KoordinatorNodeReservationUpdate) HasAllocateOnce() bool {
	return o != nil && o.AllocateOnce != nil
}

// SetAllocateOnce gets a reference to the given bool and assigns it to the AllocateOnce field.
func (o *KoordinatorNodeReservationUpdate) SetAllocateOnce(v bool) {
	o.AllocateOnce = &v
}

// GetPreAllocation returns the PreAllocation field value if set, zero value otherwise.
func (o *KoordinatorNodeReservationUpdate) GetPreAllocation() bool {
	if o == nil || o.PreAllocation == nil {
		var ret bool
		return ret
	}
	return *o.PreAllocation
}

// GetPreAllocationOk returns a tuple with the PreAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservationUpdate) GetPreAllocationOk() (*bool, bool) {
	if o == nil || o.PreAllocation == nil {
		return nil, false
	}
	return o.PreAllocation, true
}

// HasPreAllocation returns a boolean if a field has been set.
func (o *KoordinatorNodeReservationUpdate) HasPreAllocation() bool {
	return o != nil && o.PreAllocation != nil
}

// SetPreAllocation gets a reference to the given bool and assigns it to the PreAllocation field.
func (o *KoordinatorNodeReservationUpdate) SetPreAllocation(v bool) {
	o.PreAllocation = &v
}

// GetAllocatePolicy returns the AllocatePolicy field value if set, zero value otherwise.
func (o *KoordinatorNodeReservationUpdate) GetAllocatePolicy() KoordinatorReservationAllocatePolicy {
	if o == nil || o.AllocatePolicy == nil {
		var ret KoordinatorReservationAllocatePolicy
		return ret
	}
	return *o.AllocatePolicy
}

// GetAllocatePolicyOk returns a tuple with the AllocatePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorNodeReservationUpdate) GetAllocatePolicyOk() (*KoordinatorReservationAllocatePolicy, bool) {
	if o == nil || o.AllocatePolicy == nil {
		return nil, false
	}
	return o.AllocatePolicy, true
}

// HasAllocatePolicy returns a boolean if a field has been set.
func (o *KoordinatorNodeReservationUpdate) HasAllocatePolicy() bool {
	return o != nil && o.AllocatePolicy != nil
}

// SetAllocatePolicy gets a reference to the given KoordinatorReservationAllocatePolicy and assigns it to the AllocatePolicy field.
func (o *KoordinatorNodeReservationUpdate) SetAllocatePolicy(v KoordinatorReservationAllocatePolicy) {
	o.AllocatePolicy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KoordinatorNodeReservationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KoordinatorNodeReservationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Resources      *KoordinatorReservationResources      `json:"resources"`
		Ttl            *string                               `json:"ttl,omitempty"`
		AllocateOnce   *bool                                 `json:"allocateOnce,omitempty"`
		PreAllocation  *bool                                 `json:"preAllocation,omitempty"`
		AllocatePolicy *KoordinatorReservationAllocatePolicy `json:"allocatePolicy,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Resources == nil {
		return fmt.Errorf("required field resources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"resources", "ttl", "allocateOnce", "preAllocation", "allocatePolicy"})
	} else {
		return err
	}

	hasInvalidField := false
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
