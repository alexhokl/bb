# ReportAnnotation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**ExternalId** | **string** | ID of the annotation provided by the annotation creator. It can be used to identify the annotation as an alternative to it&#x27;s generated uuid. It is not used by Bitbucket, but only by the annotation creator for updating or deleting this specific annotation. Needs to be unique. | [optional] [default to null]
**Uuid** | **string** | The UUID that can be used to identify the annotation. | [optional] [default to null]
**AnnotationType** | **string** | The type of the report. | [optional] [default to null]
**Path** | **string** | The path of the file on which this annotation should be placed. This is the path of the file relative to the git repository. If no path is provided, then it will appear in the overview modal on all pull requests where the tip of the branch is the given commit, regardless of which files were modified. | [optional] [default to null]
**Line** | **int32** | The line number that the annotation should belong to. If no line number is provided, then it will default to 0 and in a pull request it will appear at the top of the file specified by the path field. | [optional] [default to null]
**Summary** | **string** | The message to display to users. | [optional] [default to null]
**Details** | **string** | The details to show to users when clicking on the annotation. | [optional] [default to null]
**Result** | **string** | The state of the report. May be set to PENDING and later updated. | [optional] [default to null]
**Severity** | **string** | The severity of the annotation. | [optional] [default to null]
**Link** | **string** | A URL linking to the annotation in an external tool. | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) | The timestamp when the report was created. | [optional] [default to null]
**UpdatedOn** | [**time.Time**](time.Time.md) | The timestamp when the report was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

