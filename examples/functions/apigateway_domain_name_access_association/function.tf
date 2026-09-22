# arn:aws:apigateway:ap-northeast-1:111111111111:/domainnameaccessassociations/domainname/domain-name/source-type/source-id
output "apigateway_domain_name_access_association" {
  value = provider::arn::apigateway_domain_name_access_association("domain-name", "source-type", "source-id")
}
