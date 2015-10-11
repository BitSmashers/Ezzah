package main

type Song struct {
  title string `json:"title"`
}

type Album struct {
  title string `json:"title"`
  songs []Song `json:"songs"`
}

type Artist struct {
  name string `json:"name"`
  albums []Album `json:"albums"`
}

type Results []Artist
