// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ServiceRef defines a serviceRef that references service provided by other cluster or external service.
// Only one of cluster or serviceDescriptor field should be set.
type ServiceRef struct {
	// refers to the serviceRef name defined in engineoption
	Name string `json:"name"`
	// the cluster name that will be used in serviceRef. The referenced cluster should
	// be at the same organization as the current cluster.
	//
	Cluster *string `json:"cluster,omitempty"`
	// serviceDescriptor that will be used in serviceRef. The field definition is in line with kubeblocks.
	ServiceDescriptor *ServiceDescriptor `json:"serviceDescriptor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceRef instantiates a new ServiceRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceRef(name string) *ServiceRef {
	this := ServiceRef{}
	this.Name = name
	return &this
}

// NewServiceRefWithDefaults instantiates a new ServiceRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceRefWithDefaults() *ServiceRef {
	this := ServiceRef{}
	return &this
}

// GetName returns the Name field value.
func (o *ServiceRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ServiceRef) SetName(v string) {
	o.Name = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ServiceRef) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRef) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ServiceRef) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *ServiceRef) SetCluster(v string) {
	o.Cluster = &v
}

// GetServiceDescriptor returns the ServiceDescriptor field value if set, zero value otherwise.
func (o *ServiceRef) GetServiceDescriptor() ServiceDescriptor {
	if o == nil || o.ServiceDescriptor == nil {
		var ret ServiceDescriptor
		return ret
	}
	return *o.ServiceDescriptor
}

// GetServiceDescriptorOk returns a tuple with the ServiceDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRef) GetServiceDescriptorOk() (*ServiceDescriptor, bool) {
	if o == nil || o.ServiceDescriptor == nil {
		return nil, false
	}
	return o.ServiceDescriptor, true
}

// HasServiceDescriptor returns a boolean if a field has been set.
func (o *ServiceRef) HasServiceDescriptor() bool {
	return o != nil && o.ServiceDescriptor != nil
}

// SetServiceDescriptor gets a reference to the given ServiceDescriptor and assigns it to the ServiceDescriptor field.
func (o *ServiceRef) SetServiceDescriptor(v ServiceDescriptor) {
	o.ServiceDescriptor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.ServiceDescriptor != nil {
		toSerialize["serviceDescriptor"] = o.ServiceDescriptor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string            `json:"name"`
		Cluster           *string            `json:"cluster,omitempty"`
		ServiceDescriptor *ServiceDescriptor `json:"serviceDescriptor,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "cluster", "serviceDescriptor"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Cluster = all.Cluster
	if all.ServiceDescriptor != nil && all.ServiceDescriptor.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ServiceDescriptor = all.ServiceDescriptor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
