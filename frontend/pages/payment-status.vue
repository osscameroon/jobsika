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
                        class="px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray50 h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4">
                        Options
                    </div>
                    <div
                        class="px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4">
                        Payments
                    </div>
                </div>
                <div class="flex flex-col md:w-11/12 mx-auto mt-14">
                    <div class="flex items-center justify-center my-6">
                        <img :src="checkImg" alt="" class="h-96 w-1/3 object-contain" />
                    </div>
                    <p style="color: #235365;" class="mt-6 text-center text-xl md:w-1/2 mx-auto font-bold">
                        Pending approuval. Your job will be published as soon as your payment is approuved
                    </p>
                </div>
            </div>
        </div>
    </main>
</template>

<script>
export default {
    name: 'OptionsPage',
    layout: 'app',
    data() {
        return {
            opened: [],
            uploadImg: require("../assets/upload.png"),
            checkImg: require("../assets/payment.png"),
        }
    },
    methods: {
    },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
