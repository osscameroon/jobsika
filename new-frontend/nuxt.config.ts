// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	modules: [
		'@pinia/nuxt'
	],
	postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
	css: [
		'@/assets/css/main.css',
	],
	build: {
		
	},
	app: {
		head: {
			title: 'Jobsika',
			htmlAttrs: {
				lang: 'en',
			},
			meta: [
				{ charset: 'utf-8' },
				{ name: 'viewport', content: 'width=device-width, initial-scale=1' },
				{ hid: 'description', name: 'description', content: '' },
				{ name: 'format-detection', content: 'telephone=no' },
			],
			link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
		},
	},
	runtimeConfig: {
		public:{
			baseURL: process.env.BASE_URL
		}
	},
	vite: {
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: '@use "@/assets/_colors.scss" as *;'
        }
      }
    }
  }
})
