# arn:aws:appsync:ap-northeast-1:111111111111:apis/source-graph-qlapi-id/mergedApiAssociations/associationid
output "appsync_merged_api_association" {
  value = provider::arn::appsync_merged_api_association("source-graph-qlapi-id", "associationid")
}
