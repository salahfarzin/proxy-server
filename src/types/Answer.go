package types

type Answer struct {
	Choice  int    `json:"choice"`
	Title   string `json:"title"`
	Type    string `json:"type"`
	Desc    string `json:"description"`
	ImgLink string `json:"imgLink"`
}
