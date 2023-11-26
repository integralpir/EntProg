<template>
  <div class="row">
    <div class="col-md-4">
      <div class="card shadow rounded my-3">
        <div><img src="/person.jpg" class="avatar" alt="10"></div>
        <div class="card-body">
          <div class="avatar-post">Медик</div>
          <div class="avatar-text">Мой любимый персонаж в TF2</div>
        </div>
      </div>
    </div>
    <div class="col-md-4">
      <p>
        Меня зовут Роман Орлов. Я студент группы ПИб-41.
        Работаю Java-разработчиком в МТС Банке на полную ставку.
        Из-за вышеперечисленных факторов фронт верстаю медленно.
      </p>
    </div>
  </div>
</template>

<style>
.avatar {
  border-radius: 50%;
  width: 100%;
}
.avatar-post {
  margin-top: 1rem;
  font-size: 1.4rem;
  text-align: center;
}
.avatar-text {
  color: rgb(140, 140, 140);
  font-size: 1.1rem;
  text-align: center;
}
</style>
