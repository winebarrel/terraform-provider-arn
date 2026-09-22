# arn:aws:comprehend:ap-northeast-1:111111111111:document-classifier/document-classifier-name
output "comprehend_document_classifier" {
  value = provider::arn::comprehend_document_classifier("document-classifier-name")
}
