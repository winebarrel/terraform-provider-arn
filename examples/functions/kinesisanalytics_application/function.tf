# arn:aws:kinesisanalytics:ap-northeast-1:111111111111:application/application-name
output "kinesisanalytics_application" {
  value = provider::arn::kinesisanalytics_application("application-name")
}
