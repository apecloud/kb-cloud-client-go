// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterInfo Cluster information
type ClusterInfo struct {
	// The name of the cluster
	ClusterName *string `json:"clusterName,omitempty"`
	// The definition of the cluster
	ClusterDefinition *string `json:"clusterDefinition,omitempty"`
	// The organization name
	OrgName *string `json:"orgName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterInfo instantiates a new ClusterInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterInfo() *ClusterInfo {
	this := ClusterInfo{}
	return &this
}

// NewClusterInfoWithDefaults instantiates a new ClusterInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterInfoWithDefaults() *ClusterInfo {
	this := ClusterInfo{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ClusterInfo) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ClusterInfo) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ClusterInfo) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterDefinition returns the ClusterDefinition field value if set, zero value otherwise.
func (o *ClusterInfo) GetClusterDefinition() string {
	if o == nil || o.ClusterDefinition == nil {
		var ret string
		return ret
	}
	return *o.ClusterDefinition
}

// GetClusterDefinitionOk returns a tuple with the ClusterDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetClusterDefinitionOk() (*string, bool) {
	if o == nil || o.ClusterDefinition == nil {
		return nil, false
	}
	return o.ClusterDefinition, true
}

// HasClusterDefinition returns a boolean if a field has been set.
func (o *ClusterInfo) HasClusterDefinition() bool {
	return o != nil && o.ClusterDefinition != nil
}

// SetClusterDefinition gets a reference to the given string and assigns it to the ClusterDefinition field.
func (o *ClusterInfo) SetClusterDefinition(v string) {
	o.ClusterDefinition = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *ClusterInfo) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInfo) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *ClusterInfo) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *ClusterInfo) SetOrgName(v string) {
	o.OrgName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.ClusterDefinition != nil {
		toSerialize["clusterDefinition"] = o.ClusterDefinition
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterName       *string `json:"clusterName,omitempty"`
		ClusterDefinition *string `json:"clusterDefinition,omitempty"`
		OrgName           *string `json:"orgName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterName", "clusterDefinition", "orgName"})
	} else {
		return err
	}
	o.ClusterName = all.ClusterName
	o.ClusterDefinition = all.ClusterDefinition
	o.OrgName = all.OrgName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
