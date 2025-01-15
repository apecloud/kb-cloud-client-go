// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// PodOwnerReference Reference to pod owner
type PodOwnerReference struct {
	// Owner type
	Kind *string `json:"kind,omitempty"`
	// Owner name
	Name *string `json:"name,omitempty"`
	// Owner UID
	Uid *string `json:"uid,omitempty"`
	// Owner API version
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Whether it is a controller
	Controller *bool `json:"controller,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPodOwnerReference instantiates a new PodOwnerReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPodOwnerReference() *PodOwnerReference {
	this := PodOwnerReference{}
	return &this
}

// NewPodOwnerReferenceWithDefaults instantiates a new PodOwnerReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPodOwnerReferenceWithDefaults() *PodOwnerReference {
	this := PodOwnerReference{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *PodOwnerReference) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodOwnerReference) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *PodOwnerReference) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *PodOwnerReference) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PodOwnerReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodOwnerReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PodOwnerReference) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PodOwnerReference) SetName(v string) {
	o.Name = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *PodOwnerReference) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodOwnerReference) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *PodOwnerReference) HasUid() bool {
	return o != nil && o.Uid != nil
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *PodOwnerReference) SetUid(v string) {
	o.Uid = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *PodOwnerReference) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodOwnerReference) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *PodOwnerReference) HasApiVersion() bool {
	return o != nil && o.ApiVersion != nil
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *PodOwnerReference) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *PodOwnerReference) GetController() bool {
	if o == nil || o.Controller == nil {
		var ret bool
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodOwnerReference) GetControllerOk() (*bool, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *PodOwnerReference) HasController() bool {
	return o != nil && o.Controller != nil
}

// SetController gets a reference to the given bool and assigns it to the Controller field.
func (o *PodOwnerReference) SetController(v bool) {
	o.Controller = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PodOwnerReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Controller != nil {
		toSerialize["controller"] = o.Controller
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PodOwnerReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kind       *string `json:"kind,omitempty"`
		Name       *string `json:"name,omitempty"`
		Uid        *string `json:"uid,omitempty"`
		ApiVersion *string `json:"apiVersion,omitempty"`
		Controller *bool   `json:"controller,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kind", "name", "uid", "apiVersion", "controller"})
	} else {
		return err
	}
	o.Kind = all.Kind
	o.Name = all.Name
	o.Uid = all.Uid
	o.ApiVersion = all.ApiVersion
	o.Controller = all.Controller

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
