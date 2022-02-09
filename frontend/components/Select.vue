<template>
  <div class="flex justify-center items-center flex-col sm:flex-row">
    <div class="mb-3 w-full sm:w-1/3 lg:w-32 ml-0 lg:ml-6">
      <select
        class="
          form-select
          appearance-none
          block
          w-full
          px-3
          py-1.5
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
          text-xs
          md:text-sm
          font-bold
        "
        aria-label="Default select example"
      >
        <option
          class="text-xs md:text-sm"
          selected
          style="font-family: 'Inter', sans-serif"
        >
          Job title
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="link in jobtitles"
          :key="link"
          value="1"
          class="text-xs md:text-sm"
        >
          {{ link }}
        </option>
      </select>
    </div>
    <div class="mb-3 w-full sm:w-1/3 lg:w-32 ml-0 sm:ml-6">
      <select
        class="
          form-select
          appearance-none
          block
          w-full
          px-3
          py-1.5
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
          text-xs
          md:text-sm
          font-bold
        "
        aria-label="Default select example"
      >
        <option
          class="text-xs md:text-sm"
          selected
          style="font-family: 'Inter', sans-serif"
        >
          Company
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="link in companies"
          :key="link.salary_id"
          value="1"
          class="text-xs md:text-sm"
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
      jobtitles: [],
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
  async created() {
    try {
      this.companies = (await axios.get(BASE_URL + '/companies')).data
      this.jobtitles = (await axios.get(BASE_URL + '/jobtitles')).data
    } catch (e) {
      console.log(e)
    }
  },
})
</script>
