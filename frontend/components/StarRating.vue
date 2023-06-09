<template>
  <div class="site__rating w-5/12 mx-auto mb-2 md:mb-8">
    <ul class="site__rating-list flex items-center justify-between">
      <li
        v-for="star in maxStars"
        :key="star.stars"
        :class="{ active: star <= stars }"
        class="star"
        style="cursor: pointer"
        @click="rate(star)"
      >
        <img
          class="w-7 h-7 px-1 md:px-0"
          :src="star <= stars ? starc : staru"
        />
      </li>
    </ul>
  </div>
</template>
<script>
export default {
  name: 'RatingComponent',
  props: {
    grade: {
      type: String,
      default: '',
    },
    maxStars: {
      type: String,
      default: '',
    },
    hasCounter: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      stars: this.grade,
      staru: require('../assets/staru.png'),
      starc: require('../assets/starc.png'),
    }
  },
  methods: {
    rate(star) {
      this.$emit('blurall')
      if (typeof star === 'number' && star <= this.maxStars && star >= 0) {
        this.stars = this.stars === star ? star - 1 : star
        this.$emit('changeGrade', this.stars)
      }
    },
  },
}
</script>
