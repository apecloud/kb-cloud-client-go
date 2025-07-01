// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ServiceRef defines a serviceRef that references service provided by other cluster or external service
type ServiceRef struct {
	// refers to the serviceRef name defined in engineoption
	Name string `json:"name"`
	// refer to engineoption for more information
	Namespace string `json:"namespace"`
	// refer to engineoption for more information. Only one of cluster or serviceDescriptor should
	// be provided.
	//
	Cluster *string `json:"cluster,omitempty"`
	// refer to engineoption for more information. Only one of cluster or serviceDescriptor should
	// be provided.
	//
	ServiceDescriptor *string `json:"serviceDescriptor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceRef instantiates a new ServiceRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceRef(name string, namespace string) *ServiceRef {
	this := ServiceRef{}
	this.Name = name
	this.Namespace = namespace
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

// GetNamespace returns the Namespace field value.
func (o *ServiceRef) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ServiceRef) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *ServiceRef) SetNamespace(v string) {
	o.Namespace = v
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
func (o *ServiceRef) GetServiceDescriptor() string {
	if o == nil || o.ServiceDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ServiceDescriptor
}

// GetServiceDescriptorOk returns a tuple with the ServiceDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRef) GetServiceDescriptorOk() (*string, bool) {
	if o == nil || o.ServiceDescriptor == nil {
		return nil, false
	}
	return o.ServiceDescriptor, true
}

// HasServiceDescriptor returns a boolean if a field has been set.
func (o *ServiceRef) HasServiceDescriptor() bool {
	return o != nil && o.ServiceDescriptor != nil
}

// SetServiceDescriptor gets a reference to the given string and assigns it to the ServiceDescriptor field.
func (o *ServiceRef) SetServiceDescriptor(v string) {
	o.ServiceDescriptor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
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
		Name              *string `json:"name"`
		Namespace         *string `json:"namespace"`
		Cluster           *string `json:"cluster,omitempty"`
		ServiceDescriptor *string `json:"serviceDescriptor,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "namespace", "cluster", "serviceDescriptor"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Namespace = *all.Namespace
	o.Cluster = all.Cluster
	o.ServiceDescriptor = all.ServiceDescriptor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
