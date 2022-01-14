<template>
  <div class="site" style="background: #e5e5e5">
    <Header name="Search..." />
    <main class="site__main pt-20">
      <div class="container mx-auto w-9/12">
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
        <div class="site__main-company pt-2">
          <Company v-s />
        </div>
      </div>
    </main>
    <footer class="site__footer py-16">
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
export default {
  name: 'IndexPage',
  components: {
    Header,
    Title,
    Select,
    Company,
  },
  data() {
    return {
      selectOne: ['Un', 'Deux', 'Trois'],
    }
  },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
