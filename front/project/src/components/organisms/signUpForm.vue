<template>
  <div class="block">
    <div class="inline text-center">
      <button
        class="bg-twitterBlue text-white font-bold py-4 px-10 rounded-full"
      >
        Twitterでログイン
      </button>
      <button
        class="bg-facebookBlue text-white font-bold py-4 px-10 ml-10 rounded-full"
      >
        FaceBookでログイン
      </button>
      <p class="text-sm mt-5">SNSに許可なく投稿されることはありません。</p>
    </div>

    <p class="my-10">または</p>

    <div class="mt-6 text-left">
      <label class="block text-gray-700 text-sm font-bold" for="mail"
        >メールアドレス</label
      >
      <a-text-field id="mail" v-model="email" placeholder="メールアドレス" />
    </div>

    <div class="mt-4 text-center">
      <button class="bg-black text-white font-bold py-4 px-10 rounded-full">
        メールアドレスで登録
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import ATextField from '@/components/atoms/ATextField.vue'
export default Vue.extend({
  components: {
    ATextField,
  },
  data() {
    return {
      email: '',
    }
  },
})
</script>
