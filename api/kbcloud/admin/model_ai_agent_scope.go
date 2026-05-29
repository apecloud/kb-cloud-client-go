// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentScope struct {
	OrgName         *string           `json:"orgName,omitempty"`
	EnvironmentName *string           `json:"environmentName,omitempty"`
	ClusterName     *string           `json:"clusterName,omitempty"`
	ObjectRef       *AiAgentObjectRef `json:"objectRef,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentScope instantiates a new AiAgentScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentScope() *AiAgentScope {
	this := AiAgentScope{}
	return &this
}

// NewAiAgentScopeWithDefaults instantiates a new AiAgentScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentScopeWithDefaults() *AiAgentScope {
	this := AiAgentScope{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiAgentScope) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiAgentScope) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiAgentScope) SetOrgName(v string) {
	o.OrgName = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *AiAgentScope) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *AiAgentScope) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *AiAgentScope) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *AiAgentScope) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *AiAgentScope) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *AiAgentScope) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetObjectRef returns the ObjectRef field value if set, zero value otherwise.
func (o *AiAgentScope) GetObjectRef() AiAgentObjectRef {
	if o == nil || o.ObjectRef == nil {
		var ret AiAgentObjectRef
		return ret
	}
	return *o.ObjectRef
}

// GetObjectRefOk returns a tuple with the ObjectRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetObjectRefOk() (*AiAgentObjectRef, bool) {
	if o == nil || o.ObjectRef == nil {
		return nil, false
	}
	return o.ObjectRef, true
}

// HasObjectRef returns a boolean if a field has been set.
func (o *AiAgentScope) HasObjectRef() bool {
	return o != nil && o.ObjectRef != nil
}

// SetObjectRef gets a reference to the given AiAgentObjectRef and assigns it to the ObjectRef field.
func (o *AiAgentScope) SetObjectRef(v AiAgentObjectRef) {
	o.ObjectRef = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.ObjectRef != nil {
		toSerialize["objectRef"] = o.ObjectRef
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName         *string           `json:"orgName,omitempty"`
		EnvironmentName *string           `json:"environmentName,omitempty"`
		ClusterName     *string           `json:"clusterName,omitempty"`
		ObjectRef       *AiAgentObjectRef `json:"objectRef,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "environmentName", "clusterName", "objectRef"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OrgName = all.OrgName
	o.EnvironmentName = all.EnvironmentName
	o.ClusterName = all.ClusterName
	if all.ObjectRef != nil && all.ObjectRef.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectRef = all.ObjectRef

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
