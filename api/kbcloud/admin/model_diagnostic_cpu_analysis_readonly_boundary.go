// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisReadonlyBoundary Explicit safety boundary for the CPU analysis collector.
type DiagnosticCPUAnalysisReadonlyBoundary struct {
	// Whether the database collector only runs readonly SELECT statements.
	OnlyReadonlyQueries bool `json:"onlyReadonlyQueries"`
	// Whether the collector executed any destructive action.
	DestructiveActionExecuted bool `json:"destructiveActionExecuted"`
	// Actions that must not appear as executable nextActions.
	ForbiddenActions []string `json:"forbiddenActions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticCPUAnalysisReadonlyBoundary instantiates a new DiagnosticCPUAnalysisReadonlyBoundary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticCPUAnalysisReadonlyBoundary(onlyReadonlyQueries bool, destructiveActionExecuted bool, forbiddenActions []string) *DiagnosticCPUAnalysisReadonlyBoundary {
	this := DiagnosticCPUAnalysisReadonlyBoundary{}
	this.OnlyReadonlyQueries = onlyReadonlyQueries
	this.DestructiveActionExecuted = destructiveActionExecuted
	this.ForbiddenActions = forbiddenActions
	return &this
}

// NewDiagnosticCPUAnalysisReadonlyBoundaryWithDefaults instantiates a new DiagnosticCPUAnalysisReadonlyBoundary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticCPUAnalysisReadonlyBoundaryWithDefaults() *DiagnosticCPUAnalysisReadonlyBoundary {
	this := DiagnosticCPUAnalysisReadonlyBoundary{}
	return &this
}

// GetOnlyReadonlyQueries returns the OnlyReadonlyQueries field value.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) GetOnlyReadonlyQueries() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.OnlyReadonlyQueries
}

// GetOnlyReadonlyQueriesOk returns a tuple with the OnlyReadonlyQueries field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) GetOnlyReadonlyQueriesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnlyReadonlyQueries, true
}

// SetOnlyReadonlyQueries sets field value.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) SetOnlyReadonlyQueries(v bool) {
	o.OnlyReadonlyQueries = v
}

// GetDestructiveActionExecuted returns the DestructiveActionExecuted field value.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) GetDestructiveActionExecuted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DestructiveActionExecuted
}

// GetDestructiveActionExecutedOk returns a tuple with the DestructiveActionExecuted field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) GetDestructiveActionExecutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestructiveActionExecuted, true
}

// SetDestructiveActionExecuted sets field value.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) SetDestructiveActionExecuted(v bool) {
	o.DestructiveActionExecuted = v
}

// GetForbiddenActions returns the ForbiddenActions field value.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) GetForbiddenActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ForbiddenActions
}

// GetForbiddenActionsOk returns a tuple with the ForbiddenActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) GetForbiddenActionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForbiddenActions, true
}

// SetForbiddenActions sets field value.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) SetForbiddenActions(v []string) {
	o.ForbiddenActions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticCPUAnalysisReadonlyBoundary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["onlyReadonlyQueries"] = o.OnlyReadonlyQueries
	toSerialize["destructiveActionExecuted"] = o.DestructiveActionExecuted
	toSerialize["forbiddenActions"] = o.ForbiddenActions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticCPUAnalysisReadonlyBoundary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OnlyReadonlyQueries       *bool     `json:"onlyReadonlyQueries"`
		DestructiveActionExecuted *bool     `json:"destructiveActionExecuted"`
		ForbiddenActions          *[]string `json:"forbiddenActions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OnlyReadonlyQueries == nil {
		return fmt.Errorf("required field onlyReadonlyQueries missing")
	}
	if all.DestructiveActionExecuted == nil {
		return fmt.Errorf("required field destructiveActionExecuted missing")
	}
	if all.ForbiddenActions == nil {
		return fmt.Errorf("required field forbiddenActions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"onlyReadonlyQueries", "destructiveActionExecuted", "forbiddenActions"})
	} else {
		return err
	}
	o.OnlyReadonlyQueries = *all.OnlyReadonlyQueries
	o.DestructiveActionExecuted = *all.DestructiveActionExecuted
	o.ForbiddenActions = *all.ForbiddenActions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
