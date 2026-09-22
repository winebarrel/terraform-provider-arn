# arn:aws:macie2:ap-northeast-1:111111111111:classification-job/resource-id
output "macie2_classification_job" {
  value = provider::arn::macie2_classification_job("resource-id")
}
