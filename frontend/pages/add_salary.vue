<template>
  <div>
    <main class="site__main pt-6 md:pt-12 pb-10 md:pb-20">
      <div class="container mx-auto w-10/12">
        <Title
          title="Add a company to the list"
          font-size="header"
          only-title="onlyTitle"
        />
        <p
          class="text-center text-xs sm:text-base pt-2 md:pt-4 font-normal"
          style="color: #2e2e2e; font-family: 'Inter', sans-serif"
        >
          NB: Your entery is anonymous. Please ensure that it is fair, accurate
          and honest.
        </p>
        <form @submit.prevent="addRating">
          <div class="site__main-row flex flex-col md:w-7/12 mx-auto">
            <div class="site__main-row w-full mt-5 md:mt-16">
              <div class="flex">
                <p
                  class="text-xs font-bold md:text-lg"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Company Name
                </p>
                <div class="flex items-center">
                  <span
                    class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs"
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
                <div class="w-full my-1">
                  <p
                    class="font-light text-xs leading-5"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    We understand you might not want to share the name of the
                    company despite the field being anonymized-please enter a
                    description of the company (e.g. "local startup",
                    "well-known bank" or legacy company, etc)
                  </p>
                </div>
              </div>

              <input
                :value="newRating.company_name"
                type="text"
                style="height: 61px"
                class="site__input-field border border-grayC mt-2 w-full rounded-md mb-4 md:mb-16"
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
                  class="text-xs md:text-lg font-bold"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Job Title
                  <span style="color: red">&#42;</span>
                </p>
                <div class="flex items-center">
                  <span
                    class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs"
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
                <div class="w-full my-1">
                  <p
                    class="font-light text-xs leading-5"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    This field requires you to enter the title of the position
                    and/or specialization (if applicable). E.g.
                    "Accountant(Audit)"
                  </p>
                </div>
              </div>
              <input
                :value="newRating.job_title"
                type="text"
                style="height: 61px"
                class="site__input-field border border-grayC mt-2 w-full rounded-md"
                @input="(event) => (newRating.job_title = event.target.value)"
                @focus="jobTitleFocus()"
              />
              <div class="mb-4 md:mb-16 mt-2">
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
                  class="text-xs md:text-lg font-bold"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Seniority
                </p>
                <div class="flex items-center">
                  <span
                    class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs"
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
                <div class="w-full my-1">
                  <p
                    class="font-light text-xs leading-5"
                    style="color: #000000; font-family: 'Inter', sans-serif"
                  >
                    This field requires you to enter the seniority of the job
                    title you have/had at the company.
                  </p>
                </div>
              </div>
              <select
                v-model="newRating.seniority"
                class="mt-2 mb-4 md:mb-16 form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-grayC rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
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
            <div class="flex">
              <p
                class="text-xs md:text-lg font-bold"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                City
                <span style="color: red">&#42;</span>
              </p>
              <div class="flex items-center">
                <span
                  class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs"
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
              <div class="w-full my-1">
                <p
                  class="font-light text-xs leading-5"
                  style="color: #000000; font-family: 'Inter', sans-serif"
                >
                  This field requires you to enter the town you are/were
                  employed in. For remote workers, please enter the town you
                  are/were based in.
                </p>
              </div>
            </div>

            <input
              :value="newRating.city"
              type="text"
              style="height: 61px"
              class="site__input-field border border-grayC mt-2 w-full rounded-md"
              @input="(event) => (newRating.city = event.target.value)"
              @focus="cityFocus()"
            />
            <div class="mb-4 md:mb-8 mt-2">
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
            <div class="flex">
              <p
                class="text-xs md:text-lg font-bold"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                Monthly Salary (FCFA)
                <span style="color: red">&#42;</span>
              </p>
              <div class="flex items-center">
                <span
                  class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs"
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
              <div class="w-full my-1">
                <p
                  class="font-light leading-5 text-xs"
                  style="color: #000000; font-family: 'Inter', sans-serif"
                >
                  This field requires you to enter the pre-tax/gross salary you
                  get, vacation money included. You can find this amount in your
                  contract or on your paycheck.
                </p>
              </div>
            </div>
            <input
              v-model="newRating.salary"
              style="height: 61px"
              class="site__input-field border border-grayC mt-2 w-full rounded-md"
              @keypress="onlyNumber"
              @focus="blurAll()"
            />
            <div class="mb-2 md:mb-8 mt-2">
              <notification :message="errorSalary" />
            </div>
            <div class="flex">
              <p
                class="text-xs md:text-lg font-bold"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                Rate
              </p>
              <div class="flex items-center">
                <span
                  class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs"
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
            <div
              v-if="opened.includes(tooltips[6].id)"
              class="w-full bg-primary"
            >
              <div class="w-full my-1">
                <p
                  class="font-light text-xs leading-5"
                  style="color: #000000; font-family: 'Inter', sans-serif"
                >
                  This field requires you to rate your overall experience with
                  the company. For example, 1 star could mean you will not
                  recommend anybody to work with that company or five stars
                  could mean your overall experience with the company was
                  excellent and you will gladly recommend anyone to work with
                  them.
                </p>
              </div>
            </div>
            <div class="flex">
              <p
                class="text-xs md:text-lg font-bold"
                style="color: #b1b1b1; font-family: 'Inter', sans-serif"
              >
                Comments
              </p>
              <div class="flex items-center">
                <span
                  class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs"
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
              <div class="w-full my-1">
                <p
                  class="font-light text-xs leading-5"
                  style="color: #000000; font-family: 'Inter', sans-serif"
                >
                  This field requires you to tell us more about your experience
                  in the company. Please share more details about your
                  contribution. Any bad or good things that happened that you
                  want to share with us?
                </p>
              </div>
            </div>

            <textarea
              v-model="newRating.comment"
              type="textarea"
              rows="5"
              style="height: 120px"
              class="resize rounded-md w-full h-full mt-2 border border-grayC mb-4"
              @focus="blurAll()"
            />

            <div
              class="site__input-btn mt-10 flex flex-col md:flex-row items-center justify-center md:justify-end"
            >
              <div class="w-full md:w-1/5">
                <NuxtLink to="/">
                  <Button
                    show-picture="nothing"
                    name="Cancel"
                    my-style="background: white; color: #000000; padding: 15px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);"
                  />
                </NuxtLink>
              </div>
              <button
                class="w-full md:w-1/5 ml-0 pt-6 md:pt-0 md:ml-6"
                type="submit"
              >
                <Button
                  my-style="background: #235365; font-family: 'Inter', sans-serif;padding: 15px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);"
                  name="Add"
                />
              </button>
            </div>
          </div>
        </form>
        <div
          class="site__warning mt-10 shadow-xl rounded-md py-4 md:py-8 px-10 md:mt-24 sm:mt-20 w-full bg-white"
        >
          <div class="flex flex-col md:flex-row md:items-center">
            <img
              :src="warning"
              alt="img"
              class="h-12 w-12 md:h-20 md:w-auto md:mr-5 mx-auto md:mx-0"
            />
            <p
              class="text-xs sm:text-sm mt-4 md:mt-0"
              style="
                color: #000000;
                font-family: 'Inter', sans-serif;
                line-height: 28px;
              "
            >
              We know that a lot of emplyers are uncomfortable sharing
              information that can identify them especially when there is not
              alot information from the company they work or if they are few
              people with thier job tittles at the company. In our will to be as
              anonymous as possible, your contributions on company name or job
              title are only shown if multiple people from the same company or
              same title share their compensation.
            </p>
          </div>
        </div>
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
      errorSalary: '',
      errorJobtitle: '',
      errorCity: '',
      opened: [],
      starsPicture: require('../assets/star.png'),
      warning: require('../assets/warning.png'),
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
        if (String(this.newRating.job_title).length === 0) {
          this.errorJobtitle = 'This field cannot be empty'
          return
        }
        if (String(this.newRating.city).length === 0) {
          this.errorCity = 'This field cannot be empty'
          return
        }
        if (String(this.newRating.salary).length === 0) {
          this.errorSalary = 'This field cannot be empty'
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
