// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// License License info
type License struct {
	// The kubernetes cluster ID
	ClusterId string `json:"clusterID"`
	// The licensed user email
	Email string `json:"email"`
	// The licensed user or organization name
	UserName string `json:"userName"`
	// The licensed unit, such as CPU or Node
	Unit string `json:"unit"`
	// The licensed total count of the unit
	Quantity string `json:"quantity"`
	// The supported engines and their quotas
	Engines []EngineQuota `json:"engines"`
	// The license expiration time
	NotAfter time.Time `json:"notAfter"`
	// The license start time
	NotBefore time.Time `json:"notBefore"`
	// The used count of the unit
	Used float64 `json:"used"`
	// The license key
	License string `json:"license"`
	// If KubeBlocks Enterprise is embedded in another platform, the license may be managed by the platform. It this case, we should extract the KubeBlocks Enterprise license from the platform and use it. Different platforms may have different ways to provide the license. The licenseMode is used to specify the different platform. For KubeSphere Enterprise, the licenseMode is `kse`.
	Mode *string `json:"mode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicense instantiates a new License object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicense(clusterId string, email string, userName string, unit string, quantity string, engines []EngineQuota, notAfter time.Time, notBefore time.Time, used float64, license string) *License {
	this := License{}
	this.ClusterId = clusterId
	this.Email = email
	this.UserName = userName
	this.Unit = unit
	this.Quantity = quantity
	this.Engines = engines
	this.NotAfter = notAfter
	this.NotBefore = notBefore
	this.Used = used
	this.License = license
	return &this
}

// NewLicenseWithDefaults instantiates a new License object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetClusterId returns the ClusterId field value.
func (o *License) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *License) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value.
func (o *License) SetClusterId(v string) {
	o.ClusterId = v
}

// GetEmail returns the Email field value.
func (o *License) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *License) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *License) SetEmail(v string) {
	o.Email = v
}

// GetUserName returns the UserName field value.
func (o *License) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *License) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *License) SetUserName(v string) {
	o.UserName = v
}

// GetUnit returns the Unit field value.
func (o *License) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *License) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *License) SetUnit(v string) {
	o.Unit = v
}

// GetQuantity returns the Quantity field value.
func (o *License) GetQuantity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *License) GetQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value.
func (o *License) SetQuantity(v string) {
	o.Quantity = v
}

// GetEngines returns the Engines field value.
func (o *License) GetEngines() []EngineQuota {
	if o == nil {
		var ret []EngineQuota
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value
// and a boolean to check if the value has been set.
func (o *License) GetEnginesOk() (*[]EngineQuota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engines, true
}

// SetEngines sets field value.
func (o *License) SetEngines(v []EngineQuota) {
	o.Engines = v
}

// GetNotAfter returns the NotAfter field value.
func (o *License) GetNotAfter() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value
// and a boolean to check if the value has been set.
func (o *License) GetNotAfterOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotAfter, true
}

// SetNotAfter sets field value.
func (o *License) SetNotAfter(v time.Time) {
	o.NotAfter = v
}

// GetNotBefore returns the NotBefore field value.
func (o *License) GetNotBefore() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value
// and a boolean to check if the value has been set.
func (o *License) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotBefore, true
}

// SetNotBefore sets field value.
func (o *License) SetNotBefore(v time.Time) {
	o.NotBefore = v
}

// GetUsed returns the Used field value.
func (o *License) GetUsed() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *License) GetUsedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value.
func (o *License) SetUsed(v float64) {
	o.Used = v
}

// GetLicense returns the License field value.
func (o *License) GetLicense() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *License) GetLicenseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value.
func (o *License) SetLicense(v string) {
	o.License = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *License) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *License) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *License) SetMode(v string) {
	o.Mode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o License) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["clusterID"] = o.ClusterId
	toSerialize["email"] = o.Email
	toSerialize["userName"] = o.UserName
	toSerialize["unit"] = o.Unit
	toSerialize["quantity"] = o.Quantity
	toSerialize["engines"] = o.Engines
	if o.NotAfter.Nanosecond() == 0 {
		toSerialize["notAfter"] = o.NotAfter.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["notAfter"] = o.NotAfter.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.NotBefore.Nanosecond() == 0 {
		toSerialize["notBefore"] = o.NotBefore.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["notBefore"] = o.NotBefore.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["used"] = o.Used
	toSerialize["license"] = o.License
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *License) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterId *string        `json:"clusterID"`
		Email     *string        `json:"email"`
		UserName  *string        `json:"userName"`
		Unit      *string        `json:"unit"`
		Quantity  *string        `json:"quantity"`
		Engines   *[]EngineQuota `json:"engines"`
		NotAfter  *time.Time     `json:"notAfter"`
		NotBefore *time.Time     `json:"notBefore"`
		Used      *float64       `json:"used"`
		License   *string        `json:"license"`
		Mode      *string        `json:"mode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ClusterId == nil {
		return fmt.Errorf("required field clusterID missing")
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Quantity == nil {
		return fmt.Errorf("required field quantity missing")
	}
	if all.Engines == nil {
		return fmt.Errorf("required field engines missing")
	}
	if all.NotAfter == nil {
		return fmt.Errorf("required field notAfter missing")
	}
	if all.NotBefore == nil {
		return fmt.Errorf("required field notBefore missing")
	}
	if all.Used == nil {
		return fmt.Errorf("required field used missing")
	}
	if all.License == nil {
		return fmt.Errorf("required field license missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterID", "email", "userName", "unit", "quantity", "engines", "notAfter", "notBefore", "used", "license", "mode"})
	} else {
		return err
	}
	o.ClusterId = *all.ClusterId
	o.Email = *all.Email
	o.UserName = *all.UserName
	o.Unit = *all.Unit
	o.Quantity = *all.Quantity
	o.Engines = *all.Engines
	o.NotAfter = *all.NotAfter
	o.NotBefore = *all.NotBefore
	o.Used = *all.Used
	o.License = *all.License
	o.Mode = all.Mode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
