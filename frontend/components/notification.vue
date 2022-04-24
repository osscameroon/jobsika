<template>
  <div
    class="notification is-danger mb-3 md:mb-6 font-light"
    style="color: #dc3545; font-family: 'Inter', sans-serif; font-size: 9px"
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
