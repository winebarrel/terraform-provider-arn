# arn:aws:lightsail:ap-northeast-1:111111111111:KeyPair/id
output "lightsail_key_pair" {
  value = provider::arn::lightsail_key_pair("id")
}
