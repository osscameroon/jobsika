<template>
  <div class="flex justify-center items-center flex-col lg:flex-row">
    <div class="mb-3 w-full sm:w-1/2 lg:w-32 ml-0 lg:ml-6">
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
        <option
          selected
          style="font-family: 'Inter', sans-serif"
          class="text-black text-sm font-medium"
        >
          Job Title
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          class="text-black text-sm font-medium"
          value="1"
        >
          Kiroo
        </option>
      </select>
    </div>
    <div class="mb-3 w-full sm:w-1/2 lg:w-32 ml-0 lg:ml-6">
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
          Company
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="link in companies"
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
import axios from 'axios'
import Vue from 'vue'
import { BASE_URL } from '../constants/api'
export default Vue.extend({
  name: 'SelectComponent',
  props: {},
  data() {
    return {
      companies: [],
      content: [
        {
          name: 'Job Title',
        },

        {
          name: 'Company',
        },
      ],
    }
  },
  mounted() {
    axios.get(BASE_URL + '/companies').then((response) => {
      this.companies = response.data
      console.log('Companies:::::', this.companies[0].name)
    })
  },
})
</script>
