# \WorkspacesApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserPermissionsWorkspacesGet**](WorkspacesApi.md#UserPermissionsWorkspacesGet) | **Get** /user/permissions/workspaces | List workspaces for the current user
[**WorkspacesGet**](WorkspacesApi.md#WorkspacesGet) | **Get** /workspaces | List workspaces for user
[**WorkspacesWorkspaceGet**](WorkspacesApi.md#WorkspacesWorkspaceGet) | **Get** /workspaces/{workspace} | Get a workspace
[**WorkspacesWorkspaceHooksGet**](WorkspacesApi.md#WorkspacesWorkspaceHooksGet) | **Get** /workspaces/{workspace}/hooks | List webhooks for a workspace
[**WorkspacesWorkspaceHooksPost**](WorkspacesApi.md#WorkspacesWorkspaceHooksPost) | **Post** /workspaces/{workspace}/hooks | Create a webhook for a workspace
[**WorkspacesWorkspaceHooksUidDelete**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidDelete) | **Delete** /workspaces/{workspace}/hooks/{uid} | Delete a webhook for a workspace
[**WorkspacesWorkspaceHooksUidGet**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidGet) | **Get** /workspaces/{workspace}/hooks/{uid} | Get a webhook for a workspace
[**WorkspacesWorkspaceHooksUidPut**](WorkspacesApi.md#WorkspacesWorkspaceHooksUidPut) | **Put** /workspaces/{workspace}/hooks/{uid} | Update a webhook for a workspace
[**WorkspacesWorkspaceMembersGet**](WorkspacesApi.md#WorkspacesWorkspaceMembersGet) | **Get** /workspaces/{workspace}/members | List users in a workspace
[**WorkspacesWorkspaceMembersMemberGet**](WorkspacesApi.md#WorkspacesWorkspaceMembersMemberGet) | **Get** /workspaces/{workspace}/members/{member} | Get user membership for a workspace
[**WorkspacesWorkspacePermissionsGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsGet) | **Get** /workspaces/{workspace}/permissions | List user permissions in a workspace
[**WorkspacesWorkspacePermissionsRepositoriesGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsRepositoriesGet) | **Get** /workspaces/{workspace}/permissions/repositories | List all repository permissions for a workspace
[**WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet**](WorkspacesApi.md#WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet) | **Get** /workspaces/{workspace}/permissions/repositories/{repo_slug} | List a repository permissions for a workspace
[**WorkspacesWorkspaceProjectsGet**](WorkspacesApi.md#WorkspacesWorkspaceProjectsGet) | **Get** /workspaces/{workspace}/projects | List projects in a workspace
[**WorkspacesWorkspaceProjectsProjectKeyGet**](WorkspacesApi.md#WorkspacesWorkspaceProjectsProjectKeyGet) | **Get** /workspaces/{workspace}/projects/{project_key} | Get a project for a workspace


# **UserPermissionsWorkspacesGet**
> PaginatedWorkspaceMemberships UserPermissionsWorkspacesGet(ctx, optional)
List workspaces for the current user

Returns an object for each workspace the caller is a member of, and their effective role - the highest level of privilege the caller has. If a user is a member of multiple groups with distinct roles, only the highest level is returned.  Permissions can be:  * `owner` * `collaborator` * `member`  **The `collaborator` role is being removed from the Bitbucket Cloud API. For more information, see the [deprecation announcement](/cloud/bitbucket/deprecation-notice-collaborator-role/).**  Example:  ``` $ curl https://api.bitbucket.org/2.0/user/permissions/workspaces  {   \"pagelen\": 10,   \"page\": 1,   \"size\": 1,   \"values\": [     {       \"type\": \"workspace_membership\",       \"permission\": \"owner\",       \"last_accessed\": \"2019-03-07T12:35:02.900024+00:00\",       \"added_on\": \"2018-10-11T17:42:02.961424+00:00\",       \"user\": {         \"type\": \"user\",         \"uuid\": \"{470c176d-3574-44ea-bb41-89e8638bcca4}\",         \"nickname\": \"evzijst\",         \"display_name\": \"Erik van Zijst\",       },       \"workspace\": {         \"type\": \"workspace\",         \"uuid\": \"{a15fb181-db1f-48f7-b41f-e1eff06929d6}\",         \"slug\": \"bbworkspace1\",         \"name\": \"Atlassian Bitbucket\",       }     }   ] } ```  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by workspace or permission by adding the following query string parameters:  * `q=workspace.slug=\"bbworkspace1\"` or `q=permission=\"owner\"` * `sort=workspace.slug`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkspacesApiUserPermissionsWorkspacesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiUserPermissionsWorkspacesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**|  Query string to narrow down the response. See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for details. | 
 **sort** | **optional.String**|  Name of a response property to sort results. See [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results) for details.  | 

### Return type

[**PaginatedWorkspaceMemberships**](paginated_workspace_memberships.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesGet**
> PaginatedWorkspaces WorkspacesGet(ctx, optional)
List workspaces for user

Returns a list of workspaces accessible by the authenticated user.  Example:  ``` $ curl https://api.bitbucket.org/2.0/workspaces  {   \"pagelen\": 10,   \"page\": 1,   \"size\": 1,   \"values\": [     {         \"uuid\": \"{a15fb181-db1f-48f7-b41f-e1eff06929d6}\",         \"links\": {             \"owners\": {                 \"href\": \"https://api.bitbucket.org/2.0/workspaces/bbworkspace1/members?q=permission%3D%22owner%22\"             },             \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/workspaces/bbworkspace1\"             },             \"repositories\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/bbworkspace1\"             },             \"snippets\": {                 \"href\": \"https://api.bitbucket.org/2.0/snippets/bbworkspace1\"             },             \"html\": {                 \"href\": \"https://bitbucket.org/bbworkspace1/\"             },             \"avatar\": {                 \"href\": \"https://bitbucket.org/workspaces/bbworkspace1/avatar/?ts=1543465801\"             },             \"members\": {                 \"href\": \"https://api.bitbucket.org/2.0/workspaces/bbworkspace1/members\"             },             \"projects\": {                 \"href\": \"https://api.bitbucket.org/2.0/workspaces/bbworkspace1/projects\"             }         },         \"created_on\": \"2018-11-14T19:15:05.058566+00:00\",         \"type\": \"workspace\",         \"slug\": \"bbworkspace1\",         \"is_private\": true,         \"name\": \"Atlassian Bitbucket\"     }   ] } ```  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by workspace or permission by adding the following query string parameters:  * `q=slug=\"bbworkspace1\"` or `q=is_private=true` * `sort=created_on`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.  **The `collaborator` role is being removed from the Bitbucket Cloud API. For more information, see the [deprecation announcement](/cloud/bitbucket/deprecation-notice-collaborator-role/).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkspacesApiWorkspacesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **optional.String**|              Filters the workspaces based on the authenticated user&#39;s role on each workspace.              * **member**: returns a list of all the workspaces which the caller is a member of                 at least one workspace group or repository             * **collaborator**: returns a list of workspaces which the caller has write access                 to at least one repository in the workspace             * **owner**: returns a list of workspaces which the caller has administrator access              | 
 **q** | **optional.String**|  Query string to narrow down the response. See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for details. | 
 **sort** | **optional.String**|  Name of a response property to sort results. See [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results) for details.  | 

### Return type

[**PaginatedWorkspaces**](paginated_workspaces.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceGet**
> Workspace WorkspacesWorkspaceGet(ctx, workspace)
Get a workspace

Returns the requested workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Workspace**](workspace.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceHooksPost**
> WebhookSubscription WorkspacesWorkspaceHooksPost(ctx, workspace)
Create a webhook for a workspace

Creates a new webhook on the specified workspace.  Workspace webhooks are fired for events from all repositories contained by that workspace.  Example:  ``` $ curl -X POST -u credentials -H 'Content-Type: application/json'   https://api.bitbucket.org/2.0/workspaces/my-workspace/hooks   -d '     {       \"description\": \"Webhook Description\",       \"url\": \"https://example.com/\",       \"active\": true,       \"events\": [         \"repo:push\",         \"issue:created\",         \"issue:updated\"       ]     }' ```  This call requires the webhook scope, as well as any scope that applies to the events that the webhook subscribes to. In the example above that means: `webhook`, `repository` and `issue`.  The `url` must properly resolve and cannot be an internal, non-routed address.  Only workspace owners can install webhooks on workspaces.

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

 - **Content-Type**: application/json
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
  **uid** | **string**| Installed webhook&#39;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
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
  **uid** | **string**| Installed webhook&#39;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
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
  **uid** | **string**| Installed webhook&#39;s ID | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WebhookSubscription**](webhook_subscription.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceMembersGet**
> PaginatedWorkspaceMemberships WorkspacesWorkspaceMembersGet(ctx, workspace)
List users in a workspace

Returns all members of the requested workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedWorkspaceMemberships**](paginated_workspace_memberships.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceMembersMemberGet**
> WorkspaceMembership WorkspacesWorkspaceMembersMemberGet(ctx, member, workspace)
Get user membership for a workspace

Returns the workspace membership, which includes a `User` object for the member and a `Workspace` object for the requested workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **member** | **string**| Member&#39;s UUID or Atlassian ID. | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**WorkspaceMembership**](workspace_membership.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspacePermissionsGet**
> PaginatedWorkspaceMemberships WorkspacesWorkspacePermissionsGet(ctx, workspace, optional)
List user permissions in a workspace

Returns the list of members in a workspace and their permission levels. Permission can be: * `owner` * `collaborator` * `member`  **The `collaborator` role is being removed from the Bitbucket Cloud API. For more information, see the [deprecation announcement](/cloud/bitbucket/deprecation-notice-collaborator-role/).**  Example:  ``` $ curl -X https://api.bitbucket.org/2.0/workspaces/bbworkspace1/permissions  {     \"pagelen\": 10,     \"values\": [         {             \"permission\": \"owner\",             \"type\": \"workspace_membership\",             \"user\": {                 \"type\": \"user\",                 \"uuid\": \"{470c176d-3574-44ea-bb41-89e8638bcca4}\",                 \"display_name\": \"Erik van Zijst\",             },             \"workspace\": {                 \"type\": \"workspace\",                 \"uuid\": \"{a15fb181-db1f-48f7-b41f-e1eff06929d6}\",                 \"slug\": \"bbworkspace1\",                 \"name\": \"Atlassian Bitbucket\",             }         },         {             \"permission\": \"member\",             \"type\": \"workspace_membership\",             \"user\": {                 \"type\": \"user\",                 \"nickname\": \"seanaty\",                 \"display_name\": \"Sean Conaty\",                 \"uuid\": \"{504c3b62-8120-4f0c-a7bc-87800b9d6f70}\"             },             \"workspace\": {                 \"type\": \"workspace\",                 \"uuid\": \"{a15fb181-db1f-48f7-b41f-e1eff06929d6}\",                 \"slug\": \"bbworkspace1\",                 \"name\": \"Atlassian Bitbucket\",             }         }     ],     \"page\": 1,     \"size\": 2 } ```  Results may be further [filtered](/cloud/bitbucket/rest/intro/#filtering) by permission by adding the following query string parameters:  * `q=permission=\"owner\"`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspacePermissionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspacePermissionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 

### Return type

[**PaginatedWorkspaceMemberships**](paginated_workspace_memberships.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspacePermissionsRepositoriesGet**
> PaginatedRepositoryPermissions WorkspacesWorkspacePermissionsRepositoriesGet(ctx, workspace, optional)
List all repository permissions for a workspace

Returns an object for each repository permission for all of a workspace's repositories.  Permissions returned are effective permissions: the highest level of permission the user has. This does not distinguish between direct and indirect (group) privileges.  Only users with admin permission for the team may access this resource.  Permissions can be:  * `admin` * `write` * `read`  Example:  ``` $ curl https://api.bitbucket.org/2.0/workspaces/atlassian_tutorial/permissions/repositories  {   \"pagelen\": 10,   \"values\": [     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Erik van Zijst\",         \"uuid\": \"{d301aafa-d676-4ee0-88be-962be7417567}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"geordi\",         \"full_name\": \"atlassian_tutorial/geordi\",         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"       },       \"permission\": \"admin\"     },     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Sean Conaty\",         \"uuid\": \"{504c3b62-8120-4f0c-a7bc-87800b9d6f70}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"geordi\",         \"full_name\": \"atlassian_tutorial/geordi\",         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"       },       \"permission\": \"write\"     },     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Jeff Zeng\",         \"uuid\": \"{47f92a9a-c3a3-4d0b-bc4e-782a969c5c72}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"whee\",         \"full_name\": \"atlassian_tutorial/whee\",         \"uuid\": \"{30ba25e9-51ff-4555-8dd0-fc7ee2fa0895}\"       },       \"permission\": \"admin\"     }   ],   \"page\": 1,   \"size\": 3 } ```  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by repository, user, or permission by adding the following query string parameters:  * `q=repository.name=\"geordi\"` or `q=permission>\"read\"` * `sort=user.display_name`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspacePermissionsRepositoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspacePermissionsRepositoriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 

### Return type

[**PaginatedRepositoryPermissions**](paginated_repository_permissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet**
> PaginatedRepositoryPermissions WorkspacesWorkspacePermissionsRepositoriesRepoSlugGet(ctx, repoSlug, workspace, optional)
List a repository permissions for a workspace

Returns an object for the repository permission of each user in the requested repository.  Permissions returned are effective permissions: the highest level of permission the user has. This does not distinguish between direct and indirect (group) privileges.  Only users with admin permission for the repository may access this resource.  Permissions can be:  * `admin` * `write` * `read`  Example:  ``` $ curl https://api.bitbucket.org/2.0/workspaces/atlassian_tutorial/permissions/repositories/geordi  {   \"pagelen\": 10,   \"values\": [     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Erik van Zijst\",         \"uuid\": \"{d301aafa-d676-4ee0-88be-962be7417567}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"geordi\",         \"full_name\": \"atlassian_tutorial/geordi\",         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"       },       \"permission\": \"admin\"     },     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Sean Conaty\",         \"uuid\": \"{504c3b62-8120-4f0c-a7bc-87800b9d6f70}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"geordi\",         \"full_name\": \"atlassian_tutorial/geordi\",         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"       },       \"permission\": \"write\"     }   ],   \"page\": 1,   \"size\": 2 } ```  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by user, or permission by adding the following query string parameters:  * `q=permission>\"read\"` * `sort=user.display_name`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***WorkspacesApiWorkspacesWorkspacePermissionsRepositoriesRepoSlugGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkspacesApiWorkspacesWorkspacePermissionsRepositoriesRepoSlugGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 

### Return type

[**PaginatedRepositoryPermissions**](paginated_repository_permissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsGet**
> PaginatedProjects WorkspacesWorkspaceProjectsGet(ctx, workspace)
List projects in a workspace

Returns the list of projects in this workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedProjects**](paginated_projects.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyGet**
> Project WorkspacesWorkspaceProjectsProjectKeyGet(ctx, projectKey, workspace)
Get a project for a workspace

Returns the requested project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Project**](project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

