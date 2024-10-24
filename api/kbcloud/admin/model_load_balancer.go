// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancer The load balancer info
// NODESCRIPTION LoadBalancer
//
// Deprecated: This model is deprecated.
type LoadBalancer struct {
	// Version of the load balancer
	Version *string `json:"version,omitempty"`
	// Type of the load balancer
	Type *string `json:"type,omitempty"`
	// Status of the load balancer
	Status LoadBalancerStatus `json:"status"`
	// Whether the loadbalancer is available in the environment.
	Available *LoadBalancerAvailableType `json:"available,omitempty"`
	// Class of the load balancer, only available on public cloud.
	Class *string `json:"class,omitempty"`
	// Charge type of the load balancer, only available on public cloud.
	ChargeType *string `json:"chargeType,omitempty"`
	// Whether the ipam is enabled.
	Ipam *LoadBalancerIpamStatus `json:"ipam,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLoadBalancer instantiates a new LoadBalancer object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLoadBalancer(status LoadBalancerStatus) *LoadBalancer {
	this := LoadBalancer{}
	this.Status = status
	return &this
}

// NewLoadBalancerWithDefaults instantiates a new LoadBalancer object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLoadBalancerWithDefaults() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LoadBalancer) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LoadBalancer) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LoadBalancer) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LoadBalancer) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LoadBalancer) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LoadBalancer) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value.
func (o *LoadBalancer) GetStatus() LoadBalancerStatus {
	if o == nil {
		var ret LoadBalancerStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetStatusOk() (*LoadBalancerStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *LoadBalancer) SetStatus(v LoadBalancerStatus) {
	o.Status = v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *LoadBalancer) GetAvailable() LoadBalancerAvailableType {
	if o == nil || o.Available == nil {
		var ret LoadBalancerAvailableType
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetAvailableOk() (*LoadBalancerAvailableType, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *LoadBalancer) HasAvailable() bool {
	return o != nil && o.Available != nil
}

// SetAvailable gets a reference to the given LoadBalancerAvailableType and assigns it to the Available field.
func (o *LoadBalancer) SetAvailable(v LoadBalancerAvailableType) {
	o.Available = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *LoadBalancer) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetClassOk() (*string, bool) {
	if o == nil || o.Class == nil {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *LoadBalancer) HasClass() bool {
	return o != nil && o.Class != nil
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *LoadBalancer) SetClass(v string) {
	o.Class = &v
}

// GetChargeType returns the ChargeType field value if set, zero value otherwise.
func (o *LoadBalancer) GetChargeType() string {
	if o == nil || o.ChargeType == nil {
		var ret string
		return ret
	}
	return *o.ChargeType
}

// GetChargeTypeOk returns a tuple with the ChargeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetChargeTypeOk() (*string, bool) {
	if o == nil || o.ChargeType == nil {
		return nil, false
	}
	return o.ChargeType, true
}

// HasChargeType returns a boolean if a field has been set.
func (o *LoadBalancer) HasChargeType() bool {
	return o != nil && o.ChargeType != nil
}

// SetChargeType gets a reference to the given string and assigns it to the ChargeType field.
func (o *LoadBalancer) SetChargeType(v string) {
	o.ChargeType = &v
}

// GetIpam returns the Ipam field value if set, zero value otherwise.
func (o *LoadBalancer) GetIpam() LoadBalancerIpamStatus {
	if o == nil || o.Ipam == nil {
		var ret LoadBalancerIpamStatus
		return ret
	}
	return *o.Ipam
}

// GetIpamOk returns a tuple with the Ipam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetIpamOk() (*LoadBalancerIpamStatus, bool) {
	if o == nil || o.Ipam == nil {
		return nil, false
	}
	return o.Ipam, true
}

// HasIpam returns a boolean if a field has been set.
func (o *LoadBalancer) HasIpam() bool {
	return o != nil && o.Ipam != nil
}

// SetIpam gets a reference to the given LoadBalancerIpamStatus and assigns it to the Ipam field.
func (o *LoadBalancer) SetIpam(v LoadBalancerIpamStatus) {
	o.Ipam = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["status"] = o.Status
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	if o.Class != nil {
		toSerialize["class"] = o.Class
	}
	if o.ChargeType != nil {
		toSerialize["chargeType"] = o.ChargeType
	}
	if o.Ipam != nil {
		toSerialize["ipam"] = o.Ipam
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LoadBalancer) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Version    *string                    `json:"version,omitempty"`
		Type       *string                    `json:"type,omitempty"`
		Status     *LoadBalancerStatus        `json:"status"`
		Available  *LoadBalancerAvailableType `json:"available,omitempty"`
		Class      *string                    `json:"class,omitempty"`
		ChargeType *string                    `json:"chargeType,omitempty"`
		Ipam       *LoadBalancerIpamStatus    `json:"ipam,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"version", "type", "status", "available", "class", "chargeType", "ipam"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Version = all.Version
	o.Type = all.Type
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.Available != nil && !all.Available.IsValid() {
		hasInvalidField = true
	} else {
		o.Available = all.Available
	}
	o.Class = all.Class
	o.ChargeType = all.ChargeType
	if all.Ipam != nil && !all.Ipam.IsValid() {
		hasInvalidField = true
	} else {
		o.Ipam = all.Ipam
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
