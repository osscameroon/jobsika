<template>
  <div class="flex items-center justify-center py-6">
    <div>
      <!-- <p class="bg-white">{{ currentPage }}</p> -->
      <v-pagination class="flex" v-model="currentPage" :page-count="30">
        <p>Hello</p>
      </v-pagination>
    </div>
  </div>
</template>

<script>
import vPagination from 'vue-plain-pagination'

export default {
  name: 'Pagination_Number',
  components: { vPagination },
  data() {
    return {
      currentPage: 1,
    }
  },
}
</script>
