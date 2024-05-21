<script setup>
import { ref } from 'vue'
import ApiService from '@/api'

import Toast from 'primevue/toast'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Dropdown from 'primevue/dropdown'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const showSuccessMessage = () => {
  toast.add({
    severity: 'success',
    summary: 'Success',
    detail: 'User successfully added',
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
const password = ref(null)
const role = ref(null)

const roles = ref([{ name: 'manager' }, { name: 'admin' }])

const firstName = ref(null)
const lastName = ref(null)
const stationName = ref(null)

const addUser = async () => {
  try {
    await ApiService.register(
      username.value,
      password.value,
      role.value,
      firstName.value,
      lastName.value,
      stationName.value,
      localStorage.getItem('token')
    )
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
      <template #title>Add user</template>
      <template #content>
        <div class="flex flex-wrap gap-2">
          <InputText type="text" placeholder="Enter username" v-model="username" />
          <InputText type="password" placeholder="Enter password" v-model="password" />
          <!-- <Dropdown
            v-model="role"
            :options="roles"
            optionLabel="name"
            placeholder="Select a Role"
          /> -->
          <InputText type="text" placeholder="Enter role" v-model="role" />
          <InputText type="text" placeholder="Enter firstname" v-model="firstName" />
          <InputText type="text" placeholder="Enter lastname" v-model="lastName" />
          <InputText type="text" placeholder="Enter station name" v-model="stationName" />
        </div>
      </template>
      <template #footer>
        <Button icon="pi pi-user-plus" label="Add" @click="addUser" />
      </template>
    </Card>
  </div>
</template>
