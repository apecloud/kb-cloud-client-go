// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Certificate certificates and it's config
type Certificate struct {
	// Name of the TLS certificates
	Name TlsName `json:"name"`
	// TLS type
	Type TlsType `json:"type"`
	// Content of the TLS certificates
	Content string `json:"content"`
	// Expiration days for this certificate
	ExpirationDays *int32 `json:"expirationDays,omitempty"`
	// DNS name in the certificate SAN
	Dns []string `json:"DNS,omitempty"`
	// IP address in the certificate SAN
	IpAddress []string `json:"ipAddress,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCertificate instantiates a new Certificate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCertificate(name TlsName, typeVar TlsType, content string) *Certificate {
	this := Certificate{}
	this.Name = name
	this.Type = typeVar
	this.Content = content
	return &this
}

// NewCertificateWithDefaults instantiates a new Certificate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCertificateWithDefaults() *Certificate {
	this := Certificate{}
	var typeVar TlsType = TlsTypeStandardTls
	this.Type = typeVar
	return &this
}

// GetName returns the Name field value.
func (o *Certificate) GetName() TlsName {
	if o == nil {
		var ret TlsName
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetNameOk() (*TlsName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Certificate) SetName(v TlsName) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *Certificate) GetType() TlsType {
	if o == nil {
		var ret TlsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetTypeOk() (*TlsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Certificate) SetType(v TlsType) {
	o.Type = v
}

// GetContent returns the Content field value.
func (o *Certificate) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *Certificate) SetContent(v string) {
	o.Content = v
}

// GetExpirationDays returns the ExpirationDays field value if set, zero value otherwise.
func (o *Certificate) GetExpirationDays() int32 {
	if o == nil || o.ExpirationDays == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationDays
}

// GetExpirationDaysOk returns a tuple with the ExpirationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetExpirationDaysOk() (*int32, bool) {
	if o == nil || o.ExpirationDays == nil {
		return nil, false
	}
	return o.ExpirationDays, true
}

// HasExpirationDays returns a boolean if a field has been set.
func (o *Certificate) HasExpirationDays() bool {
	return o != nil && o.ExpirationDays != nil
}

// SetExpirationDays gets a reference to the given int32 and assigns it to the ExpirationDays field.
func (o *Certificate) SetExpirationDays(v int32) {
	o.ExpirationDays = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *Certificate) GetDns() []string {
	if o == nil || o.Dns == nil {
		var ret []string
		return ret
	}
	return o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetDnsOk() (*[]string, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return &o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *Certificate) HasDns() bool {
	return o != nil && o.Dns != nil
}

// SetDns gets a reference to the given []string and assigns it to the Dns field.
func (o *Certificate) SetDns(v []string) {
	o.Dns = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *Certificate) GetIpAddress() []string {
	if o == nil || o.IpAddress == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetIpAddressOk() (*[]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *Certificate) HasIpAddress() bool {
	return o != nil && o.IpAddress != nil
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *Certificate) SetIpAddress(v []string) {
	o.IpAddress = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Certificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	if o.ExpirationDays != nil {
		toSerialize["expirationDays"] = o.ExpirationDays
	}
	if o.Dns != nil {
		toSerialize["DNS"] = o.Dns
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Certificate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *TlsName `json:"name"`
		Type           *TlsType `json:"type"`
		Content        *string  `json:"content"`
		ExpirationDays *int32   `json:"expirationDays,omitempty"`
		Dns            []string `json:"DNS,omitempty"`
		IpAddress      []string `json:"ipAddress,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "content", "expirationDays", "DNS", "ipAddress"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Name.IsValid() {
		hasInvalidField = true
	} else {
		o.Name = *all.Name
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Content = *all.Content
	o.ExpirationDays = all.ExpirationDays
	o.Dns = all.Dns
	o.IpAddress = all.IpAddress

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
