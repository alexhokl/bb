# BranchingModelSettingsBranchTypes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the branch type is enabled or not. A disabled branch type may contain an invalid &#x60;prefix&#x60;. | [optional] [default to null]
**Kind** | **string** | The kind of the branch type. | [default to null]
**Prefix** | **string** | The prefix for this branch type. A branch with this prefix will be classified as per &#x60;kind&#x60;. The &#x60;prefix&#x60; of an enabled branch type must be a valid branch prefix.Additionally, it cannot be blank, empty or &#x60;null&#x60;. The &#x60;prefix&#x60; for a disabled branch type can be empty or invalid. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


