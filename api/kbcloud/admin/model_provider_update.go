// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ProviderUpdate The cloud provider that needs to be updated.

type ProviderUpdate struct {
	// The logo of the cloud provider.
	Logo string `json:"logo"`
	// Whether the cloud provider is enabled.
	Enabled bool `json:"enabled"`
	// The Chinese name of the cloud provider.
	NameCn string `json:"nameCN"`
	// The English name of the cloud provider.
	NameEn string `json:"nameEN"`
	// Whether the cloud provider supports ARN.
	SupportArn bool `json:"supportARN"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProviderUpdate instantiates a new ProviderUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProviderUpdate(logo string, enabled bool, nameCn string, nameEn string, supportArn bool) *ProviderUpdate {
	this := ProviderUpdate{}
	this.Logo = logo
	this.Enabled = enabled
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.SupportArn = supportArn
	return &this
}

// NewProviderUpdateWithDefaults instantiates a new ProviderUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProviderUpdateWithDefaults() *ProviderUpdate {
	this := ProviderUpdate{}
	return &this
}

// GetLogo returns the Logo field value.
func (o *ProviderUpdate) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *ProviderUpdate) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value.
func (o *ProviderUpdate) SetLogo(v string) {
	o.Logo = v
}

// GetEnabled returns the Enabled field value.
func (o *ProviderUpdate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProviderUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ProviderUpdate) SetEnabled(v bool) {
	o.Enabled = v
}

// GetNameCn returns the NameCn field value.
func (o *ProviderUpdate) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *ProviderUpdate) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *ProviderUpdate) SetNameCn(v string) {
	o.NameCn = v
}

// GetNameEn returns the NameEn field value.
func (o *ProviderUpdate) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *ProviderUpdate) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *ProviderUpdate) SetNameEn(v string) {
	o.NameEn = v
}

// GetSupportArn returns the SupportArn field value.
func (o *ProviderUpdate) GetSupportArn() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SupportArn
}

// GetSupportArnOk returns a tuple with the SupportArn field value
// and a boolean to check if the value has been set.
func (o *ProviderUpdate) GetSupportArnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportArn, true
}

// SetSupportArn sets field value.
func (o *ProviderUpdate) SetSupportArn(v bool) {
	o.SupportArn = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProviderUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["logo"] = o.Logo
	toSerialize["enabled"] = o.Enabled
	toSerialize["nameCN"] = o.NameCn
	toSerialize["nameEN"] = o.NameEn
	toSerialize["supportARN"] = o.SupportArn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProviderUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Logo       *string `json:"logo"`
		Enabled    *bool   `json:"enabled"`
		NameCn     *string `json:"nameCN"`
		NameEn     *string `json:"nameEN"`
		SupportArn *bool   `json:"supportARN"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Logo == nil {
		return fmt.Errorf("required field logo missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.NameCn == nil {
		return fmt.Errorf("required field nameCN missing")
	}
	if all.NameEn == nil {
		return fmt.Errorf("required field nameEN missing")
	}
	if all.SupportArn == nil {
		return fmt.Errorf("required field supportARN missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"logo", "enabled", "nameCN", "nameEN", "supportARN"})
	} else {
		return err
	}
	o.Logo = *all.Logo
	o.Enabled = *all.Enabled
	o.NameCn = *all.NameCn
	o.NameEn = *all.NameEn
	o.SupportArn = *all.SupportArn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
