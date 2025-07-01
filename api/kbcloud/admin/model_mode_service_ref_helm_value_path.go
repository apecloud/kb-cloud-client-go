// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeServiceRefHelmValuePath separate values with commas.
type ModeServiceRefHelmValuePath struct {
	// the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object.
	Namespace string `json:"namespace"`
	// the name of the referenced Cluster. Other options (like serivce or port name) are omitted
	// because it is assumed to be defined in the cluster chart.
	//
	Cluster *string `json:"cluster,omitempty"`
	// the name of the referenced ServiceDescriptor
	ServiceDescriptor *string `json:"serviceDescriptor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeServiceRefHelmValuePath instantiates a new ModeServiceRefHelmValuePath object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeServiceRefHelmValuePath(namespace string) *ModeServiceRefHelmValuePath {
	this := ModeServiceRefHelmValuePath{}
	this.Namespace = namespace
	return &this
}

// NewModeServiceRefHelmValuePathWithDefaults instantiates a new ModeServiceRefHelmValuePath object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeServiceRefHelmValuePathWithDefaults() *ModeServiceRefHelmValuePath {
	this := ModeServiceRefHelmValuePath{}
	return &this
}

// GetNamespace returns the Namespace field value.
func (o *ModeServiceRefHelmValuePath) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *ModeServiceRefHelmValuePath) SetNamespace(v string) {
	o.Namespace = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ModeServiceRefHelmValuePath) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ModeServiceRefHelmValuePath) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *ModeServiceRefHelmValuePath) SetCluster(v string) {
	o.Cluster = &v
}

// GetServiceDescriptor returns the ServiceDescriptor field value if set, zero value otherwise.
func (o *ModeServiceRefHelmValuePath) GetServiceDescriptor() string {
	if o == nil || o.ServiceDescriptor == nil {
		var ret string
		return ret
	}
	return *o.ServiceDescriptor
}

// GetServiceDescriptorOk returns a tuple with the ServiceDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetServiceDescriptorOk() (*string, bool) {
	if o == nil || o.ServiceDescriptor == nil {
		return nil, false
	}
	return o.ServiceDescriptor, true
}

// HasServiceDescriptor returns a boolean if a field has been set.
func (o *ModeServiceRefHelmValuePath) HasServiceDescriptor() bool {
	return o != nil && o.ServiceDescriptor != nil
}

// SetServiceDescriptor gets a reference to the given string and assigns it to the ServiceDescriptor field.
func (o *ModeServiceRefHelmValuePath) SetServiceDescriptor(v string) {
	o.ServiceDescriptor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeServiceRefHelmValuePath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
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
func (o *ModeServiceRefHelmValuePath) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Namespace         *string `json:"namespace"`
		Cluster           *string `json:"cluster,omitempty"`
		ServiceDescriptor *string `json:"serviceDescriptor,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"namespace", "cluster", "serviceDescriptor"})
	} else {
		return err
	}
	o.Namespace = *all.Namespace
	o.Cluster = all.Cluster
	o.ServiceDescriptor = all.ServiceDescriptor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
