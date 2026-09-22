# arn:aws:textract:ap-northeast-1:111111111111:/adapters/adapter-id
output "textract_adapter" {
  value = provider::arn::textract_adapter("adapter-id")
}
