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
                title="Salary  (FCFA)"
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
                  :grade="newRating.rating"
                  :maxStars="5"
                  :hasCounter="true"
                  @changeGrade="setGrade"
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
      isCompanyNameSetted: false,
      isJobTitleSetted: false,
      isCitySetted: false,
    }
  },
  computed: {
    companies(){
      return this.$store.state.companies.companies
    },
    jobtitles(){
      return this.$store.state.jobtitles.jobtitles
    },
    cities(){
      return this.$store.state.cities.cities
    },
    seniorities(){
      return this.$store.state.seniorities.seniorities
    },
    companyNames(){
      return this.companies.map(elem => elem.name);
    },
    filteredCompanyNames(){
      if(this.newRating.company_name === ''){
        return []
      }else{
        return this.companyNames.filter(
          elem => {
            return !elem.toLowerCase().indexOf(this.newRating.company_name.toLowerCase())
          }
        );
      }
    },
    filteredJobTitles(){
      if(this.newRating.job_title === ''){
        return []
      }else{
        return this.jobtitles.filter(
          elem => {
            return !elem.toLowerCase().indexOf(this.newRating.job_title.toLowerCase())
          }
        );
      }
    },
    filteredCities(){
      if(this.newRating.city === ''){
        return []
      }else{
        return this.cities.filter(
          elem => {
            return !elem.toLowerCase().indexOf(this.newRating.city.toLowerCase())
          }
        );
      }
    },
    companyComplation(){
      return this.filteredCompanyNames.length > 0 && !this.isCompanyNameSetted
    },
    jobTitleComplation(){
      return this.filteredJobTitles.length > 0 && !this.isJobTitleSetted
    },
    cityComplation(){
      return this.filteredCities.length > 0 && !this.isCitySetted
    }
  },
  async created() {
    try {
      await this.fetchCompanies();
      await this.fetchJobtitles();
      await this.fetchCities();
      await this.fetchSeniorities();

    } catch (err) {
      console.log(err)
    }
  },
  methods: {
    goToList() {
      this.$router.push('/')
    },
    setCompanyName(name){
      this.newRating.company_name = name;
      this.isCompanyNameSetted = true
    },
    setJobTitle(job){
      this.newRating.job_title = job;
      this.isJobTitleSetted = true
    },
    setCity(city){
      this.newRating.city = city;
      this.isCitySetted = true
    },
    async fetchCompanies() {
      await this.$store.dispatch('getCompanies');
    },
    async fetchJobtitles() {
      await this.$store.dispatch('getJobtitles');
    },
    async fetchCities() {
      await this.$store.dispatch('getCities');
    },
    async fetchSeniorities() {
      await this.$store.dispatch('getSeniorities');
    },
    setGrade(value){
      console.log("call star with", value);
      this.newRating.rating = value;
    },
    addRating(){
      console.log("new :", this.newRating);
      if (this.newRating){
        this.$store.dispatch('postRating', this.newRating)
        this.$router.push('/')
      }
    },
  },
}
</script>
