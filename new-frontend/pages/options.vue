<script setup lang="ts">
	// const opened: Ref<any[]> = ref([]);
	// const uploadImg: Ref<string> = ref("../assets/upload.png");
</script>

<template>
	<main>
		<div class="container mx-auto md:w-7/12 px-4 md:px-0">
			<div class="pt-10 lg:pt-14">
				<h2 class="font-bold text-2xl md:text-4xl text-center text-black">
					Add a job
				</h2>
			</div>
			<div class="pt-8 md:pt-16"></div>
			<div>
				<div class="flex flex-col md:flex-row md:justify-center">
					<div
						class="px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray50 h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4">
						Job details
					</div>
					<div
						class="px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4">
						Options
					</div>
					<div
						class="px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray50 h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4">
						Payments
					</div>
				</div>
				<div class="flex flex-col md:w-10/12 mx-auto mt-7">
					<p class="font-bold text-base leading-10 mx-auto">
						Start hiring at FCFA 15 000 only
					</p>
					<div class="mt-12">
						<div class="flex items-center">
							<input id="default-checkbox" type="checkbox" value=""
								class="w-6 h-6 text-primaryGray  rounded focus:bg-primaryGray">
							<label for="default-checkbox" class="ml-4 text-xl font-normal text-gray-900">Company logo
								(+2.000XAF)</label>
						</div>
					</div>
					<div class="mt-4">
						<div class="flex items-center">
							<input id="default-checkbox" type="checkbox" value=""
								class="w-6 h-6 text-primaryGray  rounded focus:bg-primaryGray">
							<label for="default-checkbox" class="ml-4 text-xl font-normal text-gray-900">Diffusion on OSS
								Groups and accounts ( (8 x users in our databases) XAF)</label>
						</div>
					</div>
					<div class="mt-4">
						<div class="flex items-center">
							<input id="default-checkbox" type="checkbox" value=""
								class="w-6 h-6 text-primaryGray  rounded focus:bg-primaryGray">
							<label for="default-checkbox" class="ml-4 text-xl font-normal text-gray-900">Personal
								notifications(Email & Telegram (JobSika Bot)) ( (15 x users in our databases) XAF)</label>
						</div>
					</div>
					<div class="mt-4">
						<div class="flex items-center">
							<input id="default-checkbox" type="checkbox" value=""
								class="w-6 h-6 text-primaryGray  rounded focus:bg-primaryGray">
							<label for="default-checkbox" class="ml-4 text-xl font-normal text-gray-900">Pin the Job offer
								on OSS Cameroon groups/account (+ 5.000XAF/week)</label>
						</div>
					</div>
					<div class="mt-4">
						<div class="flex items-center">
							<input id="default-checkbox" type="checkbox" value=""
								class="w-6 h-6 text-primaryGray  rounded focus:bg-primaryGray">
							<label for="default-checkbox" class="ml-4 text-xl font-normal text-gray-900">Pin the Job offer
								on Jobsika homepage (+10.000 XAF/week)</label>
						</div>
					</div>
					<div class="mt-16 flex flex-col md:flex-row items-center justify-center xl:justify-start">
						<div class="w-full md:w-1/5">
							<div
								class="p-4 text-blueDark border-2 border-blueDark font-bold flex items-center justify-center text-xs md:text-base h-12"
								style="border-radius: 6px">
								Total: 17 500
							</div>
						</div>
						<NuxtLink class="w-full" to="/payments">
							<button class="w-full md:w-1/5 ml-0 pt-6 md:pt-0 md:ml-7" type="submit">
								<Button show-picture="nothing"
									my-style="background: #235365;padding: 20px 80px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);" name="Next" />
							</button>
						</NuxtLink>
					</div>
				</div>
			</div>
		</div>
	</main>
</template>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
