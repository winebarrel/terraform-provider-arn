# arn:aws:iam::111111111111:role/aws-service-role/es.amazonaws.com/AWSServiceRoleForAmazonOpenSearchService
output "es_es_role" {
  value = provider::arn::es_es_role()
}
