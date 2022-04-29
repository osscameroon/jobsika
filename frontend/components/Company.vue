<template>
  <div class="flex flex-col">
    <div v-if="errored" class="flex flex-col items-center justify-center">
      <Error
        title="  Oops! Something went wrong"
        content="Unable to connect to the internet"
      />
    </div>
    <div v-else class="-my-2">
      <div class="py-2 align-middle inline-block min-w-full w-16">
        <div
          class="overflow-x-auto rounded-xl border-gray-200 sm:rounded-lg w-full"
        >
          <table class="min-w-full divide-y-8 divide-primary">
            <thead style="backgound: #e5e5e5">
              <tr>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Company
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Jobs title
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Seniority
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  City
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Salary
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
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
              v-for="(company, myIndex) in ratings"
              :key="myIndex + 1"
              class="divide-y divide-primary"
            >
              <tr>
                <td
                  class="bg-white rounded-tl-xl rounded-bl-md px-4 py-3 whitespace-nowrap"
                >
                  <div
                    class="text-xs text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{
                      company.company_name !== ''
                        ? company.company_name
                        : 'A local company'
                    }}
                  </div>
                </td>
                <td class="px-4 py-3 whitespace-nowrap bg-white">
                  <div
                    class="text-xs text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{
                      company.job_title !== ''
                        ? company.job_title
                        : 'A job title'
                    }}
                  </div>
                </td>
                <td class="px-4 py-3 whitespace-nowrap bg-white">
                  <div
                    class="text-xs text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.seniority }}
                  </div>
                </td>
                <td class="px-4 py-3 whitespace-nowrap bg-white">
                  <div
                    class="text-xs text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.city }}
                  </div>
                </td>
                <td
                  class="px-4 py-3 whitespace-nowrap text-sm text-gray-500 bg-white"
                >
                  <div
                    class="text-xs text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.salary | price }}
                  </div>
                </td>
                <td
                  class="px-4 py-3 whitespace-nowrap text-sm font-medium bg-white"
                >
                  <div class="flex">
                    <div class="flex">
                      <div v-if="company.rating === 0" class="flex">
                        <div
                          v-for="item in 5"
                          :key="item"
                          class="flex flex-row"
                        >
                          <img class="w-4 h-4 mr-1" :src="starsPictureGray" />
                        </div>
                      </div>
                      <img v-else class="w-4 h-4 mr-1" :src="starsPicture" />
                    </div>
                  </div>
                </td>
                <td
                  class="px-4 py-3 whitespace-nowrap text-right text-sm font-medium bg-white rounded-tr-xl rounded-br-md"
                  :class="{ opened: opened.includes(company.salary_id) }"
                  @click="toggle(company.salary_id)"
                >
                  <button
                    v-if="String(company.comment).length === 0"
                    disable
                    class="text-grayC text-xs flex items-center space-x-3 cursor-text"
                    style="font-family: 'Inter', sans-serif"
                    type="button"
                  >
                    Details
                  </button>
                  <button
                    v-else
                    class="text-xs flex items-center space-x-3 cursor-pointer"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                    type="button"
                  >
                    Details
                  </button>
                </td>
              </tr>
              <td
                v-if="
                  opened.includes(company.salary_id) &&
                  String(company.comment).length !== 0
                "
                colspan="10"
                class="w-10/12 bg-primary"
              >
                <div class="bg-white w-full p-4 my-3 shadow-sm rounded-sm">
                  <p
                    class="py-2 text-xs md:text-sm"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    {{ company.createdat }}
                  </p>
                  <p
                    class="text-xs md:text-sm"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.comment }}
                  </p>
                </div>
              </td>
              <div v-else />
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'
import { mapMutations } from 'vuex'
export default Vue.extend({
  name: 'CompanyComponent',
  data() {
    return {
      isOpen: false,
      starsPicture: require('../assets/star.png'),
      starsPictureGray: require('../assets/stargray.png'),
      opened: [],
      errored: false,
    }
  },
  computed: {
    page() {
      return this.$store.state.ratings.page
    },
    limit() {
      return this.$store.state.ratings.limit
    },
    ratings() {
      return this.$store.state.ratings.ratings
    },
    filterjob() {
      return this.$store.state.ratings.filterjob
    },
    filtercompany() {
      return this.$store.state.ratings.filtercompany
    },
  },

  async created() {},
  methods: {
    ...mapMutations({
      setpage: 'ratings/SETPAGE',
      setlimit: 'ratings/SETLIMIT',
    }),
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
