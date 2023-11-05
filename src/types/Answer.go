package types

type Answer struct {
	Choice  int    `json:"choice"`
	Title   string `json:"title"`
	Desc    string `json:"description"`
	ImgLink string `json:"imgLink"`
}
