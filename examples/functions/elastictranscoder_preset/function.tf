# arn:aws:elastictranscoder:ap-northeast-1:111111111111:preset/preset-id
output "elastictranscoder_preset" {
  value = provider::arn::elastictranscoder_preset("preset-id")
}
