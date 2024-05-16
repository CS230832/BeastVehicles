<script setup>
import { ref } from 'vue'
import Card from 'primevue/card'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'

import { RouterLink } from 'vue-router'

import ApiService from '@/api'

const email = ref(null)
const password = ref(null)
const data = ref(null)

const login = async () => {
  if (email.value && password.value) {
    try {
      data.value = await ApiService.login(email.value, password.value)
      console.log(data.value)
    } catch (error) {
      console.log(error)
    } finally {
      localStorage.setItem('token', data.value.token)
      location.reload()
    }
  } else {
    console.log('Email or password cannot be empty')
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
            v-model="email"
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
