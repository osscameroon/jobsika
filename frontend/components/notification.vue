<template>
  <div
    class="notification is-danger mb-2 font-medium text-xs md:text-sm"
    style="color: #dc3545; font-family: 'Inter', sans-serif"
  >
    {{ message }}
  </div>
</template>

<script>
export default {
  name: 'NotificationComponent',
  props: ['message'],
}
</script>
