<template>
  <div class="data-table">
    <table class="table">
      <thead>
        <tr>
          <th>Выбрать</th>
          <th v-for="(header, index) in headers" :key="index">{{ header }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, rowIndex) in rows" :key="rowIndex">
          <td>
            <input
              type="checkbox"
              v-model="selectedRows[rowIndex]"
              @change="handleCheckboxChange(rowIndex)"
            />
          </td>
          <td v-for="(value, colIndex) in row" :key="colIndex">
            {{ value }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

interface Props {
  headers: string[]
  rows: any[]
}

const props = defineProps<Props>()
const emit = defineEmits(['selectionChange'])

const selectedRows = ref<{ [key: number]: boolean }>({})

const handleCheckboxChange = (rowIndex: number) => {
  emit('selectionChange', {
    index: rowIndex,
    selected: selectedRows.value[rowIndex],
    row: props.rows[rowIndex]
  })
}

// Очистка выбранных строк при изменении входных данных
watch(() => props.rows, () => {
  selectedRows.value = {}
})
</script>

<style scoped>
.data-table {
  margin: 20px;
}

.table {
  width: 100%;
  border-collapse: collapse;
  background-color: white;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
}

.table th,
.table td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.table th {
  background-color: #f5f5f5;
  font-weight: 600;
}

.table tr:hover {
  background-color: #f9f9f9;
}

input[type="checkbox"] {
  cursor: pointer;
  width: 18px;
  height: 18px;
}
</style> 