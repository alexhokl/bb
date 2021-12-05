# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsWorkspaceRepositoriesGet**](UsersApi.md#TeamsWorkspaceRepositoriesGet) | **Get** /teams/{workspace}/repositories | List workspace repositories
[**UserEmailsEmailGet**](UsersApi.md#UserEmailsEmailGet) | **Get** /user/emails/{email} | Get an email address for current user
[**UserEmailsGet**](UsersApi.md#UserEmailsGet) | **Get** /user/emails | List email addresses for current user
[**UserGet**](UsersApi.md#UserGet) | **Get** /user | Get current user
[**UsersSelectedUserGet**](UsersApi.md#UsersSelectedUserGet) | **Get** /users/{selected_user} | Get a user
[**UsersUsernameMembersGet**](UsersApi.md#UsersUsernameMembersGet) | **Get** /users/{username}/members | List team users
[**UsersWorkspaceRepositoriesGet**](UsersApi.md#UsersWorkspaceRepositoriesGet) | **Get** /users/{workspace}/repositories | List workspace repositories

# **TeamsWorkspaceRepositoriesGet**
> ModelError TeamsWorkspaceRepositoriesGet(ctx, workspace)
List workspace repositories

All repositories in the given workspace. This includes any private repositories the calling user has access to.  **This endpoint has been deprecated and will stop functioning soon. You should use the [repository list](/cloud/bitbucket/rest/api-group-repositories/#api-repositories-workspace-get) endpoint instead. For more information, see the [deprecation announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserEmailsEmailGet**
> ModelError UserEmailsEmailGet(ctx, email)
Get an email address for current user

Returns details about a specific one of the authenticated user's email addresses.  Details describe whether the address has been confirmed by the user and whether it is the user's primary address or not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| Email address of the user. | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserEmailsGet**
> ModelError UserEmailsGet(ctx, )
List email addresses for current user

Returns all the authenticated user's email addresses. Both confirmed and unconfirmed.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGet**
> User UserGet(ctx, )
Get current user

Returns the currently logged in user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](user.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserGet**
> User UsersSelectedUserGet(ctx, selectedUser)
Get a user

Gets the public information associated with a user account.  If the user's profile is private, `location`, `website` and `created_on` elements are omitted.  Note that the user object returned by this operation is changing significantly, due to privacy changes. See the [announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#changes-to-bitbucket-user-objects) for details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Return type

[**User**](user.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsernameMembersGet**
> User UsersUsernameMembersGet(ctx, username)
List team users

**This endpoint has been deprecated and will stop functioning soon. You should use the [workspaces](/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get) endpoint instead. For more information, see [this post](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Return type

[**User**](user.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersWorkspaceRepositoriesGet**
> ModelError UsersWorkspaceRepositoriesGet(ctx, workspace)
List workspace repositories

All repositories in the given workspace. This includes any private repositories the calling user has access to.  **This endpoint has been deprecated and will stop functioning soon. You should use the [repository list](/cloud/bitbucket/rest/api-group-repositories/#api-repositories-workspace-get) endpoint instead. For more information, see the [deprecation announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

