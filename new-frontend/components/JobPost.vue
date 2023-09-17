<script setup lang="ts">
  import { IJob } from '~/types/job';

  const DEFAULT_JOB: IJob = {
		id: 0, createdat: "", updatedat: "", company_name: "", job_title: "",
		is_remote: false, city: "", country: "", department: "", salary_range_min: 0,
		salary_range_max: 0, description: "", benefits: "", how_to_apply: "", tags: "",
		application_url: "",	application_email_address: "",	application_phone_number: ""
  }

	interface JobProps {
    job: IJob
	}

	const { job } = defineProps<JobProps>();

	const location: Ref<string> = ref('/location.png');
  const position: Ref<string> = ref('/position.png');
  const clock: Ref<string> = ref('/clock.png');
	const picture: Ref<string> = ref('/logo.png');
  const active: Ref<boolean> = ref(false);

	const mouseOver = () => {
    active.value = !active.value;
  }
</script>

<template>
    <div class="bg-white shadow-md p-3 md:p-4 rounded-md hover:shadow-xl" @mouseover="active = true"   @mouseleave="active = false">
        <div class="flex flex-col xl:flex-row">
            <div class="w-28 h-28 border rounded-md flex justify-center items-center">
                <img class="w-auto h-24" :src="picture" />
            </div>
            <div class="xl:w-7/12 pt-4 xl:pt-0 xl:ml-8 flex flex-col">
                <h4 class="text-lg font-bold">{{ job.job_title }} <span class="text-gray-600 font-light"> {{ job.city }}</span></h4>
                <div class="flex flex-col md:flex-row py-2">
                    <div class="flex">
                        <img :src="location" class="w-4 h-4 mr-1" />
                        <p class="text-xs font-extralight text-gray-600">{{ job.company_name }}</p>
                    </div>
                    <div class="xl:ml-4 pt-2 xl:pt-0 flex">
                        <img :src="position" class="w-4 h-4 mr-1" />
                        <p class="text-xs font-extralight text-gray-600">{{ job.tags }}</p>
                    </div>
                    <div class="xl:ml-4 pt-2 xl:pt-0 flex">
                        <img :src="clock" class="w-4 h-4 mr-1" />
                        <p class="text-xs font-extralight text-gray-600">{{ job.createdat }}</p>
                    </div>
                </div>
                <p class="text-xs font-extralight text-gray-600">{{ job.description }}</p>
                <div class="pt-1 flex flex-wrap">
                    <p v-for="(item, index) in job.tags?.split(',')" :key="index" class="text-xs m-2 text-gray-500">
                        <span class="bg-gray-200 rounded-full px-4 py-1">{{ item }}</span>
                    </p>
                </div>
            </div>
            <div class="xl:justify-end items-center w-full xl:w-3/12 pt-4 xl:pt-0 flex">
                <NuxtLink to="#"
                    :class="`cursor-pointer px-12 py-4 text-white font-bold items-center justify-center text-sm lg:text-base bg-blueDark h-12 rounded-lg ${active == true ? 'flex' : 'hidden'}`">
                    Apply
                </NuxtLink>
            </div>
        </div>
    </div>
</template>
