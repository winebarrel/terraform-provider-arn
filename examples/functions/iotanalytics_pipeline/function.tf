# arn:aws:iotanalytics:ap-northeast-1:111111111111:pipeline/pipeline-name
output "iotanalytics_pipeline" {
  value = provider::arn::iotanalytics_pipeline("pipeline-name")
}
