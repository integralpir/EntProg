<script setup>
</script>

<template>
  <h1>Магазин одежды</h1>
  <p>Добро пожаловать!</p>
  <div id="carouselExampleControls" class="carousel slide" data-bs-ride="carousel">
    <div class="carousel-inner">
      <div class="carousel-item active">
        <img src="/assets/img/main1.jpg" class="d-block w-100" alt="...">
      </div>
      <div class="carousel-item">
        <img src="/assets/img/main2.jpg" class="d-block w-100" alt="...">
      </div>
      <div class="carousel-item">
        <img src="/assets/img/main3.jpg" class="d-block w-100" alt="...">
      </div>
    </div>
    <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleControls" data-bs-slide="prev">
      <span class="carousel-control-prev-icon" aria-hidden="true"></span>
      <span class="visually-hidden">Previous</span>
    </button>
    <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleControls" data-bs-slide="next">
      <span class="carousel-control-next-icon" aria-hidden="true"></span>
      <span class="visually-hidden">Next</span>
    </button>
  </div>
  
</template>

<script>
export default {
  data() {
    return {
    }
  },
}
</script>
