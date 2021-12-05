# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugBranchingModelGet**](BranchingModelApi.md#RepositoriesWorkspaceRepoSlugBranchingModelGet) | **Get** /repositories/{workspace}/{repo_slug}/branching-model | Get the branching model for a repository
[**RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet**](BranchingModelApi.md#RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet) | **Get** /repositories/{workspace}/{repo_slug}/branching-model/settings | Get the branching model config for a repository
[**RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut**](BranchingModelApi.md#RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut) | **Put** /repositories/{workspace}/{repo_slug}/branching-model/settings | Update the branching model config for a repository

# **RepositoriesWorkspaceRepoSlugBranchingModelGet**
> BranchingModel RepositoriesWorkspaceRepoSlugBranchingModelGet(ctx, repoSlug, workspace)
Get the branching model for a repository

Return the branching model as applied to the repository. This view is read-only. The branching model settings can be changed using the [settings](branching-model/settings#get) API.  The returned object:  1. Always has a `development` property. `development.branch` contains    the actual repository branch object that is considered to be the    `development` branch. `development.branch` will not be present    if it does not exist. 2. Might have a `production` property. `production` will not    be present when `production` is disabled.    `production.branch` contains the actual branch object that is    considered to be the `production` branch. `production.branch` will    not be present if it does not exist. 3. Always has a `branch_types` array which contains all enabled branch    types.  Example body:  ``` {   \"development\": {     \"name\": \"master\",     \"branch\": {       \"type\": \"branch\",       \"name\": \"master\",       \"target\": {         \"hash\": \"16dffcb0de1b22e249db6799532074cf32efe80f\"       }     },     \"use_mainbranch\": true   },   \"production\": {     \"name\": \"production\",     \"branch\": {       \"type\": \"branch\",       \"name\": \"production\",       \"target\": {         \"hash\": \"16dffcb0de1b22e249db6799532074cf32efe80f\"       }     },     \"use_mainbranch\": false   },   \"branch_types\": [     {       \"kind\": \"release\",       \"prefix\": \"release/\"     },     {       \"kind\": \"hotfix\",       \"prefix\": \"hotfix/\"     },     {       \"kind\": \"feature\",       \"prefix\": \"feature/\"     },     {       \"kind\": \"bugfix\",       \"prefix\": \"bugfix/\"     }   ],   \"type\": \"branching_model\",   \"links\": {     \"self\": {       \"href\": \"https://api.bitbucket.org/.../branching-model\"     }   } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**BranchingModel**](branching_model.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet**
> BranchingModelSettings RepositoriesWorkspaceRepoSlugBranchingModelSettingsGet(ctx, repoSlug, workspace)
Get the branching model config for a repository

Return the branching model configuration for a repository. The returned object:  1. Always has a `development` property for the development branch. 2. Always a `production` property for the production branch. The    production branch can be disabled. 3. The `branch_types` contains all the branch types.  This is the raw configuration for the branching model. A client wishing to see the branching model with its actual current branches may find the [active model API](/cloud/bitbucket/rest/api-group-branching-model/#api-repositories-workspace-repo-slug-branching-model-get) more useful.  Example body:  ``` {   \"development\": {     \"is_valid\": true,     \"name\": null,     \"use_mainbranch\": true   },   \"production\": {     \"is_valid\": true,     \"name\": \"production\",     \"use_mainbranch\": false,     \"enabled\": false   },   \"branch_types\": [     {       \"kind\": \"release\",       \"enabled\": true,       \"prefix\": \"release/\"     },     {       \"kind\": \"hotfix\",       \"enabled\": true,       \"prefix\": \"hotfix/\"     },     {       \"kind\": \"feature\",       \"enabled\": true,       \"prefix\": \"feature/\"     },     {       \"kind\": \"bugfix\",       \"enabled\": false,       \"prefix\": \"bugfix/\"     }   ],   \"type\": \"branching_model_settings\",   \"links\": {     \"self\": {       \"href\": \"https://api.bitbucket.org/.../branching-model/settings\"     }   } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**BranchingModelSettings**](branching_model_settings.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut**
> BranchingModelSettings RepositoriesWorkspaceRepoSlugBranchingModelSettingsPut(ctx, repoSlug, workspace)
Update the branching model config for a repository

Update the branching model configuration for a repository.  The `development` branch can be configured to a specific branch or to track the main branch. When set to a specific branch it must currently exist. Only the passed properties will be updated. The properties not passed will be left unchanged. A request without a `development` property will leave the development branch unchanged.  It is possible for the `development` branch to be invalid. This happens when it points at a specific branch that has been deleted. This is indicated in the `is_valid` field for the branch. It is not possible to update the settings for `development` if that would leave the branch in an invalid state. Such a request will be rejected.  The `production` branch can be a specific branch, the main branch or disabled. When set to a specific branch it must currently exist. The `enabled` property can be used to enable (`true`) or disable (`false`) it. Only the passed properties will be updated. The properties not passed will be left unchanged. A request without a `production` property will leave the production branch unchanged.  It is possible for the `production` branch to be invalid. This happens when it points at a specific branch that has been deleted. This is indicated in the `is_valid` field for the branch. A request that would leave `production` enabled and invalid will be rejected. It is possible to update `production` and make it invalid if it would also be left disabled.  The `branch_types` property contains the branch types to be updated. Only the branch types passed will be updated. All updates will be rejected if it would leave the branching model in an invalid state. For branch types this means that:  1. The prefixes for all enabled branch types are valid. For example,    it is not possible to use '*' inside a Git prefix. 2. A prefix of an enabled branch type must not be a prefix of another    enabled branch type. This is to ensure that a branch can be easily    classified by its prefix unambiguously.  It is possible to store an invalid prefix if that branch type would be left disabled. Only the passed properties will be updated. The properties not passed will be left unchanged. Each branch type must have a `kind` property to identify it.  Example Body:  ```     {       \"development\": {         \"use_mainbranch\": true       },       \"production\": {         \"enabled\": true,         \"use_mainbranch\": false,         \"name\": \"production\"       },       \"branch_types\": [         {           \"kind\": \"bugfix\",           \"enabled\": true,           \"prefix\": \"bugfix/\"         },         {           \"kind\": \"feature\",           \"enabled\": true,           \"prefix\": \"feature/\"         },         {           \"kind\": \"hotfix\",           \"prefix\": \"hotfix/\"         },         {           \"kind\": \"release\",           \"enabled\": false,         }       ]     } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**BranchingModelSettings**](branching_model_settings.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

