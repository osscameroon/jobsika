<template>
  <div class="site__title" v-if="onlyTitle === 'onlyTitle'">
    <div :style="myStyle" class="pt-2 md:pt-8 lg:pt-14">
      <div class="site__title-text text-center">
        <h2
          v-if="content === 'subtitle'"
          style="color: #000000; font-family: 'Inter', sans-serif"
          class="font-bold text-xl md:text-2xl flex justify-center md:justify-start"
        >
          {{ title }}
        </h2>

        <h2
          v-else
          style="color: #000000; font-family: 'Inter', sans-serif"
          class="font-bold text-3xl md:text-4xl flex justify-center"
        >
          {{ title }}
        </h2>
      </div>
    </div>
  </div>
  <div class="site__title" v-else>
    <div class="pt-2 md:pt-8 lg:pt-14">
      <div class="site__title-text">
        <div class="flex justify-center md:justify-start">
          <h2
            style="color: #000000; font-family: 'Inter', sans-serif"
            class="font-bold text-3xl md:text-4xl"
          >
            {{ title }}
          </h2>
          <div class="-mt-2">
            <span
              class="cursor-pointer px-1 ml-2 -mb-5 text-grayC rounded-full border border-grayC text-xs"
              :class="{ opened: opened.includes(tooltips[0].id) }"
              @click="toggle(tooltips[0].id)"
            >
              !
            </span>
          </div>
        </div>
        <div v-if="opened.includes(tooltips[0].id)" class="w-full bg-primary">
          <div class="bg-white w-full px-2 py-4 my-3 shadow-sm rounded-sm">
            <p
              class="text-xs md:text-sm"
              style="
                color: #000000;
                font-family: 'Inter', sans-serif;
                line-height: 1.5rem;
              "
            >
              Company name is only displayed if at least 3 people in the same
              company contributes. Until then, a generic description like "Annon
              company 001" will be displayed. The Job title you entered is only
              displayed if at least 3 people in the same company, with the same
              Job title, from the same company share compensation details. Until
              then, a generic description like "Tier-level role, software
              engineering" is displayed instead of e.g. "Senior Frontend
              Engineer". Manual review happens before the data is published.
              Information that might help pinpointing the person sharing the
              compensation might be removed or changed during this review.
            </p>
          </div>
        </div>
        <div class="flex justify-between lg:items-end lg:flex-row flex-col">
          <div
            v-if="fontSize === 'header'"
            class="flex justify-center md:justify-start"
          >
            <h4
              style="color: #235365; font-family: 'Inter', sans-serif"
              class="pt-3 font-bold text-2xl md:text-3xl text-center"
            >
              {{ subTitle }}
            </h4>
          </div>
          <h4
            v-else
            style="color: #235365; font-family: 'Inter', sans-serif"
            class="pt-3 font-bold text-base"
          >
            {{ subTitle }}
          </h4>
          <div class="site__title-btn pt-10 lg:pt-0 w-full sm:w-1/2 lg:w-36">
            <NuxtLink to="/add_salary">
              <Button
                myStyle="background: #235365; font-family: 'Inter', sans-serif"
                name="Contribute"
              />
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
