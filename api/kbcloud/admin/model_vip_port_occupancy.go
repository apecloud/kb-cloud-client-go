// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VipPortOccupancy A port already used on the VIP
type VipPortOccupancy struct {
	// Name of the cluster owning the service
	Cluster *string `json:"cluster,omitempty"`
	// Namespace of the service
	Namespace *string `json:"namespace,omitempty"`
	// Port number
	Port int32 `json:"port"`
	// Port protocol
	Protocol string `json:"protocol"`
	// Name of the service
	Service *string `json:"service,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVipPortOccupancy instantiates a new VipPortOccupancy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVipPortOccupancy(port int32, protocol string) *VipPortOccupancy {
	this := VipPortOccupancy{}
	this.Port = port
	this.Protocol = protocol
	return &this
}

// NewVipPortOccupancyWithDefaults instantiates a new VipPortOccupancy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVipPortOccupancyWithDefaults() *VipPortOccupancy {
	this := VipPortOccupancy{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VipPortOccupancy) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipPortOccupancy) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VipPortOccupancy) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *VipPortOccupancy) SetCluster(v string) {
	o.Cluster = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *VipPortOccupancy) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipPortOccupancy) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *VipPortOccupancy) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *VipPortOccupancy) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPort returns the Port field value.
func (o *VipPortOccupancy) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *VipPortOccupancy) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value.
func (o *VipPortOccupancy) SetPort(v int32) {
	o.Port = v
}

// GetProtocol returns the Protocol field value.
func (o *VipPortOccupancy) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *VipPortOccupancy) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value.
func (o *VipPortOccupancy) SetProtocol(v string) {
	o.Protocol = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *VipPortOccupancy) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipPortOccupancy) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *VipPortOccupancy) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *VipPortOccupancy) SetService(v string) {
	o.Service = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o VipPortOccupancy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	toSerialize["port"] = o.Port
	toSerialize["protocol"] = o.Protocol
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VipPortOccupancy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster   *string `json:"cluster,omitempty"`
		Namespace *string `json:"namespace,omitempty"`
		Port      *int32  `json:"port"`
		Protocol  *string `json:"protocol"`
		Service   *string `json:"service,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Port == nil {
		return fmt.Errorf("required field port missing")
	}
	if all.Protocol == nil {
		return fmt.Errorf("required field protocol missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster", "namespace", "port", "protocol", "service"})
	} else {
		return err
	}
	o.Cluster = all.Cluster
	o.Namespace = all.Namespace
	o.Port = *all.Port
	o.Protocol = *all.Protocol
	o.Service = all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
