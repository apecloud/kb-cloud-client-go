// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type EngineOption struct {
	EngineName string `json:"engineName"`
	Title string `json:"title"`
	Description LocalizedDescription `json:"description"`
	Versions []string `json:"versions"`
	Components []ComponentOption `json:"components"`
	Modes []ModeOption `json:"modes"`
	Account AccountOption `json:"account"`
	Database *DatabaseOption `json:"database,omitempty"`
	Dms DmsOption `json:"dms"`
	Backup BackupOption `json:"backup"`
	Bench BenchOption `json:"bench"`
	Endpoints []EndpointOption `json:"endpoints"`
	Promote []ComponentOpsOption `json:"promote"`
	Stop []ComponentOpsOption `json:"stop"`
	Start []ComponentOpsOption `json:"start"`
	Restart []ComponentOpsOption `json:"restart"`
	Hscale []ComponentOpsOption `json:"hscale"`
	Vscale []ComponentOpsOption `json:"vscale"`
	License *EngineOptionLicense `json:"license,omitempty"`
	StorageExpansion []ComponentOpsOption `json:"storageExpansion"`
	RebuildInstance []ComponentOpsOption `json:"rebuildInstance,omitempty"`
	Metrics MetricsOption `json:"metrics"`
	Dashboards []DashboardOption `json:"dashboards"`
	Logs []LogOption `json:"logs"`
	Parameters []ParameterOption `json:"parameters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewEngineOption instantiates a new EngineOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOption(engineName string, title string, description LocalizedDescription, versions []string, components []ComponentOption, modes []ModeOption, account AccountOption, dms DmsOption, backup BackupOption, bench BenchOption, endpoints []EndpointOption, promote []ComponentOpsOption, stop []ComponentOpsOption, start []ComponentOpsOption, restart []ComponentOpsOption, hscale []ComponentOpsOption, vscale []ComponentOpsOption, storageExpansion []ComponentOpsOption, metrics MetricsOption, dashboards []DashboardOption, logs []LogOption, parameters []ParameterOption) *EngineOption {
	this := EngineOption{}
	this.EngineName = engineName
	this.Title = title
	this.Description = description
	this.Versions = versions
	this.Components = components
	this.Modes = modes
	this.Account = account
	this.Dms = dms
	this.Backup = backup
	this.Bench = bench
	this.Endpoints = endpoints
	this.Promote = promote
	this.Stop = stop
	this.Start = start
	this.Restart = restart
	this.Hscale = hscale
	this.Vscale = vscale
	this.StorageExpansion = storageExpansion
	this.Metrics = metrics
	this.Dashboards = dashboards
	this.Logs = logs
	this.Parameters = parameters
	return &this
}

// NewEngineOptionWithDefaults instantiates a new EngineOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineOptionWithDefaults() *EngineOption {
	this := EngineOption{}
	return &this
}
// GetEngineName returns the EngineName field value.
func (o *EngineOption) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *EngineOption) SetEngineName(v string) {
	o.EngineName = v
}


// GetTitle returns the Title field value.
func (o *EngineOption) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *EngineOption) SetTitle(v string) {
	o.Title = v
}


// GetDescription returns the Description field value.
func (o *EngineOption) GetDescription() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *EngineOption) SetDescription(v LocalizedDescription) {
	o.Description = v
}


// GetVersions returns the Versions field value.
func (o *EngineOption) GetVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Versions, true
}

// SetVersions sets field value.
func (o *EngineOption) SetVersions(v []string) {
	o.Versions = v
}


// GetComponents returns the Components field value.
func (o *EngineOption) GetComponents() []ComponentOption {
	if o == nil {
		var ret []ComponentOption
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetComponentsOk() (*[]ComponentOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *EngineOption) SetComponents(v []ComponentOption) {
	o.Components = v
}


// GetModes returns the Modes field value.
func (o *EngineOption) GetModes() []ModeOption {
	if o == nil {
		var ret []ModeOption
		return ret
	}
	return o.Modes
}

// GetModesOk returns a tuple with the Modes field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetModesOk() (*[]ModeOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modes, true
}

// SetModes sets field value.
func (o *EngineOption) SetModes(v []ModeOption) {
	o.Modes = v
}


// GetAccount returns the Account field value.
func (o *EngineOption) GetAccount() AccountOption {
	if o == nil {
		var ret AccountOption
		return ret
	}
	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetAccountOk() (*AccountOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value.
func (o *EngineOption) SetAccount(v AccountOption) {
	o.Account = v
}


// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *EngineOption) GetDatabase() DatabaseOption {
	if o == nil || o.Database == nil {
		var ret DatabaseOption
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOption) GetDatabaseOk() (*DatabaseOption, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *EngineOption) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given DatabaseOption and assigns it to the Database field.
func (o *EngineOption) SetDatabase(v DatabaseOption) {
	o.Database = &v
}


// GetDms returns the Dms field value.
func (o *EngineOption) GetDms() DmsOption {
	if o == nil {
		var ret DmsOption
		return ret
	}
	return o.Dms
}

// GetDmsOk returns a tuple with the Dms field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetDmsOk() (*DmsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dms, true
}

// SetDms sets field value.
func (o *EngineOption) SetDms(v DmsOption) {
	o.Dms = v
}


// GetBackup returns the Backup field value.
func (o *EngineOption) GetBackup() BackupOption {
	if o == nil {
		var ret BackupOption
		return ret
	}
	return o.Backup
}

// GetBackupOk returns a tuple with the Backup field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetBackupOk() (*BackupOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Backup, true
}

// SetBackup sets field value.
func (o *EngineOption) SetBackup(v BackupOption) {
	o.Backup = v
}


// GetBench returns the Bench field value.
func (o *EngineOption) GetBench() BenchOption {
	if o == nil {
		var ret BenchOption
		return ret
	}
	return o.Bench
}

// GetBenchOk returns a tuple with the Bench field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetBenchOk() (*BenchOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bench, true
}

// SetBench sets field value.
func (o *EngineOption) SetBench(v BenchOption) {
	o.Bench = v
}


// GetEndpoints returns the Endpoints field value.
func (o *EngineOption) GetEndpoints() []EndpointOption {
	if o == nil {
		var ret []EndpointOption
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetEndpointsOk() (*[]EndpointOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoints, true
}

// SetEndpoints sets field value.
func (o *EngineOption) SetEndpoints(v []EndpointOption) {
	o.Endpoints = v
}


// GetPromote returns the Promote field value.
func (o *EngineOption) GetPromote() []ComponentOpsOption {
	if o == nil {
		var ret []ComponentOpsOption
		return ret
	}
	return o.Promote
}

// GetPromoteOk returns a tuple with the Promote field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetPromoteOk() (*[]ComponentOpsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promote, true
}

// SetPromote sets field value.
func (o *EngineOption) SetPromote(v []ComponentOpsOption) {
	o.Promote = v
}


// GetStop returns the Stop field value.
func (o *EngineOption) GetStop() []ComponentOpsOption {
	if o == nil {
		var ret []ComponentOpsOption
		return ret
	}
	return o.Stop
}

// GetStopOk returns a tuple with the Stop field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetStopOk() (*[]ComponentOpsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stop, true
}

// SetStop sets field value.
func (o *EngineOption) SetStop(v []ComponentOpsOption) {
	o.Stop = v
}


// GetStart returns the Start field value.
func (o *EngineOption) GetStart() []ComponentOpsOption {
	if o == nil {
		var ret []ComponentOpsOption
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetStartOk() (*[]ComponentOpsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *EngineOption) SetStart(v []ComponentOpsOption) {
	o.Start = v
}


// GetRestart returns the Restart field value.
func (o *EngineOption) GetRestart() []ComponentOpsOption {
	if o == nil {
		var ret []ComponentOpsOption
		return ret
	}
	return o.Restart
}

// GetRestartOk returns a tuple with the Restart field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetRestartOk() (*[]ComponentOpsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Restart, true
}

// SetRestart sets field value.
func (o *EngineOption) SetRestart(v []ComponentOpsOption) {
	o.Restart = v
}


// GetHscale returns the Hscale field value.
func (o *EngineOption) GetHscale() []ComponentOpsOption {
	if o == nil {
		var ret []ComponentOpsOption
		return ret
	}
	return o.Hscale
}

// GetHscaleOk returns a tuple with the Hscale field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetHscaleOk() (*[]ComponentOpsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hscale, true
}

// SetHscale sets field value.
func (o *EngineOption) SetHscale(v []ComponentOpsOption) {
	o.Hscale = v
}


// GetVscale returns the Vscale field value.
func (o *EngineOption) GetVscale() []ComponentOpsOption {
	if o == nil {
		var ret []ComponentOpsOption
		return ret
	}
	return o.Vscale
}

// GetVscaleOk returns a tuple with the Vscale field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetVscaleOk() (*[]ComponentOpsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vscale, true
}

// SetVscale sets field value.
func (o *EngineOption) SetVscale(v []ComponentOpsOption) {
	o.Vscale = v
}


// GetLicense returns the License field value if set, zero value otherwise.
func (o *EngineOption) GetLicense() EngineOptionLicense {
	if o == nil || o.License == nil {
		var ret EngineOptionLicense
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOption) GetLicenseOk() (*EngineOptionLicense, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *EngineOption) HasLicense() bool {
	return o != nil && o.License != nil
}

// SetLicense gets a reference to the given EngineOptionLicense and assigns it to the License field.
func (o *EngineOption) SetLicense(v EngineOptionLicense) {
	o.License = &v
}


// GetStorageExpansion returns the StorageExpansion field value.
func (o *EngineOption) GetStorageExpansion() []ComponentOpsOption {
	if o == nil {
		var ret []ComponentOpsOption
		return ret
	}
	return o.StorageExpansion
}

// GetStorageExpansionOk returns a tuple with the StorageExpansion field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetStorageExpansionOk() (*[]ComponentOpsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageExpansion, true
}

// SetStorageExpansion sets field value.
func (o *EngineOption) SetStorageExpansion(v []ComponentOpsOption) {
	o.StorageExpansion = v
}


// GetRebuildInstance returns the RebuildInstance field value if set, zero value otherwise.
func (o *EngineOption) GetRebuildInstance() []ComponentOpsOption {
	if o == nil || o.RebuildInstance == nil {
		var ret []ComponentOpsOption
		return ret
	}
	return o.RebuildInstance
}

// GetRebuildInstanceOk returns a tuple with the RebuildInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOption) GetRebuildInstanceOk() (*[]ComponentOpsOption, bool) {
	if o == nil || o.RebuildInstance == nil {
		return nil, false
	}
	return &o.RebuildInstance, true
}

// HasRebuildInstance returns a boolean if a field has been set.
func (o *EngineOption) HasRebuildInstance() bool {
	return o != nil && o.RebuildInstance != nil
}

// SetRebuildInstance gets a reference to the given []ComponentOpsOption and assigns it to the RebuildInstance field.
func (o *EngineOption) SetRebuildInstance(v []ComponentOpsOption) {
	o.RebuildInstance = v
}


// GetMetrics returns the Metrics field value.
func (o *EngineOption) GetMetrics() MetricsOption {
	if o == nil {
		var ret MetricsOption
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetMetricsOk() (*MetricsOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value.
func (o *EngineOption) SetMetrics(v MetricsOption) {
	o.Metrics = v
}


// GetDashboards returns the Dashboards field value.
func (o *EngineOption) GetDashboards() []DashboardOption {
	if o == nil {
		var ret []DashboardOption
		return ret
	}
	return o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetDashboardsOk() (*[]DashboardOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dashboards, true
}

// SetDashboards sets field value.
func (o *EngineOption) SetDashboards(v []DashboardOption) {
	o.Dashboards = v
}


// GetLogs returns the Logs field value.
func (o *EngineOption) GetLogs() []LogOption {
	if o == nil {
		var ret []LogOption
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetLogsOk() (*[]LogOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logs, true
}

// SetLogs sets field value.
func (o *EngineOption) SetLogs(v []LogOption) {
	o.Logs = v
}


// GetParameters returns the Parameters field value.
func (o *EngineOption) GetParameters() []ParameterOption {
	if o == nil {
		var ret []ParameterOption
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *EngineOption) GetParametersOk() (*[]ParameterOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value.
func (o *EngineOption) SetParameters(v []ParameterOption) {
	o.Parameters = v
}



// MarshalJSON serializes the struct using spec logic.
func (o EngineOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["versions"] = o.Versions
	toSerialize["components"] = o.Components
	toSerialize["modes"] = o.Modes
	toSerialize["account"] = o.Account
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	toSerialize["dms"] = o.Dms
	toSerialize["backup"] = o.Backup
	toSerialize["bench"] = o.Bench
	toSerialize["endpoints"] = o.Endpoints
	toSerialize["promote"] = o.Promote
	toSerialize["stop"] = o.Stop
	toSerialize["start"] = o.Start
	toSerialize["restart"] = o.Restart
	toSerialize["hscale"] = o.Hscale
	toSerialize["vscale"] = o.Vscale
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	toSerialize["storageExpansion"] = o.StorageExpansion
	if o.RebuildInstance != nil {
		toSerialize["rebuildInstance"] = o.RebuildInstance
	}
	toSerialize["metrics"] = o.Metrics
	toSerialize["dashboards"] = o.Dashboards
	toSerialize["logs"] = o.Logs
	toSerialize["parameters"] = o.Parameters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName *string `json:"engineName"`
		Title *string `json:"title"`
		Description *LocalizedDescription `json:"description"`
		Versions *[]string `json:"versions"`
		Components *[]ComponentOption `json:"components"`
		Modes *[]ModeOption `json:"modes"`
		Account *AccountOption `json:"account"`
		Database *DatabaseOption `json:"database,omitempty"`
		Dms *DmsOption `json:"dms"`
		Backup *BackupOption `json:"backup"`
		Bench *BenchOption `json:"bench"`
		Endpoints *[]EndpointOption `json:"endpoints"`
		Promote *[]ComponentOpsOption `json:"promote"`
		Stop *[]ComponentOpsOption `json:"stop"`
		Start *[]ComponentOpsOption `json:"start"`
		Restart *[]ComponentOpsOption `json:"restart"`
		Hscale *[]ComponentOpsOption `json:"hscale"`
		Vscale *[]ComponentOpsOption `json:"vscale"`
		License *EngineOptionLicense `json:"license,omitempty"`
		StorageExpansion *[]ComponentOpsOption `json:"storageExpansion"`
		RebuildInstance []ComponentOpsOption `json:"rebuildInstance,omitempty"`
		Metrics *MetricsOption `json:"metrics"`
		Dashboards *[]DashboardOption `json:"dashboards"`
		Logs *[]LogOption `json:"logs"`
		Parameters *[]ParameterOption `json:"parameters"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Versions == nil {
		return fmt.Errorf("required field versions missing")
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	if all.Modes == nil {
		return fmt.Errorf("required field modes missing")
	}
	if all.Account == nil {
		return fmt.Errorf("required field account missing")
	}
	if all.Dms == nil {
		return fmt.Errorf("required field dms missing")
	}
	if all.Backup == nil {
		return fmt.Errorf("required field backup missing")
	}
	if all.Bench == nil {
		return fmt.Errorf("required field bench missing")
	}
	if all.Endpoints == nil {
		return fmt.Errorf("required field endpoints missing")
	}
	if all.Promote == nil {
		return fmt.Errorf("required field promote missing")
	}
	if all.Stop == nil {
		return fmt.Errorf("required field stop missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.Restart == nil {
		return fmt.Errorf("required field restart missing")
	}
	if all.Hscale == nil {
		return fmt.Errorf("required field hscale missing")
	}
	if all.Vscale == nil {
		return fmt.Errorf("required field vscale missing")
	}
	if all.StorageExpansion == nil {
		return fmt.Errorf("required field storageExpansion missing")
	}
	if all.Metrics == nil {
		return fmt.Errorf("required field metrics missing")
	}
	if all.Dashboards == nil {
		return fmt.Errorf("required field dashboards missing")
	}
	if all.Logs == nil {
		return fmt.Errorf("required field logs missing")
	}
	if all.Parameters == nil {
		return fmt.Errorf("required field parameters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "engineName", "title", "description", "versions", "components", "modes", "account", "database", "dms", "backup", "bench", "endpoints", "promote", "stop", "start", "restart", "hscale", "vscale", "license", "storageExpansion", "rebuildInstance", "metrics", "dashboards", "logs", "parameters",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.EngineName = *all.EngineName
	o.Title = *all.Title
	if all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = *all.Description
	o.Versions = *all.Versions
	o.Components = *all.Components
	o.Modes = *all.Modes
	if all.Account.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Account = *all.Account
	if  all.Database != nil && all.Database.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Database = all.Database
	if all.Dms.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dms = *all.Dms
	if all.Backup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Backup = *all.Backup
	if all.Bench.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Bench = *all.Bench
	o.Endpoints = *all.Endpoints
	o.Promote = *all.Promote
	o.Stop = *all.Stop
	o.Start = *all.Start
	o.Restart = *all.Restart
	o.Hscale = *all.Hscale
	o.Vscale = *all.Vscale
	if  all.License != nil && all.License.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.License = all.License
	o.StorageExpansion = *all.StorageExpansion
	o.RebuildInstance = all.RebuildInstance
	if all.Metrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metrics = *all.Metrics
	o.Dashboards = *all.Dashboards
	o.Logs = *all.Logs
	o.Parameters = *all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
