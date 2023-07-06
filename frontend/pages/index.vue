<template>
  <main>
    <div class="container mx-auto md:w-7/12 px-4 md:px-0">
      <div class="pt-10 lg:pt-14">
        <div class="flex items-center justify-center flex-col leading-10">
          <h2 class="font-bold text-2xl md:text-4xl text-center text-black">
            {{ $t('home_title') }}
          </h2>
          <h4 style="color: #7694A0;" class="text-sm md:text-base text-center lg:w-4/5 font-[500] text-gray-700 pt-4">
            {{ $t('home_page_description') }}
          </h4>
          <NuxtLink to="/add_job"
                    class="cursor-pointer p-4 text-white font-bold flex items-center justify-center lg:text-base bg-blueDark h-12 rounded-lg mt-8 text-sm w-full md:w-44">
            {{ $t('home_page_cta') }}
          </NuxtLink>
        </div>
      </div>
      <div class="pt-14 md:pt-28">
        <div class="flex flex-col xl:flex-row items-center justify-center">
          <div class="xl:w-1/4">
            <h4 class="font-bold text-lg flex items-center">
              {{ $t("label_top_jobs") }}
            </h4>
          </div>
          <div class="flex flex-col xl:flex-row xl:w-3/4 xl:justify-end">
            <p
              class="w-64 font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-gray-700">
              {{ $t("label_dont_miss_jobs") }}
            </p>
            <div class="pt-4 md:pt-2 xl:pt-0  w-full xl:w-40 xl:ml-4 flex items-center justify-center">
              <button @click="showModal = true"
                      class="cursor-pointer p-4 text-blueDark w-full font-bold flex items-center justify-center text-sm md:text-lg px-4 border-2 h-12 rounded-md border-blue">
                {{ $t("global_label_subscribe") }}
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="site__main-company py-3">
        <div class="grid grid-cols-1 md:grid-cols-2">
          <div
            v-for="(item, index) in jobslist"
            :key="index"
            class="bg-white py-6 px-7 border h-44 md:h-24"
          >
            <div>
              <div class="pt-4 xl:pt-0 flex flex-col">
                <h4 class="text-lg font-bold">{{ item.job_title }} <span class="text-gray-600 font-light"> in {{
                    item.city
                  }}</span></h4>
                <div class="flex flex-col xl:flex-row pt-2">
                  <div class="pt-1 xl:pt-0 flex flex-row items-center">
                    <img class="w-auto h-4" :src="location" />
                    <p class="text-xs mx-1 text-gray-500 ml-1">
                      {{ item.company_name }}
                    </p>
                  </div>
                  <div class="pt-1 xl:pt-0 xl:ml-2 flex flex-row items-center">
                    <img class="w-auto h-4" :src="position" />
                    <p class="text-xs mx-1 text-gray-500 ml-1">
                      {{ item.application_url }}
                    </p>
                  </div>
                  <div class="pt-1 xl:pt-0 xl:ml-2 flex flex-row items-center">
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
        <More
          :title="$t('label_see_more_jobs')"
          path="/jobs"
        />
      </div>
      <div class="pt-12 md:pt-20">
        <div class="flex flex-col xl:flex-row items-center justify-center">
          <div class="md:w-1/4">
            <h4 class="font-bold text-lg flex items-center">
              {{ $t("label_top_salaries") }}
            </h4>
          </div>
          <div class="flex flex-col xl:flex-row xl:w-3/4 xl:justify-end">
            <p
              class="font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-center text-gray-700">
              {{ $t("label_contribute_your_salary") }}
            </p>
            <div class="pt-4 md:pt-2 xl:pt-0 w-full xl:w-40 xl:ml-4">
              <NuxtLink to="/add_salary"
                        class="cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-xs md:text-lg px-4 border-2 h-12 rounded-md border-blue">
                <img :src="plusIcon" class="w-5 h-5 mr-1"/>
                {{ $t('btn_label_contribute') }}
              </NuxtLink>
            </div>
          </div>
        </div>
        <div class="py-2 md:py-3">
          <Company isHomePage/>
          <More :title="$t('label_see_more_salaries')" path="/salaries"/>
        </div>
      </div>
      <div class="rounded-md bg-white p-4 my-2 md:my-3">
        <div class="rounded-md m-2 shadow-lg flex p-4 flex-col md:flex-row items-center text-gray-500 justify-center">
          <img :src="money" class="w-14 h-14"/>
          <p class="text-sm text-medium text-gray-500 text-center md:text-left pt-2 md:pt-0">
            {{ $t("home_page_ad_0") }}
            <NuxtLink class="text-blueDark" to="#">{{ $t("home_page_ad_1") }}</NuxtLink>
            {{ $t("home_page_ad_2") }}
          </p>
        </div>
      </div>
    </div>
    <SubmitModal v-show="showModal" @close-modal="showModal = false"/>
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
    await this.fetchJobs();
    await this.fetchRatings();
  },
  methods: {
    async fetchJobs() {
      await this.$store.dispatch('getJobs', {
        page: this.page,
        limit: this.limit

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
