// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

// List of APIs:
//   - [AIApi.analyzeCluster]
//   - [accountApi.createAccount]
//   - [accountApi.createAccountOld]
//   - [accountApi.createMongoDBAccount]
//   - [accountApi.deleteAccount]
//   - [accountApi.deleteAccountOld]
//   - [accountApi.deleteMongoDBAccount]
//   - [accountApi.getRootAccountPassword]
//   - [accountApi.listAccounts]
//   - [accountApi.listAccountsOld]
//   - [accountApi.listMongoDBAccounts]
//   - [accountApi.updateAccount]
//   - [accountApi.updateAccountOld]
//   - [accountApi.updateAccountPrivileges]
//   - [accountApi.updateAccountPrivilegesOld]
//   - [adminUserApi.createAdminUser]
//   - [adminUserApi.deleteAdminUser]
//   - [adminUserApi.listAdminUsers]
//   - [adminUserApi.patchAdminUser]
//   - [adminUserApi.patchAdminUsersProfile]
//   - [adminUserApi.readAdminUser]
//   - [adminUserApi.updateAdminUserPassword]
//   - [adminUserApi.updateAdminUsersPassword]
//   - [alertApi.alertStatistic]
//   - [alertApi.batchCheckURLConnectivity]
//   - [alertConfigApi.getAlertConfig]
//   - [alertConfigApi.getAlertSMSConfig]
//   - [alertConfigApi.setAlertConfig]
//   - [alertConfigApi.updateAlertSMSConfig]
//   - [alertInhibitApi.createAlertInhibit]
//   - [alertInhibitApi.deleteAlertInhibit]
//   - [alertInhibitApi.getAlertInhibit]
//   - [alertInhibitApi.listAlertInhibits]
//   - [alertInhibitApi.patchAlertInhibit]
//   - [alertMetricsApi.listAlertMetrics]
//   - [alertObjectApi.listAlertObjects]
//   - [alertObjectApi.setAlertObjectStatus]
//   - [alertObjectApi.setAlertObjectsStatus]
//   - [alertReceiverApi.createAlertReceiver]
//   - [alertReceiverApi.deleteAlertReceiver]
//   - [alertReceiverApi.getAlertReceiver]
//   - [alertReceiverApi.listAlertReceivers]
//   - [alertReceiverApi.patchAlertReceiver]
//   - [alertRuleApi.createAlertRule]
//   - [alertRuleApi.deleteAlertRule]
//   - [alertRuleApi.getAlertRule]
//   - [alertRuleApi.listAlertRules]
//   - [alertRuleApi.updateAlertRule]
//   - [alertSMTPConfigApi.getAlertSMTPConfig]
//   - [alertSMTPConfigApi.updateAlertSMTPConfig]
//   - [alertStrategyApi.createAlertStrategy]
//   - [alertStrategyApi.deleteAlertStrategy]
//   - [alertStrategyApi.listAlertStrategies]
//   - [alertStrategyApi.patchAlertStrategy]
//   - [alertStrategyApi.updateAlertStrategy]
//   - [alertTemplateApi.createAlertTemplate]
//   - [alertTemplateApi.deleteAlertTemplate]
//   - [alertTemplateApi.getAlertTemplate]
//   - [alertTemplateApi.listAlertTemplates]
//   - [alertTemplateApi.patchAlertTemplate]
//   - [analyzeApi.analyzeBackup]
//   - [analyzeApi.analyzeClusterParam]
//   - [analyzeApi.analyzeClusterParameter]
//   - [analyzeApi.analyzeClusterRestore]
//   - [analyzeApi.analyzeLogs]
//   - [analyzeApi.analyzeOps]
//   - [analyzeApi.analyzeParam]
//   - [analyzeApi.analyzeParameter]
//   - [analyzeApi.analyzeService]
//   - [analyzeApi.analyzeSlowLogs]
//   - [autohealingApi.getAutohealing]
//   - [backupApi.createClusterBackup]
//   - [backupApi.deleteBackup]
//   - [backupApi.downloadBackup]
//   - [backupApi.downloadMultipleBackups]
//   - [backupApi.getBackup]
//   - [backupApi.getBackupLog]
//   - [backupApi.getBackupStats]
//   - [backupApi.getClusterBackupPolicy]
//   - [backupApi.listBackups]
//   - [backupApi.updateBackupPolicy]
//   - [backupApi.viewBackup]
//   - [backupMethodApi.getBackupMethod]
//   - [backupRepoApi.checkBackupRepo]
//   - [backupRepoApi.createBackupRepo]
//   - [backupRepoApi.deleteBackupRepo]
//   - [backupRepoApi.getBackupRepo]
//   - [backupRepoApi.listBackupRepoStats]
//   - [backupRepoApi.listBackupRepos]
//   - [backupRepoApi.updateBackupRepo]
//   - [backupRepoApi.viewBackupRepo]
//   - [classApi.batchClass]
//   - [classApi.createClass]
//   - [classApi.deleteClass]
//   - [classApi.listClasses]
//   - [classApi.patchClass]
//   - [clusterApi.createCluster]
//   - [clusterApi.deleteCluster]
//   - [clusterApi.describeClusterHaHistory]
//   - [clusterApi.getCluster]
//   - [clusterApi.getClusterByID]
//   - [clusterApi.getClusterInstanceLog]
//   - [clusterApi.getInstacesMetrics]
//   - [clusterApi.listCluster]
//   - [clusterApi.listClusters]
//   - [clusterApi.listEndpoints]
//   - [clusterApi.listInstance]
//   - [clusterApi.patchCluster]
//   - [clusterApi.queryClusterMetrics]
//   - [clusterAlertSwitchApi.getClusterAlertDisabled]
//   - [clusterAlertSwitchApi.setClusterAlertDisabled]
//   - [clusterLogApi.queryAuditLogs]
//   - [clusterLogApi.queryErrorLogs]
//   - [clusterLogApi.queryRunningLogs]
//   - [clusterLogApi.querySlowLogs]
//   - [clusterTagApi.createTag]
//   - [clusterTagApi.deleteTag]
//   - [clusterTagApi.getTags]
//   - [clusterTagApi.listOrgTags]
//   - [clusterTagApi.updateTag]
//   - [clusterTaskApi.getClusterTask]
//   - [clusterTaskApi.listClusterTasks]
//   - [databaseApi.createDatabase]
//   - [databaseApi.createDatabaseOld]
//   - [databaseApi.deleteDatabase]
//   - [databaseApi.deleteDatabaseOld]
//   - [databaseApi.listDatabases]
//   - [databaseApi.listDatabasesOld]
//   - [dmsApi.DataExport]
//   - [dmsApi.DataImport]
//   - [dmsApi.GetObjectInfo]
//   - [dmsApi.GetTaskList]
//   - [dmsApi.GetTaskProgress]
//   - [dmsApi.ListObjectNamesByType]
//   - [dmsApi.ListObjectTypesInSchema]
//   - [dmsApi.alterParameter]
//   - [dmsApi.closeSessions]
//   - [dmsApi.createDataSourceV2]
//   - [dmsApi.deleteDataSourceV2]
//   - [dmsApi.disableConsole]
//   - [dmsApi.enableConsole]
//   - [dmsApi.generateDDL]
//   - [dmsApi.getDataSourceV2]
//   - [dmsApi.getEngineAvailableConsole]
//   - [dmsApi.getSchemaList]
//   - [dmsApi.listDataSourceV2]
//   - [dmsApi.listParameters]
//   - [dmsApi.listQueryHistory]
//   - [dmsApi.listSessionsOld]
//   - [dmsApi.query]
//   - [dmsApi.showData]
//   - [dmsApi.sqlExplain]
//   - [dmsApi.tenantParameterHistory]
//   - [dmsApi.testDataSourceV2]
//   - [dmsApi.updateDataSourceV2]
//   - [engineApi.ListServiceVersion]
//   - [engineApi.createEngineVersion]
//   - [engineApi.deleteEngineVersion]
//   - [engineApi.engineAction]
//   - [engineApi.getEngineByNameInEnv]
//   - [engineApi.listAllEngines]
//   - [engineApi.listEngineVersions]
//   - [engineApi.listEnginesInEnv]
//   - [engineApi.patchEngineVersion]
//   - [engineLicenseApi.createEngineLicense]
//   - [engineLicenseApi.deleteEngineLicense]
//   - [engineLicenseApi.engineLicense]
//   - [engineLicenseApi.listEngineLicenses]
//   - [engineOptionApi.ListUpgradeableServiceVersion]
//   - [engineOptionApi.createEngineOption]
//   - [engineOptionApi.getEngineOption]
//   - [engineOptionApi.listEngineOptionHistory]
//   - [engineOptionApi.listEngineOptions]
//   - [engineOptionApi.updateEngineOption]
//   - [environmentApi.addNodes]
//   - [environmentApi.checkKubeconfig]
//   - [environmentApi.checkNodeSSHConfig]
//   - [environmentApi.cordonEnvironmentNode]
//   - [environmentApi.createEnvironment]
//   - [environmentApi.createNodeGroup]
//   - [environmentApi.createWorkflow]
//   - [environmentApi.deleteEnvironment]
//   - [environmentApi.deleteNodeGroup]
//   - [environmentApi.getEnvironment]
//   - [environmentApi.getEnvironmentBackupRepo]
//   - [environmentApi.getEnvironmentBootstrapManifests]
//   - [environmentApi.getEnvironmentKubeconfig]
//   - [environmentApi.getEnvironmentMetricsMonitorStats]
//   - [environmentApi.getEnvironmentModuleDetails]
//   - [environmentApi.getEnvironmentModuleInfo]
//   - [environmentApi.getEnvironmentModuleLogs]
//   - [environmentApi.getEnvironmentProvisioningProgress]
//   - [environmentApi.getEnvironmentStatus]
//   - [environmentApi.getEnvironmentStatusHistory]
//   - [environmentApi.getLatestEnvModuleVersion]
//   - [environmentApi.getNode]
//   - [environmentApi.getOptionalEnvironmentModules]
//   - [environmentApi.getWorkflow]
//   - [environmentApi.getWorkflowList]
//   - [environmentApi.getWorkflowLog]
//   - [environmentApi.listEnvNodeZone]
//   - [environmentApi.listEnvironment]
//   - [environmentApi.listEnvironmentObjectStorage]
//   - [environmentApi.listKubernetesNode]
//   - [environmentApi.listKubernetesStorageClass]
//   - [environmentApi.listNodeGroups]
//   - [environmentApi.listNodePod]
//   - [environmentApi.listNodes]
//   - [environmentApi.patchEnvironment]
//   - [environmentApi.patchNodeGroup]
//   - [environmentApi.preflightEnvironment]
//   - [environmentApi.removeNodeMaintenanceMode]
//   - [environmentApi.scaleDownEnvironmentNodes]
//   - [environmentApi.scaleOutEnvironmentNodes]
//   - [environmentApi.setNodeMaintenanceMode]
//   - [environmentApi.uncordonEnvironmentNode]
//   - [environmentApi.updateEnvironmentKubeconfig]
//   - [environmentApi.updateEnvironmentModule]
//   - [eventApi.getEvent]
//   - [eventApi.getEventFilter]
//   - [eventApi.listEvents]
//   - [featureApi.listFeature]
//   - [featureApi.listFeatureGroup]
//   - [featureApi.readFeature]
//   - [imageRegistryApi.createImageRegistry]
//   - [imageRegistryApi.deleteImageRegistry]
//   - [imageRegistryApi.getImageRegistry]
//   - [imageRegistryApi.listImageRegistries]
//   - [imageRegistryApi.patchImageRegistry]
//   - [inspectionApi.createAutoInspection]
//   - [inspectionApi.createInspectionScript]
//   - [inspectionApi.deleteInspectionScript]
//   - [inspectionApi.listAutoInspection]
//   - [inspectionApi.listInspectionScripts]
//   - [inspectionApi.listInspections]
//   - [inspectionApi.updateAutoInspection]
//   - [inspectionApi.updateInspection]
//   - [inspectionApi.updateInspectionScript]
//   - [instanceApi.listInstances]
//   - [instanceTypesApi.createInstanceType]
//   - [instanceTypesApi.deleteInstanceType]
//   - [instanceTypesApi.getInstanceTypeById]
//   - [instanceTypesApi.getInstanceTypes]
//   - [instanceTypesApi.updateInstanceType]
//   - [invitationApi.listInvitation]
//   - [ipWhitelistApi.createIPWhitelist]
//   - [ipWhitelistApi.deleteIPWhiteList]
//   - [ipWhitelistApi.listIPWhitelist]
//   - [ipWhitelistApi.updateIPWhitelist]
//   - [kafkaApi.createKafkaTopic]
//   - [kafkaApi.deleteKafkaConsumerGroup]
//   - [kafkaApi.deleteKafkaTopic]
//   - [kafkaApi.expandKafkaTopicPartitions]
//   - [kafkaApi.getKafkaBrokerConfigs]
//   - [kafkaApi.getKafkaBrokers]
//   - [kafkaApi.getKafkaConsumerGroupDescribe]
//   - [kafkaApi.getKafkaTopicBrokers]
//   - [kafkaApi.getKafkaTopicConfig]
//   - [kafkaApi.getKafkaTopicInfos]
//   - [kafkaApi.getKafkaTopicPartitions]
//   - [kafkaApi.getKafkaTopics]
//   - [kafkaApi.listKafkaConsumerGroups]
//   - [kafkaApi.listKafkaTopicConsumerGroups]
//   - [kafkaApi.listKafkaTopicConsumerOffsets]
//   - [kafkaApi.listKafkaTopicMessages]
//   - [kafkaApi.produceKafkaTopicMessage]
//   - [kafkaApi.resetKafkaTopicConsumerOffset]
//   - [kafkaApi.setKafkaTopicConfig]
//   - [kafkaApi.updateKafkaBrokerConfig]
//   - [licenseApi.getLicense]
//   - [licenseApi.updateLicense]
//   - [llmApi.addLLM]
//   - [llmApi.checkAPIKey]
//   - [llmApi.deleteLLM]
//   - [llmApi.getLLMByID]
//   - [llmApi.listAvailableModel]
//   - [llmApi.listLLM]
//   - [llmApi.updateLLM]
//   - [loadBalancerApi.checkLoadBalancer]
//   - [loadBalancerApi.getLoadBalancer]
//   - [loadBalancerApi.installLoadBalancer]
//   - [loadBalancerApi.uninstallLoadBalancer]
//   - [markClusterApi.markClusterRestoreCompleted]
//   - [metadbApi.deleteMetadbBackups]
//   - [metadbApi.getMetadbInstancesMetrics]
//   - [metadbApi.getPostgresClusterInfo]
//   - [metadbApi.listMetadbBackups]
//   - [metadbApi.metadbBackup]
//   - [metadbApi.metadbBackupConfig]
//   - [metadbApi.metadbBackupGetConfig]
//   - [metadbApi.metadbListInstances]
//   - [metadbApi.metadbRestore]
//   - [metadbApi.metadbVerticalScale]
//   - [metadbApi.metadbVolumeExpand]
//   - [metricsApi.getAggregateMetaData]
//   - [monitorDataSinkApi.createMonitorDataSink]
//   - [monitorDataSinkApi.deleteMonitorDataSink]
//   - [monitorDataSinkApi.listMonitorDataSinks]
//   - [monitorDataSinkApi.patchMonitorDataSink]
//   - [oceanbaseApi.getTenant]
//   - [oceanbaseApi.listTenants]
//   - [opsrequestApi.cancelOps]
//   - [opsrequestApi.clusterVolumeExpand]
//   - [opsrequestApi.customOps]
//   - [opsrequestApi.deleteOps]
//   - [opsrequestApi.exposeCluster]
//   - [opsrequestApi.getOpsRequestStatus]
//   - [opsrequestApi.horizontalScaleCluster]
//   - [opsrequestApi.listAvailableNodes]
//   - [opsrequestApi.promoteCluster]
//   - [opsrequestApi.rebuildInstance]
//   - [opsrequestApi.reconfigureCluster]
//   - [opsrequestApi.restartCluster]
//   - [opsrequestApi.specifyClusterIOQuotas]
//   - [opsrequestApi.startCluster]
//   - [opsrequestApi.stopCluster]
//   - [opsrequestApi.updateClusterLicense]
//   - [opsrequestApi.upgradeCluster]
//   - [opsrequestApi.verticalScaleCluster]
//   - [organizationApi.batchUpdateOrgParameters]
//   - [organizationApi.disableOrg]
//   - [organizationApi.enableOrg]
//   - [organizationApi.getOrgParameter]
//   - [organizationApi.listOrgMember]
//   - [organizationApi.listOrgParameters]
//   - [organizationApi.listOrganizations]
//   - [organizationApi.patchOrgParameter]
//   - [organizationApi.readOrg]
//   - [organizationApi.readOrgMember]
//   - [parameterApi.listParameterProps]
//   - [parameterApi.listParametersHistory]
//   - [parameterTemplateApi.createParameterTemplate]
//   - [parameterTemplateApi.deleteParameterTemplate]
//   - [parameterTemplateApi.exportParameterTemplateFromCluster]
//   - [parameterTemplateApi.getClusterParameterTemplate]
//   - [parameterTemplateApi.listParameterTemplates]
//   - [parameterTemplateApi.patchParameterTemplate]
//   - [parameterTemplateApi.readParameterTemplate]
//   - [parameterTemplateApi.updateClusterParameterTemplate]
//   - [platformParameterApi.getPlatformParameter]
//   - [platformParameterApi.listPlatformParameters]
//   - [platformParameterApi.updatePlatformParameters]
//   - [projectApi.createProject]
//   - [projectApi.deleteProject]
//   - [projectApi.listProjects]
//   - [providerApi.createCloudProvider]
//   - [providerApi.deleteCloudProvider]
//   - [providerApi.getCloudProvider]
//   - [providerApi.listCloudProviders]
//   - [providerApi.updateCloudProvider]
//   - [recycleBinClusterApi.deleteRecycleBinCluster]
//   - [recycleBinClusterApi.getRecycleBinCluster]
//   - [recycleBinClusterApi.listRecycleBinCluster]
//   - [recycleBinClusterApi.listRecycleBinClusters]
//   - [recycleBinClusterApi.restoreRecycleBinCluster]
//   - [redisApi.createRedisAccount]
//   - [redisApi.deleteRedisAccount]
//   - [redisApi.listRedisAccounts]
//   - [redisApi.updateRedisAccount]
//   - [regionApi.createRegion]
//   - [regionApi.deleteRegion]
//   - [regionApi.getRegion]
//   - [regionApi.listRegions]
//   - [regionApi.updateRegion]
//   - [regionGroupApi.createRegionGroup]
//   - [regionGroupApi.deleteRegionGroup]
//   - [regionGroupApi.listRegionGroups]
//   - [regionGroupApi.updateRegionGroup]
//   - [relationApi.createRelation]
//   - [relationApi.deleteRelation]
//   - [relationApi.listAvailableClustersForRelation]
//   - [relationApi.listRelatedClusters]
//   - [resourceStatsApi.getResourceStats]
//   - [resourceStatsApi.listInstancesResourceStats]
//   - [resourceStatsApi.listNodesResourceStats]
//   - [restoreApi.GetRestoreLog]
//   - [restoreApi.deleteRestoreObject]
//   - [restoreApi.doRestore]
//   - [restoreApi.getRestoreTimeRange]
//   - [restoreApi.listClusterRestore]
//   - [restoreApi.listRestores]
//   - [restoreApi.restoreCluster]
//   - [roleApi.getRoleByName]
//   - [roleApi.listPermissions]
//   - [roleApi.listRolePermissions]
//   - [roleApi.listRoles]
//   - [sessionApi.killSession]
//   - [sessionApi.listSessions]
//   - [storageApi.checkStorage]
//   - [storageApi.createStorage]
//   - [storageApi.deleteStorage]
//   - [storageApi.getStorage]
//   - [storageApi.importBackup]
//   - [storageApi.listStorageProviders]
//   - [storageApi.listStorages]
//   - [storageApi.updateStorage]
//   - [storageClassApi.createStorageClass]
//   - [storageClassApi.deleteStorageClass]
//   - [storageClassApi.getStorageClass]
//   - [storageClassApi.listStorageClassPvcs]
//   - [storageClassApi.listStorageClasses]
//   - [storageClassApi.listStorageProvisioners]
//   - [storageClassApi.updateStorageClass]
//   - [taskApi.getTask]
//   - [taskApi.getTaskLog]
//   - [taskApi.listTask]
//   - [taskApi.retryTask]
//   - [taskApi.stopTask]
//   - [tlsApi.getTLSCertificate]
//   - [tlsApi.tlsSwitcher]
//   - [userApi.createUserApikey]
//   - [userApi.deleteApikey]
//   - [userApi.patchAPIkey]
//   - [userApi.readUserApikeys]
//   - [vipPoolApi.createVIPPool]
//   - [vipPoolApi.deleteVIPPool]
//   - [vipPoolApi.listVIPPool]
//   - [zoneApi.createZone]
//   - [zoneApi.deleteZone]
//   - [zoneApi.getZones]
//   - [zoneApi.listZones]
//   - [zoneApi.updateZone]
package admin
