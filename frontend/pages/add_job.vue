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
                <Tabs>
                    <Tab title="Job details">
                        <JobDetail />
                    </Tab>
                    <Tab title="Options">
                        <JobOptions />
                    </Tab>
                    <Tab title="Payments">
                        <Publish />
                    </Tab>
                </Tabs>
            </div>
        </div>
    </main>
</template>

<script>
export default {
    name: 'JobsIndex',
    layout: 'app',
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
