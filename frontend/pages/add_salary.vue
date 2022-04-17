<template>
  <div>
    <main class="site__main pt-5 lg:pt-0 pb-10 md:pb-20">
      <div class="container mx-auto w-10/12">
        <Title
          title="100% Anonymous"
          font-size="header"
          only-title="onlyTitle"
        />
        <Title
          title="Add a company to the list"
          only-title="onlyTitle"
          content="subtitle"
          my-style="margin-top: 20px; font-family: 'Inter',sans-serif"
        />
        <form @submit.prevent="addRating">
          <div class="site__main-row flex flex-col-reverse md:flex-row">
            <div class="site__main-rowOne w-full md:w-1/2">
              <div class="site__main-input mt-3 md:mt-10 md:w-11/12">
                <div class="flex">
                  <p
                    class="text-xs md:text-sm font-bold"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    Company Name
                  </p>
                  <div class="flex items-center-mt-2">
                    <span
                      class="cursor-pointer h-5 text-center w-5 ml-2 text-grayC rounded-full border border-grayC text-xs"
                      :class="{ opened: opened.includes(tooltips[0].id) }"
                      @click="toggle(tooltips[0].id)"
                    >
                      !
                    </span>
                  </div>
                </div>

                <div
                  v-if="opened.includes(tooltips[0].id)"
                  class="w-full bg-primary"
                >
                  <div class="bg-white w-full p-2 my-3 shadow-sm rounded-sm">
                    <p
                      class="text-xs md:text-sm"
                      style="color: #000000; font-family: 'Inter', sans-serif"
                    >
                      This field requires you to enter the job tittle as stated
                      in your contract. To help you, we have listed some job
                      tittles in the proposition field.
                    </p>
                  </div>
                </div>

                <input
                  :value="newRating.company_name"
                  type="text"
                  style="height: 61px"
                  class="site__input-field border border-grayC mt-2 md:mt-3 w-full rounded-md mb-4 md:mb-16"
                  @input="
                    (event) => (newRating.company_name = event.target.value)
                  "
                  @focus="nameFocus()"
                />
                <ul
                  v-if="companyComplation"
                  style="background: white"
                  class="h-20 md:h-32 overflow-y-scroll rounded-lg -mt-3 md:-mt-14 mb-2 md:mb-4 cursor-pointer"
                >
                  <li
                    v-for="(name, index) in filteredCompanyNames"
                    :key="index"
                    class="shadow-sm py-2 px-4"
                    @click="setCompanyName(name)"
                  >
                    {{ name }}
                  </li>
                </ul>
                <div class="flex">
                  <p
                    class="text-xs md:text-sm font-bold"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    Monthly Salary (FCFA)
                  </p>
                  <div class="flex items-center-mt-2">
                    <span
                      class="cursor-pointer h-5 text-center w-5 ml-2 text-grayC rounded-full border border-grayC text-xs"
                      :class="{ opened: opened.includes(tooltips[0].id) }"
                      @click="toggle(tooltips[1].id)"
                    >
                      !
                    </span>
                  </div>
                </div>
                <div
                  v-if="opened.includes(tooltips[1].id)"
                  class="w-full bg-primary"
                >
                  <div class="bg-white w-full p-2 my-3 shadow-sm rounded-sm">
                    <p
                      class="text-xs md:text-sm"
                      style="color: #000000; font-family: 'Inter', sans-serif"
                    >
                      This field requires you to enter the job tittle as stated
                      in your contract. To help you, we have listed some job
                      tittles in the proposition field.
                    </p>
                  </div>
                </div>
                <input
                  v-model="newRating.salary"
                  type="number"
                  style="height: 61px"
                  class="site__input-field border border-grayC mt-2 md:mt-3 w-full rounded-md"
                  @keypress="onlyNumber"
                  @focus="blurAll()"
                />
                <div class="mb-4 md:mb-16">
                  <notification :message="errorSalary" />
                </div>
                <div class="flex">
                  <p
                    class="text-xs md:text-sm font-bold"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    Comments
                  </p>
                  <div class="flex items-center-mt-2">
                    <span
                      class="cursor-pointer h-5 text-center w-5 ml-2 text-grayC rounded-full border border-grayC text-xs"
                      :class="{ opened: opened.includes(tooltips[0].id) }"
                      @click="toggle(tooltips[2].id)"
                    >
                      !
                    </span>
                  </div>
                </div>
                <div
                  v-if="opened.includes(tooltips[2].id)"
                  class="w-full bg-primary"
                >
                  <div class="bg-white w-full p-2 my-3 shadow-sm rounded-sm">
                    <p
                      class="text-xs md:text-sm"
                      style="color: #000000; font-family: 'Inter', sans-serif"
                    >
                      This field requires you to enter monthly salary you earned
                      in FCFA
                    </p>
                  </div>
                </div>

                <textarea
                  v-model="newRating.comment"
                  type="textarea"
                  rows="5"
                  style="height: 120px"
                  class="resize rounded-md w-full h-full mt-2 md:mt-3 border border-grayC"
                  @focus="blurAll()"
                />
                <div class="site__input-btn mt-10 flex flex-col md:flex-row">
                  <div class="w-full md:w-1/4">
                    <NuxtLink to="/">
                      <Button
                        show-picture="nothing"
                        name="Cancel"
                        my-style="background: white; color: #000000"
                      />
                    </NuxtLink>
                  </div>
                  <button
                    class="w-full md:w-1/4 ml-0 pt-6 md:pt-0 md:ml-6"
                    type="submit"
                  >
                    <Button
                      my-style="background: #235365; font-family: 'Inter', sans-serif"
                      name="Add"
                    />
                  </button>
                </div>
              </div>
            </div>
            <div class="site__main-rowTwo w-full md:w-1/2">
              <div class="site__main-input mt-3 md:mt-10 md:w-11/12">
                <div class="flex">
                  <p
                    class="text-xs md:text-sm font-bold"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    Job Title
                  </p>
                  <div class="flex items-center-mt-2">
                    <span
                      class="cursor-pointer h-5 text-center w-5 ml-2 text-grayC rounded-full border border-grayC text-xs"
                      :class="{ opened: opened.includes(tooltips[0].id) }"
                      @click="toggle(tooltips[3].id)"
                    >
                      !
                    </span>
                  </div>
                </div>
                <div
                  v-if="opened.includes(tooltips[3].id)"
                  class="w-full bg-primary"
                >
                  <div class="bg-white w-full p-2 my-3 shadow-sm rounded-sm">
                    <p
                      class="text-xs md:text-sm"
                      style="color: #000000; font-family: 'Inter', sans-serif"
                    >
                      This field requires you to enter the job tittle as stated
                      in your contract. To help you, we have listed some job
                      tittles in the proposition field.
                    </p>
                  </div>
                </div>
                <input
                  :value="newRating.job_title"
                  type="text"
                  style="height: 61px"
                  class="site__input-field border border-grayC mt-2 md:mt-3 w-full rounded-md"
                  @input="(event) => (newRating.job_title = event.target.value)"
                  @focus="jobTitleFocus()"
                />
                <div class="mb-4 md:mb-16">
                  <notification :message="errorJobtitle" />
                </div>
                <ul
                  v-if="jobTitleComplation"
                  style="background: white"
                  class="h-20 md:h-32 overflow-y-scroll rounded-lg -mt-3 md:-mt-14 mb-2 md:mb-4 cursor-pointer"
                >
                  <li
                    v-for="(job, index) in filteredJobTitles"
                    :key="index"
                    class="shadow-sm py-2 px-4"
                    @click="setJobTitle(job)"
                  >
                    {{ job }}
                  </li>
                </ul>
                <div class="flex">
                  <p
                    class="text-xs md:text-sm font-bold"
                    style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                  >
                    City
                  </p>
                  <div class="flex items-center-mt-2">
                    <span
                      class="cursor-pointer h-5 text-center w-5 ml-2 text-grayC rounded-full border border-grayC text-xs"
                      :class="{ opened: opened.includes(tooltips[0].id) }"
                      @click="toggle(tooltips[4].id)"
                    >
                      !
                    </span>
                  </div>
                </div>
                <div
                  v-if="opened.includes(tooltips[4].id)"
                  class="w-full bg-primary"
                >
                  <div class="bg-white w-full p-2 my-3 shadow-sm rounded-sm">
                    <p
                      class="text-xs md:text-sm"
                      style="color: #000000; font-family: 'Inter', sans-serif"
                    >
                      This field requires you to enter the job tittle as stated
                      in your contract. To help you, we have listed some job
                      tittles in the proposition field.
                    </p>
                  </div>
                </div>

                <input
                  :value="newRating.city"
                  type="text"
                  style="height: 61px"
                  class="site__input-field border border-grayC mt-2 md:mt-3 w-full rounded-md"
                  @input="(event) => (newRating.city = event.target.value)"
                  @focus="cityFocus()"
                />
                <div class="mb-4 md:mb-16">
                  <notification :message="errorCity" />
                </div>
                <ul
                  v-if="cityComplation"
                  style="background: white"
                  class="h-20 md:h-32 overflow-y-scroll rounded-lg -mt-3 md:-mt-14 mb-2 md:mb-4 cursor-pointer"
                >
                  <li
                    v-for="(city, index) in filteredCities"
                    :key="index"
                    class="shadow-sm py-2 px-4"
                    @click="setCity(city)"
                  >
                    {{ city }}
                  </li>
                </ul>
                <div class="site__input w-full">
                  <div class="flex">
                    <p
                      class="text-xs md:text-sm font-bold"
                      style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                    >
                      Seniority
                    </p>
                    <div class="flex items-center-mt-2">
                      <span
                        class="cursor-pointer h-5 text-center w-5 ml-2 text-grayC rounded-full border border-grayC text-xs"
                        :class="{ opened: opened.includes(tooltips[0].id) }"
                        @click="toggle(tooltips[5].id)"
                      >
                        !
                      </span>
                    </div>
                  </div>
                  <div
                    v-if="opened.includes(tooltips[5].id)"
                    class="w-full bg-primary"
                  >
                    <div class="bg-white w-full p-2 my-3 shadow-sm rounded-sm">
                      <p
                        class="text-xs md:text-sm"
                        style="color: #000000; font-family: 'Inter', sans-serif"
                      >
                        This field requires you to enter the job tittle as
                        stated in your contract. To help you, we have listed
                        some job tittles in the proposition field.
                      </p>
                    </div>
                  </div>
                  <select
                    v-model="newRating.seniority"
                    class="mt-2 mb-4 md:mb-8 md:mt-3 form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-grayC rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
                    style="height: 61px; font-family: 'Inter', sans-serif"
                    aria-label="Default select example"
                    @focus="blurAll()"
                  >
                    <option
                      v-for="seniority in seniorities"
                      :key="seniority"
                      :value="seniority"
                      style="font-family: 'Inter', sans-serif"
                      class="text-xs md:text-sm"
                    >
                      {{ seniority }}
                    </option>
                  </select>
                </div>
                <div class="my-3 md:my-0 md:mt-12">
                  <div class="flex site__input w-full">
                    <div class="flex">
                      <p
                        class="text-xs md:text-sm font-bold"
                        style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                      >
                        Rate
                      </p>
                      <div class="flex items-center-mt-2">
                        <span
                          class="cursor-pointer h-5 text-center w-5 ml-2 text-grayC rounded-full border border-grayC text-xs"
                          :class="{ opened: opened.includes(tooltips[0].id) }"
                          @click="toggle(tooltips[6].id)"
                        >
                          !
                        </span>
                      </div>
                    </div>
                    <StarRating
                      :grade="newRating.rating"
                      :max-stars="5"
                      :has-counter="true"
                      @changeGrade="setGrade"
                      @blurall="blurAll()"
                    />
                  </div>
                  <notification :message="errorRating" />
                  <div
                    v-if="opened.includes(tooltips[6].id)"
                    class="w-full bg-primary"
                  >
                    <div class="bg-white w-full p-2 my-3 shadow-sm rounded-sm">
                      <p
                        class="text-xs md:text-sm"
                        style="color: #000000; font-family: 'Inter', sans-serif"
                      >
                        This field requires you to enter the job tittle as
                        stated in your contract. To help you, we have listed
                        some job tittles in the proposition field.
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </form>
      </div>
    </main>
  </div>
</template>

<script>
import Button from '../components/Button.vue'
import Notification from '~/components/notification.vue'
export default {
  name: 'AddSalary',
  components: {
    Button,
    Notification,
  },
  layout: 'addSalary',
  data() {
    return {
      isOpen: false,
      errorRating: '',
      errorSalary: '',
      errorJobtitle: '',
      errorCity: '',
      opened: [],
      starsPicture: require('../assets/star.png'),
      newRating: {
        company_name: '',
        salary: '',
        city: '',
        seniority: 'Senior',
        rating: 0,
        comment: '',
        job_title: '',
      },
      isCompanyNameFocus: false,
      isJobTitleFocus: false,
      isCityFocus: false,
      tooltips: [
        {
          id: 1,
          name: 'company',
        },
        {
          id: 2,
          name: 'monthly',
        },
        {
          id: 3,
          name: 'comment',
        },
        {
          id: 4,
          name: 'jobtitle',
        },
        {
          id: 5,
          name: 'city',
        },
        {
          id: 6,
          name: 'seniority',
        },
        {
          id: 7,
          name: 'rate',
        },
      ],
    }
  },
  computed: {
    companies() {
      return this.$store.state.companies.companies
    },
    jobtitles() {
      return this.$store.state.jobtitles.jobtitles
    },
    cities() {
      return this.$store.state.cities.cities
    },
    seniorities() {
      return this.$store.state.seniorities.seniorities
    },
    companyNames() {
      return this.companies.map((elem) => elem.name)
    },
    filteredCompanyNames() {
      if (this.newRating.company_name === '') {
        return []
      } else {
        return this.companyNames.filter((elem) => {
          return !elem
            .toLowerCase()
            .indexOf(this.newRating.company_name.toLowerCase())
        })
      }
    },
    filteredJobTitles() {
      if (this.newRating.job_title === '') {
        return []
      } else {
        return this.jobtitles.filter((elem) => {
          return !elem
            .toLowerCase()
            .indexOf(this.newRating.job_title.toLowerCase())
        })
      }
    },
    filteredCities() {
      if (this.newRating.city === '') {
        return []
      } else {
        return this.cities.filter((elem) => {
          return !elem.toLowerCase().indexOf(this.newRating.city.toLowerCase())
        })
      }
    },
    companyComplation() {
      return this.filteredCompanyNames.length > 0 && this.isCompanyNameFocus
    },
    jobTitleComplation() {
      return this.filteredJobTitles.length > 0 && this.isJobTitleFocus
    },
    cityComplation() {
      return this.filteredCities.length > 0 && this.isCityFocus
    },
  },
  async created() {
    try {
      await this.fetchCompanies()
      await this.fetchJobtitles()
      await this.fetchCities()
      await this.fetchSeniorities()
    } catch (err) {
      console.log(err)
    }
  },
  methods: {
    goToList() {
      this.$router.push('/')
    },
    setCompanyName(name) {
      this.newRating.company_name = name
      this.isCompanyNameFocus = false
    },
    setJobTitle(job) {
      this.newRating.job_title = job
      this.isJobTitleFocus = false
    },
    setCity(city) {
      this.newRating.city = city
      this.isCityFocus = false
    },
    async fetchCompanies() {
      await this.$store.dispatch('getCompanies')
    },
    async fetchJobtitles() {
      await this.$store.dispatch('getJobtitles')
    },
    async fetchCities() {
      await this.$store.dispatch('getCities')
    },
    async fetchSeniorities() {
      await this.$store.dispatch('getSeniorities')
    },
    setGrade(value) {
      this.newRating.rating = value
    },
    blurAll() {
      this.isJobTitleFocus = false
      this.isCompanyNameFocus = false
      this.isCityFocus = false
    },
    nameFocus() {
      this.blurAll()
      this.isCompanyNameFocus = true
    },
    jobTitleFocus() {
      this.blurAll()
      this.isJobTitleFocus = true
    },
    cityFocus() {
      this.blurAll()
      this.isCityFocus = true
    },
    addRating() {
      if (this.newRating) {
        if (String(this.newRating.salary).length === 0) {
          this.errorSalary = 'This field cannot be empty'
          return
        }
        if (String(this.newRating.job_title).length === 0) {
          this.errorJobtitle = 'This field cannot be empty'
          return
        }
        if (String(this.newRating.city).length === 0) {
          this.errorCity = 'This field cannot be empty'
          return
        }
        if (this.newRating.rating === 0) {
          this.errorRating = 'This field cannot be empty'
        } else {
          this.$store.dispatch('postRating', this.newRating)
          this.$router.push('/')
        }
      }
    },
    onlyNumber($event) {
      const keyCode = $event.keyCode ? $event.keyCode : $event.which
      if ((keyCode < 48 || keyCode > 57) && keyCode !== 46) {
        $event.preventDefault()
      }
    },
    toggle(id) {
      const index = this.opened.indexOf(id)
      if (index > -1) {
        this.opened.splice(index, 1)
      } else {
        this.opened.push(id)
      }
    },
  },
}
</script>
