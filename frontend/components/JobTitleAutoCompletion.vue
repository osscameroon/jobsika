<template>
  <div class="site__input w-full">
    <p
      class="text-xs md:text-sm font-bold"
      style="color: #b1b1b1; font-family: 'Inter', sans-serif"
    >
      {{ title }}
    </p>
    <input
      v-model="name"
      type="text"
      class="
        site__input-field
        border-none
        mt-2
        md:mt-3
        w-full
        rounded-md
        mb-4
        md:mb-16
      "
      :style="myStyle"
      @change="onChange"
    />
    <ul
      style="background: white"
      class="
        h-44
        overflow-y-scroll
        rounded-lg
        -mt-3
        md:-mt-14
        mb-2
        md:mb-4
        cursor-pointer
      "
      v-show="isOpen"
    >
      <li
        v-for="(job, index) in filterJobByName"
        :key="index"
        @click="setResult(job)"
        class="shadow-sm py-2 px-4"
      >
        {{ job }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'
import { BASE_URL } from '../constants/api'
export default {
  name: 'JobTitleAutoCompletion',
  props: ['title', 'myStyle'],
  data() {
    return {
      name: '',
      isOpen: false,
      jobtitles: [],
    }
  },
  computed: {
    filterJobByName() {
      return this.jobtitles.filter((job) => !job.indexOf(this.name))
    },
  },
  async created() {
    try {
      this.jobtitles = (await axios.get(BASE_URL + '/jobtitles')).data
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    onChange() {
      if (this.name === '') {
        this.jobtitles = []
      } else {
        this.$emit('input', this.name)
        this.isOpen = true
      }
    },
    setResult(result) {
      this.name = result
      this.isOpen = false
    },
  },
}
</script>
