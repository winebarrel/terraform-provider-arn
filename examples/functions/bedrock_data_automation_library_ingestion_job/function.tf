# arn:aws:bedrock:ap-northeast-1:111111111111:data-automation-library-ingestion-job/ingestion-job-id
output "bedrock_data_automation_library_ingestion_job" {
  value = provider::arn::bedrock_data_automation_library_ingestion_job("ingestion-job-id")
}
