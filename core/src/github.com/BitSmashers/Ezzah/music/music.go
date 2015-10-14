package music

type Song struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Artist string `json:"artist"`
}

type Songs []Song

type Album struct {
  Id string `json:"id"`
  Title string `json:"title"`;
  Tracks Songs `json:"tracks"`;
}

type Artist struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Details string `json:"details"`
  Albums []Album `json:"albums"`
  Country string `json:"country"`
}

type Results []Artist
