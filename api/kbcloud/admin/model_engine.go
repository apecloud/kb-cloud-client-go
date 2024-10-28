// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// NODESCRIPTION Engine
type Engine struct {
	// engine ID
	Id *string `json:"id,omitempty"`
	// the engine description
	Description *string `json:"description,omitempty"`
	// engine Name
	Name *string `json:"name,omitempty"`
	// engine version
	Version *string `json:"version,omitempty"`
	// KubeBlocks version constrain
	KbVersionConstraint *string `json:"kbVersionConstraint,omitempty"`
	// engine type
	Type *EngineType `json:"type,omitempty"`
	// whether the engine is installed
	Installed *bool `json:"installed,omitempty"`
	// engine addon provider
	Provider *string `json:"provider,omitempty"`
	// engine addon status in K8s
	Status *EngineStatus `json:"status,omitempty"`
	// engine available versions
	AvailableVersion []string `json:"availableVersion,omitempty"`
	// engine upgrade history
	UpgradeHistory *string `json:"upgradeHistory,omitempty"`
	// engine error messages when the engine status is failed
	ErrMsg *string `json:"errMsg,omitempty"`
	// clusterversion in the engines
	ClusterVersions []string `json:"clusterVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngine instantiates a new Engine object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngine() *Engine {
	this := Engine{}
	return &this
}

// NewEngineWithDefaults instantiates a new Engine object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineWithDefaults() *Engine {
	this := Engine{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Engine) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Engine) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Engine) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Engine) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Engine) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Engine) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Engine) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Engine) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Engine) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Engine) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Engine) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Engine) SetVersion(v string) {
	o.Version = &v
}

// GetKbVersionConstraint returns the KbVersionConstraint field value if set, zero value otherwise.
func (o *Engine) GetKbVersionConstraint() string {
	if o == nil || o.KbVersionConstraint == nil {
		var ret string
		return ret
	}
	return *o.KbVersionConstraint
}

// GetKbVersionConstraintOk returns a tuple with the KbVersionConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetKbVersionConstraintOk() (*string, bool) {
	if o == nil || o.KbVersionConstraint == nil {
		return nil, false
	}
	return o.KbVersionConstraint, true
}

// HasKbVersionConstraint returns a boolean if a field has been set.
func (o *Engine) HasKbVersionConstraint() bool {
	return o != nil && o.KbVersionConstraint != nil
}

// SetKbVersionConstraint gets a reference to the given string and assigns it to the KbVersionConstraint field.
func (o *Engine) SetKbVersionConstraint(v string) {
	o.KbVersionConstraint = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Engine) GetType() EngineType {
	if o == nil || o.Type == nil {
		var ret EngineType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetTypeOk() (*EngineType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Engine) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given EngineType and assigns it to the Type field.
func (o *Engine) SetType(v EngineType) {
	o.Type = &v
}

// GetInstalled returns the Installed field value if set, zero value otherwise.
func (o *Engine) GetInstalled() bool {
	if o == nil || o.Installed == nil {
		var ret bool
		return ret
	}
	return *o.Installed
}

// GetInstalledOk returns a tuple with the Installed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetInstalledOk() (*bool, bool) {
	if o == nil || o.Installed == nil {
		return nil, false
	}
	return o.Installed, true
}

// HasInstalled returns a boolean if a field has been set.
func (o *Engine) HasInstalled() bool {
	return o != nil && o.Installed != nil
}

// SetInstalled gets a reference to the given bool and assigns it to the Installed field.
func (o *Engine) SetInstalled(v bool) {
	o.Installed = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Engine) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Engine) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *Engine) SetProvider(v string) {
	o.Provider = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Engine) GetStatus() EngineStatus {
	if o == nil || o.Status == nil {
		var ret EngineStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetStatusOk() (*EngineStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Engine) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given EngineStatus and assigns it to the Status field.
func (o *Engine) SetStatus(v EngineStatus) {
	o.Status = &v
}

// GetAvailableVersion returns the AvailableVersion field value if set, zero value otherwise.
func (o *Engine) GetAvailableVersion() []string {
	if o == nil || o.AvailableVersion == nil {
		var ret []string
		return ret
	}
	return o.AvailableVersion
}

// GetAvailableVersionOk returns a tuple with the AvailableVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetAvailableVersionOk() (*[]string, bool) {
	if o == nil || o.AvailableVersion == nil {
		return nil, false
	}
	return &o.AvailableVersion, true
}

// HasAvailableVersion returns a boolean if a field has been set.
func (o *Engine) HasAvailableVersion() bool {
	return o != nil && o.AvailableVersion != nil
}

// SetAvailableVersion gets a reference to the given []string and assigns it to the AvailableVersion field.
func (o *Engine) SetAvailableVersion(v []string) {
	o.AvailableVersion = v
}

// GetUpgradeHistory returns the UpgradeHistory field value if set, zero value otherwise.
func (o *Engine) GetUpgradeHistory() string {
	if o == nil || o.UpgradeHistory == nil {
		var ret string
		return ret
	}
	return *o.UpgradeHistory
}

// GetUpgradeHistoryOk returns a tuple with the UpgradeHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetUpgradeHistoryOk() (*string, bool) {
	if o == nil || o.UpgradeHistory == nil {
		return nil, false
	}
	return o.UpgradeHistory, true
}

// HasUpgradeHistory returns a boolean if a field has been set.
func (o *Engine) HasUpgradeHistory() bool {
	return o != nil && o.UpgradeHistory != nil
}

// SetUpgradeHistory gets a reference to the given string and assigns it to the UpgradeHistory field.
func (o *Engine) SetUpgradeHistory(v string) {
	o.UpgradeHistory = &v
}

// GetErrMsg returns the ErrMsg field value if set, zero value otherwise.
func (o *Engine) GetErrMsg() string {
	if o == nil || o.ErrMsg == nil {
		var ret string
		return ret
	}
	return *o.ErrMsg
}

// GetErrMsgOk returns a tuple with the ErrMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetErrMsgOk() (*string, bool) {
	if o == nil || o.ErrMsg == nil {
		return nil, false
	}
	return o.ErrMsg, true
}

// HasErrMsg returns a boolean if a field has been set.
func (o *Engine) HasErrMsg() bool {
	return o != nil && o.ErrMsg != nil
}

// SetErrMsg gets a reference to the given string and assigns it to the ErrMsg field.
func (o *Engine) SetErrMsg(v string) {
	o.ErrMsg = &v
}

// GetClusterVersions returns the ClusterVersions field value if set, zero value otherwise.
func (o *Engine) GetClusterVersions() []string {
	if o == nil || o.ClusterVersions == nil {
		var ret []string
		return ret
	}
	return o.ClusterVersions
}

// GetClusterVersionsOk returns a tuple with the ClusterVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetClusterVersionsOk() (*[]string, bool) {
	if o == nil || o.ClusterVersions == nil {
		return nil, false
	}
	return &o.ClusterVersions, true
}

// HasClusterVersions returns a boolean if a field has been set.
func (o *Engine) HasClusterVersions() bool {
	return o != nil && o.ClusterVersions != nil
}

// SetClusterVersions gets a reference to the given []string and assigns it to the ClusterVersions field.
func (o *Engine) SetClusterVersions(v []string) {
	o.ClusterVersions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Engine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.KbVersionConstraint != nil {
		toSerialize["kbVersionConstraint"] = o.KbVersionConstraint
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Installed != nil {
		toSerialize["installed"] = o.Installed
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.AvailableVersion != nil {
		toSerialize["availableVersion"] = o.AvailableVersion
	}
	if o.UpgradeHistory != nil {
		toSerialize["upgradeHistory"] = o.UpgradeHistory
	}
	if o.ErrMsg != nil {
		toSerialize["errMsg"] = o.ErrMsg
	}
	if o.ClusterVersions != nil {
		toSerialize["clusterVersions"] = o.ClusterVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Engine) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                  *string       `json:"id,omitempty"`
		Description         *string       `json:"description,omitempty"`
		Name                *string       `json:"name,omitempty"`
		Version             *string       `json:"version,omitempty"`
		KbVersionConstraint *string       `json:"kbVersionConstraint,omitempty"`
		Type                *EngineType   `json:"type,omitempty"`
		Installed           *bool         `json:"installed,omitempty"`
		Provider            *string       `json:"provider,omitempty"`
		Status              *EngineStatus `json:"status,omitempty"`
		AvailableVersion    []string      `json:"availableVersion,omitempty"`
		UpgradeHistory      *string       `json:"upgradeHistory,omitempty"`
		ErrMsg              *string       `json:"errMsg,omitempty"`
		ClusterVersions     []string      `json:"clusterVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "description", "name", "version", "kbVersionConstraint", "type", "installed", "provider", "status", "availableVersion", "upgradeHistory", "errMsg", "clusterVersions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Description = all.Description
	o.Name = all.Name
	o.Version = all.Version
	o.KbVersionConstraint = all.KbVersionConstraint
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Installed = all.Installed
	o.Provider = all.Provider
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.AvailableVersion = all.AvailableVersion
	o.UpgradeHistory = all.UpgradeHistory
	o.ErrMsg = all.ErrMsg
	o.ClusterVersions = all.ClusterVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
