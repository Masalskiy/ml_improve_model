<script setup lang="ts">
import { ref } from 'vue'
import DataTable from './components/DataTable.vue'

// Пример данных (замените на реальные данные из вашей базы данных)
const headers = ref(['ID', 'Название', 'Описание', 'Статус'])
const rows = ref([
  [1, 'Запись 1', 'Описание записи 1', 'Активно'],
  [2, 'Запись 2', 'Описание записи 2', 'Неактивно'],
  [3, 'Запись 3', 'Описание записи 3', 'Активно']
])

const handleSelectionChange = (event: { index: number; selected: boolean; row: any[] }) => {
  console.log('Выбрана строка:', event.row)
  console.log('Состояние:', event.selected)
}
</script>

<template>
  <div class="app">
    <h1>Данные из базы</h1>
    <DataTable
      :headers="headers"
      :rows="rows"
      @selection-change="handleSelectionChange"
    />
  </div>
</template>

<style>
.app {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  color: #2c3e50;
  margin-bottom: 20px;
}
</style>
