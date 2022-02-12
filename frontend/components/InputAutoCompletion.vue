<template>
  <div class="site__input w-full">
    <p
      class="text-xs md:text-sm font-bold"
      style="color: #b1b1b1; font-family: 'Inter', sans-serif"
    >
      {{ title }}
    </p>
    <input
      v-model="name"
      type="text"
      class="site__input-field border-none mt-3 w-full rounded-md mb-8 md:mb-16"
      :style="myStyle"
      @change="onChange"
    />
    <ul
      style="background: white"
      class="
        h-20
        md:h-32
        overflow-y-scroll
        rounded-lg
        -mt-3
        md:-mt-14
        mb-2
        md:mb-4
        cursor-pointer
      "
      v-show="isOpen"
    >
      <li
        :v-if="endPoint === 'cities'"
        v-for="(city, index) in filterCityByName"
        :key="index"
        @click="setResult(city)"
        class="shadow-sm py-2 px-4"
      >
        {{ city }}
      </li>
      <li
        :v-if="endPoint === 'jobtitles'"
        v-for="(job, index) in filterJobByName"
        :key="index"
        @click="setResult(job)"
        class="shadow-sm py-2 px-4"
      >
        {{ job }}
      </li>
      <li
        :v-if="endPoint === 'seniority'"
        v-for="(seniority, index) in filterSeniorityByName"
        :key="index"
        @click="setResult(seniority)"
        class="shadow-sm py-1 px-4"
      >
        {{ seniority }}
      </li>
      <li
        :v-if="endPoint === 'companies'"
        v-for="company in filterCompanyByName"
        :key="company.id"
        @click="setResult(company.name)"
        class="shadow-sm py-2 px-4"
      >
        {{ company.name }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'
import { BASE_URL } from '../constants/api'
export default {
  name: 'JobTitleAutoCompletion',
  props: ['title', 'myStyle', 'endPoint'],
  data() {
    return {
      name: '',
      isOpen: false,
      jobtitles: [],
      companies: [],
      seniorities: [],
      cities: [],
    }
  },
  computed: {
    filterJobByName() {
      return this.jobtitles.filter((job) => !job.indexOf(this.name))
    },
    filterCityByName() {
      return this.cities.filter((city) => !city.indexOf(this.name))
    },
    filterSeniorityByName() {
      return this.seniorities.filter(
        (seniority) => !seniority.indexOf(this.name)
      )
    },
    filterCompanyByName() {
      return this.companies.filter(
        (company) => !company.name.indexOf(this.name)
      )
    },
  },
  async created() {
    try {
      if (this.endPoint === 'companies') {
        this.companies = (await axios.get(BASE_URL + `/${this.endPoint}`)).data
      }
      if (this.endPoint === 'cities') {
        this.cities = (await axios.get(BASE_URL + `/${this.endPoint}`)).data
      }
      if (this.endPoint === 'jobtitles') {
        this.jobtitles = (await axios.get(BASE_URL + `/${this.endPoint}`)).data
      }
      if (this.endPoint === 'seniority') {
        this.seniorities = (
          await axios.get(BASE_URL + `/${this.endPoint}`)
        ).data
      }
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    onChange() {
      if (this.name === '') {
        this.jobtitles = []
        this.companies = []
        this.seniorities = []
        this.cities = []
      } else {
        this.$emit('input', this.name)
        this.isOpen = true
      }
    },
    setResult(result) {
      this.name = result
      this.isOpen = false
    },
  },
}
</script>
