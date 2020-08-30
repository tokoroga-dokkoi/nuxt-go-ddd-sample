<template>
  <input
    class="shadow appearance-none border rounded w-full py-2 px-3 text-black leading-tight focus:outline-none focus:shadow-outline"
    type="text"
    :value="value"
    :placeholder="placeholder"
    :class="[{ rounded: rounded, border: border }, formBgColor]"
    v-bind="$attrs"
    v-on="$listeners"
  />
</template>
<script lang="ts">
import Vue, { PropType } from 'vue'
export default Vue.extend({
  inheritAttrs: false,
  props: {
    rounded: {
      type: Boolean as PropType<boolean>,
      default: true,
    },
    border: {
      type: Boolean as PropType<boolean>,
      default: true,
    },
    value: {
      type: [String, Number] as PropType<string | number | undefined>,
      required: true,
    },
    placeholder: {
      type: String as PropType<string | undefined>,
      default: undefined,
    },
    required: {
      type: Boolean as PropType<Boolean>,
      default: false,
    },
  },
  data() {
    return {
      isInvalidBg: 'bg-formOrange-100',
      isValidBg: 'bg-formGray-100',
    }
  },
  computed: {
    formBgColor(): string {
      if (this.required && this.isEmptyValue) {
        return this.isInvalidBg
      }
      return this.isValidBg
    },
    isEmptyValue(): boolean {
      return this.value === '' || this.value === undefined
    },
  },
})
</script>
