// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type StaticCluster struct {
	// Optional component overrides. Unknown, duplicate or other-engine components are rejected. Must be omitted when reusedClusterName is set; adding an environment does not resize an existing cluster.
	Components []StaticClusterComponent `json:"components,omitempty"`
	// Kubernetes name of an existing cluster to attach. Set together with reusedClusterNamespace. The existing cluster's actual resources are used, and component overrides are rejected.
	ReusedClusterName *string `json:"reusedClusterName,omitempty"`
	// Namespace of the cluster named by reusedClusterName. Required when reusedClusterName is set.
	ReusedClusterNamespace *string `json:"reusedClusterNamespace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStaticCluster instantiates a new StaticCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStaticCluster() *StaticCluster {
	this := StaticCluster{}
	return &this
}

// NewStaticClusterWithDefaults instantiates a new StaticCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStaticClusterWithDefaults() *StaticCluster {
	this := StaticCluster{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *StaticCluster) GetComponents() []StaticClusterComponent {
	if o == nil || o.Components == nil {
		var ret []StaticClusterComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCluster) GetComponentsOk() (*[]StaticClusterComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *StaticCluster) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []StaticClusterComponent and assigns it to the Components field.
func (o *StaticCluster) SetComponents(v []StaticClusterComponent) {
	o.Components = v
}

// GetReusedClusterName returns the ReusedClusterName field value if set, zero value otherwise.
func (o *StaticCluster) GetReusedClusterName() string {
	if o == nil || o.ReusedClusterName == nil {
		var ret string
		return ret
	}
	return *o.ReusedClusterName
}

// GetReusedClusterNameOk returns a tuple with the ReusedClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCluster) GetReusedClusterNameOk() (*string, bool) {
	if o == nil || o.ReusedClusterName == nil {
		return nil, false
	}
	return o.ReusedClusterName, true
}

// HasReusedClusterName returns a boolean if a field has been set.
func (o *StaticCluster) HasReusedClusterName() bool {
	return o != nil && o.ReusedClusterName != nil
}

// SetReusedClusterName gets a reference to the given string and assigns it to the ReusedClusterName field.
func (o *StaticCluster) SetReusedClusterName(v string) {
	o.ReusedClusterName = &v
}

// GetReusedClusterNamespace returns the ReusedClusterNamespace field value if set, zero value otherwise.
func (o *StaticCluster) GetReusedClusterNamespace() string {
	if o == nil || o.ReusedClusterNamespace == nil {
		var ret string
		return ret
	}
	return *o.ReusedClusterNamespace
}

// GetReusedClusterNamespaceOk returns a tuple with the ReusedClusterNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCluster) GetReusedClusterNamespaceOk() (*string, bool) {
	if o == nil || o.ReusedClusterNamespace == nil {
		return nil, false
	}
	return o.ReusedClusterNamespace, true
}

// HasReusedClusterNamespace returns a boolean if a field has been set.
func (o *StaticCluster) HasReusedClusterNamespace() bool {
	return o != nil && o.ReusedClusterNamespace != nil
}

// SetReusedClusterNamespace gets a reference to the given string and assigns it to the ReusedClusterNamespace field.
func (o *StaticCluster) SetReusedClusterNamespace(v string) {
	o.ReusedClusterNamespace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StaticCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.ReusedClusterName != nil {
		toSerialize["reusedClusterName"] = o.ReusedClusterName
	}
	if o.ReusedClusterNamespace != nil {
		toSerialize["reusedClusterNamespace"] = o.ReusedClusterNamespace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StaticCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components             []StaticClusterComponent `json:"components,omitempty"`
		ReusedClusterName      *string                  `json:"reusedClusterName,omitempty"`
		ReusedClusterNamespace *string                  `json:"reusedClusterNamespace,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"components", "reusedClusterName", "reusedClusterNamespace"})
	} else {
		return err
	}
	o.Components = all.Components
	o.ReusedClusterName = all.ReusedClusterName
	o.ReusedClusterNamespace = all.ReusedClusterNamespace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
