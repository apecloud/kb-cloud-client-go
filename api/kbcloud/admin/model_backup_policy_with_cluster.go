// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupPolicyWithCluster Backup policy with owning cluster context for aggregate admin queries
type BackupPolicyWithCluster struct {
	// organization name of the cluster that owns the backup policy
	OrgName *string `json:"orgName,omitempty"`
	// cluster name that owns the backup policy
	ClusterName *string `json:"clusterName,omitempty"`
	// cluster ID that owns the backup policy
	ClusterId *string `json:"clusterID,omitempty"`
	// BackupPolicy is the payload for KubeBlocks cluster backup policy
	Policy *BackupPolicy `json:"policy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupPolicyWithCluster instantiates a new BackupPolicyWithCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupPolicyWithCluster() *BackupPolicyWithCluster {
	this := BackupPolicyWithCluster{}
	return &this
}

// NewBackupPolicyWithClusterWithDefaults instantiates a new BackupPolicyWithCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupPolicyWithClusterWithDefaults() *BackupPolicyWithCluster {
	this := BackupPolicyWithCluster{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *BackupPolicyWithCluster) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicyWithCluster) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *BackupPolicyWithCluster) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *BackupPolicyWithCluster) SetOrgName(v string) {
	o.OrgName = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *BackupPolicyWithCluster) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicyWithCluster) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *BackupPolicyWithCluster) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *BackupPolicyWithCluster) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *BackupPolicyWithCluster) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicyWithCluster) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *BackupPolicyWithCluster) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *BackupPolicyWithCluster) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *BackupPolicyWithCluster) GetPolicy() BackupPolicy {
	if o == nil || o.Policy == nil {
		var ret BackupPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicyWithCluster) GetPolicyOk() (*BackupPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *BackupPolicyWithCluster) HasPolicy() bool {
	return o != nil && o.Policy != nil
}

// SetPolicy gets a reference to the given BackupPolicy and assigns it to the Policy field.
func (o *BackupPolicyWithCluster) SetPolicy(v BackupPolicy) {
	o.Policy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupPolicyWithCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.ClusterId != nil {
		toSerialize["clusterID"] = o.ClusterId
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupPolicyWithCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName     *string       `json:"orgName,omitempty"`
		ClusterName *string       `json:"clusterName,omitempty"`
		ClusterId   *string       `json:"clusterID,omitempty"`
		Policy      *BackupPolicy `json:"policy,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "clusterName", "clusterID", "policy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OrgName = all.OrgName
	o.ClusterName = all.ClusterName
	o.ClusterId = all.ClusterId
	if all.Policy != nil && all.Policy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Policy = all.Policy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
