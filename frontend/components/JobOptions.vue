<template>
    <div class="flex flex-col md:w-11/12 mx-auto mt-20">
        <ul>
            <li v-for="(item, index) in options" :key="index">
                <p class="font-bold text-lg leading-10">
                    <span class="mr-2">
                        &#x25CF;
                    </span>
                    {{ item }}
                </p>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: "JobOptionComponent",
    data() {
        return {
            options: [
                "Company logo (+2.000XAF)",
                "Diffusion on OSS Groups and accounts ( (8 x users in our databases) XAF)",
                "Personal notifications(Email & Telegram (JobSika Bot)) ( (15 x users in our databases) XAF)",
                "Pin the Job offer on OSS Cameroon groups/account (+ 5.000XAF/week)",
                "Pin the Job offer on Jobsika homepage (+10.000 XAF/week)"
            ]
        }
    }
}
</script>
