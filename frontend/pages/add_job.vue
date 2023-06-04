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
                        :class="isActiveJob === false ? 'px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray50 h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4' : 'px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4'">
                        Job details
                    </div>
                    <!-- <div :class="isActiveOption === false ? 'px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray50 h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4 cursor-pointer' : 'px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4'"
                        @click='selectTabOption()'>
                        Options
                    </div> -->
                    <div
                        :class="isActivePayment === false ? 'px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray50 h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4' : 'px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4'">
                        Payments
                    </div>
                </div>
                <form v-if="isActiveJob">
                    <div class="flex flex-col md:w-7/12 mx-auto">
                        <div class="grid gap-7 grid-cols-1 mt-20">
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Company Name
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[0].id) }"
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[0].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-6" style="font-size: 10px;">Enter your company's brand/trade
                                            name:
                                            without Inc.,
                                            Ltd., B.V., Pte., etc.</p>
                                    </div>
                                </div>
                                <input type="text" style="height: 61px"
                                    class="border border-grayC mt-2 w-full rounded-md mb-4" />
                            </div>
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Job title
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[1].id) }"
                                            @click="toggle(tooltips[1].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[1].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-6" style="font-size: 10px;">This field requires you to enter
                                            the
                                            title of the
                                            position and/or specialization (if applicable). E.g. "Accountant(Audit)"</p>
                                    </div>
                                </div>
                                <input type="text" style="height: 61px"
                                    class="border border-grayC mt-2 w-full rounded-md mb-4" />
                            </div>
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Description of role
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[2].id) }"
                                            @click="toggle(tooltips[2].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[2].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-6" style="font-size: 10px;">Please add a brief description
                                            of
                                            the
                                            role you are
                                            hiring for, make it straight forward and concise</p>
                                    </div>
                                </div>
                                <textarea cols="6" rows="5" type="text"
                                    class="border border-grayC mt-2 w-full rounded-md mb-4">
                                    </textarea>
                            </div>
                            <div>
                                <div class="flex mb-2">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Employment type
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[3].id) }"
                                            @click="toggle(tooltips[3].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[3].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-14" style="font-size: 10px;">Please add a brief description
                                            of
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
                                        </p>
                                    </div>
                                </div>
                                <select
                                    class="form-select mt-2 mb-4 appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-grayC rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
                                    style="height: 61px;" aria-label="Default select example">
                                    <option class="text-xs md:text-sm">
                                        Partial time
                                    </option>
                                    <option class="text-xs md:text-sm">
                                        Full time
                                    </option>
                                </select>
                            </div>
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Location
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[4].id) }"
                                            @click="toggle(tooltips[4].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[4].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-6" style="font-size: 10px;">
                                            This field requires you to enter the town you are/were employed in. For remote
                                            workers,
                                            please enter
                                            the town you are/were based in.
                                        </p>
                                    </div>
                                </div>
                                <select
                                    class="mt-2 mb-4 form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-grayC rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
                                    style="height: 61px;" aria-label="Default select example">
                                    <option class="text-xs md:text-sm">
                                        Remote
                                    </option>
                                    <option class="text-xs md:text-sm">
                                        On Site
                                    </option>
                                </select>
                            </div>
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Company email
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[5].id) }"
                                            @click="toggle(tooltips[5].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[5].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-6" style="font-size: 10px;">
                                            This field requires you to enter the pre-tax/gross salary you get, vacation
                                            money
                                            included. You can
                                            find this amount in your contract or on your paycheck.
                                        </p>
                                    </div>
                                </div>
                                <input type="text" style="height: 61px"
                                    class="border border-grayC mt-2 w-full rounded-md mb-4" />
                            </div>
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Salary range
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[6].id) }"
                                            @click="toggle(tooltips[6].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[6].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-6" style="font-size: 10px;">
                                            This field requires you to enter the pre-tax/gross salary you get, vacation
                                            money
                                            included. You can
                                            find this amount in your contract or on your paycheck.
                                        </p>
                                    </div>
                                </div>
                                <div class="flex flex-col md:flex-row">
                                    <input type="number" style="height: 61px"
                                        class="border border-grayC mt-2 w-full md:w-1/2 rounded-md mb-4 md:mr-2"
                                        placeholder="Minimum" />
                                    <input type="number" style="height: 61px"
                                        class="border border-grayC mt-2 w-full md:w-1/2 rounded-md mb-4 md:ml-2"
                                        placeholder="Maximum" />
                                </div>
                            </div>
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Apply URL
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[7].id) }"
                                            @click="toggle(tooltips[7].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[7].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-6" style="font-size: 10px;">
                                            How will you want people to apply for this job?
                                        </p>
                                    </div>
                                </div>
                                <input type="text" style="height: 61px"
                                    class="border border-grayC mt-2 w-full rounded-md mb-4" />
                            </div>
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Keywords
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[8].id) }"
                                            @click="toggle(tooltips[8].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[8].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-normal h-6" style="font-size: 10px;">
                                            How will you want people to apply for this job? Enter the email or any other
                                            infos that
                                            will help
                                            people apply for the job
                                        </p>
                                    </div>
                                </div>
                                <input type="text" style="height: 61px"
                                    class="border border-grayC mt-2 w-full rounded-md mb-4" />
                            </div>
                            <div>
                                <p class="font-normal h-8 md:w-3/4" style="font-size: 10px;">
                                    Please upload your logo. If no logo is upload , a default logo will be used. Using a
                                    personalised
                                    logo increases your job offer visibility by 60%
                                </p>
                                <div class="w-40 h-32 bg-white rounded-lg mt-3 flex items-center">
                                    <div class="flex justify-center items-center mx-auto cursor-pointer">
                                        <img :src="uploadImg" class="w-6 h-6" />
                                        <h4 class="ml-1 text-base font-semibold text-grayC">Upload logo</h4>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="mt-10 flex flex-col md:flex-row items-center justify-center xl:justify-start">
                            <div class="w-full md:w-1/5 ml-0 pt-6 md:pt-0 md:ml-40" @click='selectTabPayment()'>
                                <Button show-picture="nothing"
                                    my-style="background: #235365;padding: 20px 80px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);"
                                    name="Next" />
                            </div>
                        </div>
                    </div>
                </form>
                <div v-if="isActivePayment === true">
                    <div>
                        <div class="flex flex-col md:w-11/12 mx-auto mt-14">
                            <h4 class="font-bold text-5xl leading-10 mx-auto text-center">
                                Payment instruction
                            </h4>
                            <p style="color: #5b5b5b;" class="mt-6 text-center text-base mx-10 font-light">
                                Before you proceed with your payment, here is what you need to know!
                            </p>
                            <div class="mt-12 flex justify-center items-center">
                                <div class="bg-primaryGray p-2 rounded-md md:w-3/5">
                                    <p style="font-size: 10px;" class="bg-gray-100 w-40 text-center py-1 rounded-md">
                                        Powered by Open Collective
                                    </p>
                                    <div
                                        class="flex flex-col items-center justify-center text-white text-2xl md:text-5xl mt-14">
                                        <h1 class="font-extrabold">10 000<span class="text-lg">FCFA</span></h1>
                                        <div class="mt-10 text-xs leading-8">
                                            <p v-for="(item, index) in instructions" :key="index">&#8226; {{ item }}</p>
                                        </div>
                                        <div class="md:w-1/4 py-6" @click='selectTabPaymentStatus()'>
                                            <Button show-picture="nothing"
                                                my-style="background: #235365;padding: 20px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2); width: 164px;"
                                                name="Proceed payment" />
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
                <div v-if="statusPayment === true">
                    <div>
                        <div class="flex flex-col md:w-11/12 mx-auto mt-14">
                            <h4 class="font-bold text-5xl leading-10 mx-auto text-center">
                                Payment successfully
                            </h4>
                            <p style="color: #5b5b5b;" class="mt-6 text-center text-base mx-10 font-light">
                                Thanks for making the payment and supporting OSS Caemeroon
                            </p>
                            <div class="flex items-center justify-center my-6">
                                <img :src="checkImg" alt="" class="h-96 w-1/3 object-contain" />
                            </div>
                            <p style="color: #235365;" class="mt-6 text-center text-xl md:w-1/2 mx-auto font-bold">
                                Your job is now approved. Go to the job page to see your job details
                            </p>
                        </div>

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
            isActiveOption: false,
            isActiveJob: true,
            isActivePayment: false,
            statusPayment: false,
            tooltips: [
                {
                    id: 1,
                    name: 'company',
                },
                {
                    id: 2,
                    name: 'monthly',
                },
                {
                    id: 3,
                    name: 'comment',
                },
                {
                    id: 4,
                    name: 'jobtitle',
                },
                {
                    id: 5,
                    name: 'city',
                },
                {
                    id: 6,
                    name: 'seniority',
                },
                {
                    id: 7,
                    name: 'rate',
                },
                {
                    id: 8,
                    name: 'rate',
                },
                {
                    id: 9,
                    name: 'rate',
                },
                {
                    id: 10,
                    name: 'rate',
                },
                {
                    id: 11,
                    name: 'rate',
                },
                {
                    id: 12,
                    name: 'rate',
                },
                {
                    id: 13,
                    name: 'rate',
                },
                {
                    id: 14,
                    name: 'rate',
                },
                {
                    id: 15,
                    name: 'rate',
                },
                {
                    id: 16,
                    name: 'rate',
                },
                {
                    id: 17,
                    name: 'rate',
                },
            ],
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
    methods: {
        toggle(id) {
            const index = this.opened.indexOf(id)
            if (index > -1) {
                this.opened.splice(index, 1)
            } else {
                this.opened.push(id)
            }
        },

        selectTabJob() {
            this.isActiveOption = false
            this.isActivePayment = false
            this.isActiveJob = true
            this.statusPayment = false
        },
        selectTabOption() {
            this.isActiveOption = true
            this.isActivePayment = false
            this.isActiveJob = false
            this.statusPayment = false
        },
        selectTabPayment() {
            this.isActivePayment = true
            this.isActiveOption = false
            this.isActiveJob = false
            this.statusPayment = false
        },
        selectTabPaymentStatus() {
            this.isActivePayment = false
            this.statusPayment = true
        }

    },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
