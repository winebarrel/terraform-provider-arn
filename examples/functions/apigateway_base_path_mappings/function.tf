# arn:aws:apigateway:ap-northeast-1::/domainnames/domain-name/basepathmappings
output "apigateway_base_path_mappings" {
  value = provider::arn::apigateway_base_path_mappings("domain-name")
}
