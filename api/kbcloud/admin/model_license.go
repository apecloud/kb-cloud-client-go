// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "time"

// License License info
type License struct {
	// The kubernetes cluster ID
	ClusterId *string `json:"clusterID,omitempty"`
	// The licensed user email
	Email *string `json:"email,omitempty"`
	// The licensed user or organization name
	UserName *string `json:"userName,omitempty"`
	// The licensed unit, such as CPU or Node
	Unit *string `json:"unit,omitempty"`
	// The licensed total count of the unit
	Quantity *string `json:"quantity,omitempty"`
	// The supported engines
	Engines []string `json:"engines,omitempty"`
	// The license expiration time
	NotAfter *time.Time `json:"notAfter,omitempty"`
	// The license start time
	NotBefore *time.Time `json:"notBefore,omitempty"`
	// The used count of the unit
	Used *float64 `json:"used,omitempty"`
	// The license key
	License *string `json:"license,omitempty"`
	// The license mode
	Mode *string `json:"mode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicense instantiates a new License object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicense() *License {
	this := License{}
	return &this
}

// NewLicenseWithDefaults instantiates a new License object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *License) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *License) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *License) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *License) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *License) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *License) SetEmail(v string) {
	o.Email = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *License) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *License) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *License) SetUserName(v string) {
	o.UserName = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *License) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *License) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *License) SetUnit(v string) {
	o.Unit = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *License) GetQuantity() string {
	if o == nil || o.Quantity == nil {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetQuantityOk() (*string, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *License) HasQuantity() bool {
	return o != nil && o.Quantity != nil
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *License) SetQuantity(v string) {
	o.Quantity = &v
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *License) GetEngines() []string {
	if o == nil || o.Engines == nil {
		var ret []string
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetEnginesOk() (*[]string, bool) {
	if o == nil || o.Engines == nil {
		return nil, false
	}
	return &o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *License) HasEngines() bool {
	return o != nil && o.Engines != nil
}

// SetEngines gets a reference to the given []string and assigns it to the Engines field.
func (o *License) SetEngines(v []string) {
	o.Engines = v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *License) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *License) HasNotAfter() bool {
	return o != nil && o.NotAfter != nil
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *License) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *License) GetNotBefore() time.Time {
	if o == nil || o.NotBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *License) HasNotBefore() bool {
	return o != nil && o.NotBefore != nil
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *License) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *License) GetUsed() float64 {
	if o == nil || o.Used == nil {
		var ret float64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetUsedOk() (*float64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *License) HasUsed() bool {
	return o != nil && o.Used != nil
}

// SetUsed gets a reference to the given float64 and assigns it to the Used field.
func (o *License) SetUsed(v float64) {
	o.Used = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *License) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *License) HasLicense() bool {
	return o != nil && o.License != nil
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *License) SetLicense(v string) {
	o.License = &v
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
	if o.ClusterId != nil {
		toSerialize["clusterID"] = o.ClusterId
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Engines != nil {
		toSerialize["engines"] = o.Engines
	}
	if o.NotAfter != nil {
		if o.NotAfter.Nanosecond() == 0 {
			toSerialize["notAfter"] = o.NotAfter.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["notAfter"] = o.NotAfter.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.NotBefore != nil {
		if o.NotBefore.Nanosecond() == 0 {
			toSerialize["notBefore"] = o.NotBefore.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["notBefore"] = o.NotBefore.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
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
		ClusterId *string    `json:"clusterID,omitempty"`
		Email     *string    `json:"email,omitempty"`
		UserName  *string    `json:"userName,omitempty"`
		Unit      *string    `json:"unit,omitempty"`
		Quantity  *string    `json:"quantity,omitempty"`
		Engines   []string   `json:"engines,omitempty"`
		NotAfter  *time.Time `json:"notAfter,omitempty"`
		NotBefore *time.Time `json:"notBefore,omitempty"`
		Used      *float64   `json:"used,omitempty"`
		License   *string    `json:"license,omitempty"`
		Mode      *string    `json:"mode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterID", "email", "userName", "unit", "quantity", "engines", "notAfter", "notBefore", "used", "license", "mode"})
	} else {
		return err
	}
	o.ClusterId = all.ClusterId
	o.Email = all.Email
	o.UserName = all.UserName
	o.Unit = all.Unit
	o.Quantity = all.Quantity
	o.Engines = all.Engines
	o.NotAfter = all.NotAfter
	o.NotBefore = all.NotBefore
	o.Used = all.Used
	o.License = all.License
	o.Mode = all.Mode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
