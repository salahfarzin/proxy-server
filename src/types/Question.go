package types

type Question struct {
	TestId       int      `json:"test_id"`
	Type         string   `json:"type"`
	Number       int      `json:"number"`
	Title        string   `json:"title"`
	QuestionType string   `json:"questionType"`
	Mandatory    int8     `json:"mandatory"`
	Desc         string   `json:"description"`
	ImgLink      string   `json:"imgLink"`
	Answers      []Answer `json:"answers"`
}
