# arn:aws:transfer:ap-northeast-1:111111111111:agreement/server-id/agreement-id
output "transfer_agreement" {
  value = provider::arn::transfer_agreement("server-id", "agreement-id")
}
