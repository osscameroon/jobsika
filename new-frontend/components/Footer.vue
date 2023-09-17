<script setup lang="ts">
  const github: Ref<string> = ref('/github.png');
</script>

<template>
  <footer
    class="site__footer mt-10 md:mt-28 flex flex-col md:flex-row items-center justify-center"
  >
    <a href="https://github.com/osscameroon/jobsika">
      <img :src="github" alt="logo" class="site__footer-img w-8 h-8" />
    </a>
    <p class="site__footer-copyrigth text-xs md:text-sm md:ml-7 mt-3 md:mt-0">
      © {{ new Date().getFullYear() }} OSS Cameroon, All rights reserved.
    </p>
  </footer>
</template>