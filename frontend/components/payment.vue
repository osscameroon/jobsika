<template>
  <main>
    <div class="container mx-auto md:w-7/12 px-4 md:px-0">
      <div>
        <div class="flex flex-col md:w-11/12 mx-auto mt-14">
          <h4 class="font-bold text-5xl leading-10 mx-auto text-center">
            Payment successfully
          </h4>
          <p
            style="color: #5b5b5b"
            class="mt-6 text-center text-base mx-10 font-light"
          >
            Thanks for making the payment and supporting OSS Caemeroon
          </p>
          <div class="flex items-center justify-center my-6">
            <img :src="checkImg" alt="" class="h-96 w-1/3 object-contain" />
          </div>
          <p
            style="color: #235365"
            class="mt-6 text-center text-xl md:w-1/2 mx-auto font-bold"
          >
            Your job is now approved. Go to the job page to see your job details
          </p>
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
      uploadImg: require('../assets/upload.png'),
      checkImg: require('../assets/payment.png'),
    }
  },
}
</script>
<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@500;700&display=swap');
</style>
