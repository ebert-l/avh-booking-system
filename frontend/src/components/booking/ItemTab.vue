<template>
  <div>
    <ItemSearch :mode="'multiple'" />
    <ItemList />
  </div>
</template>

<script lang="ts">
import ItemSearch from "./ItemSearch.vue";
import ItemList from "./ItemList.vue";

export default {
  components: {
    ItemSearch,
    ItemList,
  },
};
</script>