<template>
  <transition name="modal-fade">
    <div class="modal-overlay" @click="$emit('close-modal')">
      <div class="modal" @click.stop>
        <div />
        <div>
          <h4 class="font-semibold text-2xl xl:text-4xl text-center text-black">
            Subscription<br />sucessful
          </h4>
          <h4 class="font-light text-base leading-5 text-center mt-6">
            You will now receive job offers<br />right to your inbox
          </h4>
        </div>
        <div class="flex items-center justify-center w-11/12">
          <button
            class="cursor-pointer p-4 text-white font-bold flex items-center justify-center text-sm lg:text-base bg-blueDark h-12 rounded-lg w-full mx-7"
            @click="$emit('close-modal')"
          >
            Alright
          </button>
        </div>
      </div>
      <div class="close" @click="$emit('close-modal')">
        <p class="font-semibold text-lg text-gray-500">X</p>
      </div>
    </div>
  </transition>
</template>

<script>
export default {}
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
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
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
