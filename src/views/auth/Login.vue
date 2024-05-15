<script setup>
import Card from 'primevue/card'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'

import { RouterLink } from 'vue-router'
</script>

<template>
  <div class="h-[100vh] flex justify-center items-center">
    <Card style="width: 25rem">
      <template #title>Login</template>
      <template #content>
        <form class="flex flex-col gap-2">
          <label for="username">Username: </label>
          <InputText type="text" id="username" required placeholder="Enter username" />

          <label for="password">Password: </label>
          <InputText type="password" id="password" required placeholder="Enter password" />
        </form>
      </template>
      <template #footer>
        <div class="flex justify-between items-center">
          <p class="text-sm">
            Don't have an account?
            <RouterLink to="/signup" class="text-blue-500 underline">Sign Up</RouterLink>
          </p>
          <Button label="Sign Up" />
        </div>
      </template>
    </Card>
  </div>
</template>
