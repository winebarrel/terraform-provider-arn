# arn:aws:textract:ap-northeast-1:111111111111:/adapters/adapter-id/versions/adapter-version
output "textract_adapterversion" {
  value = provider::arn::textract_adapterversion("adapter-id", "adapter-version")
}
