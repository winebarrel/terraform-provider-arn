# arn:aws:bedrock-mantle:ap-northeast-1:111111111111:reservation/resource-id
output "bedrock_mantle_reservation" {
  value = provider::arn::bedrock_mantle_reservation("resource-id")
}
