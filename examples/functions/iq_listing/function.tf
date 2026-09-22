# arn:aws:iq:ap-northeast-1::listing/listing-id
output "iq_listing" {
  value = provider::arn::iq_listing("listing-id")
}
