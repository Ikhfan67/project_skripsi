<template>
    <div class="list">
      <div id="menu" class="menu-main pad-top-100 pad-bottom-100">
        <div class="container">
            <div class="">
                <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
                    <div class="wow fadeIn" data-wow-duration="1s" data-wow-delay="0.1s">
                        <h2 class="block-title text-center">Our Menu</h2>
                        <div class="diet">
                            <p class="">"I am on a seafood diet. I see food and I eat it." </p>
                        </div>
                    </div>
                    <div class="tab-menu">
                        <div class="slider slider-nav">
                            <div class="tab-title-menu">
                                <button type="button" class="btn btn-outline-warning animated infinite pulse delay-4s" :class="{ active: filter == 'all' }" @click="filter = 'all'">All</button>
                                <p><i class="fas fa-utensils"></i></p>
                            </div>
                            <div class="tab-title-menu">          
                                <button type="button" class="btn btn-outline-warning animated infinite pulse delay-4s" :class="{ active: filter == 'active' }" @click="filter = 'active'">MAIN DISHES</button>
                                <p><i class="fas fa-hamburger"></i></p>
                            </div>
                            <div class="tab-title-menu">                               
                                <button type="button" class="btn btn-outline-warning animated infinite pulse delay-4s" :class="{ active: filter == 'completed' }" @click="filter = 'completed'">DRINKS</button>
                                <p><i class="fas fa-coffee"></i> </p>
                            </div>
                        </div>
                        <transition-group class="row" name="fade" enter-active-class="animated jackInTheBox" leave-active-class="animated zoomOutUp">
                            <div class="col-lg-6 col-md-6 col-sm-12 col-xs-12" v-for="todo in todosFiltered" :key="todo.id">
                                <div class="offer-item">
                                    <img :src="todo.src" alt="" class="img-responsive">
                                    <div >
                                        <h3>{{todo.name}}</h3>
                                    </div>
                                    <span class="offer-price">{{todo.harga}}</span>
                                </div>
                            </div>
                        </transition-group>
                    </div>
                </div>
            </div>
        </div>
      </div>
    </div>
</template>
<script>
import TodoFiltered from './todoFiltered'
export default {
    name:'todo-list',
    data(){
        return{
            filter: 'all',
            todos:[
                {
                    id:1,
                    completed:false,
                    name:'GARLIC BREAD',
                    harga:'15K',
                    src:require('./images/menu-item-thumbnail-01.jpg')
                },
                {
                    id:2,
                    completed:false,
                    name:'MIXED SALAD',
                    harga:'16K',
                    src:require('./images/menu-item-thumbnail-02.jpg')
                },
                {
                    id:3,
                    completed:false,
                    name:'BBQ CHICKEN WINGS',
                    harga:'20K',
                    src:require('./images/menu-item-thumbnail-03.jpg')
                },
                {
                    id:4,
                    completed:true,
                    name:'BBQ BEFF WINGS',
                    harga:'35K',
                    src:require('./images/menu-item-thumbnail-03.jpg')
                },
            ]
        }
    },
    components:{
        TodoFiltered
    },
    computed:{
        todosFiltered(){
            if(this.filter == 'all'){
                return this.todos
            }else if(this.filter == 'active'){
                return this.todos.filter(todos => !todos.completed)
            }else if(this.filter == 'completed'){
                return this.todos.filter(todos => todos.completed)                
            }
            return this.todos
        }
    }
}
</script>
<style lang="scss">
@import url("https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.7.2/animate.min.css");
.fade-enter-active, .fade-leave-active{
    transition: opacity .2s;
}
.fade-enter, .fade-leave-to{
    opacity: 0 ;
}
</style>


