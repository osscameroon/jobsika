<template>
  <main>
    <div class="container mx-auto md:w-11/12 px-4 md:px-0">
      <div class="pt-10 lg:pt-14">
        <div class="flex items-center justify-center flex-col leading-10">
          <h2 class="font-bold text-2xl md:text-4xl text-center text-black">
            Find the perfect talent for your project
          </h2>
          <h4 style="color: #7694A0;" class="text-sm md:text-base text-center lg:w-4/5 font-[500] text-gray-700 pt-4">
            Are you looking for the perfect talent to help you achieve your
            project goals? Look no further - the right person is out there and
            ready to take on the challenge.
          </h4>
          <NuxtLink to="/add_job"
            class="cursor-pointer p-4 text-white font-bold flex items-center justify-center lg:text-base bg-blueDark h-12 rounded-lg mt-8 text-sm w-full md:w-44">
            Post a job now
          </NuxtLink>
        </div>
      </div>
      <div class="pt-14 md:pt-28">
        <div class="flex flex-col xl:flex-row items-center justify-center">
          <div class="xl:w-1/4">
            <h4 class="font-bold text-lg flex items-center">
              Top jobs
            </h4>
          </div>
          <div class="flex flex-col xl:flex-row xl:w-3/4 xl:justify-end">
            <p
              class="w-64 font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-gray-700">
              Don't miss out on any job
            </p>
            <div class="pt-4 md:pt-2 xl:pt-0  w-full xl:w-40 xl:ml-4 flex items-center justify-center">
              <button @click="showModal = true"
                class="cursor-pointer p-4 text-blueDark w-full font-bold flex items-center justify-center text-sm md:text-lg px-4 border-2 h-12 rounded-md border-blue">
                Subscribe
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="site__main-company py-3">
        <div class="grid grid-cols-1 md:grid-cols-2">
          <div v-for="(item, index) in jobs" :key="index" class="bg-white py-6 px-7 border h-44 md:h-24">
            <div>
              <div class="pt-4 xl:pt-0 flex flex-col">
                <h4 class="text-lg font-bold">{{ item.title }} <span class="text-gray-600 font-light"> {{
                  item.city
                }}</span></h4>
                <div class="flex flex-col xl:flex-row pt-2">
                  <div class="pt-1 xl:pt-0 flex flex-row items-center">
                    <img class="w-auto h-4" :src="location"/>
                    <p class="text-xs mx-1 text-gray-500 ml-1">
                      {{ item.structure }}
                    </p>
                  </div>
                  <div class="pt-1 xl:pt-0 xl:ml-2 flex flex-row items-center">
                    <img class="w-auto h-4" :src="position"/>
                    <p class="text-xs mx-1 text-gray-500 ml-1">
                      {{ item.site }}
                    </p>
                  </div>
                  <div class="pt-1 xl:pt-0 xl:ml-2 flex flex-row items-center">
                    <img class="w-auto h-4" :src="clock"/>
                    <p class="text-xs mx-1 text-gray-500 ml-1">
                      {{ item.delay }}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>

        </div>
        <More />
      </div>
      <div class="pt-12 md:pt-20">
        <div class="flex flex-col xl:flex-row items-center justify-center">
          <div class="md:w-1/4">
            <h4 class="font-bold text-lg flex items-center">
              Top salaries
            </h4>
          </div>
          <div class="flex flex-col xl:flex-row xl:w-3/4 xl:justify-end">
            <p
              class="font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-center text-gray-700">
              Contribute your salary to help create awareness
            </p>
            <div class="pt-4 md:pt-2 xl:pt-0 w-full xl:w-40 xl:ml-4">
              <NuxtLink to="/add_salary"
                class="cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-xs md:text-lg px-4 border-2 h-12 rounded-md border-blue">
                <img :src="plusIcon" class="w-5 h-5 mr-1" />
                Contribute
              </NuxtLink>
            </div>
          </div>
        </div>
        <div class="py-2 md:py-3">
          <Company />
          <More />
        </div>
      </div>
      <div class="rounded-md bg-white p-4 my-2 md:my-3">
        <div class="rounded-md m-2 shadow-lg flex p-4 flex-col md:flex-row items-center text-gray-500 justify-center">
          <img :src="money" class="w-14 h-14" />
          <p class="text-sm text-medium text-gray-500 text-center md:text-left pt-2 md:pt-0">
            Are you a UX Designer near Douala? <NuxtLink class="text-blueDark" to="#">Add your salary</NuxtLink> to help
            others like yourself
            make more informed decisions on where they go to work, including better decisions around pay and
            compensation.
          </p>
        </div>
      </div>
    </div>
    <SubmitModal v-show="showModal" @close-modal="showModal = false" />
  </main>
</template>

<script>
export default {
  name: 'IndexPage',
  layout: 'app',
  data() {
    return {
      showModal: false,
      plusIcon: require('../assets/plus_dark.png'),
      location: require('../assets/location.png'),
      position: require('../assets/position.png'),
      clock: require('../assets/clock.png'),
      money: require('../assets/money.png'),
      tags: ["Figma", "Nuxt", "UI/UX Design"],
      jobs: [
        {
          title: "Flutter Developer",
          city: "in Douala",
          site: "Hybrid and remote",
          delay: "2 Months",
          structure: "Mboa Digital",
        },
        {
          title: "Senior compliance audit manager",
          site: "Onsite",
          delay: "11 days",
          city: "in Yaounde",
          structure: "MTN",
        },
        {
          title: "Frontend Dev in Douala",
          city: "in Douala",
          site: "Onsite",
          delay: "11 days",
          structure: "Bjongo hub",
        },
        {
          title: "Frontend Developer",
          city: "in Yaounde",
          site: "Onsite",
          delay: "2 Months",
          structure: "OSS Cameroon",
        },
        {
          title: "Frontend Dev in Douala",
          city: "in Douala",
          site: "Onsite",
          delay: "11 days",
          structure: "Bjongo hub",
        },
        {
          title: "Frontend Developer",
          city: "in Yaounde",
          site: "Onsite",
          delay: "2 Months",
          structure: "OSS Cameroon",
        },
      ]
    }
  },
}
</script>

