// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TlsOption TLS options
type TlsOption struct {
	// certificate issuer
	Issuer CertificateIssuer `json:"issuer"`
	// the service name that will be set in certificate SAN
	ServiceName []string `json:"serviceName,omitempty"`
	// CA certificate content
	CaContent *string `json:"caContent,omitempty"`
	// Certificate content
	CertificateContent *string `json:"certificateContent,omitempty"`
	// Private key content
	PrivateKeyContent *string `json:"privateKeyContent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTlsOption instantiates a new TlsOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTlsOption(issuer CertificateIssuer) *TlsOption {
	this := TlsOption{}
	this.Issuer = issuer
	return &this
}

// NewTlsOptionWithDefaults instantiates a new TlsOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTlsOptionWithDefaults() *TlsOption {
	this := TlsOption{}
	return &this
}

// GetIssuer returns the Issuer field value.
func (o *TlsOption) GetIssuer() CertificateIssuer {
	if o == nil {
		var ret CertificateIssuer
		return ret
	}
	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *TlsOption) GetIssuerOk() (*CertificateIssuer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value.
func (o *TlsOption) SetIssuer(v CertificateIssuer) {
	o.Issuer = v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *TlsOption) GetServiceName() []string {
	if o == nil || o.ServiceName == nil {
		var ret []string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsOption) GetServiceNameOk() (*[]string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *TlsOption) HasServiceName() bool {
	return o != nil && o.ServiceName != nil
}

// SetServiceName gets a reference to the given []string and assigns it to the ServiceName field.
func (o *TlsOption) SetServiceName(v []string) {
	o.ServiceName = v
}

// GetCaContent returns the CaContent field value if set, zero value otherwise.
func (o *TlsOption) GetCaContent() string {
	if o == nil || o.CaContent == nil {
		var ret string
		return ret
	}
	return *o.CaContent
}

// GetCaContentOk returns a tuple with the CaContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsOption) GetCaContentOk() (*string, bool) {
	if o == nil || o.CaContent == nil {
		return nil, false
	}
	return o.CaContent, true
}

// HasCaContent returns a boolean if a field has been set.
func (o *TlsOption) HasCaContent() bool {
	return o != nil && o.CaContent != nil
}

// SetCaContent gets a reference to the given string and assigns it to the CaContent field.
func (o *TlsOption) SetCaContent(v string) {
	o.CaContent = &v
}

// GetCertificateContent returns the CertificateContent field value if set, zero value otherwise.
func (o *TlsOption) GetCertificateContent() string {
	if o == nil || o.CertificateContent == nil {
		var ret string
		return ret
	}
	return *o.CertificateContent
}

// GetCertificateContentOk returns a tuple with the CertificateContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsOption) GetCertificateContentOk() (*string, bool) {
	if o == nil || o.CertificateContent == nil {
		return nil, false
	}
	return o.CertificateContent, true
}

// HasCertificateContent returns a boolean if a field has been set.
func (o *TlsOption) HasCertificateContent() bool {
	return o != nil && o.CertificateContent != nil
}

// SetCertificateContent gets a reference to the given string and assigns it to the CertificateContent field.
func (o *TlsOption) SetCertificateContent(v string) {
	o.CertificateContent = &v
}

// GetPrivateKeyContent returns the PrivateKeyContent field value if set, zero value otherwise.
func (o *TlsOption) GetPrivateKeyContent() string {
	if o == nil || o.PrivateKeyContent == nil {
		var ret string
		return ret
	}
	return *o.PrivateKeyContent
}

// GetPrivateKeyContentOk returns a tuple with the PrivateKeyContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsOption) GetPrivateKeyContentOk() (*string, bool) {
	if o == nil || o.PrivateKeyContent == nil {
		return nil, false
	}
	return o.PrivateKeyContent, true
}

// HasPrivateKeyContent returns a boolean if a field has been set.
func (o *TlsOption) HasPrivateKeyContent() bool {
	return o != nil && o.PrivateKeyContent != nil
}

// SetPrivateKeyContent gets a reference to the given string and assigns it to the PrivateKeyContent field.
func (o *TlsOption) SetPrivateKeyContent(v string) {
	o.PrivateKeyContent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TlsOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["issuer"] = o.Issuer
	if o.ServiceName != nil {
		toSerialize["serviceName"] = o.ServiceName
	}
	if o.CaContent != nil {
		toSerialize["caContent"] = o.CaContent
	}
	if o.CertificateContent != nil {
		toSerialize["certificateContent"] = o.CertificateContent
	}
	if o.PrivateKeyContent != nil {
		toSerialize["privateKeyContent"] = o.PrivateKeyContent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TlsOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Issuer             *CertificateIssuer `json:"issuer"`
		ServiceName        []string           `json:"serviceName,omitempty"`
		CaContent          *string            `json:"caContent,omitempty"`
		CertificateContent *string            `json:"certificateContent,omitempty"`
		PrivateKeyContent  *string            `json:"privateKeyContent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Issuer == nil {
		return fmt.Errorf("required field issuer missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"issuer", "serviceName", "caContent", "certificateContent", "privateKeyContent"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Issuer.IsValid() {
		hasInvalidField = true
	} else {
		o.Issuer = *all.Issuer
	}
	o.ServiceName = all.ServiceName
	o.CaContent = all.CaContent
	o.CertificateContent = all.CertificateContent
	o.PrivateKeyContent = all.PrivateKeyContent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
