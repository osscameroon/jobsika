<script setup lang="ts">
import { useJobtitlesStore } from "../store/jobtitles";
import { useRatingsStore } from "../store/ratings";
import { useCompaniesStore } from "../store/companies";

const ratingStore = useRatingsStore();
const jobtitleStore = useJobtitlesStore();
const companyStore = useCompaniesStore();

const page = computed(() => ratingStore.page);
const limit = computed(() => ratingStore.limit);
const nbHits = computed(() => ratingStore.nbHits);
const nbOfItems = computed(() => ratingStore.items.length);
const numberPage = computed(() => {
  return Math.ceil(ratingStore.nbHits / ratingStore.limit);
});
const leftSide = computed(() => {
  const result = []
  for (let i = 5; i >= 1; i--) {
    if (page.value - i > 0) {
      result.push(page.value - i)
    }
  }
  return result
});
const rightSide = computed(() => {
  const result = []
  for (let i = 1; i <= 5; i++) {
    if (page.value + i < numberPage.value) {
      result.push(page.value + i)
    }
  }
  return result
});
const filterjob = computed(() => jobtitleStore.filter);
const filtercompany = computed(() => companyStore.filtercompany);

// DATA
const current1: Ref<number> = ref(1);

onMounted(async () => {
  await fetchRatings();
});

// METHODS
const fetchRatings = async () => {
  await ratingStore.getItems({
    page: page.value,
    limit: limit.value,
    company: filtercompany.value,
    jobtitle: filterjob.value,
  });
}

const setpage = (value: number) => {
  ratingStore.page = value;
  current1.value = value;
}

const changepage = async (value: number) => {
  setpage(value);
  await fetchRatings();
}

const nextpage = async () => {
  if (page.value + 1 <= numberPage.value) {
    await changepage(page.value + 1);
  }
}

const previewpage = async () => {
  if (page.value - 1 > 0) {
    await changepage(page.value - 1);
  }
}

const startpage = async () => {
  if (page.value === 1) {
    return null
  } else {
    await changepage(1)
  }
}

const limitpage = async () => {
  if (numberPage.value === 1) {
    return null;
  } else {
    changepage(numberPage.value);
  }
}
</script>

<template>
  <div class="bg-primary py-4 md:py-8 flex items-center justify-between sm:px-6">
    <div class="flex-1 flex justify-center sm:hidden">
      <span style="border-color: #919191" class="
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
            " @click="startpage()">
        <span style="color: #919191" class="font-bold">&laquo;</span>
      </span>

      <span style="font-family: 'Inter, sans-serif'" :class="`
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
                ${page === 1
          ? 'bg-blue text-white border-none'
          : 'bg-blueDark text-white border-none'
        }
          `" @click="previewpage()">
        Previous
      </span>
      <span style="font-family: 'Inter, sans-serif'" :class="`
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
                      ${page === numberPage
          ? 'bg-blue text-white border-none'
          : 'bg-blueDark text-white border-none'
        }
                `" @click="nextpage()">
        Next
      </span>
      <span style="border-color: #919191" class="
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
            " @click="limitpage()">
        <span style="color: #919191" class="font-bold">&raquo;</span>
      </span>
    </div>
    <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
      <div>
        <p v-if="nbOfItems == 0" class="text-xs md:text-sm text-gray-700">
          No items found
        </p>
        <p v-else class="text-xs md:text-sm text-gray-700">
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
        <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
          <span class="
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
                " @click="startpage()">
            <span class="font-bold">&laquo;</span>
          </span>
          <div @click="previewpage()">
            <span :class="`z-10
                                          border 
                                          border-gray-300
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
                            `">
              <span class="sr-only">Previous</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"
                aria-hidden="true">
                <path fill-rule="evenodd"
                  d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>
            </span>
          </div>
          <div v-for="currL in leftSide" :key="currL" @click="changepage(currL)">
            <span :class="`z-10
                                          border 
                                          border-gray-300
                              text-gray-500
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
                                        ${page === currL ? 'bg-blue' : 'bg-white'}
                                        `">
              {{ currL }}
            </span>
          </div>
          <div>
            <span class="
                    z-10
                              border 
                              border-gray-300
                  text-gray-500
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
                  ">
              {{ page }}
            </span>
          </div>
          <div v-for="currR in rightSide" :key="currR" @click="changepage(currR)">
            <span :class="`
                                          z-10
                                                      border 
                                                      border-gray-300
                                          text-gray-500
                                          relative
                                          inline-flex
                                          items-center
                                          px-4
                                          py-2
                                          text-xs
                                          md:text-sm
                                          font-medium
                                          cursor-pointer
                                          ${page === currR ? 'bg-blue' : 'bg-white'}
                                        `">
              {{ currR }}
            </span>
          </div>
          <div @click="nextpage()">
            <span :class="`
                                          z-10
                                                      border-gray-300
                                          text-gray-500
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
                                        `">
              <span class="sr-only">Next</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"
                aria-hidden="true">
                <path fill-rule="evenodd"
                  d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                  clip-rule="evenodd" />
              </svg>
            </span>
          </div>
          <span class="
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
                " @click="limitpage()">
            <span class="font-bold">&raquo;</span>
          </span>
        </nav>
      </div>
    </div>
  </div>
</template>