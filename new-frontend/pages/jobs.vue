<script setup lang="ts">
	import { useJobsStore } from "../store/jobs";

	const jobStore = useJobsStore();

	const showModal: Ref<boolean> = ref(false);
	const image: Ref<string> = ref('/job.png');

	const jobs = computed(() => jobStore.items);
	const page = computed(() => jobStore.page);
	const limit = computed(() => jobStore.limit);

	onMounted(async () => {
		await fetchJobs();
	});

	const fetchJobs = async () => {
		await jobStore.getItems({
			page: page.value,
			limit: limit.value
		});
	}
</script>

<template>
	<main>
		<div class="container mx-auto md:w-7/12 px-4 md:px-0">
			<div class="pt-10 lg:pt-14">
				<h2 class="font-bold text-2xl md:text-4xl text-center text-black">
					Find the job that fits your skills and passion
				</h2>
			</div>
			<div class="pt-8 md:pt-16">
				<div class="flex flex-col xl:flex-row xl:justify-end">
					<p
						class="font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-center text-gray-700">
						Receive the jobs right in your inbox as soon as they are posted
					</p>
					<div class="pt-4 md:pt-2 xl:pt-0 w-full xl:w-36 xl:ml-4">
						<button @click="showModal = true"
							class="mx-auto xl:mx-0 w-1/2 xl:w-36 cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-xs md:text-lg px-4 border-2 h-12 rounded-md border-blue">
							Subscribe
						</button>
					</div>
				</div>
			</div>
			<div class="pt-8">
				<div v-for="(item, index) in jobs" :key="index" class="mb-6">
					<JobPost :job="item"/>
				</div>
			</div>
			<JobPagination />
		</div>
		<SubmitModal v-show="showModal" @close-modal="showModal = false" />
	</main>
</template>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
