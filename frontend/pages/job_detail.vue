<template>
  <main>
    <div class="container mx-auto px-4 md:px-0">
      <div class="pt-10 lg:pt-14 flex flex-col md:flex-row">
        <div
          class="bg-white cursor-pointer shadow-md p-3 md:p-4 rounded-md hover:shadow-xl w-full md:w-8/12"
        >
          <div class="flex flex-col xl:flex-row">
            <div
              class="w-32 h-32 border rounded-md flex justify-center items-center"
            >
              <img class="w-auto h-24" :src="image" />
            </div>
            <div class="xl:w-9/12 pt-4 xl:pt-0 xl:ml-8 flex flex-col">
              <h4 class="text-lg font-light">
                Ejara is hiring
                <span class="text-gray-600 font-bold">Flutter Developer</span>
              </h4>
              <div class="flex flex-col md:flex-row py-2">
                <div class="flex">
                  <img :src="location" class="w-4 h-4 mr-1" />
                  <p class="text-xs font-extralight text-gray-600">Ejara</p>
                </div>
                <div class="xl:ml-4 pt-2 xl:pt-0 flex">
                  <img :src="position" class="w-4 h-4 mr-1" />
                  <p class="text-xs font-extralight text-gray-600">Onsite</p>
                </div>
                <div class="xl:ml-4 pt-2 xl:pt-0 flex">
                  <img :src="clock" class="w-4 h-4 mr-1" />
                  <p class="text-xs font-extralight text-gray-600">2 days</p>
                </div>
              </div>
              <div class="pt-1 flex flex-wrap">
                <!-- <p v-for="(item, index) in tags.split(',')" :key="index" class="text-xs m-2 text-gray-500">
                                    <span class="bg-gray-200 rounded-full px-4 py-1">{{ item }}</span>
                                </p> -->
                <p
                  v-for="(item, index) in tags"
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
                    About {{ structure }}
                  </h4>
                  <p class="text-sm font-extralight pt-4 text-gray-600">
                    {{ description }}
                    We believe the world needs a place where people can exchange
                    their loved goods. That's why we built Mercari—to connect
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
                    <span class="font-bold text-black">Apply at: </span>
                    devjobs@ejara.com
                  </p>
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
            <div
              class="xl:justify-end justify-center items-start w-full xl:w-3/12 pt-4 flex"
            >
              <NuxtLink
                to="#"
                :class="`cursor-pointer px-16 text-white font-bold items-center justify-center text-sm lg:text-base bg-primaryGray h-12 rounded flex justify-center'`"
              >
                Apply
              </NuxtLink>
            </div>
          </div>
        </div>

        <div class="w-full md:w-3/12 md:ml-12 mt-8 md:mt-0">
          <div class="bg-white p-3 md:p-4 rounded-md shadow-md">
            <h4
              class="font-semibold text-xl xl:text-3xl text-center text-black"
            >
              Subscribe
            </h4>
            <p class="font-light text-xs leading-8 pt-2 text-center">
              Subscribe to receive the fresh jobs JobSika
            </p>
            <div class="mt-10">
              <div class="mx-7">
                <h4 class="text-lg font-semibold text-grayC text-left">
                  Enter your email
                </h4>
                <p style="font-size: 9px" class="pt-2 font-normal text-left">
                  All new job offers will be sent to your email
                </p>
                <input
                  type="text"
                  class="border border-grayC w-full rounded-md h-11 mt-2"
                />
              </div>
              <div class="flex items-center justify-center my-12">
                <div
                  class="cursor-pointer border-2 border-blueDark p-4 text-blueDark font-bold flex items-center justify-center text-sm h-12 w-full mx-7"
                >
                  Sign up
                </div>
              </div>
            </div>
          </div>
          <p class="font-light text-xs leading-8 py-4 text-center md:text-left">
            Subscribe to receive the fresh jobs JobSika
          </p>
          <div class="flex flex-row items-center space-x-6">
            <NuxtLink
              v-for="(item, index) in socials"
              :key="index"
              :to="item.link"
            >
              <img :src="item.icon" alt="img" class="w-5 h-5" />
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: 'JobsDetail',
  layout: 'app',
  data() {
    return {
      image: require('../assets/job.png'),
      tags: [
        'research scientist',
        'virology',
        'molecular biology',
        'JSON',
        'Figma',
        'Backend',
      ],
      location: require('../assets/location.png'),
      position: require('../assets/position.png'),
      clock: require('../assets/clock.png'),
      arrow: require('../assets/arrowd.png'),
      socials: [
        {
          icon: require('../assets/twitter.png'),
          link: '#twitter',
        },
        {
          icon: require('../assets/telegram.png'),
          link: '#telegram',
        },
        {
          icon: require('../assets/facebook.png'),
          link: '#facebook',
        },
      ],
    }
  },
}
</script>
