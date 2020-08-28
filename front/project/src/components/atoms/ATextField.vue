<template>
  <input
    class="shadow appearance-none border rounded w-full py-2 px-3 text-black leading-tight focus:outline-none focus:shadow-outline"
    type="text"
    :value="value"
    :placeholder="placeholder"
    :class="{ rounded: rounded, border: border }"
    v-bind="$attrs"
    v-on="$listeners"
  />
</template>
<script lang="ts">
import Vue, { PropType } from 'vue'
export default Vue.extend({
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
  },
})
</script>
