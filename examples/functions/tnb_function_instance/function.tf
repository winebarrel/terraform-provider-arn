# arn:aws:tnb:ap-northeast-1:111111111111:function-instance/function-instance-id
output "tnb_function_instance" {
  value = provider::arn::tnb_function_instance("function-instance-id")
}
