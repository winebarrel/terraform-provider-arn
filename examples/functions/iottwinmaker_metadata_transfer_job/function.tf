# arn:aws:iottwinmaker:ap-northeast-1:111111111111:metadata-transfer-job/metadata-transfer-job-id
output "iottwinmaker_metadata_transfer_job" {
  value = provider::arn::iottwinmaker_metadata_transfer_job("metadata-transfer-job-id")
}
