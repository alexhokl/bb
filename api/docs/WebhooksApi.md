# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HookEventsGet**](WebhooksApi.md#HookEventsGet) | **Get** /hook_events | Get a webhook resource
[**HookEventsSubjectTypeGet**](WebhooksApi.md#HookEventsSubjectTypeGet) | **Get** /hook_events/{subject_type} | List subscribable webhook types
[**RepositoriesWorkspaceRepoSlugHooksGet**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksGet) | **Get** /repositories/{workspace}/{repo_slug}/hooks | List webhooks for a repository
[**RepositoriesWorkspaceRepoSlugHooksPost**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksPost) | **Post** /repositories/{workspace}/{repo_slug}/hooks | Create a webhook for a repository
[**RepositoriesWorkspaceRepoSlugHooksUidDelete**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksUidDelete) | **Delete** /repositories/{workspace}/{repo_slug}/hooks/{uid} | Delete a webhook for a repository
[**RepositoriesWorkspaceRepoSlugHooksUidGet**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksUidGet) | **Get** /repositories/{workspace}/{repo_slug}/hooks/{uid} | Get a webhook for a repository
[**RepositoriesWorkspaceRepoSlugHooksUidPut**](WebhooksApi.md#RepositoriesWorkspaceRepoSlugHooksUidPut) | **Put** /repositories/{workspace}/{repo_slug}/hooks/{uid} | Update a webhook for a repository
[**WorkspacesWorkspaceHooksGet**](WebhooksApi.md#WorkspacesWorkspaceHooksGet) | **Get** /workspaces/{workspace}/hooks | List webhooks for a workspace
[**WorkspacesWorkspaceHooksPost**](WebhooksApi.md#WorkspacesWorkspaceHooksPost) | **Post** /workspaces/{workspace}/hooks | Create a webhook for a workspace
[**WorkspacesWorkspaceHooksUidDelete**](WebhooksApi.md#WorkspacesWorkspaceHooksUidDelete) | **Delete** /workspaces/{workspace}/hooks/{uid} | Delete a webhook for a workspace
[**WorkspacesWorkspaceHooksUidGet**](WebhooksApi.md#WorkspacesWorkspaceHooksUidGet) | **Get** /workspaces/{workspace}/hooks/{uid} | Get a webhook for a workspace
[**WorkspacesWorkspaceHooksUidPut**](WebhooksApi.md#WorkspacesWorkspaceHooksUidPut) | **Put** /workspaces/{workspace}/hooks/{uid} | Update a webhook for a workspace

# **HookEventsGet**
> SubjectTypes HookEventsGet(ctx, )
Get a webhook resource

Returns the webhook resource or subject types on which webhooks can be registered.  Each resource/subject type contains an `events` link that returns the paginated list of specific events each individual subject type can emit.  This endpoint is publicly accessible and does not require authentication or scopes.  Example:  ``` $ curl https://api.bitbucket.org/2.0/hook_events  {     \"repository\": {         \"links\": {             \"events\": {                 \"href\": \"https://api.bitbucket.org/2.0/hook_events/repository\"             }         }     },     \"team\": {         \"links\": {             \"events\": {                 \"href\": \"https://api.bitbucket.org/2.0/hook_events/team\"             }         }     },     \"user\": {         \"links\": {             \"events\": {                 \"href\": \"https://api.bitbucket.org/2.0/hook_events/user\"             }         }     } } ```

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SubjectTypes**](subject_types.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HookEventsSubjectTypeGet**
> PaginatedHookEvents HookEventsSubjectTypeGet(ctx, subjectType)
List subscribable webhook types

Returns a paginated list of all valid webhook events for the specified entity. **The team and user webhooks are deprecated, and you should use workspace instead. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**  This is public data that does not require any scopes or authentication.  Example:  NOTE: The following example is a truncated response object for the `workspace` `subject_type`. We return the same structure for the other `subject_type` objects.  ``` $ curl https://api.bitbucket.org/2.0/hook_events/workspace {     \"page\": 1,     \"pagelen\": 30,     \"size\": 21,     \"values\": [         {             \"category\": \"Repository\",             \"description\": \"Whenever a repository push occurs\",             \"event\": \"repo:push\",             \"label\": \"Push\"         },         {             \"category\": \"Repository\",             \"description\": \"Whenever a repository fork occurs\",             \"event\": \"repo:fork\",             \"label\": \"Fork\"         },         {             \"category\": \"Repository\",             \"description\": \"Whenever a repository import occurs\",             \"event\": \"repo:imported\",             \"label\": \"Import\"         },         ...         {             \"category\":\"Pull Request\",             \"label\":\"Approved\",             \"description\":\"When someone has approved a pull request\",             \"event\":\"pullrequest:approved\"         },     ] } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subjectType** | **string**| A resource or subject type. | 

### Return type

[**PaginatedHookEvents**](paginated_hook_events.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugHooksGet**
> PaginatedWebhookSubscriptions RepositoriesWorkspaceRepoSlugHooksGet(ctx, repoSlug, workspace)
List webhooks for a repository

Returns a paginated list of webhooks installed on this repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedWebhookSubscriptions**](paginated_webhook_subscriptions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugHooksPost**
> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksPost(ctx, repoSlug, workspace)
Create a webhook for a repository

Creates a new webhook on the specified repository.  Example:  ``` $ curl -X POST -u credentials -H 'Content-Type: application/json'   https://api.bitbucket.org/2.0/repositories/my-workspace/my-repo-slug/hooks   -d '     {       \"description\": \"Webhook Description\",       \"url\": \"https://example.com/\",       \"active\": true,       \"events\": [         \"repo:push\",         \"issue:created\",         \"issue:updated\"       ]     }' ```  Note that this call requires the webhook scope, as well as any scope that applies to the events that the webhook subscribes to. In the example above that means: `webhook`, `repository` and `issue`.  Also note that the `url` must properly resolve and cannot be an internal, non-routed address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugHooksUidDelete**
> RepositoriesWorkspaceRepoSlugHooksUidDelete(ctx, repoSlug, uid, workspace)
Delete a webhook for a repository

Deletes the specified webhook subscription from the given repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **uid** | **string**| Installed webhook&#x27;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugHooksUidGet**
> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksUidGet(ctx, repoSlug, uid, workspace)
Get a webhook for a repository

Returns the webhook with the specified id installed on the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **uid** | **string**| Installed webhook&#x27;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugHooksUidPut**
> WebhookSubscription RepositoriesWorkspaceRepoSlugHooksUidPut(ctx, repoSlug, uid, workspace)
Update a webhook for a repository

Updates the specified webhook subscription.  The following properties can be mutated:  * `description` * `url` * `active` * `events`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **uid** | **string**| Installed webhook&#x27;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksGet**
> PaginatedWebhookSubscriptions WorkspacesWorkspaceHooksGet(ctx, workspace)
List webhooks for a workspace

Returns a paginated list of webhooks installed on this workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedWebhookSubscriptions**](paginated_webhook_subscriptions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksPost**
> WebhookSubscription WorkspacesWorkspaceHooksPost(ctx, workspace)
Create a webhook for a workspace

Creates a new webhook on the specified workspace.  Workspace webhooks are fired for events from all repositories contained by that workspace.  Note that only owners can install webhooks on workspaces.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksUidDelete**
> WorkspacesWorkspaceHooksUidDelete(ctx, uid, workspace)
Delete a webhook for a workspace

Deletes the specified webhook subscription from the given workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Installed webhook&#x27;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksUidGet**
> WebhookSubscription WorkspacesWorkspaceHooksUidGet(ctx, uid, workspace)
Get a webhook for a workspace

Returns the webhook with the specified id installed on the given workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Installed webhook&#x27;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksUidPut**
> WebhookSubscription WorkspacesWorkspaceHooksUidPut(ctx, uid, workspace)
Update a webhook for a workspace

Updates the specified webhook subscription.  The following properties can be mutated:  * `description` * `url` * `active` * `events`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**| Installed webhook&#x27;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

