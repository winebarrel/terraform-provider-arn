# arn:aws:apigateway:ap-northeast-1::/domainnames/domain-name/basepathmappings/base-path
output "apigateway_base_path_mapping" {
  value = provider::arn::apigateway_base_path_mapping("domain-name", "base-path")
}
