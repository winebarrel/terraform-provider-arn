# arn:aws:apigateway:ap-northeast-1::/domainnames/domain-name+domain-identifier/basepathmappings
output "apigateway_private_base_path_mappings" {
  value = provider::arn::apigateway_private_base_path_mappings("domain-name", "domain-identifier")
}
