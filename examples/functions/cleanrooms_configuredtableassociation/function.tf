# arn:aws:cleanrooms:ap-northeast-1:111111111111:membership/membership-id/configuredtableassociation/configured-table-association-id
output "cleanrooms_configuredtableassociation" {
  value = provider::arn::cleanrooms_configuredtableassociation("membership-id", "configured-table-association-id")
}
