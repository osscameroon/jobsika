<template>
  <transition name="modal-fade">
    <div class="modal-overlay" @click="$emit('close-modal')">
      <div class="modal" @click.stop>
        <h4 class="font-semibold text-2xl xl:text-4xl text-center text-black">
          Subscribe
        </h4>
        <h4 class="font-light text-base leading-8 pt-2 text-center">
          Don't miss any job from JobSika
        </h4>
        <div class="mt-10">
          <div class="mx-7">
            <h4 class="text-lg font-semibold text-grayC text-left">
              Enter your email
            </h4>
            <p style="font-size: 9px" class="pt-1 font-normal text-left">
              All new job offers will be sent to your email
            </p>
            <input
              v-model="email"
              type="text"
              class="border border-grayC w-full rounded-md h-11 mt-1"
            />
          </div>
          <div class="flex items-center justify-center my-12">
            <div
              class="cursor-pointer p-4 text-white font-bold flex items-center justify-center text-sm lg:text-base bg-blueDark h-12 rounded-lg w-full mx-7"
              @click="submitform()"
            >
              Sign up
            </div>
          </div>
        </div>
      </div>
      <div class="close" @click="$emit('close-modal')">
        <p class="font-semibold text-lg text-gray-500">X</p>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  data() {
    return {
      email: '',
    }
  },
  methods: {
    submitform() {
      this.$emit('success-modal', this.email)
    },
  },
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  background-color: #000000da;
}

.modal {
  text-align: center;
  background-color: white;
  height: 550px;
  width: 479px;
  margin-top: 10%;
  padding: 60px 0;
  border-radius: 20px;
}

.close {
  margin: 11% 0 0 -50px;
  cursor: pointer;
}

.close-img {
  width: 25px;
}

.check {
  width: 150px;
}

.modal-fade-enter,
.modal-fade-leave-to {
  opacity: 0;
}

.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.5s ease;
}
</style>
