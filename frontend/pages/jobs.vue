<template>
    <main>
        <div class="container mx-auto md:w-7/12 px-4 md:px-0">
            <div class="pt-10 lg:pt-14">
                <h2 class="font-bold text-2xl md:text-4xl text-center text-black">
                    Find the job that fits your skills and passion
                </h2>
            </div>
            <div class="pt-8 md:pt-16">
                <div class="flex flex-col xl:flex-row xl:justify-end">
                    <p
                        class="font-normal text-sm md:text-lg flex justify-center items-center pt-4 md:pt-2 xl:pt-0 text-center text-gray-700">
                        Receive the jobs right in your inbox as soon as they are posted
                    </p>
                    <div class="pt-4 md:pt-2 xl:pt-0 w-full xl:w-36 xl:ml-4">
                        <button @click="showModal = true"
                            class="mx-auto xl:mx-0 w-1/2 xl:w-36 cursor-pointer p-4 text-blueDark font-bold flex items-center justify-center text-xs md:text-lg px-4 border-2 h-12 rounded-md border-blue">
                            Subscribe
                        </button>
                    </div>
                </div>
            </div>
            <div class="pt-8">
                <div v-for="(item, index) in jobs" :key="index" class="mb-6">
                    <JobPost :picture="item.picture" :title="item.title" :city="item.city" :structure="item.structure"
                        :marker="item.marker" :time="item.time" :description="item.description" :tags="item.tags" />
                </div>
            </div>
          <JobPagination />
        </div>
        <SubmitModal v-show="showModal" @close-modal="showModal = false" />
    </main>
</template>

<script>
export default {
    name: 'JobsIndex',
    layout: 'app',
    data() {
        return {
            showModal: false,
            jobs: [
                {
                    title: "Flutter Developer",
                    structure: "Ejara",
                    marker: "Hybrid or remote",
                    city: "in Douala",
                    time: "2 days",
                    description: "Au-delà du business, Ejara est un rêve, un élan, une énergie qui vise à fournir à l'Afrique francophone des outils efficaces pour faciliter et faire de la ...",
                    picture: require('../assets/ejara.png'),
                    tags: ["Figma", "Vue", "UI/UX Design"]
                },
                {
                    title: "Frontend Developer",
                    structure: "OSS Cameroon",
                    marker: "Remote",
                    city: "in Yaounde",
                    time: "4 days",
                    description: "Au-delà du business, Ejara est un rêve, un élan, une énergie qui vise à fournir à l'Afrique francophone des outils efficaces pour faciliter et faire de la ...",
                    picture: require('../assets/job.png'),
                    tags: ["Figma", "Nuxt", "UI/UX Design"]
                },
                {
                    title: "Flutter Developer",
                    structure: "Ejara",
                    marker: "Hybrid or remote",
                    city: "in Douala",
                    time: "2 days",
                    description: "Au-delà du business, Ejara est un rêve, un élan, une énergie qui vise à fournir à l'Afrique francophone des outils efficaces pour faciliter et faire de la ...",
                    picture: require('../assets/ejara.png'),
                    tags: ["Figma", "Vue", "UI/UX Design"]
                },
                {
                    title: "Frontend Developer",
                    structure: "OSS Cameroon",
                    marker: "Remote",
                    city: "in Yaounde",
                    time: "4 days",
                    description: "Au-delà du business, Ejara est un rêve, un élan, une énergie qui vise à fournir à l'Afrique francophone des outils efficaces pour faciliter et faire de la ...",
                    picture: require('../assets/job.png'),
                    tags: ["Figma", "Nuxt", "UI/UX Design"]
                },
            ]
        }
    },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
