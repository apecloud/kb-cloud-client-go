// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ProviderCreate The cloud provider that needs to be created.
type ProviderCreate struct {
	// The name of the cloud provider.
	Name string `json:"name"`
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

// NewProviderCreate instantiates a new ProviderCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProviderCreate(name string, logo string, enabled bool, nameCn string, nameEn string, supportArn bool) *ProviderCreate {
	this := ProviderCreate{}
	this.Name = name
	this.Logo = logo
	this.Enabled = enabled
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.SupportArn = supportArn
	return &this
}

// NewProviderCreateWithDefaults instantiates a new ProviderCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProviderCreateWithDefaults() *ProviderCreate {
	this := ProviderCreate{}
	return &this
}

// GetName returns the Name field value.
func (o *ProviderCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProviderCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProviderCreate) SetName(v string) {
	o.Name = v
}

// GetLogo returns the Logo field value.
func (o *ProviderCreate) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *ProviderCreate) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value.
func (o *ProviderCreate) SetLogo(v string) {
	o.Logo = v
}

// GetEnabled returns the Enabled field value.
func (o *ProviderCreate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ProviderCreate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ProviderCreate) SetEnabled(v bool) {
	o.Enabled = v
}

// GetNameCn returns the NameCn field value.
func (o *ProviderCreate) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *ProviderCreate) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *ProviderCreate) SetNameCn(v string) {
	o.NameCn = v
}

// GetNameEn returns the NameEn field value.
func (o *ProviderCreate) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *ProviderCreate) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *ProviderCreate) SetNameEn(v string) {
	o.NameEn = v
}

// GetSupportArn returns the SupportArn field value.
func (o *ProviderCreate) GetSupportArn() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SupportArn
}

// GetSupportArnOk returns a tuple with the SupportArn field value
// and a boolean to check if the value has been set.
func (o *ProviderCreate) GetSupportArnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportArn, true
}

// SetSupportArn sets field value.
func (o *ProviderCreate) SetSupportArn(v bool) {
	o.SupportArn = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProviderCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
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
func (o *ProviderCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string `json:"name"`
		Logo       *string `json:"logo"`
		Enabled    *bool   `json:"enabled"`
		NameCn     *string `json:"nameCN"`
		NameEn     *string `json:"nameEN"`
		SupportArn *bool   `json:"supportARN"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"name", "logo", "enabled", "nameCN", "nameEN", "supportARN"})
	} else {
		return err
	}
	o.Name = *all.Name
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
