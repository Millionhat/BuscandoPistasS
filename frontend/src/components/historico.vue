<template>
    <v-container>
        <v-select v-model="contenido"
        :items="sitios"
        name="sitio"
        item-value="Datos"
        item-text="Nombre"
        menu-props="auto"
        label="Elegir un sitio"
        v-on:change="getData()"
        hide-details
        single-line/>
        <div v-show="c">
            {{contenido}}
        </div>
    </v-container>
</template>

<script>
import axios from 'axios';

export default {
    data(){
        return{
            sitios:[],
            contenido:[],
            selectedHost:"",
            lista:[],
            valores:[],
            informacion:"",
            hShow: false,
            c: false,
        }
    },
    methods:{
    getSitios: function(){
        axios.get('http://localhost:8087/historial').then(response =>{
            console.log(response.data);
            this.sitios = response.data;
        },error =>{
            console.log(error);
        });
    },
    getData: function(){
        this.valores= Object.values(this.contenido)
        this.informacion= this.valores[1]
        this.c=true
    }
    },
    created(){
    this.getSitios();
    }
}

</script>
