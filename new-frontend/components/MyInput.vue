<script setup lang="ts">
  const { myStyle, title, myInput, typeInput } = defineProps({
		myStyle: { type: String, required: true},
    title: { type: String, required: true},
    myInput: { type: String, required: true},
    typeInput: { type: String, required: true}
  });

  const nameInput: Ref<string> =  ref('');
  const nameTextArea: Ref<string> = ref('');

  const onChangeInput = () => {
    //this.$store.dispatch('selectValueSalary', nameInput)
  }
  
  const onChangeTextArea = () => {
    //this.$store.dispatch('selectValueComment', nameTextArea)
  }
</script>

<template>
  <div class="site__input w-full">
    <p
      class="text-xs md:text-sm font-bold"
      style="color: #b1b1b1; font-family: 'Inter', sans-serif"
    >
      {{ title }}
    </p>
    <input
      v-if="myInput === 'input'"
      v-model="nameInput"
      :type="typeInput"
      class="
        site__input-field
        border-none
        mt-2
        md:mt-3
        w-full
        rounded-md
        mb-4
        md:mb-16
      "
      :style="myStyle"
      data-type="currency"
      @change="onChangeInput"
    />

    <textarea
      v-else
      v-model="nameTextArea"
      class="resize rounded-md w-full h-full mt-2 md:mt-3 border-none"
      name="story"
      rows="5"
      @change="onChangeTextArea"
    />
  </div>
</template>
