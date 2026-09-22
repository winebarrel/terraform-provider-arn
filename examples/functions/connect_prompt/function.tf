# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/prompt/prompt-id
output "connect_prompt" {
  value = provider::arn::connect_prompt("instance-id", "prompt-id")
}
