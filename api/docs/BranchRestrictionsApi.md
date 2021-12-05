# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsGet**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsGet) | **Get** /repositories/{workspace}/{repo_slug}/branch-restrictions | List branch restrictions
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete) | **Delete** /repositories/{workspace}/{repo_slug}/branch-restrictions/{id} | Delete a branch restriction rule
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet) | **Get** /repositories/{workspace}/{repo_slug}/branch-restrictions/{id} | Get a branch restriction rule
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut) | **Put** /repositories/{workspace}/{repo_slug}/branch-restrictions/{id} | Update a branch restriction rule
[**RepositoriesWorkspaceRepoSlugBranchRestrictionsPost**](BranchRestrictionsApi.md#RepositoriesWorkspaceRepoSlugBranchRestrictionsPost) | **Post** /repositories/{workspace}/{repo_slug}/branch-restrictions | Create a branch restriction rule

# **RepositoriesWorkspaceRepoSlugBranchRestrictionsGet**
> PaginatedBranchrestrictions RepositoriesWorkspaceRepoSlugBranchRestrictionsGet(ctx, repoSlug, workspace, optional)
List branch restrictions

Returns a paginated list of all branch restrictions on the repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***BranchRestrictionsApiRepositoriesWorkspaceRepoSlugBranchRestrictionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BranchRestrictionsApiRepositoriesWorkspaceRepoSlugBranchRestrictionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kind** | **optional.String**| Branch restrictions of this type | 
 **pattern** | **optional.String**| Branch restrictions applied to branches of this pattern | 

### Return type

[**PaginatedBranchrestrictions**](paginated_branchrestrictions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete**
> RepositoriesWorkspaceRepoSlugBranchRestrictionsIdDelete(ctx, id, repoSlug, workspace)
Delete a branch restriction rule

Deletes an existing branch restriction rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The restriction rule&#x27;s id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet**
> Branchrestriction RepositoriesWorkspaceRepoSlugBranchRestrictionsIdGet(ctx, id, repoSlug, workspace)
Get a branch restriction rule

Returns a specific branch restriction rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The restriction rule&#x27;s id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Branchrestriction**](branchrestriction.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut**
> Branchrestriction RepositoriesWorkspaceRepoSlugBranchRestrictionsIdPut(ctx, body, id, repoSlug, workspace)
Update a branch restriction rule

Updates an existing branch restriction rule.  Fields not present in the request body are ignored.  See [`POST`](/cloud/bitbucket/rest/api-group-branch-restrictions/#api-repositories-workspace-repo-slug-branch-restrictions-post) for details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Branchrestriction**](Branchrestriction.md)| The new version of the existing rule | 
  **id** | **string**| The restriction rule&#x27;s id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Branchrestriction**](branchrestriction.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugBranchRestrictionsPost**
> Branchrestriction RepositoriesWorkspaceRepoSlugBranchRestrictionsPost(ctx, body, repoSlug, workspace)
Create a branch restriction rule

Creates a new branch restriction rule for a repository.  `kind` describes what will be restricted. Allowed values include: `push`, `force`, `delete` and `restrict_merges`.  Different kinds of branch restrictions have different requirements:  * `push` and `restrict_merges` require `users` and `groups` to be   specified. Empty lists are allowed, in which case permission is   denied for everybody.  The restriction applies to all branches that match. There are two ways to match a branch. It is configured in `branch_match_kind`:  1. `glob`: Matches a branch against the `pattern`. A `'*'` in    `pattern` will expand to match zero or more characters, and every    other character matches itself. For example, `'foo*'` will match    `'foo'` and `'foobar'`, but not `'barfoo'`. `'*'` will match all    branches. 2. `branching_model`: Matches a branch against the repository's    branching model. The `branch_type` controls the type of branch    to match. Allowed values include: `production`, `development`,    `bugfix`, `release`, `feature` and `hotfix`.  The combination of `kind` and match must be unique. This means that two `glob` restrictions in a repository cannot have the same `kind` and `pattern`. Additionally, two `branching_model` restrictions in a repository cannot have the same `kind` and `branch_type`.  `users` and `groups` are lists of users and groups that are except from the restriction. They can only be configured in `push` and `restrict_merges` restrictions. The `push` restriction stops a user pushing to matching branches unless that user is in `users` or is a member of a group in `groups`. The `restrict_merges` stops a user merging pull requests to matching branches unless that user is in `users` or is a member of a group in `groups`. Adding new users or groups to an existing restriction should be done via `PUT`.  Note that branch restrictions with overlapping matchers is allowed, but the resulting behavior may be surprising.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Branchrestriction**](Branchrestriction.md)| The new rule | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Branchrestriction**](branchrestriction.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

