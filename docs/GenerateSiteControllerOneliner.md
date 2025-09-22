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
**OobHttpProxy** | **bool** | Enable OOB HTTP proxy capability | [default to false]
**InbandHttpProxy** | **bool** | Enable inband HTTP proxy capability | [default to false]
**FileTransfer** | **bool** | Enable file transfer capability | [default to false]
**InbandFileTransfer** | **bool** | Enable inband file transfer capability | [default to false]
**SwitchSubscription** | **bool** | Enable switch subscription capability | [default to false]
**CommandExecution** | **bool** | Enable command execution capability | [default to false]
**Netconf** | **bool** | Enable NETCONF capability | [default to false]
**Vnc** | **bool** | Enable VNC capability | [default to false]
**Spice** | **bool** | Enable SPICE capability | [default to false]
**Syslog** | **bool** | Enable syslog capability | [default to false]
**DhcpOob** | **bool** | Enable DHCP OOB capability | [default to false]
**AnsibleRunner** | **bool** | Enable Ansible runner capability | [default to false]
**HttpRequest** | **bool** | Enable HTTP request capability | [default to false]
**SshCommand** | **bool** | Enable SSH command capability | 
**SecondIp** | Pointer to **string** | Second IP address | [optional] 

## Methods

### NewGenerateSiteControllerOneliner

`func NewGenerateSiteControllerOneliner(usePodman bool, inbandMode bool, dockerEnv bool, registry string, gitHubTag string, localScript bool, sslHostname string, imagesTag string, msTunnelSecret string, oobHttpProxy bool, inbandHttpProxy bool, fileTransfer bool, inbandFileTransfer bool, switchSubscription bool, commandExecution bool, netconf bool, vnc bool, spice bool, syslog bool, dhcpOob bool, ansibleRunner bool, httpRequest bool, sshCommand bool, ) *GenerateSiteControllerOneliner`

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


### GetOobHttpProxy

`func (o *GenerateSiteControllerOneliner) GetOobHttpProxy() bool`

GetOobHttpProxy returns the OobHttpProxy field if non-nil, zero value otherwise.

### GetOobHttpProxyOk

`func (o *GenerateSiteControllerOneliner) GetOobHttpProxyOk() (*bool, bool)`

GetOobHttpProxyOk returns a tuple with the OobHttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobHttpProxy

`func (o *GenerateSiteControllerOneliner) SetOobHttpProxy(v bool)`

SetOobHttpProxy sets OobHttpProxy field to given value.


### GetInbandHttpProxy

`func (o *GenerateSiteControllerOneliner) GetInbandHttpProxy() bool`

GetInbandHttpProxy returns the InbandHttpProxy field if non-nil, zero value otherwise.

### GetInbandHttpProxyOk

`func (o *GenerateSiteControllerOneliner) GetInbandHttpProxyOk() (*bool, bool)`

GetInbandHttpProxyOk returns a tuple with the InbandHttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandHttpProxy

`func (o *GenerateSiteControllerOneliner) SetInbandHttpProxy(v bool)`

SetInbandHttpProxy sets InbandHttpProxy field to given value.


### GetFileTransfer

`func (o *GenerateSiteControllerOneliner) GetFileTransfer() bool`

GetFileTransfer returns the FileTransfer field if non-nil, zero value otherwise.

### GetFileTransferOk

`func (o *GenerateSiteControllerOneliner) GetFileTransferOk() (*bool, bool)`

GetFileTransferOk returns a tuple with the FileTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTransfer

`func (o *GenerateSiteControllerOneliner) SetFileTransfer(v bool)`

SetFileTransfer sets FileTransfer field to given value.


### GetInbandFileTransfer

`func (o *GenerateSiteControllerOneliner) GetInbandFileTransfer() bool`

GetInbandFileTransfer returns the InbandFileTransfer field if non-nil, zero value otherwise.

### GetInbandFileTransferOk

`func (o *GenerateSiteControllerOneliner) GetInbandFileTransferOk() (*bool, bool)`

GetInbandFileTransferOk returns a tuple with the InbandFileTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandFileTransfer

`func (o *GenerateSiteControllerOneliner) SetInbandFileTransfer(v bool)`

SetInbandFileTransfer sets InbandFileTransfer field to given value.


### GetSwitchSubscription

`func (o *GenerateSiteControllerOneliner) GetSwitchSubscription() bool`

GetSwitchSubscription returns the SwitchSubscription field if non-nil, zero value otherwise.

### GetSwitchSubscriptionOk

`func (o *GenerateSiteControllerOneliner) GetSwitchSubscriptionOk() (*bool, bool)`

GetSwitchSubscriptionOk returns a tuple with the SwitchSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchSubscription

`func (o *GenerateSiteControllerOneliner) SetSwitchSubscription(v bool)`

SetSwitchSubscription sets SwitchSubscription field to given value.


### GetCommandExecution

`func (o *GenerateSiteControllerOneliner) GetCommandExecution() bool`

GetCommandExecution returns the CommandExecution field if non-nil, zero value otherwise.

### GetCommandExecutionOk

`func (o *GenerateSiteControllerOneliner) GetCommandExecutionOk() (*bool, bool)`

GetCommandExecutionOk returns a tuple with the CommandExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandExecution

`func (o *GenerateSiteControllerOneliner) SetCommandExecution(v bool)`

SetCommandExecution sets CommandExecution field to given value.


### GetNetconf

`func (o *GenerateSiteControllerOneliner) GetNetconf() bool`

GetNetconf returns the Netconf field if non-nil, zero value otherwise.

### GetNetconfOk

`func (o *GenerateSiteControllerOneliner) GetNetconfOk() (*bool, bool)`

GetNetconfOk returns a tuple with the Netconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetconf

`func (o *GenerateSiteControllerOneliner) SetNetconf(v bool)`

SetNetconf sets Netconf field to given value.


### GetVnc

`func (o *GenerateSiteControllerOneliner) GetVnc() bool`

GetVnc returns the Vnc field if non-nil, zero value otherwise.

### GetVncOk

`func (o *GenerateSiteControllerOneliner) GetVncOk() (*bool, bool)`

GetVncOk returns a tuple with the Vnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnc

`func (o *GenerateSiteControllerOneliner) SetVnc(v bool)`

SetVnc sets Vnc field to given value.


### GetSpice

`func (o *GenerateSiteControllerOneliner) GetSpice() bool`

GetSpice returns the Spice field if non-nil, zero value otherwise.

### GetSpiceOk

`func (o *GenerateSiteControllerOneliner) GetSpiceOk() (*bool, bool)`

GetSpiceOk returns a tuple with the Spice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpice

`func (o *GenerateSiteControllerOneliner) SetSpice(v bool)`

SetSpice sets Spice field to given value.


### GetSyslog

`func (o *GenerateSiteControllerOneliner) GetSyslog() bool`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *GenerateSiteControllerOneliner) GetSyslogOk() (*bool, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *GenerateSiteControllerOneliner) SetSyslog(v bool)`

SetSyslog sets Syslog field to given value.


### GetDhcpOob

`func (o *GenerateSiteControllerOneliner) GetDhcpOob() bool`

GetDhcpOob returns the DhcpOob field if non-nil, zero value otherwise.

### GetDhcpOobOk

`func (o *GenerateSiteControllerOneliner) GetDhcpOobOk() (*bool, bool)`

GetDhcpOobOk returns a tuple with the DhcpOob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOob

`func (o *GenerateSiteControllerOneliner) SetDhcpOob(v bool)`

SetDhcpOob sets DhcpOob field to given value.


### GetAnsibleRunner

`func (o *GenerateSiteControllerOneliner) GetAnsibleRunner() bool`

GetAnsibleRunner returns the AnsibleRunner field if non-nil, zero value otherwise.

### GetAnsibleRunnerOk

`func (o *GenerateSiteControllerOneliner) GetAnsibleRunnerOk() (*bool, bool)`

GetAnsibleRunnerOk returns a tuple with the AnsibleRunner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsibleRunner

`func (o *GenerateSiteControllerOneliner) SetAnsibleRunner(v bool)`

SetAnsibleRunner sets AnsibleRunner field to given value.


### GetHttpRequest

`func (o *GenerateSiteControllerOneliner) GetHttpRequest() bool`

GetHttpRequest returns the HttpRequest field if non-nil, zero value otherwise.

### GetHttpRequestOk

`func (o *GenerateSiteControllerOneliner) GetHttpRequestOk() (*bool, bool)`

GetHttpRequestOk returns a tuple with the HttpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequest

`func (o *GenerateSiteControllerOneliner) SetHttpRequest(v bool)`

SetHttpRequest sets HttpRequest field to given value.


### GetSshCommand

`func (o *GenerateSiteControllerOneliner) GetSshCommand() bool`

GetSshCommand returns the SshCommand field if non-nil, zero value otherwise.

### GetSshCommandOk

`func (o *GenerateSiteControllerOneliner) GetSshCommandOk() (*bool, bool)`

GetSshCommandOk returns a tuple with the SshCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCommand

`func (o *GenerateSiteControllerOneliner) SetSshCommand(v bool)`

SetSshCommand sets SshCommand field to given value.


### GetSecondIp

`func (o *GenerateSiteControllerOneliner) GetSecondIp() string`

GetSecondIp returns the SecondIp field if non-nil, zero value otherwise.

### GetSecondIpOk

`func (o *GenerateSiteControllerOneliner) GetSecondIpOk() (*string, bool)`

GetSecondIpOk returns a tuple with the SecondIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondIp

`func (o *GenerateSiteControllerOneliner) SetSecondIp(v string)`

SetSecondIp sets SecondIp field to given value.

### HasSecondIp

`func (o *GenerateSiteControllerOneliner) HasSecondIp() bool`

HasSecondIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


