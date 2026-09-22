# arn:aws:iq:ap-northeast-1::attachment/attachment-id
output "iq_attachment" {
  value = provider::arn::iq_attachment("attachment-id")
}
