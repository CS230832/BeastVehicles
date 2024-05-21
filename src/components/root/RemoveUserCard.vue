<script setup>
import { ref } from 'vue'
import ApiService from '@/api'

import Toast from 'primevue/toast'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const showSuccessMessage = () => {
  toast.add({
    severity: 'success',
    summary: 'Success',
    detail: 'User successfully removed',
    life: 3000
  })
}

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

const removeUser = async () => {
  try {
    await ApiService.removeUser(username.value, localStorage.getItem('token'))
    showSuccessMessage()
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}
</script>

<template>
  <Toast />
  <div class="p-4">
    <Card>
      <template #title>Remove a user</template>
      <template #content>
        <InputText type="text" placeholder="Enter username" v-model="username" />
      </template>
      <template #footer>
        <Button icon="pi pi-user-minus" label="Remove" severity="danger" @click="removeUser" />
      </template>
    </Card>
  </div>
</template>
