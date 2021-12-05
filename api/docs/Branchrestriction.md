# Branchrestriction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** |  | [default to null]
**Links** | [***BranchingModelSettingsLinks**](branching_model_settings_links.md) |  | [optional] [default to null]
**Id** | **int32** | The branch restriction status&#x27; id. | [optional] [default to null]
**Kind** | **string** | The type of restriction that is being applied. | [default to null]
**BranchMatchKind** | **string** | Indicates how the restriction is matched against a branch. The default is &#x60;glob&#x60;. | [default to null]
**BranchType** | **string** | Apply the restriction to branches of this type. Active when &#x60;branch_match_kind&#x60; is &#x60;branching_model&#x60;. The branch type will be calculated using the branching model configured for the repository. | [optional] [default to null]
**Pattern** | **string** | Apply the restriction to branches that match this pattern. Active when &#x60;branch_match_kind&#x60; is &#x60;glob&#x60;. Will be empty when &#x60;branch_match_kind&#x60; is &#x60;branching_model&#x60;. | [default to null]
**Users** | [**[]Account**](account.md) |  | [optional] [default to null]
**Groups** | [**[]Group**](group.md) |  | [optional] [default to null]
**Value** | **int32** | &lt;staticmethod object at 0x7f1446ac7690&gt; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

