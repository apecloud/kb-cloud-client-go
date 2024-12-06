// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// EngineVersionCreate EngineVersionRecord create option
type EngineVersionCreate struct {
	// Name of the engine
	EngineName string `json:"engineName"`
	// Version of the engine
	Version string `json:"version"`
	// Version constraint for KB
	KbVersionConstraint string `json:"kbVersionConstraint"`
	// URL for the cluster chart
	ClusterChartUrl common.NullableString `json:"clusterChartUrl,omitempty"`
	// URL for the chart
	ChartUrl string `json:"chartUrl"`
	// Configuration values set for the engine
	SetValues common.NullableString `json:"setValues,omitempty"`
	// Image associated with the charts
	ChartsImage common.NullableString `json:"chartsImage,omitempty"`
	// Determines if the image registry is set
	SetImageRegistry common.NullableBool `json:"setImageRegistry,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineVersionCreate instantiates a new EngineVersionCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineVersionCreate(engineName string, version string, kbVersionConstraint string, chartUrl string) *EngineVersionCreate {
	this := EngineVersionCreate{}
	this.EngineName = engineName
	this.Version = version
	this.KbVersionConstraint = kbVersionConstraint
	this.ChartUrl = chartUrl
	var setImageRegistry bool = true
	this.SetImageRegistry = *common.NewNullableBool(&setImageRegistry)
	return &this
}

// NewEngineVersionCreateWithDefaults instantiates a new EngineVersionCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineVersionCreateWithDefaults() *EngineVersionCreate {
	this := EngineVersionCreate{}
	var setImageRegistry bool = true
	this.SetImageRegistry = *common.NewNullableBool(&setImageRegistry)
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *EngineVersionCreate) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineVersionCreate) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineVersionCreate) SetEngineName(v string) {
	o.EngineName = v
}

// GetVersion returns the Version field value.
func (o *EngineVersionCreate) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EngineVersionCreate) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *EngineVersionCreate) SetVersion(v string) {
	o.Version = v
}

// GetKbVersionConstraint returns the KbVersionConstraint field value.
func (o *EngineVersionCreate) GetKbVersionConstraint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.KbVersionConstraint
}

// GetKbVersionConstraintOk returns a tuple with the KbVersionConstraint field value
// and a boolean to check if the value has been set.
func (o *EngineVersionCreate) GetKbVersionConstraintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KbVersionConstraint, true
}

// SetKbVersionConstraint sets field value.
func (o *EngineVersionCreate) SetKbVersionConstraint(v string) {
	o.KbVersionConstraint = v
}

// GetClusterChartUrl returns the ClusterChartUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionCreate) GetClusterChartUrl() string {
	if o == nil || o.ClusterChartUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterChartUrl.Get()
}

// GetClusterChartUrlOk returns a tuple with the ClusterChartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionCreate) GetClusterChartUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterChartUrl.Get(), o.ClusterChartUrl.IsSet()
}

// HasClusterChartUrl returns a boolean if a field has been set.
func (o *EngineVersionCreate) HasClusterChartUrl() bool {
	return o != nil && o.ClusterChartUrl.IsSet()
}

// SetClusterChartUrl gets a reference to the given common.NullableString and assigns it to the ClusterChartUrl field.
func (o *EngineVersionCreate) SetClusterChartUrl(v string) {
	o.ClusterChartUrl.Set(&v)
}

// SetClusterChartUrlNil sets the value for ClusterChartUrl to be an explicit nil.
func (o *EngineVersionCreate) SetClusterChartUrlNil() {
	o.ClusterChartUrl.Set(nil)
}

// UnsetClusterChartUrl ensures that no value is present for ClusterChartUrl, not even an explicit nil.
func (o *EngineVersionCreate) UnsetClusterChartUrl() {
	o.ClusterChartUrl.Unset()
}

// GetChartUrl returns the ChartUrl field value.
func (o *EngineVersionCreate) GetChartUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChartUrl
}

// GetChartUrlOk returns a tuple with the ChartUrl field value
// and a boolean to check if the value has been set.
func (o *EngineVersionCreate) GetChartUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartUrl, true
}

// SetChartUrl sets field value.
func (o *EngineVersionCreate) SetChartUrl(v string) {
	o.ChartUrl = v
}

// GetSetValues returns the SetValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionCreate) GetSetValues() string {
	if o == nil || o.SetValues.Get() == nil {
		var ret string
		return ret
	}
	return *o.SetValues.Get()
}

// GetSetValuesOk returns a tuple with the SetValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionCreate) GetSetValuesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SetValues.Get(), o.SetValues.IsSet()
}

// HasSetValues returns a boolean if a field has been set.
func (o *EngineVersionCreate) HasSetValues() bool {
	return o != nil && o.SetValues.IsSet()
}

// SetSetValues gets a reference to the given common.NullableString and assigns it to the SetValues field.
func (o *EngineVersionCreate) SetSetValues(v string) {
	o.SetValues.Set(&v)
}

// SetSetValuesNil sets the value for SetValues to be an explicit nil.
func (o *EngineVersionCreate) SetSetValuesNil() {
	o.SetValues.Set(nil)
}

// UnsetSetValues ensures that no value is present for SetValues, not even an explicit nil.
func (o *EngineVersionCreate) UnsetSetValues() {
	o.SetValues.Unset()
}

// GetChartsImage returns the ChartsImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionCreate) GetChartsImage() string {
	if o == nil || o.ChartsImage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChartsImage.Get()
}

// GetChartsImageOk returns a tuple with the ChartsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionCreate) GetChartsImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChartsImage.Get(), o.ChartsImage.IsSet()
}

// HasChartsImage returns a boolean if a field has been set.
func (o *EngineVersionCreate) HasChartsImage() bool {
	return o != nil && o.ChartsImage.IsSet()
}

// SetChartsImage gets a reference to the given common.NullableString and assigns it to the ChartsImage field.
func (o *EngineVersionCreate) SetChartsImage(v string) {
	o.ChartsImage.Set(&v)
}

// SetChartsImageNil sets the value for ChartsImage to be an explicit nil.
func (o *EngineVersionCreate) SetChartsImageNil() {
	o.ChartsImage.Set(nil)
}

// UnsetChartsImage ensures that no value is present for ChartsImage, not even an explicit nil.
func (o *EngineVersionCreate) UnsetChartsImage() {
	o.ChartsImage.Unset()
}

// GetSetImageRegistry returns the SetImageRegistry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionCreate) GetSetImageRegistry() bool {
	if o == nil || o.SetImageRegistry.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SetImageRegistry.Get()
}

// GetSetImageRegistryOk returns a tuple with the SetImageRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionCreate) GetSetImageRegistryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SetImageRegistry.Get(), o.SetImageRegistry.IsSet()
}

// HasSetImageRegistry returns a boolean if a field has been set.
func (o *EngineVersionCreate) HasSetImageRegistry() bool {
	return o != nil && o.SetImageRegistry.IsSet()
}

// SetSetImageRegistry gets a reference to the given common.NullableBool and assigns it to the SetImageRegistry field.
func (o *EngineVersionCreate) SetSetImageRegistry(v bool) {
	o.SetImageRegistry.Set(&v)
}

// SetSetImageRegistryNil sets the value for SetImageRegistry to be an explicit nil.
func (o *EngineVersionCreate) SetSetImageRegistryNil() {
	o.SetImageRegistry.Set(nil)
}

// UnsetSetImageRegistry ensures that no value is present for SetImageRegistry, not even an explicit nil.
func (o *EngineVersionCreate) UnsetSetImageRegistry() {
	o.SetImageRegistry.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineVersionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["version"] = o.Version
	toSerialize["kbVersionConstraint"] = o.KbVersionConstraint
	if o.ClusterChartUrl.IsSet() {
		toSerialize["clusterChartUrl"] = o.ClusterChartUrl.Get()
	}
	toSerialize["chartUrl"] = o.ChartUrl
	if o.SetValues.IsSet() {
		toSerialize["setValues"] = o.SetValues.Get()
	}
	if o.ChartsImage.IsSet() {
		toSerialize["chartsImage"] = o.ChartsImage.Get()
	}
	if o.SetImageRegistry.IsSet() {
		toSerialize["setImageRegistry"] = o.SetImageRegistry.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineVersionCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName          *string               `json:"engineName"`
		Version             *string               `json:"version"`
		KbVersionConstraint *string               `json:"kbVersionConstraint"`
		ClusterChartUrl     common.NullableString `json:"clusterChartUrl,omitempty"`
		ChartUrl            *string               `json:"chartUrl"`
		SetValues           common.NullableString `json:"setValues,omitempty"`
		ChartsImage         common.NullableString `json:"chartsImage,omitempty"`
		SetImageRegistry    common.NullableBool   `json:"setImageRegistry,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	if all.KbVersionConstraint == nil {
		return fmt.Errorf("required field kbVersionConstraint missing")
	}
	if all.ChartUrl == nil {
		return fmt.Errorf("required field chartUrl missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "version", "kbVersionConstraint", "clusterChartUrl", "chartUrl", "setValues", "chartsImage", "setImageRegistry"})
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Version = *all.Version
	o.KbVersionConstraint = *all.KbVersionConstraint
	o.ClusterChartUrl = all.ClusterChartUrl
	o.ChartUrl = *all.ChartUrl
	o.SetValues = all.SetValues
	o.ChartsImage = all.ChartsImage
	o.SetImageRegistry = all.SetImageRegistry

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
