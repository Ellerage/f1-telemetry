<template>
    <div v-if="state" class="container">
        <div>
            <h3 class="title">
                App
            </h3>
            <Input label="App port" id="port" type="number" v-model.number="state.Port"></Input>
        </div>
        <div>
            <h3 class="title">Files</h3>
            <div class="input-group">
                 <Input label="Telemetry file" id="telemetry-file" v-model="state.TelemetryFileName"></Input>
                <Input label="Laps file" id="laps-file" v-model="state.LapsFileName"></Input>

                <Input label="TelemetryFMBufferRows" id="telemetry-fm-buffer-rows" type="number" v-model="state.TelemetryFMBufferRows"></Input>
            </div>
        </div>
        <div>
            <h3 class="input-group">OBS</h3>

            <div class="input-group">
                <Input label="Use obs" id="use-obs" type="number" v-model.number="state.UseObs"></Input>
                <Input v-if="state.UseObs" label="Obs buffer sec" id="obs-buffer" type="number" v-model.number="state.ObsBufferSeconds"></Input>
                <Input v-if="state.UseObs" label="Obs password" id="obs-password" v-model="state.ObsPassword"></Input>
                <Input v-if="state.UseObs" label="Obs port" id="obs-port" type="number" v-model.number="state.ObsPort"></Input>
                <Input v-if="state.UseObs" label="Obs addr" id="obs-addr" v-model="state.ObsAddr"></Input>
            </div>
        </div>
    </div>

    <div class="button-container">
        <Button class="button" @click="save">Save</Button>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import type { Config } from '../utils/types/config';
import Input from "../components/Input.vue"

const state = ref<Config | null>(null)

onMounted(async () => {
    const data = await fetch("/api/config")
    const config = await data.json() as Config
    state.value = config
})

async function save() {
    const response = await fetch("/api/config", {
        method: "POST", 
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(state.value)
    })

     if (response.ok) {
        alert("Saved!")
    }
}

</script>

<style scoped>
    .container {
        display: flex;
        flex-direction: column;
        gap: 15px;
    }

    .button-container {
        margin-top: 20px;
        display: flex;
        justify-content: center;
    }

    .button {
        width: 200px;
        height: 44px;
    }

    .input-group {
        display: flex;
        flex-wrap: wrap;
        gap: 14px;
    }
</style>