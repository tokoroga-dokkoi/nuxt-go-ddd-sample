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
      <a-labeled-item label="メールアドレス" required>
        <a-text-field id="mail" v-model="email" placeholder="メールアドレス" />
      </a-labeled-item>
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
import ALabeledItem from '@/components/atoms/ALabeledItem.vue'
export default Vue.extend({
  components: {
    ATextField,
    ALabeledItem,
  },
  data() {
    return {
      email: '',
    }
  },
})
</script>
