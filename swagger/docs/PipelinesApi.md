# \PipelinesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentVariable**](PipelinesApi.md#CreateDeploymentVariable) | **Post** /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables | Create a variable for an environment
[**CreatePipelineForRepository**](PipelinesApi.md#CreatePipelineForRepository) | **Post** /repositories/{workspace}/{repo_slug}/pipelines/ | Run a pipeline
[**CreatePipelineVariableForTeam**](PipelinesApi.md#CreatePipelineVariableForTeam) | **Post** /teams/{username}/pipelines_config/variables/ | Create a variable for a user
[**CreatePipelineVariableForUser**](PipelinesApi.md#CreatePipelineVariableForUser) | **Post** /users/{selected_user}/pipelines_config/variables/ | Create a variable for a user
[**CreatePipelineVariableForWorkspace**](PipelinesApi.md#CreatePipelineVariableForWorkspace) | **Post** /workspaces/{workspace}/pipelines-config/variables | Create a variable for a workspace
[**CreateRepositoryPipelineKnownHost**](PipelinesApi.md#CreateRepositoryPipelineKnownHost) | **Post** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/ | Create a known host
[**CreateRepositoryPipelineSchedule**](PipelinesApi.md#CreateRepositoryPipelineSchedule) | **Post** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/ | Create a schedule
[**CreateRepositoryPipelineVariable**](PipelinesApi.md#CreateRepositoryPipelineVariable) | **Post** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/ | Create a variable for a repository
[**DeleteDeploymentVariable**](PipelinesApi.md#DeleteDeploymentVariable) | **Delete** /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables/{variable_uuid} | Delete a variable for an environment
[**DeletePipelineVariableForTeam**](PipelinesApi.md#DeletePipelineVariableForTeam) | **Delete** /teams/{username}/pipelines_config/variables/{variable_uuid} | Delete a variable for a team
[**DeletePipelineVariableForUser**](PipelinesApi.md#DeletePipelineVariableForUser) | **Delete** /users/{selected_user}/pipelines_config/variables/{variable_uuid} | Delete a variable for a user
[**DeletePipelineVariableForWorkspace**](PipelinesApi.md#DeletePipelineVariableForWorkspace) | **Delete** /workspaces/{workspace}/pipelines-config/variables/{variable_uuid} | Delete a variable for a workspace
[**DeleteRepositoryPipelineCache**](PipelinesApi.md#DeleteRepositoryPipelineCache) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines-config/caches/{cache_uuid} | Delete a cache
[**DeleteRepositoryPipelineCaches**](PipelinesApi.md#DeleteRepositoryPipelineCaches) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines-config/caches | Delete caches
[**DeleteRepositoryPipelineKeyPair**](PipelinesApi.md#DeleteRepositoryPipelineKeyPair) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/key_pair | Delete SSH key pair
[**DeleteRepositoryPipelineKnownHost**](PipelinesApi.md#DeleteRepositoryPipelineKnownHost) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/{known_host_uuid} | Delete a known host
[**DeleteRepositoryPipelineSchedule**](PipelinesApi.md#DeleteRepositoryPipelineSchedule) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid} | Delete a schedule
[**DeleteRepositoryPipelineVariable**](PipelinesApi.md#DeleteRepositoryPipelineVariable) | **Delete** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid} | Delete a variable for a repository
[**GetDeploymentVariables**](PipelinesApi.md#GetDeploymentVariables) | **Get** /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables | List variables for an environment
[**GetOIDCConfiguration**](PipelinesApi.md#GetOIDCConfiguration) | **Get** /workspaces/{workspace}/pipelines-config/identity/oidc/.well-known/openid-configuration | Get OpenID configuration for OIDC in Pipelines
[**GetOIDCKeys**](PipelinesApi.md#GetOIDCKeys) | **Get** /workspaces/{workspace}/pipelines-config/identity/oidc/keys.json | Get keys for OIDC in Pipelines
[**GetPipelineContainerLog**](PipelinesApi.md#GetPipelineContainerLog) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/logs/{log_uuid} | Get the logs for the build container or a service container for a given step of a pipeline.
[**GetPipelineForRepository**](PipelinesApi.md#GetPipelineForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid} | Get a pipeline
[**GetPipelineStepForRepository**](PipelinesApi.md#GetPipelineStepForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid} | Get a step of a pipeline
[**GetPipelineStepLogForRepository**](PipelinesApi.md#GetPipelineStepLogForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/log | Get log file for a step
[**GetPipelineStepsForRepository**](PipelinesApi.md#GetPipelineStepsForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/ | List steps for a pipeline
[**GetPipelineTestReportTestCaseReasons**](PipelinesApi.md#GetPipelineTestReportTestCaseReasons) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/test_reports/test_cases/{test_case_uuid}/test_case_reasons | Get test case reasons (output) for a given test case in a step of a pipeline.
[**GetPipelineTestReportTestCases**](PipelinesApi.md#GetPipelineTestReportTestCases) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/test_reports/test_cases | Get test cases for a given step of a pipeline.
[**GetPipelineTestReports**](PipelinesApi.md#GetPipelineTestReports) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/steps/{step_uuid}/test_reports | Get a summary of test reports for a given step of a pipeline.
[**GetPipelineVariableForTeam**](PipelinesApi.md#GetPipelineVariableForTeam) | **Get** /teams/{username}/pipelines_config/variables/{variable_uuid} | Get a variable for a team
[**GetPipelineVariableForUser**](PipelinesApi.md#GetPipelineVariableForUser) | **Get** /users/{selected_user}/pipelines_config/variables/{variable_uuid} | Get a variable for a user
[**GetPipelineVariableForWorkspace**](PipelinesApi.md#GetPipelineVariableForWorkspace) | **Get** /workspaces/{workspace}/pipelines-config/variables/{variable_uuid} | Get variable for a workspace
[**GetPipelineVariablesForTeam**](PipelinesApi.md#GetPipelineVariablesForTeam) | **Get** /teams/{username}/pipelines_config/variables/ | List variables for an account
[**GetPipelineVariablesForUser**](PipelinesApi.md#GetPipelineVariablesForUser) | **Get** /users/{selected_user}/pipelines_config/variables/ | List variables for a user
[**GetPipelineVariablesForWorkspace**](PipelinesApi.md#GetPipelineVariablesForWorkspace) | **Get** /workspaces/{workspace}/pipelines-config/variables | List variables for a workspace
[**GetPipelinesForRepository**](PipelinesApi.md#GetPipelinesForRepository) | **Get** /repositories/{workspace}/{repo_slug}/pipelines/ | List pipelines
[**GetRepositoryPipelineCacheContentURI**](PipelinesApi.md#GetRepositoryPipelineCacheContentURI) | **Get** /repositories/{workspace}/{repo_slug}/pipelines-config/caches/{cache_uuid}/content-uri | Get cache content URI
[**GetRepositoryPipelineCaches**](PipelinesApi.md#GetRepositoryPipelineCaches) | **Get** /repositories/{workspace}/{repo_slug}/pipelines-config/caches/ | List caches
[**GetRepositoryPipelineConfig**](PipelinesApi.md#GetRepositoryPipelineConfig) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config | Get configuration
[**GetRepositoryPipelineKnownHost**](PipelinesApi.md#GetRepositoryPipelineKnownHost) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/{known_host_uuid} | Get a known host
[**GetRepositoryPipelineKnownHosts**](PipelinesApi.md#GetRepositoryPipelineKnownHosts) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/ | List known hosts
[**GetRepositoryPipelineSchedule**](PipelinesApi.md#GetRepositoryPipelineSchedule) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid} | Get a schedule
[**GetRepositoryPipelineScheduleExecutions**](PipelinesApi.md#GetRepositoryPipelineScheduleExecutions) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid}/executions/ | List executions of a schedule
[**GetRepositoryPipelineSchedules**](PipelinesApi.md#GetRepositoryPipelineSchedules) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/ | List schedules
[**GetRepositoryPipelineSshKeyPair**](PipelinesApi.md#GetRepositoryPipelineSshKeyPair) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/key_pair | Get SSH key pair
[**GetRepositoryPipelineVariable**](PipelinesApi.md#GetRepositoryPipelineVariable) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid} | Get a variable for a repository
[**GetRepositoryPipelineVariables**](PipelinesApi.md#GetRepositoryPipelineVariables) | **Get** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/ | List variables for a repository
[**StopPipeline**](PipelinesApi.md#StopPipeline) | **Post** /repositories/{workspace}/{repo_slug}/pipelines/{pipeline_uuid}/stopPipeline | Stop a pipeline
[**UpdateDeploymentVariable**](PipelinesApi.md#UpdateDeploymentVariable) | **Put** /repositories/{workspace}/{repo_slug}/deployments_config/environments/{environment_uuid}/variables/{variable_uuid} | Update a variable for an environment
[**UpdatePipelineVariableForTeam**](PipelinesApi.md#UpdatePipelineVariableForTeam) | **Put** /teams/{username}/pipelines_config/variables/{variable_uuid} | Update a variable for a team
[**UpdatePipelineVariableForUser**](PipelinesApi.md#UpdatePipelineVariableForUser) | **Put** /users/{selected_user}/pipelines_config/variables/{variable_uuid} | Update a variable for a user
[**UpdatePipelineVariableForWorkspace**](PipelinesApi.md#UpdatePipelineVariableForWorkspace) | **Put** /workspaces/{workspace}/pipelines-config/variables/{variable_uuid} | Update variable for a workspace
[**UpdateRepositoryBuildNumber**](PipelinesApi.md#UpdateRepositoryBuildNumber) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/build_number | Update the next build number
[**UpdateRepositoryPipelineConfig**](PipelinesApi.md#UpdateRepositoryPipelineConfig) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config | Update configuration
[**UpdateRepositoryPipelineKeyPair**](PipelinesApi.md#UpdateRepositoryPipelineKeyPair) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/key_pair | Update SSH key pair
[**UpdateRepositoryPipelineKnownHost**](PipelinesApi.md#UpdateRepositoryPipelineKnownHost) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/ssh/known_hosts/{known_host_uuid} | Update a known host
[**UpdateRepositoryPipelineSchedule**](PipelinesApi.md#UpdateRepositoryPipelineSchedule) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/schedules/{schedule_uuid} | Update a schedule
[**UpdateRepositoryPipelineVariable**](PipelinesApi.md#UpdateRepositoryPipelineVariable) | **Put** /repositories/{workspace}/{repo_slug}/pipelines_config/variables/{variable_uuid} | Update a variable for a repository


# **CreateDeploymentVariable**
> DeploymentVariable CreateDeploymentVariable(ctx, workspace, repoSlug, environmentUuid, body)
Create a variable for an environment

Create a deployment environment level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **environmentUuid** | **string**| The environment. | 
  **body** | [**DeploymentVariable**](DeploymentVariable.md)| The variable to create | 

### Return type

[**DeploymentVariable**](deployment_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePipelineForRepository**
> Pipeline CreatePipelineForRepository(ctx, workspace, repoSlug, body)
Run a pipeline

Endpoint to create and initiate a pipeline. There are a couple of different options to initiate a pipeline, where the payload of the request will determine which type of pipeline will be instantiated. # Trigger a Pipeline for a branch One way to trigger pipelines is by specifying the branch for which you want to trigger a pipeline. The specified branch will be used to determine which pipeline definition from the `bitbucket-pipelines.yml` file will be applied to initiate the pipeline. The pipeline will then do a clone of the repository and checkout the latest revision of the specified branch.  ### Example  ``` $ curl -X POST -is -u username:password \\   -H 'Content-Type: application/json' \\  https://api.bitbucket.org/2.0/repositories/jeroendr/meat-demo2/pipelines/ \\   -d '   {     \"target\": {       \"ref_type\": \"branch\",       \"type\": \"pipeline_ref_target\",       \"ref_name\": \"master\"     }   }' ``` # Trigger a Pipeline for a commit on a branch or tag You can initiate a pipeline for a specific commit and in the context of a specified reference (e.g. a branch, tag or bookmark). The specified reference will be used to determine which pipeline definition from the bitbucket-pipelines.yml file will be applied to initiate the pipeline. The pipeline will clone the repository and then do a checkout the specified reference.  The following reference types are supported:  * `branch` * `named_branch` * `bookmark`  * `tag`  ### Example  ``` $ curl -X POST -is -u username:password \\   -H 'Content-Type: application/json' \\   https://api.bitbucket.org/2.0/repositories/jeroendr/meat-demo2/pipelines/ \\   -d '   {     \"target\": {       \"commit\": {         \"type\": \"commit\",         \"hash\": \"ce5b7431602f7cbba007062eeb55225c6e18e956\"       },       \"ref_type\": \"branch\",       \"type\": \"pipeline_ref_target\",       \"ref_name\": \"master\"     }   }' ``` # Trigger a specific pipeline definition for a commit You can trigger a specific pipeline that is defined in your `bitbucket-pipelines.yml` file for a specific commit. In addition to the commit revision, you specify the type and pattern of the selector that identifies the pipeline definition. The resulting pipeline will then clone the repository and checkout the specified revision.  ### Example  ``` $ curl -X POST -is -u username:password \\   -H 'Content-Type: application/json' \\  https://api.bitbucket.org/2.0/repositories/jeroendr/meat-demo2/pipelines/ \\  -d '   {      \"target\": {       \"commit\": {          \"hash\":\"a3c4e02c9a3755eccdc3764e6ea13facdf30f923\",          \"type\":\"commit\"        },         \"selector\": {            \"type\":\"custom\",               \"pattern\":\"Deploy to production\"           },         \"type\":\"pipeline_commit_target\"    }   }' ``` # Trigger a specific pipeline definition for a commit on a branch or tag You can trigger a specific pipeline that is defined in your `bitbucket-pipelines.yml` file for a specific commit in the context of a specified reference. In addition to the commit revision, you specify the type and pattern of the selector that identifies the pipeline definition, as well as the reference information. The resulting pipeline will then clone the repository a checkout the specified reference.  ### Example  ``` $ curl -X POST -is -u username:password \\   -H 'Content-Type: application/json' \\  https://api.bitbucket.org/2.0/repositories/jeroendr/meat-demo2/pipelines/ \\  -d '   {      \"target\": {       \"commit\": {          \"hash\":\"a3c4e02c9a3755eccdc3764e6ea13facdf30f923\",          \"type\":\"commit\"        },        \"selector\": {           \"type\": \"custom\",           \"pattern\": \"Deploy to production\"        },        \"type\": \"pipeline_ref_target\",        \"ref_name\": \"master\",        \"ref_type\": \"branch\"      }   }' ```   # Trigger a custom pipeline with variables In addition to triggering a custom pipeline that is defined in your `bitbucket-pipelines.yml` file as shown in the examples above, you can specify variables that will be available for your build. In the request, provide a list of variables, specifying the following for each variable: key, value, and whether it should be secured or not (this field is optional and defaults to not secured).  ### Example  ``` $ curl -X POST -is -u username:password \\   -H 'Content-Type: application/json' \\  https://api.bitbucket.org/2.0/repositories/{workspace}/{repo_slug}/pipelines/ \\  -d '   {     \"target\": {       \"type\": \"pipeline_ref_target\",       \"ref_type\": \"branch\",       \"ref_name\": \"master\",       \"selector\": {         \"type\": \"custom\",         \"pattern\": \"Deploy to production\"       }     },     \"variables\": [       {         \"key\": \"var1key\",         \"value\": \"var1value\",         \"secured\": true       },       {         \"key\": \"var2key\",         \"value\": \"var2value\"       }     ]   }' ```  # Trigger a pull request pipeline  You can also initiate a pipeline for a specific pull request.  ### Example  ``` $ curl -X POST -is -u username:password \\   -H 'Content-Type: application/json' \\  https://api.bitbucket.org/2.0/repositories/{workspace}/{repo_slug}/pipelines/ \\  -d '   {  \"target\": {       \"type\": \"pipeline_pullrequest_target\",    \"source\": \"pull-request-branch\",       \"destination\": \"master\",       \"destination_commit\": {         \"hash\" : \"9f848b7\"       },       \"commit\": {        \"hash\" : \"1a372fc\"       },       \"pullrequest\" : {        \"id\" : \"3\"       },    \"selector\": {         \"type\": \"pull-requests\",         \"pattern\": \"**\"       }     }   }' ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **body** | [**Pipeline**](Pipeline.md)| The pipeline to initiate. | 

### Return type

[**Pipeline**](pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePipelineVariableForTeam**
> PipelineVariable CreatePipelineVariableForTeam(ctx, username, optional)
Create a variable for a user

Create an account level variable. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
 **optional** | ***PipelinesApiCreatePipelineVariableForTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiCreatePipelineVariableForTeamOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PipelineVariable**](PipelineVariable.md)| The variable to create. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePipelineVariableForUser**
> PipelineVariable CreatePipelineVariableForUser(ctx, selectedUser, optional)
Create a variable for a user

Create a user level variable. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
 **optional** | ***PipelinesApiCreatePipelineVariableForUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiCreatePipelineVariableForUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PipelineVariable**](PipelineVariable.md)| The variable to create. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePipelineVariableForWorkspace**
> PipelineVariable CreatePipelineVariableForWorkspace(ctx, workspace, optional)
Create a variable for a workspace

Create a workspace level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
 **optional** | ***PipelinesApiCreatePipelineVariableForWorkspaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelinesApiCreatePipelineVariableForWorkspaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PipelineVariable**](PipelineVariable.md)| The variable to create. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepositoryPipelineKnownHost**
> PipelineKnownHost CreateRepositoryPipelineKnownHost(ctx, workspace, repoSlug, body)
Create a known host

Create a repository level known host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **body** | [**PipelineKnownHost**](PipelineKnownHost.md)| The known host to create. | 

### Return type

[**PipelineKnownHost**](pipeline_known_host.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepositoryPipelineSchedule**
> PipelineSchedule CreateRepositoryPipelineSchedule(ctx, workspace, repoSlug, body)
Create a schedule

Create a schedule for the given repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **body** | [**PipelineSchedule**](PipelineSchedule.md)| The schedule to create. | 

### Return type

[**PipelineSchedule**](pipeline_schedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepositoryPipelineVariable**
> PipelineVariable CreateRepositoryPipelineVariable(ctx, workspace, repoSlug, body)
Create a variable for a repository

Create a repository level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **body** | [**PipelineVariable**](PipelineVariable.md)| The variable to create. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeploymentVariable**
> DeleteDeploymentVariable(ctx, workspace, repoSlug, environmentUuid, variableUuid)
Delete a variable for an environment

Delete a deployment environment level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **environmentUuid** | **string**| The environment. | 
  **variableUuid** | **string**| The UUID of the variable to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePipelineVariableForTeam**
> DeletePipelineVariableForTeam(ctx, username, variableUuid)
Delete a variable for a team

Delete a team level variable. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **variableUuid** | **string**| The UUID of the variable to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePipelineVariableForUser**
> DeletePipelineVariableForUser(ctx, selectedUser, variableUuid)
Delete a variable for a user

Delete an account level variable. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
  **variableUuid** | **string**| The UUID of the variable to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePipelineVariableForWorkspace**
> DeletePipelineVariableForWorkspace(ctx, workspace, variableUuid)
Delete a variable for a workspace

Delete a workspace level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **variableUuid** | **string**| The UUID of the variable to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryPipelineCache**
> DeleteRepositoryPipelineCache(ctx, workspace, repoSlug, cacheUuid)
Delete a cache

Delete a repository cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **cacheUuid** | **string**| The UUID of the cache to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryPipelineCaches**
> DeleteRepositoryPipelineCaches(ctx, workspace, repoSlug, name)
Delete caches

Delete repository cache versions by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **name** | **string**| The cache name. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryPipelineKeyPair**
> DeleteRepositoryPipelineKeyPair(ctx, workspace, repoSlug)
Delete SSH key pair

Delete the repository SSH key pair.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryPipelineKnownHost**
> DeleteRepositoryPipelineKnownHost(ctx, workspace, repoSlug, knownHostUuid)
Delete a known host

Delete a repository level known host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **knownHostUuid** | **string**| The UUID of the known host to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryPipelineSchedule**
> DeleteRepositoryPipelineSchedule(ctx, workspace, repoSlug, scheduleUuid)
Delete a schedule

Delete a schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **scheduleUuid** | **string**| The uuid of the schedule. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryPipelineVariable**
> DeleteRepositoryPipelineVariable(ctx, workspace, repoSlug, variableUuid)
Delete a variable for a repository

Delete a repository level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **variableUuid** | **string**| The UUID of the variable to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeploymentVariables**
> PaginatedDeploymentVariable GetDeploymentVariables(ctx, workspace, repoSlug, environmentUuid)
List variables for an environment

Find deployment environment level variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **environmentUuid** | **string**| The environment. | 

### Return type

[**PaginatedDeploymentVariable**](paginated_deployment_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOIDCConfiguration**
> GetOIDCConfiguration(ctx, workspace)
Get OpenID configuration for OIDC in Pipelines

This is part of OpenID Connect for Pipelines, see https://support.atlassian.com/bitbucket-cloud/docs/integrate-pipelines-with-resource-servers-using-oidc/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOIDCKeys**
> GetOIDCKeys(ctx, workspace)
Get keys for OIDC in Pipelines

This is part of OpenID Connect for Pipelines, see https://support.atlassian.com/bitbucket-cloud/docs/integrate-pipelines-with-resource-servers-using-oidc/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineContainerLog**
> GetPipelineContainerLog(ctx, workspace, repoSlug, pipelineUuid, stepUuid, logUuid)
Get the logs for the build container or a service container for a given step of a pipeline.

Retrieve the log file for a build container or service container.  This endpoint supports (and encourages!) the use of [HTTP Range requests](https://tools.ietf.org/html/rfc7233) to deal with potentially very large log files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The UUID of the pipeline. | 
  **stepUuid** | **string**| The UUID of the step. | 
  **logUuid** | **string**| For the main build container specify the step UUID; for a service container specify the service container UUID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineForRepository**
> Pipeline GetPipelineForRepository(ctx, workspace, repoSlug, pipelineUuid)
Get a pipeline

Retrieve a specified pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The pipeline UUID. | 

### Return type

[**Pipeline**](pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineStepForRepository**
> PipelineStep GetPipelineStepForRepository(ctx, workspace, repoSlug, pipelineUuid, stepUuid)
Get a step of a pipeline

Retrieve a given step of a pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The UUID of the pipeline. | 
  **stepUuid** | **string**| The UUID of the step. | 

### Return type

[**PipelineStep**](pipeline_step.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineStepLogForRepository**
> GetPipelineStepLogForRepository(ctx, workspace, repoSlug, pipelineUuid, stepUuid)
Get log file for a step

Retrieve the log file for a given step of a pipeline.  This endpoint supports (and encourages!) the use of [HTTP Range requests](https://tools.ietf.org/html/rfc7233) to deal with potentially very large log files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The UUID of the pipeline. | 
  **stepUuid** | **string**| The UUID of the step. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineStepsForRepository**
> PaginatedPipelineSteps GetPipelineStepsForRepository(ctx, workspace, repoSlug, pipelineUuid)
List steps for a pipeline

Find steps for the given pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The UUID of the pipeline. | 

### Return type

[**PaginatedPipelineSteps**](paginated_pipeline_steps.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineTestReportTestCaseReasons**
> GetPipelineTestReportTestCaseReasons(ctx, workspace, repoSlug, pipelineUuid, stepUuid, testCaseUuid)
Get test case reasons (output) for a given test case in a step of a pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The UUID of the pipeline. | 
  **stepUuid** | **string**| The UUID of the step. | 
  **testCaseUuid** | **string**| The UUID of the test case. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineTestReportTestCases**
> GetPipelineTestReportTestCases(ctx, workspace, repoSlug, pipelineUuid, stepUuid)
Get test cases for a given step of a pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The UUID of the pipeline. | 
  **stepUuid** | **string**| The UUID of the step. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineTestReports**
> GetPipelineTestReports(ctx, workspace, repoSlug, pipelineUuid, stepUuid)
Get a summary of test reports for a given step of a pipeline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The UUID of the pipeline. | 
  **stepUuid** | **string**| The UUID of the step. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineVariableForTeam**
> PipelineVariable GetPipelineVariableForTeam(ctx, username, variableUuid)
Get a variable for a team

Retrieve a team level variable. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **variableUuid** | **string**| The UUID of the variable to retrieve. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineVariableForUser**
> PipelineVariable GetPipelineVariableForUser(ctx, selectedUser, variableUuid)
Get a variable for a user

Retrieve a user level variable. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
  **variableUuid** | **string**| The UUID of the variable to retrieve. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineVariableForWorkspace**
> PipelineVariable GetPipelineVariableForWorkspace(ctx, workspace, variableUuid)
Get variable for a workspace

Retrieve a workspace level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **variableUuid** | **string**| The UUID of the variable to retrieve. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineVariablesForTeam**
> PaginatedPipelineVariables GetPipelineVariablesForTeam(ctx, username)
List variables for an account

Find account level variables. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 

### Return type

[**PaginatedPipelineVariables**](paginated_pipeline_variables.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineVariablesForUser**
> PaginatedPipelineVariables GetPipelineVariablesForUser(ctx, selectedUser)
List variables for a user

Find user level variables. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 

### Return type

[**PaginatedPipelineVariables**](paginated_pipeline_variables.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineVariablesForWorkspace**
> PaginatedPipelineVariables GetPipelineVariablesForWorkspace(ctx, workspace)
List variables for a workspace

Find workspace level variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 

### Return type

[**PaginatedPipelineVariables**](paginated_pipeline_variables.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelinesForRepository**
> PaginatedPipelines GetPipelinesForRepository(ctx, workspace, repoSlug)
List pipelines

Find pipelines

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 

### Return type

[**PaginatedPipelines**](paginated_pipelines.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineCacheContentURI**
> PipelineCacheContentUri GetRepositoryPipelineCacheContentURI(ctx, workspace, repoSlug, cacheUuid)
Get cache content URI

Retrieve the URI of the content of the specified cache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 
  **cacheUuid** | **string**| The UUID of the cache. | 

### Return type

[**PipelineCacheContentUri**](pipeline_cache_content_uri.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineCaches**
> PaginatedPipelineCaches GetRepositoryPipelineCaches(ctx, workspace, repoSlug)
List caches

Retrieve the repository pipelines caches.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 

### Return type

[**PaginatedPipelineCaches**](paginated_pipeline_caches.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineConfig**
> PipelinesConfig GetRepositoryPipelineConfig(ctx, workspace, repoSlug)
Get configuration

Retrieve the repository pipelines configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The account. | 
  **repoSlug** | **string**| The repository. | 

### Return type

[**PipelinesConfig**](pipelines_config.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineKnownHost**
> PipelineKnownHost GetRepositoryPipelineKnownHost(ctx, workspace, repoSlug, knownHostUuid)
Get a known host

Retrieve a repository level known host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **knownHostUuid** | **string**| The UUID of the known host to retrieve. | 

### Return type

[**PipelineKnownHost**](pipeline_known_host.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineKnownHosts**
> PaginatedPipelineKnownHosts GetRepositoryPipelineKnownHosts(ctx, workspace, repoSlug)
List known hosts

Find repository level known hosts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 

### Return type

[**PaginatedPipelineKnownHosts**](paginated_pipeline_known_hosts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineSchedule**
> PipelineSchedule GetRepositoryPipelineSchedule(ctx, workspace, repoSlug, scheduleUuid)
Get a schedule

Retrieve a schedule by its UUID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **scheduleUuid** | **string**| The uuid of the schedule. | 

### Return type

[**PipelineSchedule**](pipeline_schedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineScheduleExecutions**
> PaginatedPipelineScheduleExecutions GetRepositoryPipelineScheduleExecutions(ctx, workspace, repoSlug, scheduleUuid)
List executions of a schedule

Retrieve the executions of a given schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **scheduleUuid** | **string**| The uuid of the schedule. | 

### Return type

[**PaginatedPipelineScheduleExecutions**](paginated_pipeline_schedule_executions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineSchedules**
> PaginatedPipelineSchedules GetRepositoryPipelineSchedules(ctx, workspace, repoSlug)
List schedules

Retrieve the configured schedules for the given repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 

### Return type

[**PaginatedPipelineSchedules**](paginated_pipeline_schedules.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineSshKeyPair**
> PipelineSshKeyPair GetRepositoryPipelineSshKeyPair(ctx, workspace, repoSlug)
Get SSH key pair

Retrieve the repository SSH key pair excluding the SSH private key. The private key is a write only field and will never be exposed in the logs or the REST API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 

### Return type

[**PipelineSshKeyPair**](pipeline_ssh_key_pair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineVariable**
> PipelineVariable GetRepositoryPipelineVariable(ctx, workspace, repoSlug, variableUuid)
Get a variable for a repository

Retrieve a repository level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **variableUuid** | **string**| The UUID of the variable to retrieve. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryPipelineVariables**
> PaginatedPipelineVariables GetRepositoryPipelineVariables(ctx, workspace, repoSlug)
List variables for a repository

Find repository level variables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 

### Return type

[**PaginatedPipelineVariables**](paginated_pipeline_variables.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopPipeline**
> StopPipeline(ctx, workspace, repoSlug, pipelineUuid)
Stop a pipeline

Signal the stop of a pipeline and all of its steps that not have completed yet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **pipelineUuid** | **string**| The UUID of the pipeline. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDeploymentVariable**
> DeploymentVariable UpdateDeploymentVariable(ctx, workspace, repoSlug, environmentUuid, variableUuid, body)
Update a variable for an environment

Update a deployment environment level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **environmentUuid** | **string**| The environment. | 
  **variableUuid** | **string**| The UUID of the variable to update. | 
  **body** | [**DeploymentVariable**](DeploymentVariable.md)| The updated deployment variable. | 

### Return type

[**DeploymentVariable**](deployment_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipelineVariableForTeam**
> PipelineVariable UpdatePipelineVariableForTeam(ctx, username, variableUuid, body)
Update a variable for a team

Update a team level variable. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account. | 
  **variableUuid** | **string**| The UUID of the variable. | 
  **body** | [**PipelineVariable**](PipelineVariable.md)| The updated variable. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipelineVariableForUser**
> PipelineVariable UpdatePipelineVariableForUser(ctx, selectedUser, variableUuid, body)
Update a variable for a user

Update a user level variable. This endpoint has been deprecated, and you should use the new workspaces endpoint. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
  **variableUuid** | **string**| The UUID of the variable. | 
  **body** | [**PipelineVariable**](PipelineVariable.md)| The updated variable. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePipelineVariableForWorkspace**
> PipelineVariable UpdatePipelineVariableForWorkspace(ctx, workspace, variableUuid, body)
Update variable for a workspace

Update a workspace level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **variableUuid** | **string**| The UUID of the variable. | 
  **body** | [**PipelineVariable**](PipelineVariable.md)| The updated variable. | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryBuildNumber**
> PipelineBuildNumber UpdateRepositoryBuildNumber(ctx, workspace, repoSlug, body)
Update the next build number

Update the next build number that should be assigned to a pipeline. The next build number that will be configured has to be strictly higher than the current latest build number for this repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **body** | [**PipelineBuildNumber**](PipelineBuildNumber.md)| The build number to update. | 

### Return type

[**PipelineBuildNumber**](pipeline_build_number.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryPipelineConfig**
> PipelinesConfig UpdateRepositoryPipelineConfig(ctx, workspace, repoSlug, body)
Update configuration

Update the pipelines configuration for a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **body** | [**PipelinesConfig**](PipelinesConfig.md)| The updated repository pipelines configuration. | 

### Return type

[**PipelinesConfig**](pipelines_config.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryPipelineKeyPair**
> PipelineSshKeyPair UpdateRepositoryPipelineKeyPair(ctx, workspace, repoSlug, body)
Update SSH key pair

Create or update the repository SSH key pair. The private key will be set as a default SSH identity in your build container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **body** | [**PipelineSshKeyPair**](PipelineSshKeyPair.md)| The created or updated SSH key pair. | 

### Return type

[**PipelineSshKeyPair**](pipeline_ssh_key_pair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryPipelineKnownHost**
> PipelineKnownHost UpdateRepositoryPipelineKnownHost(ctx, workspace, repoSlug, knownHostUuid, body)
Update a known host

Update a repository level known host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **knownHostUuid** | **string**| The UUID of the known host to update. | 
  **body** | [**PipelineKnownHost**](PipelineKnownHost.md)| The updated known host. | 

### Return type

[**PipelineKnownHost**](pipeline_known_host.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryPipelineSchedule**
> PipelineSchedule UpdateRepositoryPipelineSchedule(ctx, workspace, repoSlug, scheduleUuid, body)
Update a schedule

Update a schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **scheduleUuid** | **string**| The uuid of the schedule. | 
  **body** | [**PipelineSchedule**](PipelineSchedule.md)| The schedule to update. | 

### Return type

[**PipelineSchedule**](pipeline_schedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryPipelineVariable**
> PipelineVariable UpdateRepositoryPipelineVariable(ctx, workspace, repoSlug, variableUuid, body)
Update a variable for a repository

Update a repository level variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **variableUuid** | **string**| The UUID of the variable to update. | 
  **body** | [**PipelineVariable**](PipelineVariable.md)| The updated variable | 

### Return type

[**PipelineVariable**](pipeline_variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

