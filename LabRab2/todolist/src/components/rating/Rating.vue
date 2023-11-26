<script setup>
import StarRating from "vue-star-rating"
</script>

<template>
  <div @mouseleave="mouseOverRating = 0" style="display:inline-block;">
    <star-rating :show-rating="false" @hover:rating="mouseOverRating = $event" :increment="0.5"></star-rating>
  </div>
  <div style="margin-top:10px;font-weight:bold;">{{mouseOverRatingText}}</div>
</template>

<script>
export default {
  computed: {
    mouseOverRatingText() {
      switch (this.mouseOverRating) {
        case 0: return "Нет рейтинга"
        case 0.5:
        case 1: return "Ужасно"
        case 1.5:
        case 2: return "Плохо"
        case 2.5:
        case 3: return "Средне"
        case 3.5:
        case 4: return "Хорошо"
        case 4.5:
        case 5: return "Отлично"
      }
    }
  },
  data() {
    return {
      rating: 0,
      resetableRating: 2,
      currentRating: "No Rating",
      mouseOverRating: null
    };
  }
}
</script>