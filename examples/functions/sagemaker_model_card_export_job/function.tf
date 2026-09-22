# arn:aws:sagemaker:ap-northeast-1:111111111111:model-card/model-card-name/export-job/export-job-name
output "sagemaker_model_card_export_job" {
  value = provider::arn::sagemaker_model_card_export_job("model-card-name", "export-job-name")
}
