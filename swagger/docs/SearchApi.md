# \SearchApi

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAccount**](SearchApi.md#SearchAccount) | **Get** /users/{selected_user}/search/code | Search for code in a user&#39;s repositories
[**SearchTeam**](SearchApi.md#SearchTeam) | **Get** /teams/{username}/search/code | Search for code in a team&#39;s repositories
[**SearchWorkspace**](SearchApi.md#SearchWorkspace) | **Get** /workspaces/{workspace}/search/code | Search for code in a workspace


# **SearchAccount**
> SearchResultPage SearchAccount(ctx, selectedUser, searchQuery, optional)
Search for code in a user's repositories

Search for code in the repositories of the specified user.  Searching across all repositories:  ``` curl 'https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code?search_query=foo' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 2,       \"content_matches\": [         {           \"lines\": [             {               \"line\": 2,               \"segments\": []             },             {               \"line\": 3,               \"segments\": [                 {                   \"text\": \"def \"                 },                 {                   \"text\": \"foo\",                   \"match\": true                 },                 {                   \"text\": \"():\"                 }               ]             },             {               \"line\": 4,               \"segments\": [                 {                   \"text\": \"    print(\\\"snek\\\")\"                 }               ]             },             {               \"line\": 5,               \"segments\": []             }           ]         }       ],       \"path_matches\": [         {           \"text\": \"src/\"         },         {           \"text\": \"foo\",           \"match\": true         },         {           \"text\": \".py\"         }       ],       \"file\": {         \"path\": \"src/foo.py\",         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         }       }     }   ] } ```  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  ``` curl 'https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code?search_query=foo+repo:demo' # results from the \"demo\" repository ```  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files (the `%2B` is a URL-encoded `+`):  ``` curl 'https://api.bitbucket.org/2.0/users/{ed08f5e1-605b-4f4a-aee4-6c97628a673e}/search/code'\\      '?search_query=foo&fields=%2Bvalues.file.commit.repository' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 1,       \"content_matches\": [...],       \"path_matches\": [...],       \"file\": {         \"commit\": {           \"type\": \"commit\",           \"hash\": \"ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\",           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             },             \"html\": {               \"href\": \"https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             }           },           \"repository\": {             \"name\": \"demo\",             \"type\": \"repository\",             \"full_name\": \"my-workspace/demo\",             \"links\": {               \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo\"               },               \"html\": {                 \"href\": \"https://bitbucket.org/my-workspace/demo\"               },               \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts=default\"               }             },             \"uuid\": \"{850e1749-781a-4115-9316-df39d0600e7a}\"           }         },         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         },         \"path\": \"src/foo.py\"       }     }   ] } ```  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectedUser** | **string**| Either the UUID of the account surrounded by curly-braces, for example &#x60;{account UUID}&#x60;, OR an Atlassian Account ID. | 
  **searchQuery** | **string**| The search query | 
 **optional** | ***SearchApiSearchAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page of the search results to retrieve | [default to 1]
 **pagelen** | **optional.Int32**| How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](search_result_page.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTeam**
> SearchResultPage SearchTeam(ctx, username, searchQuery, optional)
Search for code in a team's repositories

Search for code in the repositories of the specified team.  Searching across all repositories:  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query=foo' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 2,       \"content_matches\": [         {           \"lines\": [             {               \"line\": 2,               \"segments\": []             },             {               \"line\": 3,               \"segments\": [                 {                   \"text\": \"def \"                 },                 {                   \"text\": \"foo\",                   \"match\": true                 },                 {                   \"text\": \"():\"                 }               ]             },             {               \"line\": 4,               \"segments\": [                 {                   \"text\": \"    print(\\\"snek\\\")\"                 }               ]             },             {               \"line\": 5,               \"segments\": []             }           ]         }       ],       \"path_matches\": [         {           \"text\": \"src/\"         },         {           \"text\": \"foo\",           \"match\": true         },         {           \"text\": \".py\"         }       ],       \"file\": {         \"path\": \"src/foo.py\",         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         }       }     }   ] } ```  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code?search_query=foo+repo:demo' # results from the \"demo\" repository ```  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files (the `%2B` is a URL-encoded `+`):  ``` curl 'https://api.bitbucket.org/2.0/teams/team_name/search/code'\\      '?search_query=foo&fields=%2Bvalues.file.commit.repository' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 1,       \"content_matches\": [...],       \"path_matches\": [...],       \"file\": {         \"commit\": {           \"type\": \"commit\",           \"hash\": \"ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\",           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             },             \"html\": {               \"href\": \"https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             }           },           \"repository\": {             \"name\": \"demo\",             \"type\": \"repository\",             \"full_name\": \"my-workspace/demo\",             \"links\": {               \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo\"               },               \"html\": {                 \"href\": \"https://bitbucket.org/my-workspace/demo\"               },               \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts=default\"               }             },             \"uuid\": \"{850e1749-781a-4115-9316-df39d0600e7a}\"           }         },         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         },         \"path\": \"src/foo.py\"       }     }   ] } ```  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| The account to search in; either the username or the UUID in curly braces | 
  **searchQuery** | **string**| The search query | 
 **optional** | ***SearchApiSearchTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchTeamOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page of the search results to retrieve | [default to 1]
 **pagelen** | **optional.Int32**| How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](search_result_page.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchWorkspace**
> SearchResultPage SearchWorkspace(ctx, workspace, searchQuery, optional)
Search for code in a workspace

Search for code in the repositories of the specified workspace.  Searching across all repositories:  ``` curl 'https://api.bitbucket.org/2.0/workspaces/workspace_slug_or_uuid/search/code?search_query=foo' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 2,       \"content_matches\": [         {           \"lines\": [             {               \"line\": 2,               \"segments\": []             },             {               \"line\": 3,               \"segments\": [                 {                   \"text\": \"def \"                 },                 {                   \"text\": \"foo\",                   \"match\": true                 },                 {                   \"text\": \"():\"                 }               ]             },             {               \"line\": 4,               \"segments\": [                 {                   \"text\": \"    print(\\\"snek\\\")\"                 }               ]             },             {               \"line\": 5,               \"segments\": []             }           ]         }       ],       \"path_matches\": [         {           \"text\": \"src/\"         },         {           \"text\": \"foo\",           \"match\": true         },         {           \"text\": \".py\"         }       ],       \"file\": {         \"path\": \"src/foo.py\",         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         }       }     }   ] } ```  Note that searches can match in the file's text (`content_matches`), the path (`path_matches`), or both as in the example above.  You can use the same syntax for the search query as in the UI, e.g. to only search within a specific repository:  ``` curl 'https://api.bitbucket.org/2.0/workspaces/my-workspace/search/code?search_query=foo+repo:demo' # results from the \"demo\" repository ```  Similar to other APIs, you can request more fields using a `fields` query parameter. E.g. to get some more information about the repository of matched files (the `%2B` is a URL-encoded `+`):  ``` curl 'https://api.bitbucket.org/2.0/workspaces/my-workspace/search/code'\\      '?search_query=foo&fields=%2Bvalues.file.commit.repository' {   \"size\": 1,   \"page\": 1,   \"pagelen\": 10,   \"query_substituted\": false,   \"values\": [     {       \"type\": \"code_search_result\",       \"content_match_count\": 1,       \"content_matches\": [...],       \"path_matches\": [...],       \"file\": {         \"commit\": {           \"type\": \"commit\",           \"hash\": \"ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\",           \"links\": {             \"self\": {               \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/commit/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             },             \"html\": {               \"href\": \"https://bitbucket.org/my-workspace/demo/commits/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b\"             }           },           \"repository\": {             \"name\": \"demo\",             \"type\": \"repository\",             \"full_name\": \"my-workspace/demo\",             \"links\": {               \"self\": {                 \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo\"               },               \"html\": {                 \"href\": \"https://bitbucket.org/my-workspace/demo\"               },               \"avatar\": {                 \"href\": \"https://bytebucket.org/ravatar/%7B850e1749-781a-4115-9316-df39d0600e7a%7D?ts=default\"               }             },             \"uuid\": \"{850e1749-781a-4115-9316-df39d0600e7a}\"           }         },         \"type\": \"commit_file\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/my-workspace/demo/src/ad6964b5fe2880dbd9ddcad1c89000f1dbcbc24b/src/foo.py\"           }         },         \"path\": \"src/foo.py\"       }     }   ] } ```  Try `fields=%2Bvalues.*.*.*.*` to get an idea what's possible. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workspace** | **string**| The workspace to search in; either the slug or the UUID in curly braces | 
  **searchQuery** | **string**| The search query | 
 **optional** | ***SearchApiSearchWorkspaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchWorkspaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page of the search results to retrieve | [default to 1]
 **pagelen** | **optional.Int32**| How many search results to retrieve per page | [default to 10]

### Return type

[**SearchResultPage**](search_result_page.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

