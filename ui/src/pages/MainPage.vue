<script setup lang="ts">
import { TracksList } from '../utils/enum/track-ids-map';

</script>

<template>
    <h1>Tracks</h1>
    <div>
        <ul class="list">
            <li v-for="[index, name] in TracksList" class="list-item"><router-link class="link" :to="index">{{ name }}</router-link></li>
        </ul>
    </div>
</template>

<style scoped>
    .list {
        padding-inline-start: 0;
        column-count: 2;      
        column-gap: 16px;
    }
    .list-item {
        list-style-type: none;
        padding: 10px 15px;
        border: 1px solid white;
        break-inside: avoid;   /* чтобы элементы не разрывались */
        margin-bottom: 8px;
        background: grey;
        color: white;
    }

    .link {
        color: white;
    }
</style>
