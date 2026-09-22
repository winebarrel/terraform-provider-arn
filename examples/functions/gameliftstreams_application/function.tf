# arn:aws:gameliftstreams:ap-northeast-1:111111111111:application/application-id
output "gameliftstreams_application" {
  value = provider::arn::gameliftstreams_application("application-id")
}
