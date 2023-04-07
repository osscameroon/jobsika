<script setup lang="ts">
  const { grade, maxStars, hasCounter } = defineProps({
		grade: {type: Number, required: true, default: ""},
    maxStars: {type: Number, required: true},
    hasCounter: {type: Boolean, required: true},
	});

  const emit = defineEmits(['blurall', 'changeGrade']);

  const stars: Ref<number> = ref(grade);
  const staru: Ref<string> = ref('/staru.png');
  const starc: Ref<string> = ref('/starc.png');

  const rate = (star: any) => {
    emit('blurall');
    if (typeof star === 'number' && star <= maxStars && star >= 0) {
      stars.value = stars.value === star ? star - 1 : star
      emit('blurall', stars);
    }
  }
</script>

<template>
  <div class="site__rating w-5/12 mx-auto mb-2 md:mb-8">
    <ul class="site__rating-list flex items-center justify-between">
      <li
        v-for="star in maxStars"
        @click="rate(star)"
        :class="{ active: star <= stars }"
        :key="star"
        class="star"
        style="cursor: pointer"
      >
        <img
          class="w-7 h-7 px-1 md:px-0"
          :src="star <= stars ? starc : staru"
        />
      </li>
    </ul>
  </div>
</template>