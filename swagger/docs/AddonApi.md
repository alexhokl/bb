# \AddonApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddonDelete**](AddonApi.md#AddonDelete) | **Delete** /addon | Delete an app
[**AddonLinkersGet**](AddonApi.md#AddonLinkersGet) | **Get** /addon/linkers | List linkers for an app
[**AddonLinkersLinkerKeyGet**](AddonApi.md#AddonLinkersLinkerKeyGet) | **Get** /addon/linkers/{linker_key} | Get a linker for an app
[**AddonLinkersLinkerKeyValuesDelete**](AddonApi.md#AddonLinkersLinkerKeyValuesDelete) | **Delete** /addon/linkers/{linker_key}/values | Delete all linker values
[**AddonLinkersLinkerKeyValuesGet**](AddonApi.md#AddonLinkersLinkerKeyValuesGet) | **Get** /addon/linkers/{linker_key}/values | List linker values for a linker
[**AddonLinkersLinkerKeyValuesPost**](AddonApi.md#AddonLinkersLinkerKeyValuesPost) | **Post** /addon/linkers/{linker_key}/values | Create a linker value
[**AddonLinkersLinkerKeyValuesPut**](AddonApi.md#AddonLinkersLinkerKeyValuesPut) | **Put** /addon/linkers/{linker_key}/values | Update a linker value
[**AddonLinkersLinkerKeyValuesValueIdDelete**](AddonApi.md#AddonLinkersLinkerKeyValuesValueIdDelete) | **Delete** /addon/linkers/{linker_key}/values/{value_id} | Delete a linker value
[**AddonLinkersLinkerKeyValuesValueIdGet**](AddonApi.md#AddonLinkersLinkerKeyValuesValueIdGet) | **Get** /addon/linkers/{linker_key}/values/{value_id} | Get a linker value
[**AddonPut**](AddonApi.md#AddonPut) | **Put** /addon | Update an installed app


# **AddonDelete**
> AddonDelete(ctx, )
Delete an app

Deletes the application for the user.  This endpoint is intended to be used by Bitbucket Connect apps and only supports JWT authentication -- that is how Bitbucket identifies the particular installation of the app. Developers with applications registered in the \"Develop Apps\" section of Bitbucket Marketplace need not use this endpoint as updates for those applications can be sent out via the UI of that section.  ``` $ curl -X DELETE https://api.bitbucket.org/2.0/addon \\   -H \"Authorization: JWT <JWT Token>\" ```

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonLinkersGet**
> AddonLinkersGet(ctx, )
List linkers for an app

Gets a list of all [linkers](/cloud/bitbucket/modules/linker/) for the authenticated application.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonLinkersLinkerKeyGet**
> AddonLinkersLinkerKeyGet(ctx, linkerKey)
Get a linker for an app

Gets a [linker](/cloud/bitbucket/modules/linker/) specified by `linker_key` for the authenticated application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkerKey** | **string**| The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonLinkersLinkerKeyValuesDelete**
> AddonLinkersLinkerKeyValuesDelete(ctx, linkerKey)
Delete all linker values

Delete all [linker](/cloud/bitbucket/modules/linker/) values for the specified linker of the authenticated application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkerKey** | **string**| The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonLinkersLinkerKeyValuesGet**
> AddonLinkersLinkerKeyValuesGet(ctx, linkerKey)
List linker values for a linker

Gets a list of all [linker](/cloud/bitbucket/modules/linker/) values for the specified linker of the authenticated application.  A linker value lets applications supply values to modify its regular expression.  The base regular expression must use a Bitbucket-specific match group `(?K)` which will be translated to `([\\w\\-]+)`. A value must match this pattern.  [Read more about linker values](/cloud/bitbucket/modules/linker/#usingthebitbucketapitosupplyvalues)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkerKey** | **string**| The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonLinkersLinkerKeyValuesPost**
> AddonLinkersLinkerKeyValuesPost(ctx, linkerKey)
Create a linker value

Creates a [linker](/cloud/bitbucket/modules/linker/) value for the specified linker of authenticated application.  A linker value lets applications supply values to modify its regular expression.  The base regular expression must use a Bitbucket-specific match group `(?K)` which will be translated to `([\\w\\-]+)`. A value must match this pattern.  [Read more about linker values](/cloud/bitbucket/modules/linker/#usingthebitbucketapitosupplyvalues)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkerKey** | **string**| The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonLinkersLinkerKeyValuesPut**
> AddonLinkersLinkerKeyValuesPut(ctx, linkerKey)
Update a linker value

Bulk update [linker](/cloud/bitbucket/modules/linker/) values for the specified linker of the authenticated application.  A linker value lets applications supply values to modify its regular expression.  The base regular expression must use a Bitbucket-specific match group `(?K)` which will be translated to `([\\w\\-]+)`. A value must match this pattern.  [Read more about linker values](/cloud/bitbucket/modules/linker/#usingthebitbucketapitosupplyvalues)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkerKey** | **string**| The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonLinkersLinkerKeyValuesValueIdDelete**
> AddonLinkersLinkerKeyValuesValueIdDelete(ctx, linkerKey, valueId)
Delete a linker value

Delete a single [linker](/cloud/bitbucket/modules/linker/) value of the authenticated application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkerKey** | **string**| The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 
  **valueId** | **int32**| The numeric ID of the linker value. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonLinkersLinkerKeyValuesValueIdGet**
> AddonLinkersLinkerKeyValuesValueIdGet(ctx, linkerKey, valueId)
Get a linker value

Get a single [linker](/cloud/bitbucket/modules/linker/) value of the authenticated application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkerKey** | **string**| The unique key of a [linker module](/cloud/bitbucket/modules/linker/) as defined in an application descriptor. | 
  **valueId** | **int32**| The numeric ID of the linker value. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonPut**
> AddonPut(ctx, )
Update an installed app

Updates the application installation for the user.  This endpoint is intended to be used by Bitbucket Connect apps and only supports JWT authentication -- that is how Bitbucket identifies the particular installation of the app. Developers with applications registered in the \"Develop Apps\" section of Bitbucket need not use this endpoint as updates for those applications can be sent out via the UI of that section.  Passing an empty body will update the installation using the existing descriptor URL.  ``` $ curl -X PUT https://api.bitbucket.org/2.0/addon \\   -H \"Authorization: JWT <JWT Token>\" \\   --header \"Content-Type: application/json\" \\   --data '{}' ```  The new `descriptor` for the installation can be also provided in the body directly.  ``` $ curl -X PUT https://api.bitbucket.org/2.0/addon \\   -H \"Authorization: JWT <JWT Token>\" \\   --header \"Content-Type: application/json\" \\   --data '{\"descriptor\": $NEW_DESCRIPTOR}' ```  In both these modes the URL of the descriptor cannot be changed. To change the descriptor location and upgrade an installation the request must be made exclusively with a `descriptor_url`.   ``` $ curl -X PUT https://api.bitbucket.org/2.0/addon \\   -H \"Authorization: JWT <JWT Token>\" \\   --header \"Content-Type: application/json\" \\   --data '{\"descriptor_url\": $NEW_URL}' ```  The `descriptor_url` must exactly match the marketplace registration that Atlassian has for the application. Contact your Atlassian developer advocate to update this registration. Once the registration has been updated you may call this resource for each installation.  Note that the scopes of the application cannot be increased in the new descriptor nor reduced to none.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

