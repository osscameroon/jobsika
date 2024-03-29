<template>
  <div
    class="bg-white shadow-md p-3 md:p-4 rounded-md hover:shadow-xl"
    @mouseover="active = true"
    @mouseleave="active = false"
  >
    <div
      class="flex flex-col justify-center items-center md:items-start md:flex-row"
    >
      <div class="w-28 h-28 border rounded-md flex justify-center items-center">
        <img class="w-auto h-24" :src="picture" />
      </div>
      <div
        class="md:w-7/12 pt-4 md:pt-0 md:ml-8 flex flex-col items-center md:items-start"
      >
        <h4 class="text-lg font-bold text-center md:text-left">
          {{ title }} <span class="text-gray-600 font-light"> {{ city }}</span>
        </h4>
        <div class="flex flex-col space-x-4 md:flex-row">
          <div class="flex pt-2 md:pt-0">
            <img :src="location" class="w-4 h-4" />
            <p class="text-xs font-extralight ml-2 text-gray-600">
              {{ structure }}
            </p>
          </div>
          <div class="flex pt-2 md:pt-0">
            <img :src="position" class="w-4 h-4" />
            <p class="text-xs font-extralight ml-2 text-gray-600">
              {{ marker }}
            </p>
          </div>
          <div class="flex pt-2 md:pt-0">
            <img :src="clock" class="w-4 h-4" />
            <p class="text-xs font-extralight ml-2 text-gray-600">
              {{ new Date(time).toLocaleDateString() }}
            </p>
          </div>
          <div v-if="opened.includes(details[myIndex].id)" class="w-full">
            <div class="my-4">
              <h4 class="text-gray-800 py-4 font-semibold">
                About {{ structure }}
              </h4>
              <p class="text-sm font-extralight pt-4 text-gray-600">
                {{ description }}
              </p>
              <p class="text-sm font-extralight pt-4 text-gray-600">
                <span class="font-bold text-black">Employment type: </span>Full
                time
              </p>
              <p class="text-sm font-extralight pt-4 text-gray-600">
                <span class="font-bold text-black">Apply at: </span>{{ mail }}
              </p>
              <p class="text-sm font-extralight pt-4 text-gray-600">
                <span class="font-bold text-black">Salary range: </span>Min({{
                  minSalary
                }}
                FCFA), Max({{ maxSalary }} FCFA)
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
        <p
          class="text-xs pt-2 font-extralight text-gray-600 text-center md:text-left w-10/12"
        >
          {{ cleanText(description) }}
        </p>
        <div class="pt-1 flex flex-wrap">
          <p
            v-for="(item, index) in tags.split(',')"
            :key="index"
            class="text-xs m-2 text-gray-500"
          >
            <span class="bg-gray-200 rounded-full px-4 py-1">{{ item }}</span>
          </p>
        </div>
      </div>
      <div
        class="md:justify-end items-center w-full md:w-3/12 pt-4 md:pt-0 flex"
      >
        <NuxtLink
          to="#"
          :class="`cursor-pointer px-12 py-4 text-white font-bold items-center justify-center text-sm md:text-base bg-blueDark h-12 rounded-lg ${
            active == true ? 'flex' : 'hidden'
          }`"
        >
          Apply
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'JobComponent',
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
      type: [Array],
      default: null,
    },
  },
  data() {
    return {
      location: require('../assets/location.png'),
      position: require('../assets/position.png'),
      clock: require('../assets/clock.png'),
      arrow: require('../assets/arrowd.png'),
      active: false,
      opened: [],
    }
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
    cleanText(text) {
      return text.slice(0, 250) + (text.length > 250 ? '..' : '')
    },
  },
}
</script>
