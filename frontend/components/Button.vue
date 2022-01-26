<template>
  <div
    :style="myStyle"
    class="
      site__btn
      cursor-pointer
      py-2
      text-white
      rounded-sm
      font-bold
      flex
      items-center
      justify-center
      text-xs
      md:text-sm
    "
  >
    <div v-if="showPicture === 'nothing'" />
    <img v-else :src="plusIcon" class="w-5 h-5 mr-1" />
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
    showPicture: String,
  },
  data() {
    return {
      plusIcon: require('../assets/plus.png'),
    }
  },
})
</script>
