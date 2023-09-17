<script setup lang="ts">
	import { INav } from "../types/common";

	const logo: Ref<string> = ref('../assets/logo.png');
	const plusIcon: Ref<string> = ref('../assets/plus_dark.png');
	const setID: Ref<string> = ref('home');
	const navs: Ref<INav[]> = ref([
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
	]);
</script>

<template>
	<div>
		<header class="bg-white flex justify-center items-center lg:justify-start py-3 lg:py-2 z-50 sticky top-0 left-0">
			<div class="lg:container lg:mx-auto lg:w-10/12">
				<div class="flex flex-col lg:flex-row lg:justify-between items-center">
					<a href="/" class="cursor-pointer">
						<img :src="logo" alt="logo" class="site__logo-pic w-32 h-12 lg:w-32 lg:h-auto" />
					</a>
					<nav>
						<ul class="flex flex-col lg:flex-row items-center py-4 lg:py-0 justify-center">
							<NuxtLink v-for="(item, index) in navs" :key="index" :to="`${item.link}`" @click.native="setID = item.id"
								:class="`h-7 lg:w-20 lg:h-16 py-2 flex justify-center items-center text-sm lg:text-base font-bold ${setID === item.id ? 'nav-select' : 'nav-noselect'
									}`">
								{{ item.name }}
							</NuxtLink>
						</ul>
					</nav>
					<div class="flex flex-col lg:flex-row">
						<div class="lg:pt-0 w-full lg:w-36 lg:mr-4">
							<NuxtLink to="/add_job"
								class="cursor-pointer p-4 text-white font-bold flex items-center justify-center text-sm lg:text-base bg-blueDark h-12 rounded-lg">
								Post a job
							</NuxtLink>
						</div>
						<div class="pt-4 lg:pt-0 w-full lg:w-36 lg:ml-4">
							<NuxtLink to="/add_salary"
								class="cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-sm lg:text-base border-2 h-12 rounded-lg border-blue">
								<img :src="plusIcon" class="w-5 h-5 mr-1" />
								Contribute
							</NuxtLink>
						</div>
					</div>
				</div>
			</div>
		</header>
		<Nuxt />
		<footer class="site__footer pb-16">
			<div class="container mx-auto w-10/12 text-center">
				<Footer />
			</div>
		</footer>
	</div>
</template>  