# arn:aws:iq:ap-northeast-1::seller/seller-aws-account-id
output "iq_seller" {
  value = provider::arn::iq_seller("seller-aws-account-id")
}
