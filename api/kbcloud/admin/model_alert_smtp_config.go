// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertSMTPConfig Alert smtp config
// NODESCRIPTION AlertSMTPConfig
//
// Deprecated: This model is deprecated.
type AlertSMTPConfig struct {
	// SMTP authentication information
	SmtpAuthPassword string `json:"smtp_auth_password"`
	// SMTP authentication information
	SmtpAuthUsername string `json:"smtp_auth_username"`
	// The sender's address
	SmtpFrom string `json:"smtp_from"`
	// The SMTP host through which emails are sent
	SmtpSmarthost string `json:"smtp_smarthost"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertSMTPConfig instantiates a new AlertSMTPConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertSMTPConfig(smtpAuthPassword string, smtpAuthUsername string, smtpFrom string, smtpSmarthost string) *AlertSMTPConfig {
	this := AlertSMTPConfig{}
	this.SmtpAuthPassword = smtpAuthPassword
	this.SmtpAuthUsername = smtpAuthUsername
	this.SmtpFrom = smtpFrom
	this.SmtpSmarthost = smtpSmarthost
	return &this
}

// NewAlertSMTPConfigWithDefaults instantiates a new AlertSMTPConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertSMTPConfigWithDefaults() *AlertSMTPConfig {
	this := AlertSMTPConfig{}
	return &this
}

// GetSmtpAuthPassword returns the SmtpAuthPassword field value.
func (o *AlertSMTPConfig) GetSmtpAuthPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SmtpAuthPassword
}

// GetSmtpAuthPasswordOk returns a tuple with the SmtpAuthPassword field value
// and a boolean to check if the value has been set.
func (o *AlertSMTPConfig) GetSmtpAuthPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmtpAuthPassword, true
}

// SetSmtpAuthPassword sets field value.
func (o *AlertSMTPConfig) SetSmtpAuthPassword(v string) {
	o.SmtpAuthPassword = v
}

// GetSmtpAuthUsername returns the SmtpAuthUsername field value.
func (o *AlertSMTPConfig) GetSmtpAuthUsername() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SmtpAuthUsername
}

// GetSmtpAuthUsernameOk returns a tuple with the SmtpAuthUsername field value
// and a boolean to check if the value has been set.
func (o *AlertSMTPConfig) GetSmtpAuthUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmtpAuthUsername, true
}

// SetSmtpAuthUsername sets field value.
func (o *AlertSMTPConfig) SetSmtpAuthUsername(v string) {
	o.SmtpAuthUsername = v
}

// GetSmtpFrom returns the SmtpFrom field value.
func (o *AlertSMTPConfig) GetSmtpFrom() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SmtpFrom
}

// GetSmtpFromOk returns a tuple with the SmtpFrom field value
// and a boolean to check if the value has been set.
func (o *AlertSMTPConfig) GetSmtpFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmtpFrom, true
}

// SetSmtpFrom sets field value.
func (o *AlertSMTPConfig) SetSmtpFrom(v string) {
	o.SmtpFrom = v
}

// GetSmtpSmarthost returns the SmtpSmarthost field value.
func (o *AlertSMTPConfig) GetSmtpSmarthost() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SmtpSmarthost
}

// GetSmtpSmarthostOk returns a tuple with the SmtpSmarthost field value
// and a boolean to check if the value has been set.
func (o *AlertSMTPConfig) GetSmtpSmarthostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmtpSmarthost, true
}

// SetSmtpSmarthost sets field value.
func (o *AlertSMTPConfig) SetSmtpSmarthost(v string) {
	o.SmtpSmarthost = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertSMTPConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["smtp_auth_password"] = o.SmtpAuthPassword
	toSerialize["smtp_auth_username"] = o.SmtpAuthUsername
	toSerialize["smtp_from"] = o.SmtpFrom
	toSerialize["smtp_smarthost"] = o.SmtpSmarthost

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertSMTPConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SmtpAuthPassword *string `json:"smtp_auth_password"`
		SmtpAuthUsername *string `json:"smtp_auth_username"`
		SmtpFrom         *string `json:"smtp_from"`
		SmtpSmarthost    *string `json:"smtp_smarthost"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SmtpAuthPassword == nil {
		return fmt.Errorf("required field smtp_auth_password missing")
	}
	if all.SmtpAuthUsername == nil {
		return fmt.Errorf("required field smtp_auth_username missing")
	}
	if all.SmtpFrom == nil {
		return fmt.Errorf("required field smtp_from missing")
	}
	if all.SmtpSmarthost == nil {
		return fmt.Errorf("required field smtp_smarthost missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"smtp_auth_password", "smtp_auth_username", "smtp_from", "smtp_smarthost"})
	} else {
		return err
	}
	o.SmtpAuthPassword = *all.SmtpAuthPassword
	o.SmtpAuthUsername = *all.SmtpAuthUsername
	o.SmtpFrom = *all.SmtpFrom
	o.SmtpSmarthost = *all.SmtpSmarthost

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
