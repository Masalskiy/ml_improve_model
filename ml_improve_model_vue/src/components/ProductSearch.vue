<template>
  <div class="product-search">
    <div class="search-form">
      <div class="search-type">
        <label>
          <input 
            type="radio" 
            v-model="searchType" 
            value="id" 
          > Поиск по ID
        </label>
        <label>
          <input 
            type="radio" 
            v-model="searchType" 
            value="name" 
          > Поиск по названию
        </label>
      </div>

      <div class="search-input">
        <input 
          v-model="searchQuery" 
          :placeholder="searchType === 'id' ? 'Введите ID продукта' : 'Введите название продукта'"
          class="input"
        >
        <button @click="searchProducts" class="button">Поиск</button>
      </div>
    </div>

    <div v-if="error" class="error">
      {{ error }}
    </div>

    <div v-if="loading" class="loading">
      Загрузка...
    </div>

    <table v-if="products.length && !loading" class="products-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Название</th>
          <th>SKU</th>
          <th>Ссылка</th>
          <th>Изображение</th>
          <th>Описание</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.id_product">
          <td>{{ product.id_product }}</td>
          <td>{{ product.name }}</td>
          <td>{{ product.sku }}</td>
          <td>
            <a :href="product.link" target="_blank">Ссылка</a>
          </td>
          <td>
            <img 
              v-if="product.image_link" 
              :src="product.image_link" 
              :alt="product.name"
              class="product-image"
            >
          </td>
          <td>{{ product.description }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Product {
  id_product: number
  name: string
  sku: string
  link: string
  image_link: string
  description: string
}

const searchType = ref<'id' | 'name'>('id')
const searchQuery = ref('')
const products = ref<Product[]>([])
const loading = ref(false)
const error = ref('')

const searchProducts = async () => {
  if (!searchQuery.value.trim()) {
    error.value = 'Пожалуйста, введите запрос для поиска'
    return
  }

  loading.value = true
  error.value = ''
  products.value = []

  try {
    let url = 'http://localhost:8000/products'
    
    if (searchType.value === 'id') {
      url += `/${searchQuery.value}`
    } else {
      url += `/name/${searchQuery.value}`
    }

    const response = await fetch(url)
    const data = await response.json()

    if (!response.ok) {
      throw new Error(data.error || 'Ошибка при поиске продуктов')
    }

    products.value = Array.isArray(data) ? data : [data]
  } catch (err: unknown) {
    error.value = err instanceof Error ? err.message : 'Произошла ошибка при поиске'
    products.value = []
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.product-search {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.search-form {
  margin-bottom: 20px;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 8px;
}

.search-type {
  margin-bottom: 15px;
}

.search-type label {
  margin-right: 20px;
  cursor: pointer;
}

.search-input {
  display: flex;
  gap: 10px;
}

.input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

.button {
  padding: 8px 20px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.button:hover {
  background-color: #45a049;
}

.error {
  color: #ff4444;
  margin-bottom: 15px;
  padding: 10px;
  background-color: #ffebee;
  border-radius: 4px;
}

.loading {
  text-align: center;
  padding: 20px;
  color: #666;
}

.products-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

.products-table th,
.products-table td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.products-table th {
  background-color: #f5f5f5;
  font-weight: 600;
}

.products-table tr:hover {
  background-color: #f9f9f9;
}

.product-image {
  max-width: 100px;
  max-height: 100px;
  object-fit: cover;
  border-radius: 4px;
}
</style> 