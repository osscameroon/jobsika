<template>
  <div class="site__error">
    <header>
      <h1
        style="color: #bb321f; font-family: 'Inter', sans-serif"
        class="py-3 text-lg font-bold"
      >
        {{ title }}
      </h1>
    </header>
    <main>
      <p
        style="color: #000; font-family: 'Inter', sans-serif"
        class="text-sm text-center leading-6 mx-auto"
      >
        {{ content }}
      </p>
    </main>
  </div>
</template>

<script>
export default {
  name: 'ErrorComponent',
  props: {
    title: {
      type: String,
      default: '',
    },
    content: {
      type: String,
      default: '',
    },
  },
}
</script>
