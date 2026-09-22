# arn:aws:cleanrooms:ap-northeast-1:111111111111:membership/membership-id/idmappingtable/id-mapping-table-id
output "cleanrooms_idmappingtable" {
  value = provider::arn::cleanrooms_idmappingtable("membership-id", "id-mapping-table-id")
}
