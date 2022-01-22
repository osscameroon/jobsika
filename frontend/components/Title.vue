<template>
  <div class="site__title">
    <div
      class="
        flex
        justify-between
        lg:items-end lg:flex-row
        flex-col
        pt-2
        md:pt-12
        lg:pt-24
        items-center
      "
    >
      <div class="site__title-text text-center lg:text-left">
        <h2 class="font-bold text-3xl md:text-4xl">
          {{ title }}
        </h2>
        <h4
          v-if="fontSize === 'header'"
          style="color: #235365; font-family: 'Inter', sans-serif"
          class="pt-3 font-bold text-2xl md:text-3xl"
        >
          {{ subTitle }}
        </h4>
        <h4 v-else style="color: #235365" class="pt-3 font-bold text-base">
          {{ subTitle }}
        </h4>
      </div>
      <div class="site__title-btn pt-10 lg:pt-0 w-full sm:w-1/2 lg:w-36">
        <Button name="Contribute" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Button from '../components/Button.vue'
export default Vue.extend({
  name: 'TitleComponent',
  props: {
    title: String,
    subTitle: String,
    fontSize: String,
  },
  components: {
    Button,
  },
})
</script>
