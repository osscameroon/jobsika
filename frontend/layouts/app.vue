<template>
  <div class="site bg-primary">
    <header
      class="site__logo bg-blueDark flex justify-center items-center md:justify-start py-3 md:py-2 z-50 sticky top-0 left-0 shadow-md"
    >
      <div class="md:container md:mx-auto md:w-10/12">
        <nuxt-link to="/" class="cursor-pointer">
          <img
            :src="logo"
            alt="logo"
            class="site__logo-pic w-32 h-12 md:w-36 md:h-14"
          />
        </nuxt-link>
      </div>
    </header>
    <Nuxt />
  </div>
</template>

<script>
export default {
  data() {
    return {
      logo: require('../assets/logo_trans.svg'),
    }
  },
}
</script>
