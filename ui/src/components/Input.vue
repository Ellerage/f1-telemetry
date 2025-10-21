<template>
    <div style="margin-bottom: 1rem; max-width: 200px;">
        <label :for="id" style="display: block; margin-bottom: 0.5rem; font-weight: 500;">
            {{ label }}
        </label>
          <input 
            :id="id"
            :value="modelValue"
            :type="type === 'number' ? 'number' : 'text'"
            @input="handleInput"
            style="
                width: 100%; 
                padding: 0.75rem; 
                border: 2px solid #ddd; 
                border-radius: 6px; 
                font-size: 1rem;
                box-sizing: border-box;
            "
        />
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    label: string
    id: string
    modelValue: string | number
    type?: 'text' | 'number'
}>()
const emit = defineEmits<{
    'update:modelValue': [value: string | number]
}>()

const handleInput = (event: Event) => {
    const target = event.target as HTMLInputElement

     if (props.type === 'number') {
        const value = parseInt(target.value, 10)
        emit('update:modelValue', isNaN(value) ? 0 : value)
    } else {
        emit('update:modelValue', target.value)
    }
}
</script>