<script setup lang="ts">
	const { name } = defineProps({
		name: {type: String, required: true, default: ""}
	});
</script>

<template>
  <div class="h-24 flex items-center justify-center" style="background: #d9d9d9;">
    <p class="text-base font-bold">
      {{ name }}
    </p>
	</div>
</template>
