# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/channel-handshake/identifier
output "partnercentral_channel_handshake" {
  value = provider::arn::partnercentral_channel_handshake("catalog", "identifier")
}
