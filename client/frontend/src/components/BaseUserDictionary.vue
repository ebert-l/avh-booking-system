<template>
  <div class="user-list">
    <div
      class="columns"
      v-for="(users, key) in userDict"
      :key="key"
      style="margin-bottom: 0cm"
    >
      <div class="column is-1">
        <h3 class="subtitle is-3 has-text-grey-light has-text-weight-bold">
          {{ key }}
        </h3>
      </div>
      <div class="column">
          <!-- :key="user" -->
        <div class="buttons">
          <button
            class="button"
            v-for="user in users"
            :class="buttonColor(selectedUser, user)"
            @click="selectUser(user)">
            {{ displayUserName(user) }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { buttonColor, selectUser, displayUserName } from "../helper"
import User from "../models/User"
import UserDict from "../models/UserDict"
import { useUserStore } from '../stores/userStore'
const userStore = useUserStore()

export default defineComponent({
    data() {
      return {
          userDict:{} as UserDict
      }
    },
    computed: {
      selectedUser(): User {
          return userStore.getSelectedUser
      }
    },
    methods: {
        displayUserName,
        buttonColor,
        selectUser
    }
});
</script>