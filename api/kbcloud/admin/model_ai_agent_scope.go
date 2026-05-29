// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentScope struct {
	OrgName             *string                  `json:"orgName,omitempty"`
	EnvironmentName     *string                  `json:"environmentName,omitempty"`
	ClusterName         *string                  `json:"clusterName,omitempty"`
	KubernetesName      *string                  `json:"kubernetesName,omitempty"`
	KubernetesNamespace *string                  `json:"kubernetesNamespace,omitempty"`
	ClusterDefinition   *string                  `json:"clusterDefinition,omitempty"`
	EngineMode          *string                  `json:"engineMode,omitempty"`
	Components          []map[string]interface{} `json:"components,omitempty"`
	ServiceRefs         []map[string]interface{} `json:"serviceRefs,omitempty"`
	Label               *string                  `json:"label,omitempty"`
	Namespace           *string                  `json:"namespace,omitempty"`
	ObjectRef           *AiAgentObjectRef        `json:"objectRef,omitempty"`
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

// GetKubernetesName returns the KubernetesName field value if set, zero value otherwise.
func (o *AiAgentScope) GetKubernetesName() string {
	if o == nil || o.KubernetesName == nil {
		var ret string
		return ret
	}
	return *o.KubernetesName
}

// GetKubernetesNameOk returns a tuple with the KubernetesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetKubernetesNameOk() (*string, bool) {
	if o == nil || o.KubernetesName == nil {
		return nil, false
	}
	return o.KubernetesName, true
}

// HasKubernetesName returns a boolean if a field has been set.
func (o *AiAgentScope) HasKubernetesName() bool {
	return o != nil && o.KubernetesName != nil
}

// SetKubernetesName gets a reference to the given string and assigns it to the KubernetesName field.
func (o *AiAgentScope) SetKubernetesName(v string) {
	o.KubernetesName = &v
}

// GetKubernetesNamespace returns the KubernetesNamespace field value if set, zero value otherwise.
func (o *AiAgentScope) GetKubernetesNamespace() string {
	if o == nil || o.KubernetesNamespace == nil {
		var ret string
		return ret
	}
	return *o.KubernetesNamespace
}

// GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetKubernetesNamespaceOk() (*string, bool) {
	if o == nil || o.KubernetesNamespace == nil {
		return nil, false
	}
	return o.KubernetesNamespace, true
}

// HasKubernetesNamespace returns a boolean if a field has been set.
func (o *AiAgentScope) HasKubernetesNamespace() bool {
	return o != nil && o.KubernetesNamespace != nil
}

// SetKubernetesNamespace gets a reference to the given string and assigns it to the KubernetesNamespace field.
func (o *AiAgentScope) SetKubernetesNamespace(v string) {
	o.KubernetesNamespace = &v
}

// GetClusterDefinition returns the ClusterDefinition field value if set, zero value otherwise.
func (o *AiAgentScope) GetClusterDefinition() string {
	if o == nil || o.ClusterDefinition == nil {
		var ret string
		return ret
	}
	return *o.ClusterDefinition
}

// GetClusterDefinitionOk returns a tuple with the ClusterDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetClusterDefinitionOk() (*string, bool) {
	if o == nil || o.ClusterDefinition == nil {
		return nil, false
	}
	return o.ClusterDefinition, true
}

// HasClusterDefinition returns a boolean if a field has been set.
func (o *AiAgentScope) HasClusterDefinition() bool {
	return o != nil && o.ClusterDefinition != nil
}

// SetClusterDefinition gets a reference to the given string and assigns it to the ClusterDefinition field.
func (o *AiAgentScope) SetClusterDefinition(v string) {
	o.ClusterDefinition = &v
}

// GetEngineMode returns the EngineMode field value if set, zero value otherwise.
func (o *AiAgentScope) GetEngineMode() string {
	if o == nil || o.EngineMode == nil {
		var ret string
		return ret
	}
	return *o.EngineMode
}

// GetEngineModeOk returns a tuple with the EngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetEngineModeOk() (*string, bool) {
	if o == nil || o.EngineMode == nil {
		return nil, false
	}
	return o.EngineMode, true
}

// HasEngineMode returns a boolean if a field has been set.
func (o *AiAgentScope) HasEngineMode() bool {
	return o != nil && o.EngineMode != nil
}

// SetEngineMode gets a reference to the given string and assigns it to the EngineMode field.
func (o *AiAgentScope) SetEngineMode(v string) {
	o.EngineMode = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *AiAgentScope) GetComponents() []map[string]interface{} {
	if o == nil || o.Components == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetComponentsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *AiAgentScope) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []map[string]interface{} and assigns it to the Components field.
func (o *AiAgentScope) SetComponents(v []map[string]interface{}) {
	o.Components = v
}

// GetServiceRefs returns the ServiceRefs field value if set, zero value otherwise.
func (o *AiAgentScope) GetServiceRefs() []map[string]interface{} {
	if o == nil || o.ServiceRefs == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.ServiceRefs
}

// GetServiceRefsOk returns a tuple with the ServiceRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetServiceRefsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ServiceRefs == nil {
		return nil, false
	}
	return &o.ServiceRefs, true
}

// HasServiceRefs returns a boolean if a field has been set.
func (o *AiAgentScope) HasServiceRefs() bool {
	return o != nil && o.ServiceRefs != nil
}

// SetServiceRefs gets a reference to the given []map[string]interface{} and assigns it to the ServiceRefs field.
func (o *AiAgentScope) SetServiceRefs(v []map[string]interface{}) {
	o.ServiceRefs = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AiAgentScope) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AiAgentScope) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AiAgentScope) SetLabel(v string) {
	o.Label = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AiAgentScope) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentScope) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AiAgentScope) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *AiAgentScope) SetNamespace(v string) {
	o.Namespace = &v
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
	if o.KubernetesName != nil {
		toSerialize["kubernetesName"] = o.KubernetesName
	}
	if o.KubernetesNamespace != nil {
		toSerialize["kubernetesNamespace"] = o.KubernetesNamespace
	}
	if o.ClusterDefinition != nil {
		toSerialize["clusterDefinition"] = o.ClusterDefinition
	}
	if o.EngineMode != nil {
		toSerialize["engineMode"] = o.EngineMode
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.ServiceRefs != nil {
		toSerialize["serviceRefs"] = o.ServiceRefs
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
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
		OrgName             *string                  `json:"orgName,omitempty"`
		EnvironmentName     *string                  `json:"environmentName,omitempty"`
		ClusterName         *string                  `json:"clusterName,omitempty"`
		KubernetesName      *string                  `json:"kubernetesName,omitempty"`
		KubernetesNamespace *string                  `json:"kubernetesNamespace,omitempty"`
		ClusterDefinition   *string                  `json:"clusterDefinition,omitempty"`
		EngineMode          *string                  `json:"engineMode,omitempty"`
		Components          []map[string]interface{} `json:"components,omitempty"`
		ServiceRefs         []map[string]interface{} `json:"serviceRefs,omitempty"`
		Label               *string                  `json:"label,omitempty"`
		Namespace           *string                  `json:"namespace,omitempty"`
		ObjectRef           *AiAgentObjectRef        `json:"objectRef,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "environmentName", "clusterName", "kubernetesName", "kubernetesNamespace", "clusterDefinition", "engineMode", "components", "serviceRefs", "label", "namespace", "objectRef"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OrgName = all.OrgName
	o.EnvironmentName = all.EnvironmentName
	o.ClusterName = all.ClusterName
	o.KubernetesName = all.KubernetesName
	o.KubernetesNamespace = all.KubernetesNamespace
	o.ClusterDefinition = all.ClusterDefinition
	o.EngineMode = all.EngineMode
	o.Components = all.Components
	o.ServiceRefs = all.ServiceRefs
	o.Label = all.Label
	o.Namespace = all.Namespace
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
