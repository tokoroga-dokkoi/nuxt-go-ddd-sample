package response

import (
  "time"
)

type TodoResponse struct {
  ID uint `json:"id"`
  Name string `json:"name"`
  Priority string `json:"priority"`
  Description string `json:"description"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}
