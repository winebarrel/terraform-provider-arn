# arn:aws:comprehend:ap-northeast-1:111111111111:document-classifier-endpoint/document-classifier-endpoint-name
output "comprehend_document_classifier_endpoint" {
  value = provider::arn::comprehend_document_classifier_endpoint("document-classifier-endpoint-name")
}
