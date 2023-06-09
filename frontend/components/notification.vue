<template>
  <div
    class="notification is-danger mb-3 md:mb-6 font-light text-xs"
    style="color: #dc3545; font-family: 'Inter', sans-serif"
  >
    {{ message }}
  </div>
</template>

<script>
export default {
  name: 'NotificationComponent',
  props: {
    message: {
      type: String,
      default: '',
    },
  },
}
</script>
