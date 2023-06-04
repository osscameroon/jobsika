<template>
    <div :class="{ opened: opened.includes(details[myIndex].id) }" @click="toggle(details[myIndex].id)">
        <div class="bg-white cursor-pointer shadow-md p-3 md:p-4 rounded-md hover:shadow-xl" @mouseover="active = true"
            @mouseleave="active = false">
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
                            <p class="text-xs font-extralight text-gray-600">{{ new Date(time).toLocaleDateString() }}</p>
                        </div>
                    </div>
                    <div class="pt-1 flex flex-wrap">
                        <p v-for="(item, index) in tags.split(',')" :key="index" class="text-xs m-2 text-gray-500">
                            <span class="bg-gray-200 rounded-full px-4 py-1">{{ item }}</span>
                        </p>
                    </div>
                    <div v-if="opened.includes(details[myIndex].id)" class="w-full">
                        <div class="my-4">
                            <h4 class="text-gray-600 font-light">{{ city }} is hiring <span class=" text-lg font-bold"> {{
                                title }}</span></h4>

                            <p class="text-sm font-extralight pt-4 text-gray-600">What is {{ structure }}?</p>
                            <p class="text-sm font-extralight pt-4 text-gray-600">{{ description }}</p>
                            <p class="text-sm font-extralight pt-4 text-gray-600"><span class="font-bold">Contact: </span>{{
                                mail }}</p>
                            <p class="text-sm font-extralight pt-4 text-gray-600"><span class="font-bold">Salary range:
                                </span>Min({{ minSalary }} FCFA), Max({{ maxSalary }} FCFA)</p>
                        </div>
                        <div class="my-8">
                            <NuxtLink to="#"
                                class="cursor-pointer w-full mx-auto px-10 py-3 text-white font-bold flex items-center justify-center text-xs lg:text-sm bg-primaryGray h-10 rounded-lg">
                                <p>
                                    Click here to apply
                                </p>
                            </NuxtLink>
                        </div>
                    </div>
                </div>
                <div class="xl:justify-end items-center w-full xl:w-3/12 pt-4 xl:pt-0 flex">
                    <NuxtLink to="#"
                        :class="`cursor-pointer px-12 py-4 text-white font-bold items-center justify-center text-sm lg:text-base bg-primaryGray h-12 rounded-lg ${active == true ? 'flex' : 'hidden'}`">
                        Apply
                    </NuxtLink>
                </div>
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
        minSalary: {
            type: String,
            default: ""
        },
        maxSalary: {
            type: String,
            default: ""
        },
        description: {
            type: String,
            default: ""
        },
        mail: {
            type: String,
            default: ""
        },
        tags: {
            type: [String],
            default: null
        },
        myIndex: {
            type: Number,
            default: 0
        },
        details: {
            type: [],
            default: null
        }
    },
    data() {
        return {
            location: require('../assets/location.png'),
            position: require('../assets/position.png'),
            clock: require('../assets/clock.png'),
            active: false,
            opened: [],
        }
    },
    methods: {
        mouseOver() {
            this.active = !this.active;
        },
        toggle(id) {
            const index = this.opened.indexOf(id)
            if (index > -1) {
                this.opened.splice(index, 1)
            } else {
                this.opened.push(id)
            }
        },
    }

}
</script>
