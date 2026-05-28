// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticEnvironmentDiagnosisReadonlyBoundary Read-only safety boundary for environment diagnosis v0.2.
type DiagnosticEnvironmentDiagnosisReadonlyBoundary struct {
	// Whether the collector only used read-only Kubernetes list/get APIs.
	OnlyReadOnlyKubernetesApIs bool `json:"onlyReadOnlyKubernetesAPIs"`
	// Whether any destructive action was executed by this diagnosis.
	DestructiveActionExecuted bool `json:"destructiveActionExecuted"`
	// Actions that must not appear as executable nextActions.
	ForbiddenActions []string `json:"forbiddenActions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticEnvironmentDiagnosisReadonlyBoundary instantiates a new DiagnosticEnvironmentDiagnosisReadonlyBoundary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticEnvironmentDiagnosisReadonlyBoundary(onlyReadOnlyKubernetesApIs bool, destructiveActionExecuted bool, forbiddenActions []string) *DiagnosticEnvironmentDiagnosisReadonlyBoundary {
	this := DiagnosticEnvironmentDiagnosisReadonlyBoundary{}
	this.OnlyReadOnlyKubernetesApIs = onlyReadOnlyKubernetesApIs
	this.DestructiveActionExecuted = destructiveActionExecuted
	this.ForbiddenActions = forbiddenActions
	return &this
}

// NewDiagnosticEnvironmentDiagnosisReadonlyBoundaryWithDefaults instantiates a new DiagnosticEnvironmentDiagnosisReadonlyBoundary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticEnvironmentDiagnosisReadonlyBoundaryWithDefaults() *DiagnosticEnvironmentDiagnosisReadonlyBoundary {
	this := DiagnosticEnvironmentDiagnosisReadonlyBoundary{}
	return &this
}

// GetOnlyReadOnlyKubernetesApIs returns the OnlyReadOnlyKubernetesApIs field value.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) GetOnlyReadOnlyKubernetesApIs() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.OnlyReadOnlyKubernetesApIs
}

// GetOnlyReadOnlyKubernetesApIsOk returns a tuple with the OnlyReadOnlyKubernetesApIs field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) GetOnlyReadOnlyKubernetesApIsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnlyReadOnlyKubernetesApIs, true
}

// SetOnlyReadOnlyKubernetesApIs sets field value.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) SetOnlyReadOnlyKubernetesApIs(v bool) {
	o.OnlyReadOnlyKubernetesApIs = v
}

// GetDestructiveActionExecuted returns the DestructiveActionExecuted field value.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) GetDestructiveActionExecuted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DestructiveActionExecuted
}

// GetDestructiveActionExecutedOk returns a tuple with the DestructiveActionExecuted field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) GetDestructiveActionExecutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestructiveActionExecuted, true
}

// SetDestructiveActionExecuted sets field value.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) SetDestructiveActionExecuted(v bool) {
	o.DestructiveActionExecuted = v
}

// GetForbiddenActions returns the ForbiddenActions field value.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) GetForbiddenActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ForbiddenActions
}

// GetForbiddenActionsOk returns a tuple with the ForbiddenActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) GetForbiddenActionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForbiddenActions, true
}

// SetForbiddenActions sets field value.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) SetForbiddenActions(v []string) {
	o.ForbiddenActions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticEnvironmentDiagnosisReadonlyBoundary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["onlyReadOnlyKubernetesAPIs"] = o.OnlyReadOnlyKubernetesApIs
	toSerialize["destructiveActionExecuted"] = o.DestructiveActionExecuted
	toSerialize["forbiddenActions"] = o.ForbiddenActions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticEnvironmentDiagnosisReadonlyBoundary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OnlyReadOnlyKubernetesApIs *bool     `json:"onlyReadOnlyKubernetesAPIs"`
		DestructiveActionExecuted  *bool     `json:"destructiveActionExecuted"`
		ForbiddenActions           *[]string `json:"forbiddenActions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OnlyReadOnlyKubernetesApIs == nil {
		return fmt.Errorf("required field onlyReadOnlyKubernetesAPIs missing")
	}
	if all.DestructiveActionExecuted == nil {
		return fmt.Errorf("required field destructiveActionExecuted missing")
	}
	if all.ForbiddenActions == nil {
		return fmt.Errorf("required field forbiddenActions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"onlyReadOnlyKubernetesAPIs", "destructiveActionExecuted", "forbiddenActions"})
	} else {
		return err
	}
	o.OnlyReadOnlyKubernetesApIs = *all.OnlyReadOnlyKubernetesApIs
	o.DestructiveActionExecuted = *all.DestructiveActionExecuted
	o.ForbiddenActions = *all.ForbiddenActions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
