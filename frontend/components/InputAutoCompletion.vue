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
      class="site__input-field border-none mt-3 w-full rounded-md mb-8 md:mb-16"
      :style="myStyle"
    />
    <ul
      v-if="opened"
      style="background: white"
      class="h-20 md:h-32 overflow-y-scroll rounded-lg -mt-3 md:-mt-14 mb-2 md:mb-4 cursor-pointer"
    >
      <li
        v-for="(elem, index) in filteredData"
        :key="index"
        class="shadow-sm py-2 px-4"
        @click="setResult(elem)"
      >
        {{ elem }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'AutoCompletion',
  props: {
    title: {
      required: true,
      type: String,
    },
    /* name: {
      required: true,
      type: String
    }, */
    myStyle: {
      required: true,
      type: String,
    },
    datas: {
      required: true,
      type: [Array],
    },
  },
  data() {
    return {
      name: '',
      isOpen: false,
    }
  },
  computed: {
    filteredData() {
      if (this.name === '') {
        return []
      } else {
        return this.datas
          .map((elem) => elem.toLowerCase())
          .filter((elem) => !elem.indexOf(this.name.toLowerCase()))
      }
    },
    opened() {
      return this.filteredData.length > 0
    },
  },
  created() {},
  methods: {
    setResult(result) {
      this.name = result
      this.isOpen = false
    },
  },
}
</script>
