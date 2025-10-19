<script setup>
import Header from './components/Header.vue' 

</script>

<template>
  <div id="app">
    <Header></Header>
    <main class="app-container">
      <router-view />
    </main>
  </div>
</template>


<style scoped> 
.app-container {
  padding-top: 44px;
}
</style>