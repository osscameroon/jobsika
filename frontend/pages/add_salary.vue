<template>
  <div>
    <main class="site__main pt-5 lg:pt-0 pb-10 md:pb-20">
      <div class="container mx-auto w-10/12">
        <Title title="100% Anonymous" fontSize="header" onlyTitle="onlyTitle" />
        <Title
          title="Add a company to the list"
          onlyTitle="onlyTitle"
          content="subtitle"
          myStyle="margin-top: 20px; font-family: 'Inter',sans-serif"
        />
        <div class="site__main-row flex flex-col-reverse md:flex-row">
          <div class="site__main-rowOne w-full md:w-1/2">
            <div class="site__main-input mt-3 md:mt-10 md:w-11/12">
              <InputAutoCompletion
                title="Company Name"
                myStyle="height: 61px;"
                endPoint="companies"
              />
              <MyInput
                title="Monthly Salary  (FCFA)"
                myStyle="height: 61px;"
                content="subtitle"
                myInput="input"
                typeInput="number"
              />
              <MyInput
                title="Comments"
                myStyle="height: 120px;"
                content="subtitle"
              />
              <div class="site__input-btn mt-10 flex flex-col md:flex-row">
                <div class="w-full md:w-1/4">
                  <NuxtLink to="/">
                    <Button
                      showPicture="nothing"
                      name="Cancel"
                      myStyle="background: white; color: #000000"
                    />
                  </NuxtLink>
                </div>
                <div
                  @click="addRating"
                  class="w-full md:w-1/4 ml-0 pt-6 md:pt-0 md:ml-6"
                >
                  <Button
                    myStyle="background: #235365; font-family: 'Inter', sans-serif"
                    name="Add"
                  />
                </div>
              </div>
            </div>
          </div>
          <div class="site__main-rowTwo w-full md:w-1/2">
            <div class="site__main-input mt-3 md:mt-10 md:w-11/12">
              <InputAutoCompletion
                title="Job title"
                myStyle="height: 61px;"
                endPoint="jobtitles"
              />
              <InputAutoCompletion
                title="City"
                myStyle="height: 61px;"
                endPoint="cities"
              />
              <div class="site__input w-full">
                <p
                  class="text-xs md:text-sm font-bold"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Seniority
                </p>
                <select
                  class="
                    mt-2
                    mb-4
                    md:mb-8 md:mt-3
                    form-select
                    appearance-none
                    block
                    w-full
                    px-3
                    py-1.5
                    text-gray-700
                    bg-white bg-clip-padding bg-no-repeat
                    border-none
                    rounded
                    transition
                    ease-in-out
                    focus:text-gray-700
                    focus:bg-white
                    focus:border-blue-600
                    focus:outline-none
                    text-xs
                    md:text-sm
                    font-bold
                  "
                  style="height: 61px; font-family: 'Inter', sans-serif"
                  v-model="newRating.seniority"
                  aria-label="Default select example"
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
              <div class="site__input w-full flex my-3 md:my-0 md:mt-12">
                <p
                  class="text-xs md:text-sm font-bold"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Rate
                </p>
                <StarRating
                  :grade="selectvaluestars"
                  :maxStars="5"
                  :hasCounter="true"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import axios from 'axios'
import Button from '../components/Button.vue'
import { BASE_URL } from '../constants/api'
export default {
  name: 'AddSalary',
  components: {
    Button,
  },
  layout: 'app',
  data() {
    return {
      starsPicture: require('../assets/star.png'),
      seniorities: [],
      newRating: {
        company_name: '',
        salary: 0,
        city: '',
        seniority: 'Senior',
        rating: 3,
        comment: '',
        job_title: '',
      },
    }
  },
  computed: {
    selectvaluecompany() {
      return this.$store.state.ratings.selectvaluecompany
    },
    selectvaluejob() {
      return this.$store.state.ratings.selectvaluejob
    },
    selectvaluecity() {
      return this.$store.state.ratings.selectvaluecity
    },
    selectvalueseniority() {
      return this.newRating.seniority
    },
    selectvaluesalary() {
      return parseInt(this.$store.state.ratings.selectvaluesalary)
    },
    selectvaluecomment() {
      return this.$store.state.ratings.selectvaluecomment
    },
    selectvaluestars() {
      return this.$store.state.ratings.selectvaluestars
    },
  },
  async created() {
    try {
      this.seniorities = (await axios.get(BASE_URL + '/seniority')).data
    } catch (err) {
      console.log(err)
    }
  },
  methods: {
    goToList() {
      this.$router.push('/')
    },
    addRating() {
      if (this.newRating) {
        this.$store.dispatch('postCompany', {
          company_name: this.selectvaluecompany,
          salary: this.selectvaluesalary,
          city: this.selectvaluecity,
          seniority: this.selectvalueseniority,
          rating: this.selectvaluestars,
          comment: this.selectvaluecomment,
          job_title: this.selectvaluejob,
        })
        this.$router.push('/')
      }
    },
  },
}
</script>
