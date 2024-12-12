// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

// TlsConfig Configuration for TLS including all certificates and keys
type TlsConfig struct {
	// CA Certificate
	CaCert *string `json:"caCert,omitempty"`
	// Client Side Certificate
	ClientCert *string `json:"clientCert,omitempty"`
	// Client Side key
	ClientKey *string `json:"clientKey,omitempty"`
	// Keystore in PEM format
	Keystore *string `json:"keystore,omitempty"`
	// Truststore in PEM format
	Truststore *string `json:"truststore,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTlsConfig instantiates a new TlsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTlsConfig() *TlsConfig {
	this := TlsConfig{}
	return &this
}

// NewTlsConfigWithDefaults instantiates a new TlsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTlsConfigWithDefaults() *TlsConfig {
	this := TlsConfig{}
	return &this
}

// GetCaCert returns the CaCert field value if set, zero value otherwise.
func (o *TlsConfig) GetCaCert() string {
	if o == nil || o.CaCert == nil {
		var ret string
		return ret
	}
	return *o.CaCert
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfig) GetCaCertOk() (*string, bool) {
	if o == nil || o.CaCert == nil {
		return nil, false
	}
	return o.CaCert, true
}

// HasCaCert returns a boolean if a field has been set.
func (o *TlsConfig) HasCaCert() bool {
	return o != nil && o.CaCert != nil
}

// SetCaCert gets a reference to the given string and assigns it to the CaCert field.
func (o *TlsConfig) SetCaCert(v string) {
	o.CaCert = &v
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise.
func (o *TlsConfig) GetClientCert() string {
	if o == nil || o.ClientCert == nil {
		var ret string
		return ret
	}
	return *o.ClientCert
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfig) GetClientCertOk() (*string, bool) {
	if o == nil || o.ClientCert == nil {
		return nil, false
	}
	return o.ClientCert, true
}

// HasClientCert returns a boolean if a field has been set.
func (o *TlsConfig) HasClientCert() bool {
	return o != nil && o.ClientCert != nil
}

// SetClientCert gets a reference to the given string and assigns it to the ClientCert field.
func (o *TlsConfig) SetClientCert(v string) {
	o.ClientCert = &v
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise.
func (o *TlsConfig) GetClientKey() string {
	if o == nil || o.ClientKey == nil {
		var ret string
		return ret
	}
	return *o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfig) GetClientKeyOk() (*string, bool) {
	if o == nil || o.ClientKey == nil {
		return nil, false
	}
	return o.ClientKey, true
}

// HasClientKey returns a boolean if a field has been set.
func (o *TlsConfig) HasClientKey() bool {
	return o != nil && o.ClientKey != nil
}

// SetClientKey gets a reference to the given string and assigns it to the ClientKey field.
func (o *TlsConfig) SetClientKey(v string) {
	o.ClientKey = &v
}

// GetKeystore returns the Keystore field value if set, zero value otherwise.
func (o *TlsConfig) GetKeystore() string {
	if o == nil || o.Keystore == nil {
		var ret string
		return ret
	}
	return *o.Keystore
}

// GetKeystoreOk returns a tuple with the Keystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfig) GetKeystoreOk() (*string, bool) {
	if o == nil || o.Keystore == nil {
		return nil, false
	}
	return o.Keystore, true
}

// HasKeystore returns a boolean if a field has been set.
func (o *TlsConfig) HasKeystore() bool {
	return o != nil && o.Keystore != nil
}

// SetKeystore gets a reference to the given string and assigns it to the Keystore field.
func (o *TlsConfig) SetKeystore(v string) {
	o.Keystore = &v
}

// GetTruststore returns the Truststore field value if set, zero value otherwise.
func (o *TlsConfig) GetTruststore() string {
	if o == nil || o.Truststore == nil {
		var ret string
		return ret
	}
	return *o.Truststore
}

// GetTruststoreOk returns a tuple with the Truststore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfig) GetTruststoreOk() (*string, bool) {
	if o == nil || o.Truststore == nil {
		return nil, false
	}
	return o.Truststore, true
}

// HasTruststore returns a boolean if a field has been set.
func (o *TlsConfig) HasTruststore() bool {
	return o != nil && o.Truststore != nil
}

// SetTruststore gets a reference to the given string and assigns it to the Truststore field.
func (o *TlsConfig) SetTruststore(v string) {
	o.Truststore = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TlsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CaCert != nil {
		toSerialize["caCert"] = o.CaCert
	}
	if o.ClientCert != nil {
		toSerialize["clientCert"] = o.ClientCert
	}
	if o.ClientKey != nil {
		toSerialize["clientKey"] = o.ClientKey
	}
	if o.Keystore != nil {
		toSerialize["keystore"] = o.Keystore
	}
	if o.Truststore != nil {
		toSerialize["truststore"] = o.Truststore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TlsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaCert     *string `json:"caCert,omitempty"`
		ClientCert *string `json:"clientCert,omitempty"`
		ClientKey  *string `json:"clientKey,omitempty"`
		Keystore   *string `json:"keystore,omitempty"`
		Truststore *string `json:"truststore,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"caCert", "clientCert", "clientKey", "keystore", "truststore"})
	} else {
		return err
	}
	o.CaCert = all.CaCert
	o.ClientCert = all.ClientCert
	o.ClientKey = all.ClientKey
	o.Keystore = all.Keystore
	o.Truststore = all.Truststore

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
