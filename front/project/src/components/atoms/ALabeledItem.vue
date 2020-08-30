<template>
  <div :class="{ 'mb-5': !dense }" class="b-labeled-item">
    <component
      :is="tag"
      :class="{ small: small }"
      class="flex-row flex-col mb-1 text-left"
    >
      <span class="d-flex">
        <template v-if="required">
          <span
            class="b-labeled-item-required border-solid border-alert border-2 text-alert px-2"
            >必須</span
          >
        </template>
        <template v-if="label">
          <span class="b-labeled-item-label">{{ label }}</span>
        </template>
        <slot v-else name="label" />
      </span>
      <span v-if="hint" class="caption">{{ hint }}</span>
      <div class="caption">
        <slot name="hint" />
      </div>
      <span v-if="!$props.for" class="mt-2">
        <slot />
      </span>
    </component>
    <div v-if="$props.for" class="mt-2">
      <slot :id="$props.for" />
    </div>
    <div class="errors">
      <slot name="errors" />
    </div>
  </div>
</template>
<script lang="ts">
import Vue, { PropType } from 'vue'
export default Vue.extend({
  props: {
    label: {
      type: String as PropType<string | undefined>,
      default: undefined,
    },
    required: {
      type: Boolean as PropType<boolean>,
      default: false,
    },
    optional: {
      type: Boolean as PropType<boolean>,
      default: false,
    },
    secret: {
      type: Boolean as PropType<boolean>,
      default: false,
    },
    small: {
      type: Boolean as PropType<boolean>,
      default: false,
    },
    for: {
      type: String as PropType<string | undefined>,
      default: undefined,
    },
    id: {
      type: String as PropType<string | undefined>,
      default: undefined,
    },
    tag: {
      type: String as PropType<string>,
      default: 'label',
    },
    hint: {
      type: String as PropType<string | undefined>,
      default: undefined,
    },
    dense: {
      type: Boolean as PropType<boolean>,
      default: false,
    },
  },
})
</script>
<style lang="scss" scoped>
.small {
  font-size: 15px;
}
</style>
