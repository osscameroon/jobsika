<template>
  <div class="site__input w-full">
    <p
      class="text-xs md:text-sm font-bold"
      style="color: #b1b1b1; font-family: 'Inter', sans-serif"
    >
      {{ title }}
    </p>
    <input
      v-if="myInput === 'input'"
      :v-model="myModelInput"
      type="text"
      class="
        site__input-field
        border-none
        mt-2
        md:mt-3
        w-full
        rounded-md
        mb-4
        md:mb-16
      "
      :style="myStyle"
    />

    <textarea
      v-else
      :v-model="myModelTextArea"
      class="resize rounded-md w-full h-full mt-2 md:mt-3 border-none"
      name="story"
      rows="5"
    />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'MyInputComponent',
  props: ['title', 'myStyle', 'myInput', 'myModelInput', 'myModelTextArea'],
})
</script>e
