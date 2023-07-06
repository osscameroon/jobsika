<template>
    <main>
        <div class="container mx-auto md:w-7/12 px-4 md:px-0">
            <div class="pt-10 lg:pt-14">
                <h2 class="font-bold text-2xl md:text-4xl text-center text-black">
                    {{ $t('jobs_page_title') }}
                </h2>
            </div>
            <div class="pt-8 md:pt-16">
                <div class="flex flex-col xl:flex-row xl:justify-end">
                    <p
                        class="font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-center text-gray-700">
                        {{ $t('jobs_page_cta') }}
                    </p>
                    <div class="pt-4 md:pt-2 xl:pt-0 w-full xl:w-36 xl:ml-4">
                        <button @click="showModal = true"
                            class="mx-auto xl:mx-0 w-1/2 xl:w-36 cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-xs md:text-lg px-4 border-2 h-12 rounded-md border-blue">
                            {{ $t('global_label_subscribe') }}
                        </button>
                    </div>
                </div>
            </div>
            <div class="pt-8">
                <div v-for="(item, index) in jobs" :key="index" class="mb-6">
                    <JobPost :picture="image" :title="item.job_title" :city="item.city" :structure="item.company_name"
                        :marker="item.is_remote ? 'Remote' : 'Local' " :time="item.createdat" :description="item.description" :tags="item.tags" />
                </div>
            </div>
          <JobPagination />
        </div>
      </div>
      <div class="pt-8">
        <div v-for="(item, index) in jobs" :key="index" class="mb-6">
          <JobPost
            :details="jobs"
            :my-index="index"
            :picture="image"
            :title="item.job_title"
            :city="item.city"
            :structure="item.company_name"
            min-salary="150 000"
            max-salary="200 000"
            :marker="item.is_remote ? 'Remote' : 'Local'"
            :time="item.createdat"
            :description="item.description"
            :tags="item.tags"
            mail="devjobs@ejara.com"
          />
        </div>
      </div>
      <JobPagination />
    </div>
    <SubmitModal
      v-show="showModal"
      @close-modal="showModal = false"
      @success-modal="showModalSuccess = true"
    />
    <SubmitModalSucces
      v-show="showModalSuccess"
      @close-modal=";(showModalSuccess = false) & (showModal = false)"
    />
  </main>
</template>

<script>
export default {
  name: 'JobsIndex',
  layout: 'app',
  data() {
    return {
      showModal: false,
      showModalSuccess: false,
      image: require('../assets/job.png'),
    }
  },
  computed: {
    jobs() {
      return this.$store.state.jobs.jobs
    },
    limit() {
      return this.$store.state.jobs.limit
    },
  },
  async created() {
    await this.fetchJobs()
  },
  methods: {
    async fetchJobs() {
      await this.$store.dispatch('getJobs', {
        page: this.page,
        limit: this.limit,
      })
    },
  },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
