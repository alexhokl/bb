# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build/{key} | Get a build status for a commit
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build/{key} | Update a build status for a commit
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses/build | Create a build status for a commit
[**RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/statuses | List commit statuses for a commit
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet**](CommitStatusesApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/statuses | List commit statuses for a pull request

# **RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet**
> Commitstatus RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyGet(ctx, commit, key, repoSlug, workspace)
Get a build status for a commit

Returns the specified build status for a commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **key** | **string**| The build status&#x27; unique key | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Commitstatus**](commitstatus.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut**
> Commitstatus RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPut(ctx, commit, key, repoSlug, workspace, optional)
Update a build status for a commit

Used to update the current status of a build status object on the specific commit.  This operation can also be used to change other properties of the build status:  * `state` * `name` * `description` * `url` * `refname`  The `key` cannot be changed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **key** | **string**| The build status&#x27; unique key | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildKeyPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of Commitstatus**](Commitstatus.md)| The updated build status object | 

### Return type

[**Commitstatus**](commitstatus.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost**
> Commitstatus RepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPost(ctx, commit, repoSlug, workspace, optional)
Create a build status for a commit

Creates a new build status against the specified commit.  If the specified key already exists, the existing status object will be overwritten.  Example:  ``` curl https://api.bitbucket.org/2.0/repositories/my-workspace/my-repo/commit/e10dae226959c2194f2b07b077c07762d93821cf/statuses/build/           -X POST -u jdoe -H 'Content-Type: application/json'           -d '{     \"key\": \"MY-BUILD\",     \"state\": \"SUCCESSFUL\",     \"description\": \"42 tests passed\",     \"url\": \"https://www.example.org/my-build-result\"   }' ```  When creating a new commit status, you can use a URI template for the URL. Templates are URLs that contain variable names that Bitbucket will evaluate at runtime whenever the URL is displayed anywhere similar to parameter substitution in [Bitbucket Connect](https://developer.atlassian.com/bitbucket/concepts/context-parameters.html). For example, one could use `https://foo.com/builds/{repository.full_name}` which Bitbucket will turn into `https://foo.com/builds/foo/bar` at render time. The context variables available are `repository` and `commit`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesBuildPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of Commitstatus**](Commitstatus.md)| The new commit status object. | 

### Return type

[**Commitstatus**](commitstatus.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet**
> PaginatedCommitstatuses RepositoriesWorkspaceRepoSlugCommitCommitStatusesGet(ctx, commit, repoSlug, workspace, optional)
List commit statuses for a commit

Returns all statuses (e.g. build results) for a specific commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitStatusesApiRepositoriesWorkspaceRepoSlugCommitCommitStatusesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **optional.String**| Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering).  | 
 **sort** | **optional.String**| Field by which the results should be sorted as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). Defaults to &#x60;created_on&#x60;.  | 

### Return type

[**PaginatedCommitstatuses**](paginated_commitstatuses.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet**
> PaginatedCommitstatuses RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet(ctx, pullRequestId, repoSlug, workspace, optional)
List commit statuses for a pull request

Returns all statuses (e.g. build results) for the given pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitStatusesApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **optional.String**| Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering).  | 
 **sort** | **optional.String**| Field by which the results should be sorted as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). Defaults to &#x60;created_on&#x60;.  | 

### Return type

[**PaginatedCommitstatuses**](paginated_commitstatuses.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

