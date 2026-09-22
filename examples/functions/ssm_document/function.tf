# arn:aws:ssm:ap-northeast-1:111111111111:document/document-name
output "ssm_document" {
  value = provider::arn::ssm_document("document-name")
}
