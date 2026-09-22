# arn:aws:iottwinmaker:ap-northeast-1:111111111111:workspace/workspace-id/scene/scene-id
output "iottwinmaker_scene" {
  value = provider::arn::iottwinmaker_scene("workspace-id", "scene-id")
}
