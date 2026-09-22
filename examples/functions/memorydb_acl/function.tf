# arn:aws:memorydb:ap-northeast-1:111111111111:acl/acl-name
output "memorydb_acl" {
  value = provider::arn::memorydb_acl("acl-name")
}
