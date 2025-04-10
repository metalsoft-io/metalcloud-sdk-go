# GenerateSiteControllerOneliner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsePodman** | **bool** | Use Podman instead of Docker | [default to false]
**InbandMode** | **bool** | Install in inband mode | [default to false]
**DockerEnv** | **bool** | Use Docker environment variables | [default to false]
**Registry** | **string** | Registry URL | [default to "registry.metalsoft.dev"]
**GitHubTag** | **string** | GitHub tag for deploy-agents.sh script | [default to "main"]
**LocalScript** | **bool** | Use local script instead of curling it | [default to false]
**SslHostname** | **string** | SSL hostname | 
**ImagesTag** | **string** | Docker images tag version ( The version of tunnel microservice ) | 
**MsTunnelSecret** | **string** | MS Tunnel Secret for secure connections | 

## Methods

### NewGenerateSiteControllerOneliner

`func NewGenerateSiteControllerOneliner(usePodman bool, inbandMode bool, dockerEnv bool, registry string, gitHubTag string, localScript bool, sslHostname string, imagesTag string, msTunnelSecret string, ) *GenerateSiteControllerOneliner`

NewGenerateSiteControllerOneliner instantiates a new GenerateSiteControllerOneliner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSiteControllerOnelinerWithDefaults

`func NewGenerateSiteControllerOnelinerWithDefaults() *GenerateSiteControllerOneliner`

NewGenerateSiteControllerOnelinerWithDefaults instantiates a new GenerateSiteControllerOneliner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsePodman

`func (o *GenerateSiteControllerOneliner) GetUsePodman() bool`

GetUsePodman returns the UsePodman field if non-nil, zero value otherwise.

### GetUsePodmanOk

`func (o *GenerateSiteControllerOneliner) GetUsePodmanOk() (*bool, bool)`

GetUsePodmanOk returns a tuple with the UsePodman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePodman

`func (o *GenerateSiteControllerOneliner) SetUsePodman(v bool)`

SetUsePodman sets UsePodman field to given value.


### GetInbandMode

`func (o *GenerateSiteControllerOneliner) GetInbandMode() bool`

GetInbandMode returns the InbandMode field if non-nil, zero value otherwise.

### GetInbandModeOk

`func (o *GenerateSiteControllerOneliner) GetInbandModeOk() (*bool, bool)`

GetInbandModeOk returns a tuple with the InbandMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandMode

`func (o *GenerateSiteControllerOneliner) SetInbandMode(v bool)`

SetInbandMode sets InbandMode field to given value.


### GetDockerEnv

`func (o *GenerateSiteControllerOneliner) GetDockerEnv() bool`

GetDockerEnv returns the DockerEnv field if non-nil, zero value otherwise.

### GetDockerEnvOk

`func (o *GenerateSiteControllerOneliner) GetDockerEnvOk() (*bool, bool)`

GetDockerEnvOk returns a tuple with the DockerEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerEnv

`func (o *GenerateSiteControllerOneliner) SetDockerEnv(v bool)`

SetDockerEnv sets DockerEnv field to given value.


### GetRegistry

`func (o *GenerateSiteControllerOneliner) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *GenerateSiteControllerOneliner) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *GenerateSiteControllerOneliner) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetGitHubTag

`func (o *GenerateSiteControllerOneliner) GetGitHubTag() string`

GetGitHubTag returns the GitHubTag field if non-nil, zero value otherwise.

### GetGitHubTagOk

`func (o *GenerateSiteControllerOneliner) GetGitHubTagOk() (*string, bool)`

GetGitHubTagOk returns a tuple with the GitHubTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubTag

`func (o *GenerateSiteControllerOneliner) SetGitHubTag(v string)`

SetGitHubTag sets GitHubTag field to given value.


### GetLocalScript

`func (o *GenerateSiteControllerOneliner) GetLocalScript() bool`

GetLocalScript returns the LocalScript field if non-nil, zero value otherwise.

### GetLocalScriptOk

`func (o *GenerateSiteControllerOneliner) GetLocalScriptOk() (*bool, bool)`

GetLocalScriptOk returns a tuple with the LocalScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScript

`func (o *GenerateSiteControllerOneliner) SetLocalScript(v bool)`

SetLocalScript sets LocalScript field to given value.


### GetSslHostname

`func (o *GenerateSiteControllerOneliner) GetSslHostname() string`

GetSslHostname returns the SslHostname field if non-nil, zero value otherwise.

### GetSslHostnameOk

`func (o *GenerateSiteControllerOneliner) GetSslHostnameOk() (*string, bool)`

GetSslHostnameOk returns a tuple with the SslHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslHostname

`func (o *GenerateSiteControllerOneliner) SetSslHostname(v string)`

SetSslHostname sets SslHostname field to given value.


### GetImagesTag

`func (o *GenerateSiteControllerOneliner) GetImagesTag() string`

GetImagesTag returns the ImagesTag field if non-nil, zero value otherwise.

### GetImagesTagOk

`func (o *GenerateSiteControllerOneliner) GetImagesTagOk() (*string, bool)`

GetImagesTagOk returns a tuple with the ImagesTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagesTag

`func (o *GenerateSiteControllerOneliner) SetImagesTag(v string)`

SetImagesTag sets ImagesTag field to given value.


### GetMsTunnelSecret

`func (o *GenerateSiteControllerOneliner) GetMsTunnelSecret() string`

GetMsTunnelSecret returns the MsTunnelSecret field if non-nil, zero value otherwise.

### GetMsTunnelSecretOk

`func (o *GenerateSiteControllerOneliner) GetMsTunnelSecretOk() (*string, bool)`

GetMsTunnelSecretOk returns a tuple with the MsTunnelSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsTunnelSecret

`func (o *GenerateSiteControllerOneliner) SetMsTunnelSecret(v string)`

SetMsTunnelSecret sets MsTunnelSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


