# arn:aws:s3:ap-northeast-1:111111111111:storage-lens/config-id
output "s3_storagelensconfiguration" {
  value = provider::arn::s3_storagelensconfiguration("config-id")
}
