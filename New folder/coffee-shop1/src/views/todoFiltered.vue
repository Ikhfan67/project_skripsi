<template>
    <div class="slider slider-nav">
        <div class="tab-title-menu">
            <button type="button" class="btn btn-outline-warning" :class="{ active: filter == 'all' }" @click="filter = 'all'">All</button>
            <p><i class="fas fa-utensils"></i></p>
        </div>
        <div class="tab-title-menu">          
            <button type="button" class="btn btn-outline-warning" :class="{ active: filter == 'active' }" @click="filter = 'active'">MAIN DISHES</button>
            <p><i class="fas fa-hamburger"></i></p>
        </div>
        <div class="tab-title-menu">                               
            <button type="button" class="btn btn-outline-warning" :class="{ active: filter == 'completed' }" @click="filter = 'completed'">DRINKS</button>
            <p><i class="fas fa-coffee"></i> </p>
        </div>
    </div>
</template>

<script>
export default {
    name: 'todo-filtered',
    computed:{
        filter(){
            return this.$store.state.filter
        }
    },
    methods:{
        changeFilter(filter){
            this.$store.dispatch('updateFilter', filter)
        }
    }
}
</script>

