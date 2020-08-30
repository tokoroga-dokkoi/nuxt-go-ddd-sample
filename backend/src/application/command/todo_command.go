package command

type TodoCreateCommand struct {
  Name string `json:"name"`
  Priority string `json:"priority"`
  Descrition string `json:"description"`
}


type TodoUpdateCommand struct {
  Id int `json:"id" query:"id"`
  Name string `json:"name"`
  Priority string `json:"priority"`
  Descrition string `json:"description"`
}
