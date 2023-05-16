# InlineResponse200

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventsCount** | **int32** | The total number of ACTIVATED and SUSPENDED events to a max depth of 1 for all of the users in the query. A value larger than the limit specified on the query indicates that additional calls are needed, using a skip greater than 0, to retrieve the full set of results. | [optional] [default to null]
**Results** | [**[]ScheduledUserstateResult**](scheduled-userstate-result.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


