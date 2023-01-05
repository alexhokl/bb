# Report

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Uuid** | **string** | The UUID that can be used to identify the report. | [optional] [default to null]
**Title** | **string** | The title of the report. | [optional] [default to null]
**Details** | **string** | A string to describe the purpose of the report. | [optional] [default to null]
**ExternalId** | **string** | ID of the report provided by the report creator. It can be used to identify the report as an alternative to it&#39;s generated uuid. It is not used by Bitbucket, but only by the report creator for updating or deleting this specific report. Needs to be unique. | [optional] [default to null]
**Reporter** | **string** | A string to describe the tool or company who created the report. | [optional] [default to null]
**Link** | **string** | A URL linking to the results of the report in an external tool. | [optional] [default to null]
**RemoteLinkEnabled** | **bool** | If enabled, a remote link is created in Jira for the issue associated with the commit the report belongs to. | [optional] [default to null]
**LogoUrl** | **string** | A URL to the report logo. If none is provided, the default insights logo will be used. | [optional] [default to null]
**ReportType** | **string** | The type of the report. | [optional] [default to null]
**Result** | **string** | The state of the report. May be set to PENDING and later updated. | [optional] [default to null]
**Data** | [**[]ReportData**](report_data.md) | An array of data fields to display information on the report. Maximum 10. | [optional] [default to null]
**CreatedOn** | [**time.Time**](time.Time.md) | The timestamp when the report was created. | [optional] [default to null]
**UpdatedOn** | [**time.Time**](time.Time.md) | The timestamp when the report was updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


