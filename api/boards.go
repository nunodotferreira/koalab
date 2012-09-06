package main

import (
	"fmt"
  "encoding/json"
	"net/http"
  "io/ioutil"
)


// == Models

type Board struct {
  Id string
  Postits []*Postit
  // FIXME : h|vrules
}

func (b *Board) Unmarshal(bytes []byte) error {
  return json.Unmarshal(bytes, &b)
}

func (b *Board) toJSON() ([]byte, error) {
  return json.Marshal(&b)
}

type BoardCollection struct {
  Boards []Board
}

func (bc *BoardCollection) addBoard(board *Board) {
  bc.Boards = append(boards.Boards, *board)
}

func (bc *BoardCollection) toJSON() ([]byte, error) {
  return json.Marshal(&bc)
}

func (bc *BoardCollection) findById(id string) *Board {
  for _, b := range bc.Boards {
    if b.Id == id {
      return &b
    }
  }
  return nil
}

// == Controller

var boards = new(BoardCollection)

func ListBoards(w http.ResponseWriter, req *http.Request) {
  bytes, _ := boards.toJSON()
  fmt.Fprintf(w, "%s", bytes)
}

func ShowBoard(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":Id")
  board := boards.findById(id)
  if board != nil {
    bytes, _ := (*board).toJSON()
    fmt.Fprintf(w, "%s", bytes)
  }
}

func CreateBoard(w http.ResponseWriter, req *http.Request) {
	board := new(Board)
  bytes, _ := ioutil.ReadAll(req.Body)
  err := board.Unmarshal(bytes)
  if err == nil {
    boards.addBoard(board)
  } else {
    fmt.Println("%s", err)
  }
}