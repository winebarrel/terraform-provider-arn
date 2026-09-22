# arn:aws:tnb:ap-northeast-1:111111111111:function-package/function-package-id
output "tnb_function_package" {
  value = provider::arn::tnb_function_package("function-package-id")
}
