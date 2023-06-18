<template lang="html">
  <div v-show="isActive">
    <slot></slot>
  </div>
</template>
<script>
export default {
  name: 'TabComponent',
  props: {
    title: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      isActive: true,
    }
  },
}
</script>
