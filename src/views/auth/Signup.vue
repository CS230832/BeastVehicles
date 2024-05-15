<template>
  <div class="h-[100vh] flex justify-center items-center">
    <Card style="width: 25rem">
      <template #title>Sign Up</template>
      <template #content>
        <form class="flex flex-col gap-2">
          <label for="name">Full name: </label>
          <InputText type="text" id="name" required placeholder="Enter your full name" />

          <label for="username">Username: </label>
          <InputText type="text" id="username" required placeholder="Create username" />

          <label for="password">Password: </label>
          <InputText type="password" id="password" required placeholder="Create password" />
        </form>
      </template>
      <template #footer>
        <div class="flex justify-between items-center">
          <p class="text-sm">
            Already have an account?
            <RouterLink to="/login" class="text-blue-500 underline">Login</RouterLink>
          </p>
          <Button label="Sign Up" />
        </div>
      </template>
    </Card>
  </div>
</template>

<script setup>
import Card from 'primevue/card'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'

import { RouterLink } from 'vue-router'
</script>
