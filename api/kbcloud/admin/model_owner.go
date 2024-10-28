// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// Owner Owner related to the Event

type Owner struct {
	// APIVersion is the API version of the owner.
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is the type of the owner.
	Kind *string `json:"kind,omitempty"`
	// Name is the name of the owner.
	Name *string `json:"name,omitempty"`
	// UID is the unique identifier of the owner.
	Uid *string `json:"uid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwner instantiates a new Owner object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwner() *Owner {
	this := Owner{}
	return &this
}

// NewOwnerWithDefaults instantiates a new Owner object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnerWithDefaults() *Owner {
	this := Owner{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *Owner) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owner) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *Owner) HasApiVersion() bool {
	return o != nil && o.ApiVersion != nil
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *Owner) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Owner) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owner) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Owner) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Owner) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Owner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Owner) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Owner) SetName(v string) {
	o.Name = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *Owner) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Owner) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *Owner) HasUid() bool {
	return o != nil && o.Uid != nil
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *Owner) SetUid(v string) {
	o.Uid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Owner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Owner) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiVersion *string `json:"apiVersion,omitempty"`
		Kind       *string `json:"kind,omitempty"`
		Name       *string `json:"name,omitempty"`
		Uid        *string `json:"uid,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"apiVersion", "kind", "name", "uid"})
	} else {
		return err
	}
	o.ApiVersion = all.ApiVersion
	o.Kind = all.Kind
	o.Name = all.Name
	o.Uid = all.Uid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
