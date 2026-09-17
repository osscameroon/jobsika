<script setup lang="ts">
  const { title, myStyle, datas } = defineProps({
		title: {type: String, required: true},
    myStyle: {type: String, required: true},
    datas: {type: Array<string>, required: true},
	});

  const name: Ref<string> = ref("");
  const isOpen: Ref<boolean> = ref(false);
  
  const filteredData = computed(() => {
    if(name.value === ""){
        return []
      }else{
        return datas.map((elem: string) => elem.toLowerCase()).filter(
          (elem) => !elem.indexOf(name.value.toLowerCase())
        )
      }
  });
  const opened = computed(() => filteredData.value.length > 0);

  const setResult = (result: string) => {
    name.value = result;
    isOpen.value = false;
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
      v-model="name"
      type="text"
      class="site__input-field border-none mt-3 w-full rounded-md mb-8 md:mb-16"
      :style="myStyle"
    />
    <ul
      style="background: white"
      class="
        h-20
        md:h-32
        overflow-y-scroll
        rounded-lg
        -mt-3
        md:-mt-14
        mb-2
        md:mb-4
        cursor-pointer
      "
      v-if="opened"
    >
      <li
        v-for="(elem, index) in filteredData"
        :key="index"
        @click="setResult(elem)"
        class="shadow-sm py-2 px-4"
      >
        {{ elem }}
      </li>
    </ul>
  </div>
</template>