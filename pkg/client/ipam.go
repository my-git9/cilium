// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package client

import (
	"github.com/cilium/cilium/api/v1/client/ipam"
	"github.com/cilium/cilium/api/v1/models"
	"github.com/cilium/cilium/pkg/api"
)

const (
	AddressFamilyIPv6 = "ipv6"
	AddressFamilyIPv4 = "ipv4"

	// FIXME: merge with respective consts in pkg/ipam or extract to common shared package?
	PoolPodCIDR     = ""
	PoolMultihoming = "multihoming"
)

// IPAMAllocate allocates an IP address out of address family specific pool.
func (c *Client) IPAMAllocate(family, owner, pool string, expiration bool) (*models.IPAMResponse, error) {
	params := ipam.NewPostIpamParams().WithTimeout(api.ClientTimeout)

	if family != "" {
		params.SetFamily(&family)
	}
	if owner != "" {
		params.SetOwner(&owner)
	}
	if pool != "" {
		params.SetPool(&pool)
	}
	params.SetExpiration(&expiration)

	resp, err := c.Ipam.PostIpam(params)
	if err != nil {
		return nil, Hint(err)
	}
	return resp.Payload, nil
}

// IPAMAllocateIP tries to allocate a particular IP address.
func (c *Client) IPAMAllocateIP(ip, owner, pool string) error {
	params := ipam.NewPostIpamIPParams().WithIP(ip).WithOwner(&owner).WithPool(&pool).WithTimeout(api.ClientTimeout)
	_, err := c.Ipam.PostIpamIP(params)
	return Hint(err)
}

// IPAMReleaseIP releases a IP address back to the pool.
func (c *Client) IPAMReleaseIP(ip, pool string) error {
	params := ipam.NewDeleteIpamIPParams().WithIP(ip).WithPool(&pool).WithTimeout(api.ClientTimeout)
	_, err := c.Ipam.DeleteIpamIP(params)
	return Hint(err)
}
