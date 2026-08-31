// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeServiceRef Defines a ServiceRef for a cluster, enabling access to both external services and
// Services provided by other Clusters. The defined serviceRef must be provided when creating cluster.
type ModeServiceRef struct {
	// The name will be referenced in clusterCreate request. The name should also be defined in `.components`
	// so that frontend can use it to get proper localized title.
	//
	Name string `json:"name"`
	// The localized title of the serviceRef.
	Title map[string]string `json:"title,omitempty"`
	// whether this serviceRef is optional. If set to true, the cluster can be created without providing this serviceRef.
	Optional *bool `json:"optional,omitempty"`
	// The default engine to be used in serviceRef. This field is used as the fallback engine filter and default create entry.
	EngineName string `json:"engineName"`
	// The mode to be used in serviceRef. This field is used to filter clusters. If not set, it means all modes are supported.
	Modes []string `json:"modes,omitempty"`
	// specify the style that will be used in servicedescriptor.
	// "hostport" will ask user to provide both host and port.
	// "endpoint" will ask user to provide an endpoint.
	// When using a serviceSelector, this option will not be effective.
	//
	AddressStyle ServiceDescriptorAddressStyle `json:"addressStyle"`
	// whether to disable manual input of the service reference. If set to true, users can only select from the serviceRefs provided by list clusters api.
	DisableManualInput *bool `json:"disableManualInput,omitempty"`
	// whether to disable auto creating the ServiceDescriptor object for this serviceRef.
	// If set to false, the serviceDescriptor content (host/port/username/password) will be
	// set directly into the helm values under the helmValuePath.serviceDescriptor path.
	// If omitted, the generated ServiceDescriptor name will be set instead.
	//
	DisableAutoCreateServiceDescriptor *bool `json:"disableAutoCreateServiceDescriptor,omitempty"`
	// whether to disable selecting a related cluster created in the platform for this
	// serviceRef. If set to true, the frontend should only allow manual input instead
	// of selecting from the clusters provided by the list clusters api.
	//
	DisableRelatedCluster *bool `json:"disableRelatedCluster,omitempty"`
	// Extra form field definitions for manual input of the serviceRef. The keys are the
	// field names and the frontend renders them as input items. The user-entered values
	// are passed in ClusterCreate serviceRefs[].extraForManualInput and mapped to helm
	// values via helmValuePath.extraForManualInput.
	//
	ExtraForManualInput map[string]interface{} `json:"extraForManualInput,omitempty"`
	// The path to be used in values. Separated with commas. ClusterCreate API will use these path to override values in the cluster chart.
	HelmValuePath ModeServiceRefHelmValuePath `json:"helmValuePath"`
	// ServiceSelectors will map cluster's mode to a serviceSelector. The serviceSelector
	// will be used to provide the corresponding helm values.
	// If no serviceSelector is matched, the corresponding helm value will not be set.
	//
	ServiceSelectors []ServiceSelector `json:"serviceSelectors,omitempty"`
	// Service version compatibility rules for this serviceRef. The create API uses
	// these rules to reject incompatible referenced clusters, and the frontend can
	// use them to filter or explain selectable referenced clusters.
	//
	ServiceVersionCompatibility []ModeServiceRefVersionCompatibility `json:"serviceVersionCompatibility,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeServiceRef instantiates a new ModeServiceRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeServiceRef(name string, engineName string, addressStyle ServiceDescriptorAddressStyle, helmValuePath ModeServiceRefHelmValuePath) *ModeServiceRef {
	this := ModeServiceRef{}
	this.Name = name
	this.EngineName = engineName
	this.AddressStyle = addressStyle
	this.HelmValuePath = helmValuePath
	return &this
}

// NewModeServiceRefWithDefaults instantiates a new ModeServiceRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeServiceRefWithDefaults() *ModeServiceRef {
	this := ModeServiceRef{}
	return &this
}

// GetName returns the Name field value.
func (o *ModeServiceRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ModeServiceRef) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ModeServiceRef) GetTitle() map[string]string {
	if o == nil || o.Title == nil {
		var ret map[string]string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetTitleOk() (*map[string]string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return &o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModeServiceRef) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given map[string]string and assigns it to the Title field.
func (o *ModeServiceRef) SetTitle(v map[string]string) {
	o.Title = v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *ModeServiceRef) GetOptional() bool {
	if o == nil || o.Optional == nil {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetOptionalOk() (*bool, bool) {
	if o == nil || o.Optional == nil {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *ModeServiceRef) HasOptional() bool {
	return o != nil && o.Optional != nil
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *ModeServiceRef) SetOptional(v bool) {
	o.Optional = &v
}

// GetEngineName returns the EngineName field value.
func (o *ModeServiceRef) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *ModeServiceRef) SetEngineName(v string) {
	o.EngineName = v
}

// GetModes returns the Modes field value if set, zero value otherwise.
func (o *ModeServiceRef) GetModes() []string {
	if o == nil || o.Modes == nil {
		var ret []string
		return ret
	}
	return o.Modes
}

// GetModesOk returns a tuple with the Modes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetModesOk() (*[]string, bool) {
	if o == nil || o.Modes == nil {
		return nil, false
	}
	return &o.Modes, true
}

// HasModes returns a boolean if a field has been set.
func (o *ModeServiceRef) HasModes() bool {
	return o != nil && o.Modes != nil
}

// SetModes gets a reference to the given []string and assigns it to the Modes field.
func (o *ModeServiceRef) SetModes(v []string) {
	o.Modes = v
}

// GetAddressStyle returns the AddressStyle field value.
func (o *ModeServiceRef) GetAddressStyle() ServiceDescriptorAddressStyle {
	if o == nil {
		var ret ServiceDescriptorAddressStyle
		return ret
	}
	return o.AddressStyle
}

// GetAddressStyleOk returns a tuple with the AddressStyle field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetAddressStyleOk() (*ServiceDescriptorAddressStyle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressStyle, true
}

// SetAddressStyle sets field value.
func (o *ModeServiceRef) SetAddressStyle(v ServiceDescriptorAddressStyle) {
	o.AddressStyle = v
}

// GetDisableManualInput returns the DisableManualInput field value if set, zero value otherwise.
func (o *ModeServiceRef) GetDisableManualInput() bool {
	if o == nil || o.DisableManualInput == nil {
		var ret bool
		return ret
	}
	return *o.DisableManualInput
}

// GetDisableManualInputOk returns a tuple with the DisableManualInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetDisableManualInputOk() (*bool, bool) {
	if o == nil || o.DisableManualInput == nil {
		return nil, false
	}
	return o.DisableManualInput, true
}

// HasDisableManualInput returns a boolean if a field has been set.
func (o *ModeServiceRef) HasDisableManualInput() bool {
	return o != nil && o.DisableManualInput != nil
}

// SetDisableManualInput gets a reference to the given bool and assigns it to the DisableManualInput field.
func (o *ModeServiceRef) SetDisableManualInput(v bool) {
	o.DisableManualInput = &v
}

// GetDisableAutoCreateServiceDescriptor returns the DisableAutoCreateServiceDescriptor field value if set, zero value otherwise.
func (o *ModeServiceRef) GetDisableAutoCreateServiceDescriptor() bool {
	if o == nil || o.DisableAutoCreateServiceDescriptor == nil {
		var ret bool
		return ret
	}
	return *o.DisableAutoCreateServiceDescriptor
}

// GetDisableAutoCreateServiceDescriptorOk returns a tuple with the DisableAutoCreateServiceDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetDisableAutoCreateServiceDescriptorOk() (*bool, bool) {
	if o == nil || o.DisableAutoCreateServiceDescriptor == nil {
		return nil, false
	}
	return o.DisableAutoCreateServiceDescriptor, true
}

// HasDisableAutoCreateServiceDescriptor returns a boolean if a field has been set.
func (o *ModeServiceRef) HasDisableAutoCreateServiceDescriptor() bool {
	return o != nil && o.DisableAutoCreateServiceDescriptor != nil
}

// SetDisableAutoCreateServiceDescriptor gets a reference to the given bool and assigns it to the DisableAutoCreateServiceDescriptor field.
func (o *ModeServiceRef) SetDisableAutoCreateServiceDescriptor(v bool) {
	o.DisableAutoCreateServiceDescriptor = &v
}

// GetDisableRelatedCluster returns the DisableRelatedCluster field value if set, zero value otherwise.
func (o *ModeServiceRef) GetDisableRelatedCluster() bool {
	if o == nil || o.DisableRelatedCluster == nil {
		var ret bool
		return ret
	}
	return *o.DisableRelatedCluster
}

// GetDisableRelatedClusterOk returns a tuple with the DisableRelatedCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetDisableRelatedClusterOk() (*bool, bool) {
	if o == nil || o.DisableRelatedCluster == nil {
		return nil, false
	}
	return o.DisableRelatedCluster, true
}

// HasDisableRelatedCluster returns a boolean if a field has been set.
func (o *ModeServiceRef) HasDisableRelatedCluster() bool {
	return o != nil && o.DisableRelatedCluster != nil
}

// SetDisableRelatedCluster gets a reference to the given bool and assigns it to the DisableRelatedCluster field.
func (o *ModeServiceRef) SetDisableRelatedCluster(v bool) {
	o.DisableRelatedCluster = &v
}

// GetExtraForManualInput returns the ExtraForManualInput field value if set, zero value otherwise.
func (o *ModeServiceRef) GetExtraForManualInput() map[string]interface{} {
	if o == nil || o.ExtraForManualInput == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraForManualInput
}

// GetExtraForManualInputOk returns a tuple with the ExtraForManualInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetExtraForManualInputOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraForManualInput == nil {
		return nil, false
	}
	return &o.ExtraForManualInput, true
}

// HasExtraForManualInput returns a boolean if a field has been set.
func (o *ModeServiceRef) HasExtraForManualInput() bool {
	return o != nil && o.ExtraForManualInput != nil
}

// SetExtraForManualInput gets a reference to the given map[string]interface{} and assigns it to the ExtraForManualInput field.
func (o *ModeServiceRef) SetExtraForManualInput(v map[string]interface{}) {
	o.ExtraForManualInput = v
}

// GetHelmValuePath returns the HelmValuePath field value.
func (o *ModeServiceRef) GetHelmValuePath() ModeServiceRefHelmValuePath {
	if o == nil {
		var ret ModeServiceRefHelmValuePath
		return ret
	}
	return o.HelmValuePath
}

// GetHelmValuePathOk returns a tuple with the HelmValuePath field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetHelmValuePathOk() (*ModeServiceRefHelmValuePath, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HelmValuePath, true
}

// SetHelmValuePath sets field value.
func (o *ModeServiceRef) SetHelmValuePath(v ModeServiceRefHelmValuePath) {
	o.HelmValuePath = v
}

// GetServiceSelectors returns the ServiceSelectors field value if set, zero value otherwise.
func (o *ModeServiceRef) GetServiceSelectors() []ServiceSelector {
	if o == nil || o.ServiceSelectors == nil {
		var ret []ServiceSelector
		return ret
	}
	return o.ServiceSelectors
}

// GetServiceSelectorsOk returns a tuple with the ServiceSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetServiceSelectorsOk() (*[]ServiceSelector, bool) {
	if o == nil || o.ServiceSelectors == nil {
		return nil, false
	}
	return &o.ServiceSelectors, true
}

// HasServiceSelectors returns a boolean if a field has been set.
func (o *ModeServiceRef) HasServiceSelectors() bool {
	return o != nil && o.ServiceSelectors != nil
}

// SetServiceSelectors gets a reference to the given []ServiceSelector and assigns it to the ServiceSelectors field.
func (o *ModeServiceRef) SetServiceSelectors(v []ServiceSelector) {
	o.ServiceSelectors = v
}

// GetServiceVersionCompatibility returns the ServiceVersionCompatibility field value if set, zero value otherwise.
func (o *ModeServiceRef) GetServiceVersionCompatibility() []ModeServiceRefVersionCompatibility {
	if o == nil || o.ServiceVersionCompatibility == nil {
		var ret []ModeServiceRefVersionCompatibility
		return ret
	}
	return o.ServiceVersionCompatibility
}

// GetServiceVersionCompatibilityOk returns a tuple with the ServiceVersionCompatibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRef) GetServiceVersionCompatibilityOk() (*[]ModeServiceRefVersionCompatibility, bool) {
	if o == nil || o.ServiceVersionCompatibility == nil {
		return nil, false
	}
	return &o.ServiceVersionCompatibility, true
}

// HasServiceVersionCompatibility returns a boolean if a field has been set.
func (o *ModeServiceRef) HasServiceVersionCompatibility() bool {
	return o != nil && o.ServiceVersionCompatibility != nil
}

// SetServiceVersionCompatibility gets a reference to the given []ModeServiceRefVersionCompatibility and assigns it to the ServiceVersionCompatibility field.
func (o *ModeServiceRef) SetServiceVersionCompatibility(v []ModeServiceRefVersionCompatibility) {
	o.ServiceVersionCompatibility = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeServiceRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Optional != nil {
		toSerialize["optional"] = o.Optional
	}
	toSerialize["engineName"] = o.EngineName
	if o.Modes != nil {
		toSerialize["modes"] = o.Modes
	}
	toSerialize["addressStyle"] = o.AddressStyle
	if o.DisableManualInput != nil {
		toSerialize["disableManualInput"] = o.DisableManualInput
	}
	if o.DisableAutoCreateServiceDescriptor != nil {
		toSerialize["disableAutoCreateServiceDescriptor"] = o.DisableAutoCreateServiceDescriptor
	}
	if o.DisableRelatedCluster != nil {
		toSerialize["disableRelatedCluster"] = o.DisableRelatedCluster
	}
	if o.ExtraForManualInput != nil {
		toSerialize["extraForManualInput"] = o.ExtraForManualInput
	}
	toSerialize["helmValuePath"] = o.HelmValuePath
	if o.ServiceSelectors != nil {
		toSerialize["serviceSelectors"] = o.ServiceSelectors
	}
	if o.ServiceVersionCompatibility != nil {
		toSerialize["serviceVersionCompatibility"] = o.ServiceVersionCompatibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeServiceRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                               *string                              `json:"name"`
		Title                              map[string]string                    `json:"title,omitempty"`
		Optional                           *bool                                `json:"optional,omitempty"`
		EngineName                         *string                              `json:"engineName"`
		Modes                              []string                             `json:"modes,omitempty"`
		AddressStyle                       *ServiceDescriptorAddressStyle       `json:"addressStyle"`
		DisableManualInput                 *bool                                `json:"disableManualInput,omitempty"`
		DisableAutoCreateServiceDescriptor *bool                                `json:"disableAutoCreateServiceDescriptor,omitempty"`
		DisableRelatedCluster              *bool                                `json:"disableRelatedCluster,omitempty"`
		ExtraForManualInput                map[string]interface{}               `json:"extraForManualInput,omitempty"`
		HelmValuePath                      *ModeServiceRefHelmValuePath         `json:"helmValuePath"`
		ServiceSelectors                   []ServiceSelector                    `json:"serviceSelectors,omitempty"`
		ServiceVersionCompatibility        []ModeServiceRefVersionCompatibility `json:"serviceVersionCompatibility,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.AddressStyle == nil {
		return fmt.Errorf("required field addressStyle missing")
	}
	if all.HelmValuePath == nil {
		return fmt.Errorf("required field helmValuePath missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "title", "optional", "engineName", "modes", "addressStyle", "disableManualInput", "disableAutoCreateServiceDescriptor", "disableRelatedCluster", "extraForManualInput", "helmValuePath", "serviceSelectors", "serviceVersionCompatibility"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Title = all.Title
	o.Optional = all.Optional
	o.EngineName = *all.EngineName
	o.Modes = all.Modes
	if !all.AddressStyle.IsValid() {
		hasInvalidField = true
	} else {
		o.AddressStyle = *all.AddressStyle
	}
	o.DisableManualInput = all.DisableManualInput
	o.DisableAutoCreateServiceDescriptor = all.DisableAutoCreateServiceDescriptor
	o.DisableRelatedCluster = all.DisableRelatedCluster
	o.ExtraForManualInput = all.ExtraForManualInput
	if all.HelmValuePath.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.HelmValuePath = *all.HelmValuePath
	o.ServiceSelectors = all.ServiceSelectors
	o.ServiceVersionCompatibility = all.ServiceVersionCompatibility

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
