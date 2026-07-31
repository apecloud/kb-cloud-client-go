// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// KeyCluster a cluster which uses the key
type KeyCluster struct {
	// the id of the cluster
	ClusterId *string `json:"clusterId,omitempty"`
	// the name of the cluster
	Name *string `json:"name,omitempty"`
	// the cluster definition (engine type) of the cluster
	ClusterDefinition *string `json:"clusterDefinition,omitempty"`
	// the status of the cluster
	Status *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKeyCluster instantiates a new KeyCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKeyCluster() *KeyCluster {
	this := KeyCluster{}
	return &this
}

// NewKeyClusterWithDefaults instantiates a new KeyCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKeyClusterWithDefaults() *KeyCluster {
	this := KeyCluster{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *KeyCluster) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyCluster) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *KeyCluster) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *KeyCluster) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeyCluster) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyCluster) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeyCluster) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeyCluster) SetName(v string) {
	o.Name = &v
}

// GetClusterDefinition returns the ClusterDefinition field value if set, zero value otherwise.
func (o *KeyCluster) GetClusterDefinition() string {
	if o == nil || o.ClusterDefinition == nil {
		var ret string
		return ret
	}
	return *o.ClusterDefinition
}

// GetClusterDefinitionOk returns a tuple with the ClusterDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyCluster) GetClusterDefinitionOk() (*string, bool) {
	if o == nil || o.ClusterDefinition == nil {
		return nil, false
	}
	return o.ClusterDefinition, true
}

// HasClusterDefinition returns a boolean if a field has been set.
func (o *KeyCluster) HasClusterDefinition() bool {
	return o != nil && o.ClusterDefinition != nil
}

// SetClusterDefinition gets a reference to the given string and assigns it to the ClusterDefinition field.
func (o *KeyCluster) SetClusterDefinition(v string) {
	o.ClusterDefinition = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KeyCluster) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyCluster) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KeyCluster) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KeyCluster) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KeyCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ClusterDefinition != nil {
		toSerialize["clusterDefinition"] = o.ClusterDefinition
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
func (o *KeyCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterId         *string `json:"clusterId,omitempty"`
		Name              *string `json:"name,omitempty"`
		ClusterDefinition *string `json:"clusterDefinition,omitempty"`
		Status            *string `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterId", "name", "clusterDefinition", "status"})
	} else {
		return err
	}
	o.ClusterId = all.ClusterId
	o.Name = all.Name
	o.ClusterDefinition = all.ClusterDefinition
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
