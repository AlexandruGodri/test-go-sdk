# SnapshotProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] [default to null]
**Description** | **string** | Human readable description | [optional] [default to null]
**Location** | **string** | Location of that image/snapshot.  | [optional] [default to null]
**Size** | **float64** | The size of the image in GB | [optional] [default to null]
**SecAuthProtection** | **bool** | Boolean value representing if the snapshot requires extra protection e.g. two factor protection | [optional] [default to null]
**CpuHotPlug** | **bool** | Is capable of CPU hot plug (no reboot required) | [optional] [default to null]
**CpuHotUnplug** | **bool** | Is capable of CPU hot unplug (no reboot required) | [optional] [default to null]
**RamHotPlug** | **bool** | Is capable of memory hot plug (no reboot required) | [optional] [default to null]
**RamHotUnplug** | **bool** | Is capable of memory hot unplug (no reboot required) | [optional] [default to null]
**NicHotPlug** | **bool** | Is capable of nic hot plug (no reboot required) | [optional] [default to null]
**NicHotUnplug** | **bool** | Is capable of nic hot unplug (no reboot required) | [optional] [default to null]
**DiscVirtioHotPlug** | **bool** | Is capable of Virt-IO drive hot plug (no reboot required) | [optional] [default to null]
**DiscVirtioHotUnplug** | **bool** | Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. | [optional] [default to null]
**DiscScsiHotPlug** | **bool** | Is capable of SCSI drive hot plug (no reboot required) | [optional] [default to null]
**DiscScsiHotUnplug** | **bool** | Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. | [optional] [default to null]
**LicenceType** | **string** | OS type of this Snapshot | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

