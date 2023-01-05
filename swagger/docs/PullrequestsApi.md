# \PullrequestsApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPullrequestsForCommit**](PullrequestsApi.md#GetPullrequestsForCommit) | **Get** /repositories/{workspace}/{repo_slug}/commit/{commit}/pullrequests | List pull requests that contain a commit
[**PullrequestsSelectedUserGet**](PullrequestsApi.md#PullrequestsSelectedUserGet) | **Get** /pullrequests/{selected_user} | List pull requests for a user
[**RepositoriesWorkspaceRepoSlugDefaultReviewersGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugDefaultReviewersGet) | **Get** /repositories/{workspace}/{repo_slug}/default-reviewers | List default reviewers
[**RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete) | **Delete** /repositories/{workspace}/{repo_slug}/default-reviewers/{target_username} | Remove a user from the default reviewers
[**RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet) | **Get** /repositories/{workspace}/{repo_slug}/default-reviewers/{target_username} | Get a default reviewer
[**RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut) | **Put** /repositories/{workspace}/{repo_slug}/default-reviewers/{target_username} | Add a user to the default reviewers
[**RepositoriesWorkspaceRepoSlugEffectiveDefaultReviewersGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugEffectiveDefaultReviewersGet) | **Get** /repositories/{workspace}/{repo_slug}/effective-default-reviewers | List effective default reviewers
[**RepositoriesWorkspaceRepoSlugPullrequestsActivityGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsActivityGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/activity | List a pull request activity log
[**RepositoriesWorkspaceRepoSlugPullrequestsGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests | List pull requests
[**RepositoriesWorkspaceRepoSlugPullrequestsPost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests | Create a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/activity | List a pull request activity log
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/approve | Unapprove a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/approve | Approve a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments/{comment_id} | Delete a comment on a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments/{comment_id} | Get a comment on a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut) | **Put** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments/{comment_id} | Update a comment on a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments | List comments on a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/comments | Create a comment on a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/commits | List commits on a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/decline | Decline a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/diff | List changes in a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/diffstat | Get the diff stat for a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id} | Get a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/merge | Merge a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/merge/task-status/{task_id} | Get the merge task status for a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/patch | Get the patch for a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut) | **Put** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id} | Update a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete) | **Delete** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/request-changes | Remove change request for a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost) | **Post** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/request-changes | Request changes for a pull request
[**RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet**](PullrequestsApi.md#RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet) | **Get** /repositories/{workspace}/{repo_slug}/pullrequests/{pull_request_id}/statuses | List commit statuses for a pull request


# **GetPullrequestsForCommit**
> PaginatedPullrequests GetPullrequestsForCommit(ctx, workspace, repoSlug, commit, optional)
List pull requests that contain a commit

Returns a paginated list of all pull requests as part of which this commit was reviewed. Pull Request Commit Links app must be installed first before using this API; installation automatically occurs when 'Go to pull request' is clicked from the web interface for a commit's details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces | 
  **repoSlug** | **string**| The repository; either the UUID in curly braces, or the slug | 
  **commit** | **string**| The SHA1 of the commit | 
 **optional** | ***PullrequestsApiGetPullrequestsForCommitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullrequestsApiGetPullrequestsForCommitOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int32**| Which page to retrieve | [default to 1]
 **pagelen** | **optional.Int32**| How many pull requests to retrieve per page | [default to 30]

### Return type

[**PaginatedPullrequests**](paginated_pullrequests.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PullrequestsSelectedUserGet**
> PaginatedPullrequests PullrequestsSelectedUserGet(ctx, selectedUser, optional)
List pull requests for a user

Returns all pull requests authored by the specified user.  By default only open pull requests are returned. This can be controlled using the `state` query parameter. To retrieve pull requests that are in one of multiple states, repeat the `state` parameter for each individual state.  This endpoint also supports filtering and sorting of the results. See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| This can either be the username of the pull request author, the author&#39;s UUID surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;, or the author&#39;s Atlassian ID.  | 
 **optional** | ***PullrequestsApiPullrequestsSelectedUserGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullrequestsApiPullrequestsSelectedUserGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**| Only return pull requests that are in this state. This parameter can be repeated. | 

### Return type

[**PaginatedPullrequests**](paginated_pullrequests.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDefaultReviewersGet**
> PaginatedAccounts RepositoriesWorkspaceRepoSlugDefaultReviewersGet(ctx, repoSlug, workspace)
List default reviewers

Returns the repository's default reviewers.  These are the users that are automatically added as reviewers on every new pull request that is created. To obtain the repository's default reviewers as well as the default reviewers inherited from the project, use the [effective-default-reveiwers](#api-repositories-workspace-repo-slug-effective-default-reviewers-get) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedAccounts**](paginated_accounts.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete**
> RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameDelete(ctx, repoSlug, targetUsername, workspace)
Remove a user from the default reviewers

Removes a default reviewer from the repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **targetUsername** | **string**| This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet**
> Account RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernameGet(ctx, repoSlug, targetUsername, workspace)
Get a default reviewer

Returns the specified reviewer.  This can be used to test whether a user is among the repository's default reviewers list. A 404 indicates that that specified user is not a default reviewer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **targetUsername** | **string**| This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Account**](account.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut**
> Account RepositoriesWorkspaceRepoSlugDefaultReviewersTargetUsernamePut(ctx, repoSlug, targetUsername, workspace)
Add a user to the default reviewers

Adds the specified user to the repository's list of default reviewers.  This method is idempotent. Adding a user a second time has no effect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **targetUsername** | **string**| This can either be the username or the UUID of the default reviewer, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Account**](account.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugEffectiveDefaultReviewersGet**
> PaginatedDefaultReviewerAndType RepositoriesWorkspaceRepoSlugEffectiveDefaultReviewersGet(ctx, repoSlug, workspace)
List effective default reviewers

Returns the repository's effective default reviewers. This includes both default reviewers defined at the repository level as well as those inherited from its project.  These are the users that are automatically added as reviewers on every new pull request that is created.  ``` $ curl https://api.bitbucket.org/2.0/repositories/{workspace_slug}/{repo_slug}/effective-default-reviewers?page=1&pagelen=20 {     \"pagelen\": 20,     \"values\": [         {             \"user\": {                 \"display_name\": \"Patrick Wolf\",                 \"uuid\": \"{9565301a-a3cf-4b5d-88f4-dd6af8078d7e}\"             },             \"reviewer_type\": \"project\",             \"type\": \"default_reviewer\",         },         {             \"user\": {                 \"display_name\": \"Davis Lee\",                 \"uuid\": \"{f0e0e8e9-66c1-4b85-a784-44a9eb9ef1a6}\"             },             \"reviewer_type\": \"repository\",             \"type\": \"default_reviewer\",         }     ],     \"page\": 1,     \"size\": 2 } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedDefaultReviewerAndType**](paginated_default_reviewer_and_type.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsActivityGet**
> RepositoriesWorkspaceRepoSlugPullrequestsActivityGet(ctx, repoSlug, workspace)
List a pull request activity log

Returns a paginated list of the pull request's activity log.  This handler serves both a v20 and internal endpoint. The v20 endpoint returns reviewer comments, updates, approvals and request changes. The internal endpoint includes those plus tasks and attachments.  Comments created on a file or a line of code have an inline property.  Comment example: ``` {     \"pagelen\": 20,     \"values\": [         {             \"comment\": {                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695/comments/118571088\"                     },                     \"html\": {                         \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695/_/diff#comment-118571088\"                     }                 },                 \"deleted\": false,                 \"pullrequest\": {                     \"type\": \"pullrequest\",                     \"id\": 5695,                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                         }                     },                     \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"                 },                 \"content\": {                     \"raw\": \"inline with to a dn from lines\",                     \"markup\": \"markdown\",                     \"html\": \"<p>inline with to a dn from lines</p>\",                     \"type\": \"rendered\"                 },                 \"created_on\": \"2019-09-27T00:33:46.039178+00:00\",                 \"user\": {                     \"display_name\": \"Name Lastname\",                     \"uuid\": \"{}\",                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/users/%7B%7D\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/%7B%7D/\"                         },                         \"avatar\": {                             \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/:/128\"                         }                     },                     \"type\": \"user\",                     \"nickname\": \"Name\",                     \"account_id\": \"\"                 },                 \"created_on\": \"2019-09-27T00:33:46.039178+00:00\",                 \"user\": {                     \"display_name\": \"Name Lastname\",                     \"uuid\": \"{}\",                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/users/%7B%7D\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/%7B%7D/\"                         },                         \"avatar\": {                             \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/:/128\"                         }                     },                     \"type\": \"user\",                     \"nickname\": \"Name\",                     \"account_id\": \"\"                 },                 \"updated_on\": \"2019-09-27T00:33:46.055384+00:00\",                 \"inline\": {                     \"context_lines\": \"\",                     \"to\": null,                     \"path\": \"\",                     \"outdated\": false,                     \"from\": 211                 },                 \"type\": \"pullrequest_comment\",                 \"id\": 118571088             },             \"pull_request\": {                 \"type\": \"pullrequest\",                 \"id\": 5695,                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                     },                     \"html\": {                         \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                     }                 },                 \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"             }         }     ] } ```  Updates include a state property of OPEN, MERGED, or DECLINED.  Update example: ``` {     \"pagelen\": 20,     \"values\": [         {             \"update\": {                 \"description\": \"\",                 \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\",                 \"destination\": {                     \"commit\": {                         \"type\": \"commit\",                         \"hash\": \"6a2c16e4a152\",                         \"links\": {                             \"self\": {                                 \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/commit/6a2c16e4a152\"                             },                             \"html\": {                                 \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/commits/6a2c16e4a152\"                             }                         }                     },                     \"branch\": {                         \"name\": \"master\"                     },                     \"repository\": {                         \"name\": \"Atlaskit-MK-2\",                         \"type\": \"repository\",                         \"full_name\": \"atlassian/atlaskit-mk-2\",                         \"links\": {                             \"self\": {                                 \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2\"                             },                             \"html\": {                                 \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2\"                             },                             \"avatar\": {                                 \"href\": \"https://bytebucket.org/ravatar/%7B%7D?ts=js\"                             }                         },                         \"uuid\": \"{}\"                     }                 },                 \"reason\": \"\",                 \"source\": {                     \"commit\": {                         \"type\": \"commit\",                         \"hash\": \"728c8bad1813\",                         \"links\": {                             \"self\": {                                 \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/commit/728c8bad1813\"                             },                             \"html\": {                                 \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/commits/728c8bad1813\"                             }                         }                     },                     \"branch\": {                         \"name\": \"username/NONE-add-onClick-prop-for-accessibility\"                     },                     \"repository\": {                         \"name\": \"Atlaskit-MK-2\",                         \"type\": \"repository\",                         \"full_name\": \"atlassian/atlaskit-mk-2\",                         \"links\": {                             \"self\": {                                 \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2\"                             },                             \"html\": {                                 \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2\"                             },                             \"avatar\": {                                 \"href\": \"https://bytebucket.org/ravatar/%7B%7D?ts=js\"                             }                         },                         \"uuid\": \"{}\"                     }                 },                 \"state\": \"OPEN\",                 \"author\": {                     \"display_name\": \"Name Lastname\",                     \"uuid\": \"{}\",                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/users/%7B%7D\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/%7B%7D/\"                         },                         \"avatar\": {                             \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/:/128\"                         }                     },                     \"type\": \"user\",                     \"nickname\": \"Name\",                     \"account_id\": \"\"                 },                 \"date\": \"2019-05-10T06:48:25.305565+00:00\"             },             \"pull_request\": {                 \"type\": \"pullrequest\",                 \"id\": 5695,                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                     },                     \"html\": {                         \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                     }                 },                 \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"             }         }     ] } ```  Approval example: ``` {     \"pagelen\": 20,     \"values\": [         {             \"approval\": {                 \"date\": \"2019-09-27T00:37:19.849534+00:00\",                 \"pullrequest\": {                     \"type\": \"pullrequest\",                     \"id\": 5695,                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                         }                     },                     \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"                 },                 \"user\": {                     \"display_name\": \"Name Lastname\",                     \"uuid\": \"{}\",                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/users/%7B%7D\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/%7B%7D/\"                         },                         \"avatar\": {                             \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/:/128\"                         }                     },                     \"type\": \"user\",                     \"nickname\": \"Name\",                     \"account_id\": \"\"                 }             },             \"pull_request\": {                 \"type\": \"pullrequest\",                 \"id\": 5695,                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                     },                     \"html\": {                         \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                     }                 },                 \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"             }         }     ] } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **RepositoriesWorkspaceRepoSlugPullrequestsGet**
> PaginatedPullrequests RepositoriesWorkspaceRepoSlugPullrequestsGet(ctx, repoSlug, workspace, optional)
List pull requests

Returns all pull requests on the specified repository.  By default only open pull requests are returned. This can be controlled using the `state` query parameter. To retrieve pull requests that are in one of multiple states, repeat the `state` parameter for each individual state.  This endpoint also supports filtering and sorting of the results. See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **optional.String**| Only return pull requests that are in this state. This parameter can be repeated. | 

### Return type

[**PaginatedPullrequests**](paginated_pullrequests.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPost**
> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPost(ctx, repoSlug, workspace, optional)
Create a pull request

Creates a new pull request where the destination repository is this repository and the author is the authenticated user.  The minimum required fields to create a pull request are `title` and `source`, specified by a branch name.  ``` curl https://api.bitbucket.org/2.0/repositories/my-workspace/my-repository/pullrequests \\     -u my-username:my-password \\     --request POST \\     --header 'Content-Type: application/json' \\     --data '{         \"title\": \"My Title\",         \"source\": {             \"branch\": {                 \"name\": \"staging\"             }         }     }' ```  If the pull request's `destination` is not specified, it will default to the `repository.mainbranch`. To open a pull request to a different branch, say from a feature branch to a staging branch, specify a `destination` (same format as the `source`):  ``` {     \"title\": \"My Title\",     \"source\": {         \"branch\": {             \"name\": \"my-feature-branch\"         }     },     \"destination\": {         \"branch\": {             \"name\": \"staging\"         }     } } ```  Reviewers can be specified by adding an array of user objects as the `reviewers` property.  ``` {     \"title\": \"My Title\",     \"source\": {         \"branch\": {             \"name\": \"my-feature-branch\"         }     },     \"reviewers\": [         {             \"uuid\": \"{504c3b62-8120-4f0c-a7bc-87800b9d6f70}\"         }     ] } ```  Other fields:  * `description` - a string * `close_source_branch` - boolean that specifies if the source branch should be closed upon merging

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Pullrequest**](Pullrequest.md)| The new pull request.  The request URL you POST to becomes the destination repository URL. For this reason, you must specify an explicit source repository in the request object if you want to pull from a different repository (fork).  Since not all elements are required or even mutable, you only need to include the elements you want to initialize, such as the source branch and the title. | 

### Return type

[**Pullrequest**](pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdActivityGet(ctx, pullRequestId, repoSlug, workspace)
List a pull request activity log

Returns a paginated list of the pull request's activity log.  This handler serves both a v20 and internal endpoint. The v20 endpoint returns reviewer comments, updates, approvals and request changes. The internal endpoint includes those plus tasks and attachments.  Comments created on a file or a line of code have an inline property.  Comment example: ``` {     \"pagelen\": 20,     \"values\": [         {             \"comment\": {                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695/comments/118571088\"                     },                     \"html\": {                         \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695/_/diff#comment-118571088\"                     }                 },                 \"deleted\": false,                 \"pullrequest\": {                     \"type\": \"pullrequest\",                     \"id\": 5695,                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                         }                     },                     \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"                 },                 \"content\": {                     \"raw\": \"inline with to a dn from lines\",                     \"markup\": \"markdown\",                     \"html\": \"<p>inline with to a dn from lines</p>\",                     \"type\": \"rendered\"                 },                 \"created_on\": \"2019-09-27T00:33:46.039178+00:00\",                 \"user\": {                     \"display_name\": \"Name Lastname\",                     \"uuid\": \"{}\",                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/users/%7B%7D\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/%7B%7D/\"                         },                         \"avatar\": {                             \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/:/128\"                         }                     },                     \"type\": \"user\",                     \"nickname\": \"Name\",                     \"account_id\": \"\"                 },                 \"created_on\": \"2019-09-27T00:33:46.039178+00:00\",                 \"user\": {                     \"display_name\": \"Name Lastname\",                     \"uuid\": \"{}\",                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/users/%7B%7D\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/%7B%7D/\"                         },                         \"avatar\": {                             \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/:/128\"                         }                     },                     \"type\": \"user\",                     \"nickname\": \"Name\",                     \"account_id\": \"\"                 },                 \"updated_on\": \"2019-09-27T00:33:46.055384+00:00\",                 \"inline\": {                     \"context_lines\": \"\",                     \"to\": null,                     \"path\": \"\",                     \"outdated\": false,                     \"from\": 211                 },                 \"type\": \"pullrequest_comment\",                 \"id\": 118571088             },             \"pull_request\": {                 \"type\": \"pullrequest\",                 \"id\": 5695,                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                     },                     \"html\": {                         \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                     }                 },                 \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"             }         }     ] } ```  Updates include a state property of OPEN, MERGED, or DECLINED.  Update example: ``` {     \"pagelen\": 20,     \"values\": [         {             \"update\": {                 \"description\": \"\",                 \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\",                 \"destination\": {                     \"commit\": {                         \"type\": \"commit\",                         \"hash\": \"6a2c16e4a152\",                         \"links\": {                             \"self\": {                                 \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/commit/6a2c16e4a152\"                             },                             \"html\": {                                 \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/commits/6a2c16e4a152\"                             }                         }                     },                     \"branch\": {                         \"name\": \"master\"                     },                     \"repository\": {                         \"name\": \"Atlaskit-MK-2\",                         \"type\": \"repository\",                         \"full_name\": \"atlassian/atlaskit-mk-2\",                         \"links\": {                             \"self\": {                                 \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2\"                             },                             \"html\": {                                 \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2\"                             },                             \"avatar\": {                                 \"href\": \"https://bytebucket.org/ravatar/%7B%7D?ts=js\"                             }                         },                         \"uuid\": \"{}\"                     }                 },                 \"reason\": \"\",                 \"source\": {                     \"commit\": {                         \"type\": \"commit\",                         \"hash\": \"728c8bad1813\",                         \"links\": {                             \"self\": {                                 \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/commit/728c8bad1813\"                             },                             \"html\": {                                 \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/commits/728c8bad1813\"                             }                         }                     },                     \"branch\": {                         \"name\": \"username/NONE-add-onClick-prop-for-accessibility\"                     },                     \"repository\": {                         \"name\": \"Atlaskit-MK-2\",                         \"type\": \"repository\",                         \"full_name\": \"atlassian/atlaskit-mk-2\",                         \"links\": {                             \"self\": {                                 \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2\"                             },                             \"html\": {                                 \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2\"                             },                             \"avatar\": {                                 \"href\": \"https://bytebucket.org/ravatar/%7B%7D?ts=js\"                             }                         },                         \"uuid\": \"{}\"                     }                 },                 \"state\": \"OPEN\",                 \"author\": {                     \"display_name\": \"Name Lastname\",                     \"uuid\": \"{}\",                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/users/%7B%7D\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/%7B%7D/\"                         },                         \"avatar\": {                             \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/:/128\"                         }                     },                     \"type\": \"user\",                     \"nickname\": \"Name\",                     \"account_id\": \"\"                 },                 \"date\": \"2019-05-10T06:48:25.305565+00:00\"             },             \"pull_request\": {                 \"type\": \"pullrequest\",                 \"id\": 5695,                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                     },                     \"html\": {                         \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                     }                 },                 \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"             }         }     ] } ```  Approval example: ``` {     \"pagelen\": 20,     \"values\": [         {             \"approval\": {                 \"date\": \"2019-09-27T00:37:19.849534+00:00\",                 \"pullrequest\": {                     \"type\": \"pullrequest\",                     \"id\": 5695,                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                         }                     },                     \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"                 },                 \"user\": {                     \"display_name\": \"Name Lastname\",                     \"uuid\": \"{}\",                     \"links\": {                         \"self\": {                             \"href\": \"https://api.bitbucket.org/2.0/users/%7B%7D\"                         },                         \"html\": {                             \"href\": \"https://bitbucket.org/%7B%7D/\"                         },                         \"avatar\": {                             \"href\": \"https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/:/128\"                         }                     },                     \"type\": \"user\",                     \"nickname\": \"Name\",                     \"account_id\": \"\"                 }             },             \"pull_request\": {                 \"type\": \"pullrequest\",                 \"id\": 5695,                 \"links\": {                     \"self\": {                         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/atlaskit-mk-2/pullrequests/5695\"                     },                     \"html\": {                         \"href\": \"https://bitbucket.org/atlassian/atlaskit-mk-2/pull-requests/5695\"                     }                 },                 \"title\": \"username/NONE: small change from onFocus to onClick to handle tabbing through the page and not expand the editor unless a click event triggers it\"             }         }     ] } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
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

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveDelete(ctx, pullRequestId, repoSlug, workspace)
Unapprove a pull request

Redact the authenticated user's approval of the specified pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
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

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost**
> Participant RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApprovePost(ctx, pullRequestId, repoSlug, workspace)
Approve a pull request

Approve the specified pull request as the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Participant**](participant.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdDelete(ctx, commentId, pullRequestId, repoSlug, workspace)
Delete a comment on a pull request

Deletes a specific pull request comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **int32**| The id of the comment. | 
  **pullRequestId** | **int32**| The id of the pull request. | 
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

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet**
> PullrequestComment RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdGet(ctx, commentId, pullRequestId, repoSlug, workspace)
Get a comment on a pull request

Returns a specific pull request comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **int32**| The id of the comment. | 
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PullrequestComment**](pullrequest_comment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut**
> PullrequestComment RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdPut(ctx, commentId, pullRequestId, repoSlug, workspace, body)
Update a comment on a pull request

Updates a specific pull request comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **int32**| The id of the comment. | 
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
  **body** | [**PullrequestComment**](PullrequestComment.md)| The contents of the updated comment. | 

### Return type

[**PullrequestComment**](pullrequest_comment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet**
> PaginatedPullrequestComments RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsGet(ctx, pullRequestId, repoSlug, workspace)
List comments on a pull request

Returns a paginated list of the pull request's comments.  This includes both global, inline comments and replies.  The default sorting is oldest to newest and can be overridden with the `sort` query parameter.  This endpoint also supports filtering and sorting of the results. See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**PaginatedPullrequestComments**](paginated_pullrequest_comments.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost**
> PullrequestComment RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsPost(ctx, pullRequestId, repoSlug, workspace, body)
Create a comment on a pull request

Creates a new pull request comment.  Returns the newly created pull request comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
  **body** | [**PullrequestComment**](PullrequestComment.md)| The comment object. | 

### Return type

[**PullrequestComment**](pullrequest_comment.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommitsGet(ctx, pullRequestId, repoSlug, workspace)
List commits on a pull request

Returns a paginated list of the pull request's commits.  These are the commits that are being merged into the destination branch when the pull requests gets accepted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
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

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost**
> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDeclinePost(ctx, pullRequestId, repoSlug, workspace)
Decline a pull request

Declines the pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Pullrequest**](pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffGet(ctx, pullRequestId, repoSlug, workspace)
List changes in a pull request

Redirects to the [repository diff](/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diff-spec-get) with the revspec that corresponds to the pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
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

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffstatGet(ctx, pullRequestId, repoSlug, workspace)
Get the diff stat for a pull request

Redirects to the [repository diffstat](/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get) with the revspec that corresponds to the pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
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

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet**
> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdGet(ctx, pullRequestId, repoSlug, workspace)
Get a pull request

Returns the specified pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Pullrequest**](pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost**
> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePost(ctx, pullRequestId, repoSlug, workspace, optional)
Merge a pull request

Merges the pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of PullrequestMergeParameters**](PullrequestMergeParameters.md)|  | 
 **async** | **optional.Bool**| Default value is false.   When set to true, runs merge asynchronously and immediately returns a 202 with polling link to the task-status API in the Location header.   When set to false, runs merge and waits for it to complete, returning 200 when it succeeds. If the duration of the merge exceeds a timeout threshold, the API returns a 202 with polling link to the task-status API in the Location header. | 

### Return type

[**Pullrequest**](pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeTaskStatusTaskIdGet(ctx, pullRequestId, repoSlug, taskId, workspace)
Get the merge task status for a pull request

When merging a pull request takes too long, the client receives a task ID along with a 202 status code. The task ID can be used in a call to this endpoint to check the status of a merge task.  ``` curl -X GET https://api.bitbucket.org/2.0/repositories/atlassian/bitbucket/pullrequests/2286/merge/task-status/<task_id> ```  If the merge task is not yet finished, a PENDING status will be returned.  ``` HTTP/2 200 {     \"task_status\": \"PENDING\",     \"links\": {         \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bitbucket/pullrequests/2286/merge/task-status/<task_id>\"         }     } } ```  If the merge was successful, a SUCCESS status will be returned.  ``` HTTP/2 200 {     \"task_status\": \"SUCCESS\",     \"links\": {         \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bitbucket/pullrequests/2286/merge/task-status/<task_id>\"         }     },     \"merge_result\": <the merged pull request object> } ```  If the merge task failed, an error will be returned.  ``` {     \"type\": \"error\",     \"error\": {         \"message\": \"<error message>\"     } } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **taskId** | **string**| ID of the merge task | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPatchGet(ctx, pullRequestId, repoSlug, workspace)
Get the patch for a pull request

Redirects to the [repository patch](/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-patch-spec-get) with the revspec that corresponds to pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
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

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut**
> Pullrequest RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPut(ctx, pullRequestId, repoSlug, workspace, optional)
Update a pull request

Mutates the specified pull request.  This can be used to change the pull request's branches or description.  Only open pull requests can be mutated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of Pullrequest**](Pullrequest.md)| The pull request that is to be updated. | 

### Return type

[**Pullrequest**](pullrequest.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete**
> RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesDelete(ctx, pullRequestId, repoSlug, workspace)
Remove change request for a pull request



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
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

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost**
> Participant RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdRequestChangesPost(ctx, pullRequestId, repoSlug, workspace)
Request changes for a pull request



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 

### Return type

[**Participant**](participant.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet**
> PaginatedCommitstatuses RepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGet(ctx, pullRequestId, repoSlug, workspace, optional)
List commit statuses for a pull request

Returns all statuses (e.g. build results) for the given pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The id of the pull request. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PullrequestsApiRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **q** | **optional.String**| Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering).  | 
 **sort** | **optional.String**| Field by which the results should be sorted as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). Defaults to &#x60;created_on&#x60;.  | 

### Return type

[**PaginatedCommitstatuses**](paginated_commitstatuses.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

