<template>
  <div class="mt-20 text-center signup-form">
    <h2 class="text-5xl font-bold">ユーザ登録</h2>
    <p class="font-bold text-black mt-2">
      サブテキストサブテキストサブテキストサブテキスト
    </p>
    <sign-up-form class="mt-4" />
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
import SignUpForm from '@/components/organisms/signUpForm.vue'
export default Vue.extend({
  components: {
    SignUpForm,
  },
})
</script>
<style lang="scss" scoped>
.signup-form {
  max-width: 560px;
  margin: 60px auto;
}
</style>
