<script setup lang="ts">

</script>

<template>
	<div>
		<main class="site__main pt-10 lg:pt-4">
			<div class="container mx-auto md:w-7/12 px-4 md:px-0">
				<Title title="100% Anonymous" subTitle="Company ratings and salaries in Cameroon" fontSize="header"
					onlyTitle="header" header="header" />
				<div class="site__main-select flex flex-col lg:flex-row justify-end pt-20">
					<div class="site__main-selectItem">
						<Select />
					</div>
				</div>
				<div class="site__main-company py-2 md:py-8">
					<Company />
					<Average title="Average" />
					<Pagination />
				</div>
			</div>
		</main>
		<footer class="site__footer pb-16">
			<div class="container mx-auto md:w-7/12 text-center px-4 md:px-0">
				<Title title="Let's change the narative"
					subTitle="When no one shares their salary, employers have the upper hand. We can change this."
					content="subtitle" />
			</div>
		</footer>
	</div>
</template>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>