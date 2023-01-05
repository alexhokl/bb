# \UsersApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserEmailsEmailGet**](UsersApi.md#UserEmailsEmailGet) | **Get** /user/emails/{email} | Get an email address for current user
[**UserEmailsGet**](UsersApi.md#UserEmailsGet) | **Get** /user/emails | List email addresses for current user
[**UserGet**](UsersApi.md#UserGet) | **Get** /user | Get current user
[**UsersSelectedUserGet**](UsersApi.md#UsersSelectedUserGet) | **Get** /users/{selected_user} | Get a user


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

 - **Content-Type**: application/json
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGet**
> Account UserGet(ctx, )
Get current user

Returns the currently logged in user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Account**](account.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSelectedUserGet**
> Account UsersSelectedUserGet(ctx, selectedUser)
Get a user

Gets the public information associated with a user account.  If the user's profile is private, `location`, `website` and `created_on` elements are omitted.  Note that the user object returned by this operation is changing significantly, due to privacy changes. See the [announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#changes-to-bitbucket-user-objects) for details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| This can either be the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, OR an Atlassian Account ID.  | 

### Return type

[**Account**](account.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

