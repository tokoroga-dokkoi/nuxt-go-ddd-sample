package testsignuprequest

import (
	"testing"

	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/application/usecase/signup_request"
	domain_model_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/model/users"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/repository"
	domain_service_users "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/domain/service/users"
	"github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/infra/mysql"
	testshared "github.com/MikiWaraMiki/nuxt-go-ddd-sample/backend/src/spec/shared"

	"github.com/bxcodec/faker/v3"
)

type testDBUS struct {
	repo repository.IUserEmailVerificationRepository
	ds   *domain_service_users.EmailVerificationService
	uc   signup_request.IUserAuthSignUpRequestUsecase
}

// newTestRepo は、test用の仮登録リポジトリを返す
func newTestRepo() repository.IUserEmailVerificationRepository {
	conn := testshared.NewTestDb()
	repo := mysql.NewUserEmailVerificationRepository(conn)
	return repo
}

// newTestModule はUserSignUpreRequestUsecaseを返す
func newTestModule() *testDBUS {
	repo := newTestRepo()
	ds := domain_service_users.NewEmailVerificationService(repo)
	uc := signup_request.NewUserSignUpRequestUsecase(&repo, ds)

	testObj := testDBUS{
		repo: repo,
		ds:   ds,
		uc:   uc,
	}

	return &testObj
}

// parameterが不足いないかつ、登録済ではない場合は、ユーザが登録できる
func TestIsSuccessWhenParameterIsValidAndUserNotExist(t *testing.T) {
	mod := newTestModule()
	command := signup_request.UserSignUpRequestInputCommand{
		Email: faker.Email(),
	}

	result, err := mod.uc.SignUpRequest(&command)

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	resEmail, _ := domain_model_users.NewEmail(result.Email)
	isRegistered := mod.ds.Exist(resEmail)

	if !isRegistered {
		t.Errorf("ユーザが登録されていませんでした email: %s", result.Email)
	}
}
