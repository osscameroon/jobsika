<template>
    <main>
        <div class="container mx-auto md:w-7/12 px-4 md:px-0">
            <div>
                <div class="flex flex-col md:w-11/12 mx-auto mt-14">
                    <h4 class="font-bold text-5xl leading-10 mx-auto text-center">
                        Payment instructions
                    </h4>
                    <p style="color: #5b5b5b;" class="mt-6 text-center text-base mx-10 font-light">
                        Before you proceed with your payment, here is what you need to know!
                    </p>
                    <div class="mt-12 flex justify-center items-center">
                        <div class="bg-primaryGray p-2 rounded-md md:w-3/5">
                            <p style="font-size: 10px;" class="bg-gray-100 w-40 text-center py-1 rounded-md">
                                Powered by Open Collective
                            </p>
                            <div class="flex flex-col items-center justify-center text-white text-2xl md:text-5xl mt-14">
                                <h1 class="font-extrabold">10 000<span class="text-lg">FCFA</span></h1>
                                <div class="mt-10 text-xs leading-8">
                                    <p v-for="(item, index) in instructions" :key="index">&#8226; {{ item }}</p>
                                </div>
                                <div class="md:w-1/4 py-6">
                                    <NuxtLink to="/payment">
                                        <Button show-picture="nothing"
                                            my-style="background: #235365;padding: 20px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2); width: 164px;"
                                            name="Proceed payment" />
                                    </NuxtLink>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="mt-10 text-xs leading-8 mx-auto flex flex-col items-center">
                        <p v-for="(item, index) in instructions1" :key="index" style="color: #5b5b5b;"
                            class="text-left text-base font-light md:w-4/6 mb-4">&#8226; {{ item }}</p>
                    </div>
                </div>
            </div>

        </div>
    </main>
</template>

<script>
export default {
    name: 'JobsIndex',
    layout: 'app',
    data() {
        return {
            opened: [],
            uploadImg: require("../assets/upload.png"),
            checkImg: require("../assets/payment.png"),
            instructions: [
                "Company logo (+2.000XAF)",
                "Diffusion on OSS Groups and accounts ( (8 x users in our databases) XAF)"
            ],
            instructions1: [
                "After you choose an option, you will will be redirected to the open collective website where you can finalize payment and enter your payment infos",
                "You will receive an email of your job offer"
            ]
        }
    },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
