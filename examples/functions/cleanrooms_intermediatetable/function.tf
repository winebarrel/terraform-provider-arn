# arn:aws:cleanrooms:ap-northeast-1:111111111111:membership/membership-id/intermediatetable/intermediate-table-id
output "cleanrooms_intermediatetable" {
  value = provider::arn::cleanrooms_intermediatetable("membership-id", "intermediate-table-id")
}
