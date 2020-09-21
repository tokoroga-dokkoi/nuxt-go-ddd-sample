<template>
  <div class="mt-20 text-center signup-form">
    <h2 class="text-5xl font-bold">ユーザ登録</h2>
    <p class="font-bold text-black mt-2">
      サブテキストサブテキストサブテキストサブテキスト
    </p>
    <sign-up-form
      class="mt-4"
      :email.sync="email"
      @signup-email="signUpWithEmail"
      @signup-google="signUpWithGoogle"
      @signup-facebook="signUpWithFacebook"
    />
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
  data() {
    return {
      email: '',
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
    async signUpWithEmail() {
      try {
        const email: string = 'sample_user_2@example.com'
        const password: string = 'passWord1234!'
        // test user created
        const res = await this.$auth.createUserWithEmailAndPassword(
          email,
          password
        )

        const idToken = await res.user.getIdToken()
        console.log(idToken)
        // backendにリクエスト送信
        this.$axios.post('/v1/auth/signup', {
          email: this.email,
          idToken,
        })

      } catch (e) {
        console.error(e)
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
