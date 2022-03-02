<template>
  <div class="site__rating ml-6">
    <ul class="site__rating-list flex">
      <li
        v-for="star in maxStars"
        @click="rate(star)"
        :class="{ active: star <= stars }"
        :key="star.stars"
        class="star"
      >
        <img class="w-4 h-4 mr-1" :src="star <= stars ? starc : staru" />
      </li>
    </ul>
  </div>
</template>
<script>
export default {
  name: 'RatingComponent',
  props: ['grade', 'maxStars', 'hasCounter'],
  data() {
    return {
      stars: this.grade,
      staru: require('../assets/staru.png'),
      starc: require('../assets/starc.png'),
    }
  },
  methods: {
    rate(star) {
      this.$emit("blurall");
      if (typeof star === 'number' && star <= this.maxStars && star >= 0) {
        this.stars = this.stars === star ? star - 1 : star
        this.$emit("changeGrade", this.stars);
      }
    },
  },
}
</script>
