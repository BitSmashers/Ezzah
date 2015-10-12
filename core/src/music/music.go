package music

type Song struct {
  Title string `json:"title"`;
}

type Songs []Song

type Album struct {
  Title string `json:"title"`;
  Songs `json:"songs"`;
}

type Artist struct {
  Name string `json:"name"`;
  Albums []Album `json:"albums"`;
}

type Results []Artist
