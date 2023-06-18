<template>
  <div class="flex justify-center items-center flex-col sm:flex-row">
    <div class="mb-3 w-full sm:w-1/3 lg:w-40 ml-0 lg:ml-6">
      <h4
        style="color: #b1b1b1"
        class="pb-1 font-semibold text-xs md:text-base"
      >
        Seniority
      </h4>
      <select
        v-model="myfilterseniority"
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        @change="onChangeSeniority"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          {{ $t('global_label_all_seniorties') }}
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          {{ $t('global_label_any') }}
        </option>
        <option
          v-for="(seniority, index) in seniorities"
          v-show="seniority !== ''"
          :key="index"
          style="font-family: 'Inter', sans-serif"
          :value="seniority"
          class="text-xs md:text-sm"
        >
          {{ seniority }}
        </option>
      </select>
    </div>
    <div class="mb-3 w-full sm:w-1/3 lg:w-40 ml-0 lg:ml-6">
      <h4
        style="color: #b1b1b1"
        class="pb-1 font-semibold text-xs md:text-base"
      >
    {{ $t('salary_table_header_job_title')}}
      </h4>
      <select
        v-model="myfilterjob"
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        @change="onChangeJobTitle"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          {{ $t('global_label_all_job_title') }}
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          {{ $t('global_label_any') }}
        </option>
        <option
          v-for="(link, index) in jobtitles"
          v-show="link !== ''"
          :key="index"
          style="font-family: 'Inter', sans-serif"
          :value="link"
          class="text-xs md:text-sm"
        >
          {{ link }}
        </option>
      </select>
    </div>
    <div class="mb-3 w-full sm:w-1/3 lg:w-40 ml-0 sm:ml-6">
      <h4
        style="color: #b1b1b1"
        class="pb-1 font-semibold text-xs md:text-base"
      >
        {{ $t('global_label_companies')}}
      </h4>
      <select
        v-model="myfiltercompany"
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        @change="onChangeCompany"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          {{ $t('global_label_all_companies') }}
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          {{ $t('global_label_any') }}
        </option>
        <option
          v-for="link in companies"
          v-show="link.name !== ''"
          :key="link.id"
          style="font-family: 'Inter', sans-serif"
          :value="link.name"
          class="text-xs md:text-sm"
        >
          {{ link.name }}
        </option>
      </select>
    </div>
    <div class="mb-3 w-full sm:w-1/3 lg:w-40 ml-0 sm:ml-6">
      <h4
        style="color: #b1b1b1"
        class="pb-1 font-semibold text-xs md:text-base"
      >
        {{ $t('global_label_cities') }}
      </h4>
      <select
        v-model="myfiltercity"
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        @change="onChangeCity"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          {{ $t('global_label_all_cities') }}
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          {{ $t('global_label_any') }}
        </option>
        <option
          v-for="(city, index) in cities"
          v-show="city !== ''"
          :key="index"
          style="font-family: 'Inter', sans-serif"
          :value="city"
          class="text-xs md:text-sm"
        >
          {{ city }}
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
      myfilterjob: '',
      myfiltercompany: '',
      myfilterseniority: '',
      myfiltercity: '',
    }
  },
  computed: {
    companies() {
      return this.$store.state.companies.companies
    },
    cities() {
      return this.$store.state.cities.cities
    },
    jobtitles() {
      return this.$store.state.jobtitles.jobtitles
    },
    seniorities() {
      return this.$store.state.seniorities.seniorities
    },
    filterjob() {
      return this.$store.state.jobtitles.filterjob
    },
    filtercompany() {
      return this.$store.state.companies.filtercompany
    },
    filterseniority() {
      return this.$store.state.seniorities.filterseniority
    },
    filtercity() {
      return this.$store.state.cities.filtercity
    },
    page() {
      return this.$store.state.ratings.page
    },
    limit() {
      return this.$store.state.ratings.limit
    },
  },
  async created() {
    await this.fetchCompanies()
    await this.fetchCities()
    await this.fetchJobtitles()
    await this.fetchSeniorities()
  },
  methods: {
    onChangeJobTitle() {
      this.$store.dispatch('filterJob', this.myfilterjob)
      this.$store.dispatch('getRatings', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.filterseniority,
        city: this.filtercity,
      })
      this.$store.dispatch('fetchAverage', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.filterseniority,
        city: this.filtercity,
      })
      this.$store.commit('ratings/SETPAGE', 1)
    },
    onChangeSeniority() {
      this.$store.dispatch('filterSeniority', this.myfilterseniority)
      this.$store.dispatch('getRatings', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.filterseniority,
        city: this.filtercity,
      })
      this.$store.dispatch('fetchAverage', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.myfilterseniority,
        city: this.filtercity,
      })
      this.$store.commit('ratings/SETPAGE', 1)
    },
    onChangeCompany() {
      this.$store.dispatch('filterCompany', this.myfiltercompany)
      this.$store.dispatch('getRatings', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.filterseniority,
        city: this.filtercity,
      })
      this.$store.dispatch('fetchAverage', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.filterseniority,
        city: this.filtercity,
      })
      this.$store.commit('ratings/SETPAGE', 1)
    },
    onChangeCity() {
      this.$store.dispatch('filterCity', this.myfiltercity)
      this.$store.dispatch('getRatings', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.filterseniority,
        city: this.filtercity,
      })
      this.$store.dispatch('fetchAverage', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.filterseniority,
        city: this.filtercity,
      })
      this.$store.commit('ratings/SETPAGE', 1)
    },
    async fetchCompanies() {
      await this.$store.dispatch('getCompanies')
    },
    async fetchCities() {
      await this.$store.dispatch('getCities')
    },
    async fetchJobtitles() {
      await this.$store.dispatch('getJobtitles')
    },
    async fetchSeniorities() {
      await this.$store.dispatch('getSeniorities')
    },
  },
})
</script>
