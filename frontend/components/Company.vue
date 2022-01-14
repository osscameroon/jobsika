<template>
  <div class="flex flex-col w-full">
    <div class="-my-2 overflow-x-auto">
      <div class="py-2 align-middle inline-block min-w-full">
        <div
          class="
            overflow-hidden
            border-b
            rounded-xl
            border-gray-200
            sm:rounded-lg
            w-full
          "
        >
          <table class="min-w-full divide-y divide-gray-200">
            <thead style="backgound: #e5e5e5">
              <tr>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    font-semibold
                    text-gray-500
                    uppercase
                    tracking-wider
                  "
                  style="color: #b1b1b1"
                >
                  Position
                </th>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    font-medium
                    text-gray-500
                    uppercase
                    tracking-wider
                  "
                  style="color: #b1b1b1"
                >
                  Company Name
                </th>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    font-medium
                    text-gray-500
                    uppercase
                    tracking-wider
                  "
                  style="color: #b1b1b1"
                >
                  Jobs title
                </th>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    font-medium
                    text-gray-500
                    uppercase
                    tracking-wider
                  "
                  style="color: #b1b1b1"
                >
                  Salary
                </th>
                <th
                  scope="col"
                  class="
                    px-6
                    py-3
                    text-left text-xs
                    font-medium
                    text-gray-500
                    uppercase
                    tracking-wider
                  "
                  style="color: #b1b1b1"
                >
                  Rating
                </th>
                <th scope="col" class="relative px-6 py-3">
                  <span class="sr-only">Edit</span>
                </th>
              </tr>
            </thead>
            <tbody
              class="bg-white divide-y divide-white"
              v-for="company in companies.slice(0, 4)"
              :key="company"
            >
              <tr>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900">
                        {{ company.salary_id }}
                      </div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">
                    {{ company.company_name }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">
                    {{ company.job_title }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <div class="text-sm text-gray-900">
                    {{ company.salary }} FCFA
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <div class="text-sm text-gray-900">{{ company.rating }}</div>
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
                    @click="toggleAccordion()"
                    class="flex items-center space-x-3 cursor-pointer"
                    :aria-expanded="isOpen"
                    :aria-controls="`collapse${_uid}`"
                    >Details</a
                  >
                </td>
              </tr>
              <div
                v-show="isOpen"
                class="w-10/12 bg-red-500"
                :id="`collapse${_uid}`"
              >
                <!-- <slot name="{{company.comment}}" /> -->
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
export default Vue.extend({
  name: 'CompanyComponent',
  data() {
    return {
      companies: [],
      isOpen: false,
    }
  },
  mounted() {
    axios.get('http://localhost:7000/ratings').then((response) => {
      this.companies = response.data
    })
  },
  methods: {
    toggleAccordion() {
      this.isOpen = !this.isOpen
    },
  },
})
</script>
