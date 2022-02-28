<template>
  <div
    class="bg-primary py-4 md:py-8 flex items-center justify-between sm:px-6"
  >
    <div class="flex-1 flex justify-center sm:hidden">
      <span
        @click="startpage()"
        style="border-color: #919191"
        class="
          relative
          inline-flex
          items-center
          px-3
          py-2
          rounded-md
          border
          bg-white
          text-xs
          md:text-sm
          font-medium
          text-gray-500
          hover:bg-gray-50
          cursor-pointer
          mx-2
        "
      >
        <span style="color: #919191" class="font-bold">&laquo;</span>
      </span>

      <span
        style="font-family: 'Inter', sans-serif"
        :class="`
          relative
          inline-flex
          items-center
          px-4
          py-1
          md:py-2
          border border-gray-300
          text-xs
          md:text-sm
          font-medium
          rounded-md
          text-gray-700
          cursor-pointer
          ${
            page === 1
              ? 'bg-blue text-white border-none'
              : 'bg-blueDark text-white border-none'
          }
          `"
        @click="previewpage()"
      >
        Previous
      </span>
      <span
        style="font-family: 'Inter', sans-serif"
        :class="`
          ml-3
          relative
          inline-flex
          items-center
          px-4
          py-1
          md:py-2
          text-xs
          md:text-sm
          font-medium
          rounded-md
          cursor-pointer
          ${
            page === numberPage
              ? 'bg-blue text-white border-none'
              : 'bg-blueDark text-white border-none'
          }
          `"
        @click="nextpage()"
      >
        Next
      </span>
      <span
        @click="limitpage()"
        style="border-color: #919191"
        class="
          relative
          inline-flex
          items-center
          px-3
          py-2
          rounded-md
          border
          bg-white
          text-xs
          md:text-sm
          font-medium
          text-gray-500
          hover:bg-gray-50
          cursor-pointer
          mx-2
        "
      >
        <span style="color: #919191" class="font-bold">&raquo;</span>
      </span>
    </div>
    <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
      <div>
        <p
          v-if="nbOfItems == 0"
          style="font-family: 'Inter', sans-serif"
          class="text-xs md:text-sm text-gray-700"
        >
          No items found
        </p>
        <p
          v-else
          style="font-family: 'Inter', sans-serif"
          class="text-xs md:text-sm text-gray-700"
        >
          Showing
          <span class="font-medium">{{ (page - 1) * limit + 1 }}</span>
          to
          <span class="font-medium">{{ limit * (page - 1) + nbOfItems }}</span>
          of
          <span class="font-medium">{{ nbHits }}</span>
          results
        </p>
      </div>
      <div>
        <nav
          class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px"
          aria-label="Pagination"
        >
          <span
            @click="startpage()"
            class="
              relative
              inline-flex
              items-center
              px-4
              py-2
              rounded-l-md
              border border-gray-300
              bg-white
              text-xs
              md:text-sm
              font-medium
              text-gray-500
              hover:bg-gray-50
              cursor-pointer
            "
          >
            <span class="font-bold">&laquo;</span>
          </span>
          <div @click="previewpage()">
            <span
              :class="`z-10
                bg-indigo-50
                border-indigo-500
                text-indigo-600
                relative
                inline-flex
                items-center
                px-2
                py-2
                border
                text-xs
                md:text-sm
                font-medium
                cursor-pointer
                ${page === current1 ? 'bg-blue' : 'bg-white'}
                `"
            >
              <span class="sr-only">Previous</span>
              <svg
                class="h-5 w-5"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
              >
                <path
                  fill-rule="evenodd"
                  d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                  clip-rule="evenodd"
                />
              </svg>
            </span>
          </div>
          <div
            v-for="current1 in leftSide"
            :key="current1"
            @click="changepage(current1)"
          >
            <span
              :class="`z-10
                bg-indigo-50
                border-indigo-500
                text-indigo-600
                relative
                inline-flex
                items-center
                px-4
                py-2
                border
                text-xs
                md:text-sm
                font-medium
                cursor-pointer
                ${page === current1 ? 'bg-blue' : 'bg-white'}
                `"
            >
              {{ current1 }}
            </span>
          </div>
          <div>
            <span
              class="
                z-10
                bg-blue
                border-indigo-500
                text-indigo-600
                relative
                inline-flex
                items-center
                px-4
                py-2
                border
                text-xs
                md:text-sm
                font-medium
                cursor-pointer
              "
            >
              {{ page }}
            </span>
          </div>
          <div
            v-for="current in rightSide"
            :key="current"
            @click="changepage(current)"
          >
            <span
              :class="`
                z-10
                bg-indigo-50
                border-indigo-500
                text-indigo-600
                relative
                inline-flex
                items-center
                px-4
                py-2
                border
                text-xs
                md:text-sm
                font-medium
                cursor-pointer
                ${page === current1 ? 'bg-blue' : 'bg-white'}
              `"
            >
              {{ current }}
            </span>
          </div>
          <div @click="nextpage()">
            <span
              :class="`
                z-10
                bg-indigo-50
                border-indigo-500
                text-indigo-600
                relative
                inline-flex
                items-center
                px-2
                py-2
                border
                text-xs
                md:text-sm
                font-medium
                cursor-pointer
                ${page === current1 ? 'bg-blue' : 'bg-white'}
              `"
            >
              <span class="sr-only">Next</span>
              <svg
                class="h-5 w-5"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
              >
                <path
                  fill-rule="evenodd"
                  d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                  clip-rule="evenodd"
                />
              </svg>
            </span>
          </div>
          <span
            @click="limitpage()"
            class="
              relative
              inline-flex
              items-center
              px-4
              py-2
              rounded-r-md
              border border-gray-300
              bg-white
              text-xs
              md:text-sm
              font-medium
              text-gray-500
              hover:bg-gray-50
              cursor-pointer
            "
          >
            <span class="font-bold">&raquo;</span>
          </span>
        </nav>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Pagination_Number',
  data() {
    return {
      current1: '',
    }
  },
  computed: {
    page() {
      return this.$store.state.ratings.page
    },
    limit() {
      return this.$store.state.ratings.limit
    },
    nbHits() {
      return this.$store.state.ratings.nbHits
    },
    nbOfItems() {
      return this.$store.state.ratings.ratings.length
    },
    numberPage() {
      return Math.ceil(this.nbHits / this.limit)
    },
    leftSide() {
      const result = []
      for (let i = 5; i >= 1; i--) {
        if (this.page - i > 0) {
          result.push(this.page - i)
        }
      }
      return result
    },
    rightSide() {
      const result = []
      for (let i = 1; i <= 5; i++) {
        if (this.page + i < this.numberPage) {
          result.push(this.page + i)
        }
      }
      return result
    },
    filterjob() {
      return this.$store.state.ratings.filterjob
    },
    filtercompany() {
      return this.$store.state.ratings.filtercompany
    },
  },
  methods: {
    async fetchRatings() {
      await this.$store.dispatch('getRatings', {
        page: this.page,
        limit: this.limit,
        company: this.filtercompany,
        jobtitle: this.filterjob,
      })
    },
    async setpage(value) {
      await this.$store.commit('ratings/SETPAGE', value)
    },
    changepage(value) {
      this.setpage(value)
      this.fetchRatings()
    },
    nextpage() {
      if (this.page + 1 <= this.numberPage) {
        this.changepage(this.page + 1)
      }
    },
    previewpage() {
      if (this.page - 1 > 0) {
        this.changepage(this.page - 1)
      }
    },
    startpage() {
      this.changepage(1)
    },
    limitpage() {
      this.changepage(this.numberPage)
    },
  },
  async created() {
    await this.fetchRatings();
    console.log('number of page', this.numberPage);
  },
}
</script>
