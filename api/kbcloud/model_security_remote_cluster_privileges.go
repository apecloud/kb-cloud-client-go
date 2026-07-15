// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SecurityRemoteClusterPrivileges struct {
	Clusters   []string `json:"clusters"`
	Privileges []string `json:"privileges"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityRemoteClusterPrivileges instantiates a new SecurityRemoteClusterPrivileges object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityRemoteClusterPrivileges(clusters []string, privileges []string) *SecurityRemoteClusterPrivileges {
	this := SecurityRemoteClusterPrivileges{}
	this.Clusters = clusters
	this.Privileges = privileges
	return &this
}

// NewSecurityRemoteClusterPrivilegesWithDefaults instantiates a new SecurityRemoteClusterPrivileges object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityRemoteClusterPrivilegesWithDefaults() *SecurityRemoteClusterPrivileges {
	this := SecurityRemoteClusterPrivileges{}
	return &this
}

// GetClusters returns the Clusters field value.
func (o *SecurityRemoteClusterPrivileges) GetClusters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *SecurityRemoteClusterPrivileges) GetClustersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// SetClusters sets field value.
func (o *SecurityRemoteClusterPrivileges) SetClusters(v []string) {
	o.Clusters = v
}

// GetPrivileges returns the Privileges field value.
func (o *SecurityRemoteClusterPrivileges) GetPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *SecurityRemoteClusterPrivileges) GetPrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *SecurityRemoteClusterPrivileges) SetPrivileges(v []string) {
	o.Privileges = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityRemoteClusterPrivileges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["clusters"] = o.Clusters
	toSerialize["privileges"] = o.Privileges

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityRemoteClusterPrivileges) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Clusters   *[]string `json:"clusters"`
		Privileges *[]string `json:"privileges"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Clusters == nil {
		return fmt.Errorf("required field clusters missing")
	}
	if all.Privileges == nil {
		return fmt.Errorf("required field privileges missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusters", "privileges"})
	} else {
		return err
	}
	o.Clusters = *all.Clusters
	o.Privileges = *all.Privileges

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
