# arn:aws:mediatailor:ap-northeast-1:111111111111:program/channel-name/program-name
output "mediatailor_program" {
  value = provider::arn::mediatailor_program("channel-name", "program-name")
}
