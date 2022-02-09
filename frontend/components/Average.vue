<template>
  <div>
    <div class="text-center py-4 md:py-10 font-bold text-sm" v-if="errored">
      <p class="text-red-500">
        We're sorry, we're not able to retrieve this information at the moment,
        please try back later
      </p>
    </div>
    <div
      v-else
      style="background: #235365; font-family: 'Inter', sans-serif"
      class="site__average py-2 md:py-4 opacity-60 rounded-md my-3"
    >
      <div class="flex flex-row justify-between items-center">
        <div class="site__average__title w-full md:w-1/6 ml-10">
          <p
            class="text-white text-xs lg:text-sm font-medium"
            style="font-family: 'Inter', sans-serif"
          >
            {{ title }}
          </p>
        </div>
        <div
          class="
            site__average__content
            w-full
            flex flex-col
            md:flex-row
            items-center
            md:w-1/2
            ml-10
            mr-2
            md:ml-0 md:mr-0
          "
        >
          <p
            class="
              text-white text-xs
              lg:text-sm
              font-medium
              md:ml-10
              lg:ml-20
              text-center
            "
            style="font-family: 'Inter', sans-serif"
          >
            {{ salaryRating.salary }} FCFA
          </p>
          <div class="flex pt-1 md:pt-0 ml-0 md:ml-24 md:mr-10">
            <div
              v-for="(item, index) in salaryRating.rating"
              :key="index"
              class="flex"
            >
              <img class="w-4 h-4 mr-1" :src="starsPicture" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import axios from 'axios'
import Vue from 'vue'
import { BASE_URL } from '../constants/api'
export default Vue.extend({
  name: 'AverageComponent',
  props: ['title'],
  data() {
    return {
      starsPicture: require('../assets/star.png'),
      salaryRating: '',
      loading: true,
      errored: false,
    }
  },
  async created() {
    try {
      this.salaryRating = (await axios.get(BASE_URL + '/average-rating')).data
    } catch (err) {
      console.log(err)
      this.errored = true
    }
  },
})
</script>
