package testsignuprequest

import (
	"fmt"
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
	conn *mysql.SQLHandler
	repo repository.IUserEmailVerificationRepository
	ds   *domain_service_users.EmailVerificationService
	uc   signup_request.IUserAuthSignUpRequestUsecase
}

// newTestRepo は、test用の仮登録リポジトリを返す
func newTestRepo() (*mysql.SQLHandler, repository.IUserEmailVerificationRepository) {
	conn := testshared.NewTestDb()
	repo := mysql.NewUserEmailVerificationRepository(conn)
	return conn, repo
}

// newTestModule はUserSignUpreRequestUsecaseを返す
func newTestModule() *testDBUS {
	conn, repo := newTestRepo()
	ds := domain_service_users.NewEmailVerificationService(repo)
	uc := signup_request.NewUserSignUpRequestUsecase(repo, ds)

	testObj := testDBUS{
		conn: conn,
		repo: repo,
		ds:   ds,
		uc:   uc,
	}

	return &testObj
}

// 共通処理
func dbCleanUp(mod *testDBUS) {
	if mod != nil && mod.conn != nil {
		fmt.Println("db truncate and closed")
		testshared.CloseTestDbConn(mod.conn)
	}

	fmt.Println("clean up complete")
}

// parameterに不足がないかつ、登録済ではない場合は、ユーザが登録できる
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

	t.Cleanup(func() {
		dbCleanUp(mod)
	})
}

// emailが入力されていない場合は、ユーザ登録ができない
func TestIsFailedWhenEmailIsEmpty(t *testing.T) {
	mod := newTestModule()
	command := signup_request.UserSignUpRequestInputCommand{
		Email: "",
	}

	_, err := mod.uc.SignUpRequest(&command)

	if err == nil {
		t.Errorf("メールアドレスが空で登録ができました. Email: %s", command.Email)
	}

	t.Cleanup(func() {
		dbCleanUp(mod)
	})
}

// emailの入力形式が不正な場合は、ユーザ登録ができない
func TestIsFailedWhenEmailIsInValidPattern(t *testing.T) {
	mod := newTestModule()
	command := signup_request.UserSignUpRequestInputCommand{
		Email: "hogehoge",
	}

	_, err := mod.uc.SignUpRequest(&command)

	if err == nil {
		t.Errorf("メールアドレスが不正な形式で登録ができました. Email: %s", command.Email)
	}

	t.Cleanup(func() {
		dbCleanUp(mod)
	})
}

// すでに登録済の場合は、登録ができない
func TestIsFailedWhenUserIsAlreadyRegistered(t *testing.T) {
	mod := newTestModule()
	command := signup_request.UserSignUpRequestInputCommand{
		Email: faker.Email(),
	}

	// 1回目の登録
	mod.uc.SignUpRequest(&command)

	// 2回目の登録
	_, err := mod.uc.SignUpRequest(&command)

	if err == nil {
		t.Errorf("ユーザの重複登録ができました。 email: %s", command.Email)
	}

	t.Cleanup(func() {
		dbCleanUp(mod)
	})
}
