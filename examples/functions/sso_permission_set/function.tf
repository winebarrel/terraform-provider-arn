# arn:aws:sso:::permissionSet/instance-id/permission-set-id
output "sso_permission_set" {
  value = provider::arn::sso_permission_set("instance-id", "permission-set-id")
}
