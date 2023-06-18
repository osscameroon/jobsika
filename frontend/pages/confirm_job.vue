<template>
  <main>
    <div class="container mx-auto md:w-7/12 px-4 md:px-0">
      <div class="pt-10 md:pt-20 pb-5 md:pb-10">
        <h2 class="font-bold text-2xl md:text-4xl text-center text-black">
          Confirm job information
        </h2>
      </div>
      <div>
        <div class="bg-white p-3 md:p-4 rounded-md shadow-xl">
          <div class="flex flex-col xl:flex-row">
            <div
              class="w-28 h-28 border rounded-md flex justify-center items-center"
            >
              <img class="w-auto h-24" :src="image" />
            </div>
            <div class="xl:w-7/12 pt-4 xl:pt-0 xl:ml-8 flex flex-col">
              <h4 class="text-lg font-bold">
                {{ title }}
                <span class="text-gray-600 font-light"> {{ city }}</span>
              </h4>
              <div class="flex flex-col md:flex-row py-2">
                <div class="flex">
                  <img :src="location" class="w-4 h-4 mr-1" />
                  <!-- <p class="text-xs font-extralight text-gray-600">
                                        {{ structure }}
                                    </p> -->
                  <p class="text-xs font-extralight text-gray-600">Ejara</p>
                </div>
                <div class="xl:ml-4 pt-2 xl:pt-0 flex">
                  <img :src="position" class="w-4 h-4 mr-1" />
                  <!-- <p class="text-xs font-extralight text-gray-600">{{ marker }}</p> -->
                  <p class="text-xs font-extralight text-gray-600">Onsite</p>
                </div>
                <div class="xl:ml-4 pt-2 xl:pt-0 flex">
                  <img :src="clock" class="w-4 h-4 mr-1" />
                  <p class="text-xs font-extralight text-gray-600">
                    <!-- {{ new Date(time).toLocaleDateString() }} -->
                    1/5/2022
                  </p>
                </div>
              </div>
              <div class="pt-1 flex flex-wrap">
                <p
                  v-for="(item, index) in mydetails"
                  :key="index"
                  class="text-xs m-2 text-gray-500"
                >
                  <span class="bg-gray-200 rounded-full px-4 py-1">{{
                    item
                  }}</span>
                </p>
              </div>
              <div class="w-full">
                <div class="my-4">
                  <h4 class="text-gray-800 py-4 font-semibold">
                    <!-- About {{ structure }} -->
                    About Mercari
                  </h4>
                  <p class="text-sm font-extralight pt-4 text-gray-600">
                    <!-- {{ description }} -->
                    We believe the world needs a place where people can exchange
                    their loved goods. That’s why we built Mercari—to connect
                    people who want to buy and sell things of emotional value
                    through their mobile phones. Mercari has grown to become a
                    leading community-powered marketplace that brings together
                    millions of people every day.
                  </p>
                  <p class="text-sm font-extralight pt-4 text-gray-600">
                    <span class="font-bold text-black">Employment type: </span
                    >Full time
                  </p>
                  <p class="text-sm font-extralight pt-4 text-gray-600">
                    <!-- <span class="font-bold text-black">Apply at: </span>{{ mail }} -->
                    <span class="font-bold text-black">Apply at: </span
                    >devjobs@ejara.com
                  </p>
                  <!-- <p class="text-sm font-extralight pt-4 text-gray-600">
                                        <span class="font-bold text-black">Salary range: </span>Min({{
                                            minSalary
                                        }}
                                        FCFA), Max({{ maxSalary }} FCFA)
                                    </p> -->
                  <p class="text-sm font-extralight pt-4 text-gray-600">
                    <span class="font-bold text-black">Salary range: </span
                    >Min(xxx FCFA), Max(xxx FCFA)
                  </p>
                </div>
                <div class="my-8">
                  <NuxtLink
                    to="#"
                    class="cursor-pointer w-60 px-4 py-1 text-gray-500 font-bold flex items-center justify-center text-xs bg-gray-200 h-10 rounded-full hover:bg-gray-400 hover:text-white"
                  >
                    <p class="flex items-center justify-center">
                      Learn more about the job offer<span
                        ><img class="ml-1 w-2 h-auto" alt="pic" :src="arrow"
                      /></span>
                    </p>
                  </NuxtLink>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="flex flex-col md:flex-row md:justify-center mt-10">
          <div class="lg:pt-0 w-full lg:w-44 lg:mr-2">
            <NuxtLink
              to="/add_job"
              class="cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-sm lg:text-base border-2 h-12 rounded-lg border-blue"
            >
              Cancel
            </NuxtLink>
          </div>
          <div class="pt-4 lg:pt-0 w-full lg:w-44 lg:ml-2">
            <NuxtLink
              to="/payment_instruction"
              class="cursor-pointer p-4 text-white font-bold flex items-center justify-center text-sm lg:text-base bg-blueDark h-12 rounded-lg"
            >
              Confirm
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: 'ConfirmPage',
  layout: 'app',
  props: {
    title: {
      type: String,
      default: '',
    },
    picture: {
      type: String,
      default: '',
    },
    city: {
      type: String,
      default: '',
    },
    structure: {
      type: String,
      default: '',
    },
    marker: {
      type: String,
      default: '',
    },
    time: {
      type: String,
      default: '',
    },
    minSalary: {
      type: String,
      default: '',
    },
    maxSalary: {
      type: String,
      default: '',
    },
    description: {
      type: String,
      default: '',
    },
    mail: {
      type: String,
      default: '',
    },
    tags: {
      type: [String],
      default: null,
    },
    myIndex: {
      type: Number,
      default: 0,
    },
    details: {
      type: [Object],
      default: null,
    },
  },
  data() {
    return {
      location: require('../assets/location.png'),
      position: require('../assets/position.png'),
      clock: require('../assets/clock.png'),
      arrow: require('../assets/arrowd.png'),
      image: require('../assets/job.png'),
      mydetails: ['research scientist', 'virology', 'molecular biology'],
    }
  },
  methods: {},
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
