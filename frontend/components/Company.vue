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
              v-for="(company, myIndex) in companies.slice(0, 4)"
              :key="myIndex + 1"
              class="bg-white shadow-md divide-y divide-white"
            >
              <tr>
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
                      <img class="w-4 h-4 mr-1" :src="starsPicture" />
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
                  @click="toggle(company.salary_id)"
                  :class="{ opened: opened.includes(company.salary_id) }"
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
                    >Details</a
                  >
                </td>
              </tr>
              <td
                v-if="opened.includes(company.salary_id)"
                colspan="6"
                class="w-10/12 bg-primary"
              >
                <div class="bg-white w-full p-4 my-3 shadow-sm rounded-sm">
                  <p
                    class="py-2 text-xs md:text-sm"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    22/03/2022 12:34
                  </p>
                  <p
                    class="text-xs md:text-sm"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.comment }}
                  </p>
                </div>
              </td>
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
      starsPicture: require('../assets/star.png'),
      opened: [],
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
    toggle(id) {
      const index = this.opened.indexOf(id)
      if (index > -1) {
        this.opened.splice(index, 1)
      } else {
        this.opened.push(id)
      }
    },
  },
})
</script>
