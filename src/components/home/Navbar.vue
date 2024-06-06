<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import ApiService from '@/api'
import checkIfUserIsAuthenticated from '@/views/auth/checkAuth'
import checkIfUserIsRoot from '@/views/auth/checkRoot'
import checkIfUserIsManager from '@/views/auth/checkManager'

import Toast from 'primevue/toast'
import Menubar from 'primevue/menubar'
import Button from 'primevue/button'
import Menu from 'primevue/menu'
import Badge from 'primevue/badge'
import Avatar from 'primevue/avatar'
import { useToast } from 'primevue/usetoast'

const toast = useToast()
const router = useRouter()

const errorMessage = ref(null)
const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: errorMessage.value,
    life: 3000
  })
}

const logout = async () => {
  if (checkIfUserIsAuthenticated()) {
    try {
      await ApiService.logout(localStorage.getItem('token'))
    } catch (error) {
      errorMessage.value = error.response.data.data
      showErrorMessage()
    } finally {
      localStorage.removeItem('token')
      localStorage.removeItem('username')
      location.reload()
    }
  } else {
    errorMessage.value = 'Something went wrong'
    showErrorMessage()
  }
}

const navbarItems = ref([
  {
    label: 'Home',
    icon: 'pi pi-home',
    command: () => {
      router.push('/home')
    }
  },

  {
    label: 'Find Vehicle',
    icon: 'pi pi-search',
    command: () => {
      router.push('/search')
    }
  }
])

const menu = ref()
const menuItems = ref([
  {
    label: 'Profile',
    items: [
      {
        label: 'Logout',
        icon: 'pi pi-sign-out',
        command: () => {
          logout()
        }
      }
    ]
  }
])

const toggle = (event) => {
  menu.value.toggle(event)
}

const username = ref(null)
const role = ref(null)

const getUserDetails = async () => {
  try {
    const response = await ApiService.getUser(
      localStorage.getItem('username'),
      localStorage.getItem('token')
    )

    username.value = response.data.username
    role.value = response.data.role
  } catch (error) {
    console.log(`Error getting user details: ${error.response.data.data}`)
  }
}

onMounted(async () => {
  const isAuth = checkIfUserIsAuthenticated()
  const isRoot = await checkIfUserIsRoot()
  const isManager = await checkIfUserIsManager()

  if (isAuth) {
    navbarItems.value.push(
      {
        label: 'New Vehicle',
        icon: 'pi pi-car',
        command: () => {
          router.push('/add')
        }
      },

      {
        label: 'Remove Vehicle',
        icon: 'pi pi-trash',
        command: () => {
          router.push('/remove')
        }
      },

      {
        label: 'Free Slots',
        icon: 'pi pi-circle',
        command: () => {
          router.push('/free')
        }
      },

      {
        label: 'Full Slots',
        icon: 'pi pi-circle-fill',
        command: () => {
          router.push('/full')
        }
      }
    )
  }

  if (isRoot) {
    navbarItems.value.push(
      {
        label: 'New Station',
        icon: 'pi pi-plus',
        command: () => {
          router.push('/add-station')
        }
      },
      {
        label: 'Remove Station',
        icon: 'pi pi-minus-circle',
        command: () => {
          router.push('/remove-station')
        }
      },
      {
        label: 'Add User',
        icon: 'pi pi-user-plus',
        command: () => {
          router.push('/add-user')
        }
      },
      {
        label: 'Remove User',
        icon: 'pi pi-user-minus',
        command: () => {
          router.push('/remove-user')
        }
      }
    )
  } else if (isManager) {
    navbarItems.value.push(
      {
        label: 'Add User',
        icon: 'pi pi-user-plus',
        command: () => {
          router.push('/add-user')
        }
      },
      {
        label: 'Remove User',
        icon: 'pi pi-user-minus',
        command: () => {
          router.push('/remove-user')
        }
      }
    )
  }

  if (isAuth) {
    await getUserDetails()
  }
})
</script>

<template>
  <Toast />
  <Menubar :model="navbarItems">
    <template #end v-if="checkIfUserIsAuthenticated()">
      <Button
        type="button"
        icon="pi pi-user"
        rounded
        @click="toggle"
        aria-haspopup="true"
        aria-controls="overlay_menu"
      />
      <Menu ref="menu" id="overlay_menu" :model="menuItems" :popup="true">
        <template #submenuheader="{ item }">
          <span class="text-primary font-bold">{{ item.label }}</span>
        </template>
        <template #item="{ item, props }">
          <a class="flex items-center" v-bind="props.action">
            <span :class="item.icon" />
            <span class="ml-2">{{ item.label }}</span>
            <Badge v-if="item.badge" class="ml-auto" :value="item.badge" />
            <span
              v-if="item.shortcut"
              class="ml-auto border border-surface-200 dark:border-surface-700 rounded-md bg-surface-100 dark:bg-surface-700 text-xs p-1"
              >{{ item.shortcut }}</span
            >
          </a>
        </template>
        <template #end>
          <div class="w-full flex items-center p-2 pl-3 text-surface-700 dark:text-surface-0/80">
            <Avatar image="/user.png" class="mr-2" shape="circle" />
            <span class="inline-flex flex-col justify-start">
              <span class="font-bold">{{ username }}</span>
              <span class="text-sm flex items-center gap-1"
                >Role:<span class="font-bold">{{ role }}</span
                ><span class="pi pi-verified"></span
              ></span>
            </span>
          </div>
        </template>
      </Menu>
    </template>

    <template #end v-else>
      <Button
        type="button"
        label="Login"
        @click="
          () => {
            router.push('/login')
          }
        "
      />
    </template>
  </Menubar>
</template>
