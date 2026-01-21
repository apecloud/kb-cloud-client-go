// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineVersion EngineVersionRecord
type EngineVersion struct {
	// Primary Key for the EngineVersionRecord
	Id string `json:"id"`
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
	// Service versions supported by this engine version
	ServiceVersions []string `json:"serviceVersions"`
	// Timestamp when the record was created
	CreatedAt common.NullableTime `json:"createdAt,omitempty"`
	// Timestamp when the record was last updated
	UpdatedAt common.NullableTime `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineVersion instantiates a new EngineVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineVersion(id string, engineName string, version string, kbVersionConstraint string, chartUrl string, serviceVersions []string) *EngineVersion {
	this := EngineVersion{}
	this.Id = id
	this.EngineName = engineName
	this.Version = version
	this.KbVersionConstraint = kbVersionConstraint
	this.ChartUrl = chartUrl
	var setImageRegistry bool = true
	this.SetImageRegistry = *common.NewNullableBool(&setImageRegistry)
	this.ServiceVersions = serviceVersions
	return &this
}

// NewEngineVersionWithDefaults instantiates a new EngineVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineVersionWithDefaults() *EngineVersion {
	this := EngineVersion{}
	var setImageRegistry bool = true
	this.SetImageRegistry = *common.NewNullableBool(&setImageRegistry)
	return &this
}

// GetId returns the Id field value.
func (o *EngineVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EngineVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *EngineVersion) SetId(v string) {
	o.Id = v
}

// GetEngineName returns the EngineName field value.
func (o *EngineVersion) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineVersion) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineVersion) SetEngineName(v string) {
	o.EngineName = v
}

// GetVersion returns the Version field value.
func (o *EngineVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EngineVersion) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *EngineVersion) SetVersion(v string) {
	o.Version = v
}

// GetKbVersionConstraint returns the KbVersionConstraint field value.
func (o *EngineVersion) GetKbVersionConstraint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.KbVersionConstraint
}

// GetKbVersionConstraintOk returns a tuple with the KbVersionConstraint field value
// and a boolean to check if the value has been set.
func (o *EngineVersion) GetKbVersionConstraintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KbVersionConstraint, true
}

// SetKbVersionConstraint sets field value.
func (o *EngineVersion) SetKbVersionConstraint(v string) {
	o.KbVersionConstraint = v
}

// GetClusterChartUrl returns the ClusterChartUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersion) GetClusterChartUrl() string {
	if o == nil || o.ClusterChartUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterChartUrl.Get()
}

// GetClusterChartUrlOk returns a tuple with the ClusterChartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersion) GetClusterChartUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterChartUrl.Get(), o.ClusterChartUrl.IsSet()
}

// HasClusterChartUrl returns a boolean if a field has been set.
func (o *EngineVersion) HasClusterChartUrl() bool {
	return o != nil && o.ClusterChartUrl.IsSet()
}

// SetClusterChartUrl gets a reference to the given common.NullableString and assigns it to the ClusterChartUrl field.
func (o *EngineVersion) SetClusterChartUrl(v string) {
	o.ClusterChartUrl.Set(&v)
}

// SetClusterChartUrlNil sets the value for ClusterChartUrl to be an explicit nil.
func (o *EngineVersion) SetClusterChartUrlNil() {
	o.ClusterChartUrl.Set(nil)
}

// UnsetClusterChartUrl ensures that no value is present for ClusterChartUrl, not even an explicit nil.
func (o *EngineVersion) UnsetClusterChartUrl() {
	o.ClusterChartUrl.Unset()
}

// GetChartUrl returns the ChartUrl field value.
func (o *EngineVersion) GetChartUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChartUrl
}

// GetChartUrlOk returns a tuple with the ChartUrl field value
// and a boolean to check if the value has been set.
func (o *EngineVersion) GetChartUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartUrl, true
}

// SetChartUrl sets field value.
func (o *EngineVersion) SetChartUrl(v string) {
	o.ChartUrl = v
}

// GetSetValues returns the SetValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersion) GetSetValues() string {
	if o == nil || o.SetValues.Get() == nil {
		var ret string
		return ret
	}
	return *o.SetValues.Get()
}

// GetSetValuesOk returns a tuple with the SetValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersion) GetSetValuesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SetValues.Get(), o.SetValues.IsSet()
}

// HasSetValues returns a boolean if a field has been set.
func (o *EngineVersion) HasSetValues() bool {
	return o != nil && o.SetValues.IsSet()
}

// SetSetValues gets a reference to the given common.NullableString and assigns it to the SetValues field.
func (o *EngineVersion) SetSetValues(v string) {
	o.SetValues.Set(&v)
}

// SetSetValuesNil sets the value for SetValues to be an explicit nil.
func (o *EngineVersion) SetSetValuesNil() {
	o.SetValues.Set(nil)
}

// UnsetSetValues ensures that no value is present for SetValues, not even an explicit nil.
func (o *EngineVersion) UnsetSetValues() {
	o.SetValues.Unset()
}

// GetChartsImage returns the ChartsImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersion) GetChartsImage() string {
	if o == nil || o.ChartsImage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChartsImage.Get()
}

// GetChartsImageOk returns a tuple with the ChartsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersion) GetChartsImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChartsImage.Get(), o.ChartsImage.IsSet()
}

// HasChartsImage returns a boolean if a field has been set.
func (o *EngineVersion) HasChartsImage() bool {
	return o != nil && o.ChartsImage.IsSet()
}

// SetChartsImage gets a reference to the given common.NullableString and assigns it to the ChartsImage field.
func (o *EngineVersion) SetChartsImage(v string) {
	o.ChartsImage.Set(&v)
}

// SetChartsImageNil sets the value for ChartsImage to be an explicit nil.
func (o *EngineVersion) SetChartsImageNil() {
	o.ChartsImage.Set(nil)
}

// UnsetChartsImage ensures that no value is present for ChartsImage, not even an explicit nil.
func (o *EngineVersion) UnsetChartsImage() {
	o.ChartsImage.Unset()
}

// GetSetImageRegistry returns the SetImageRegistry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersion) GetSetImageRegistry() bool {
	if o == nil || o.SetImageRegistry.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SetImageRegistry.Get()
}

// GetSetImageRegistryOk returns a tuple with the SetImageRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersion) GetSetImageRegistryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SetImageRegistry.Get(), o.SetImageRegistry.IsSet()
}

// HasSetImageRegistry returns a boolean if a field has been set.
func (o *EngineVersion) HasSetImageRegistry() bool {
	return o != nil && o.SetImageRegistry.IsSet()
}

// SetSetImageRegistry gets a reference to the given common.NullableBool and assigns it to the SetImageRegistry field.
func (o *EngineVersion) SetSetImageRegistry(v bool) {
	o.SetImageRegistry.Set(&v)
}

// SetSetImageRegistryNil sets the value for SetImageRegistry to be an explicit nil.
func (o *EngineVersion) SetSetImageRegistryNil() {
	o.SetImageRegistry.Set(nil)
}

// UnsetSetImageRegistry ensures that no value is present for SetImageRegistry, not even an explicit nil.
func (o *EngineVersion) UnsetSetImageRegistry() {
	o.SetImageRegistry.Unset()
}

// GetServiceVersions returns the ServiceVersions field value.
func (o *EngineVersion) GetServiceVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ServiceVersions
}

// GetServiceVersionsOk returns a tuple with the ServiceVersions field value
// and a boolean to check if the value has been set.
func (o *EngineVersion) GetServiceVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceVersions, true
}

// SetServiceVersions sets field value.
func (o *EngineVersion) SetServiceVersions(v []string) {
	o.ServiceVersions = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersion) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EngineVersion) HasCreatedAt() bool {
	return o != nil && o.CreatedAt.IsSet()
}

// SetCreatedAt gets a reference to the given common.NullableTime and assigns it to the CreatedAt field.
func (o *EngineVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil.
func (o *EngineVersion) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil.
func (o *EngineVersion) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineVersion) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineVersion) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EngineVersion) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt.IsSet()
}

// SetUpdatedAt gets a reference to the given common.NullableTime and assigns it to the UpdatedAt field.
func (o *EngineVersion) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil.
func (o *EngineVersion) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil.
func (o *EngineVersion) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
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
	toSerialize["serviceVersions"] = o.ServiceVersions
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                  *string               `json:"id"`
		EngineName          *string               `json:"engineName"`
		Version             *string               `json:"version"`
		KbVersionConstraint *string               `json:"kbVersionConstraint"`
		ClusterChartUrl     common.NullableString `json:"clusterChartUrl,omitempty"`
		ChartUrl            *string               `json:"chartUrl"`
		SetValues           common.NullableString `json:"setValues,omitempty"`
		ChartsImage         common.NullableString `json:"chartsImage,omitempty"`
		SetImageRegistry    common.NullableBool   `json:"setImageRegistry,omitempty"`
		ServiceVersions     *[]string             `json:"serviceVersions"`
		CreatedAt           common.NullableTime   `json:"createdAt,omitempty"`
		UpdatedAt           common.NullableTime   `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	if all.ServiceVersions == nil {
		return fmt.Errorf("required field serviceVersions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "engineName", "version", "kbVersionConstraint", "clusterChartUrl", "chartUrl", "setValues", "chartsImage", "setImageRegistry", "serviceVersions", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.EngineName = *all.EngineName
	o.Version = *all.Version
	o.KbVersionConstraint = *all.KbVersionConstraint
	o.ClusterChartUrl = all.ClusterChartUrl
	o.ChartUrl = *all.ChartUrl
	o.SetValues = all.SetValues
	o.ChartsImage = all.ChartsImage
	o.SetImageRegistry = all.SetImageRegistry
	o.ServiceVersions = *all.ServiceVersions
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
