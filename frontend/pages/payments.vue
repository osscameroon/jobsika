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
                    <h4 class="font-bold text-5xl leading-10 mx-auto text-center">
                        How do you want to make your payment?
                    </h4>
                    <p style="color: #7694A0;" class="mt-6 text-center text-base mx-10">
                        Last step! please choose a payment method. Your job offer will be published as soon as your payment
                        is received and approved by admins
                    </p>
                    <div class="mt-12">
                        <NuxtLink to="#">
                            <Payment name="Open collective" />
                        </NuxtLink>
                        <div class="mt-10 md:mt-20" />
                        <NuxtLink to="#">
                            <Payment name="Mobile money/ Visa/Mastercard" />
                        </NuxtLink>
                    </div>
                    <NuxtLink class="w-full mt-16 flex items-center justify-center" to="/payment-status">
                        <button class="w-full md:w-1/4 pt-6" type="submit">
                            <Button show-picture="nothing"
                                my-style="background: #235365;padding: 20px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);"
                                name="Make payment" />
                        </button>
                    </NuxtLink>
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
        }
    },
    methods: {
    },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
