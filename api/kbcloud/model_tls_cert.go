// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION TlsCert
type TlsCert struct {
	// Configuration for TLS including all certificates and keys
	TlsConfig *TlsConfig `json:"tlsConfig,omitempty"`
	// Component these certs belong to
	ComponentName *string `json:"componentName,omitempty"`
	// Expiration days of CA certificate
	ExpirationDaysCa *int32 `json:"expirationDaysCA,omitempty"`
	// Expiration days of client certificate
	ExpirationDaysClient *int32 `json:"expirationDaysClient,omitempty"`
	// Expiration days of Keystore certificate
	ExpirationDaysKeystore *int32 `json:"expirationDaysKeystore,omitempty"`
	// Expiration days of Truststore certificate
	ExpirationDaysTruststore *int32 `json:"expirationDaysTruststore,omitempty"`
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

// GetTlsConfig returns the TlsConfig field value if set, zero value otherwise.
func (o *TlsCert) GetTlsConfig() TlsConfig {
	if o == nil || o.TlsConfig == nil {
		var ret TlsConfig
		return ret
	}
	return *o.TlsConfig
}

// GetTlsConfigOk returns a tuple with the TlsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCert) GetTlsConfigOk() (*TlsConfig, bool) {
	if o == nil || o.TlsConfig == nil {
		return nil, false
	}
	return o.TlsConfig, true
}

// HasTlsConfig returns a boolean if a field has been set.
func (o *TlsCert) HasTlsConfig() bool {
	return o != nil && o.TlsConfig != nil
}

// SetTlsConfig gets a reference to the given TlsConfig and assigns it to the TlsConfig field.
func (o *TlsCert) SetTlsConfig(v TlsConfig) {
	o.TlsConfig = &v
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

// GetExpirationDaysCa returns the ExpirationDaysCa field value if set, zero value otherwise.
func (o *TlsCert) GetExpirationDaysCa() int32 {
	if o == nil || o.ExpirationDaysCa == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationDaysCa
}

// GetExpirationDaysCaOk returns a tuple with the ExpirationDaysCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCert) GetExpirationDaysCaOk() (*int32, bool) {
	if o == nil || o.ExpirationDaysCa == nil {
		return nil, false
	}
	return o.ExpirationDaysCa, true
}

// HasExpirationDaysCa returns a boolean if a field has been set.
func (o *TlsCert) HasExpirationDaysCa() bool {
	return o != nil && o.ExpirationDaysCa != nil
}

// SetExpirationDaysCa gets a reference to the given int32 and assigns it to the ExpirationDaysCa field.
func (o *TlsCert) SetExpirationDaysCa(v int32) {
	o.ExpirationDaysCa = &v
}

// GetExpirationDaysClient returns the ExpirationDaysClient field value if set, zero value otherwise.
func (o *TlsCert) GetExpirationDaysClient() int32 {
	if o == nil || o.ExpirationDaysClient == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationDaysClient
}

// GetExpirationDaysClientOk returns a tuple with the ExpirationDaysClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCert) GetExpirationDaysClientOk() (*int32, bool) {
	if o == nil || o.ExpirationDaysClient == nil {
		return nil, false
	}
	return o.ExpirationDaysClient, true
}

// HasExpirationDaysClient returns a boolean if a field has been set.
func (o *TlsCert) HasExpirationDaysClient() bool {
	return o != nil && o.ExpirationDaysClient != nil
}

// SetExpirationDaysClient gets a reference to the given int32 and assigns it to the ExpirationDaysClient field.
func (o *TlsCert) SetExpirationDaysClient(v int32) {
	o.ExpirationDaysClient = &v
}

// GetExpirationDaysKeystore returns the ExpirationDaysKeystore field value if set, zero value otherwise.
func (o *TlsCert) GetExpirationDaysKeystore() int32 {
	if o == nil || o.ExpirationDaysKeystore == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationDaysKeystore
}

// GetExpirationDaysKeystoreOk returns a tuple with the ExpirationDaysKeystore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCert) GetExpirationDaysKeystoreOk() (*int32, bool) {
	if o == nil || o.ExpirationDaysKeystore == nil {
		return nil, false
	}
	return o.ExpirationDaysKeystore, true
}

// HasExpirationDaysKeystore returns a boolean if a field has been set.
func (o *TlsCert) HasExpirationDaysKeystore() bool {
	return o != nil && o.ExpirationDaysKeystore != nil
}

// SetExpirationDaysKeystore gets a reference to the given int32 and assigns it to the ExpirationDaysKeystore field.
func (o *TlsCert) SetExpirationDaysKeystore(v int32) {
	o.ExpirationDaysKeystore = &v
}

// GetExpirationDaysTruststore returns the ExpirationDaysTruststore field value if set, zero value otherwise.
func (o *TlsCert) GetExpirationDaysTruststore() int32 {
	if o == nil || o.ExpirationDaysTruststore == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationDaysTruststore
}

// GetExpirationDaysTruststoreOk returns a tuple with the ExpirationDaysTruststore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCert) GetExpirationDaysTruststoreOk() (*int32, bool) {
	if o == nil || o.ExpirationDaysTruststore == nil {
		return nil, false
	}
	return o.ExpirationDaysTruststore, true
}

// HasExpirationDaysTruststore returns a boolean if a field has been set.
func (o *TlsCert) HasExpirationDaysTruststore() bool {
	return o != nil && o.ExpirationDaysTruststore != nil
}

// SetExpirationDaysTruststore gets a reference to the given int32 and assigns it to the ExpirationDaysTruststore field.
func (o *TlsCert) SetExpirationDaysTruststore(v int32) {
	o.ExpirationDaysTruststore = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TlsCert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TlsConfig != nil {
		toSerialize["tlsConfig"] = o.TlsConfig
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	if o.ExpirationDaysCa != nil {
		toSerialize["expirationDaysCA"] = o.ExpirationDaysCa
	}
	if o.ExpirationDaysClient != nil {
		toSerialize["expirationDaysClient"] = o.ExpirationDaysClient
	}
	if o.ExpirationDaysKeystore != nil {
		toSerialize["expirationDaysKeystore"] = o.ExpirationDaysKeystore
	}
	if o.ExpirationDaysTruststore != nil {
		toSerialize["expirationDaysTruststore"] = o.ExpirationDaysTruststore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TlsCert) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TlsConfig                *TlsConfig `json:"tlsConfig,omitempty"`
		ComponentName            *string    `json:"componentName,omitempty"`
		ExpirationDaysCa         *int32     `json:"expirationDaysCA,omitempty"`
		ExpirationDaysClient     *int32     `json:"expirationDaysClient,omitempty"`
		ExpirationDaysKeystore   *int32     `json:"expirationDaysKeystore,omitempty"`
		ExpirationDaysTruststore *int32     `json:"expirationDaysTruststore,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"tlsConfig", "componentName", "expirationDaysCA", "expirationDaysClient", "expirationDaysKeystore", "expirationDaysTruststore"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TlsConfig != nil && all.TlsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TlsConfig = all.TlsConfig
	o.ComponentName = all.ComponentName
	o.ExpirationDaysCa = all.ExpirationDaysCa
	o.ExpirationDaysClient = all.ExpirationDaysClient
	o.ExpirationDaysKeystore = all.ExpirationDaysKeystore
	o.ExpirationDaysTruststore = all.ExpirationDaysTruststore

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
