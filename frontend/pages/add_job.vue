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
                        class="px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4">
                        Job details
                    </div>
                    <div
                        class="px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray50 h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4">
                        Options
                    </div>
                    <div
                        class="px-14 py-4 text-white font-bold flex items-center justify-center lg:text-base bg-primaryGray50 h-14 rounded-lg mt-8 text-sm w-full md:w-56 md:mx-4">
                        Payments
                    </div>
                </div>
                <form>
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
                                <p class="font-normal h-6" style="font-size: 9px;">Enter your company's brand/trade name:
                                    without Inc.,
                                    Ltd., B.V., Pte., etc.</p>
                                <div v-if="opened.includes(tooltips[0].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p style="font-size: 9px;" class="font-light text-xs leading-5">
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
                                        </p>
                                    </div>
                                </div>
                                <input type="text" style="height: 61px"
                                    class="border border-grayC mt-2 w-full rounded-md mb-4" />
                            </div>
                            <div>
                                <div class="flex">
                                    <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                                        Job tittle
                                    </p>
                                    <div class="flex items-center">
                                        <span
                                            class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                            :class="{ opened: opened.includes(tooltips[1].id) }"
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <p class="font-normal h-6" style="font-size: 9px;">This field requires you to enter the
                                    title of the
                                    position and/or specialization (if applicable). E.g. "Accountant(Audit)"</p>
                                <div v-if="opened.includes(tooltips[1].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p style="font-size: 9px;" class="font-light text-xs leading-5">
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
                                        </p>
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
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <p class="font-normal h-6" style="font-size: 9px;">Please add a brief description of the
                                    role you are
                                    hiring for, make it straight forward and concise</p>
                                <div v-if="opened.includes(tooltips[2].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-light text-xs leading-5">
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
                                        </p>
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
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <div v-if="opened.includes(tooltips[3].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-light text-xs leading-5">
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
                                            :class="{ opened: opened.includes(tooltips[3].id) }"
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <p class="font-normal h-6" style="font-size: 9px;">
                                    This field requires you to enter the town you are/were employed in. For remote workers,
                                    please enter
                                    the town you are/were based in.
                                </p>
                                <div v-if="opened.includes(tooltips[4].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-light text-xs leading-5">
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
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
                                            :class="{ opened: opened.includes(tooltips[4].id) }"
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <p class="font-normal h-6" style="font-size: 9px;">
                                    This field requires you to enter the pre-tax/gross salary you get, vacation money
                                    included. You can
                                    find this amount in your contract or on your paycheck.
                                </p>
                                <div v-if="opened.includes(tooltips[4].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-light text-xs leading-5">
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
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
                                            :class="{ opened: opened.includes(tooltips[4].id) }"
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <p class="font-normal h-6" style="font-size: 9px;">
                                    This field requires you to enter the pre-tax/gross salary you get, vacation money
                                    included. You can
                                    find this amount in your contract or on your paycheck.
                                </p>
                                <div v-if="opened.includes(tooltips[4].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-light text-xs leading-5">
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
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
                                            :class="{ opened: opened.includes(tooltips[6].id) }"
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <p class="font-normal h-6" style="font-size: 9px;">
                                    How will you want people to apply for this job?
                                </p>
                                <div v-if="opened.includes(tooltips[6].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-light text-xs leading-5">
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
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
                                            :class="{ opened: opened.includes(tooltips[5].id) }"
                                            @click="toggle(tooltips[0].id)">
                                            !
                                        </span>
                                    </div>
                                </div>
                                <p class="font-normal h-6" style="font-size: 9px;">
                                    How will you want people to apply for this job? Enter the email or any other infos that
                                    will help
                                    people apply for the job
                                </p>
                                <div v-if="opened.includes(tooltips[5].id)" class="w-full bg-primary">
                                    <div class="w-full my-1">
                                        <p class="font-light text-xs leading-5">
                                            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum
                                            amet non
                                            doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates
                                            nesciunt
                                            consectetur dignissimos aut nobis, magni et voluptas aperiam!
                                        </p>
                                    </div>
                                </div>
                                <input type="text" style="height: 61px"
                                    class="border border-grayC mt-2 w-full rounded-md mb-4" />
                            </div>
                            <div>
                                <p class="font-normal h-6 md:w-3/4" style="font-size: 9px;">
                                    Please upload your logo. If no logo is upload , a default logo will be used. Using a
                                    personalised
                                    logo increases your job offer visibility by 60%
                                </p>
                                <div class="w-40 h-32 bg-white rounded-lg mt-3 flex items-center">
                                    <div class="flex justify-center items-center mx-auto">
                                        <img :src="uploadImg" class="w-6 h-6" />
                                        <h4 class="ml-1 text-base font-semibold text-grayC">Upload logo</h4>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="mt-10 flex flex-col md:flex-row items-center justify-center xl:justify-start">
                            <NuxtLink to="/options">
                                <button class="w-full md:w-1/5 ml-0 pt-6 md:pt-0 md:ml-40" type="submit">
                                    <Button show-picture="nothing"
                                        my-style="background: #235365;padding: 20px 80px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);"
                                        name="Next" />
                                </button>
                            </NuxtLink>
                        </div>
                    </div>
                </form>
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
            ],
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

    },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
