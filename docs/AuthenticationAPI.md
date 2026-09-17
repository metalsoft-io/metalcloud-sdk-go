# \AuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthenticationProvider**](AuthenticationAPI.md#GetAuthenticationProvider) | **Get** /api/v2/authentication-provider | Retrieve authentication provider for an email
[**GetBranding**](AuthenticationAPI.md#GetBranding) | **Get** /api/v2/branding | Get branding configuration
[**GetCaptchaConfig**](AuthenticationAPI.md#GetCaptchaConfig) | **Get** /api/v2/captcha | Get CAPTCHA configuration
[**GetCurrentUser**](AuthenticationAPI.md#GetCurrentUser) | **Get** /api/v2/user | Get current user
[**GetPublicSignupDisabled**](AuthenticationAPI.md#GetPublicSignupDisabled) | **Get** /api/v2/public-signup-disabled | Check if public signup is disabled
[**Login**](AuthenticationAPI.md#Login) | **Post** /api/v2/login | User login
[**Login2FA**](AuthenticationAPI.md#Login2FA) | **Post** /api/v2/login-2fa | User login with 2FA
[**LoginOidc**](AuthenticationAPI.md#LoginOidc) | **Get** /api/v2/login-oidc | Start an OIDC login
[**LoginOidcCallback**](AuthenticationAPI.md#LoginOidcCallback) | **Get** /api/v2/login-oidc/callback | OIDC login callback URL
[**LoginSaml**](AuthenticationAPI.md#LoginSaml) | **Post** /api/v2/login-saml | User SAML login callback URL
[**Logout**](AuthenticationAPI.md#Logout) | **Post** /api/v2/logout | Logout user
[**SignUp**](AuthenticationAPI.md#SignUp) | **Post** /api/v2/sign-up | Sign up
[**TestLdap**](AuthenticationAPI.md#TestLdap) | **Post** /api/v2/authentication/ldap/actions/test | Test authentication against the stored LDAP configuration
[**ValidateLdap**](AuthenticationAPI.md#ValidateLdap) | **Post** /api/v2/authentication/ldap/actions/validate | Validate the stored LDAP configuration
[**ValidateOidc**](AuthenticationAPI.md#ValidateOidc) | **Post** /api/v2/authentication/oidc/actions/validate | Validate the stored OIDC configuration
[**ValidateSaml**](AuthenticationAPI.md#ValidateSaml) | **Post** /api/v2/authentication/saml/actions/validate | Validate the stored SAML configuration



## GetAuthenticationProvider

> AuthenticationUserProvider GetAuthenticationProvider(ctx).Email(email).Execute()

Retrieve authentication provider for an email



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
	email := "email_example" // string | The email address to look up

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetAuthenticationProvider(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetAuthenticationProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticationProvider`: AuthenticationUserProvider
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetAuthenticationProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticationProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | The email address to look up | 

### Return type

[**AuthenticationUserProvider**](AuthenticationUserProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranding

> Branding GetBranding(ctx).Execute()

Get branding configuration



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
	resp, r, err := apiClient.AuthenticationAPI.GetBranding(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetBranding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranding`: Branding
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetBranding`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingRequest struct via the builder pattern


### Return type

[**Branding**](Branding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptchaConfig

> CaptchaConfig GetCaptchaConfig(ctx).Execute()

Get CAPTCHA configuration



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
	resp, r, err := apiClient.AuthenticationAPI.GetCaptchaConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetCaptchaConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptchaConfig`: CaptchaConfig
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetCaptchaConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaConfigRequest struct via the builder pattern


### Return type

[**CaptchaConfig**](CaptchaConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> User GetCurrentUser(ctx).Execute()

Get current user



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
	resp, r, err := apiClient.AuthenticationAPI.GetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUser`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicSignupDisabled

> PublicSignupDisabled GetPublicSignupDisabled(ctx).Execute()

Check if public signup is disabled



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
	resp, r, err := apiClient.AuthenticationAPI.GetPublicSignupDisabled(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetPublicSignupDisabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicSignupDisabled`: PublicSignupDisabled
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetPublicSignupDisabled`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicSignupDisabledRequest struct via the builder pattern


### Return type

[**PublicSignupDisabled**](PublicSignupDisabled.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> User Login(ctx).AuthenticationRequest(authenticationRequest).Execute()

User login



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
	authenticationRequest := *openapiclient.NewAuthenticationRequest("Provider_example", openapiclient.AuthenticationRequest_properties{AuthenticationRequestPropertiesLdap: openapiclient.NewAuthenticationRequestPropertiesLdap("test.email@metalsoft.io", "secret-password")}) // AuthenticationRequest | The authentication request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Login(context.Background()).AuthenticationRequest(authenticationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationRequest** | [**AuthenticationRequest**](AuthenticationRequest.md) | The authentication request | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login2FA

> User Login2FA(ctx).TwoFactorAuthenticationToken(twoFactorAuthenticationToken).Execute()

User login with 2FA



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
	twoFactorAuthenticationToken := *openapiclient.NewTwoFactorAuthenticationToken() // TwoFactorAuthenticationToken | The 2FA token request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Login2FA(context.Background()).TwoFactorAuthenticationToken(twoFactorAuthenticationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Login2FA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login2FA`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Login2FA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogin2FARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthenticationToken** | [**TwoFactorAuthenticationToken**](TwoFactorAuthenticationToken.md) | The 2FA token request | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginOidc

> LoginOidc(ctx).Execute()

Start an OIDC login



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
	r, err := apiClient.AuthenticationAPI.LoginOidc(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.LoginOidc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoginOidcRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginOidcCallback

> LoginOidcCallback(ctx).Error_(error_).Code(code).State(state).Execute()

OIDC login callback URL



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
	error_ := "error__example" // string | 
	code := "code_example" // string | Authorization code from the IdP (optional)
	state := "state_example" // string | State issued when the flow started (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.LoginOidcCallback(context.Background()).Error_(error_).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.LoginOidcCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginOidcCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **error_** | **string** |  | 
 **code** | **string** | Authorization code from the IdP | 
 **state** | **string** | State issued when the flow started | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginSaml

> LoginSaml(ctx).AuthenticationRequestPropertiesSaml(authenticationRequestPropertiesSaml).Execute()

User SAML login callback URL



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
	authenticationRequestPropertiesSaml := *openapiclient.NewAuthenticationRequestPropertiesSaml("SAMLResponse_example") // AuthenticationRequestPropertiesSaml | The SAML response

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.LoginSaml(context.Background()).AuthenticationRequestPropertiesSaml(authenticationRequestPropertiesSaml).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.LoginSaml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationRequestPropertiesSaml** | [**AuthenticationRequestPropertiesSaml**](AuthenticationRequestPropertiesSaml.md) | The SAML response | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> LogoutRequest Logout(ctx).Execute()

Logout user



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
	resp, r, err := apiClient.AuthenticationAPI.Logout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Logout`: LogoutRequest
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Logout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


### Return type

[**LogoutRequest**](LogoutRequest.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignUp

> User SignUp(ctx).CreateUser(createUser).Execute()

Sign up



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
	createUser := *openapiclient.NewCreateUser("DisplayName_example", "AccessLevel_example", "Email_example") // CreateUser | The user to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.SignUp(context.Background()).CreateUser(createUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.SignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignUp`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.SignUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUser** | [**CreateUser**](CreateUser.md) | The user to create | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestLdap

> LdapTestResult TestLdap(ctx).LdapTestRequest(ldapTestRequest).Execute()

Test authentication against the stored LDAP configuration



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
	ldapTestRequest := *openapiclient.NewLdapTestRequest("Username_example", "Password_example") // LdapTestRequest | Credentials to authenticate against the configured LDAP directory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.TestLdap(context.Background()).LdapTestRequest(ldapTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.TestLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestLdap`: LdapTestResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.TestLdap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapTestRequest** | [**LdapTestRequest**](LdapTestRequest.md) | Credentials to authenticate against the configured LDAP directory | 

### Return type

[**LdapTestResult**](LdapTestResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateLdap

> AuthConfigValidationResult ValidateLdap(ctx).Execute()

Validate the stored LDAP configuration



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
	resp, r, err := apiClient.AuthenticationAPI.ValidateLdap(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ValidateLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateLdap`: AuthConfigValidationResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ValidateLdap`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiValidateLdapRequest struct via the builder pattern


### Return type

[**AuthConfigValidationResult**](AuthConfigValidationResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateOidc

> AuthConfigValidationResult ValidateOidc(ctx).Execute()

Validate the stored OIDC configuration



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
	resp, r, err := apiClient.AuthenticationAPI.ValidateOidc(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ValidateOidc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateOidc`: AuthConfigValidationResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ValidateOidc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiValidateOidcRequest struct via the builder pattern


### Return type

[**AuthConfigValidationResult**](AuthConfigValidationResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSaml

> AuthConfigValidationResult ValidateSaml(ctx).Execute()

Validate the stored SAML configuration



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
	resp, r, err := apiClient.AuthenticationAPI.ValidateSaml(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ValidateSaml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSaml`: AuthConfigValidationResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ValidateSaml`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiValidateSamlRequest struct via the builder pattern


### Return type

[**AuthConfigValidationResult**](AuthConfigValidationResult.md)

### Authorization

[apiKey](../README.md#apiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

