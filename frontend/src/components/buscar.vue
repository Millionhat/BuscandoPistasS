<template>
    <v-container>
        <v-row>
        <v-text-field v-model="sitio" label="Que sitio deseas buscar"></v-text-field>
        <v-btn color="accent" @click="buscarInfo"> Buscar </v-btn>
        </v-row>
        <div>
            {{contenido}}
        </div>
    </v-container>
</template>

<script>
import axios from 'axios';

export default {
    data(){
        return{
            contenido:[],
            sitio:'',
            link:'',

        }
    },
    methods:{
        buscarInfo:function(){
        axios.get('http://localhost:8087/nuevo/'+this.sitio).then(response =>{
            console.log(response.data);
            this.contenido= response.data;
        },error =>{
            console.log(error);
        });
        }
    }
}

</script>