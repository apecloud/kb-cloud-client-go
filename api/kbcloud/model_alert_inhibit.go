// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// AlertInhibit Alert object information
type AlertInhibit struct {
	Id          *int32              `json:"id,omitempty"`
	Name        *string             `json:"name,omitempty"`
	Description *string             `json:"description,omitempty"`
	OrgName     *string             `json:"orgName,omitempty"`
	Envs        []string            `json:"envs,omitempty"`
	SourceMatch map[string][]string `json:"sourceMatch,omitempty"`
	TargetMatch map[string][]string `json:"targetMatch,omitempty"`
	Equal       []string            `json:"equal,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertInhibit instantiates a new AlertInhibit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertInhibit() *AlertInhibit {
	this := AlertInhibit{}
	return &this
}

// NewAlertInhibitWithDefaults instantiates a new AlertInhibit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertInhibitWithDefaults() *AlertInhibit {
	this := AlertInhibit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertInhibit) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInhibit) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertInhibit) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlertInhibit) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertInhibit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInhibit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertInhibit) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertInhibit) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlertInhibit) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInhibit) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlertInhibit) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlertInhibit) SetDescription(v string) {
	o.Description = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AlertInhibit) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInhibit) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AlertInhibit) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AlertInhibit) SetOrgName(v string) {
	o.OrgName = &v
}

// GetEnvs returns the Envs field value if set, zero value otherwise.
func (o *AlertInhibit) GetEnvs() []string {
	if o == nil || o.Envs == nil {
		var ret []string
		return ret
	}
	return o.Envs
}

// GetEnvsOk returns a tuple with the Envs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInhibit) GetEnvsOk() (*[]string, bool) {
	if o == nil || o.Envs == nil {
		return nil, false
	}
	return &o.Envs, true
}

// HasEnvs returns a boolean if a field has been set.
func (o *AlertInhibit) HasEnvs() bool {
	return o != nil && o.Envs != nil
}

// SetEnvs gets a reference to the given []string and assigns it to the Envs field.
func (o *AlertInhibit) SetEnvs(v []string) {
	o.Envs = v
}

// GetSourceMatch returns the SourceMatch field value if set, zero value otherwise.
func (o *AlertInhibit) GetSourceMatch() map[string][]string {
	if o == nil || o.SourceMatch == nil {
		var ret map[string][]string
		return ret
	}
	return o.SourceMatch
}

// GetSourceMatchOk returns a tuple with the SourceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInhibit) GetSourceMatchOk() (*map[string][]string, bool) {
	if o == nil || o.SourceMatch == nil {
		return nil, false
	}
	return &o.SourceMatch, true
}

// HasSourceMatch returns a boolean if a field has been set.
func (o *AlertInhibit) HasSourceMatch() bool {
	return o != nil && o.SourceMatch != nil
}

// SetSourceMatch gets a reference to the given map[string][]string and assigns it to the SourceMatch field.
func (o *AlertInhibit) SetSourceMatch(v map[string][]string) {
	o.SourceMatch = v
}

// GetTargetMatch returns the TargetMatch field value if set, zero value otherwise.
func (o *AlertInhibit) GetTargetMatch() map[string][]string {
	if o == nil || o.TargetMatch == nil {
		var ret map[string][]string
		return ret
	}
	return o.TargetMatch
}

// GetTargetMatchOk returns a tuple with the TargetMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInhibit) GetTargetMatchOk() (*map[string][]string, bool) {
	if o == nil || o.TargetMatch == nil {
		return nil, false
	}
	return &o.TargetMatch, true
}

// HasTargetMatch returns a boolean if a field has been set.
func (o *AlertInhibit) HasTargetMatch() bool {
	return o != nil && o.TargetMatch != nil
}

// SetTargetMatch gets a reference to the given map[string][]string and assigns it to the TargetMatch field.
func (o *AlertInhibit) SetTargetMatch(v map[string][]string) {
	o.TargetMatch = v
}

// GetEqual returns the Equal field value if set, zero value otherwise.
func (o *AlertInhibit) GetEqual() []string {
	if o == nil || o.Equal == nil {
		var ret []string
		return ret
	}
	return o.Equal
}

// GetEqualOk returns a tuple with the Equal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertInhibit) GetEqualOk() (*[]string, bool) {
	if o == nil || o.Equal == nil {
		return nil, false
	}
	return &o.Equal, true
}

// HasEqual returns a boolean if a field has been set.
func (o *AlertInhibit) HasEqual() bool {
	return o != nil && o.Equal != nil
}

// SetEqual gets a reference to the given []string and assigns it to the Equal field.
func (o *AlertInhibit) SetEqual(v []string) {
	o.Equal = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertInhibit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.Envs != nil {
		toSerialize["envs"] = o.Envs
	}
	if o.SourceMatch != nil {
		toSerialize["sourceMatch"] = o.SourceMatch
	}
	if o.TargetMatch != nil {
		toSerialize["targetMatch"] = o.TargetMatch
	}
	if o.Equal != nil {
		toSerialize["equal"] = o.Equal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertInhibit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *int32              `json:"id,omitempty"`
		Name        *string             `json:"name,omitempty"`
		Description *string             `json:"description,omitempty"`
		OrgName     *string             `json:"orgName,omitempty"`
		Envs        []string            `json:"envs,omitempty"`
		SourceMatch map[string][]string `json:"sourceMatch,omitempty"`
		TargetMatch map[string][]string `json:"targetMatch,omitempty"`
		Equal       []string            `json:"equal,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "description", "orgName", "envs", "sourceMatch", "targetMatch", "equal"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.Description = all.Description
	o.OrgName = all.OrgName
	o.Envs = all.Envs
	o.SourceMatch = all.SourceMatch
	o.TargetMatch = all.TargetMatch
	o.Equal = all.Equal

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
