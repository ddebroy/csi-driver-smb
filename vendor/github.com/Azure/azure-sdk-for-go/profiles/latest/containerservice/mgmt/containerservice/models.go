// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package containerservice

import original "github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2017-09-30/containerservice"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ContainerServicesClient = original.ContainerServicesClient
type ManagedClustersClient = original.ManagedClustersClient
type OrchestratorTypes = original.OrchestratorTypes

const (
	Custom     OrchestratorTypes = original.Custom
	DCOS       OrchestratorTypes = original.DCOS
	DockerCE   OrchestratorTypes = original.DockerCE
	Kubernetes OrchestratorTypes = original.Kubernetes
	Swarm      OrchestratorTypes = original.Swarm
)

type OSType = original.OSType

const (
	Linux   OSType = original.Linux
	Windows OSType = original.Windows
)

type StorageProfileTypes = original.StorageProfileTypes

const (
	ManagedDisks   StorageProfileTypes = original.ManagedDisks
	StorageAccount StorageProfileTypes = original.StorageAccount
)

type VMSizeTypes = original.VMSizeTypes

const (
	StandardA1          VMSizeTypes = original.StandardA1
	StandardA10         VMSizeTypes = original.StandardA10
	StandardA11         VMSizeTypes = original.StandardA11
	StandardA1V2        VMSizeTypes = original.StandardA1V2
	StandardA2          VMSizeTypes = original.StandardA2
	StandardA2mV2       VMSizeTypes = original.StandardA2mV2
	StandardA2V2        VMSizeTypes = original.StandardA2V2
	StandardA3          VMSizeTypes = original.StandardA3
	StandardA4          VMSizeTypes = original.StandardA4
	StandardA4mV2       VMSizeTypes = original.StandardA4mV2
	StandardA4V2        VMSizeTypes = original.StandardA4V2
	StandardA5          VMSizeTypes = original.StandardA5
	StandardA6          VMSizeTypes = original.StandardA6
	StandardA7          VMSizeTypes = original.StandardA7
	StandardA8          VMSizeTypes = original.StandardA8
	StandardA8mV2       VMSizeTypes = original.StandardA8mV2
	StandardA8V2        VMSizeTypes = original.StandardA8V2
	StandardA9          VMSizeTypes = original.StandardA9
	StandardB2ms        VMSizeTypes = original.StandardB2ms
	StandardB2s         VMSizeTypes = original.StandardB2s
	StandardB4ms        VMSizeTypes = original.StandardB4ms
	StandardB8ms        VMSizeTypes = original.StandardB8ms
	StandardD1          VMSizeTypes = original.StandardD1
	StandardD11         VMSizeTypes = original.StandardD11
	StandardD11V2       VMSizeTypes = original.StandardD11V2
	StandardD11V2Promo  VMSizeTypes = original.StandardD11V2Promo
	StandardD12         VMSizeTypes = original.StandardD12
	StandardD12V2       VMSizeTypes = original.StandardD12V2
	StandardD12V2Promo  VMSizeTypes = original.StandardD12V2Promo
	StandardD13         VMSizeTypes = original.StandardD13
	StandardD13V2       VMSizeTypes = original.StandardD13V2
	StandardD13V2Promo  VMSizeTypes = original.StandardD13V2Promo
	StandardD14         VMSizeTypes = original.StandardD14
	StandardD14V2       VMSizeTypes = original.StandardD14V2
	StandardD14V2Promo  VMSizeTypes = original.StandardD14V2Promo
	StandardD15V2       VMSizeTypes = original.StandardD15V2
	StandardD16sV3      VMSizeTypes = original.StandardD16sV3
	StandardD16V3       VMSizeTypes = original.StandardD16V3
	StandardD1V2        VMSizeTypes = original.StandardD1V2
	StandardD2          VMSizeTypes = original.StandardD2
	StandardD2sV3       VMSizeTypes = original.StandardD2sV3
	StandardD2V2        VMSizeTypes = original.StandardD2V2
	StandardD2V2Promo   VMSizeTypes = original.StandardD2V2Promo
	StandardD2V3        VMSizeTypes = original.StandardD2V3
	StandardD3          VMSizeTypes = original.StandardD3
	StandardD32sV3      VMSizeTypes = original.StandardD32sV3
	StandardD32V3       VMSizeTypes = original.StandardD32V3
	StandardD3V2        VMSizeTypes = original.StandardD3V2
	StandardD3V2Promo   VMSizeTypes = original.StandardD3V2Promo
	StandardD4          VMSizeTypes = original.StandardD4
	StandardD4sV3       VMSizeTypes = original.StandardD4sV3
	StandardD4V2        VMSizeTypes = original.StandardD4V2
	StandardD4V2Promo   VMSizeTypes = original.StandardD4V2Promo
	StandardD4V3        VMSizeTypes = original.StandardD4V3
	StandardD5V2        VMSizeTypes = original.StandardD5V2
	StandardD5V2Promo   VMSizeTypes = original.StandardD5V2Promo
	StandardD64sV3      VMSizeTypes = original.StandardD64sV3
	StandardD64V3       VMSizeTypes = original.StandardD64V3
	StandardD8sV3       VMSizeTypes = original.StandardD8sV3
	StandardD8V3        VMSizeTypes = original.StandardD8V3
	StandardDS1         VMSizeTypes = original.StandardDS1
	StandardDS11        VMSizeTypes = original.StandardDS11
	StandardDS11V2      VMSizeTypes = original.StandardDS11V2
	StandardDS11V2Promo VMSizeTypes = original.StandardDS11V2Promo
	StandardDS12        VMSizeTypes = original.StandardDS12
	StandardDS12V2      VMSizeTypes = original.StandardDS12V2
	StandardDS12V2Promo VMSizeTypes = original.StandardDS12V2Promo
	StandardDS13        VMSizeTypes = original.StandardDS13
	StandardDS132V2     VMSizeTypes = original.StandardDS132V2
	StandardDS134V2     VMSizeTypes = original.StandardDS134V2
	StandardDS13V2      VMSizeTypes = original.StandardDS13V2
	StandardDS13V2Promo VMSizeTypes = original.StandardDS13V2Promo
	StandardDS14        VMSizeTypes = original.StandardDS14
	StandardDS144V2     VMSizeTypes = original.StandardDS144V2
	StandardDS148V2     VMSizeTypes = original.StandardDS148V2
	StandardDS14V2      VMSizeTypes = original.StandardDS14V2
	StandardDS14V2Promo VMSizeTypes = original.StandardDS14V2Promo
	StandardDS15V2      VMSizeTypes = original.StandardDS15V2
	StandardDS1V2       VMSizeTypes = original.StandardDS1V2
	StandardDS2         VMSizeTypes = original.StandardDS2
	StandardDS2V2       VMSizeTypes = original.StandardDS2V2
	StandardDS2V2Promo  VMSizeTypes = original.StandardDS2V2Promo
	StandardDS3         VMSizeTypes = original.StandardDS3
	StandardDS3V2       VMSizeTypes = original.StandardDS3V2
	StandardDS3V2Promo  VMSizeTypes = original.StandardDS3V2Promo
	StandardDS4         VMSizeTypes = original.StandardDS4
	StandardDS4V2       VMSizeTypes = original.StandardDS4V2
	StandardDS4V2Promo  VMSizeTypes = original.StandardDS4V2Promo
	StandardDS5V2       VMSizeTypes = original.StandardDS5V2
	StandardDS5V2Promo  VMSizeTypes = original.StandardDS5V2Promo
	StandardE16sV3      VMSizeTypes = original.StandardE16sV3
	StandardE16V3       VMSizeTypes = original.StandardE16V3
	StandardE2sV3       VMSizeTypes = original.StandardE2sV3
	StandardE2V3        VMSizeTypes = original.StandardE2V3
	StandardE3216sV3    VMSizeTypes = original.StandardE3216sV3
	StandardE328sV3     VMSizeTypes = original.StandardE328sV3
	StandardE32sV3      VMSizeTypes = original.StandardE32sV3
	StandardE32V3       VMSizeTypes = original.StandardE32V3
	StandardE4sV3       VMSizeTypes = original.StandardE4sV3
	StandardE4V3        VMSizeTypes = original.StandardE4V3
	StandardE6416sV3    VMSizeTypes = original.StandardE6416sV3
	StandardE6432sV3    VMSizeTypes = original.StandardE6432sV3
	StandardE64sV3      VMSizeTypes = original.StandardE64sV3
	StandardE64V3       VMSizeTypes = original.StandardE64V3
	StandardE8sV3       VMSizeTypes = original.StandardE8sV3
	StandardE8V3        VMSizeTypes = original.StandardE8V3
	StandardF1          VMSizeTypes = original.StandardF1
	StandardF16         VMSizeTypes = original.StandardF16
	StandardF16s        VMSizeTypes = original.StandardF16s
	StandardF16sV2      VMSizeTypes = original.StandardF16sV2
	StandardF1s         VMSizeTypes = original.StandardF1s
	StandardF2          VMSizeTypes = original.StandardF2
	StandardF2s         VMSizeTypes = original.StandardF2s
	StandardF2sV2       VMSizeTypes = original.StandardF2sV2
	StandardF32sV2      VMSizeTypes = original.StandardF32sV2
	StandardF4          VMSizeTypes = original.StandardF4
	StandardF4s         VMSizeTypes = original.StandardF4s
	StandardF4sV2       VMSizeTypes = original.StandardF4sV2
	StandardF64sV2      VMSizeTypes = original.StandardF64sV2
	StandardF72sV2      VMSizeTypes = original.StandardF72sV2
	StandardF8          VMSizeTypes = original.StandardF8
	StandardF8s         VMSizeTypes = original.StandardF8s
	StandardF8sV2       VMSizeTypes = original.StandardF8sV2
	StandardG1          VMSizeTypes = original.StandardG1
	StandardG2          VMSizeTypes = original.StandardG2
	StandardG3          VMSizeTypes = original.StandardG3
	StandardG4          VMSizeTypes = original.StandardG4
	StandardG5          VMSizeTypes = original.StandardG5
	StandardGS1         VMSizeTypes = original.StandardGS1
	StandardGS2         VMSizeTypes = original.StandardGS2
	StandardGS3         VMSizeTypes = original.StandardGS3
	StandardGS4         VMSizeTypes = original.StandardGS4
	StandardGS44        VMSizeTypes = original.StandardGS44
	StandardGS48        VMSizeTypes = original.StandardGS48
	StandardGS5         VMSizeTypes = original.StandardGS5
	StandardGS516       VMSizeTypes = original.StandardGS516
	StandardGS58        VMSizeTypes = original.StandardGS58
	StandardH16         VMSizeTypes = original.StandardH16
	StandardH16m        VMSizeTypes = original.StandardH16m
	StandardH16mr       VMSizeTypes = original.StandardH16mr
	StandardH16r        VMSizeTypes = original.StandardH16r
	StandardH8          VMSizeTypes = original.StandardH8
	StandardH8m         VMSizeTypes = original.StandardH8m
	StandardL16s        VMSizeTypes = original.StandardL16s
	StandardL32s        VMSizeTypes = original.StandardL32s
	StandardL4s         VMSizeTypes = original.StandardL4s
	StandardL8s         VMSizeTypes = original.StandardL8s
	StandardM12832ms    VMSizeTypes = original.StandardM12832ms
	StandardM12864ms    VMSizeTypes = original.StandardM12864ms
	StandardM128ms      VMSizeTypes = original.StandardM128ms
	StandardM128s       VMSizeTypes = original.StandardM128s
	StandardM6416ms     VMSizeTypes = original.StandardM6416ms
	StandardM6432ms     VMSizeTypes = original.StandardM6432ms
	StandardM64ms       VMSizeTypes = original.StandardM64ms
	StandardM64s        VMSizeTypes = original.StandardM64s
	StandardNC12        VMSizeTypes = original.StandardNC12
	StandardNC12sV2     VMSizeTypes = original.StandardNC12sV2
	StandardNC12sV3     VMSizeTypes = original.StandardNC12sV3
	StandardNC24        VMSizeTypes = original.StandardNC24
	StandardNC24r       VMSizeTypes = original.StandardNC24r
	StandardNC24rsV2    VMSizeTypes = original.StandardNC24rsV2
	StandardNC24rsV3    VMSizeTypes = original.StandardNC24rsV3
	StandardNC24sV2     VMSizeTypes = original.StandardNC24sV2
	StandardNC24sV3     VMSizeTypes = original.StandardNC24sV3
	StandardNC6         VMSizeTypes = original.StandardNC6
	StandardNC6sV2      VMSizeTypes = original.StandardNC6sV2
	StandardNC6sV3      VMSizeTypes = original.StandardNC6sV3
	StandardND12s       VMSizeTypes = original.StandardND12s
	StandardND24rs      VMSizeTypes = original.StandardND24rs
	StandardND24s       VMSizeTypes = original.StandardND24s
	StandardND6s        VMSizeTypes = original.StandardND6s
	StandardNV12        VMSizeTypes = original.StandardNV12
	StandardNV24        VMSizeTypes = original.StandardNV24
	StandardNV6         VMSizeTypes = original.StandardNV6
)

type AccessProfile = original.AccessProfile
type AgentPoolProfile = original.AgentPoolProfile
type ContainerService = original.ContainerService
type ContainerServicesCreateOrUpdateFutureType = original.ContainerServicesCreateOrUpdateFutureType
type ContainerServicesDeleteFutureType = original.ContainerServicesDeleteFutureType
type CustomProfile = original.CustomProfile
type DiagnosticsProfile = original.DiagnosticsProfile
type KeyVaultSecretRef = original.KeyVaultSecretRef
type LinuxProfile = original.LinuxProfile
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type ManagedCluster = original.ManagedCluster
type ManagedClusterAccessProfile = original.ManagedClusterAccessProfile
type ManagedClusterListResult = original.ManagedClusterListResult
type ManagedClusterListResultIterator = original.ManagedClusterListResultIterator
type ManagedClusterListResultPage = original.ManagedClusterListResultPage
type ManagedClusterPoolUpgradeProfile = original.ManagedClusterPoolUpgradeProfile
type ManagedClusterProperties = original.ManagedClusterProperties
type ManagedClustersCreateOrUpdateFuture = original.ManagedClustersCreateOrUpdateFuture
type ManagedClustersDeleteFuture = original.ManagedClustersDeleteFuture
type ManagedClusterUpgradeProfile = original.ManagedClusterUpgradeProfile
type ManagedClusterUpgradeProfileProperties = original.ManagedClusterUpgradeProfileProperties
type MasterProfile = original.MasterProfile
type OrchestratorProfile = original.OrchestratorProfile
type OrchestratorProfileType = original.OrchestratorProfileType
type OrchestratorVersionProfile = original.OrchestratorVersionProfile
type OrchestratorVersionProfileListResult = original.OrchestratorVersionProfileListResult
type OrchestratorVersionProfileProperties = original.OrchestratorVersionProfileProperties
type Properties = original.Properties
type Resource = original.Resource
type ServicePrincipalProfile = original.ServicePrincipalProfile
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type VMDiagnostics = original.VMDiagnostics
type WindowsProfile = original.WindowsProfile

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewContainerServicesClient(subscriptionID string) ContainerServicesClient {
	return original.NewContainerServicesClient(subscriptionID)
}
func NewContainerServicesClientWithBaseURI(baseURI string, subscriptionID string) ContainerServicesClient {
	return original.NewContainerServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedClustersClient(subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClient(subscriptionID)
}
func NewManagedClustersClientWithBaseURI(baseURI string, subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleOrchestratorTypesValues() []OrchestratorTypes {
	return original.PossibleOrchestratorTypesValues()
}
func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}
func PossibleStorageProfileTypesValues() []StorageProfileTypes {
	return original.PossibleStorageProfileTypesValues()
}
func PossibleVMSizeTypesValues() []VMSizeTypes {
	return original.PossibleVMSizeTypesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
