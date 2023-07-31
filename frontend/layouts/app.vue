<template>
  <div>
    <header class="bg-white flex items-center justify-start py-3 lg:py-6 z-50 sticky top-0 left-0">
      <div class="container mx-auto w-11/12 md:w-10/12 px-4 lg:px-0">
        <div class="flex flex-row justify-between items-center">
          <a href="/" class="cursor-pointer">
            <img :src="logo" alt="logo" class="w-32 h-12 lg:w-32 lg:h-auto" />
          </a>
          <div class="flex md:hidden">
            <img @click="openNav" :src="menu" alt="menu" :class="`w-6 h-auto ${navbar === false ? 'flex': 'hidden'}`" />
            <img @click="openNav" :src="close" alt="menu" :class="`w-8 h-auto ${navbar === true ? 'flex': 'hidden'}`" />
          </div>
          <nav class="hidden md:flex">
            <ul class="flex flex-col md:flex-row items-center justify-center space-x-4 xl:space-x-8">
              <NuxtLink v-for="(item, index) in navs" :key="index" :to="`${item.link}`" :class="`flex justify-center items-center text-sm lg:text-base font-bold ${setID === item.id ? 'nav-select' : 'nav-noselect'
                }`" @click.native="setID = item.id">
                {{ item.name }}
              </NuxtLink>
            </ul>
          </nav>
          <div class="hidden lg:flex flex-col space-y-4 lg:space-y-0 lg:space-x-4 lg:flex-row">
            <div class="w-64 lg:w-36">
              <NuxtLink to="/add_job"
                class="cursor-pointer p-4 text-white font-bold flex items-center justify-center text-sm lg:text-base bg-blueDark h-12 rounded-lg">
                Post a job
              </NuxtLink>
            </div>
            <div class="w-full lg:w-36">
              <NuxtLink to="/add_salary"
                class="cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-sm lg:text-base border-2 h-12 rounded-lg border-blue">
                <img :src="plusIcon" class="w-5 h-5 mr-1" />
                Contribute
              </NuxtLink>
            </div>
          </div>
        </div>
        <div :class="`${navbar === true ? 'flex' : 'hidden'}`">
          <div class="pt-10 w-full">
            <ul class="flex flex-col items-center justify-center space-y-6">
              <NuxtLink v-for="(item, index) in navs" :key="index" :to="`${item.link}`" :class="`flex justify-center items-center text-sm lg:text-base font-bold ${setID === item.id ? 'nav-select' : 'nav-noselect'
                }`" @click.native="setID = item.id">
                {{ item.name }}
              </NuxtLink>
              <div class="flex flex-col space-y-4 lg:space-y-0 lg:space-x-4">
                <div class="w-64 lg:w-36">
                  <NuxtLink to="/add_job"
                    class="cursor-pointer p-4 text-white font-bold flex items-center justify-center text-sm lg:text-base bg-blueDark h-12 rounded-lg">
                    Post a job
                  </NuxtLink>
                </div>
                <div class="w-full lg:w-36">
                  <NuxtLink to="/add_salary"
                    class="cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-sm lg:text-base border-2 h-12 rounded-lg border-blue">
                    <img :src="plusIcon" class="w-5 h-5 mr-1" />
                    Contribute
                  </NuxtLink>
                </div>
              </div>
            </ul>
          </div>
        </div>
      </div>
    </header>
    <Nuxt />
    <footer class="pb-16">
      <div class="container mx-auto w-10/12 text-center">
        <Footer />
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      menu: require('../assets/menu.png'),
      close: require('../assets/close.png'),
      logo: require('../assets/logo.png'),
      plusIcon: require('../assets/plus_dark.png'),
      navs: [
        {
          id: 'home',
          link: '/',
          name: 'HOME',
        },
        {
          id: '#jobs',
          link: '/jobs',
          name: 'JOBS',
        },
        {
          id: '#salaries',
          link: '/salaries',
          name: 'SALARIES',
        },
      ],
      setID: 'home',
      navbar: false,
      closeIcon: false
    }
  },
  methods: {
    openNav() {
      this.navbar = !this.navbar
    }
  }
}
</script>
