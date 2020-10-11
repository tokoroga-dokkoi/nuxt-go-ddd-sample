<template>
  <div class="mt-20 text-center signup-form">
    <template v-if="!isSuccess">
      <h2 class="text-5xl font-bold">ユーザ登録</h2>
      <p class="font-bold text-black mt-2">
        サブテキストサブテキストサブテキストサブテキスト
      </p>
      <div v-if="isError" class="bg-red-600 font-bold text-white mt-3">
        <ul class="list-disc">
          <li v-for="(mes, index) in errors" :key="index">
            {{ mes }}
          </li>
        </ul>
      </div>
      <sign-up-form
        class="mt-4"
        :email.sync="email"
        @signup-email="signUpWithEmail"
        @signup-google="signUpWithGoogle"
        @signup-facebook="signUpWithFacebook"
      />
    </template>
    <template v-else>
      <h2 class="text-5xl font-bold">仮登録が完了しました</h2>
      <div class="font-bold text-black mt-2">
        <p>{{ email }}宛に本登録用のURLを送らせていただきました。</p>
        <p>メールに記載されたURLから、本登録を行ってください</p>
      </div>
    </template>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
import * as firebase from 'firebase'
import 'firebase/auth'
import SignUpForm from '@/components/organisms/signUpForm.vue'

export default Vue.extend({
  components: {
    SignUpForm,
  },
  fetch(context) {
    console.log(context.store.state)
  },
  data() {
    return {
      email: '',
      isError: false,
      isSuccess: false,
      errors: [],
    }
  },
  mounted() {
    this.$auth.onAuthStateChanged(user => {
      // 登録完了のタイミング
      if (user) {
        console.log(user)
        // email検証
        if (this.$auth.currentUser.emailVerified) {
          this.$auth.currentUser.sendEmailVerification()
        }
      } else {
        console.error('hogehoge')
      }
    })
  },
  methods: {
    // Eメールでの仮登録
    async signUpWithEmail() {
      this.isError = false
      this.errors = []
      try {
        await this.$axios.post('/v1/users/signup_request', {
          email: this.email,
        })
        this.isSuccess = true
      } catch (e) {
        this.errors = e.response.data.errors
        this.isError = true
      }
    },
    signUpWithGoogle() {
      const provider = new firebase.auth.GoogleAuthProvider()
      this.$auth.signInWithRedirect(provider)
    },
    signUpWithFacebook() {
      const provider = new firebase.auth.FacebookAuthProvider()
      this.$auth.signInWithRedirect(provider)
    },
  },
})
</script>
<style lang="scss" scoped>
.signup-form {
  max-width: 560px;
  margin: 60px auto;
}
</style>
