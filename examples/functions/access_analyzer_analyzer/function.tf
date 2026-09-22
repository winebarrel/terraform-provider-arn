# arn:aws:access-analyzer:ap-northeast-1:111111111111:analyzer/analyzer-name
output "access_analyzer_analyzer" {
  value = provider::arn::access_analyzer_analyzer("analyzer-name")
}
