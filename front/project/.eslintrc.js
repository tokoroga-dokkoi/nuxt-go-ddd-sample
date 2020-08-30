module.exports = {
  root: true,
  env: {
    browser: true,
    node: true,
  },
  extends: [
    '@nuxtjs/eslint-config-typescript',
    'prettier',
    'prettier/vue',
    'plugin:prettier/recommended',
    'plugin:nuxt/recommended',
  ],
  plugins: ['prettier'],
  // add your custom rules here
  rules: {
    // 実運用に合わせて追加していく
    'no-plusplus': 'off',
    'func-names': 'off',
    'vue/component-name-in-template-casing': 'off',
    'no-console': 'error',
    'no-debugger': 'error',
    'import/prefer-default-export': 'off', // default exportだと変数・関数が使えないのでOFF
    'import/no-unresolved': 'off', // Vueの@とか~が使えなくなるため追加
    // セミコロンを省略する
    semi: ['error', 'never', { beforeStatementContinuationChars: 'never' }],
    'semi-spacing': ['error', { after: true, before: false }],
    'semi-style': ['error', 'first'],
    'no-extra-semi': 'error',
    'no-unexpected-multiline': 'error',
    'no-unreachable': 'error',
    // 文字列はシングルクォートにする
    quotes: ['warn', 'single'],
  },
}
