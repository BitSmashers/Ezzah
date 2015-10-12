package music

type Song struct {
  Id string `json:"id"`
  Title string `json:"title"`
}

type Songs []Song

type Album struct {
  Id string `json:"id"`
  Title string `json:"title"`;
  Tracks Songs `json:"tracks"`;
}

type Artist struct {
  Id string `json:"id"`
  Name string `json:"name"`;
  Albums []Album `json:"albums"`;
}

type Results []Artist
