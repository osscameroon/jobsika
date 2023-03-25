<template>
  <div class="site__title" v-if="onlyTitle === 'onlyTitle'">
    <div :style="myStyle" class="pt-2 md:pt-8 lg:pt-14">
      <div class="site__title-text text-center">
        <h2 v-if="content === 'subtitle'" style="color: #000000; font-family: 'Inter', sans-serif"
          class="font-bold text-xl md:text-2xl flex justify-center md:justify-start">
          {{ title }}
        </h2>

        <h2 v-else style="color: #000000; font-family: 'Inter', sans-serif"
          class="font-bold text-3xl md:text-4xl flex justify-center">
          {{ title }}
        </h2>
      </div>
    </div>
  </div>

  <div class="site__title" v-else>
    <div v-if="header === 'header'" class="pt-2 md:pt-8 lg:pt-14">
      <div class="site__title-text">
        <div class="flex justify-center md:justify-start">
          <h2 style="color: #000000; font-family: 'Inter', sans-serif" class="font-bold text-3xl md:text-4xl">
            {{ title }}
          </h2>
          <div class="flex items-center -mb-1">
            <span
              class="cursor-pointer h-5 text-center w-5 ml-2 text-grayDark rounded-full border border-grayDark text-xs"
              :class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[0].id)">
              !
            </span>
          </div>
        </div>
        <div v-if="opened.includes(tooltips[0].id)" class="w-full bg-primary">
          <div class="w-full my-3 ml-3 md:ml-0">
            <p class="text-xs md:text-base font-bold" style="
                color: #235365;
                font-family: 'Inter', sans-serif;
                line-height: 1.5rem;
              ">
              ...But here is what you need to know:
            </p>
            <div class="ml-0 md:ml-8 my-1">
              <u class="no-underline">
                <li>
                  <span class="text-xs md:text-sm mt-1 leading-5" style="color: #808081">
                    <span class="font-bold">The company name</span> is only
                    displayed if at least 3 people in the same company with the
                    same job title contribute. Until then, a generic description
                    like "A local company" will be displayed.
                  </span>
                </li>
                <li>
                  <span class="text-xs md:text-sm mt-2 leading-5" style="color: #808081">
                    <span class="font-bold">The comment</span> is only displayed
                    if at least 3 people in the same company with the same job
                    title contribute.
                  </span>
                </li>
                <li>
                  <span class="text-xs md:text-sm mt-2 leading-5" style="color: #808081">
                    <span class="font-bold">The rating</span> is only displayed
                    if at least 3 people in the same company with the same job
                    title contribute.
                  </span>
                </li>
              </u>
            </div>
          </div>
        </div>
        <div class="flex justify-between lg:items-end lg:flex-row flex-col">
          <div v-if="fontSize === 'header'" class="flex justify-center md:justify-start">
            <h4 style="color: #235365; font-family: 'Inter', sans-serif"
              class="pt-3 font-bold text-2xl md:text-3xl text-center">
              {{ subTitle }}
            </h4>
          </div>
          <h4 v-else style="color: #235365; font-family: 'Inter', sans-serif" class="pt-3 font-bold text-base">
            {{ subTitle }}
          </h4>
          <div class="site__title-btn pt-10 lg:pt-0 w-full sm:w-1/2 lg:w-36">
            <NuxtLink to="/add_salary">
              <Button myStyle="background: #235365; font-family: 'Inter', sans-serif" name="Contribute" />
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
    <div v-else class="pt-2 md:pt-8 lg:pt-14">
      <div class="site__title-text">
        <div class="flex justify-center md:justify-start">
          <h2 style="color: #000000; font-family: 'Inter', sans-serif" class="font-bold text-3xl md:text-4xl">
            {{ title }}
          </h2>
        </div>
        <div class="flex justify-between lg:items-end lg:flex-row flex-col">
          <div v-if="fontSize === 'header'" class="flex justify-center md:justify-start">
            <h4 style="color: #235365; font-family: 'Inter', sans-serif"
              class="pt-3 font-bold text-2xl md:text-3xl text-center">
              {{ subTitle }}
            </h4>
          </div>
          <h4 v-else style="color: #235365; font-family: 'Inter', sans-serif" class="pt-3 font-bold text-base">
            {{ subTitle }}
          </h4>
          <div class="site__title-btn pt-10 lg:pt-0 w-full sm:w-1/2 lg:w-36">
            <NuxtLink to="/add_salary">
              <Button myStyle="background: #235365;" name="Contribute" />
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Button from '../components/Button.vue'
export default Vue.extend({
  name: 'TitleComponent',
  props: {
    title: String,
    subTitle: String,
    fontSize: String,
    onlyTitle: String,
    content: String,
    myStyle: String,
    header: String,
  },
  components: {
    Button,
  },
  data() {
    return {
      isOpen: false,
      opened: [],
      tooltips: [
        {
          id: 0,
          name: 'title',
        },
      ],
    }
  },
  methods: {
    toggle(id) {
      const index = this.opened.indexOf(id)
      if (index > -1) {
        this.opened.splice(index, 1)
      } else {
        this.opened.push(id)
      }
    },
  },
})
</script>
