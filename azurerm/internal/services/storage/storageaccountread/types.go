package storageaccountread

import (
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// Account the storage account.
type Account struct {
	autorest.Response `json:"-"`
	// Sku - READ-ONLY; Gets the SKU.
	Sku *storage.Sku `json:"sku,omitempty"`
	// Kind - READ-ONLY; Gets the Kind. Possible values include: 'KindStorage', 'KindStorageV2', 'KindBlobStorage', 'KindFileStorage', 'KindBlockBlobStorage'
	Kind storage.Kind `json:"kind,omitempty"`
	// Identity - The identity of the resource.
	Identity *storage.Identity `json:"identity,omitempty"`
	// // ExtendedLocation - The extendedLocation of the resource.
	// ExtendedLocation *storage.ExtendedLocation `json:"extendedLocation,omitempty"`
	// AccountProperties - Properties of the storage account.
	*AccountProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// AccountProperties properties of the storage account.
type AccountProperties struct {
	// ProvisioningState - READ-ONLY; Gets the status of the storage account at the time the operation was called. Possible values include: 'Creating', 'ResolvingDNS', 'Succeeded'
	ProvisioningState storage.ProvisioningState `json:"provisioningState,omitempty"`
	// PrimaryEndpoints - READ-ONLY; Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints *storage.Endpoints `json:"primaryEndpoints,omitempty"`
	// PrimaryLocation - READ-ONLY; Gets the location of the primary data center for the storage account.
	PrimaryLocation *string `json:"primaryLocation,omitempty"`
	// StatusOfPrimary - READ-ONLY; Gets the status indicating whether the primary location of the storage account is available or unavailable. Possible values include: 'Available', 'Unavailable'
	StatusOfPrimary storage.AccountStatus `json:"statusOfPrimary,omitempty"`
	// LastGeoFailoverTime - READ-ONLY; Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime *date.Time `json:"lastGeoFailoverTime,omitempty"`
	// SecondaryLocation - READ-ONLY; Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`
	// StatusOfSecondary - READ-ONLY; Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS. Possible values include: 'Available', 'Unavailable'
	StatusOfSecondary storage.AccountStatus `json:"statusOfSecondary,omitempty"`
	// CreationTime - READ-ONLY; Gets the creation date and time of the storage account in UTC.
	CreationTime *date.Time `json:"creationTime,omitempty"`
	// CustomDomain - READ-ONLY; Gets the custom domain the user assigned to this storage account.
	CustomDomain *storage.CustomDomain `json:"customDomain,omitempty"`
	// SecondaryEndpoints - READ-ONLY; Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints *storage.Endpoints `json:"secondaryEndpoints,omitempty"`
	// Encryption - READ-ONLY; Gets the encryption settings on the account. If unspecified, the account is unencrypted.
	Encryption *storage.Encryption `json:"encryption,omitempty"`
	// AccessTier - READ-ONLY; Required for storage accounts where kind = BlobStorage. The access tier used for billing. Possible values include: 'Hot', 'Cool'
	AccessTier storage.AccessTier `json:"accessTier,omitempty"`
	// AzureFilesIdentityBasedAuthentication - Provides the identity based authentication settings for Azure Files.
	AzureFilesIdentityBasedAuthentication *storage.AzureFilesIdentityBasedAuthentication `json:"azureFilesIdentityBasedAuthentication,omitempty"`
	// EnableHTTPSTrafficOnly - Allows https traffic only to storage service if sets to true.
	EnableHTTPSTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
	// NetworkRuleSet - READ-ONLY; Network rule set
	NetworkRuleSet *storage.NetworkRuleSet `json:"networkAcls,omitempty"`
	// IsHnsEnabled - Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled *bool `json:"isHnsEnabled,omitempty"`
	// GeoReplicationStats - READ-ONLY; Geo Replication Stats
	GeoReplicationStats *storage.GeoReplicationStats `json:"geoReplicationStats,omitempty"`
	// FailoverInProgress - READ-ONLY; If the failover is in progress, the value will be true, otherwise, it will be null.
	FailoverInProgress *bool `json:"failoverInProgress,omitempty"`
	// LargeFileSharesState - Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled. Possible values include: 'Disabled', 'Enabled'
	LargeFileSharesState storage.LargeFileSharesState `json:"largeFileSharesState,omitempty"`

	// AllowBlobPublicAccess - Allow or disallow public access to all blobs or containers in the storage account. The default interpretation is true for this property.
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess,omitempty"`
	// AllowSharedKeyAccess - Indicates whether the storage account permits requests to be authorized with the account access key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess *bool `json:"allowSharedKeyAccess,omitempty"`
}
