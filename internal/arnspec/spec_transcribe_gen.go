// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: transcribe
// Source: https://servicereference.us-east-1.amazonaws.com/v1/transcribe/transcribe.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "transcribe_callanalyticscategory", Service: "transcribe", Resource: "callanalyticscategory", Template: "arn:${Partition}:transcribe:${Region}:${Account}:analytics-category/${CategoryName}"},
		{Name: "transcribe_callanalyticsjob", Service: "transcribe", Resource: "callanalyticsjob", Template: "arn:${Partition}:transcribe:${Region}:${Account}:analytics/${JobName}"},
		{Name: "transcribe_languagemodel", Service: "transcribe", Resource: "languagemodel", Template: "arn:${Partition}:transcribe:${Region}:${Account}:language-model/${ModelName}"},
		{Name: "transcribe_medicalscribejob", Service: "transcribe", Resource: "medicalscribejob", Template: "arn:${Partition}:transcribe:${Region}:${Account}:medical-scribe-job/${JobName}"},
		{Name: "transcribe_medicaltranscriptionjob", Service: "transcribe", Resource: "medicaltranscriptionjob", Template: "arn:${Partition}:transcribe:${Region}:${Account}:medical-transcription-job/${JobName}"},
		{Name: "transcribe_medicalvocabulary", Service: "transcribe", Resource: "medicalvocabulary", Template: "arn:${Partition}:transcribe:${Region}:${Account}:medical-vocabulary/${VocabularyName}"},
		{Name: "transcribe_transcriptionjob", Service: "transcribe", Resource: "transcriptionjob", Template: "arn:${Partition}:transcribe:${Region}:${Account}:transcription-job/${JobName}"},
		{Name: "transcribe_vocabulary", Service: "transcribe", Resource: "vocabulary", Template: "arn:${Partition}:transcribe:${Region}:${Account}:vocabulary/${VocabularyName}"},
		{Name: "transcribe_vocabularyfilter", Service: "transcribe", Resource: "vocabularyfilter", Template: "arn:${Partition}:transcribe:${Region}:${Account}:vocabulary-filter/${VocabularyFilterName}"},
	})
}
