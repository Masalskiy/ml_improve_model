<script setup lang="ts">
import ProductSearch from './components/ProductSearch.vue'
</script>

<template>
  <div class="app">
    <h1>Поиск продуктов</h1>
    <ProductSearch />
  </div>
</template>

<style>
.app {
  font-family: Arial, sans-serif;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  color: #2c3e50;
  text-align: center;
  margin-bottom: 30px;
}

/* Глобальные стили */
body {
  margin: 0;
  padding: 0;
  background-color: #f8f9fa;
}
</style>
