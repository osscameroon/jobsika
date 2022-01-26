<template>
  <div class="site__title" v-if="onlyTitle === 'onlyTitle'">
    <div :style="myStyle" class="pt-2 md:pt-8 lg:pt-14">
      <div class="site__title-text text-center">
        <h2
          v-if="content === 'subtitle'"
          style="color: #000000; font-family: 'Inter', sans-serif"
          class="
            font-bold
            text-xl
            md:text-2xl
            flex
            justify-center
            md:justify-start
          "
        >
          {{ title }}
        </h2>
        <h2
          v-else
          style="color: #000000; font-family: 'Inter', sans-serif"
          class="font-bold text-3xl md:text-4xl flex justify-center"
        >
          {{ title }}
        </h2>
      </div>
    </div>
  </div>
  <div class="site__title" v-else>
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
        <h2
          style="color: #000000; font-family: 'Inter', sans-serif"
          class="font-bold text-3xl md:text-4xl"
        >
          {{ title }}
        </h2>
        <h4
          v-if="fontSize === 'header'"
          style="color: #235365; font-family: 'Inter', sans-serif"
          class="pt-3 font-bold text-2xl md:text-3xl"
        >
          {{ subTitle }}
        </h4>
        <h4
          v-else
          style="color: #235365; font-family: 'Inter', sans-serif"
          class="pt-3 font-bold text-base"
        >
          {{ subTitle }}
        </h4>
      </div>
      <div class="site__title-btn pt-10 lg:pt-0 w-full sm:w-1/2 lg:w-36">
        <a :href="myLink">
          <Button
            myStyle="background: #235365; font-family: 'Inter', sans-serif"
            name="Contribute"
          />
        </a>
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
    onlyTitle: String,
    myLink: String,
    content: String,
    myStyle: String,
  },
  components: {
    Button,
  },
})
</script>
