# arn:aws:cleanrooms:ap-northeast-1:111111111111:membership/membership-id
output "cleanrooms_membership" {
  value = provider::arn::cleanrooms_membership("membership-id")
}
