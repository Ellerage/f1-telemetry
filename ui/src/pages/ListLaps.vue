<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router';
import type { Lap } from '../utils/types/lap';
import { TrackIDMap } from '../utils/enum/track-ids-map';
import { sessionTypeMap } from '../utils/enum/session-type';
import { getLapTime, getSectorString } from '../utils/func/lap';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

interface TableRow {
    track: string
    lapNum: number
    sector1: string
    sector2: string
    sector3: string
    total: string
    sessionType: string
    sessionUID: number
}


const route = useRoute();

const laps = ref<TableRow[]>([])

onMounted(async () => {
  const res = await fetch(`/api/laps?track_id=${route.params.trackId}`)
  const result: Lap[] = await res.json()

  laps.value = result.map((item) => {
    return {
        track: TrackIDMap[item.TrackId] || "",
        lapNum: item.CurrentLapNum,
        sector1: getSectorString(item.Sector1Minutes, item.Sector1MS),
        sector2: getSectorString(item.Sector2Minutes, item.Sector2MS),
        sector3: getSectorString(item.Sector3Minutes, item.Sector3MS),
        total: getLapTime(item.Total),
        sessionType: sessionTypeMap[item.SessionType] || "",
        sessionUID: item.SessionUID
    }
  })
})
</script>

<template>
    <div v-show="laps.length > 0">
        <DataTable :value="laps" :tableStyle="{ minWidth: '600px' }" style="width: 100%;"   >
            <Column field="track" header="Track"></Column>
            <Column field="lapNum" header="Lap num"></Column>
            <Column field="sector1" header="Sector 1"></Column>
            <Column field="sector2" header="Sector 2"></Column>
            <Column field="sector3" header="Sector 3"></Column>
            <Column field="total" header="Total"></Column>
            <Column field="sessionType" header="SessionType"></Column>
            <Column field="sessionUID" header="SessionUID"></Column>
        </DataTable>
    </div>

    <div v-show="laps.length === 0">
        <p>No results for this track</p>
    </div>
</template>

<style scoped>
    .list-item-container {
        list-style-type: none;
    }

    .column-item {
        padding: 7px;
        display: inline-block;
    }
</style>
