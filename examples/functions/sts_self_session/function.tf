# arn:aws:sts::111111111111:self
output "sts_self_session" {
  value = provider::arn::sts_self_session()
}
