# arn:aws:iam::111111111111:role/aws-service-role/opensearchservice.amazonaws.com/AWSServiceRoleForAmazonOpenSearchService
output "es_opensearchservice_role" {
  value = provider::arn::es_opensearchservice_role()
}
