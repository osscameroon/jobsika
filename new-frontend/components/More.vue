<script setup lang="ts">
  const { path, title } = defineProps({
		path: { type: String, required: true},
		title: {type: String, required: true}
	});
  const arrow: string = '/arrowright.png';
</script>

<template>
  <div class="bg-white shadow-md h-14 py-2 flex justify-center items-center border-t-2">
    <NuxtLink :to="path" class="font-normal text-lg ext-gray-700 flex items-center">
      {{ title }} <span class="ml-4"><img class="w-auto h-4" :src="arrow" /></span>
    </NuxtLink>
  </div>
</template>
