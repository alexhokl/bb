# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateOrUpdateAnnotations**](CommitsApi.md#BulkCreateOrUpdateAnnotations) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations | Bulk create or update annotations
[**CreateOrUpdateAnnotation**](CommitsApi.md#CreateOrUpdateAnnotation) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | Create or update an annotation
[**CreateOrUpdateReport**](CommitsApi.md#CreateOrUpdateReport) | **Put** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | Create or update a report
[**DeleteAnnotation**](CommitsApi.md#DeleteAnnotation) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | Delete an annotation
[**DeleteReport**](CommitsApi.md#DeleteReport) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | Delete a report
[**GetAnnotation**](CommitsApi.md#GetAnnotation) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations/{annotationId} | Get an annotation
[**GetAnnotationsForReport**](CommitsApi.md#GetAnnotationsForReport) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId}/annotations | List annotations
[**GetReport**](CommitsApi.md#GetReport) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports/{reportId} | Get a report
[**GetReportsForCommit**](CommitsApi.md#GetReportsForCommit) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/reports | List reports
[**RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete) | **Delete** /repositories/{workspace}/{repo_slug}/commit/{commit}/approve | Unapprove a commit
[**RepositoriesWorkspaceRepoSlugCommitCommitApprovePost**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitApprovePost) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/approve | Approve a commit
[**RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/comments/{comment_id} | Get a commit comment
[**RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/comments | List a commit&#x27;s comments
[**RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost) | **Post** /repositories/{workspace}/{repo_slug}/commit/{commit}/comments | Create comment for a commit
[**RepositoriesWorkspaceRepoSlugCommitCommitGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitCommitGet) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit} | Get a commit
[**RepositoriesWorkspaceRepoSlugCommitsGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitsGet) | **Get** /repositories/{workspace}/{repo_slug}/commits | List commits
[**RepositoriesWorkspaceRepoSlugCommitsPost**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitsPost) | **Post** /repositories/{workspace}/{repo_slug}/commits | List commits with include/exclude
[**RepositoriesWorkspaceRepoSlugCommitsRevisionGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitsRevisionGet) | **Get** /repositories/{workspace}/{repo_slug}/commits/{revision} | List commits for revision
[**RepositoriesWorkspaceRepoSlugCommitsRevisionPost**](CommitsApi.md#RepositoriesWorkspaceRepoSlugCommitsRevisionPost) | **Post** /repositories/{workspace}/{repo_slug}/commits/{revision} | List commits for revision using include/exclude
[**RepositoriesWorkspaceRepoSlugDiffSpecGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugDiffSpecGet) | **Get** /repositories/{workspace}/{repo_slug}/diff/{spec} | Compare two commits
[**RepositoriesWorkspaceRepoSlugDiffstatSpecGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugDiffstatSpecGet) | **Get** /repositories/{workspace}/{repo_slug}/diffstat/{spec} | Compare two commit diff stats
[**RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet) | **Get** /repositories/{workspace}/{repo_slug}/merge-base/{revspec} | Get the common ancestor between two commits
[**RepositoriesWorkspaceRepoSlugPatchSpecGet**](CommitsApi.md#RepositoriesWorkspaceRepoSlugPatchSpecGet) | **Get** /repositories/{workspace}/{repo_slug}/patch/{spec} | Get a patch for two commits

# **BulkCreateOrUpdateAnnotations**
> []ReportAnnotation BulkCreateOrUpdateAnnotations(ctx, body, workspace, repoSlug, commit, reportId)
Bulk create or update annotations

Bulk upload of annotations. Annotations are individual findings that have been identified as part of a report, for example, a line of code that represents a vulnerability. These annotations can be attached to a specific file and even a specific line in that file, however, that is optional. Annotations are not mandatory and a report can contain up to 1000 annotations.  Add the annotations you want to upload as objects in a JSON array and make sure each annotation has the external_id field set to a unique value. If you want to use an existing id from your own system, we recommend prefixing it with your system's name to avoid collisions, for example, mySystem-annotation001. The external id can later be used to identify the report as an alternative to the generated [UUID](https://developer.atlassian.com/bitbucket/api/2/reference/meta/uri-uuid#uuid). You can upload up to 100 annotations per POST request.  ### Sample cURL request: ``` curl --location 'https://api.bitbucket.org/2.0/repositories/<username>/<reposity-name>/commit/<commit-hash>/reports/mysystem-001/annotations' \\ --header 'Content-Type: application/json' \\ --data-raw '[   {         \"external_id\": \"mysystem-annotation001\",         \"title\": \"Security scan report\",         \"annotation_type\": \"VULNERABILITY\",         \"summary\": \"This line represents a security threat.\",         \"severity\": \"HIGH\",       \"path\": \"my-service/src/main/java/com/myCompany/mysystem/logic/Main.java\",         \"line\": 42   },   {         \"external_id\": \"mySystem-annotation002\",         \"title\": \"Bug report\",         \"annotation_type\": \"BUG\",         \"result\": \"FAILED\",         \"summary\": \"This line might introduce a bug.\",         \"severity\": \"MEDIUM\",       \"path\": \"my-service/src/main/java/com/myCompany/mysystem/logic/Helper.java\",         \"line\": 13   } ]' ```  ### Possible field values: annotation_type: VULNERABILITY, CODE_SMELL, BUG result: PASSED, FAILED, IGNORED, SKIPPED severity: HIGH, MEDIUM, LOW, CRITICAL  Please refer to the [Code Insights documentation](https://confluence.atlassian.com/bitbucket/code-insights-994316785.html) for more information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ReportAnnotation**](report_annotation.md)| The annotations to create or update | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit for which to retrieve reports. | 
  **reportId** | **string**| Uuid or external-if of the report for which to get annotations for. | 

### Return type

[**[]ReportAnnotation**](report_annotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateAnnotation**
> ReportAnnotation CreateOrUpdateAnnotation(ctx, body, workspace, repoSlug, commit, reportId, annotationId)
Create or update an annotation

Creates or updates an individual annotation for the specified report. Annotations are individual findings that have been identified as part of a report, for example, a line of code that represents a vulnerability. These annotations can be attached to a specific file and even a specific line in that file, however, that is optional. Annotations are not mandatory and a report can contain up to 1000 annotations.  Just as reports, annotation needs to be uploaded with a unique ID that can later be used to identify the report as an alternative to the generated [UUID](https://developer.atlassian.com/bitbucket/api/2/reference/meta/uri-uuid#uuid). If you want to use an existing id from your own system, we recommend prefixing it with your system's name to avoid collisions, for example, mySystem-annotation001.  ### Sample cURL request: ``` curl --request PUT 'https://api.bitbucket.org/2.0/repositories/<username>/<reposity-name>/commit/<commit-hash>/reports/mySystem-001/annotations/mysystem-annotation001' \\ --header 'Content-Type: application/json' \\ --data-raw '{     \"title\": \"Security scan report\",     \"annotation_type\": \"VULNERABILITY\",     \"summary\": \"This line represents a security thread.\",     \"severity\": \"HIGH\",     \"path\": \"my-service/src/main/java/com/myCompany/mysystem/logic/Main.java\",     \"line\": 42 }' ```  ### Possible field values: annotation_type: VULNERABILITY, CODE_SMELL, BUG result: PASSED, FAILED, IGNORED, SKIPPED severity: HIGH, MEDIUM, LOW, CRITICAL  Please refer to the [Code Insights documentation](https://confluence.atlassian.com/bitbucket/code-insights-994316785.html) for more information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReportAnnotation**](ReportAnnotation.md)| The annotation to create or update | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit the report belongs to. | 
  **reportId** | **string**| Either the uuid or external-id of the report. | 
  **annotationId** | **string**| Either the uuid or external-id of the annotation. | 

### Return type

[**ReportAnnotation**](report_annotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrUpdateReport**
> Report CreateOrUpdateReport(ctx, body, workspace, repoSlug, commit, reportId)
Create or update a report

Creates or updates a report for the specified commit. To upload a report, make sure to generate an ID that is unique across all reports for that commit. If you want to use an existing id from your own system, we recommend prefixing it with your system's name to avoid collisions, for example, mySystem-001.  ### Sample cURL request: ``` curl --request PUT 'https://api.bitbucket.org/2.0/repositories/<username>/<reposity-name>/commit/<commit-hash>/reports/mysystem-001' \\ --header 'Content-Type: application/json' \\ --data-raw '{     \"title\": \"Security scan report\",     \"details\": \"This pull request introduces 10 new dependency vulnerabilities.\",     \"report_type\": \"SECURITY\",     \"reporter\": \"mySystem\",     \"link\": \"http://www.mysystem.com/reports/001\",     \"result\": \"FAILED\",     \"data\": [         {             \"title\": \"Duration (seconds)\",             \"type\": \"DURATION\",             \"value\": 14         },         {             \"title\": \"Safe to merge?\",             \"type\": \"BOOLEAN\",             \"value\": false         }     ] }' ```  ### Possible field values: report_type: SECURITY, COVERAGE, TEST, BUG result: PASSED, FAILED, PENDING data.type: BOOLEAN, DATE, DURATION, LINK, NUMBER, PERCENTAGE, TEXT  #### Data field formats | Type  Field   | Value Field Type  | Value Field Display | |:--------------|:------------------|:--------------------| | None/ Omitted | Number, String or Boolean (not an array or object) | Plain text | | BOOLEAN | Boolean | The value will be read as a JSON boolean and displayed as 'Yes' or 'No'. | | DATE  | Number | The value will be read as a JSON number in the form of a Unix timestamp (milliseconds) and will be displayed as a relative date if the date is less than one week ago, otherwise  it will be displayed as an absolute date. | | DURATION | Number | The value will be read as a JSON number in milliseconds and will be displayed in a human readable duration format. | | LINK | Object: `{\"text\": \"Link text here\", \"href\": \"https://link.to.annotation/in/external/tool\"}` | The value will be read as a JSON object containing the fields \"text\" and \"href\" and will be displayed as a clickable link on the report. | | NUMBER | Number | The value will be read as a JSON number and large numbers will be  displayed in a human readable format (e.g. 14.3k). | | PERCENTAGE | Number (between 0 and 100) | The value will be read as a JSON number between 0 and 100 and will be displayed with a percentage sign. | | TEXT | String | The value will be read as a JSON string and will be displayed as-is |  Please refer to the [Code Insights documentation](https://confluence.atlassian.com/bitbucket/code-insights-994316785.html) for more information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Report**](Report.md)| The report to create or update | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit the report belongs to. | 
  **reportId** | **string**| Either the uuid or external-id of the report. | 

### Return type

[**Report**](report.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAnnotation**
> DeleteAnnotation(ctx, workspace, repoSlug, commit, reportId, annotationId)
Delete an annotation

Deletes a single Annotation matching the provided ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit the annotation belongs to. | 
  **reportId** | **string**| Either the uuid or external-id of the annotation. | 
  **annotationId** | **string**| Either the uuid or external-id of the annotation. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReport**
> DeleteReport(ctx, workspace, repoSlug, commit, reportId)
Delete a report

Deletes a single Report matching the provided ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit the report belongs to. | 
  **reportId** | **string**| Either the uuid or external-id of the report. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnnotation**
> ReportAnnotation GetAnnotation(ctx, workspace, repoSlug, commit, reportId, annotationId)
Get an annotation

Returns a single Annotation matching the provided ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit the report belongs to. | 
  **reportId** | **string**| Either the uuid or external-id of the report. | 
  **annotationId** | **string**| Either the uuid or external-id of the annotation. | 

### Return type

[**ReportAnnotation**](report_annotation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnnotationsForReport**
> PaginatedAnnotations GetAnnotationsForReport(ctx, workspace, repoSlug, commit, reportId)
List annotations

Returns a paginated list of Annotations for a specified report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit for which to retrieve reports. | 
  **reportId** | **string**| Uuid or external-if of the report for which to get annotations for. | 

### Return type

[**PaginatedAnnotations**](paginated_annotations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReport**
> Report GetReport(ctx, workspace, repoSlug, commit, reportId)
Get a report

Returns a single Report matching the provided ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit the report belongs to. | 
  **reportId** | **string**| Either the uuid or external-id of the report. | 

### Return type

[**Report**](report.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsForCommit**
> PaginatedReports GetReportsForCommit(ctx, workspace, repoSlug, commit)
List reports

Returns a paginated list of Reports linked to this commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example &#x60;{workspace UUID}&#x60;. | 
  **repoSlug** | **string**| The repository. | 
  **commit** | **string**| The commit for which to retrieve reports. | 

### Return type

[**PaginatedReports**](paginated_reports.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete**
> RepositoriesWorkspaceRepoSlugCommitCommitApproveDelete(ctx, commit, repoSlug, workspace)
Unapprove a commit

Redact the authenticated user's approval of the specified commit.  This operation is only available to users that have explicit access to the repository. In contrast, just the fact that a repository is publicly accessible to users does not give them the ability to approve commits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitApprovePost**
> Participant RepositoriesWorkspaceRepoSlugCommitCommitApprovePost(ctx, commit, repoSlug, workspace)
Approve a commit

Approve the specified commit as the authenticated user.  This operation is only available to users that have explicit access to the repository. In contrast, just the fact that a repository is publicly accessible to users does not give them the ability to approve commits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Participant**](participant.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet**
> CommitComment RepositoriesWorkspaceRepoSlugCommitCommitCommentsCommentIdGet(ctx, commentId, commit, repoSlug, workspace)
Get a commit comment

Returns the specified commit comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **int32**| The id of the comment. | 
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**CommitComment**](commit_comment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet**
> PaginatedCommitComments RepositoriesWorkspaceRepoSlugCommitCommitCommentsGet(ctx, commit, repoSlug, workspace, optional)
List a commit's comments

Returns the commit's comments.  This includes both global and inline comments.  The default sorting is oldest to newest and can be overridden with the `sort` query parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***CommitsApiRepositoriesWorkspaceRepoSlugCommitCommitCommentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiRepositoriesWorkspaceRepoSlugCommitCommitCommentsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **optional.String**| Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering).  | 
 **sort** | **optional.String**| Field by which the results should be sorted as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering).  | 

### Return type

[**PaginatedCommitComments**](paginated_commit_comments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost**
> RepositoriesWorkspaceRepoSlugCommitCommitCommentsPost(ctx, body, commit, repoSlug, workspace)
Create comment for a commit

Creates new comment on the specified commit.  To post a reply to an existing comment, include the `parent.id` field:  ``` $ curl https://api.bitbucket.org/2.0/repositories/atlassian/prlinks/commit/db9ba1e031d07a02603eae0e559a7adc010257fc/comments/ \\   -X POST -u evzijst \\   -H 'Content-Type: application/json' \\   -d '{\"content\": {\"raw\": \"One more thing!\"},        \"parent\": {\"id\": 5728901}}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CommitComment**](CommitComment.md)| The specified comment. | 
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitCommitGet**
> Commit RepositoriesWorkspaceRepoSlugCommitCommitGet(ctx, commit, repoSlug, workspace)
Get a commit

Returns the specified commit.  Example:  ``` $ curl https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/commit/f7591a1 {     \"rendered\": {         \"message\": {         \"raw\": \"Add a GEORDI_OUTPUT_DIR setting\",         \"markup\": \"markdown\",         \"html\": \"<p>Add a GEORDI_OUTPUT_DIR setting</p>\",         \"type\": \"rendered\"         }     },     \"hash\": \"f7591a13eda445d9a9167f98eb870319f4b6c2d8\",     \"repository\": {         \"name\": \"geordi\",         \"type\": \"repository\",         \"full_name\": \"bitbucket/geordi\",         \"links\": {             \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi\"             },             \"html\": {                 \"href\": \"https://bitbucket.org/bitbucket/geordi\"             },             \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B85d08b4e-571d-44e9-a507-fa476535aa98%7D?ts=1730260\"             }         },         \"uuid\": \"{85d08b4e-571d-44e9-a507-fa476535aa98}\"     },     \"links\": {         \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/commit/f7591a13eda445d9a9167f98eb870319f4b6c2d8\"         },         \"comments\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/commit/f7591a13eda445d9a9167f98eb870319f4b6c2d8/comments\"         },         \"patch\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/patch/f7591a13eda445d9a9167f98eb870319f4b6c2d8\"         },         \"html\": {             \"href\": \"https://bitbucket.org/bitbucket/geordi/commits/f7591a13eda445d9a9167f98eb870319f4b6c2d8\"         },         \"diff\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/diff/f7591a13eda445d9a9167f98eb870319f4b6c2d8\"         },         \"approve\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/commit/f7591a13eda445d9a9167f98eb870319f4b6c2d8/approve\"         },         \"statuses\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/commit/f7591a13eda445d9a9167f98eb870319f4b6c2d8/statuses\"         }     },     \"author\": {         \"raw\": \"Brodie Rao <a@b.c>\",         \"type\": \"author\",         \"user\": {             \"display_name\": \"Brodie Rao\",             \"uuid\": \"{9484702e-c663-4afd-aefb-c93a8cd31c28}\",             \"links\": {                 \"self\": {                     \"href\": \"https://api.bitbucket.org/2.0/users/%7B9484702e-c663-4afd-aefb-c93a8cd31c28%7D\"                 },                 \"html\": {                     \"href\": \"https://bitbucket.org/%7B9484702e-c663-4afd-aefb-c93a8cd31c28%7D/\"                 },                 \"avatar\": {                     \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:3aae1e05-702a-41e5-81c8-f36f29afb6ca/613070db-28b0-421f-8dba-ae8a87e2a5c7/128\"                 }             },             \"type\": \"user\",             \"nickname\": \"brodie\",             \"account_id\": \"557058:3aae1e05-702a-41e5-81c8-f36f29afb6ca\"         }     },     \"summary\": {         \"raw\": \"Add a GEORDI_OUTPUT_DIR setting\",         \"markup\": \"markdown\",         \"html\": \"<p>Add a GEORDI_OUTPUT_DIR setting</p>\",         \"type\": \"rendered\"     },     \"participants\": [],     \"parents\": [         {             \"type\": \"commit\",             \"hash\": \"f06941fec4ef6bcb0c2456927a0cf258fa4f899b\",             \"links\": {                 \"self\": {                     \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/commit/f06941fec4ef6bcb0c2456927a0cf258fa4f899b\"                 },                 \"html\": {                     \"href\": \"https://bitbucket.org/bitbucket/geordi/commits/f06941fec4ef6bcb0c2456927a0cf258fa4f899b\"                 }             }         }     ],     \"date\": \"2012-07-16T19:37:54+00:00\",     \"message\": \"Add a GEORDI_OUTPUT_DIR setting\",     \"type\": \"commit\" } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Commit**](commit.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitsGet**
> PaginatedChangeset RepositoriesWorkspaceRepoSlugCommitsGet(ctx, repoSlug, workspace)
List commits

These are the repository's commits. They are paginated and returned in reverse chronological order, similar to the output of `git log`. Like these tools, the DAG can be filtered.  #### GET /repositories/{workspace}/{repo_slug}/commits/  Returns all commits in the repo in topological order (newest commit first). All branches and tags are included (similar to `git log --all`).  #### GET /repositories/{workspace}/{repo_slug}/commits/?exclude=master  Returns all commits in the repo that are not on master (similar to `git log --all ^master`).  #### GET /repositories/{workspace}/{repo_slug}/commits/?include=foo&include=bar&exclude=fu&exclude=fubar  Returns all commits that are on refs `foo` or `bar`, but not on `fu` or `fubar` (similar to `git log foo bar ^fu ^fubar`).  An optional `path` parameter can be specified that will limit the results to commits that affect that path. `path` can either be a file or a directory. If a directory is specified, commits are returned that have modified any file in the directory tree rooted by `path`. It is important to note that if the `path` parameter is specified, the commits returned by this endpoint may no longer be a DAG, parent commits that do not modify the path will be omitted from the response.  #### GET /repositories/{workspace}/{repo_slug}/commits/?path=README.md&include=foo&include=bar&exclude=master  Returns all commits that are on refs `foo` or `bar`, but not on `master` that changed the file README.md.  #### GET /repositories/{workspace}/{repo_slug}/commits/?path=src/&include=foo&include=bar&exclude=master  Returns all commits that are on refs `foo` or `bar`, but not on `master` that changed to a file in any file in the directory src or its children.  Because the response could include a very large number of commits, it is paginated. Follow the 'next' link in the response to navigate to the next page of commits. As with other paginated resources, do not construct your own links.  When the include and exclude parameters are more than can fit in a query string, clients can use a `x-www-form-urlencoded` POST instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedChangeset**](paginated_changeset.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitsPost**
> PaginatedChangeset RepositoriesWorkspaceRepoSlugCommitsPost(ctx, repoSlug, workspace)
List commits with include/exclude

Identical to `GET /repositories/{workspace}/{repo_slug}/commits`, except that POST allows clients to place the include and exclude parameters in the request body to avoid URL length issues.  **Note that this resource does NOT support new commit creation.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedChangeset**](paginated_changeset.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitsRevisionGet**
> PaginatedChangeset RepositoriesWorkspaceRepoSlugCommitsRevisionGet(ctx, repoSlug, revision, workspace)
List commits for revision

These are the repository's commits. They are paginated and returned in reverse chronological order, similar to the output of `git log`. Like these tools, the DAG can be filtered.  #### GET /repositories/{workspace}/{repo_slug}/commits/master  Returns all commits on rev `master` (similar to `git log master`).  #### GET /repositories/{workspace}/{repo_slug}/commits/dev?include=foo&exclude=master  Returns all commits on ref `dev` or `foo`, except those that are reachable on `master` (similar to `git log dev foo ^master`).  An optional `path` parameter can be specified that will limit the results to commits that affect that path. `path` can either be a file or a directory. If a directory is specified, commits are returned that have modified any file in the directory tree rooted by `path`. It is important to note that if the `path` parameter is specified, the commits returned by this endpoint may no longer be a DAG, parent commits that do not modify the path will be omitted from the response.  #### GET /repositories/{workspace}/{repo_slug}/commits/dev?path=README.md&include=foo&include=bar&exclude=master  Returns all commits that are on refs `dev` or `foo` or `bar`, but not on `master` that changed the file README.md.  #### GET /repositories/{workspace}/{repo_slug}/commits/dev?path=src/&include=foo&exclude=master  Returns all commits that are on refs `dev` or `foo`, but not on `master` that changed to a file in any file in the directory src or its children.  Because the response could include a very large number of commits, it is paginated. Follow the 'next' link in the response to navigate to the next page of commits. As with other paginated resources, do not construct your own links.  When the include and exclude parameters are more than can fit in a query string, clients can use a `x-www-form-urlencoded` POST instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **revision** | **string**| The commit&#x27;s SHA1. | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedChangeset**](paginated_changeset.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugCommitsRevisionPost**
> PaginatedChangeset RepositoriesWorkspaceRepoSlugCommitsRevisionPost(ctx, repoSlug, revision, workspace)
List commits for revision using include/exclude

Identical to `GET /repositories/{workspace}/{repo_slug}/commits/{revision}`, except that POST allows clients to place the include and exclude parameters in the request body to avoid URL length issues.  **Note that this resource does NOT support new commit creation.**

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **revision** | **string**| The commit&#x27;s SHA1. | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedChangeset**](paginated_changeset.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDiffSpecGet**
> RepositoriesWorkspaceRepoSlugDiffSpecGet(ctx, repoSlug, spec, workspace, optional)
Compare two commits

Produces a raw git-style diff.  #### Single commit spec  If the `spec` argument to this API is a single commit, the diff is produced against the first parent of the specified commit.  #### Two commit spec  Two commits separated by `..` may be provided as the `spec`, e.g., `3a8b42..9ff173`. When two commits are provided and the `merge` query parameter is true or absent, this API produces a 3-way diff, also referred to as a merge diff. This is equivalent to merging the left branch into the right branch and then computing the diff of the merge commit against its first parent (the right branch). These diffs have the same behavior as pull requests that show the 3-way diff, such as the [Bitbucket Cloud Pull Request](https://blog.developer.atlassian.com/a-better-pull-request/). For a simple git-style diff, add `merge=false` to the query.  The two commits are interpreted as follows:  * First commit: the commit containing the changes we wish to preview * Second commit: the commit representing the state to which we want to   compare the first commit * **Note**: This is the opposite of the order used in `git diff`.  #### Comparison to patches  While similar to patches, diffs:  * Don't have a commit header (username, commit message, etc) * Support the optional `path=foo/bar.py` query param to filter   the diff to just that one file diff  #### Response  The raw diff is returned as-is, in whatever encoding the files in the repository use. It is not decoded into unicode. As such, the content-type is `text/plain`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **spec** | **string**| A commit SHA (e.g. &#x60;3a8b42&#x60;) or a commit range using double dot notation (e.g. &#x60;3a8b42..9ff173&#x60;).  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***CommitsApiRepositoriesWorkspaceRepoSlugDiffSpecGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiRepositoriesWorkspaceRepoSlugDiffSpecGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **context** | **optional.Int32**| Generate diffs with &lt;n&gt; lines of context instead of the usual three. | 
 **path** | **optional.String**| Limit the diff to a particular file (this parameter can be repeated for multiple paths). | 
 **ignoreWhitespace** | **optional.Bool**| Generate diffs that ignore whitespace. | 
 **binary** | **optional.Bool**| Generate diffs that include binary files, true if omitted. | 
 **renames** | **optional.Bool**| Whether to perform rename detection, true if omitted. | 
 **merge** | **optional.Bool**| If true, the source commit is merged into the destination commit, and then a diff from the destination to the merge result is returned. If false, a simple &#x27;two dot&#x27; diff between the source and destination is returned. True if omitted. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDiffstatSpecGet**
> PaginatedDiffstats RepositoriesWorkspaceRepoSlugDiffstatSpecGet(ctx, repoSlug, spec, workspace, optional)
Compare two commit diff stats

Produces a response in JSON format with a record for every path modified, including information on the type of the change and the number of lines added and removed.  #### Single commit spec  If the `spec` argument to this API is a single commit, the diff is produced against the first parent of the specified commit.  #### Two commit spec  Two commits separated by `..` may be provided as the `spec`, e.g., `3a8b42..9ff173`. When two commits are provided and the `merge` query parameter is true or absent, this API produces a 3-way diff, also referred to as a merge diff. This is equivalent to merging the left branch into the right branch and then computing the diff of the merge commit against its first parent (the right branch). These diffs have the same behavior as pull requests that show the 3-way diff, such as the [Bitbucket Cloud Pull Request](https://blog.developer.atlassian.com/a-better-pull-request/). For a simple git-style diff, add `merge=false` to the query.  The two commits are interpreted as follows:  * First commit: the commit containing the changes we wish to preview * Second commit: the commit representing the state to which we want to   compare the first commit * **Note**: This is the opposite of the order used in `git diff`.  #### Sample output ``` curl https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/diffstat/d222fa2..e174964 {     \"pagelen\": 500,     \"values\": [         {             \"type\": \"diffstat\",             \"status\": \"modified\",             \"lines_removed\": 1,             \"lines_added\": 2,             \"old\": {                 \"path\": \"setup.py\",                 \"escaped_path\": \"setup.py\",                 \"type\": \"commit_file\",                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/src/e1749643d655d7c7014001a6c0f58abaf42ad850/setup.py\"                     }                 }             },             \"new\": {                 \"path\": \"setup.py\",                 \"escaped_path\": \"setup.py\",                 \"type\": \"commit_file\",                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/bitbucket/geordi/src/d222fa235229c55dad20b190b0b571adf737d5a6/setup.py\"                     }                 }             }         }     ],     \"page\": 1,     \"size\": 1 } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **spec** | **string**| A commit SHA (e.g. &#x60;3a8b42&#x60;) or a commit range using double dot notation (e.g. &#x60;3a8b42..9ff173&#x60;).  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***CommitsApiRepositoriesWorkspaceRepoSlugDiffstatSpecGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommitsApiRepositoriesWorkspaceRepoSlugDiffstatSpecGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ignoreWhitespace** | **optional.Bool**| Generate diffs that ignore whitespace | 
 **merge** | **optional.Bool**| If true, the source commit is merged into the destination commit, and then a diffstat from the destination to the merge result is returned. If false, a simple &#x27;two dot&#x27; diffstat between the source and destination is returned. True if omitted. | 
 **path** | **optional.String**| Limit the diffstat to a particular file (this parameter can be repeated for multiple paths). | 
 **renames** | **optional.Bool**| Whether to perform rename detection, true if omitted. | 

### Return type

[**PaginatedDiffstats**](paginated_diffstats.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet**
> Commit RepositoriesWorkspaceRepoSlugMergeBaseRevspecGet(ctx, repoSlug, revspec, workspace)
Get the common ancestor between two commits

Returns the best common ancestor between two commits, specified in a revspec of 2 commits (e.g. 3a8b42..9ff173).  If more than one best common ancestor exists, only one will be returned. It is unspecified which will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **revspec** | **string**| A commit range using double dot notation (e.g. &#x60;3a8b42..9ff173&#x60;).  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Commit**](commit.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPatchSpecGet**
> RepositoriesWorkspaceRepoSlugPatchSpecGet(ctx, repoSlug, spec, workspace)
Get a patch for two commits

Produces a raw patch for a single commit (diffed against its first parent), or a patch-series for a revspec of 2 commits (e.g. `3a8b42..9ff173` where the first commit represents the source and the second commit the destination).  In case of the latter (diffing a revspec), a patch series is returned for the commits on the source branch (`3a8b42` and its ancestors in our example).  While similar to diffs, patches:  * Have a commit header (username, commit message, etc) * Do not support the `path=foo/bar.py` query parameter  The raw patch is returned as-is, in whatever encoding the files in the repository use. It is not decoded into unicode. As such, the content-type is `text/plain`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **spec** | **string**| A commit SHA (e.g. &#x60;3a8b42&#x60;) or a commit range using double dot notation (e.g. &#x60;3a8b42..9ff173&#x60;).  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

