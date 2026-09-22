# arn:aws:chime:ap-northeast-1:111111111111:media-insights-pipeline-configuration/configuration-name
output "chime_media_insights_pipeline_configuration" {
  value = provider::arn::chime_media_insights_pipeline_configuration("configuration-name")
}
