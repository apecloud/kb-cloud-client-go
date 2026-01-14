// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PgExtension struct {
	// The name of the PostgreSQL extension.
	Name *string `json:"name,omitempty"`
	// The database where the extension is installed.
	Database *string `json:"database,omitempty"`
	// The default version of the extension.
	DefaultVersion *string `json:"default_version,omitempty"`
	// The currently installed version of the extension.
	InstalledVersion *string `json:"installed_version,omitempty"`
	// The schema where the extension is installed.
	Schema *string `json:"schema,omitempty"`
	// The category of the extension.
	Category *string `json:"category,omitempty"`
	// Whether the extension needs a restart to take effect.
	NeedRestart *bool `json:"need_restart,omitempty"`
	// The version of PostgreSQL that the extension is compatible with.
	PgVersion []string `json:"pg_version,omitempty"`
	// A description of the extension in Chinese.
	DescriptionZn *string `json:"description_zn,omitempty"`
	// A description of the extension in English.
	DescriptionEn *string `json:"description_en,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPgExtension instantiates a new PgExtension object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPgExtension() *PgExtension {
	this := PgExtension{}
	return &this
}

// NewPgExtensionWithDefaults instantiates a new PgExtension object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPgExtensionWithDefaults() *PgExtension {
	this := PgExtension{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PgExtension) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PgExtension) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PgExtension) SetName(v string) {
	o.Name = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *PgExtension) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *PgExtension) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *PgExtension) SetDatabase(v string) {
	o.Database = &v
}

// GetDefaultVersion returns the DefaultVersion field value if set, zero value otherwise.
func (o *PgExtension) GetDefaultVersion() string {
	if o == nil || o.DefaultVersion == nil {
		var ret string
		return ret
	}
	return *o.DefaultVersion
}

// GetDefaultVersionOk returns a tuple with the DefaultVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetDefaultVersionOk() (*string, bool) {
	if o == nil || o.DefaultVersion == nil {
		return nil, false
	}
	return o.DefaultVersion, true
}

// HasDefaultVersion returns a boolean if a field has been set.
func (o *PgExtension) HasDefaultVersion() bool {
	return o != nil && o.DefaultVersion != nil
}

// SetDefaultVersion gets a reference to the given string and assigns it to the DefaultVersion field.
func (o *PgExtension) SetDefaultVersion(v string) {
	o.DefaultVersion = &v
}

// GetInstalledVersion returns the InstalledVersion field value if set, zero value otherwise.
func (o *PgExtension) GetInstalledVersion() string {
	if o == nil || o.InstalledVersion == nil {
		var ret string
		return ret
	}
	return *o.InstalledVersion
}

// GetInstalledVersionOk returns a tuple with the InstalledVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetInstalledVersionOk() (*string, bool) {
	if o == nil || o.InstalledVersion == nil {
		return nil, false
	}
	return o.InstalledVersion, true
}

// HasInstalledVersion returns a boolean if a field has been set.
func (o *PgExtension) HasInstalledVersion() bool {
	return o != nil && o.InstalledVersion != nil
}

// SetInstalledVersion gets a reference to the given string and assigns it to the InstalledVersion field.
func (o *PgExtension) SetInstalledVersion(v string) {
	o.InstalledVersion = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PgExtension) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PgExtension) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PgExtension) SetSchema(v string) {
	o.Schema = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *PgExtension) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *PgExtension) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *PgExtension) SetCategory(v string) {
	o.Category = &v
}

// GetNeedRestart returns the NeedRestart field value if set, zero value otherwise.
func (o *PgExtension) GetNeedRestart() bool {
	if o == nil || o.NeedRestart == nil {
		var ret bool
		return ret
	}
	return *o.NeedRestart
}

// GetNeedRestartOk returns a tuple with the NeedRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetNeedRestartOk() (*bool, bool) {
	if o == nil || o.NeedRestart == nil {
		return nil, false
	}
	return o.NeedRestart, true
}

// HasNeedRestart returns a boolean if a field has been set.
func (o *PgExtension) HasNeedRestart() bool {
	return o != nil && o.NeedRestart != nil
}

// SetNeedRestart gets a reference to the given bool and assigns it to the NeedRestart field.
func (o *PgExtension) SetNeedRestart(v bool) {
	o.NeedRestart = &v
}

// GetPgVersion returns the PgVersion field value if set, zero value otherwise.
func (o *PgExtension) GetPgVersion() []string {
	if o == nil || o.PgVersion == nil {
		var ret []string
		return ret
	}
	return o.PgVersion
}

// GetPgVersionOk returns a tuple with the PgVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetPgVersionOk() (*[]string, bool) {
	if o == nil || o.PgVersion == nil {
		return nil, false
	}
	return &o.PgVersion, true
}

// HasPgVersion returns a boolean if a field has been set.
func (o *PgExtension) HasPgVersion() bool {
	return o != nil && o.PgVersion != nil
}

// SetPgVersion gets a reference to the given []string and assigns it to the PgVersion field.
func (o *PgExtension) SetPgVersion(v []string) {
	o.PgVersion = v
}

// GetDescriptionZn returns the DescriptionZn field value if set, zero value otherwise.
func (o *PgExtension) GetDescriptionZn() string {
	if o == nil || o.DescriptionZn == nil {
		var ret string
		return ret
	}
	return *o.DescriptionZn
}

// GetDescriptionZnOk returns a tuple with the DescriptionZn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetDescriptionZnOk() (*string, bool) {
	if o == nil || o.DescriptionZn == nil {
		return nil, false
	}
	return o.DescriptionZn, true
}

// HasDescriptionZn returns a boolean if a field has been set.
func (o *PgExtension) HasDescriptionZn() bool {
	return o != nil && o.DescriptionZn != nil
}

// SetDescriptionZn gets a reference to the given string and assigns it to the DescriptionZn field.
func (o *PgExtension) SetDescriptionZn(v string) {
	o.DescriptionZn = &v
}

// GetDescriptionEn returns the DescriptionEn field value if set, zero value otherwise.
func (o *PgExtension) GetDescriptionEn() string {
	if o == nil || o.DescriptionEn == nil {
		var ret string
		return ret
	}
	return *o.DescriptionEn
}

// GetDescriptionEnOk returns a tuple with the DescriptionEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PgExtension) GetDescriptionEnOk() (*string, bool) {
	if o == nil || o.DescriptionEn == nil {
		return nil, false
	}
	return o.DescriptionEn, true
}

// HasDescriptionEn returns a boolean if a field has been set.
func (o *PgExtension) HasDescriptionEn() bool {
	return o != nil && o.DescriptionEn != nil
}

// SetDescriptionEn gets a reference to the given string and assigns it to the DescriptionEn field.
func (o *PgExtension) SetDescriptionEn(v string) {
	o.DescriptionEn = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PgExtension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.DefaultVersion != nil {
		toSerialize["default_version"] = o.DefaultVersion
	}
	if o.InstalledVersion != nil {
		toSerialize["installed_version"] = o.InstalledVersion
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.NeedRestart != nil {
		toSerialize["need_restart"] = o.NeedRestart
	}
	if o.PgVersion != nil {
		toSerialize["pg_version"] = o.PgVersion
	}
	if o.DescriptionZn != nil {
		toSerialize["description_zn"] = o.DescriptionZn
	}
	if o.DescriptionEn != nil {
		toSerialize["description_en"] = o.DescriptionEn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PgExtension) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string  `json:"name,omitempty"`
		Database         *string  `json:"database,omitempty"`
		DefaultVersion   *string  `json:"default_version,omitempty"`
		InstalledVersion *string  `json:"installed_version,omitempty"`
		Schema           *string  `json:"schema,omitempty"`
		Category         *string  `json:"category,omitempty"`
		NeedRestart      *bool    `json:"need_restart,omitempty"`
		PgVersion        []string `json:"pg_version,omitempty"`
		DescriptionZn    *string  `json:"description_zn,omitempty"`
		DescriptionEn    *string  `json:"description_en,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "database", "default_version", "installed_version", "schema", "category", "need_restart", "pg_version", "description_zn", "description_en"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Database = all.Database
	o.DefaultVersion = all.DefaultVersion
	o.InstalledVersion = all.InstalledVersion
	o.Schema = all.Schema
	o.Category = all.Category
	o.NeedRestart = all.NeedRestart
	o.PgVersion = all.PgVersion
	o.DescriptionZn = all.DescriptionZn
	o.DescriptionEn = all.DescriptionEn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
