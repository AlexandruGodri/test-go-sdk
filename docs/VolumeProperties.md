# VolumeProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] [default to null]
**Type_** | **string** | Hardware type of the volume. Default is HDD | [optional] [default to null]
**Size** | **float64** | The size of the volume in GB | [default to null]
**AvailabilityZone** | **string** | The availability zone in which the volume should exist. The storage volume will be provisioned on as less physical storages as possible but cannot guarantee upfront | [optional] [default to null]
**Image** | **string** | Image or snapshot ID to be used as template for this volume | [optional] [default to null]
**ImagePassword** | **string** | Initial password to be set for installed OS. Works with public images only. Not modifiable, forbidden in update requests. Password rules allows all characters from a-z, A-Z, 0-9 | [optional] [default to null]
**SshKeys** | **[]string** | Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation. | [optional] [default to null]
**Bus** | **string** | The bus type of the volume. Default is VIRTIO | [optional] [default to null]
**LicenceType** | **string** | OS type of this volume | [optional] [default to null]
**CpuHotPlug** | **bool** | Is capable of CPU hot plug (no reboot required) | [optional] [default to null]
**RamHotPlug** | **bool** | Is capable of memory hot plug (no reboot required) | [optional] [default to null]
**NicHotPlug** | **bool** | Is capable of nic hot plug (no reboot required) | [optional] [default to null]
**NicHotUnplug** | **bool** | Is capable of nic hot unplug (no reboot required) | [optional] [default to null]
**DiscVirtioHotPlug** | **bool** | Is capable of Virt-IO drive hot plug (no reboot required) | [optional] [default to null]
**DiscVirtioHotUnplug** | **bool** | Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. | [optional] [default to null]
**DeviceNumber** | **int64** | The LUN ID of the storage volume. Null for volumes not mounted to any VM | [optional] [default to null]
**BackupunitId** | **string** | The uuid of the Backup Unit that user has access to. The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provied either public image or imageAlias in conjunction with this property. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

