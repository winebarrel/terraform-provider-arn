// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: wisdom
// Source: https://servicereference.us-east-1.amazonaws.com/v1/wisdom/wisdom.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "wisdom_ai_agent", Service: "wisdom", Resource: "AIAgent", Template: "arn:${Partition}:wisdom:${Region}:${Account}:ai-agent/${AssistantId}/${AIAgentId}"},
		{Name: "wisdom_ai_guardrail", Service: "wisdom", Resource: "AIGuardrail", Template: "arn:${Partition}:wisdom:${Region}:${Account}:ai-guardrail/${AssistantId}/${AIGuardrailId}"},
		{Name: "wisdom_ai_prompt", Service: "wisdom", Resource: "AIPrompt", Template: "arn:${Partition}:wisdom:${Region}:${Account}:ai-prompt/${AssistantId}/${AIPromptId}"},
		{Name: "wisdom_assistant", Service: "wisdom", Resource: "Assistant", Template: "arn:${Partition}:wisdom:${Region}:${Account}:assistant/${AssistantId}"},
		{Name: "wisdom_assistant_association", Service: "wisdom", Resource: "AssistantAssociation", Template: "arn:${Partition}:wisdom:${Region}:${Account}:association/${AssistantId}/${AssistantAssociationId}"},
		{Name: "wisdom_content", Service: "wisdom", Resource: "Content", Template: "arn:${Partition}:wisdom:${Region}:${Account}:content/${KnowledgeBaseId}/${ContentId}"},
		{Name: "wisdom_content_association", Service: "wisdom", Resource: "ContentAssociation", Template: "arn:${Partition}:wisdom:${Region}:${Account}:content-association/${KnowledgeBaseId}/${ContentId}/${ContentAssociationId}"},
		{Name: "wisdom_knowledge_base", Service: "wisdom", Resource: "KnowledgeBase", Template: "arn:${Partition}:wisdom:${Region}:${Account}:knowledge-base/${KnowledgeBaseId}"},
		{Name: "wisdom_message_template", Service: "wisdom", Resource: "MessageTemplate", Template: "arn:${Partition}:wisdom:${Region}:${Account}:message-template/${KnowledgeBaseId}/${MessageTemplateId}"},
		{Name: "wisdom_quick_response", Service: "wisdom", Resource: "QuickResponse", Template: "arn:${Partition}:wisdom:${Region}:${Account}:quick-response/${KnowledgeBaseId}/${QuickResponseId}"},
		{Name: "wisdom_session", Service: "wisdom", Resource: "Session", Template: "arn:${Partition}:wisdom:${Region}:${Account}:session/${AssistantId}/${SessionId}"},
	})
}
