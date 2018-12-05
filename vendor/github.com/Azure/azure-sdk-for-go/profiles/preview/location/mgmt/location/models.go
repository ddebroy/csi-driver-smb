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

package location

import original "github.com/Azure/azure-sdk-for-go/services/location/mgmt/2017-01-01-preview/location"

type AccountsClient = original.AccountsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type KeyType = original.KeyType

const (
	Primary   KeyType = original.Primary
	Secondary KeyType = original.Secondary
)

type BasedServicesAccount = original.BasedServicesAccount
type BasedServicesAccountCreateParameters = original.BasedServicesAccountCreateParameters
type BasedServicesAccountKeys = original.BasedServicesAccountKeys
type BasedServicesAccounts = original.BasedServicesAccounts
type BasedServicesAccountsMoveRequest = original.BasedServicesAccountsMoveRequest
type BasedServicesAccountUpdateParameters = original.BasedServicesAccountUpdateParameters
type BasedServicesKeySpecification = original.BasedServicesKeySpecification
type BasedServicesOperations = original.BasedServicesOperations
type BasedServicesOperationsValueItem = original.BasedServicesOperationsValueItem
type BasedServicesOperationsValueItemDisplay = original.BasedServicesOperationsValueItemDisplay
type Error = original.Error
type ErrorDetailsItem = original.ErrorDetailsItem
type Resource = original.Resource
type Sku = original.Sku

func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
