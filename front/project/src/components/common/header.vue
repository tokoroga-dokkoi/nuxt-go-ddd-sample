<template>
  <header class="bg-gray-100">
    <div
      class="header-container flex items-center justify-between flex-wrap p-6 max-w-screen-lg mx-auto"
    >
      <div class="flex items-center text-brack mr-6">
        <span class="font-semibold text-xl a-top-link">テキスト</span>
        <a-text-field v-model="searchWord" placeholder="検索する" />
      </div>

      <div class="w-full block lg:flex lg:items-left lg:w-auto">
        <button
          class="bg-gray1-500 hover:bg-gray-100 text-black font-bold py-2 px-4 rounted border border-black mr-4"
        >
          ログイン
        </button>
        <button
          class="bg-black hover:bg-black text-white font-bold py-2 px-4 rounted"
        >
          会員登録
        </button>
      </div>
    </div>
  </header>
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
      searchWord: '',
    }
  },
})
</script>
<style lang="scss" scoped>
.a-top-link {
  width: 200px;
}
</style>
