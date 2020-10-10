package mysql

// MySQLError は、Repository層のエラーオブジェクト
type mySQLError struct {
	// ログに表示するメッセージ
	message string
	// エラー発生下のerrorオブジェクト
	originalError error
}

// NewMySQLError はMySQLErrorのコンストラクタ
func NewMySQLError(message string, originalError error) *mySQLError {
	errObj := mySQLError{
		message:       message,
		originalError: originalError,
	}

	return &errObj
}

// Internal は内部エラー(500)であることを示す
func (e *mySQLError) Internal() bool {
	return true
}

// Error は、ログ表示用のエラーメッセージを返す
func (e *mySQLError) Error() string {
	return e.message + ":" + e.originalError.Error()
}
