# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/use-case/use-case-id
output "connect_use_case" {
  value = provider::arn::connect_use_case("instance-id", "use-case-id")
}
