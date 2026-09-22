# arn:aws:fsx:ap-northeast-1:111111111111:storage-virtual-machine/file-system-id/storage-virtual-machine-id
output "fsx_storage_virtual_machine" {
  value = provider::arn::fsx_storage_virtual_machine("file-system-id", "storage-virtual-machine-id")
}
