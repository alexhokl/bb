# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsGet**](TeamsApi.md#TeamsGet) | **Get** /teams | List teams a user is part of
[**TeamsUsernameFollowersGet**](TeamsApi.md#TeamsUsernameFollowersGet) | **Get** /teams/{username}/followers | List team followers
[**TeamsUsernameFollowingGet**](TeamsApi.md#TeamsUsernameFollowingGet) | **Get** /teams/{username}/following | List accounts a team is following
[**TeamsUsernameGet**](TeamsApi.md#TeamsUsernameGet) | **Get** /teams/{username} | Get a team
[**TeamsUsernameMembersGet**](TeamsApi.md#TeamsUsernameMembersGet) | **Get** /teams/{username}/members | List team members
[**TeamsUsernamePermissionsGet**](TeamsApi.md#TeamsUsernamePermissionsGet) | **Get** /teams/{username}/permissions | List team permissions for a user 
[**TeamsUsernamePermissionsRepositoriesGet**](TeamsApi.md#TeamsUsernamePermissionsRepositoriesGet) | **Get** /teams/{username}/permissions/repositories | List repository permissions for a team
[**TeamsUsernamePermissionsRepositoriesRepoSlugGet**](TeamsApi.md#TeamsUsernamePermissionsRepositoriesRepoSlugGet) | **Get** /teams/{username}/permissions/repositories/{repo_slug} | List repository permissions for a team
[**TeamsWorkspaceRepositoriesGet**](TeamsApi.md#TeamsWorkspaceRepositoriesGet) | **Get** /teams/{workspace}/repositories | List workspace repositories
[**UserPermissionsTeamsGet**](TeamsApi.md#UserPermissionsTeamsGet) | **Get** /user/permissions/teams | List team permissions for the user
[**UsersWorkspaceRepositoriesGet**](TeamsApi.md#UsersWorkspaceRepositoriesGet) | **Get** /users/{workspace}/repositories | List workspace repositories

# **TeamsGet**
> PaginatedTeams TeamsGet(ctx, optional)
List teams a user is part of

Returns all the teams that the authenticated user is associated with.  **This endpoint has been deprecated and will stop functioning soon. You should use the [workspaces](/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-get) endpoint instead. For more information, see [this post](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiTeamsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **optional.String**|  Filters the teams based on the authenticated user&#x27;s role on each team.  * **member**: returns a list of all the teams which the caller is a member of   at least one team group or repository owned by the team * **contributor**: returns a list of teams which the caller has write access   to at least one repository owned by the team * **admin**: returns a list teams which the caller has team administrator access  | 

### Return type

[**PaginatedTeams**](paginated_teams.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsUsernameFollowersGet**
> PaginatedUsers TeamsUsernameFollowersGet(ctx, username)
List team followers

Returns the list of accounts that are following this team.  **This endpoint has been deprecated and will stop functioning soon. There is no replacement endpoint. For more information, see [this post](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Return type

[**PaginatedUsers**](paginated_users.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsUsernameFollowingGet**
> PaginatedUsers TeamsUsernameFollowingGet(ctx, username)
List accounts a team is following

Returns the list of accounts this team is following.  **This endpoint has been deprecated and will stop functioning soon. There is no replacement endpoint. For more information, see [this post](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Return type

[**PaginatedUsers**](paginated_users.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsUsernameGet**
> Team TeamsUsernameGet(ctx, username)
Get a team

Gets the public information associated with a team.  If the team's profile is private, `location`, `website` and `created_on` elements are omitted.  **This endpoint has been deprecated and will stop functioning on August 25th, 2021. You should use the [workspace](/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-get) endpoint instead. For more information, see [this post](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 

### Return type

[**Team**](team.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsUsernameMembersGet**
> User TeamsUsernameMembersGet(ctx, username)
List team members

Returns all members of the specified team. Any member of any of the team's groups is considered a member of the team. This includes users in groups that may not actually have access to any of the team's repositories.  **This operation has been deprecated due to privacy changes. See the [announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/) for details. You should this [workspaces](/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get) endpoint as a replacement.**

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

# **TeamsUsernamePermissionsGet**
> PaginatedTeamPermissions TeamsUsernamePermissionsGet(ctx, username, optional)
List team permissions for a user 

Returns an object for each team permission a user on the team has.  **This endpoint has been deprecated and will stop functioning soon. You should use the [workspace permissions](/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-member-get) endpoint instead. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**  Permissions returned are effective permissions — if a user is a member of multiple groups with distinct roles, only the highest level is returned.  Permissions can be:  * `admin` * `collaborator`  Only users with admin permission for the team may access this resource.  Example:  ``` $ curl https://api.bitbucket.org/2.0/teams/atlassian_tutorial/permissions  {   \"pagelen\": 10,   \"values\": [     {       \"permission\": \"admin\",       \"type\": \"team_permission\",       \"user\": {         \"type\": \"user\",         \"nickname\": \"evzijst\",         \"display_name\": \"Erik van Zijst\",         \"uuid\": \"{d301aafa-d676-4ee0-88be-962be7417567}\"       },       \"team\": {         \"display_name\": \"Atlassian Bitbucket\",         \"uuid\": \"{4cc6108a-a241-4db0-96a5-64347ac04f87}\"       }     },     {       \"permission\": \"collaborator\",       \"type\": \"team_permission\",       \"user\": {         \"type\": \"user\",         \"nickname\": \"seanaty\",         \"display_name\": \"Sean Conaty\",         \"uuid\": \"{504c3b62-8120-4f0c-a7bc-87800b9d6f70}\"       },       \"team\": {         \"display_name\": \"Atlassian Bitbucket\",         \"uuid\": \"{4cc6108a-a241-4db0-96a5-64347ac04f87}\"       }     }   ],   \"page\": 1,   \"size\": 2 } ```  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by team, user, or permission by adding the following query string parameters:  * `q=user.uuid=\"{d301aafa-d676-4ee0-88be-962be7417567}\"` or `q=permission=\"admin\"` * `sort=team.display_name`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 
 **optional** | ***TeamsApiTeamsUsernamePermissionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsUsernamePermissionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 

### Return type

[**PaginatedTeamPermissions**](paginated_team_permissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsUsernamePermissionsRepositoriesGet**
> PaginatedRepositoryPermissions TeamsUsernamePermissionsRepositoriesGet(ctx, username, optional)
List repository permissions for a team

Returns an object for each repository permission for all of a team’s repositories.  **This endpoint has been deprecated and will stop functioning soon. You should use the [workspace repository permissions](/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-permissions-repositories-get) endpoint instead. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**  If the username URL parameter refers to a user account instead of a team account, an object containing the repository permissions of all the username's repositories will be returned.  Permissions returned are effective permissions — the highest level of permission the user has. This does not include public repositories that users are not granted any specific permission in, and does not distinguish between direct and indirect privileges.  Only users with admin permission for the team may access this resource.  Permissions can be:  * `admin` * `write` * `read`  Example:  ``` $ curl https://api.bitbucket.org/2.0/teams/atlassian_tutorial/permissions/repositories  {   \"pagelen\": 10,   \"values\": [     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Erik van Zijst\",         \"uuid\": \"{d301aafa-d676-4ee0-88be-962be7417567}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"geordi\",         \"full_name\": \"bitbucket/geordi\",         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"       },       \"permission\": \"admin\"     },     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Sean Conaty\",         \"uuid\": \"{504c3b62-8120-4f0c-a7bc-87800b9d6f70}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"geordi\",         \"full_name\": \"bitbucket/geordi\",         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"       },       \"permission\": \"write\"     }   ],   \"page\": 1,   \"size\": 2 } ```  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by repository, user, or permission by adding the following query string parameters:  * `q=repository.name=\"geordi\"` or `q=permission>\"read\"` * `sort=user.display_name`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 
 **optional** | ***TeamsApiTeamsUsernamePermissionsRepositoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsUsernamePermissionsRepositoriesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 

### Return type

[**PaginatedRepositoryPermissions**](paginated_repository_permissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamsUsernamePermissionsRepositoriesRepoSlugGet**
> PaginatedRepositoryPermissions TeamsUsernamePermissionsRepositoriesRepoSlugGet(ctx, repoSlug, username, optional)
List repository permissions for a team

Returns an object for each repository permission of a given repository.  **This endpoint has been deprecated and will stop functioning soon. You should use the [workspace repository permissions](/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-permissions-repositories-repo-slug-get) endpoint instead. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**  If the username URL parameter refers to a user account instead of a team account, an object containing the repository permissions of the username's repository will be returned.  Permissions returned are effective permissions — the highest level of permission the user has. This does not include public repositories that users are not granted any specific permission in, and does not distinguish between direct and indirect privileges.  Only users with admin permission for the repository may access this resource.  Permissions can be:  * `admin` * `write` * `read`  Example:  ``` $ curl https://api.bitbucket.org/2.0/teams/atlassian_tutorial/permissions/repositories/geordi  {   \"pagelen\": 10,   \"values\": [     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Erik van Zijst\",         \"uuid\": \"{d301aafa-d676-4ee0-88be-962be7417567}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"geordi\",         \"full_name\": \"bitbucket/geordi\",         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"       },       \"permission\": \"admin\"     },     {       \"type\": \"repository_permission\",       \"user\": {         \"type\": \"user\",         \"display_name\": \"Sean Conaty\",         \"uuid\": \"{504c3b62-8120-4f0c-a7bc-87800b9d6f70}\"       },       \"repository\": {         \"type\": \"repository\",         \"name\": \"geordi\",         \"full_name\": \"bitbucket/geordi\",         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"       },       \"permission\": \"write\"     }   ],   \"page\": 1,   \"size\": 2 } ```  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by user, or permission by adding the following query string parameters:  * `q=permission>\"read\"` * `sort=user.display_name`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **username** | **string**| This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user.  | 
 **optional** | ***TeamsApiTeamsUsernamePermissionsRepositoriesRepoSlugGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamsUsernamePermissionsRepositoriesRepoSlugGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 

### Return type

[**PaginatedRepositoryPermissions**](paginated_repository_permissions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **UserPermissionsTeamsGet**
> PaginatedTeamPermissions UserPermissionsTeamsGet(ctx, optional)
List team permissions for the user

Returns an object for each team the caller is a member of, and their effective role — the highest level of privilege the caller has. If a user is a member of multiple groups with distinct roles, only the highest level is returned.  **This endpoint has been deprecated and will stop functioning soon. You should use the [workspace permissions](/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-get) endpoint instead. For more information, see [the announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-teams-deprecation/).**  Permissions can be:  * `admin` * `collaborator`  Example:  ``` $ curl https://api.bitbucket.org/2.0/user/permissions/teams  {   \"pagelen\": 10,   \"values\": [     {       \"permission\": \"admin\",       \"type\": \"team_permission\",       \"user\": {         \"type\": \"user\",         \"nickname\": \"evzijst\",         \"display_name\": \"Erik van Zijst\",         \"uuid\": \"{d301aafa-d676-4ee0-88be-962be7417567}\"       },       \"team\": {         \"display_name\": \"Atlassian Bitbucket\",         \"uuid\": \"{4cc6108a-a241-4db0-96a5-64347ac04f87}\"       }     }   ],   \"page\": 1,   \"size\": 1 } ```  Results may be further [filtered or sorted](/cloud/bitbucket/rest/intro/#filtering) by team or permission by adding the following query string parameters:  * `q=team.uuid=\"{4cc6108a-a241-4db0-96a5-64347ac04f87}\"` or `q=permission=\"admin\"` * `sort=team.display_name`  Note that the query parameter values need to be URL escaped so that `=` would become `%3D`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiUserPermissionsTeamsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiUserPermissionsTeamsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 

### Return type

[**PaginatedTeamPermissions**](paginated_team_permissions.md)

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

