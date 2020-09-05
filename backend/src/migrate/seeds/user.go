package migrate

import (
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/user"
	"github.com/bxcodec/faker/v3"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetUp() {
	email := user.NewEmail(faker.Email())
	lastName := user.NewLastName(faker.LastName())
	firstName := user.NewFirstName(faker.FirstName())
	displayName := user.NewDisplayName(faker.UserName())

	user := user.New(email, lastName, firstName, displayName)
}
