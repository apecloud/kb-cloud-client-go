// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// WorkflowList component management workflow list
type WorkflowList struct {
	// component management workflow
	Install *Workflow `json:"install,omitempty"`
	// component management workflow
	Uninstall *Workflow `json:"uninstall,omitempty"`
	// component management workflow
	UpgradeKubeblocks *Workflow `json:"upgradeKubeblocks,omitempty"`
	// component management workflow
	UpgradeGemini *Workflow `json:"upgradeGemini,omitempty"`
	// component management workflow
	UpgradeVictoriaMetrics *Workflow `json:"upgradeVictoriaMetrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowList instantiates a new WorkflowList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowList() *WorkflowList {
	this := WorkflowList{}
	return &this
}

// NewWorkflowListWithDefaults instantiates a new WorkflowList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowListWithDefaults() *WorkflowList {
	this := WorkflowList{}
	return &this
}

// GetInstall returns the Install field value if set, zero value otherwise.
func (o *WorkflowList) GetInstall() Workflow {
	if o == nil || o.Install == nil {
		var ret Workflow
		return ret
	}
	return *o.Install
}

// GetInstallOk returns a tuple with the Install field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowList) GetInstallOk() (*Workflow, bool) {
	if o == nil || o.Install == nil {
		return nil, false
	}
	return o.Install, true
}

// HasInstall returns a boolean if a field has been set.
func (o *WorkflowList) HasInstall() bool {
	return o != nil && o.Install != nil
}

// SetInstall gets a reference to the given Workflow and assigns it to the Install field.
func (o *WorkflowList) SetInstall(v Workflow) {
	o.Install = &v
}

// GetUninstall returns the Uninstall field value if set, zero value otherwise.
func (o *WorkflowList) GetUninstall() Workflow {
	if o == nil || o.Uninstall == nil {
		var ret Workflow
		return ret
	}
	return *o.Uninstall
}

// GetUninstallOk returns a tuple with the Uninstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowList) GetUninstallOk() (*Workflow, bool) {
	if o == nil || o.Uninstall == nil {
		return nil, false
	}
	return o.Uninstall, true
}

// HasUninstall returns a boolean if a field has been set.
func (o *WorkflowList) HasUninstall() bool {
	return o != nil && o.Uninstall != nil
}

// SetUninstall gets a reference to the given Workflow and assigns it to the Uninstall field.
func (o *WorkflowList) SetUninstall(v Workflow) {
	o.Uninstall = &v
}

// GetUpgradeKubeblocks returns the UpgradeKubeblocks field value if set, zero value otherwise.
func (o *WorkflowList) GetUpgradeKubeblocks() Workflow {
	if o == nil || o.UpgradeKubeblocks == nil {
		var ret Workflow
		return ret
	}
	return *o.UpgradeKubeblocks
}

// GetUpgradeKubeblocksOk returns a tuple with the UpgradeKubeblocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowList) GetUpgradeKubeblocksOk() (*Workflow, bool) {
	if o == nil || o.UpgradeKubeblocks == nil {
		return nil, false
	}
	return o.UpgradeKubeblocks, true
}

// HasUpgradeKubeblocks returns a boolean if a field has been set.
func (o *WorkflowList) HasUpgradeKubeblocks() bool {
	return o != nil && o.UpgradeKubeblocks != nil
}

// SetUpgradeKubeblocks gets a reference to the given Workflow and assigns it to the UpgradeKubeblocks field.
func (o *WorkflowList) SetUpgradeKubeblocks(v Workflow) {
	o.UpgradeKubeblocks = &v
}

// GetUpgradeGemini returns the UpgradeGemini field value if set, zero value otherwise.
func (o *WorkflowList) GetUpgradeGemini() Workflow {
	if o == nil || o.UpgradeGemini == nil {
		var ret Workflow
		return ret
	}
	return *o.UpgradeGemini
}

// GetUpgradeGeminiOk returns a tuple with the UpgradeGemini field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowList) GetUpgradeGeminiOk() (*Workflow, bool) {
	if o == nil || o.UpgradeGemini == nil {
		return nil, false
	}
	return o.UpgradeGemini, true
}

// HasUpgradeGemini returns a boolean if a field has been set.
func (o *WorkflowList) HasUpgradeGemini() bool {
	return o != nil && o.UpgradeGemini != nil
}

// SetUpgradeGemini gets a reference to the given Workflow and assigns it to the UpgradeGemini field.
func (o *WorkflowList) SetUpgradeGemini(v Workflow) {
	o.UpgradeGemini = &v
}

// GetUpgradeVictoriaMetrics returns the UpgradeVictoriaMetrics field value if set, zero value otherwise.
func (o *WorkflowList) GetUpgradeVictoriaMetrics() Workflow {
	if o == nil || o.UpgradeVictoriaMetrics == nil {
		var ret Workflow
		return ret
	}
	return *o.UpgradeVictoriaMetrics
}

// GetUpgradeVictoriaMetricsOk returns a tuple with the UpgradeVictoriaMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowList) GetUpgradeVictoriaMetricsOk() (*Workflow, bool) {
	if o == nil || o.UpgradeVictoriaMetrics == nil {
		return nil, false
	}
	return o.UpgradeVictoriaMetrics, true
}

// HasUpgradeVictoriaMetrics returns a boolean if a field has been set.
func (o *WorkflowList) HasUpgradeVictoriaMetrics() bool {
	return o != nil && o.UpgradeVictoriaMetrics != nil
}

// SetUpgradeVictoriaMetrics gets a reference to the given Workflow and assigns it to the UpgradeVictoriaMetrics field.
func (o *WorkflowList) SetUpgradeVictoriaMetrics(v Workflow) {
	o.UpgradeVictoriaMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Install != nil {
		toSerialize["install"] = o.Install
	}
	if o.Uninstall != nil {
		toSerialize["uninstall"] = o.Uninstall
	}
	if o.UpgradeKubeblocks != nil {
		toSerialize["upgradeKubeblocks"] = o.UpgradeKubeblocks
	}
	if o.UpgradeGemini != nil {
		toSerialize["upgradeGemini"] = o.UpgradeGemini
	}
	if o.UpgradeVictoriaMetrics != nil {
		toSerialize["upgradeVictoriaMetrics"] = o.UpgradeVictoriaMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Install                *Workflow `json:"install,omitempty"`
		Uninstall              *Workflow `json:"uninstall,omitempty"`
		UpgradeKubeblocks      *Workflow `json:"upgradeKubeblocks,omitempty"`
		UpgradeGemini          *Workflow `json:"upgradeGemini,omitempty"`
		UpgradeVictoriaMetrics *Workflow `json:"upgradeVictoriaMetrics,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"install", "uninstall", "upgradeKubeblocks", "upgradeGemini", "upgradeVictoriaMetrics"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Install != nil && all.Install.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Install = all.Install
	if all.Uninstall != nil && all.Uninstall.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Uninstall = all.Uninstall
	if all.UpgradeKubeblocks != nil && all.UpgradeKubeblocks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UpgradeKubeblocks = all.UpgradeKubeblocks
	if all.UpgradeGemini != nil && all.UpgradeGemini.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UpgradeGemini = all.UpgradeGemini
	if all.UpgradeVictoriaMetrics != nil && all.UpgradeVictoriaMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UpgradeVictoriaMetrics = all.UpgradeVictoriaMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
