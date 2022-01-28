<template>
  <div>
    <main class="site__main pt-5 lg:pt-0 mb-10 md:mb-20">
      <div class="container mx-auto w-10/12">
        <Title title="100% Anonymous" fontSize="header" onlyTitle="onlyTitle" />
        <Title
          title="Add a company to the list"
          onlyTitle="onlyTitle"
          content="subtitle"
          myStyle="margin-top: 20px; font-family: 'Inter',sans-serif"
        />
        <div class="site__main-row flex flex-col-reverse md:flex-row">
          <div class="site__main-rowOne w-full md:w-1/2">
            <div class="site__main-input mt-3 md:mt-10 md:w-11/12">
              <MyInput
                title="Company Name"
                myStyle="height: 61px;"
                content="subtitle"
                myInput="input"
              />
              <MyInput
                title="Salary"
                myStyle="height: 61px;"
                content="subtitle"
                myInput="input"
              />
              <MyInput
                title="Comments"
                myStyle="height: 120px;"
                content="subtitle"
              />
              <div class="site__input-btn mt-10 flex flex-col md:flex-row">
                <div class="w-full md:w-1/4">
                  <Button
                    showPicture="nothing"
                    name="Cancel"
                    myStyle="background: white; color: #000000"
                  />
                </div>
                <div class="w-full md:w-1/4 ml-0 pt-6 md:pt-0 md:ml-6">
                  <Button
                    myStyle="background: #235365; font-family: 'Inter', sans-serif"
                    name="Add"
                  />
                </div>
              </div>
            </div>
          </div>
          <div class="site__main-rowTwo w-full md:w-1/2">
            <div class="site__main-input mt-3 md:mt-10 md:w-11/12">
              <MyInput
                title="Job title"
                myStyle="height: 61px;"
                content="subtitle"
                myInput="input"
              />
              <div class="site__input w-full flex my-3 md:my-0 md:mt-12">
                <p
                  class="text-xs md:text-sm font-bold"
                  style="color: #b1b1b1; font-family: 'Inter', sans-serif"
                >
                  Rate
                </p>
                <div class="flex mt-0 ml-6">
                  <div v-for="item in 5" :key="item" class="flex">
                    <img class="w-4 h-4 mr-1" :src="starsPicture" />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Button from '../components/Button.vue'
export default Vue.extend({
  name: 'AddSalary',
  components: {
    Button,
  },
  data() {
    return {
      starsPicture: require('../assets/star.png'),
    }
  },
})
</script>
