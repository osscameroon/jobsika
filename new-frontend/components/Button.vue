<script setup lang="ts">
  const { name, myStyle, showPicture } = defineProps({
		name: {type: String, required: true},
    myStyle: {type: String, required: true},
    showPicture: {type: String, default: ""}
	});

  const plusIcon: Ref<string> = ref('/plus.png');
</script>

<template>
  <div
    :style="myStyle"
    class="site__btn cursor-pointer p-4 text-white font-bold flex items-center justify-center text-xs md:text-sm h-12"
    style="border-radius: 6px"
  >
    <div v-if="showPicture === 'nothing'" />
    <img v-else :src="plusIcon" class="w-5 h-5 mr-1" />
    {{ name }}
  </div>
</template>