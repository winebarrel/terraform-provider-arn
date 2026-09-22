# arn:aws:config:ap-northeast-1:111111111111:conformance-pack/conformance-pack-name/conformance-pack-id
output "config_conformance_pack" {
  value = provider::arn::config_conformance_pack("conformance-pack-name", "conformance-pack-id")
}
