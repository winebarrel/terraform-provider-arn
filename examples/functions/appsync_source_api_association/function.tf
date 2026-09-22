# arn:aws:appsync:ap-northeast-1:111111111111:apis/merged-graph-qlapi-id/sourceApiAssociations/associationid
output "appsync_source_api_association" {
  value = provider::arn::appsync_source_api_association("merged-graph-qlapi-id", "associationid")
}
