# \SecurityAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePermission**](SecurityAPI.md#CreatePermission) | **Post** /api/v2/permissions | Create a new permission
[**CreateQuotaProfile**](SecurityAPI.md#CreateQuotaProfile) | **Post** /api/v2/quota-profiles | Create a new quota profile
[**CreateRole**](SecurityAPI.md#CreateRole) | **Post** /api/v2/roles | Create a new role
[**DeletePermission**](SecurityAPI.md#DeletePermission) | **Delete** /api/v2/permissions/{permissionName} | Delete a permission by name
[**DeleteQuotaProfile**](SecurityAPI.md#DeleteQuotaProfile) | **Delete** /api/v2/quota-profiles/{profileId} | Delete a quota profile by id
[**DeleteRole**](SecurityAPI.md#DeleteRole) | **Delete** /api/v2/roles/{roleName} | Delete a role by name
[**GetPermissions**](SecurityAPI.md#GetPermissions) | **Get** /api/v2/permissions | Get all permissions
[**GetProviders**](SecurityAPI.md#GetProviders) | **Get** /api/v2/authentication/providers | Get available authentication providers
[**GetQuotaProfile**](SecurityAPI.md#GetQuotaProfile) | **Get** /api/v2/quota-profiles/{profileId} | Get a quota profile by id
[**GetQuotaProfiles**](SecurityAPI.md#GetQuotaProfiles) | **Get** /api/v2/quota-profiles | Get all quota profiles
[**GetRole**](SecurityAPI.md#GetRole) | **Get** /api/v2/roles/{roleName} | Get a role by name
[**GetRoles**](SecurityAPI.md#GetRoles) | **Get** /api/v2/roles | Get all roles
[**UpdateProvider**](SecurityAPI.md#UpdateProvider) | **Patch** /api/v2/authentication/providers/{name} | Updates authentication provider
[**UpdateQuotaProfile**](SecurityAPI.md#UpdateQuotaProfile) | **Patch** /api/v2/quota-profiles/{profileId} | Update a quota profile by id
[**UpdateRole**](SecurityAPI.md#UpdateRole) | **Put** /api/v2/roles/{roleName} | Update a role by name



## CreatePermission

> Permission CreatePermission(ctx).CreatePermission(createPermission).Execute()

Create a new permission

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	createPermission := *openapiclient.NewCreatePermission("Name_example", "Label_example") // CreatePermission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreatePermission(context.Background()).CreatePermission(createPermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreatePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePermission`: Permission
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreatePermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPermission** | [**CreatePermission**](CreatePermission.md) |  | 

### Return type

[**Permission**](Permission.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuotaProfile

> QuotaProfile CreateQuotaProfile(ctx).CreateQuotaProfile(createQuotaProfile).Execute()

Create a new quota profile

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	createQuotaProfile := *openapiclient.NewCreateQuotaProfile("My Profile") // CreateQuotaProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateQuotaProfile(context.Background()).CreateQuotaProfile(createQuotaProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateQuotaProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuotaProfile`: QuotaProfile
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateQuotaProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuotaProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createQuotaProfile** | [**CreateQuotaProfile**](CreateQuotaProfile.md) |  | 

### Return type

[**QuotaProfile**](QuotaProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> Role CreateRole(ctx).CreateRole(createRole).Execute()

Create a new role

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	createRole := *openapiclient.NewCreateRole("Label_example", []openapiclient.MetalsoftPermissions{openapiclient.MetalsoftPermissions("admin_maintenance_write")}) // CreateRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateRole(context.Background()).CreateRole(createRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRole** | [**CreateRole**](CreateRole.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePermission

> DeletePermission(ctx, permissionName).Execute()

Delete a permission by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	permissionName := "permissionName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeletePermission(context.Background(), permissionName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeletePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuotaProfile

> DeleteQuotaProfile(ctx, profileId).Execute()

Delete a quota profile by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	profileId := "profileId_example" // string | The quota profile id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteQuotaProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteQuotaProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The quota profile id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQuotaProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, roleName).Execute()

Delete a role by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteRole(context.Background(), roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissions

> PermissionPaginatedList GetPermissions(ctx).Page(page).Limit(limit).FilterType(filterType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all permissions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterType := []string{"Inner_example"} // []string | Filter by type query param.  **Format:** filter.type={$not}:OPERATION:VALUE    **Example:** filter.type=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** label:ASC  **Available Fields** - id  - name  - label  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,label,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  - description  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetPermissions(context.Background()).Page(page).Limit(limit).FilterType(filterType).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissions`: PermissionPaginatedList
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterType** | **[]string** | Filter by type query param.  **Format:** filter.type&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.type&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** label:ASC  **Available Fields** - id  - name  - label  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,label,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  - description  | 

### Return type

[**PermissionPaginatedList**](PermissionPaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviders

> []AuthenticationProvider GetProviders(ctx).Execute()

Get available authentication providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviders`: []AuthenticationProvider
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersRequest struct via the builder pattern


### Return type

[**[]AuthenticationProvider**](AuthenticationProvider.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotaProfile

> QuotaProfile GetQuotaProfile(ctx, profileId).Execute()

Get a quota profile by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	profileId := "profileId_example" // string | The quota profile id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetQuotaProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetQuotaProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotaProfile`: QuotaProfile
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetQuotaProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The quota profile id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotaProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuotaProfile**](QuotaProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotaProfiles

> QuotaProfilePaginatedList GetQuotaProfiles(ctx).Page(page).Limit(limit).FilterId(filterId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all quota profiles

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterId := []string{"Inner_example"} // []string | Filter by id query param.  **Format:** filter.id={$not}:OPERATION:VALUE    **Example:** filter.id=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** name:ASC  **Available Fields** - id  - name  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** id,name,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - description  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetQuotaProfiles(context.Background()).Page(page).Limit(limit).FilterId(filterId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetQuotaProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotaProfiles`: QuotaProfilePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetQuotaProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotaProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterId** | **[]string** | Filter by id query param.  **Format:** filter.id&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.id&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** name:ASC  **Available Fields** - id  - name  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** id,name,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - id  - name  - description  | 

### Return type

[**QuotaProfilePaginatedList**](QuotaProfilePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx, roleName).Execute()

Get a role by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetRole(context.Background(), roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoles

> RolePaginatedList GetRoles(ctx).Page(page).Limit(limit).FilterType(filterType).FilterQuotaProfileId(filterQuotaProfileId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()

Get all roles

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	page := float32(8.14) // float32 | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   (optional)
	limit := float32(8.14) // float32 | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  (optional)
	filterType := []string{"Inner_example"} // []string | Filter by type query param.  **Format:** filter.type={$not}:OPERATION:VALUE    **Example:** filter.type=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	filterQuotaProfileId := []string{"Inner_example"} // []string | Filter by quotaProfileId query param.  **Format:** filter.quotaProfileId={$not}:OPERATION:VALUE    **Example:** filter.quotaProfileId=$eq:John Doe  **Available Operations** - $eq  - $and  - $or (optional)
	sortBy := []string{"SortBy_example"} // []string | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy=id:DESC&sortBy=name:DESC   **Default Value:** label:ASC  **Available Fields** - id  - name  - label  (optional)
	search := "search_example" // string | Search term to filter result values  **Example:** John   **Default Value:** No default value   (optional)
	searchBy := []string{"Inner_example"} // []string | List of fields to search by term to filter result values  **Example:** name,label,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  - description  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetRoles(context.Background()).Page(page).Limit(limit).FilterType(filterType).FilterQuotaProfileId(filterQuotaProfileId).SortBy(sortBy).Search(search).SearchBy(searchBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoles`: RolePaginatedList
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** | Page number to retrieve. If you provide invalid value the default page number will applied  **Example:** 1   **Default Value:** 1   | 
 **limit** | **float32** | Number of records per page.   **Example:** 20    **Default Value:** 20    **Max Value:** 100   If provided value is greater than max value, max value will be applied.  | 
 **filterType** | **[]string** | Filter by type query param.  **Format:** filter.type&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.type&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **filterQuotaProfileId** | **[]string** | Filter by quotaProfileId query param.  **Format:** filter.quotaProfileId&#x3D;{$not}:OPERATION:VALUE    **Example:** filter.quotaProfileId&#x3D;$eq:John Doe  **Available Operations** - $eq  - $and  - $or | 
 **sortBy** | **[]string** | Parameter to sort by. To sort by multiple fields, just provide query param multiple types. The order in url defines an order of sorting  **Format:** {fieldName}:{DIRECTION}   **Example:** sortBy&#x3D;id:DESC&amp;sortBy&#x3D;name:DESC   **Default Value:** label:ASC  **Available Fields** - id  - name  - label  | 
 **search** | **string** | Search term to filter result values  **Example:** John   **Default Value:** No default value   | 
 **searchBy** | **[]string** | List of fields to search by term to filter result values  **Example:** name,label,description   **Default Value:** By default all fields mentioned below will be used to search by term  **Available Fields** - name  - label  - description  | 

### Return type

[**RolePaginatedList**](RolePaginatedList.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvider

> AuthenticationProvider UpdateProvider(ctx, name).AuthenticationProviderUpdate(authenticationProviderUpdate).Execute()

Updates authentication provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	name := "name_example" // string | The provider name
	authenticationProviderUpdate := *openapiclient.NewAuthenticationProviderUpdate(false) // AuthenticationProviderUpdate | The provider updates

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.UpdateProvider(context.Background(), name).AuthenticationProviderUpdate(authenticationProviderUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProvider`: AuthenticationProvider
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.UpdateProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The provider name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationProviderUpdate** | [**AuthenticationProviderUpdate**](AuthenticationProviderUpdate.md) | The provider updates | 

### Return type

[**AuthenticationProvider**](AuthenticationProvider.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuotaProfile

> QuotaProfile UpdateQuotaProfile(ctx, profileId).EditQuotaProfile(editQuotaProfile).Execute()

Update a quota profile by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	profileId := "profileId_example" // string | The quota profile id
	editQuotaProfile := *openapiclient.NewEditQuotaProfile() // EditQuotaProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.UpdateQuotaProfile(context.Background(), profileId).EditQuotaProfile(editQuotaProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateQuotaProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQuotaProfile`: QuotaProfile
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.UpdateQuotaProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The quota profile id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuotaProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editQuotaProfile** | [**EditQuotaProfile**](EditQuotaProfile.md) |  | 

### Return type

[**QuotaProfile**](QuotaProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> Role UpdateRole(ctx, roleName).EditRole(editRole).Execute()

Update a role by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metalsoft-io/metalcloud-sdk-go"
)

func main() {
	roleName := "roleName_example" // string | 
	editRole := *openapiclient.NewEditRole() // EditRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.UpdateRole(context.Background(), roleName).EditRole(editRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editRole** | [**EditRole**](EditRole.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

