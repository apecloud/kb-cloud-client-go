// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertSMSConfig Alert sms config
type AlertSMSConfig struct {
	Name               *string `json:"name,omitempty"`
	AccessKeyId        string  `json:"accessKeyID"`
	AccessKeySecret    string  `json:"accessKeySecret"`
	Endpoint           *string `json:"endpoint,omitempty"`
	AlertTemplateCode  *string `json:"alertTemplateCode,omitempty"`
	VerifyTemplateCode *string `json:"verifyTemplateCode,omitempty"`
	SignName           *string `json:"signName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertSMSConfig instantiates a new AlertSMSConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertSMSConfig(accessKeyId string, accessKeySecret string) *AlertSMSConfig {
	this := AlertSMSConfig{}
	this.AccessKeyId = accessKeyId
	this.AccessKeySecret = accessKeySecret
	return &this
}

// NewAlertSMSConfigWithDefaults instantiates a new AlertSMSConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertSMSConfigWithDefaults() *AlertSMSConfig {
	this := AlertSMSConfig{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertSMSConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSMSConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertSMSConfig) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertSMSConfig) SetName(v string) {
	o.Name = &v
}

// GetAccessKeyId returns the AccessKeyId field value.
func (o *AlertSMSConfig) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *AlertSMSConfig) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value.
func (o *AlertSMSConfig) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetAccessKeySecret returns the AccessKeySecret field value.
func (o *AlertSMSConfig) GetAccessKeySecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessKeySecret
}

// GetAccessKeySecretOk returns a tuple with the AccessKeySecret field value
// and a boolean to check if the value has been set.
func (o *AlertSMSConfig) GetAccessKeySecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeySecret, true
}

// SetAccessKeySecret sets field value.
func (o *AlertSMSConfig) SetAccessKeySecret(v string) {
	o.AccessKeySecret = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *AlertSMSConfig) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSMSConfig) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *AlertSMSConfig) HasEndpoint() bool {
	return o != nil && o.Endpoint != nil
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *AlertSMSConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetAlertTemplateCode returns the AlertTemplateCode field value if set, zero value otherwise.
func (o *AlertSMSConfig) GetAlertTemplateCode() string {
	if o == nil || o.AlertTemplateCode == nil {
		var ret string
		return ret
	}
	return *o.AlertTemplateCode
}

// GetAlertTemplateCodeOk returns a tuple with the AlertTemplateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSMSConfig) GetAlertTemplateCodeOk() (*string, bool) {
	if o == nil || o.AlertTemplateCode == nil {
		return nil, false
	}
	return o.AlertTemplateCode, true
}

// HasAlertTemplateCode returns a boolean if a field has been set.
func (o *AlertSMSConfig) HasAlertTemplateCode() bool {
	return o != nil && o.AlertTemplateCode != nil
}

// SetAlertTemplateCode gets a reference to the given string and assigns it to the AlertTemplateCode field.
func (o *AlertSMSConfig) SetAlertTemplateCode(v string) {
	o.AlertTemplateCode = &v
}

// GetVerifyTemplateCode returns the VerifyTemplateCode field value if set, zero value otherwise.
func (o *AlertSMSConfig) GetVerifyTemplateCode() string {
	if o == nil || o.VerifyTemplateCode == nil {
		var ret string
		return ret
	}
	return *o.VerifyTemplateCode
}

// GetVerifyTemplateCodeOk returns a tuple with the VerifyTemplateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSMSConfig) GetVerifyTemplateCodeOk() (*string, bool) {
	if o == nil || o.VerifyTemplateCode == nil {
		return nil, false
	}
	return o.VerifyTemplateCode, true
}

// HasVerifyTemplateCode returns a boolean if a field has been set.
func (o *AlertSMSConfig) HasVerifyTemplateCode() bool {
	return o != nil && o.VerifyTemplateCode != nil
}

// SetVerifyTemplateCode gets a reference to the given string and assigns it to the VerifyTemplateCode field.
func (o *AlertSMSConfig) SetVerifyTemplateCode(v string) {
	o.VerifyTemplateCode = &v
}

// GetSignName returns the SignName field value if set, zero value otherwise.
func (o *AlertSMSConfig) GetSignName() string {
	if o == nil || o.SignName == nil {
		var ret string
		return ret
	}
	return *o.SignName
}

// GetSignNameOk returns a tuple with the SignName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertSMSConfig) GetSignNameOk() (*string, bool) {
	if o == nil || o.SignName == nil {
		return nil, false
	}
	return o.SignName, true
}

// HasSignName returns a boolean if a field has been set.
func (o *AlertSMSConfig) HasSignName() bool {
	return o != nil && o.SignName != nil
}

// SetSignName gets a reference to the given string and assigns it to the SignName field.
func (o *AlertSMSConfig) SetSignName(v string) {
	o.SignName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertSMSConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["accessKeyID"] = o.AccessKeyId
	toSerialize["accessKeySecret"] = o.AccessKeySecret
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.AlertTemplateCode != nil {
		toSerialize["alertTemplateCode"] = o.AlertTemplateCode
	}
	if o.VerifyTemplateCode != nil {
		toSerialize["verifyTemplateCode"] = o.VerifyTemplateCode
	}
	if o.SignName != nil {
		toSerialize["signName"] = o.SignName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertSMSConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name               *string `json:"name,omitempty"`
		AccessKeyId        *string `json:"accessKeyID"`
		AccessKeySecret    *string `json:"accessKeySecret"`
		Endpoint           *string `json:"endpoint,omitempty"`
		AlertTemplateCode  *string `json:"alertTemplateCode,omitempty"`
		VerifyTemplateCode *string `json:"verifyTemplateCode,omitempty"`
		SignName           *string `json:"signName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.AccessKeyId == nil {
		return fmt.Errorf("required field accessKeyID missing")
	}
	if all.AccessKeySecret == nil {
		return fmt.Errorf("required field accessKeySecret missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "accessKeyID", "accessKeySecret", "endpoint", "alertTemplateCode", "verifyTemplateCode", "signName"})
	} else {
		return err
	}
	o.Name = all.Name
	o.AccessKeyId = *all.AccessKeyId
	o.AccessKeySecret = *all.AccessKeySecret
	o.Endpoint = all.Endpoint
	o.AlertTemplateCode = all.AlertTemplateCode
	o.VerifyTemplateCode = all.VerifyTemplateCode
	o.SignName = all.SignName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
