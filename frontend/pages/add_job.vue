<template>
  <main>
    <div class="container mx-auto md:w-7/12 px-4 md:px-0">
      <div class="pt-10 lg:pt-20">
        <h2 class="font-bold text-2xl md:text-4xl text-center text-black">
          Add a job
        </h2>
      </div>
      <div>
        <form @submit.prevent="saveNewJob">
          <div class="flex flex-col md:w-7/12 mx-auto">
            <div class="grid gap-7 grid-cols-1 mt-10">
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    Company Name
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
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
                      class="font-normal h-6 text-gray-700"
                      style="font-size: 12px"
                    >
                      Enter your company's brand/trade name: without Inc., Ltd.,
                      B.V., Pte., etc.
                    </p>
                  </div>
                </div>
                <input
                  :value="newJob.company_name"
                  type="text"
                  style="height: 61px"
                  class="border border-grayC mt-2 w-full rounded-md mb-4"
                  @input="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        company_name: event.target.value,
                      })
                  "
                />
              </div>
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    Company Email
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[1].id) }"
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
                      class="font-normal h-6 text-gray-700"
                      style="font-size: 12px"
                    >
                      Enter your company's Email:
                    </p>
                  </div>
                </div>
                <input
                  :value="newJob.company_email"
                  type="text"
                  style="height: 61px"
                  class="border border-grayC mt-2 w-full rounded-md mb-4"
                  @input="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        company_email: event.target.value,
                      })
                  "
                />
              </div>
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    Job title
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[2].id) }"
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
                      class="font-normal h-6 text-gray-700"
                      style="font-size: 12px"
                    >
                      This field requires you to enter the title of the position
                      and/or specialization (if applicable). E.g.
                      "Accountant(Audit)"
                    </p>
                  </div>
                </div>
                <input
                  :value="newJob.job_title"
                  type="text"
                  style="height: 61px"
                  class="border border-grayC mt-2 w-full rounded-md mb-4"
                  @input="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        job_title: event.target.value,
                      })
                  "
                />
              </div>
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    Description of role
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[3].id) }"
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
                      class="font-normal h-6 text-gray-700"
                      style="font-size: 12px"
                    >
                      Please add a brief description of the role you are hiring
                      for, make it straight forward and concise
                    </p>
                  </div>
                </div>
                <textarea
                  :value="newJob.description"
                  cols="6"
                  rows="5"
                  type="text"
                  class="border border-grayC mt-2 w-full rounded-md mb-4"
                  @input="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        description: event.target.value,
                      })
                  "
                >
                </textarea>
              </div>
              <div>
                <div class="flex mb-2">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    City
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[4].id) }"
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
                    <p class="font-normal h-14" style="font-size: 12px">
                      The City of the Job
                    </p>
                  </div>
                </div>
                <select
                  :value="newJob.city"
                  class="form-select mt-2 mb-4 appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-grayC rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
                  style="height: 61px"
                  aria-label="Default select example"
                  @input="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        city: event.target.value,
                      })
                  "
                >
                  <option class="text-xs md:text-sm" value="">
                    Select the City
                  </option>
                  <option
                    v-for="city in cities"
                    :key="city"
                    :value="city"
                    class="text-xs md:text-sm"
                  >
                    {{ city }}
                  </option>
                </select>
              </div>
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    Location
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[5].id) }"
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
                      class="font-normal h-10 text-gray-700"
                      style="font-size: 12px"
                    >
                      This field requires you to enter the town you are/were
                      employed in. For remote workers, please enter the town you
                      are/were based in.
                    </p>
                  </div>
                </div>
                <select
                  :value="newJob.is_remote"
                  class="mt-2 mb-4 form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-grayC rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
                  style="height: 61px"
                  aria-label="Default select example"
                  @change="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        is_remote: event.target.value === 'true' ? true : false,
                      })
                  "
                >
                  <option class="text-xs md:text-sm" :value="true">
                    Remote
                  </option>
                  <option class="text-xs md:text-sm" :value="false">
                    On Site
                  </option>
                </select>
              </div>
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    Application email
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[6].id) }"
                      @click="toggle(tooltips[6].id)"
                    >
                      !
                    </span>
                  </div>
                </div>
                <div
                  v-if="opened.includes(tooltips[6].id)"
                  class="w-full bg-primary"
                >
                  <div class="w-full my-1">
                    <p
                      class="font-normal h-10 text-gray-700"
                      style="font-size: 12px"
                    >
                      This field requires you to enter the pre-tax/gross salary
                      you get, vacation money included. You can find this amount
                      in your contract or on your paycheck.
                    </p>
                  </div>
                </div>
                <input
                  :value="newJob.application_email_address"
                  type="text"
                  style="height: 61px"
                  class="border border-grayC mt-2 w-full rounded-md mb-4"
                  @input="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        application_email_address: event.target.value,
                      })
                  "
                />
              </div>
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    Phone Number
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[7].id) }"
                      @click="toggle(tooltips[7].id)"
                    >
                      !
                    </span>
                  </div>
                </div>
                <div
                  v-if="opened.includes(tooltips[7].id)"
                  class="w-full bg-primary"
                >
                  <div class="w-full my-1">
                    <p
                      class="font-normal h-10 text-gray-700"
                      style="font-size: 12px"
                    >
                      This field requires you to enter the pre-tax/gross salary
                      you get, vacation money included. You can find this amount
                      in your contract or on your paycheck.
                    </p>
                  </div>
                </div>
                <input
                  :value="newJob.application_phone_number"
                  type="text"
                  style="height: 61px"
                  class="border border-grayC mt-2 w-full rounded-md mb-4"
                  @input="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        application_phone_number: event.target.value,
                      })
                  "
                />
              </div>
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    Salary range
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[8].id) }"
                      @click="toggle(tooltips[8].id)"
                    >
                      !
                    </span>
                  </div>
                </div>
                <div
                  v-if="opened.includes(tooltips[8].id)"
                  class="w-full bg-primary"
                >
                  <div class="w-full my-1">
                    <p
                      class="font-normal h-10 text-gray-700"
                      style="font-size: 12px"
                    >
                      This field requires you to enter the pre-tax/gross salary
                      you get, vacation money included. You can find this amount
                      in your contract or on your paycheck.
                    </p>
                  </div>
                </div>
                <div class="flex flex-col md:flex-row">
                  <input
                    :value="newJob.salary_range_min"
                    type="number"
                    style="height: 61px"
                    class="border border-grayC mt-2 w-full md:w-1/2 rounded-md mb-4 md:mr-2"
                    placeholder="Minimum"
                    @input="
                      (event) =>
                        $store.commit('jobs/SETNEWJOB', {
                          ...newJob,
                          salary_range_min: event.target.value,
                        })
                    "
                  />
                  <input
                    :value="newJob.salary_range_max"
                    type="number"
                    style="height: 61px"
                    class="border border-grayC mt-2 w-full md:w-1/2 rounded-md mb-4 md:ml-2"
                    placeholder="Maximum"
                    @input="
                      (event) =>
                        $store.commit('jobs/SETNEWJOB', {
                          ...newJob,
                          salary_range_max: event.target.value,
                        })
                    "
                  />
                </div>
              </div>
              <div>
                <div class="flex">
                  <p
                    class="text-xs font-bold md:text-lg"
                    style="color: #b1b1b1"
                  >
                    How to apply
                  </p>
                  <div class="flex items-center">
                    <span
                      class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                      :class="{ opened: opened.includes(tooltips[9].id) }"
                      @click="toggle(tooltips[9].id)"
                    >
                      !
                    </span>
                  </div>
                </div>
                <div
                  v-if="opened.includes(tooltips[9].id)"
                  class="w-full bg-primary"
                >
                  <div class="w-full my-1">
                    <p
                      class="font-normal h-6 text-gray-700"
                      style="font-size: 12px"
                    >
                      How will you want people to apply for this job? Enter the
                      email or any other infos that will help people apply for
                      the job.
                    </p>
                  </div>
                </div>
                <textarea
                  :value="newJob.how_to_apply"
                  cols="6"
                  rows="5"
                  type="text"
                  class="border border-grayC mt-2 w-full rounded-md mb-4"
                  @input="
                    (event) =>
                      $store.commit('jobs/SETNEWJOB', {
                        ...newJob,
                        how_to_apply: event.target.value,
                      })
                  "
                ></textarea>
              </div>
              <div>
                <p
                  class="font-normal h-8 md:w-3/4 text-gray-700"
                  style="font-size: 12px"
                >
                  Please upload your logo. If no logo is upload , a default logo
                  will be used. Using a personalised logo increases your job
                  offer visibility by 60%
                </p>
                <div
                  class="w-40 h-32 bg-white rounded-lg mt-3 flex items-center"
                >
                  <div
                    class="flex justify-center items-center mx-auto cursor-pointer"
                  >
                    <img :src="uploadImg" class="w-6 h-6" />
                    <h4 class="ml-1 text-base font-semibold text-grayC">
                      Upload logo
                    </h4>
                  </div>
                </div>
              </div>
            </div>
            <div
              class="mt-10 flex flex-col md:flex-row items-center justify-center xl:justify-start"
            >
              <button type="submit">
                <div class="w-full md:w-1/5 ml-0 pt-6 md:pt-0 md:ml-40">
                  <Button
                    show-picture="nothing"
                    my-style="background: #235365;padding: 20px 80px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);"
                    name="Next"
                  />
                </div>
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: 'JobsIndex',
  layout: 'app',
  data() {
    return {
      opened: [],
      uploadImg: require('../assets/upload.png'),
      checkImg: require('../assets/payment.png'),
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
        {
          id: 8,
          name: 'rate',
        },
        {
          id: 9,
          name: 'rate',
        },
        {
          id: 10,
          name: 'rate',
        },
      ],
    }
  },
  computed: {
    cities() {
      return this.$store.state.cities.cities
    },
    newJob() {
      return this.$store.state.jobs.newjob
    },
  },
  async created() {
    await this.fetchCities()
  },
  methods: {
    toggle(id) {
      const index = this.opened.indexOf(id)
      if (index > -1) {
        this.opened.splice(index, 1)
      } else {
        this.opened.push(id)
      }
    },
    async fetchCities() {
      await this.$store.dispatch('getCities')
    },
    saveNewJob() {
      if (this.newJob) {
        this.$store.dispatch('setNewJob', {
          ...this.newJob,
          salary_range_min: parseFloat(this.newJob.salary_range_min),
          salary_range_max: parseFloat(this.newJob.salary_range_max),
        })
        this.$router.push('/confirm_job')
      }
    },
  },
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
