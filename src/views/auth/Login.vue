<script setup>
import { ref } from 'vue'
import Card from 'primevue/card'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'

import { RouterLink } from 'vue-router'

import ApiService from '@/api'

const username = ref(null)
const password = ref(null)
const data = ref(null)

const login = async () => {
  if (username.value && password.value) {
    try {
      data.value = await ApiService.login(username.value, password.value)
    } catch (error) {
      console.log(error)
    } finally {
      localStorage.setItem('token', data.value.data)
      localStorage.setItem('username', username.value)
      location.reload()
    }
  } else {
    console.log('Username or password cannot be empty')
  }
}
</script>

<template>
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
        <div class="flex justify-between items-center">
          <p class="text-sm">
            Don't have an account?
            <RouterLink to="/signup" class="text-blue-500 underline">Sign Up</RouterLink>
          </p>
          <Button label="Login" @click="login" />
        </div>
      </template>
    </Card>
  </div>
</template>
