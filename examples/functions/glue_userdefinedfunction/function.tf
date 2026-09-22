# arn:aws:glue:ap-northeast-1:111111111111:userDefinedFunction/database-name/user-defined-function-name
output "glue_userdefinedfunction" {
  value = provider::arn::glue_userdefinedfunction("database-name", "user-defined-function-name")
}
