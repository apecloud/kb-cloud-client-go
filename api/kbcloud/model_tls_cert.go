// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TlsCert struct {
	// TLS certificates
	Certificates []Certificate `json:"certificates,omitempty"`
	// Component these certs belong to
	ComponentName *string `json:"componentName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTlsCert instantiates a new TlsCert object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTlsCert() *TlsCert {
	this := TlsCert{}
	return &this
}

// NewTlsCertWithDefaults instantiates a new TlsCert object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTlsCertWithDefaults() *TlsCert {
	this := TlsCert{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *TlsCert) GetCertificates() []Certificate {
	if o == nil || o.Certificates == nil {
		var ret []Certificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCert) GetCertificatesOk() (*[]Certificate, bool) {
	if o == nil || o.Certificates == nil {
		return nil, false
	}
	return &o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *TlsCert) HasCertificates() bool {
	return o != nil && o.Certificates != nil
}

// SetCertificates gets a reference to the given []Certificate and assigns it to the Certificates field.
func (o *TlsCert) SetCertificates(v []Certificate) {
	o.Certificates = v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *TlsCert) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCert) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *TlsCert) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *TlsCert) SetComponentName(v string) {
	o.ComponentName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TlsCert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Certificates != nil {
		toSerialize["certificates"] = o.Certificates
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TlsCert) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Certificates  []Certificate `json:"certificates,omitempty"`
		ComponentName *string       `json:"componentName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"certificates", "componentName"})
	} else {
		return err
	}
	o.Certificates = all.Certificates
	o.ComponentName = all.ComponentName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
