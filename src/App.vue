<template>
    <div id="app">
        <Index msg="db haha"/>
        ${ dd }
    </div>
</template>

<script>
    import Index from './components/index.vue'

    export default {
        name: 'App',
        components: {
            Index
        },
        data() {
            return {
                dd: "dsfdsf"
            }
        }
    }
</script>

<style>
    #app {
        font-family: Avenir, Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
    }
</style>
