// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineVersionUpdate EngineVersionRecord update option
type EngineVersionUpdate struct {
	// Name of the engine
	EngineName string `json:"engineName"`
	// Version of the engine
	Version string `json:"version"`
	// Version constraint for KB
	KbVersionConstraint common.NullableString `json:"kbVersionConstraint,omitempty"`
	// URL for the cluster chart
	ClusterChartUrl common.NullableString `json:"clusterChartUrl,omitempty"`
	// URL for the chart
	ChartUrl common.NullableString `json:"chartUrl,omitempty"`
	// Configuration values set for the engine
	SetValues common.NullableString `json:"setValues,omitempty"`
	// Image associated with the charts
	ChartsImage common.NullableString `json:"chartsImage,omitempty"`
	// Determines if the image registry is set
	SetImageRegistry common.NullableBool `json:"setImageRegistry,omitempty"`
	// Service versions supported by this engine version
	ServiceVersions common.NullableList[string] `json:"serviceVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineVersionUpdate instantiates a new EngineVersionUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineVersionUpdate(engineName string, version string) *EngineVersionUpdate {
	this := EngineVersionUpdate{}
	this.EngineName = engineName
	this.Version = version
	var setImageRegistry bool = true
	this.SetImageRegistry = *common.NewNullableBool(&setImageRegistry)
	return &this
}

// NewEngineVersionUpdateWithDefaults instantiates a new EngineVersionUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineVersionUpdateWithDefaults() *EngineVersionUpdate {
	this := EngineVersionUpdate{}
	var setImageRegistry bool = true
	this.SetImageRegistry = *common.NewNullableBool(&setImageRegistry)
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *EngineVersionUpdate) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineVersionUpdate) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineVersionUpdate) SetEngineName(v string) {
	o.EngineName = v
}

// GetVersion returns the Version field value.
func (o *EngineVersionUpdate) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EngineVersionUpdate) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *EngineVersionUpdate) SetVersion(v string) {
	o.Version = v
}

// GetKbVersionConstraint returns the KbVersionConstraint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionUpdate) GetKbVersionConstraint() string {
	if o == nil || o.KbVersionConstraint.Get() == nil {
		var ret string
		return ret
	}
	return *o.KbVersionConstraint.Get()
}

// GetKbVersionConstraintOk returns a tuple with the KbVersionConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionUpdate) GetKbVersionConstraintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KbVersionConstraint.Get(), o.KbVersionConstraint.IsSet()
}

// HasKbVersionConstraint returns a boolean if a field has been set.
func (o *EngineVersionUpdate) HasKbVersionConstraint() bool {
	return o != nil && o.KbVersionConstraint.IsSet()
}

// SetKbVersionConstraint gets a reference to the given common.NullableString and assigns it to the KbVersionConstraint field.
func (o *EngineVersionUpdate) SetKbVersionConstraint(v string) {
	o.KbVersionConstraint.Set(&v)
}

// SetKbVersionConstraintNil sets the value for KbVersionConstraint to be an explicit nil.
func (o *EngineVersionUpdate) SetKbVersionConstraintNil() {
	o.KbVersionConstraint.Set(nil)
}

// UnsetKbVersionConstraint ensures that no value is present for KbVersionConstraint, not even an explicit nil.
func (o *EngineVersionUpdate) UnsetKbVersionConstraint() {
	o.KbVersionConstraint.Unset()
}

// GetClusterChartUrl returns the ClusterChartUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionUpdate) GetClusterChartUrl() string {
	if o == nil || o.ClusterChartUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterChartUrl.Get()
}

// GetClusterChartUrlOk returns a tuple with the ClusterChartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionUpdate) GetClusterChartUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterChartUrl.Get(), o.ClusterChartUrl.IsSet()
}

// HasClusterChartUrl returns a boolean if a field has been set.
func (o *EngineVersionUpdate) HasClusterChartUrl() bool {
	return o != nil && o.ClusterChartUrl.IsSet()
}

// SetClusterChartUrl gets a reference to the given common.NullableString and assigns it to the ClusterChartUrl field.
func (o *EngineVersionUpdate) SetClusterChartUrl(v string) {
	o.ClusterChartUrl.Set(&v)
}

// SetClusterChartUrlNil sets the value for ClusterChartUrl to be an explicit nil.
func (o *EngineVersionUpdate) SetClusterChartUrlNil() {
	o.ClusterChartUrl.Set(nil)
}

// UnsetClusterChartUrl ensures that no value is present for ClusterChartUrl, not even an explicit nil.
func (o *EngineVersionUpdate) UnsetClusterChartUrl() {
	o.ClusterChartUrl.Unset()
}

// GetChartUrl returns the ChartUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionUpdate) GetChartUrl() string {
	if o == nil || o.ChartUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChartUrl.Get()
}

// GetChartUrlOk returns a tuple with the ChartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionUpdate) GetChartUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChartUrl.Get(), o.ChartUrl.IsSet()
}

// HasChartUrl returns a boolean if a field has been set.
func (o *EngineVersionUpdate) HasChartUrl() bool {
	return o != nil && o.ChartUrl.IsSet()
}

// SetChartUrl gets a reference to the given common.NullableString and assigns it to the ChartUrl field.
func (o *EngineVersionUpdate) SetChartUrl(v string) {
	o.ChartUrl.Set(&v)
}

// SetChartUrlNil sets the value for ChartUrl to be an explicit nil.
func (o *EngineVersionUpdate) SetChartUrlNil() {
	o.ChartUrl.Set(nil)
}

// UnsetChartUrl ensures that no value is present for ChartUrl, not even an explicit nil.
func (o *EngineVersionUpdate) UnsetChartUrl() {
	o.ChartUrl.Unset()
}

// GetSetValues returns the SetValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionUpdate) GetSetValues() string {
	if o == nil || o.SetValues.Get() == nil {
		var ret string
		return ret
	}
	return *o.SetValues.Get()
}

// GetSetValuesOk returns a tuple with the SetValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionUpdate) GetSetValuesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SetValues.Get(), o.SetValues.IsSet()
}

// HasSetValues returns a boolean if a field has been set.
func (o *EngineVersionUpdate) HasSetValues() bool {
	return o != nil && o.SetValues.IsSet()
}

// SetSetValues gets a reference to the given common.NullableString and assigns it to the SetValues field.
func (o *EngineVersionUpdate) SetSetValues(v string) {
	o.SetValues.Set(&v)
}

// SetSetValuesNil sets the value for SetValues to be an explicit nil.
func (o *EngineVersionUpdate) SetSetValuesNil() {
	o.SetValues.Set(nil)
}

// UnsetSetValues ensures that no value is present for SetValues, not even an explicit nil.
func (o *EngineVersionUpdate) UnsetSetValues() {
	o.SetValues.Unset()
}

// GetChartsImage returns the ChartsImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionUpdate) GetChartsImage() string {
	if o == nil || o.ChartsImage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChartsImage.Get()
}

// GetChartsImageOk returns a tuple with the ChartsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionUpdate) GetChartsImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChartsImage.Get(), o.ChartsImage.IsSet()
}

// HasChartsImage returns a boolean if a field has been set.
func (o *EngineVersionUpdate) HasChartsImage() bool {
	return o != nil && o.ChartsImage.IsSet()
}

// SetChartsImage gets a reference to the given common.NullableString and assigns it to the ChartsImage field.
func (o *EngineVersionUpdate) SetChartsImage(v string) {
	o.ChartsImage.Set(&v)
}

// SetChartsImageNil sets the value for ChartsImage to be an explicit nil.
func (o *EngineVersionUpdate) SetChartsImageNil() {
	o.ChartsImage.Set(nil)
}

// UnsetChartsImage ensures that no value is present for ChartsImage, not even an explicit nil.
func (o *EngineVersionUpdate) UnsetChartsImage() {
	o.ChartsImage.Unset()
}

// GetSetImageRegistry returns the SetImageRegistry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionUpdate) GetSetImageRegistry() bool {
	if o == nil || o.SetImageRegistry.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SetImageRegistry.Get()
}

// GetSetImageRegistryOk returns a tuple with the SetImageRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionUpdate) GetSetImageRegistryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SetImageRegistry.Get(), o.SetImageRegistry.IsSet()
}

// HasSetImageRegistry returns a boolean if a field has been set.
func (o *EngineVersionUpdate) HasSetImageRegistry() bool {
	return o != nil && o.SetImageRegistry.IsSet()
}

// SetSetImageRegistry gets a reference to the given common.NullableBool and assigns it to the SetImageRegistry field.
func (o *EngineVersionUpdate) SetSetImageRegistry(v bool) {
	o.SetImageRegistry.Set(&v)
}

// SetSetImageRegistryNil sets the value for SetImageRegistry to be an explicit nil.
func (o *EngineVersionUpdate) SetSetImageRegistryNil() {
	o.SetImageRegistry.Set(nil)
}

// UnsetSetImageRegistry ensures that no value is present for SetImageRegistry, not even an explicit nil.
func (o *EngineVersionUpdate) UnsetSetImageRegistry() {
	o.SetImageRegistry.Unset()
}

// GetServiceVersions returns the ServiceVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersionUpdate) GetServiceVersions() []string {
	if o == nil || o.ServiceVersions.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ServiceVersions.Get()
}

// GetServiceVersionsOk returns a tuple with the ServiceVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersionUpdate) GetServiceVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceVersions.Get(), o.ServiceVersions.IsSet()
}

// HasServiceVersions returns a boolean if a field has been set.
func (o *EngineVersionUpdate) HasServiceVersions() bool {
	return o != nil && o.ServiceVersions.IsSet()
}

// SetServiceVersions gets a reference to the given common.NullableList[string] and assigns it to the ServiceVersions field.
func (o *EngineVersionUpdate) SetServiceVersions(v []string) {
	o.ServiceVersions.Set(&v)
}

// SetServiceVersionsNil sets the value for ServiceVersions to be an explicit nil.
func (o *EngineVersionUpdate) SetServiceVersionsNil() {
	o.ServiceVersions.Set(nil)
}

// UnsetServiceVersions ensures that no value is present for ServiceVersions, not even an explicit nil.
func (o *EngineVersionUpdate) UnsetServiceVersions() {
	o.ServiceVersions.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineVersionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["version"] = o.Version
	if o.KbVersionConstraint.IsSet() {
		toSerialize["kbVersionConstraint"] = o.KbVersionConstraint.Get()
	}
	if o.ClusterChartUrl.IsSet() {
		toSerialize["clusterChartUrl"] = o.ClusterChartUrl.Get()
	}
	if o.ChartUrl.IsSet() {
		toSerialize["chartUrl"] = o.ChartUrl.Get()
	}
	if o.SetValues.IsSet() {
		toSerialize["setValues"] = o.SetValues.Get()
	}
	if o.ChartsImage.IsSet() {
		toSerialize["chartsImage"] = o.ChartsImage.Get()
	}
	if o.SetImageRegistry.IsSet() {
		toSerialize["setImageRegistry"] = o.SetImageRegistry.Get()
	}
	if o.ServiceVersions.IsSet() {
		toSerialize["serviceVersions"] = o.ServiceVersions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineVersionUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName          *string                     `json:"engineName"`
		Version             *string                     `json:"version"`
		KbVersionConstraint common.NullableString       `json:"kbVersionConstraint,omitempty"`
		ClusterChartUrl     common.NullableString       `json:"clusterChartUrl,omitempty"`
		ChartUrl            common.NullableString       `json:"chartUrl,omitempty"`
		SetValues           common.NullableString       `json:"setValues,omitempty"`
		ChartsImage         common.NullableString       `json:"chartsImage,omitempty"`
		SetImageRegistry    common.NullableBool         `json:"setImageRegistry,omitempty"`
		ServiceVersions     common.NullableList[string] `json:"serviceVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "version", "kbVersionConstraint", "clusterChartUrl", "chartUrl", "setValues", "chartsImage", "setImageRegistry", "serviceVersions"})
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Version = *all.Version
	o.KbVersionConstraint = all.KbVersionConstraint
	o.ClusterChartUrl = all.ClusterChartUrl
	o.ChartUrl = all.ChartUrl
	o.SetValues = all.SetValues
	o.ChartsImage = all.ChartsImage
	o.SetImageRegistry = all.SetImageRegistry
	o.ServiceVersions = all.ServiceVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
