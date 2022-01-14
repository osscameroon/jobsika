<template>
  <div class="flex justify-center items-center flex-col lg:flex-row">
    <div
      v-for="item in content"
      :key="item"
      class="mb-3 w-full sm:w-1/2 lg:w-32 ml-0 lg:ml-6"
    >
      <select
        class="
          form-select
          appearance-none
          block
          w-full
          px-3
          py-1.5
          text-base
          font-normal
          text-gray-700
          bg-white bg-clip-padding bg-no-repeat
          border border-solid border-gray-300
          rounded
          transition
          ease-in-out
          m-0
          focus:text-gray-700
          focus:bg-white
          focus:border-blue-600
          focus:outline-none
        "
        aria-label="Default select example"
      >
        <option selected style="font-family: 'Inter', sans-serif">
          {{ item.name }}
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="link in item.links"
          :key="link"
          value="1"
        >
          {{ link.name }}
        </option>
      </select>
    </div>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'SelectComponent',
  props: {},
  data() {
    return {
      content: [
        {
          name: 'Job Title',
          links: [
            {
              name: 'OSS Cameroon',
              link: '#',
            },
            {
              name: 'Kiroo Games',
              link: '#',
            },
          ],
        },
        {
          name: 'Company',
          links: [
            {
              name: 'OSS Cameroon Company',
              link: '#',
            },
            {
              name: 'Kiroo Games Company',
              link: '#',
            },
          ],
        },
      ],
    }
  },
})
</script>
