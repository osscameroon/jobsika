<template>
  <div class="flex flex-col">
    <div v-if="errored" class="flex flex-col items-center justify-center">
      <Error
        title="  Oops! Something went wrong"
        content="Unable to connect to the internet"
      />
    </div>
    <div v-else class="-my-2">
      <div class="inline-block w-16 min-w-full py-2 align-middle">
        <div
          class="w-full overflow-x-auto border-gray-200 rounded-xl sm:rounded-lg"
        >
          <table class="min-w-full divide-y-8 divide-primary">
            <thead style="backgound: #e5e5e5">
              <tr>
                <th
                  scope="col"
                  class="px-4 py-3 text-xs font-bold tracking-wider text-left md:text-sm"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Company
                </th>
                <th
                  scope="col"
                  class="px-4 py-3 text-xs font-bold tracking-wider text-left md:text-sm"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Jobs title
                </th>
                <th
                  scope="col"
                  class="px-4 py-3 text-xs font-bold tracking-wider text-left md:text-sm"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Seniority
                </th>
                <th
                  scope="col"
                  class="px-4 py-3 text-xs font-bold tracking-wider text-left md:text-sm"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  City
                </th>
                <th
                  scope="col"
                  class="px-4 py-3 text-xs font-bold tracking-wider text-left md:text-sm"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Salary
                </th>
                <th
                  scope="col"
                  class="px-4 py-3 text-xs font-bold tracking-wider text-left md:text-sm"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Rating
                </th>
                <th scope="col" class="relative flex px-6 py-3">
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
                  class="px-4 py-3 bg-white rounded-tl-xl rounded-bl-md whitespace-nowrap"
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
                <td class="px-4 py-3 bg-white whitespace-nowrap">
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
                <td class="px-4 py-3 bg-white whitespace-nowrap">
                  <div
                    class="text-xs text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.seniority }}
                  </div>
                </td>
                <td class="px-4 py-3 bg-white whitespace-nowrap">
                  <div
                    class="text-xs text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.city }}
                  </div>
                </td>
                <td
                  class="px-4 py-3 text-sm text-gray-500 bg-white whitespace-nowrap"
                >
                  <div
                    class="text-xs text-gray-900"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{
                      parseInt(company.salary.toString()) < 1000
                        ? parseInt(company.salary.toString()) * 1000
                        : company.salary | price
                    }}
                  </div>
                </td>
                <td
                  class="px-4 py-3 text-sm font-medium bg-white whitespace-nowrap"
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

                      <div v-else class="flex">
                        <div
                          v-for="item in company.rating"
                          :key="item"
                          class="flex"
                        >
                          <img class="w-4 h-4 mr-1" :src="starsPicture" />
                        </div>
                      </div>
                    </div>
                  </div>
                </td>
                <td
                  class="px-4 py-3 text-sm font-medium text-right bg-white whitespace-nowrap rounded-tr-xl rounded-br-md"
                  :class="{ opened: opened.includes(company.salary_id) }"
                  @click="toggle(company.salary_id)"
                >
                  <button
                    v-if="String(company.comment).length === 0"
                    disable
                    class="flex items-center text-xs text-grayC space-x-3 cursor-text"
                    style="font-family: 'Inter', sans-serif"
                    type="button"
                  >
                    Comments
                    <span class="flex items-center">
                      <span
                        class="flex items-center justify-center w-3 h-3 ml-1 text-center border rounded-full cursor-pointer text-grayC border-grayC"
                        style="font-size: 8px"
                      >
                        !
                      </span>
                    </span>
                  </button>
                  <button
                    v-else
                    class="flex items-center text-xs cursor-pointer space-x-3"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                    type="button"
                  >
                    Comments
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
                <div class="w-full p-4 my-3 bg-white rounded-sm shadow-sm">
                  <p
                    class="py-2 text-xs md:text-sm"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    {{
                      new Date(
                        Date.parse(
                          company.createdat !== undefined
                            ? company.createdat
                            : Date.now()
                        )
                      ).toDateString()
                    }}
                  </p>
                  <p
                    class="text-xs md:text-sm"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    {{ company.comment }}
                  </p>
                </div>
              </td>
              <td
                v-if="
                  opened.includes(company.salary_id) &&
                  String(company.comment).length === 0
                "
                colspan="10"
                class="w-10/12 bg-primary"
              >
                <div class="w-full p-4 my-3 bg-white rounded-sm shadow-sm">
                  <p
                    class="py-2 text-xs md:text-sm"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    Details will show up if there is more than three entries
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
