<template>
  <div class="block">
    <div class="inline text-center">
      <button
        class="bg-facebookBlue text-white font-bold py-4 px-10 rounded-full"
        @click="$emit('signup-facebook')"
      >
        Facebookでログイン
      </button>
      <button
        class="bg-white text-black font-bold py-4 px-10 ml-10 rounded-full border-2"
        @click="$emit('signup-google')"
      >
        Googleでログイン
      </button>
      <p class="text-sm mt-5">SNSに許可なく投稿されることはありません。</p>
    </div>

    <p class="my-10">または</p>

    <div class="mt-6 text-left">
      <a-labeled-item label="メールアドレス" required>
        <a-text-field
          :value="email"
          placeholder="メールアドレス"
          class="mt-2"
          required
          @input="val => $emit('update:email', val.target.value)"
        />
        <template v-if="isError" v-slot:errors>
          <p class="text-sm text-alert">{{ errorMessage }}</p>
        </template>

        <p class="text-sm mt-5">
          登録したメールアドレスに本登録用の案内メールが届きます。
        </p>
      </a-labeled-item>
    </div>

    <div class="mt-4 text-center">
      <button
        class="bg-black text-white font-bold py-4 px-10 rounded-full"
        @click="checkEmail"
      >
        メールアドレスで登録
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue'
import ATextField from '@/components/atoms/ATextField.vue'
import ALabeledItem from '@/components/atoms/ALabeledItem.vue'
export default Vue.extend({
  components: {
    ATextField,
    ALabeledItem,
  },
  props: {
    email: {
      type: String as PropType<string>,
      required: true,
    },
  },
  data() {
    return {
      isError: false,
      errorMessage: '',
    }
  },
  methods: {
    checkEmail() {
      if (this.email === '' || this.email === undefined) {
        this.isError = true
        this.errorMessage = 'メールアドレスの入力は必須です'
        return
      }

      const regexpEmail: RegExp = /^[A-Za-z0-9]{1}[A-Za-z0-9_.-]*@{1}[A-Za-z0-9_.-]{1,}\.[A-Za-z0-9]{1,}$/
      if (!regexpEmail.test(this.email)) {
        this.isError = true
        this.errorMessage = 'メールアドレスの形式が不正です'
        return
      }

      this.$emit('signup-email')
    },
  },
})
</script>
