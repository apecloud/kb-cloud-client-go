// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsExpose OpsExpose is the payload to expose a KubeBlocks cluster
type OpsExpose struct {
	Component string `json:"component"`
	Enable    bool   `json:"enable"`
	Readonly  *bool  `json:"readonly,omitempty"`
	// Specifies the type of exposure for the KubeBlocks cluster.
	Type OpsExposeType `json:"type"`
	// Specifies the type of service for the KubeBlocks cluster.
	VpcServiceType *OpsExposeVPCServiceType    `json:"vpcServiceType,omitempty"`
	PortsMapping   []OpsExposePortsMappingItem `json:"portsMapping,omitempty"`
	// The IP address of the LoadBalancer service. If not set, the IP address will be assigned by the system. Only available when vpcServiceType is LoadBalancer.
	LoadBalancerIp common.NullableString `json:"loadBalancerIP,omitempty"`
	// The IP pool ID of the LoadBalancer service. If not set, the IP pool will be assigned by the system. Only available when vpcServiceType is LoadBalancer.
	LoadBalancerIpPoolId common.NullableString `json:"loadBalancerIPPoolID,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsExpose instantiates a new OpsExpose object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsExpose(component string, enable bool, typeVar OpsExposeType) *OpsExpose {
	this := OpsExpose{}
	this.Component = component
	this.Enable = enable
	this.Type = typeVar
	var vpcServiceType OpsExposeVPCServiceType = OpsExposeVPCServiceTypeLoadBalancer
	this.VpcServiceType = &vpcServiceType
	return &this
}

// NewOpsExposeWithDefaults instantiates a new OpsExpose object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsExposeWithDefaults() *OpsExpose {
	this := OpsExpose{}
	var vpcServiceType OpsExposeVPCServiceType = OpsExposeVPCServiceTypeLoadBalancer
	this.VpcServiceType = &vpcServiceType
	return &this
}

// GetComponent returns the Component field value.
func (o *OpsExpose) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OpsExpose) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *OpsExpose) SetComponent(v string) {
	o.Component = v
}

// GetEnable returns the Enable field value.
func (o *OpsExpose) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *OpsExpose) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value.
func (o *OpsExpose) SetEnable(v bool) {
	o.Enable = v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *OpsExpose) GetReadonly() bool {
	if o == nil || o.Readonly == nil {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsExpose) GetReadonlyOk() (*bool, bool) {
	if o == nil || o.Readonly == nil {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *OpsExpose) HasReadonly() bool {
	return o != nil && o.Readonly != nil
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *OpsExpose) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetType returns the Type field value.
func (o *OpsExpose) GetType() OpsExposeType {
	if o == nil {
		var ret OpsExposeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OpsExpose) GetTypeOk() (*OpsExposeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OpsExpose) SetType(v OpsExposeType) {
	o.Type = v
}

// GetVpcServiceType returns the VpcServiceType field value if set, zero value otherwise.
func (o *OpsExpose) GetVpcServiceType() OpsExposeVPCServiceType {
	if o == nil || o.VpcServiceType == nil {
		var ret OpsExposeVPCServiceType
		return ret
	}
	return *o.VpcServiceType
}

// GetVpcServiceTypeOk returns a tuple with the VpcServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsExpose) GetVpcServiceTypeOk() (*OpsExposeVPCServiceType, bool) {
	if o == nil || o.VpcServiceType == nil {
		return nil, false
	}
	return o.VpcServiceType, true
}

// HasVpcServiceType returns a boolean if a field has been set.
func (o *OpsExpose) HasVpcServiceType() bool {
	return o != nil && o.VpcServiceType != nil
}

// SetVpcServiceType gets a reference to the given OpsExposeVPCServiceType and assigns it to the VpcServiceType field.
func (o *OpsExpose) SetVpcServiceType(v OpsExposeVPCServiceType) {
	o.VpcServiceType = &v
}

// GetPortsMapping returns the PortsMapping field value if set, zero value otherwise.
func (o *OpsExpose) GetPortsMapping() []OpsExposePortsMappingItem {
	if o == nil || o.PortsMapping == nil {
		var ret []OpsExposePortsMappingItem
		return ret
	}
	return o.PortsMapping
}

// GetPortsMappingOk returns a tuple with the PortsMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsExpose) GetPortsMappingOk() (*[]OpsExposePortsMappingItem, bool) {
	if o == nil || o.PortsMapping == nil {
		return nil, false
	}
	return &o.PortsMapping, true
}

// HasPortsMapping returns a boolean if a field has been set.
func (o *OpsExpose) HasPortsMapping() bool {
	return o != nil && o.PortsMapping != nil
}

// SetPortsMapping gets a reference to the given []OpsExposePortsMappingItem and assigns it to the PortsMapping field.
func (o *OpsExpose) SetPortsMapping(v []OpsExposePortsMappingItem) {
	o.PortsMapping = v
}

// GetLoadBalancerIp returns the LoadBalancerIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsExpose) GetLoadBalancerIp() string {
	if o == nil || o.LoadBalancerIp.Get() == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerIp.Get()
}

// GetLoadBalancerIpOk returns a tuple with the LoadBalancerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsExpose) GetLoadBalancerIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoadBalancerIp.Get(), o.LoadBalancerIp.IsSet()
}

// HasLoadBalancerIp returns a boolean if a field has been set.
func (o *OpsExpose) HasLoadBalancerIp() bool {
	return o != nil && o.LoadBalancerIp.IsSet()
}

// SetLoadBalancerIp gets a reference to the given common.NullableString and assigns it to the LoadBalancerIp field.
func (o *OpsExpose) SetLoadBalancerIp(v string) {
	o.LoadBalancerIp.Set(&v)
}

// SetLoadBalancerIpNil sets the value for LoadBalancerIp to be an explicit nil.
func (o *OpsExpose) SetLoadBalancerIpNil() {
	o.LoadBalancerIp.Set(nil)
}

// UnsetLoadBalancerIp ensures that no value is present for LoadBalancerIp, not even an explicit nil.
func (o *OpsExpose) UnsetLoadBalancerIp() {
	o.LoadBalancerIp.Unset()
}

// GetLoadBalancerIpPoolId returns the LoadBalancerIpPoolId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsExpose) GetLoadBalancerIpPoolId() string {
	if o == nil || o.LoadBalancerIpPoolId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerIpPoolId.Get()
}

// GetLoadBalancerIpPoolIdOk returns a tuple with the LoadBalancerIpPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsExpose) GetLoadBalancerIpPoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoadBalancerIpPoolId.Get(), o.LoadBalancerIpPoolId.IsSet()
}

// HasLoadBalancerIpPoolId returns a boolean if a field has been set.
func (o *OpsExpose) HasLoadBalancerIpPoolId() bool {
	return o != nil && o.LoadBalancerIpPoolId.IsSet()
}

// SetLoadBalancerIpPoolId gets a reference to the given common.NullableString and assigns it to the LoadBalancerIpPoolId field.
func (o *OpsExpose) SetLoadBalancerIpPoolId(v string) {
	o.LoadBalancerIpPoolId.Set(&v)
}

// SetLoadBalancerIpPoolIdNil sets the value for LoadBalancerIpPoolId to be an explicit nil.
func (o *OpsExpose) SetLoadBalancerIpPoolIdNil() {
	o.LoadBalancerIpPoolId.Set(nil)
}

// UnsetLoadBalancerIpPoolId ensures that no value is present for LoadBalancerIpPoolId, not even an explicit nil.
func (o *OpsExpose) UnsetLoadBalancerIpPoolId() {
	o.LoadBalancerIpPoolId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsExpose) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["enable"] = o.Enable
	if o.Readonly != nil {
		toSerialize["readonly"] = o.Readonly
	}
	toSerialize["type"] = o.Type
	if o.VpcServiceType != nil {
		toSerialize["vpcServiceType"] = o.VpcServiceType
	}
	if o.PortsMapping != nil {
		toSerialize["portsMapping"] = o.PortsMapping
	}
	if o.LoadBalancerIp.IsSet() {
		toSerialize["loadBalancerIP"] = o.LoadBalancerIp.Get()
	}
	if o.LoadBalancerIpPoolId.IsSet() {
		toSerialize["loadBalancerIPPoolID"] = o.LoadBalancerIpPoolId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsExpose) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component            *string                     `json:"component"`
		Enable               *bool                       `json:"enable"`
		Readonly             *bool                       `json:"readonly,omitempty"`
		Type                 *OpsExposeType              `json:"type"`
		VpcServiceType       *OpsExposeVPCServiceType    `json:"vpcServiceType,omitempty"`
		PortsMapping         []OpsExposePortsMappingItem `json:"portsMapping,omitempty"`
		LoadBalancerIp       common.NullableString       `json:"loadBalancerIP,omitempty"`
		LoadBalancerIpPoolId common.NullableString       `json:"loadBalancerIPPoolID,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Enable == nil {
		return fmt.Errorf("required field enable missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "enable", "readonly", "type", "vpcServiceType", "portsMapping", "loadBalancerIP", "loadBalancerIPPoolID"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = *all.Component
	o.Enable = *all.Enable
	o.Readonly = all.Readonly
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.VpcServiceType != nil && !all.VpcServiceType.IsValid() {
		hasInvalidField = true
	} else {
		o.VpcServiceType = all.VpcServiceType
	}
	o.PortsMapping = all.PortsMapping
	o.LoadBalancerIp = all.LoadBalancerIp
	o.LoadBalancerIpPoolId = all.LoadBalancerIpPoolId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
