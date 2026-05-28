// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DiagnosticRelatedResource Resource connected to one diagnostic check.
type DiagnosticRelatedResource struct {
	// Kubernetes or KBE resource kind.
	Kind *string `json:"kind,omitempty"`
	// Kubernetes namespace, when applicable.
	Namespace *string `json:"namespace,omitempty"`
	// Resource name.
	Name *string `json:"name,omitempty"`
	// Why this resource is related to the check.
	Relation *string `json:"relation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticRelatedResource instantiates a new DiagnosticRelatedResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticRelatedResource() *DiagnosticRelatedResource {
	this := DiagnosticRelatedResource{}
	return &this
}

// NewDiagnosticRelatedResourceWithDefaults instantiates a new DiagnosticRelatedResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticRelatedResourceWithDefaults() *DiagnosticRelatedResource {
	this := DiagnosticRelatedResource{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *DiagnosticRelatedResource) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticRelatedResource) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *DiagnosticRelatedResource) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *DiagnosticRelatedResource) SetKind(v string) {
	o.Kind = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *DiagnosticRelatedResource) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticRelatedResource) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *DiagnosticRelatedResource) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *DiagnosticRelatedResource) SetNamespace(v string) {
	o.Namespace = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DiagnosticRelatedResource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticRelatedResource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DiagnosticRelatedResource) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DiagnosticRelatedResource) SetName(v string) {
	o.Name = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *DiagnosticRelatedResource) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticRelatedResource) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *DiagnosticRelatedResource) HasRelation() bool {
	return o != nil && o.Relation != nil
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *DiagnosticRelatedResource) SetRelation(v string) {
	o.Relation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticRelatedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticRelatedResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kind      *string `json:"kind,omitempty"`
		Namespace *string `json:"namespace,omitempty"`
		Name      *string `json:"name,omitempty"`
		Relation  *string `json:"relation,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kind", "namespace", "name", "relation"})
	} else {
		return err
	}
	o.Kind = all.Kind
	o.Namespace = all.Namespace
	o.Name = all.Name
	o.Relation = all.Relation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
