package migrate

import (
  "os"
  "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/todos"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/bxcodec/faker/v3"
)

func SetUp() {
  name := todos.NewName(faker.Word())
  priority := todos.NewPriority("high")
  description := todos.NewDescription(faker.paragraph(240))

  todo := todos.New(name, priority, description)


}
