package model

type Song struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

type Songs []Song

type Album struct {
	Id     string `json:"id"`
	Title  string `json:"title"`;
	Tracks Songs `json:"tracks"`;
}

type Artist struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Details string `json:"Details"`
	Albums  []Album `json:"Albums"`
	Country string `json:"Country"`
}

type Results []Artist
