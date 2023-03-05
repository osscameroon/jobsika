<template>
    <form>
        <div class="flex flex-col md:w-11/12 mx-auto">
            <div class="grid gap-7 grid-cols-1 xl:grid-cols-2 mt-20">
                <div>
                    <div class="flex">
                        <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                            Company Name
                        </p>
                        <div class="flex items-center">
                            <span
                                class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                :class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[0].id)">
                                !
                            </span>
                        </div>
                    </div>
                    <p class="font-normal h-6" style="font-size: 9px;">Enter your company's brand/trade name: without Inc.,
                        Ltd., B.V.,
                        Pte., etc.</p>
                    <div v-if="opened.includes(tooltips[0].id)" class="w-full bg-primary">
                        <div class="w-full my-1">
                            <p class="font-light text-xs leading-5">
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum amet non
                                doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates nesciunt
                                consectetur dignissimos aut nobis, magni et voluptas aperiam!
                            </p>
                        </div>
                    </div>
                    <input type="text" style="height: 61px" class="border border-grayC mt-2 w-full rounded-md mb-4" />
                </div>
                <div>
                    <div class="flex">
                        <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                            Company email
                        </p>
                        <div class="flex items-center">
                            <span
                                class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                :class="{ opened: opened.includes(tooltips[1].id) }" @click="toggle(tooltips[0].id)">
                                !
                            </span>
                        </div>
                    </div>
                    <p class="font-normal h-6" style="font-size: 9px;">This email is private and will be used to send the
                        invoice</p>
                    <div v-if="opened.includes(tooltips[1].id)" class="w-full bg-primary">
                        <div class="w-full my-1">
                            <p class="font-light text-xs leading-5">
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum amet non
                                doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates nesciunt
                                consectetur dignissimos aut nobis, magni et voluptas aperiam!
                            </p>
                        </div>
                    </div>
                    <input type="text" style="height: 61px" class="border border-grayC mt-2 w-full rounded-md mb-4" />
                </div>
                <div>
                    <div class="flex">
                        <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                            Job title
                        </p>
                        <div class="flex items-center">
                            <span
                                class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                :class="{ opened: opened.includes(tooltips[2].id) }" @click="toggle(tooltips[0].id)">
                                !
                            </span>
                        </div>
                    </div>
                    <p class="font-normal h-6" style="font-size: 9px;">This field requires you to enter the title of the
                        position and/or
                        specialization (if applicable). E.g. "Accountant(Audit)"</p>
                    <div v-if="opened.includes(tooltips[2].id)" class="w-full bg-primary">
                        <div class="w-full my-1">
                            <p class="font-light text-xs leading-5">
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum amet non
                                doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates nesciunt
                                consectetur dignissimos aut nobis, magni et voluptas aperiam!
                            </p>
                        </div>
                    </div>
                    <input type="text" style="height: 61px" class="border border-grayC mt-2 w-full rounded-md mb-4" />
                </div>
                <div>
                    <div class="flex">
                        <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                            Employment type
                        </p>
                        <div class="flex items-center">
                            <span
                                class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                :class="{ opened: opened.includes(tooltips[3].id) }" @click="toggle(tooltips[0].id)">
                                !
                            </span>
                        </div>
                    </div>
                    <p class="font-normal h-6" style="font-size: 9px;"></p>
                    <div v-if="opened.includes(tooltips[3].id)" class="w-full bg-primary">
                        <div class="w-full my-1">
                            <p class="font-light text-xs leading-5">
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum amet non
                                doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates nesciunt
                                consectetur dignissimos aut nobis, magni et voluptas aperiam!
                            </p>
                        </div>
                    </div>
                    <input type="text" style="height: 61px" class="border border-grayC mt-2 w-full rounded-md mb-4" />
                </div>
                <div>
                    <div class="flex">
                        <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                            Monthly salary
                        </p>
                        <div class="flex items-center">
                            <span
                                class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                :class="{ opened: opened.includes(tooltips[4].id) }" @click="toggle(tooltips[0].id)">
                                !
                            </span>
                        </div>
                    </div>
                    <p class="font-normal h-6" style="font-size: 9px;">
                        This field requires you to enter the pre-tax/gross salary you get, vacation money included. You can
                        find this amount in your contract or on your paycheck.
                    </p>
                    <div v-if="opened.includes(tooltips[4].id)" class="w-full bg-primary">
                        <div class="w-full my-1">
                            <p class="font-light text-xs leading-5">
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum amet non
                                doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates nesciunt
                                consectetur dignissimos aut nobis, magni et voluptas aperiam!
                            </p>
                        </div>
                    </div>
                    <input type="text" style="height: 61px" class="border border-grayC mt-2 w-full rounded-md mb-4" />
                </div>
                <div>
                    <div class="flex">
                        <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                            Location
                        </p>
                        <div class="flex items-center">
                            <span
                                class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                :class="{ opened: opened.includes(tooltips[5].id) }" @click="toggle(tooltips[0].id)">
                                !
                            </span>
                        </div>
                    </div>
                    <p class="font-normal h-6" style="font-size: 9px;">
                        This field requires you to enter the town you are/were employed in. For remote workers, please enter
                        the town you are/were based in.
                    </p>
                    <div v-if="opened.includes(tooltips[5].id)" class="w-full bg-primary">
                        <div class="w-full my-1">
                            <p class="font-light text-xs leading-5">
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum amet non
                                doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates nesciunt
                                consectetur dignissimos aut nobis, magni et voluptas aperiam!
                            </p>
                        </div>
                    </div>
                    <input type="text" style="height: 61px" class="border border-grayC mt-2 w-full rounded-md mb-4" />
                </div>
                <div>
                    <div class="flex">
                        <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                            Apply URL
                        </p>
                        <div class="flex items-center">
                            <span
                                class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                :class="{ opened: opened.includes(tooltips[6].id) }" @click="toggle(tooltips[0].id)">
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
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum amet non
                                doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates nesciunt
                                consectetur dignissimos aut nobis, magni et voluptas aperiam!
                            </p>
                        </div>
                    </div>
                    <input type="text" style="height: 61px" class="border border-grayC mt-2 w-full rounded-md mb-4" />
                </div>
                <div>
                    <div class="flex">
                        <p class="text-xs font-bold md:text-lg" style="color: #b1b1b1;">
                            Upload logo
                        </p>
                        <div class="flex items-center">
                            <span
                                class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
                                :class="{ opened: opened.includes(tooltips[7].id) }" @click="toggle(tooltips[0].id)">
                                !
                            </span>
                        </div>
                    </div>
                    <p class="font-normal h-6" style="font-size: 9px;">
                        Please upload your logo
                    </p>
                    <div v-if="opened.includes(tooltips[7].id)" class="w-full bg-primary">
                        <div class="w-full my-1">
                            <p class="font-light text-xs leading-5">
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat fugiat harum amet non
                                doloribus tenetur aliquid fuga pariatur, repudiandae, voluptatem voluptates nesciunt
                                consectetur dignissimos aut nobis, magni et voluptas aperiam!
                            </p>
                        </div>
                    </div>
                    <input type="text" style="height: 61px" class="border border-grayC mt-2 w-full rounded-md mb-4" />
                </div>
            </div>
            <div class="mt-10 flex flex-col md:flex-row items-center justify-center xl:justify-start">
                <div class="w-full md:w-1/5">
                    <NuxtLink to="/">
                        <Button show-picture="nothing" name="Cancel"
                            my-style="background: white; color: #000000; padding: 15px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);" />
                    </NuxtLink>
                </div>
                <button class="w-full md:w-1/5 ml-0 pt-6 md:pt-0 md:ml-6" type="submit">
                    <Button
                        my-style="background: #235365; font-family: 'Inter', sans-serif;padding: 15px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);"
                        name="Add" />
                </button>
            </div>
        </div>
    </form>
</template>

<script>
export default {
    name: "JobDetailCompoent",
    data() {
        return {
            opened: [],
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
