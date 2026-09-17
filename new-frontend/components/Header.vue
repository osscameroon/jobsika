<script setup lang="ts">
  const { name } = defineProps({
		name: {type: String, required: true},
	});

  const searchPicture: Ref<string> = ref('/search.png');
</script>

<template>
  <header
    style="box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.14); height: 67px"
    class="site__header left-0 top-0 bg-white w-full sticky z-50"
  >
    <div class="container mx-auto w-11/12 h-full">
      <div class="flex items-center py-4">
        <img class="w-6 h-6" :src="searchPicture" />
        <input
          type="text"
          name="name"
          :placeholder="name"
          class="
            w-full
            border-none
            text-gray-600
            font-medium
            text-sm
            md:text-base
          "
          style="font-family: 'Inter', sans-serif"
        />
      </div>
    </div>
  </header>
</template>