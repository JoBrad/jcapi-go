# ImportUsersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUserReactivation** | **bool** | A boolean value to allow the reactivation of suspended users | [optional] [default to null]
**Operations** | [**[]ImportOperation**](ImportOperation.md) | Operations to be performed on the user list returned from the application | [optional] [default to null]
**QueryString** | **string** | Query string to filter and sort the user list returned from the application.  The supported filtering and sorting varies by application.  If no value is sent, all users are returned. **Example:** \&quot;location&#x3D;Chicago&amp;department&#x3D;IT\&quot;Query string used to retrieve users from service | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


