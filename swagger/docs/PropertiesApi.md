# \PropertiesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCommitHostedPropertyValue**](PropertiesApi.md#DeleteCommitHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | Delete a commit application property
[**DeletePullRequestHostedPropertyValue**](PropertiesApi.md#DeletePullRequestHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | Delete a pull request application property
[**DeleteRepositoryHostedPropertyValue**](PropertiesApi.md#DeleteRepositoryHostedPropertyValue) | **Delete** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | Delete a repository application property
[**DeleteUserHostedPropertyValue**](PropertiesApi.md#DeleteUserHostedPropertyValue) | **Delete** /users/{selected_user}/properties/{app_key}/{property_name} | Delete a user application property
[**GetCommitHostedPropertyValue**](PropertiesApi.md#GetCommitHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | Get a commit application property
[**GetPullRequestHostedPropertyValue**](PropertiesApi.md#GetPullRequestHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | Get a pull request application property
[**GetRepositoryHostedPropertyValue**](PropertiesApi.md#GetRepositoryHostedPropertyValue) | **Get** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | Get a repository application property
[**RetrieveUserHostedPropertyValue**](PropertiesApi.md#RetrieveUserHostedPropertyValue) | **Get** /users/{selected_user}/properties/{app_key}/{property_name} | Get a user application property
[**UpdateCommitHostedPropertyValue**](PropertiesApi.md#UpdateCommitHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/properties/{app_key}/{property_name} | Update a commit application property
[**UpdatePullRequestHostedPropertyValue**](PropertiesApi.md#UpdatePullRequestHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/pullrequests/{pullrequest_id}/properties/{app_key}/{property_name} | Update a pull request application property
[**UpdateRepositoryHostedPropertyValue**](PropertiesApi.md#UpdateRepositoryHostedPropertyValue) | **Put** /repositories/{workspace}/{repo_slug}/properties/{app_key}/{property_name} | Update a repository application property
[**UpdateUserHostedPropertyValue**](PropertiesApi.md#UpdateUserHostedPropertyValue) | **Put** /users/{selected_user}/properties/{app_key}/{property_name} | Update a user application property


# **DeleteCommitHostedPropertyValue**
> DeleteCommitHostedPropertyValue(ctx, workspace, repoSlug, commit, appKey, propertyName)
Delete a commit application property

Delete an [application property](/cloud/bitbucket/application-properties/) value stored against a commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePullRequestHostedPropertyValue**
> DeletePullRequestHostedPropertyValue(ctx, workspace, repoSlug, pullrequestId, appKey, propertyName)
Delete a pull request application property

Delete an [application property](/cloud/bitbucket/application-properties/) value stored against a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **pullrequestId** | **string**| The pull request ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryHostedPropertyValue**
> DeleteRepositoryHostedPropertyValue(ctx, workspace, repoSlug, appKey, propertyName)
Delete a repository application property

Delete an [application property](/cloud/bitbucket/application-properties/) value stored against a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserHostedPropertyValue**
> DeleteUserHostedPropertyValue(ctx, selectedUser, appKey, propertyName)
Delete a user application property

Delete an [application property](/cloud/bitbucket/application-properties/) value stored against a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitHostedPropertyValue**
> ApplicationProperty GetCommitHostedPropertyValue(ctx, workspace, repoSlug, commit, appKey, propertyName)
Get a commit application property

Retrieve an [application property](/cloud/bitbucket/application-properties/) value stored against a commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

[**ApplicationProperty**](application_property.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestHostedPropertyValue**
> ApplicationProperty GetPullRequestHostedPropertyValue(ctx, workspace, repoSlug, pullrequestId, appKey, propertyName)
Get a pull request application property

Retrieve an [application property](/cloud/bitbucket/application-properties/) value stored against a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **pullrequestId** | **string**| The pull request ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

[**ApplicationProperty**](application_property.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepositoryHostedPropertyValue**
> ApplicationProperty GetRepositoryHostedPropertyValue(ctx, workspace, repoSlug, appKey, propertyName)
Get a repository application property

Retrieve an [application property](/cloud/bitbucket/application-properties/) value stored against a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

[**ApplicationProperty**](application_property.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveUserHostedPropertyValue**
> ApplicationProperty RetrieveUserHostedPropertyValue(ctx, selectedUser, appKey, propertyName)
Get a user application property

Retrieve an [application property](/cloud/bitbucket/application-properties/) value stored against a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 

### Return type

[**ApplicationProperty**](application_property.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCommitHostedPropertyValue**
> UpdateCommitHostedPropertyValue(ctx, workspace, repoSlug, commit, appKey, propertyName, body)
Update a commit application property

Update an [application property](/cloud/bitbucket/application-properties/) value stored against a commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 
  **body** | [**ApplicationProperty**](ApplicationProperty.md)| The application property to create or update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePullRequestHostedPropertyValue**
> UpdatePullRequestHostedPropertyValue(ctx, workspace, repoSlug, pullrequestId, appKey, propertyName, body)
Update a pull request application property

Update an [application property](/cloud/bitbucket/application-properties/) value stored against a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **pullrequestId** | **string**| The pull request ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 
  **body** | [**ApplicationProperty**](ApplicationProperty.md)| The application property to create or update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryHostedPropertyValue**
> UpdateRepositoryHostedPropertyValue(ctx, workspace, repoSlug, appKey, propertyName, body)
Update a repository application property

Update an [application property](/cloud/bitbucket/application-properties/) value stored against a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The repository container; either the workspace slug or the UUID in curly braces. | 
  **repoSlug** | **string**| The repository. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 
  **body** | [**ApplicationProperty**](ApplicationProperty.md)| The application property to create or update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserHostedPropertyValue**
> UpdateUserHostedPropertyValue(ctx, selectedUser, appKey, propertyName, body)
Update a user application property

Update an [application property](/cloud/bitbucket/application-properties/) value stored against a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
  **appKey** | **string**| The key of the Connect app. | 
  **propertyName** | **string**| The name of the property. | 
  **body** | [**ApplicationProperty**](ApplicationProperty.md)| The application property to create or update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

