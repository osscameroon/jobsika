<template>
  <div
    style="background: #235365"
    class="
      site__btn
      cursor-pointer
      px-
      py-2
      text-white text-sm
      rounded-sm
      font-bold
      flex
      items-center
      justify-center
    "
  >
    <img :src="plusIcon" class="w-5 h-5 mr-1" />
    {{ name }}
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'ButtonComponent',
  props: {
    name: String,
    myStyle: String,
  },
  data() {
    return {
      plusIcon: require('../assets/plus.png'),
    }
  },
})
</script>
