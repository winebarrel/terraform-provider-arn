# arn:aws:mgn:ap-northeast-1:111111111111:wave/wave-id
output "mgn_wave_resource" {
  value = provider::arn::mgn_wave_resource("wave-id")
}
