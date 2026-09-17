<script setup lang="ts">
	const { title } = defineProps({
		title: {type: String, required: true, default: ""}
	});

	const isActive: Ref<boolean> = ref(true); 
</script>

<template>
  <div v-show="isActive">
    <slot></slot>
  </div>
</template>