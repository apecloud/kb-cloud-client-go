// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

// List of APIs:
//   - [AIApi.chatBIV2SSE]
//   - [AIApi.chatV2SSE]
//   - [AIApi.deleteConversation]
//   - [AIApi.getConversationMessages]
//   - [AIApi.listConversations]
//   - [AIApi.newAIConversation]
//   - [accountApi.createAccount]
//   - [accountApi.createAccountOld]
//   - [accountApi.createMongoDBAccount]
//   - [accountApi.deleteAccount]
//   - [accountApi.deleteAccountOld]
//   - [accountApi.deleteMongoDBAccount]
//   - [accountApi.deleteMssqlAccount]
//   - [accountApi.getRootAccountPassword]
//   - [accountApi.listAccounts]
//   - [accountApi.listAccountsOld]
//   - [accountApi.listMongoDBAccounts]
//   - [accountApi.updateAccount]
//   - [accountApi.updateAccountOld]
//   - [accountApi.updateAccountPrivileges]
//   - [accountApi.updateAccountPrivilegesOld]
//   - [accountApi.updateMssqlAccount]
//   - [alertApi.alertStatistic]
//   - [alertApi.batchCheckURLConnectivity]
//   - [alertConfigApi.getAlertConfig]
//   - [alertConfigApi.setAlertConfig]
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
//   - [alertRuleApi.downloadOrgAlertRuleFile]
//   - [alertRuleApi.getAlertRule]
//   - [alertRuleApi.listAlertRules]
//   - [alertRuleApi.restoreOrgAlertRuleToDefault]
//   - [alertRuleApi.updateAlertRule]
//   - [alertRuleApi.updateRuleConfig]
//   - [alertStrategyApi.createAlertStrategy]
//   - [alertStrategyApi.deleteAlertStrategy]
//   - [alertStrategyApi.listAlertStrategies]
//   - [alertStrategyApi.patchAlertStrategy]
//   - [alertStrategyApi.updateAlertStrategy]
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
//   - [backupRepoApi.getBackupRepo]
//   - [backupRepoApi.listBackupRepos]
//   - [benchmarkApi.createPgbench]
//   - [benchmarkApi.createSysbench]
//   - [benchmarkApi.createTpcc]
//   - [benchmarkApi.createYcsb]
//   - [benchmarkApi.deleteBenchmark]
//   - [benchmarkApi.getBenchmark]
//   - [benchmarkApi.listBenchmark]
//   - [classApi.listClasses]
//   - [classTypesApi.getClassType]
//   - [classTypesApi.getClassTypes]
//   - [clusterApi.createCluster]
//   - [clusterApi.deleteCluster]
//   - [clusterApi.describeClusterHaHistory]
//   - [clusterApi.getCluster]
//   - [clusterApi.getClusterByID]
//   - [clusterApi.getClusterInstanceLog]
//   - [clusterApi.getInstacesMetrics]
//   - [clusterApi.listCluster]
//   - [clusterApi.listEndpoints]
//   - [clusterApi.listInstance]
//   - [clusterApi.listInstanceEvents]
//   - [clusterApi.patchCluster]
//   - [clusterApi.queryClusterMetrics]
//   - [clusterApi.validateClusterCreation]
//   - [clusterAlertSwitchApi.getClusterAlertDisabled]
//   - [clusterAlertSwitchApi.setClusterAlertDisabled]
//   - [clusterLogApi.queryAuditLogs]
//   - [clusterLogApi.queryErrorLogs]
//   - [clusterLogApi.queryPodLogs]
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
//   - [databaseApi.getDatabase]
//   - [databaseApi.listDatabases]
//   - [databaseApi.listDatabasesOld]
//   - [databaseApi.updateDatabaseConfig]
//   - [disasterRecoveryApi.createDisasterRecovery]
//   - [disasterRecoveryApi.deleteDisasterRecovery]
//   - [disasterRecoveryApi.getDisasterRecoveryHistory]
//   - [disasterRecoveryApi.getDisasterRecoveryStatus]
//   - [disasterRecoveryApi.listDisasterRecovery]
//   - [disasterRecoveryApi.promoteDisasterRecovery]
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
//   - [engineApi.ListUpgradeableServiceVersion]
//   - [engineApi.listEngineLicenses]
//   - [engineApi.listEngineResourceConstraints]
//   - [engineApi.listEnginesInEnv]
//   - [engineOptionApi.getEngineOption]
//   - [engineOptionApi.listEngineOptions]
//   - [environmentApi.getEnvironment]
//   - [environmentApi.getEnvironmentModuleInfo]
//   - [environmentApi.listEnvNodeZone]
//   - [environmentApi.listEnvironment]
//   - [environmentApi.listNodeGroup]
//   - [eventApi.getEventFilter]
//   - [eventApi.getOrgEvent]
//   - [eventApi.listOrgEvents]
//   - [featureApi.listFeature]
//   - [featureApi.readFeature]
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
//   - [keyApi.listKeys]
//   - [llmApi.ListLLMForOrg]
//   - [llmApi.addLLMForOrg]
//   - [llmApi.checkAPIKeyForOrg]
//   - [llmApi.deleteLLM]
//   - [llmApi.getLLMByIDInOrg]
//   - [llmApi.listAvailableModelInOrg]
//   - [llmApi.updateLLM]
//   - [loadBalancerApi.getLoadBalancer]
//   - [markClusterApi.markClusterRestoreCompleted]
//   - [memberApi.addOrgMember]
//   - [memberApi.deleteOrgMember]
//   - [memberApi.listOrgMember]
//   - [memberApi.listOrgMemberPermission]
//   - [memberApi.patchOrgMember]
//   - [memberApi.readOrgMember]
//   - [mssqlApi.manageMssqlTDEDatabase]
//   - [oceanbaseApi.getTenant]
//   - [oceanbaseApi.listTenants]
//   - [opsrequestApi.cancelOps]
//   - [opsrequestApi.clusterVolumeExpand]
//   - [opsrequestApi.customOps]
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
//   - [organizationApi.createOrg]
//   - [organizationApi.deleteOrg]
//   - [organizationApi.freezeMember]
//   - [organizationApi.listOrg]
//   - [organizationApi.patchOrg]
//   - [organizationApi.readOrg]
//   - [organizationApi.unfreezeMember]
//   - [organizationParameterApi.batchUpdateOrgParameters]
//   - [organizationParameterApi.getOrgParameter]
//   - [organizationParameterApi.listOrgParameters]
//   - [organizationParameterApi.patchOrgParameter]
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
//   - [projectApi.listProjects]
//   - [rdbmsApi.generateDemoData]
//   - [rdbmsApi.getDatabaseOptions]
//   - [recycleBinClusterApi.deleteRecycleBinCluster]
//   - [recycleBinClusterApi.getRecycleBinCluster]
//   - [recycleBinClusterApi.listRecycleBinCluster]
//   - [recycleBinClusterApi.restoreRecycleBinCluster]
//   - [redisApi.createRedisAccount]
//   - [redisApi.deleteRedisAccount]
//   - [redisApi.listRedisAccounts]
//   - [redisApi.updateRedisAccount]
//   - [relationApi.createRelation]
//   - [relationApi.deleteRelation]
//   - [relationApi.listAvailableClustersForRelation]
//   - [relationApi.listRelatedClusters]
//   - [resourceStatsApi.listNodesResourceStats]
//   - [restoreApi.GetRestoreLog]
//   - [restoreApi.deleteRestoreObject]
//   - [restoreApi.doRestore]
//   - [restoreApi.getRestoreTimeRange]
//   - [restoreApi.listClusterRestore]
//   - [restoreApi.listRestores]
//   - [restoreApi.restoreCluster]
//   - [roleApi.batchAddRolePermissions]
//   - [roleApi.batchRemoveRolePermissions]
//   - [roleApi.createRole]
//   - [roleApi.deleteRoleByName]
//   - [roleApi.getRoleByName]
//   - [roleApi.listPermissions]
//   - [roleApi.listRolePermissions]
//   - [roleApi.listRoles]
//   - [roleApi.updateRoleByName]
//   - [sessionApi.killSession]
//   - [sessionApi.listSessions]
//   - [storageApi.createStorage]
//   - [storageApi.deleteStorage]
//   - [storageApi.getStorage]
//   - [storageApi.listStorages]
//   - [storageApi.updateStorage]
//   - [storageClassApi.getStorageClassStats]
//   - [tlsApi.getTLSCertificate]
//   - [tlsApi.tlsSwitcher]
//   - [userApi.createUserApikey]
//   - [userApi.deleteApikey]
//   - [userApi.getAuthenticatedUser]
//   - [userApi.patchAPIkey]
//   - [userApi.patchAuthenticatedUser]
//   - [userApi.phoneVerification]
//   - [userApi.readUserApikeys]
//   - [userApi.updateAuthenticatedUserPassword]
package kbcloud
