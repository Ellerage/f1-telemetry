<script setup>
    const navItems = [
        {text: "Home", url: "/"},
        {text: "List laps", url: "/0"},
        {text: "Config", url: "/config"}
    ]
</script>

<template>
    <nav class="container">
        <ui class="list">
            <li class="list-item" v-for="item in navItems"><router-link :to="item.url">{{ item.text }}</router-link></li>
        </ui>
    </nav>
</template>

<style scoped>
    .container {
        width: 100vw;
        position: fixed;
        top: 0;
        left: 0;
    }

    .list {
        display: flex;
        flex-direction: row;
        padding: 10px 20px;
        list-style-type: none;
    }

    .list-item {
        padding-right: 20px;
    }
</style>