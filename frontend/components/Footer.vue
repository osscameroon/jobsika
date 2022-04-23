<template>
  <footer
    class="site__footer mt-10 md:mt-28 flex flex-col md:flex-row items-center justify-center"
  >
    <a href="https://github.com/osscameroon/camerdevs">
      <img :src="github" alt="logo" class="site__footer-img w-8 h-8" />
    </a>
    <p class="site__footer-copyrigth text-xs md:text-sm md:ml-7 mt-3 md:mt-0">
      © 2022 OSS Cameroon, All rights reserved.
    </p>
  </footer>
</template>
<script>
export default {
  name: 'FooterComponent',
  data() {
    return {
      github: require('../assets/github.png'),
    }
  },
}
</script>
