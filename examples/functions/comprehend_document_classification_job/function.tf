# arn:aws:comprehend:ap-northeast-1:111111111111:document-classification-job/job-id
output "comprehend_document_classification_job" {
  value = provider::arn::comprehend_document_classification_job("job-id")
}
