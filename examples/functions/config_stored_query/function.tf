# arn:aws:config:ap-northeast-1:111111111111:stored-query/stored-query-name/stored-query-id
output "config_stored_query" {
  value = provider::arn::config_stored_query("stored-query-name", "stored-query-id")
}
