# arn:aws:ssm:ap-northeast-1:111111111111:windowtarget/window-target-id
output "ssm_windowtarget" {
  value = provider::arn::ssm_windowtarget("window-target-id")
}
