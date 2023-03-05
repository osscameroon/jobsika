<template>
    <div class="bg-white shadow-md h-14 py-2 flex justify-center items-center border-t-2">
        <NuxtLink to="#" class="font-normal text-lg ext-gray-700 flex items-center">
            See more jobs <span class="ml-4"><img class="w-auto h-4" :src="arrow" /></span>
        </NuxtLink>
    </div>
</template>

<script>
export default {
    name: "MoreComponent",
    data() {
        return {
            arrow: require('../assets/arrowright.png'),
        }
    }
}
</script>
