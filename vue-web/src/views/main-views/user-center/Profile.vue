<template>
<div class="profile-body">
  <div class="profile-container">

  </div>
</div>
</template>

<script setup>

</script>

<style scoped>
.profile-body {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: linear-gradient(rgba(255, 255, 255, 0.8), rgba(255, 255, 255, 0.8)), url("@/res/login-background.jpeg") no-repeat center;
  background-size: cover;
}
.profile-container {
  width: 960px;
  height: 540px;
  max-width: 960px;
  max-height: 540px;
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0 21px 41px 0 rgba(0, 0, 0, 0.2);
}
.profile-container::before {
  content: "";
  display: block;
  padding-top: 56.25%; /* 100% / (960 / 540) = 56.25% */
}

.profile-container > * {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  /* 现在你可以在这里设置你的内容布局，比如使用 flexbox 或 grid */
}
</style>