# arn:aws:observabilityadmin:ap-northeast-1:111111111111:s3tableintegration/s3-table-integration-identifier
output "observabilityadmin_s3tableintegration" {
  value = provider::arn::observabilityadmin_s3tableintegration("s3-table-integration-identifier")
}
