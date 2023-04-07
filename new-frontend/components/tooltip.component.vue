<script setup lang="ts">
  const { title, content } = defineProps({
		title: {type: String, required: true},
    content: {type: String, required: true}
	});
</script>

<template>
  <div class="modal-overlay" @click="$emit('close-modal')">
    <div class="modal" @click.stop>
      <h3
        class="title text-3xl text-center pb-4"
        style="color: #000000; font-family: 'Inter', sans-serif"
      >
        {{ title }}
      </h3>
      <p
        class="text-sm text-center mx-auto leading-6"
        style="color: #000000; font-family: 'Inter', sans-serif"
      >
        {{ content }}
      </p>
      <div class="mt-5">
        <div class="close" @click="$emit('close-modal')">
          <button
            class="bg-blueDark py-4 w-full text-white px-4 rounded-md text-sm"
          >
            Back to site
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

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
  /* text-align: center; */
  background-color: white;
  height: 300px;
  width: 500px;
  margin-top: 6%;
  padding: 20px 30px;
  border-radius: 5px;
}
.close,
.send {
  margin: 10% 0 0 16px;
  cursor: pointer;
}

h6 {
  font-weight: 500;
  font-size: 28px;
  margin: 20px 0;
}
</style>
