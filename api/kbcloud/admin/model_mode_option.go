// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModeOption struct {
	Name             string                      `json:"name"`
	Title            LocalizedDescription        `json:"title"`
	Description      LocalizedDescription        `json:"description"`
	SchedulingPolicy *ModeOptionSchedulingPolicy `json:"schedulingPolicy,omitempty"`
	// specify the compatible kubeblocks version. If empty, it means all versions are supported.
	//
	CompatibleKubeblocksVersion *ModeCompatibleKubeblocksVersion `json:"compatibleKubeblocksVersion,omitempty"`
	Components                  []ModeComponent                  `json:"components"`
	Proxy                       *ModeOptionProxy                 `json:"proxy,omitempty"`
	Versions                    []string                         `json:"versions,omitempty"`
	Extra                       map[string]interface{}           `json:"extra,omitempty"`
	ServiceRefs                 []ModeServiceRef                 `json:"serviceRefs,omitempty"`
	// object storage related configs
	//
	ObjectStorage  *ModeObjectStorage        `json:"objectStorage,omitempty"`
	ValuesMappings *ModeOptionValuesMappings `json:"valuesMappings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeOption instantiates a new ModeOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeOption(name string, title LocalizedDescription, description LocalizedDescription, components []ModeComponent) *ModeOption {
	this := ModeOption{}
	this.Name = name
	this.Title = title
	this.Description = description
	this.Components = components
	return &this
}

// NewModeOptionWithDefaults instantiates a new ModeOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeOptionWithDefaults() *ModeOption {
	this := ModeOption{}
	return &this
}

// GetName returns the Name field value.
func (o *ModeOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModeOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ModeOption) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value.
func (o *ModeOption) GetTitle() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ModeOption) GetTitleOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ModeOption) SetTitle(v LocalizedDescription) {
	o.Title = v
}

// GetDescription returns the Description field value.
func (o *ModeOption) GetDescription() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModeOption) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ModeOption) SetDescription(v LocalizedDescription) {
	o.Description = v
}

// GetSchedulingPolicy returns the SchedulingPolicy field value if set, zero value otherwise.
func (o *ModeOption) GetSchedulingPolicy() ModeOptionSchedulingPolicy {
	if o == nil || o.SchedulingPolicy == nil {
		var ret ModeOptionSchedulingPolicy
		return ret
	}
	return *o.SchedulingPolicy
}

// GetSchedulingPolicyOk returns a tuple with the SchedulingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetSchedulingPolicyOk() (*ModeOptionSchedulingPolicy, bool) {
	if o == nil || o.SchedulingPolicy == nil {
		return nil, false
	}
	return o.SchedulingPolicy, true
}

// HasSchedulingPolicy returns a boolean if a field has been set.
func (o *ModeOption) HasSchedulingPolicy() bool {
	return o != nil && o.SchedulingPolicy != nil
}

// SetSchedulingPolicy gets a reference to the given ModeOptionSchedulingPolicy and assigns it to the SchedulingPolicy field.
func (o *ModeOption) SetSchedulingPolicy(v ModeOptionSchedulingPolicy) {
	o.SchedulingPolicy = &v
}

// GetCompatibleKubeblocksVersion returns the CompatibleKubeblocksVersion field value if set, zero value otherwise.
func (o *ModeOption) GetCompatibleKubeblocksVersion() ModeCompatibleKubeblocksVersion {
	if o == nil || o.CompatibleKubeblocksVersion == nil {
		var ret ModeCompatibleKubeblocksVersion
		return ret
	}
	return *o.CompatibleKubeblocksVersion
}

// GetCompatibleKubeblocksVersionOk returns a tuple with the CompatibleKubeblocksVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetCompatibleKubeblocksVersionOk() (*ModeCompatibleKubeblocksVersion, bool) {
	if o == nil || o.CompatibleKubeblocksVersion == nil {
		return nil, false
	}
	return o.CompatibleKubeblocksVersion, true
}

// HasCompatibleKubeblocksVersion returns a boolean if a field has been set.
func (o *ModeOption) HasCompatibleKubeblocksVersion() bool {
	return o != nil && o.CompatibleKubeblocksVersion != nil
}

// SetCompatibleKubeblocksVersion gets a reference to the given ModeCompatibleKubeblocksVersion and assigns it to the CompatibleKubeblocksVersion field.
func (o *ModeOption) SetCompatibleKubeblocksVersion(v ModeCompatibleKubeblocksVersion) {
	o.CompatibleKubeblocksVersion = &v
}

// GetComponents returns the Components field value.
func (o *ModeOption) GetComponents() []ModeComponent {
	if o == nil {
		var ret []ModeComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *ModeOption) GetComponentsOk() (*[]ModeComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *ModeOption) SetComponents(v []ModeComponent) {
	o.Components = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *ModeOption) GetProxy() ModeOptionProxy {
	if o == nil || o.Proxy == nil {
		var ret ModeOptionProxy
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetProxyOk() (*ModeOptionProxy, bool) {
	if o == nil || o.Proxy == nil {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *ModeOption) HasProxy() bool {
	return o != nil && o.Proxy != nil
}

// SetProxy gets a reference to the given ModeOptionProxy and assigns it to the Proxy field.
func (o *ModeOption) SetProxy(v ModeOptionProxy) {
	o.Proxy = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ModeOption) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetVersionsOk() (*[]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ModeOption) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *ModeOption) SetVersions(v []string) {
	o.Versions = v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *ModeOption) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return &o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *ModeOption) HasExtra() bool {
	return o != nil && o.Extra != nil
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *ModeOption) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// GetServiceRefs returns the ServiceRefs field value if set, zero value otherwise.
func (o *ModeOption) GetServiceRefs() []ModeServiceRef {
	if o == nil || o.ServiceRefs == nil {
		var ret []ModeServiceRef
		return ret
	}
	return o.ServiceRefs
}

// GetServiceRefsOk returns a tuple with the ServiceRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetServiceRefsOk() (*[]ModeServiceRef, bool) {
	if o == nil || o.ServiceRefs == nil {
		return nil, false
	}
	return &o.ServiceRefs, true
}

// HasServiceRefs returns a boolean if a field has been set.
func (o *ModeOption) HasServiceRefs() bool {
	return o != nil && o.ServiceRefs != nil
}

// SetServiceRefs gets a reference to the given []ModeServiceRef and assigns it to the ServiceRefs field.
func (o *ModeOption) SetServiceRefs(v []ModeServiceRef) {
	o.ServiceRefs = v
}

// GetObjectStorage returns the ObjectStorage field value if set, zero value otherwise.
func (o *ModeOption) GetObjectStorage() ModeObjectStorage {
	if o == nil || o.ObjectStorage == nil {
		var ret ModeObjectStorage
		return ret
	}
	return *o.ObjectStorage
}

// GetObjectStorageOk returns a tuple with the ObjectStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetObjectStorageOk() (*ModeObjectStorage, bool) {
	if o == nil || o.ObjectStorage == nil {
		return nil, false
	}
	return o.ObjectStorage, true
}

// HasObjectStorage returns a boolean if a field has been set.
func (o *ModeOption) HasObjectStorage() bool {
	return o != nil && o.ObjectStorage != nil
}

// SetObjectStorage gets a reference to the given ModeObjectStorage and assigns it to the ObjectStorage field.
func (o *ModeOption) SetObjectStorage(v ModeObjectStorage) {
	o.ObjectStorage = &v
}

// GetValuesMappings returns the ValuesMappings field value if set, zero value otherwise.
func (o *ModeOption) GetValuesMappings() ModeOptionValuesMappings {
	if o == nil || o.ValuesMappings == nil {
		var ret ModeOptionValuesMappings
		return ret
	}
	return *o.ValuesMappings
}

// GetValuesMappingsOk returns a tuple with the ValuesMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetValuesMappingsOk() (*ModeOptionValuesMappings, bool) {
	if o == nil || o.ValuesMappings == nil {
		return nil, false
	}
	return o.ValuesMappings, true
}

// HasValuesMappings returns a boolean if a field has been set.
func (o *ModeOption) HasValuesMappings() bool {
	return o != nil && o.ValuesMappings != nil
}

// SetValuesMappings gets a reference to the given ModeOptionValuesMappings and assigns it to the ValuesMappings field.
func (o *ModeOption) SetValuesMappings(v ModeOptionValuesMappings) {
	o.ValuesMappings = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	if o.SchedulingPolicy != nil {
		toSerialize["schedulingPolicy"] = o.SchedulingPolicy
	}
	if o.CompatibleKubeblocksVersion != nil {
		toSerialize["compatibleKubeblocksVersion"] = o.CompatibleKubeblocksVersion
	}
	toSerialize["components"] = o.Components
	if o.Proxy != nil {
		toSerialize["proxy"] = o.Proxy
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.ServiceRefs != nil {
		toSerialize["serviceRefs"] = o.ServiceRefs
	}
	if o.ObjectStorage != nil {
		toSerialize["objectStorage"] = o.ObjectStorage
	}
	if o.ValuesMappings != nil {
		toSerialize["valuesMappings"] = o.ValuesMappings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                        *string                          `json:"name"`
		Title                       *LocalizedDescription            `json:"title"`
		Description                 *LocalizedDescription            `json:"description"`
		SchedulingPolicy            *ModeOptionSchedulingPolicy      `json:"schedulingPolicy,omitempty"`
		CompatibleKubeblocksVersion *ModeCompatibleKubeblocksVersion `json:"compatibleKubeblocksVersion,omitempty"`
		Components                  *[]ModeComponent                 `json:"components"`
		Proxy                       *ModeOptionProxy                 `json:"proxy,omitempty"`
		Versions                    []string                         `json:"versions,omitempty"`
		Extra                       map[string]interface{}           `json:"extra,omitempty"`
		ServiceRefs                 []ModeServiceRef                 `json:"serviceRefs,omitempty"`
		ObjectStorage               *ModeObjectStorage               `json:"objectStorage,omitempty"`
		ValuesMappings              *ModeOptionValuesMappings        `json:"valuesMappings,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "title", "description", "schedulingPolicy", "compatibleKubeblocksVersion", "components", "proxy", "versions", "extra", "serviceRefs", "objectStorage", "valuesMappings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = *all.Title
	if all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = *all.Description
	if all.SchedulingPolicy != nil && all.SchedulingPolicy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SchedulingPolicy = all.SchedulingPolicy
	if all.CompatibleKubeblocksVersion != nil && !all.CompatibleKubeblocksVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.CompatibleKubeblocksVersion = all.CompatibleKubeblocksVersion
	}
	o.Components = *all.Components
	if all.Proxy != nil && all.Proxy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Proxy = all.Proxy
	o.Versions = all.Versions
	o.Extra = all.Extra
	o.ServiceRefs = all.ServiceRefs
	if all.ObjectStorage != nil && all.ObjectStorage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectStorage = all.ObjectStorage
	if all.ValuesMappings != nil && all.ValuesMappings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ValuesMappings = all.ValuesMappings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
