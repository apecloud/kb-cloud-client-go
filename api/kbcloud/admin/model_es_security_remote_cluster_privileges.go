// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ESSecurityRemoteClusterPrivileges struct {
	Clusters   []string `json:"clusters"`
	Privileges []string `json:"privileges"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewESSecurityRemoteClusterPrivileges instantiates a new ESSecurityRemoteClusterPrivileges object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewESSecurityRemoteClusterPrivileges(clusters []string, privileges []string) *ESSecurityRemoteClusterPrivileges {
	this := ESSecurityRemoteClusterPrivileges{}
	this.Clusters = clusters
	this.Privileges = privileges
	return &this
}

// NewESSecurityRemoteClusterPrivilegesWithDefaults instantiates a new ESSecurityRemoteClusterPrivileges object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewESSecurityRemoteClusterPrivilegesWithDefaults() *ESSecurityRemoteClusterPrivileges {
	this := ESSecurityRemoteClusterPrivileges{}
	return &this
}

// GetClusters returns the Clusters field value.
func (o *ESSecurityRemoteClusterPrivileges) GetClusters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *ESSecurityRemoteClusterPrivileges) GetClustersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// SetClusters sets field value.
func (o *ESSecurityRemoteClusterPrivileges) SetClusters(v []string) {
	o.Clusters = v
}

// GetPrivileges returns the Privileges field value.
func (o *ESSecurityRemoteClusterPrivileges) GetPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *ESSecurityRemoteClusterPrivileges) GetPrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *ESSecurityRemoteClusterPrivileges) SetPrivileges(v []string) {
	o.Privileges = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ESSecurityRemoteClusterPrivileges) MarshalJSON() ([]byte, error) {
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
func (o *ESSecurityRemoteClusterPrivileges) UnmarshalJSON(bytes []byte) (err error) {
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
