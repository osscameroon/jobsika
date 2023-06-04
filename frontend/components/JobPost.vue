<template>
    <div class="bg-white shadow-md p-3 md:p-4 rounded-md hover:shadow-xl" @mouseover="active = true"   @mouseleave="active = false">
        <div class="flex flex-col xl:flex-row">
            <div class="w-28 h-28 border rounded-md flex justify-center items-center">
                <img class="w-auto h-24" :src="picture" />
            </div>
            <div class="xl:w-7/12 pt-4 xl:pt-0 xl:ml-8 flex flex-col">
                <h4 class="text-lg font-bold">{{ title }} <span class="text-gray-600 font-light"> {{ city }}</span></h4>
                <div class="flex flex-col md:flex-row py-2">
                    <div class="flex">
                        <img :src="location" class="w-4 h-4 mr-1" />
                        <p class="text-xs font-extralight text-gray-600">{{ structure }}</p>
                    </div>
                    <div class="xl:ml-4 pt-2 xl:pt-0 flex">
                        <img :src="position" class="w-4 h-4 mr-1" />
                        <p class="text-xs font-extralight text-gray-600">{{ marker }}</p>
                    </div>
                    <div class="xl:ml-4 pt-2 xl:pt-0 flex">
                        <img :src="clock" class="w-4 h-4 mr-1" />
                        <p class="text-xs font-extralight text-gray-600">{{new Date(time).toLocaleDateString() }}</p>
                    </div>
                </div>
                <p class="text-xs font-extralight text-gray-600">{{ description }}</p>
                <div class="pt-1 flex flex-wrap">
                    <p v-for="(item, index) in tags.split(',')" :key="index" class="text-xs m-2 text-gray-500">
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

<script>
export default {
    name: 'JobComponent',
    props: {
        title: {
            type: String,
            default: ""
        },
        picture: {
            type: String,
            default: ""
        },
        city: {
            type: String,
            default: ""
        },
        structure: {
            type: String,
            default: ""
        },
        marker: {
            type: String,
            default: ""
        },
        time: {
            type: String,
            default: ""
        },
        description: {
            type: String,
            default: ""
        },
        tags: {
            type: [String],
            default: null
        },
    },
    data() {
        return {
            location: require('../assets/location.png'),
            position: require('../assets/position.png'),
            clock: require('../assets/clock.png'),
            active: false
        }
    },
    methods: {
        mouseOver() {
            this.active = !this.active;
        }
    }

}
</script>
