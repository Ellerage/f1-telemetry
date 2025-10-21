<template>
    <button @click="emit('submit')">
        {{ text }}
    </button>
</template>

<script setup>
    const props = defineProps(['text'])
    const emit = defineEmits(['submit'])
</script>