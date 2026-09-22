# arn:aws:access-analyzer:ap-northeast-1:111111111111:analyzer/analyzer-name/archive-rule/rule-name
output "access_analyzer_archive_rule" {
  value = provider::arn::access_analyzer_archive_rule("analyzer-name", "rule-name")
}
