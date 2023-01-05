# \ProjectsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspacesWorkspaceProjectsPost**](ProjectsApi.md#WorkspacesWorkspaceProjectsPost) | **Post** /workspaces/{workspace}/projects | Create a project in a workspace
[**WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersGet**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersGet) | **Get** /workspaces/{workspace}/projects/{project_key}/default-reviewers | List the default reviewers in a project
[**WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserDelete**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserDelete) | **Delete** /workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user} | Remove the specific user from the project&#39;s default reviewers
[**WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserGet**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserGet) | **Get** /workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user} | Get a default reviewer
[**WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserPut**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserPut) | **Put** /workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user} | Add the specific user as a default reviewer for the project
[**WorkspacesWorkspaceProjectsProjectKeyDelete**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyDelete) | **Delete** /workspaces/{workspace}/projects/{project_key} | Delete a project for a workspace
[**WorkspacesWorkspaceProjectsProjectKeyGet**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyGet) | **Get** /workspaces/{workspace}/projects/{project_key} | Get a project for a workspace
[**WorkspacesWorkspaceProjectsProjectKeyPut**](ProjectsApi.md#WorkspacesWorkspaceProjectsProjectKeyPut) | **Put** /workspaces/{workspace}/projects/{project_key} | Update a project for a workspace


# **WorkspacesWorkspaceProjectsPost**
> Project WorkspacesWorkspaceProjectsPost(ctx, workspace, body)
Create a project in a workspace

Creates a new project.  Note that the avatar has to be embedded as either a data-url or a URL to an external image as shown in the examples below:  ``` $ body=$(cat << EOF {     \"name\": \"Mars Project\",     \"key\": \"MARS\",     \"description\": \"Software for colonizing mars.\",     \"links\": {         \"avatar\": {             \"href\": \"data:image/gif;base64,R0lGODlhEAAQAMQAAORHHOVSKudfOulrSOp3WOyDZu6QdvCchPGolfO0o/...\"         }     },     \"is_private\": false } EOF ) $ curl -H \"Content-Type: application/json\" \\        -X POST \\        -d \"$body\" \\        https://api.bitbucket.org/2.0/teams/teams-in-space/projects/ | jq . {   // Serialized project document } ```  or even:  ``` $ body=$(cat << EOF {     \"name\": \"Mars Project\",     \"key\": \"MARS\",     \"description\": \"Software for colonizing mars.\",     \"links\": {         \"avatar\": {             \"href\": \"http://i.imgur.com/72tRx4w.gif\"         }     },     \"is_private\": false } EOF ) $ curl -H \"Content-Type: application/json\" \\        -X POST \\        -d \"$body\" \\        https://api.bitbucket.org/2.0/teams/teams-in-space/projects/ | jq . {   // Serialized project document } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
  **body** | [**Project**](Project.md)|  | 

### Return type

[**Project**](project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersGet**
> PaginatedDefaultReviewerAndType WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersGet(ctx, projectKey, workspace)
List the default reviewers in a project

Return a list of all default reviewers for a project. This is a list of users that will be added as default reviewers to pull requests for any repository within the project.  Example: ``` $ curl https://api.bitbucket.org/2.0/.../projects/.../default-reviewers | jq . {     \"pagelen\": 10,     \"values\": [         {             \"user\": {                 \"display_name\": \"Davis Lee\",                 \"uuid\": \"{f0e0e8e9-66c1-4b85-a784-44a9eb9ef1a6}\"             },             \"reviewer_type\": \"project\",             \"type\": \"default_reviewer\"         },         {             \"user\": {                 \"display_name\": \"Jorge Rodriguez\",                 \"uuid\": \"{1aa43376-260d-4a0b-9660-f62672b9655d}\"             },             \"reviewer_type\": \"project\",             \"type\": \"default_reviewer\"         }     ],     \"page\": 1,     \"size\": 2 } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedDefaultReviewerAndType**](paginated_default_reviewer_and_type.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserDelete**
> WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserDelete(ctx, projectKey, selectedUser, workspace)
Remove the specific user from the project's default reviewers

Removes a default reviewer from the project.  Example: ``` $ curl https://api.bitbucket.org/2.0/.../default-reviewers/%7Bf0e0e8e9-66c1-4b85-a784-44a9eb9ef1a6%7D  HTTP/1.1 204 ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This can either be the actual &#x60;key&#x60; assigned to the project or the &#x60;UUID&#x60; (surrounded by curly-braces (&#x60;{}&#x60;)).  | 
  **selectedUser** | **string**| This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserGet**
> User WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserGet(ctx, projectKey, selectedUser, workspace)
Get a default reviewer

Returns the specified default reviewer.  Example: ``` $ curl https://api.bitbucket.org/2.0/.../default-reviewers/%7Bf0e0e8e9-66c1-4b85-a784-44a9eb9ef1a6%7D {     \"display_name\": \"Davis Lee\",     \"type\": \"user\",     \"uuid\": \"{f0e0e8e9-66c1-4b85-a784-44a9eb9ef1a6}\" } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This can either be the actual &#x60;key&#x60; assigned to the project or the &#x60;UUID&#x60; (surrounded by curly-braces (&#x60;{}&#x60;)).  | 
  **selectedUser** | **string**| This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**User**](user.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserPut**
> User WorkspacesWorkspaceProjectsProjectKeyDefaultReviewersSelectedUserPut(ctx, projectKey, selectedUser, workspace)
Add the specific user as a default reviewer for the project

Adds the specified user to the project's list of default reviewers. The method is idempotent. Accepts an optional body containing the `uuid` of the user to be added.  Example: ``` $ curl -XPUT https://api.bitbucket.org/2.0/.../default-reviewers/%7Bf0e0e8e9-66c1-4b85-a784-44a9eb9ef1a6%7D -d { 'uuid': '{f0e0e8e9-66c1-4b85-a784-44a9eb9ef1a6}' }  HTTP/1.1 204 ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This can either be the actual &#x60;key&#x60; assigned to the project or the &#x60;UUID&#x60; (surrounded by curly-braces (&#x60;{}&#x60;)).  | 
  **selectedUser** | **string**| This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**User**](user.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WorkspacesWorkspaceProjectsProjectKeyDelete**
> WorkspacesWorkspaceProjectsProjectKeyDelete(ctx, projectKey, workspace)
Delete a project for a workspace

Deletes this project. This is an irreversible operation.  You cannot delete a project that still contains repositories. To delete the project, [delete](/cloud/bitbucket/rest/api-group-repositories/#api-repositories-workspace-repo-slug-delete) or transfer the repositories first.  Example: ``` $ curl -X DELETE https://api.bitbucket.org/2.0/bbworkspace1/PROJ ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

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

# **WorkspacesWorkspaceProjectsProjectKeyPut**
> Project WorkspacesWorkspaceProjectsProjectKeyPut(ctx, projectKey, workspace, body)
Update a project for a workspace

Since this endpoint can be used to both update and to create a project, the request body depends on the intent.  #### Creation  See the POST documentation for the project collection for an example of the request body.  Note: The `key` should not be specified in the body of request (since it is already present in the URL). The `name` is required, everything else is optional.  #### Update  See the POST documentation for the project collection for an example of the request body.  Note: The key is not required in the body (since it is already in the URL). The key may be specified in the body, if the intent is to change the key itself. In such a scenario, the location of the project is changed and is returned in the `Location` header of the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project in question. This is the actual &#x60;key&#x60; assigned to the project.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
  **body** | [**Project**](Project.md)|  | 

### Return type

[**Project**](project.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

