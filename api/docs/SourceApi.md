# {{classname}}

All URIs are relative to *https://api.bitbucket.org/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet**](SourceApi.md#RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet) | **Get** /repositories/{workspace}/{repo_slug}/filehistory/{commit}/{path} | List commits that modified a file
[**RepositoriesWorkspaceRepoSlugSrcCommitPathGet**](SourceApi.md#RepositoriesWorkspaceRepoSlugSrcCommitPathGet) | **Get** /repositories/{workspace}/{repo_slug}/src/{commit}/{path} | Get file or directory contents
[**RepositoriesWorkspaceRepoSlugSrcGet**](SourceApi.md#RepositoriesWorkspaceRepoSlugSrcGet) | **Get** /repositories/{workspace}/{repo_slug}/src | Get the root directory of the main branch
[**RepositoriesWorkspaceRepoSlugSrcPost**](SourceApi.md#RepositoriesWorkspaceRepoSlugSrcPost) | **Post** /repositories/{workspace}/{repo_slug}/src | Create a commit by uploading a file

# **RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet**
> PaginatedFiles RepositoriesWorkspaceRepoSlugFilehistoryCommitPathGet(ctx, commit, path, repoSlug, workspace, optional)
List commits that modified a file

Returns a paginated list of commits that modified the specified file.  Commits are returned in reverse chronological order. This is roughly equivalent to the following commands:      $ git log --follow --date-order <sha> <path>      $ hg log --follow <path>  By default, Bitbucket will follow renames and the path name in the returned entries reflects that. This can be turned off using the `?renames=false` query parameter.  Results are returned in descending chronological order by default, and like most endpoints you can [filter and sort](/cloud/bitbucket/rest/intro/#filtering) the response to only provide exactly the data you want.  For example, if you wanted to find commits made before 2011-05-18 against a file named `README.rst`, but you only wanted the path and date, your query would look like this:  ``` $ curl 'https://api.bitbucket.org/2.0/repositories/evzijst/dogslow/filehistory/master/README.rst'\\   '?fields=values.next,values.path,values.commit.date&q=commit.date<=2011-05-18' {   \"values\": [     {       \"commit\": {         \"date\": \"2011-05-17T07:32:09+00:00\"       },       \"path\": \"README.rst\"     },     {       \"commit\": {         \"date\": \"2011-05-16T06:33:28+00:00\"       },       \"path\": \"README.txt\"     },     {       \"commit\": {         \"date\": \"2011-05-16T06:15:39+00:00\"       },       \"path\": \"README.txt\"     }   ] } ```  In the response you can see that the file was renamed to `README.rst` by the commit made on 2011-05-16, and was previously named `README.txt`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **path** | **string**| Path to the file. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***SourceApiRepositoriesWorkspaceRepoSlugFilehistoryCommitPathGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceApiRepositoriesWorkspaceRepoSlugFilehistoryCommitPathGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **renames** | **optional.String**|  When &#x60;true&#x60;, Bitbucket will follow the history of the file across renames (this is the default behavior). This can be turned off by specifying &#x60;false&#x60;. | 
 **q** | **optional.String**|  Query string to narrow down the response as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**|  Name of a response property sort the result by as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results).  | 

### Return type

[**PaginatedFiles**](paginated_files.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugSrcCommitPathGet**
> PaginatedTreeentries RepositoriesWorkspaceRepoSlugSrcCommitPathGet(ctx, commit, path, repoSlug, workspace, optional)
Get file or directory contents

This endpoints is used to retrieve the contents of a single file, or the contents of a directory at a specified revision.  #### Raw file contents  When `path` points to a file, this endpoint returns the raw contents. The response's Content-Type is derived from the filename extension (not from the contents). The file contents are not processed and no character encoding/recoding is performed and as a result no character encoding is included as part of the Content-Type.  The `Content-Disposition` header will be \"attachment\" to prevent browsers from running executable files.  If the file is managed by LFS, then a 301 redirect pointing to Atlassian's media services platform is returned.  The response includes an ETag that is based on the contents of the file and its attributes. This means that an empty `__init__.py` always returns the same ETag, regardless on the directory it lives in, or the commit it is on.  #### File meta data  When the request for a file path includes the query parameter `?format=meta`, instead of returning the file's raw contents, Bitbucket instead returns the JSON object describing the file's properties:  ```javascript $ curl https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef/tests/__init__.py?format=meta {   \"links\": {     \"self\": {       \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef5d3df01aed629f650959d6706d54cd335/tests/__init__.py\"     },     \"meta\": {       \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef5d3df01aed629f650959d6706d54cd335/tests/__init__.py?format=meta\"     }   },   \"path\": \"tests/__init__.py\",   \"commit\": {     \"type\": \"commit\",     \"hash\": \"eefd5ef5d3df01aed629f650959d6706d54cd335\",     \"links\": {       \"self\": {         \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/commit/eefd5ef5d3df01aed629f650959d6706d54cd335\"       },       \"html\": {         \"href\": \"https://bitbucket.org/atlassian/bbql/commits/eefd5ef5d3df01aed629f650959d6706d54cd335\"       }     }   },   \"attributes\": [],   \"type\": \"commit_file\",   \"size\": 0 } ```  File objects contain an `attributes` element that contains a list of possible modifiers. Currently defined values are:  * `link` -- indicates that the entry is a symbolic link. The contents     of the file represent the path the link points to. * `executable` -- indicates that the file has the executable bit set. * `subrepository` -- indicates that the entry points to a submodule or     subrepo. The contents of the file is the SHA1 of the repository     pointed to. * `binary` -- indicates whether Bitbucket thinks the file is binary.  This endpoint can provide an alternative to how a HEAD request can be used to check for the existence of a file, or a file's size without incurring the overhead of receiving its full contents.   #### Directory listings  When `path` points to a directory instead of a file, the response is a paginated list of directory and file objects in the same order as the underlying SCM system would return them.  For example:  ```javascript $ curl https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef/tests {   \"pagelen\": 10,   \"values\": [     {       \"path\": \"tests/test_project\",       \"type\": \"commit_directory\",       \"links\": {         \"self\": {           \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef5d3df01aed629f650959d6706d54cd335/tests/test_project/\"         },         \"meta\": {           \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef5d3df01aed629f650959d6706d54cd335/tests/test_project/?format=meta\"         }       },       \"commit\": {         \"type\": \"commit\",         \"hash\": \"eefd5ef5d3df01aed629f650959d6706d54cd335\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/commit/eefd5ef5d3df01aed629f650959d6706d54cd335\"           },           \"html\": {             \"href\": \"https://bitbucket.org/atlassian/bbql/commits/eefd5ef5d3df01aed629f650959d6706d54cd335\"           }         }       }     },     {       \"links\": {         \"self\": {           \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef5d3df01aed629f650959d6706d54cd335/tests/__init__.py\"         },         \"meta\": {           \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef5d3df01aed629f650959d6706d54cd335/tests/__init__.py?format=meta\"         }       },       \"path\": \"tests/__init__.py\",       \"commit\": {         \"type\": \"commit\",         \"hash\": \"eefd5ef5d3df01aed629f650959d6706d54cd335\",         \"links\": {           \"self\": {             \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/commit/eefd5ef5d3df01aed629f650959d6706d54cd335\"           },           \"html\": {             \"href\": \"https://bitbucket.org/atlassian/bbql/commits/eefd5ef5d3df01aed629f650959d6706d54cd335\"           }         }       },       \"attributes\": [],       \"type\": \"commit_file\",       \"size\": 0     }   ],   \"page\": 1,   \"size\": 2 } ```  When listing the contents of the repo's root directory, the use of a trailing slash at the end of the URL is required.  The response by default is not recursive, meaning that only the direct contents of a path are returned. The response does not recurse down into subdirectories. In order to \"walk\" the entire directory tree, the client can either parse each response and follow the `self` links of each `commit_directory` object, or can specify a `max_depth` to recurse to.  The max_depth parameter will do a breadth-first search to return the contents of the subdirectories up to the depth specified. Breadth-first search was chosen as it leads to the least amount of file system operations for git. If the `max_depth` parameter is specified to be too large, the call will time out and return a 555.  Each returned object is either a `commit_file`, or a `commit_directory`, both of which contain a `path` element. This path is the absolute path from the root of the repository. Each object also contains a `commit` object which embeds the commit the file is on. Note that this is merely the commit that was used in the URL. It is *not* the commit that last modified the file.  Directory objects have 2 representations. Their `self` link returns the paginated contents of the directory. The `meta` link on the other hand returns the actual `directory` object itself, e.g.:  ```javascript {   \"path\": \"tests/test_project\",   \"type\": \"commit_directory\",   \"links\": {     \"self\": {       \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef5d3df01aed629f650959d6706d54cd335/tests/test_project/\"     },     \"meta\": {       \"href\": \"https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src/eefd5ef5d3df01aed629f650959d6706d54cd335/tests/test_project/?format=meta\"     }   },   \"commit\": { ... } } ```  #### Querying, filtering and sorting  Like most API endpoints, this API supports the Bitbucket querying/filtering syntax and so you could filter a directory listing to only include entries that match certain criteria. For instance, to list all binary files over 1kb use the expression:  `size > 1024 and attributes = \"binary\"`  which after urlencoding yields the query string:  `?q=size%3E1024+and+attributes%3D%22binary%22`  To change the ordering of the response, use the `?sort` parameter:  `.../src/eefd5ef/?sort=-size`  See [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering) for more details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commit** | **string**| The commit&#x27;s SHA1. | 
  **path** | **string**| Path to the file. | 
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***SourceApiRepositoriesWorkspaceRepoSlugSrcCommitPathGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceApiRepositoriesWorkspaceRepoSlugSrcCommitPathGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **format** | **optional.String**| If &#x27;meta&#x27; is provided, returns the (json) meta data for the contents of the file.  If &#x27;rendered&#x27; is provided, returns the contents of a non-binary file in HTML-formatted rendered markup. Since Git does not generally track what text encoding scheme is used, this endpoint attempts to detect the most appropriate character encoding. While usually correct, determining the character encoding can be ambiguous which in exceptional cases can lead to misinterpretation of the characters. As such, the raw element in the response object should not be treated as equivalent to the file&#x27;s actual contents. | 
 **q** | **optional.String**| Optional filter expression as per [filtering and sorting](/cloud/bitbucket/rest/intro/#filtering). | 
 **sort** | **optional.String**| Optional sorting parameter as per [filtering and sorting](/cloud/bitbucket/rest/intro/#sorting-query-results). | 
 **maxDepth** | **optional.Int32**| If provided, returns the contents of the repository and its subdirectories recursively until the specified max_depth of nested directories. When omitted, this defaults to 1. | 

### Return type

[**PaginatedTreeentries**](paginated_treeentries.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugSrcGet**
> PaginatedTreeentries RepositoriesWorkspaceRepoSlugSrcGet(ctx, repoSlug, workspace, optional)
Get the root directory of the main branch

This endpoint redirects the client to the directory listing of the root directory on the main branch.  This is equivalent to directly hitting [/2.0/repositories/{username}/{repo_slug}/src/{commit}/{path}](src/%7Bcommit%7D/%7Bpath%7D) without having to know the name or SHA1 of the repo's main branch.  To create new commits, [POST to this endpoint](#post)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***SourceApiRepositoriesWorkspaceRepoSlugSrcGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceApiRepositoriesWorkspaceRepoSlugSrcGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **optional.String**| Instead of returning the file&#x27;s contents, return the (json) meta data for it. | 

### Return type

[**PaginatedTreeentries**](paginated_treeentries.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesWorkspaceRepoSlugSrcPost**
> RepositoriesWorkspaceRepoSlugSrcPost(ctx, repoSlug, workspace, optional)
Create a commit by uploading a file

This endpoint is used to create new commits in the repository by uploading files.  To add a new file to a repository:  ``` $ curl https://api.bitbucket.org/2.0/repositories/username/slug/src \\   -F /repo/path/to/image.png=@image.png ```  This will create a new commit on top of the main branch, inheriting the contents of the main branch, but adding (or overwriting) the `image.png` file to the repository in the `/repo/path/to` directory.  To create a commit that deletes files, use the `files` parameter:  ``` $ curl https://api.bitbucket.org/2.0/repositories/username/slug/src \\   -F files=/file/to/delete/1.txt \\   -F files=/file/to/delete/2.txt ```  You can add/modify/delete multiple files in a request. Rename/move a file by deleting the old path and adding the content at the new path.  This endpoint accepts `multipart/form-data` (as in the examples above), as well as `application/x-www-form-urlencoded`.  #### multipart/form-data  A `multipart/form-data` post contains a series of \"form fields\" that identify both the individual files that are being uploaded, as well as additional, optional meta data.  Files are uploaded in file form fields (those that have a `Content-Disposition` parameter) whose field names point to the remote path in the repository where the file should be stored. Path field names are always interpreted to be absolute from the root of the repository, regardless whether the client uses a leading slash (as the above `curl` example did).  File contents are treated as bytes and are not decoded as text.  The commit message, as well as other non-file meta data for the request, is sent along as normal form field elements. Meta data fields share the same namespace as the file objects. For `multipart/form-data` bodies that should not lead to any ambiguity, as the `Content-Disposition` header will contain the `filename` parameter to distinguish between a file named \"message\" and the commit message field.  #### application/x-www-form-urlencoded  It is also possible to upload new files using a simple `application/x-www-form-urlencoded` POST. This can be convenient when uploading pure text files:  ``` $ curl https://api.bitbucket.org/2.0/repositories/atlassian/bbql/src \\   --data-urlencode \"/path/to/me.txt=Lorem ipsum.\" \\   --data-urlencode \"message=Initial commit\" \\   --data-urlencode \"author=Erik van Zijst <erik.van.zijst@gmail.com>\" ```  There could be a field name clash if a client were to upload a file named \"message\", as this filename clashes with the meta data property for the commit message. To avoid this and to upload files whose names clash with the meta data properties, use a leading slash for the files, e.g. `curl --data-urlencode \"/message=file contents\"`.  When an explicit slash is omitted for a file whose path matches that of a meta data parameter, then it is interpreted as meta data, not as a file.  #### Executables and links  While this API aims to facilitate the most common use cases, it is possible to perform some more advanced operations like creating a new symlink in the repository, or creating an executable file.  Files can be supplied with a `x-attributes` value in the `Content-Disposition` header. For example, to upload an executable file, as well as create a symlink from `README.txt` to `README`:  ``` --===============1438169132528273974== Content-Type: text/plain; charset=\"us-ascii\" MIME-Version: 1.0 Content-Transfer-Encoding: 7bit Content-ID: \"bin/shutdown.sh\" Content-Disposition: attachment; filename=\"shutdown.sh\"; x-attributes:\"executable\"  #!/bin/sh halt  --===============1438169132528273974== Content-Type: text/plain; charset=\"us-ascii\" MIME-Version: 1.0 Content-Transfer-Encoding: 7bit Content-ID: \"/README.txt\" Content-Disposition: attachment; filename=\"README.txt\"; x-attributes:\"link\"  README --===============1438169132528273974==-- ```  Links are files that contain the target path and have `x-attributes:\"link\"` set.  When overwriting links with files, or vice versa, the newly uploaded file determines both the new contents, as well as the attributes. That means uploading a file without specifying `x-attributes=\"link\"` will create a regular file, even if the parent commit hosted a symlink at the same path.  The same applies to executables. When modifying an existing executable file, the form-data file element must include `x-attributes=\"executable\"` in order to preserve the executable status of the file.  Note that this API does not support the creation or manipulation of subrepos / submodules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoSlug** | **string**| This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;.  | 
  **workspace** | **string**| This can either be the workspace ID (slug) or the workspace UUID surrounded by curly-braces, for example: &#x60;{workspace UUID}&#x60;.  | 
 **optional** | ***SourceApiRepositoriesWorkspaceRepoSlugSrcPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceApiRepositoriesWorkspaceRepoSlugSrcPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **message** | **optional.String**| The commit message. When omitted, Bitbucket uses a canned string. | 
 **author** | **optional.String**|  The raw string to be used as the new commit&#x27;s author. This string follows the format &#x60;Erik van Zijst &lt;evzijst@atlassian.com&gt;&#x60;.  When omitted, Bitbucket uses the authenticated user&#x27;s full/display name and primary email address. Commits cannot be created anonymously. | 
 **parents** | **optional.String**|  A comma-separated list of SHA1s of the commits that should be the parents of the newly created commit.  When omitted, the new commit will inherit from and become a child of the main branch&#x27;s tip/HEAD commit.  When more than one SHA1 is provided, the first SHA1 identifies the commit from which the content will be inherited.\&quot;. | 
 **files** | **optional.String**|  Optional field that declares the files that the request is manipulating. When adding a new file to a repo, or when overwriting an existing file, the client can just upload the full contents of the file in a normal form field and the use of this &#x60;files&#x60; meta data field is redundant. However, when the &#x60;files&#x60; field contains a file path that does not have a corresponding, identically-named form field, then Bitbucket interprets that as the client wanting to replace the named file with the null set and the file is deleted instead.  Paths in the repo that are referenced in neither files nor an individual file field, remain unchanged and carry over from the parent to the new commit.  This API does not support renaming as an explicit feature. To rename a file, simply delete it and recreate it under the new name in the same commit.  | 
 **branch** | **optional.String**|  The name of the branch that the new commit should be created on. When omitted, the commit will be created on top of the main branch and will become the main branch&#x27;s new head.  When a branch name is provided that already exists in the repo, then the commit will be created on top of that branch. In this case, *if* a parent SHA1 was also provided, then it is asserted that the parent is the branch&#x27;s tip/HEAD at the time the request is made. When this is not the case, a 409 is returned.  When a new branch name is specified (that does not already exist in the repo), and no parent SHA1s are provided, then the new commit will inherit from the current main branch&#x27;s tip/HEAD commit, but not advance the main branch. The new commit will be the new branch. When the request *also* specifies a parent SHA1, then the new commit and branch are created directly on top of the parent commit, regardless of the state of the main branch.  When a branch name is not specified, but a parent SHA1 is provided, then Bitbucket asserts that it represents the main branch&#x27;s current HEAD/tip, or a 409 is returned.  When a branch name is not specified and the repo is empty, the new commit will become the repo&#x27;s root commit and will be on the main branch.  When a branch name is specified and the repo is empty, the new commit will become the repo&#x27;s root commit and also define the repo&#x27;s main branch going forward.  This API cannot be used to create additional root commits in non-empty repos.  The branch field cannot be repeated.  As a side effect, this API can be used to create a new branch without modifying any files, by specifying a new branch name in this field, together with &#x60;parents&#x60;, but omitting the &#x60;files&#x60; fields, while not sending any files. This will create a new commit and branch with the same contents as the first parent. The diff of this commit against its first parent will be empty.  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

