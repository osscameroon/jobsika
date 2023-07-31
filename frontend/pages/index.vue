<template>
  <main>
    <div class="container mx-auto md:w-9/12 lg:w-7/12 px-4 md:px-0">
      <div class="pt-10 lg:pt-14">
        <div class="flex items-center justify-center flex-col leading-10">
          <h2 class="font-bold text-2xl md:text-4xl text-center text-black">
            Find the perfect talent for your project
          </h2>
          <h4 style="color: #7694a0" class="text-sm md:text-base text-center lg:w-4/5 font-[500] text-gray-700 pt-4">
            Are you looking for the perfect talent to help you achieve your
            project goals? Look no further - the right person is out there and
            ready to take on the challenge.
          </h4>
          <NuxtLink
to="/add_job"
            class="cursor-pointer p-4 text-white font-bold flex items-center justify-center lg:text-base bg-blueDark h-12 rounded-lg mt-8 text-sm w-64 px-4 sm:w-44">
            Post a job now
          </NuxtLink>
        </div>
      </div>
      <div class="pt-4 sm:pt-8 md:pt-12 lg:pt-16">
        <div class="flex flex-col xl:flex-row items-center justify-center">
          <div class="xl:w-1/4">
            <h4 class="font-bold text-lg flex items-center">Top jobs</h4>
          </div>
          <div class="flex flex-col xl:flex-row xl:w-3/4 xl:justify-end">
            <p
              class="w-64 font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-gray-700">
              Don't miss out on any job
            </p>
            <div class="pt-4 md:pt-2 xl:pt-0 w-full xl:w-40 xl:ml-4 flex items-center justify-center">
              <button
                class="cursor-pointer p-4 text-blueDark w-full font-bold flex items-center justify-center text-sm md:text-lg px-4 border-2 h-12 rounded-md border-blue my-4"
                @click="showModal = true">
                Subscribe
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="py-3">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-y-4 sm:gap-y-0">
          <div v-for="(item, index) in jobslist" :key="index" class="bg-white py-6 px-7 border h-44 lg:h-32 xl:h-24">
            <div>
              <div class="pt-4 lg:pt-0 flex flex-col">
                <h4 class="text-lg font-bold text-center lg:text-left">
                  {{ item.job_title }}
                  <span class="text-gray-600 font-light">
                    in {{ item.city }}</span>
                </h4>
                <div class="flex flex-col lg:flex-row pt-2">
                  <div class="pt-1 lg:pt-0 flex flex-row justify-center lg:justify-start items-center">
                    <img class="w-auto h-4" :src="location" />
                    <p class="text-xs text-gray-500 ml-2">
                      {{ item.company_name }}
                    </p>
                  </div>
                  <div class="pt-3 lg:pt-0 lg:ml-2 flex flex-row justify-center lg:justify-start items-center">
                    <img class="w-auto h-4" :src="position" />
                    <p class="text-xs text-gray-500 ml-2">
                      {{ item.application_url }}
                    </p>
                  </div>
                  <div class="pt-3 lg:pt-0 lg:ml-2 flex flex-row items-center justify-center lg:justify-start">
                    <img class="w-auto h-4" :src="clock" />
                    <p class="text-xs mx-1 text-gray-500 ml-1">
                      {{ new Date(item.createdat).toLocaleDateString() }}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <More title="See more Jobs" path="/jobs" />
      </div>
      <div class="pt-4 sm:pt-8 md:pt-12 lg:pt-16">
        <div class="flex flex-col xl:flex-row items-center justify-center">
          <div class="md:w-1/4">
            <h4 class="font-bold text-lg flex items-center">Top salaries</h4>
          </div>
          <div class="flex flex-col xl:flex-row xl:w-3/4 xl:justify-end">
            <p
              class="font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-center text-gray-700">
              Contribute your salary to help create awareness
            </p>
            <div class="pt-4 md:pt-2 xl:pt-0 w-full xl:w-40 xl:ml-4">
              <NuxtLink
to="/add_salary"
                class="cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-xs md:text-lg px-4 border-2 h-12 rounded-md border-blue my-4">
                <img :src="plusIcon" class="w-5 h-5 mr-1" />
                Contribute
              </NuxtLink>
            </div>
          </div>
        </div>
        <div class="py-2 md:py-3">
          <Company is-home-page />
          <More title="See more Salaries" path="/salaries" />
        </div>
      </div>
    </div>
    <SubmitModal v-show="showModal" @close-modal="showModal = false" @success-modal="showModalSuccess = true" />
    <SubmitModalSucces v-show="showModalSuccess" @close-modal="; (showModalSuccess = false) & (showModal = false)" />
  </main>
</template>

<script>
export default {
  name: 'IndexPage',
  layout: 'app',
  data() {
    return {
      showModal: false,
      showModalSuccess: false,
      plusIcon: require('../assets/plus_dark.png'),
      location: require('../assets/location.png'),
      position: require('../assets/position.png'),
      clock: require('../assets/clock.png'),
      money: require('../assets/money.png'),
      tags: ['Figma', 'Nuxt', 'UI/UX Design'],
    }
  },
  computed: {
    jobslist() {
      return this.$store.state.jobs.jobs.slice(0, 6)
    },
  },
  async created() {
    await this.fetchJobs()
    await this.fetchRatings()
  },
  methods: {
    async fetchJobs() {
      await this.$store.dispatch('getJobs', {
        page: this.page,
        limit: this.limit,
      })
    },
    async fetchRatings() {
      await this.$store.dispatch('getRatings', {
        page: 1,
      })
    },
  },
}
</script>
