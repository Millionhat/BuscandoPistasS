<template>
    <v-container>
        <v-row>
        <v-text-field v-model="sitio" label="Que sitio deseas buscar"></v-text-field>
        <v-btn color='#a83cde' dark @click="buscarInfo"> Buscar </v-btn>
        </v-row>
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
            contenido:[],
            sitio:'',
            link:'',
            c:false,

        }
    },
    methods:{
        buscarInfo:function(){
        axios.get('http://localhost:8087/nuevo/'+this.sitio).then(response =>{
            console.log(response.data);
            this.contenido= response.data;
            this.c=true;
        },error =>{
            console.log(error);
        });
        }
    }
}

</script>