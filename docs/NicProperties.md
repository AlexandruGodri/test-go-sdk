# NicProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] [default to null]
**Mac** | **string** | The MAC address of the NIC | [optional] [default to null]
**Ips** | **[]string** | Collection of IP addresses assigned to a nic. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically. | [optional] [default to null]
**Dhcp** | **bool** | Indicates if the nic will reserve an IP using DHCP | [optional] [default to null]
**Lan** | **int32** | The LAN ID the NIC will sit on. If the LAN ID does not exist it will be implicitly created | [default to null]
**FirewallActive** | **bool** | Activate or deactivate the firewall. By default an active firewall without any defined rules will block all incoming network traffic except for the firewall rules that explicitly allows certain protocols, ip addresses and ports. | [optional] [default to null]
**Nat** | **bool** | Indicates if NAT is enabled on this NIC | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

