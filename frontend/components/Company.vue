<template>
  <div class="flex flex-col">
    <div v-if="errored" class="flex flex-col items-center justify-center">
      <Error
        :title=" $t('error_500_title')"
        :content=" $t('error_internet_connection')"
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
                class="px-4 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                {{ $t('salary_table_header_company') }}
              </th>
              <th
                scope="col"
                class="px-4 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                {{$t('salary_table_header_job_title')}}
              </th>
              <th
                scope="col"
                class="px-4 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                {{$t('salary_table_header_seniority')}}
              </th>
              <th
                scope="col"
                class="px-4 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                {{$t('salary_table_header_city')}}
              </th>
              <th
                scope="col"
                class="px-4 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                {{$t('salary_table_header_salary')}}
              </th>
              <th
                scope="col"
                class="px-4 py-3 text-left text-xs md:text-sm font-bold tracking-wider"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                {{$t('salary_table_header_rating')}}
              </th>
              <th scope="col" class="relative px-6 py-3 flex">
                <span class="sr-only">{{$t('global_label_edit')}}</span>
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
                      : $t('label_a_local_company')
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
                      :  $t('label_a_job_title')
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
                        <img class="w-4 h-4 mr-1" :src="starsPictureGray"/>
                      </div>
                    </div>

                    <div v-else class="flex">
                      <div
                        v-for="item in company.rating"
                        :key="item"
                        class="flex"
                      >
                        <img class="w-4 h-4 mr-1" :src="starsPicture"/>
                      </div>
                    </div>
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
                  {{$t('global_label_comments')}}
                  <span class="flex items-center">
                      <span
                        class="cursor-pointer h-3 text-center w-3 ml-1 text-grayC rounded-full border border-grayC flex items-center justify-center"
                        style="font-size: 8px"
                      >
                        !
                      </span>
                    </span>
                </button>
                <button
                  v-else
                  class="text-xs flex items-center space-x-3 cursor-pointer"
                  style="color: #000000; font-family: 'Inter', sans-serif"
                  type="button"
                >
                  {{$t('global_label_comments')}}
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
              <div class="bg-white w-full p-4 my-3 shadow-sm rounded-sm">
                <p
                  class="py-2 text-xs md:text-sm"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  {{$t('label_not_enough_data')}}
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
import {mapMutations} from 'vuex'

export default Vue.extend({
  name: 'CompanyComponent',
  props: {
    isHomePage: {
      type: Boolean,
      required: false,
      default: false
    }
  },
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
      if (this.isHomePage) {
        return this.$store.state.ratings.ratings.slice(0, 5);
      } else {
        return this.$store.state.ratings.ratings;
      }
    },
    filterjob() {
      return this.$store.state.ratings.filterjob
    },
    filtercompany() {
      return this.$store.state.ratings.filtercompany
    },
  },

  async created() {
  },
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
