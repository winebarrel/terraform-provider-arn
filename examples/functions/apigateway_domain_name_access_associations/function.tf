# arn:aws:apigateway:ap-northeast-1:111111111111:/domainnameaccessassociations
output "apigateway_domain_name_access_associations" {
  value = provider::arn::apigateway_domain_name_access_associations()
}
