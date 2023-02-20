<template>
    <transition name="modal-fade">
        <div class="modal-overlay" @click="$emit('close-modal')">
            <div class="modal" @click.stop>
                <h4 class="font-semibold text-2xl text-center text-black">Subscribe</h4>
                <h4 class="font-light text-2xl leading-8 pt-2 text-center">Don't miss any job from JobSika</h4>
                <form class="mt-12">
                    <div class="mx-7">
                        <div>
                            <p class="text-base font-normal text-black text-left">
                                Enter you Telegram handle
                            </p>
                            <input type="text" class="border border-grayC w-full rounded-md h-16 mt-3" />
                        </div>
                        <div class="mt-9">
                            <p class="text-base font-normal text-black text-left">
                                Enter your email
                            </p>
                            <input type="text" class="border border-grayC w-full rounded-md h-16 mt-3" />
                        </div>
                    </div>
                    <div class="flex items-center justify-center mt-12">
                        <button
                            class="cursor-pointer p-4 text-white font-bold flex items-center justify-center text-sm lg:text-base bg-blueDark h-12 rounded-lg w-80">
                            Sign up
                        </button>
                    </div>
                </form>
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
    width: 500px;
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



