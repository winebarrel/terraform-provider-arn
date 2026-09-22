# arn:aws:apigateway:ap-northeast-1::/domainnames/domain-name+domain-identifier/basepathmappings/base-path
output "apigateway_private_base_path_mapping" {
  value = provider::arn::apigateway_private_base_path_mapping("domain-name", "domain-identifier", "base-path")
}
