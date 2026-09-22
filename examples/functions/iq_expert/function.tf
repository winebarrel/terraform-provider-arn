# arn:aws:iq:ap-northeast-1::expert/expert-id
output "iq_expert" {
  value = provider::arn::iq_expert("expert-id")
}
