<template>
  <div
    class="bg-primary py-4 md:py-8 flex items-center justify-between sm:px-6"
  >
    <div class="flex-1 flex justify-between sm:hidden">
      <a
        href="#"
        style="font-family: 'Inter', sans-serif"
        class="
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
          bg-white
          hover:bg-gray-50
        "
      >
        Previous
      </a>
      <a
        href="#"
        style="font-family: 'Inter', sans-serif"
        class="
          ml-3
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
          bg-white
          hover:bg-gray-50
        "
      >
        Next
      </a>
    </div>
    <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
      <div>
        <p
          style="font-family: 'Inter', sans-serif"
          class="text-xs md:text-sm text-gray-700"
        >
          Showing
          <span class="font-medium">1</span>
          to
          <span class="font-medium">5</span>
          of
          <span class="font-medium">10</span>
          results
        </p>
      </div>
      <div>
        <nav
          class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px"
          aria-label="Pagination"
        >
          <a
            @click="previewpage()"
            href="#"
            class="
              relative
              inline-flex
              items-center
              px-2
              py-2
              rounded-l-md
              border border-gray-300
              bg-white
              text-xs
              md:text-sm
              font-medium
              text-gray-500
              hover:bg-gray-50
            "
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
          </a>
          <div
            v-for="(page, index) in numberPage"
            :key="index"
            @click="changepage(page)"
          >
            <a
              href="#"
              aria-current="page"
              class="
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
              "
            >
              {{ page }}
            </a>
          </div>
          <a
            @click="nextpage()"
            href="#"
            class="
              relative
              inline-flex
              items-center
              px-2
              py-2
              rounded-r-md
              border border-gray-300
              bg-white
              text-xs
              md:text-sm
              font-medium
              text-gray-500
              hover:bg-gray-50
            "
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
          </a>
        </nav>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Pagination_Number',
  computed: {
    page(){
      return this.$store.state.ratings.page;
    },
    limit(){
      return this.$store.state.ratings.limit;
    },
    nbHits(){
      return this.$store.state.ratings.nbHits;
    },
    numberPage(){
      return this.nbHits / this.limit;
    }
    
  },
  methods: {
    async fetchCompanies(){
      await this.$store.dispatch("getCompanies", {
        page: this.page,
        limit: this.limit
      });
    },
    async setpage(value){
      await this.$store.commit("ratings/SETPAGE", value);
    },
    changepage(value){
      this.setpage(value);
      this.fetchCompanies();
    },
    nexpage(){
      if((this.page + 1)<= this.numberPage){
        this.changepage(this.page + 1);
      }
    },
    previewpage(){
      if((this.page - 1) > 0){
        this.changepage(this.page - 1);
      }
    }
  },
  async created(){
    await this.fetchCompanies();
    console.log("number of page", this.numberPage);
  }
}
</script>
