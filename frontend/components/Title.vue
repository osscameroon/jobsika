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
        <h2 style="font-size: 32px" class="font-bold leading-10">
          {{ title }}
        </h2>
        <h4
          v-if="fontSize === 'header'"
          style="
            color: #235365;
            font-size: 32px;
            font-family: 'Inter', sans-serif;
          "
          class="pt-3 font-semibold"
        >
          {{ subTitle }}
        </h4>
        <h4
          v-else
          style="color: #235365; font-size: 16px"
          class="pt-3 font-semibold"
        >
          {{ subTitle }}
        </h4>
      </div>
      <div class="site__title-btn pt-4 lg:pt-0">
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
