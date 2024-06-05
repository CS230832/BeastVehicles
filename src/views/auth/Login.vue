<script setup>
import { ref } from 'vue'
import ApiService from '@/api'

import Toast from 'primevue/toast'
import Card from 'primevue/card'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import { useToast } from 'primevue/usetoast'

const toast = useToast()
const errorMessage = ref(null)

const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: errorMessage.value,
    life: 3000
  })
}

const username = ref(null)
const password = ref(null)
const data = ref(null)

const login = async () => {
  try {
    data.value = await ApiService.login(username.value, password.value)
  } catch (error) {
    password.value = null
    errorMessage.value = error.response.data.data
    showErrorMessage()
  } finally {
    localStorage.setItem('token', data.value.data)
    localStorage.setItem('username', username.value)
    location.reload()
  }
}
</script>

<template>
  <Toast />
  <div class="h-[100vh] flex justify-center items-center">
    <Card style="width: 25rem">
      <template #title>Login</template>
      <template #content>
        <form class="flex flex-col gap-2">
          <label for="username">Username: </label>
          <InputText
            type="text"
            id="username"
            required
            placeholder="Enter username"
            v-model="username"
          />

          <label for="password">Password: </label>
          <InputText
            type="password"
            id="password"
            required
            placeholder="Enter password"
            v-model="password"
          />
        </form>
      </template>
      <template #footer>
        <Button label="Login" @click="login" />
      </template>
    </Card>
  </div>
</template>
