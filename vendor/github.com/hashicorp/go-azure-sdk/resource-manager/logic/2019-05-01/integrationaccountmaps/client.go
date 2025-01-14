package integrationaccountmaps

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountMapsClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationAccountMapsClientWithBaseURI(api environments.Api) (*IntegrationAccountMapsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "integrationaccountmaps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationAccountMapsClient: %+v", err)
	}

	return &IntegrationAccountMapsClient{
		Client: client,
	}, nil
}
