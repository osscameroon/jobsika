<template>
  <div class="flex flex-col">
    <div class="-my-2">
      <div class="py-2 align-middle inline-block min-w-full w-16">
        <div
          class="
            overflow-x-auto
            rounded-xl
            border-gray-200
            sm:rounded-lg
            w-full
          "
        >
          <table class="min-w-full divide-y-8 divide-primary">
            <thead style="backgound: #e5e5e5">
              <tr>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    md:text-sm
                    font-bold
                    tracking-wider
                  "
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Position
                </th>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    md:text-sm
                    font-bold
                    tracking-wider
                  "
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Company
                </th>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    md:text-sm
                    font-bold
                    tracking-wider
                  "
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Jobs title
                </th>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    md:text-sm
                    font-bold
                    tracking-wider
                  "
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Salary
                </th>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    md:text-sm
                    font-bold
                    tracking-wider
                  "
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Rating
                </th>
                <th scope="col" class="relative px-6 py-3 flex">
                  <span class="sr-only">Edit</span>
                </th>
              </tr>
            </thead>
            <tbody
              v-for="company in companies.slice(0, 4)"
              :key="company.id"
              class="bg-white shadow-md divide-y divide-white"
            >
              <tr>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="ml-4">
                      <div
                        class="text-xs md:text-sm font-medium text-gray-900"
                        style="color: #000000; font-family: 'Inter', sans-serif"
                      >
                        {{ company.salary_id }}
                      </div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div
                    class="text-xs md:text-sm text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.company_name }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div
                    class="text-xs md:text-sm text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.job_title }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <div
                    class="text-xs md:text-sm text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.salary }} FCFA
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="flex">
                    <div
                      v-for="item in company.rating"
                      :key="item"
                      class="flex"
                    >
                      <img class="w-4 h-4 mr-1" :src="startPicture" />
                    </div>
                  </div>
                </td>
                <td
                  class="
                    px-6
                    py-4
                    whitespace-nowrap
                    text-right text-sm
                    font-medium
                  "
                >
                  <a
                    class="
                      text-xs
                      md:text-sm
                      flex
                      items-center
                      space-x-3
                      cursor-pointer
                    "
                    style="color: #000000; font-family: 'Inter', sans-serif"
                    :aria-expanded="isOpen"
                    :aria-controls="`collapse${_uid}`"
                    @click="toggleAccordion()"
                    >Details</a
                  >
                </td>
              </tr>
              <div
                v-show="isOpen"
                :id="`collapse${_uid}`"
                class="w-10/12 bg-red-500"
              >
                <p class="w-full bg-yellow-600">{{ company.comment }}</p>
              </div>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import axios from 'axios'
import Vue from 'vue'
import { BASE_URL } from '../constants/api'
export default Vue.extend({
  name: 'CompanyComponent',
  data() {
    return {
      companies: [],
      isOpen: false,
      startPicture: require('../assets/star.png'),
    }
  },
  async created() {
    try {
      this.companies = (await axios.get(BASE_URL + '/ratings')).data
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    toggleAccordion() {
      this.isOpen = !this.isOpen
    },
  },
})
</script>
