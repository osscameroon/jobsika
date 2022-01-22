<template>
  <div class="site" style="background: #f1f5f9">
    <Header name="Search..." />
    <main class="site__main pt-10 lg:pt-4">
      <div class="container mx-auto w-10/12">
        <Title
          title="100% Anonymous"
          subTitle="Company rating and salaries in Cameroon"
          fontSize="header"
        />
        <div
          class="site__main-select flex flex-col lg:flex-row justify-end pt-20"
        >
          <div class="site__main-selectItem">
            <Select />
          </div>
        </div>
        <div class="site__main-company py-2 md:py-8">
          <Company />
          <Pagination />
        </div>
      </div>
    </main>
    <footer class="site__footer pb-16">
      <div class="container mx-auto w-9/12">
        <Title
          title="Let's change the narative"
          subTitle="When no one shares their salary, employers have the upper hand. We can change this."
        />
      </div>
    </footer>
  </div>
</template>

<script>
import Header from '../components/Header.vue'
import Title from '../components/Title.vue'
import Select from '../components/Select.vue'
import Company from '../components/Company.vue'
import Pagination from '../components/Pagination.vue'
export default {
  name: 'IndexPage',
  components: {
    Header,
    Title,
    Select,
    Company,
    Pagination,
  },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
