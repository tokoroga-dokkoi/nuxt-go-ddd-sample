-- ユーザ仮登録用テーブルの作成
-- +migrate Up
CREATE TABLE user_email_verifications (
  `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `email` VARCHAR(200) NOT NULL COMMENT '仮登録時に利用したメールアドレス',
  `status` VARCHAR(100) NOT NULL COMMENT '仮登録状態のステータス(waiting_registration:本登録待ち registration: 登録済 expired: 期限ぎれ)',
  `registration_url_token` VARCHAR(300) NOT NULL COMMENT '認証用URLトークン',
  `registration_email_sent_at` DATETIME COMMENT '認証用URLトークンを送った日付',
  `url_token_expired_at` DATETIME COMMENT '認証用URLトークンの有効期限'
) ENGINE = InnoDB DEFAULT CHARSET=utf8;
CREATE UNIQUE INDEX index_of_unique_user_email_verifications_on_email ON user_email_verifications(email);
-- +migrate Down
DROP TABLE user_email_verifications;
