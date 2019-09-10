<template>
  <div id="app">
    <transition name="router-anim">
      <router-view/>
    </transition>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.router-anim-enter-active{
  animation: coming 1s;
  animation-delay: .5s;
  opacity: 0;
}
.router-anim-leave-active{
  animation: going 1s;
}
@keyframes going{
  from{
    transform: translateX(0);
  }
  to{
    transform: translateX(-50px);
    opacity: 0;
  }
}
@keyframes coming{
  from{
    transform: translateX(-50px);
    opacity: 0;
  }
  to{
    transform: translateX(0px);
    opacity: 1;
  }
}
</style>
