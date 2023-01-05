# \IssueTrackerApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugComponentsComponentIdGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugComponentsComponentIdGet) | **Get** /repositories/{workspace}/{repo_slug}/components/{component_id} | Get a component for issues
[**RepositoriesWorkspaceRepoSlugComponentsGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugComponentsGet) | **Get** /repositories/{workspace}/{repo_slug}/components | List components
[**RepositoriesWorkspaceRepoSlugIssuesExportPost**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesExportPost) | **Post** /repositories/{workspace}/{repo_slug}/issues/export | Export issues
[**RepositoriesWorkspaceRepoSlugIssuesExportRepoNameIssuesTaskIdZipGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesExportRepoNameIssuesTaskIdZipGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/export/{repo_name}-issues-{task_id}.zip | Check issue export status
[**RepositoriesWorkspaceRepoSlugIssuesGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesGet) | **Get** /repositories/{workspace}/{repo_slug}/issues | List issues
[**RepositoriesWorkspaceRepoSlugIssuesImportGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesImportGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/import | Check issue import status
[**RepositoriesWorkspaceRepoSlugIssuesImportPost**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesImportPost) | **Post** /repositories/{workspace}/{repo_slug}/issues/import | Import issues
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/attachments | List attachments for an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathDelete**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathDelete) | **Delete** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/attachments/{path} | Delete an attachment for an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/attachments/{path} | Get attachment for an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPost**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPost) | **Post** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/attachments | Upload an attachment to an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesChangeIdGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesChangeIdGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/changes/{change_id} | Get issue change object
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/changes | List changes on an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesPost**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesPost) | **Post** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/changes | Modify the state of an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdDelete**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdDelete) | **Delete** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/comments/{comment_id} | Delete a comment on an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/comments/{comment_id} | Get a comment on an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdPut**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdPut) | **Put** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/comments/{comment_id} | Update a comment on an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/comments | List comments on an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsPost**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsPost) | **Post** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/comments | Create a comment on an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdDelete**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdDelete) | **Delete** /repositories/{workspace}/{repo_slug}/issues/{issue_id} | Delete an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id} | Get an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdPut**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdPut) | **Put** /repositories/{workspace}/{repo_slug}/issues/{issue_id} | Update an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdVoteDelete**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdVoteDelete) | **Delete** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/vote | Remove vote for an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdVoteGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdVoteGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/vote | Check if current user voted for an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdVotePut**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdVotePut) | **Put** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/vote | Vote for an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchDelete**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchDelete) | **Delete** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch | Stop watching an issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchGet) | **Get** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch | Check if current user is watching a issue
[**RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchPut**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchPut) | **Put** /repositories/{workspace}/{repo_slug}/issues/{issue_id}/watch | Watch an issue
[**RepositoriesWorkspaceRepoSlugIssuesPost**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugIssuesPost) | **Post** /repositories/{workspace}/{repo_slug}/issues | Create an issue
[**RepositoriesWorkspaceRepoSlugMilestonesGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugMilestonesGet) | **Get** /repositories/{workspace}/{repo_slug}/milestones | List milestones
[**RepositoriesWorkspaceRepoSlugMilestonesMilestoneIdGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugMilestonesMilestoneIdGet) | **Get** /repositories/{workspace}/{repo_slug}/milestones/{milestone_id} | Get a milestone
[**RepositoriesWorkspaceRepoSlugVersionsGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugVersionsGet) | **Get** /repositories/{workspace}/{repo_slug}/versions | List defined versions for issues
[**RepositoriesWorkspaceRepoSlugVersionsVersionIdGet**](IssueTrackerApi.md#RepositoriesWorkspaceRepoSlugVersionsVersionIdGet) | **Get** /repositories/{workspace}/{repo_slug}/versions/{version_id} | Get a defined version for issues


# **RepositoriesWorkspaceRepoSlugComponentsComponentIdGet**
> Component RepositoriesWorkspaceRepoSlugComponentsComponentIdGet(ctx, componentId, repoSlug, workspace)
Get a component for issues

Returns the specified issue tracker component object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentId** | **int32**| The component&#39;s id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Component**](component.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugComponentsGet**
> PaginatedComponents RepositoriesWorkspaceRepoSlugComponentsGet(ctx, repoSlug, workspace)
List components

Returns the components that have been defined in the issue tracker.  This resource is only available on repositories that have the issue tracker enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedComponents**](paginated_components.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesExportPost**
> RepositoriesWorkspaceRepoSlugIssuesExportPost(ctx, repoSlug, workspace, optional)
Export issues

A POST request to this endpoint initiates a new background celery task that archives the repo's issues.  For example, you can run:  curl -u <username> -X POST http://api.bitbucket.org/2.0/repositories/<owner_username>/<repo_slug>/ issues/export  When the job has been accepted, it will return a 202 (Accepted) along with a unique url to this job in the 'Location' response header. This url is the endpoint for where the user can obtain their zip files.\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***IssueTrackerApiRepositoriesWorkspaceRepoSlugIssuesExportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueTrackerApiRepositoriesWorkspaceRepoSlugIssuesExportPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ExportOptions**](ExportOptions.md)| The options to apply to the export. Available options include &#x60;project_key&#x60; and &#x60;project_name&#x60; which, if specified, are used as the project key and name in the exported Jira json format. Option &#x60;send_email&#x60; specifies whether an email should be sent upon export result. Option &#x60;include_attachments&#x60; specifies whether attachments are included in the export. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesExportRepoNameIssuesTaskIdZipGet**
> IssueJobStatus RepositoriesWorkspaceRepoSlugIssuesExportRepoNameIssuesTaskIdZipGet(ctx, repoName, repoSlug, taskId, workspace)
Check issue export status

This endpoint is used to poll for the progress of an issue export job and return the zip file after the job is complete. As long as the job is running, this will return a 200 response with in the response body a description of the current status.  After the job has been scheduled, but before it starts executing, this endpoint's response is:  {  \"type\": \"issue_job_status\",  \"status\": \"ACCEPTED\",  \"phase\": \"Initializing\",  \"total\": 0,  \"count\": 0,  \"pct\": 0 }   Then once it starts running, it becomes:  {  \"type\": \"issue_job_status\",  \"status\": \"STARTED\",  \"phase\": \"Attachments\",  \"total\": 15,  \"count\": 11,  \"pct\": 73 }  Once the job has successfully completed, it returns a stream of the zip file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of the repo | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **taskId** | **string**| The ID of the export task | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**IssueJobStatus**](issue_job_status.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesGet**
> PaginatedIssues RepositoriesWorkspaceRepoSlugIssuesGet(ctx, repoSlug, workspace)
List issues

Returns the issues in the issue tracker.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedIssues**](paginated_issues.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesImportGet**
> IssueJobStatus RepositoriesWorkspaceRepoSlugIssuesImportGet(ctx, repoSlug, workspace)
Check issue import status

When using GET, this endpoint reports the status of the current import task. Request example:  ``` $ curl -u <username> -X GET https://api.bitbucket.org/2.0/repositories/<owner_username>/<repo_slug>/issues/import ```  After the job has been scheduled, but before it starts executing, this endpoint's response is:  ``` < HTTP/1.1 202 Accepted {     \"type\": \"issue_job_status\",     \"status\": \"PENDING\",     \"phase\": \"Attachments\",     \"total\": 15,     \"count\": 0,     \"percent\": 0 } ```  Once it starts running, it is a 202 response with status STARTED and progress filled.  After it is finished, it becomes a 200 response with status SUCCESS or FAILURE.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**IssueJobStatus**](issue_job_status.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesImportPost**
> IssueJobStatus RepositoriesWorkspaceRepoSlugIssuesImportPost(ctx, repoSlug, workspace)
Import issues

A POST request to this endpoint will import the zip file given by the archive parameter into the repository. All existing issues will be deleted and replaced by the contents of the imported zip file.  Imports are done through a multipart/form-data POST. There is one valid and required form field, with the name \"archive,\" which needs to be a file field:  ``` $ curl -u <username> -X POST -F archive=@/path/to/file.zip https://api.bitbucket.org/2.0/repositories/<owner_username>/<repo_slug>/issues/import ```  When the import job is accepted, here is example output:  ``` < HTTP/1.1 202 Accepted  {     \"type\": \"issue_job_status\",     \"status\": \"ACCEPTED\",     \"phase\": \"Attachments\",     \"total\": 15,     \"count\": 0,     \"percent\": 0 } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**IssueJobStatus**](issue_job_status.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsGet**
> PaginatedIssueAttachments RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsGet(ctx, issueId, repoSlug, workspace)
List attachments for an issue

Returns all attachments for this issue.  This returns the files' meta data. This does not return the files' actual contents.  The files are always ordered by their upload date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedIssueAttachments**](paginated_issue_attachments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathDelete**
> RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathDelete(ctx, issueId, path, repoSlug, workspace)
Delete an attachment for an issue

Deletes an attachment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **path** | **string**| Path to the file. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathGet**
> RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPathGet(ctx, issueId, path, repoSlug, workspace)
Get attachment for an issue

Returns the contents of the specified file attachment.  Note that this endpoint does not return a JSON response, but instead returns a redirect pointing to the actual file that in turn will return the raw contents.  The redirect URL contains a one-time token that has a limited lifetime. As a result, the link should not be persisted, stored, or shared.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **path** | **string**| Path to the file. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPost**
> RepositoriesWorkspaceRepoSlugIssuesIssueIdAttachmentsPost(ctx, issueId, repoSlug, workspace)
Upload an attachment to an issue

Upload new issue attachments.  To upload files, perform a `multipart/form-data` POST containing one or more file fields.  When a file is uploaded with the same name as an existing attachment, then the existing file will be replaced.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesChangeIdGet**
> IssueChange RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesChangeIdGet(ctx, changeId, issueId, repoSlug, workspace)
Get issue change object

Returns the specified issue change object.  This resource is only available on repositories that have the issue tracker enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changeId** | **string**| The issue change id | 
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**IssueChange**](issue_change.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesGet**
> PaginatedLogEntries RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesGet(ctx, issueId, repoSlug, workspace, optional)
List changes on an issue

Returns the list of all changes that have been made to the specified issue. Changes are returned in chronological order with the oldest change first.  Each time an issue is edited in the UI or through the API, an immutable change record is created under the `/issues/123/changes` endpoint. It also has a comment associated with the change.  Note that this operation is changing significantly, due to privacy changes. See the [announcement](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#changes-to-the-issue-changes-api) for details.  ``` $ curl -s https://api.bitbucket.org/2.0/repositories/evzijst/dogslow/issues/1/changes - | jq .  {   \"pagelen\": 20,   \"values\": [     {       \"changes\": {         \"priority\": {           \"new\": \"trivial\",           \"old\": \"major\"         },         \"assignee\": {           \"new\": \"\",           \"old\": \"evzijst\"         },         \"assignee_account_id\": {           \"new\": \"\",           \"old\": \"557058:c0b72ad0-1cb5-4018-9cdc-0cde8492c443\"         },         \"kind\": {           \"new\": \"enhancement\",           \"old\": \"bug\"         }       },       \"links\": {         \"self\": {           \"href\": \"https://api.bitbucket.org/2.0/repositories/evzijst/dogslow/issues/1/changes/2\"         },         \"html\": {           \"href\": \"https://bitbucket.org/evzijst/dogslow/issues/1#comment-2\"         }       },       \"issue\": {         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/evzijst/dogslow/issues/1\"           }         },         \"type\": \"issue\",         \"id\": 1,         \"repository\": {           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/evzijst/dogslow\"             },             \"html\": {               \"href\": \"https://bitbucket.org/evzijst/dogslow\"             },             \"avatar\": {               \"href\": \"https://bitbucket.org/evzijst/dogslow/avatar/32/\"             }           },           \"type\": \"repository\",           \"name\": \"dogslow\",           \"full_name\": \"evzijst/dogslow\",           \"uuid\": \"{988b17c6-1a47-4e70-84ee-854d5f012bf6}\"         },         \"title\": \"Updated title\"       },       \"created_on\": \"2018-03-03T00:35:28.353630+00:00\",       \"user\": {         \"username\": \"evzijst\",         \"nickname\": \"evzijst\",         \"display_name\": \"evzijst\",         \"type\": \"user\",         \"uuid\": \"{aaa7972b-38af-4fb1-802d-6e3854c95778}\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/users/evzijst\"           },           \"html\": {             \"href\": \"https://bitbucket.org/evzijst/\"           },           \"avatar\": {             \"href\": \"https://bitbucket.org/account/evzijst/avatar/32/\"           }         }       },       \"message\": {         \"raw\": \"Removed assignee, changed kind and priority.\",         \"markup\": \"markdown\",         \"html\": \"<p>Removed assignee, changed kind and priority.</p>\",         \"type\": \"rendered\"       },       \"type\": \"issue_change\",       \"id\": 2     }   ],   \"page\": 1 } ```  Changes support [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) that can be used to search for specific changes. For instance, to see when an issue transitioned to \"resolved\":  ``` $ curl -s https://api.bitbucket.org/2.0/repositories/site/master/issues/1/changes \\    -G --data-urlencode='q=changes.state.new = \"resolved\"' ```  This resource is only available on repositories that have the issue tracker enabled.  N.B.  The `changes.assignee` and `changes.assignee_account_id` fields are not a `user` object. Instead, they contain the raw `username` and `account_id` of the user. This is to protect the integrity of the audit log even after a user account gets deleted.  The `changes.assignee` field is deprecated will disappear in the future. Use `changes.assignee_account_id` instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***IssueTrackerApiRepositoriesWorkspaceRepoSlugIssuesIssueIdChangesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueTrackerApiRepositoriesWorkspaceRepoSlugIssuesIssueIdChangesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **optional.String**|  Query string to narrow down the response. See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for details. | 
 **sort** | **optional.String**|  Name of a response property to sort results. See [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results) for details.  | 

### Return type

[**PaginatedLogEntries**](paginated_log_entries.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesPost**
> IssueChange RepositoriesWorkspaceRepoSlugIssuesIssueIdChangesPost(ctx, issueId, repoSlug, workspace, body)
Modify the state of an issue

Makes a change to the specified issue.  For example, to change an issue's state and assignee, create a new change object that modifies these fields:  ``` curl https://api.bitbucket.org/2.0/site/master/issues/1234/changes \\   -s -u evzijst -X POST -H \"Content-Type: application/json\" \\   -d '{     \"changes\": {       \"assignee_account_id\": {         \"new\": \"557058:c0b72ad0-1cb5-4018-9cdc-0cde8492c443\"       },       \"state\": {         \"new\": 'resolved\"       }     }     \"message\": {       \"raw\": \"This is now resolved.\"     }   }' ```  The above example also includes a custom comment to go alongside the change. This comment will also be visible on the issue page in the UI.  The fields of the `changes` object are strings, not objects. This allows for immutable change log records, even after user accounts, milestones, or other objects recorded in a change entry, get renamed or deleted.  The `assignee_account_id` field stores the account id. When POSTing a new change and changing the assignee, the client should therefore use the user's account_id in the `changes.assignee_account_id.new` field.  This call requires authentication. Private repositories or private issue trackers require the caller to authenticate with an account that has appropriate authorization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
  **body** | [**IssueChange**](IssueChange.md)| The new issue state change. The only required elements are &#x60;changes.[].new&#x60;. All other elements can be omitted from the body. | 

### Return type

[**IssueChange**](issue_change.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdDelete**
> RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdDelete(ctx, commentId, issueId, repoSlug, workspace)
Delete a comment on an issue

Deletes the specified comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **int32**| The id of the comment. | 
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdGet**
> IssueComment RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdGet(ctx, commentId, issueId, repoSlug, workspace)
Get a comment on an issue

Returns the specified issue comment object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **int32**| The id of the comment. | 
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**IssueComment**](issue_comment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdPut**
> IssueComment RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsCommentIdPut(ctx, commentId, issueId, repoSlug, workspace, body)
Update a comment on an issue

Updates the content of the specified issue comment. Note that only the `content.raw` field can be modified.  ``` $ curl https://api.bitbucket.org/2.0/repositories/atlassian/prlinks/issues/42/comments/5728901 \\   -X PUT -u evzijst \\   -H 'Content-Type: application/json' \\   -d '{\"content\": {\"raw\": \"Lorem ipsum.\"}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **int32**| The id of the comment. | 
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
  **body** | [**IssueComment**](IssueComment.md)| The updated comment. | 

### Return type

[**IssueComment**](issue_comment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsGet**
> PaginatedIssueComments RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsGet(ctx, issueId, repoSlug, workspace, optional)
List comments on an issue

Returns a paginated list of all comments that were made on the specified issue.  The default sorting is oldest to newest and can be overridden with the `sort` query parameter.  This endpoint also supports filtering and sorting of the results. See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***IssueTrackerApiRepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueTrackerApiRepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 

### Return type

[**PaginatedIssueComments**](paginated_issue_comments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsPost**
> RepositoriesWorkspaceRepoSlugIssuesIssueIdCommentsPost(ctx, issueId, repoSlug, workspace, body)
Create a comment on an issue

Creates a new issue comment.  ``` $ curl https://api.bitbucket.org/2.0/repositories/atlassian/prlinks/issues/42/comments/ \\   -X POST -u evzijst \\   -H 'Content-Type: application/json' \\   -d '{\"content\": {\"raw\": \"Lorem ipsum.\"}}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
  **body** | [**IssueComment**](IssueComment.md)| The new issue comment object. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdDelete**
> Issue RepositoriesWorkspaceRepoSlugIssuesIssueIdDelete(ctx, issueId, repoSlug, workspace)
Delete an issue

Deletes the specified issue. This requires write access to the repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Issue**](issue.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdGet**
> Issue RepositoriesWorkspaceRepoSlugIssuesIssueIdGet(ctx, issueId, repoSlug, workspace)
Get an issue

Returns the specified issue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Issue**](issue.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdPut**
> Issue RepositoriesWorkspaceRepoSlugIssuesIssueIdPut(ctx, issueId, repoSlug, workspace)
Update an issue

Modifies the issue.  ``` $ curl https://api.bitbucket.org/2.0/repostories/evzijst/dogslow/issues/123 \\   -u evzijst -s -X PUT -H 'Content-Type: application/json' \\   -d '{   \"title\": \"Updated title\",   \"assignee\": {     \"account_id\": \"5d5355e8c6b9320d9ea5b28d\"   },   \"priority\": \"minor\",   \"version\": {     \"name\": \"1.0\"   },   \"component\": null }' ```  This example changes the `title`, `assignee`, `priority` and the `version`. It also removes the value of the `component` from the issue by setting the field to `null`. Any field not present keeps its existing value.  Each time an issue is edited in the UI or through the API, an immutable change record is created under the `/issues/123/changes` endpoint. It also has a comment associated with the change.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Issue**](issue.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdVoteDelete**
> ModelError RepositoriesWorkspaceRepoSlugIssuesIssueIdVoteDelete(ctx, issueId, repoSlug, workspace)
Remove vote for an issue

Retract your vote.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdVoteGet**
> ModelError RepositoriesWorkspaceRepoSlugIssuesIssueIdVoteGet(ctx, issueId, repoSlug, workspace)
Check if current user voted for an issue

Check whether the authenticated user has voted for this issue. A 204 status code indicates that the user has voted, while a 404 implies they haven't.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdVotePut**
> ModelError RepositoriesWorkspaceRepoSlugIssuesIssueIdVotePut(ctx, issueId, repoSlug, workspace)
Vote for an issue

Vote for this issue.  To cast your vote, do an empty PUT. The 204 status code indicates that the operation was successful.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchDelete**
> ModelError RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchDelete(ctx, issueId, repoSlug, workspace)
Stop watching an issue

Stop watching this issue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchGet**
> ModelError RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchGet(ctx, issueId, repoSlug, workspace)
Check if current user is watching a issue

Indicated whether or not the authenticated user is watching this issue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchPut**
> ModelError RepositoriesWorkspaceRepoSlugIssuesIssueIdWatchPut(ctx, issueId, repoSlug, workspace)
Watch an issue

Start watching this issue.  To start watching this issue, do an empty PUT. The 204 status code indicates that the operation was successful.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| The issue id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**ModelError**](error.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugIssuesPost**
> Issue RepositoriesWorkspaceRepoSlugIssuesPost(ctx, repoSlug, workspace, body)
Create an issue

Creates a new issue.  This call requires authentication. Private repositories or private issue trackers require the caller to authenticate with an account that has appropriate authorization.  The authenticated user is used for the issue's `reporter` field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
  **body** | [**Issue**](Issue.md)| The new issue. The only required element is &#x60;title&#x60;. All other elements can be omitted from the body. | 

### Return type

[**Issue**](issue.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugMilestonesGet**
> PaginatedMilestones RepositoriesWorkspaceRepoSlugMilestonesGet(ctx, repoSlug, workspace)
List milestones

Returns the milestones that have been defined in the issue tracker.  This resource is only available on repositories that have the issue tracker enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedMilestones**](paginated_milestones.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugMilestonesMilestoneIdGet**
> Milestone RepositoriesWorkspaceRepoSlugMilestonesMilestoneIdGet(ctx, milestoneId, repoSlug, workspace)
Get a milestone

Returns the specified issue tracker milestone object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **milestoneId** | **int32**| The milestone&#39;s id | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Milestone**](milestone.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugVersionsGet**
> PaginatedVersions RepositoriesWorkspaceRepoSlugVersionsGet(ctx, repoSlug, workspace)
List defined versions for issues

Returns the versions that have been defined in the issue tracker.  This resource is only available on repositories that have the issue tracker enabled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedVersions**](paginated_versions.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugVersionsVersionIdGet**
> Version RepositoriesWorkspaceRepoSlugVersionsVersionIdGet(ctx, repoSlug, versionId, workspace)
Get a defined version for issues

Returns the specified issue tracker version object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **versionId** | **int32**| The version&#39;s id | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Version**](version.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

