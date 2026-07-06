// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type KoordinatorReservationSummaryItem struct {
	ResourceClass *KoordinatorReservationResourceClass `json:"resourceClass,omitempty"`
	NodeCount     *int64                               `json:"nodeCount,omitempty"`
	// Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
	Total *KoordinatorReservationResources `json:"total,omitempty"`
	// Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
	Allocated *KoordinatorReservationResources `json:"allocated,omitempty"`
	// Resource requests supported by KBE for Koordinator Reservations. Values are Kubernetes quantities and responses use Kubernetes canonical quantity strings.
	Available *KoordinatorReservationResources `json:"available,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKoordinatorReservationSummaryItem instantiates a new KoordinatorReservationSummaryItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKoordinatorReservationSummaryItem() *KoordinatorReservationSummaryItem {
	this := KoordinatorReservationSummaryItem{}
	return &this
}

// NewKoordinatorReservationSummaryItemWithDefaults instantiates a new KoordinatorReservationSummaryItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKoordinatorReservationSummaryItemWithDefaults() *KoordinatorReservationSummaryItem {
	this := KoordinatorReservationSummaryItem{}
	return &this
}

// GetResourceClass returns the ResourceClass field value if set, zero value otherwise.
func (o *KoordinatorReservationSummaryItem) GetResourceClass() KoordinatorReservationResourceClass {
	if o == nil || o.ResourceClass == nil {
		var ret KoordinatorReservationResourceClass
		return ret
	}
	return *o.ResourceClass
}

// GetResourceClassOk returns a tuple with the ResourceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummaryItem) GetResourceClassOk() (*KoordinatorReservationResourceClass, bool) {
	if o == nil || o.ResourceClass == nil {
		return nil, false
	}
	return o.ResourceClass, true
}

// HasResourceClass returns a boolean if a field has been set.
func (o *KoordinatorReservationSummaryItem) HasResourceClass() bool {
	return o != nil && o.ResourceClass != nil
}

// SetResourceClass gets a reference to the given KoordinatorReservationResourceClass and assigns it to the ResourceClass field.
func (o *KoordinatorReservationSummaryItem) SetResourceClass(v KoordinatorReservationResourceClass) {
	o.ResourceClass = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *KoordinatorReservationSummaryItem) GetNodeCount() int64 {
	if o == nil || o.NodeCount == nil {
		var ret int64
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummaryItem) GetNodeCountOk() (*int64, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *KoordinatorReservationSummaryItem) HasNodeCount() bool {
	return o != nil && o.NodeCount != nil
}

// SetNodeCount gets a reference to the given int64 and assigns it to the NodeCount field.
func (o *KoordinatorReservationSummaryItem) SetNodeCount(v int64) {
	o.NodeCount = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *KoordinatorReservationSummaryItem) GetTotal() KoordinatorReservationResources {
	if o == nil || o.Total == nil {
		var ret KoordinatorReservationResources
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummaryItem) GetTotalOk() (*KoordinatorReservationResources, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *KoordinatorReservationSummaryItem) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given KoordinatorReservationResources and assigns it to the Total field.
func (o *KoordinatorReservationSummaryItem) SetTotal(v KoordinatorReservationResources) {
	o.Total = &v
}

// GetAllocated returns the Allocated field value if set, zero value otherwise.
func (o *KoordinatorReservationSummaryItem) GetAllocated() KoordinatorReservationResources {
	if o == nil || o.Allocated == nil {
		var ret KoordinatorReservationResources
		return ret
	}
	return *o.Allocated
}

// GetAllocatedOk returns a tuple with the Allocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummaryItem) GetAllocatedOk() (*KoordinatorReservationResources, bool) {
	if o == nil || o.Allocated == nil {
		return nil, false
	}
	return o.Allocated, true
}

// HasAllocated returns a boolean if a field has been set.
func (o *KoordinatorReservationSummaryItem) HasAllocated() bool {
	return o != nil && o.Allocated != nil
}

// SetAllocated gets a reference to the given KoordinatorReservationResources and assigns it to the Allocated field.
func (o *KoordinatorReservationSummaryItem) SetAllocated(v KoordinatorReservationResources) {
	o.Allocated = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *KoordinatorReservationSummaryItem) GetAvailable() KoordinatorReservationResources {
	if o == nil || o.Available == nil {
		var ret KoordinatorReservationResources
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KoordinatorReservationSummaryItem) GetAvailableOk() (*KoordinatorReservationResources, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *KoordinatorReservationSummaryItem) HasAvailable() bool {
	return o != nil && o.Available != nil
}

// SetAvailable gets a reference to the given KoordinatorReservationResources and assigns it to the Available field.
func (o *KoordinatorReservationSummaryItem) SetAvailable(v KoordinatorReservationResources) {
	o.Available = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KoordinatorReservationSummaryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ResourceClass != nil {
		toSerialize["resourceClass"] = o.ResourceClass
	}
	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
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
func (o *KoordinatorReservationSummaryItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceClass *KoordinatorReservationResourceClass `json:"resourceClass,omitempty"`
		NodeCount     *int64                               `json:"nodeCount,omitempty"`
		Total         *KoordinatorReservationResources     `json:"total,omitempty"`
		Allocated     *KoordinatorReservationResources     `json:"allocated,omitempty"`
		Available     *KoordinatorReservationResources     `json:"available,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"resourceClass", "nodeCount", "total", "allocated", "available"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ResourceClass != nil && !all.ResourceClass.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceClass = all.ResourceClass
	}
	o.NodeCount = all.NodeCount
	if all.Total != nil && all.Total.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Total = all.Total
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
