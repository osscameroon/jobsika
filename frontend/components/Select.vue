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
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        v-model="myfilterseniority"
        @change="onChangeSeniority"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          All Seniorities
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          any
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="(seniority, index) in seniorities"
          :key="index"
          :value="seniority"
          v-show="seniority !== ''"
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
        Job Title
      </h4>
      <select
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        v-model="myfilterjob"
        @change="onChangeJobTitle"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          All titles
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          any
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="(link, index) in jobtitles"
          :key="index"
          :value="link"
          v-show="link !== ''"
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
        Companies
      </h4>
      <select
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        v-model="myfiltercompany"
        @change="onChangeCompany"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          All Compagnies
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          any
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="link in companies"
          :key="link.id"
          :value="link.name"
          v-show="link.name !== ''"
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
        Cities
      </h4>
      <select
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        v-model="myfiltercity"
        @change="onChangeCity"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          All Cities
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          any
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="(city, index) in cities"
          :key="index"
          :value="city"
          v-show="city !== ''"
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
      })
      this.$store.dispatch('fetchAverage', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
      })
      this.$store.commit('ratings/SETPAGE', 1)
    },
    onChangeSeniority() {
      this.$store.dispatch('filterSeniority', this.myfilterseniority)
      this.$store.dispatch('getRatings', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.filterseniority,
      })
      this.$store.dispatch('fetchAverage', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        seniority: this.myfilterseniority,
      })
      this.$store.commit('ratings/SETPAGE', 1)
    },
    onChangeCompany() {
      this.$store.dispatch('filterCompany', this.myfiltercompany)
      this.$store.dispatch('getRatings', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
      })
      this.$store.dispatch('fetchAverage', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
      })
      this.$store.commit('ratings/SETPAGE', 1)
    },
    onChangeCity() {
      this.$store.dispatch('filterCity', this.myfiltercity)
      this.$store.dispatch('getRatings', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
        city: this.filtercity,
      })
      this.$store.dispatch('fetchAverage', {
        company: this.filtercompany,
        jobtitle: this.filterjob,
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
