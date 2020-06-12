<template>
  <div>
    <v-app-bar class="mx-auto" color='#a83cde' grow dark>Proyecto Buscando Pistas S</v-app-bar>
    <v-container class="contenedor">
     <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb3 text-center">
          Bienvenido que deseas hacer
        </h1>
        <h3 class="text-center titulo">
          Deseas buscar la informacion de un dominio o las busquedas anteriores?
        </h3>
    </v-col>
    <v-col class="mb-4 text-center subtitulo">
          <v-btn color='#a83cde' dark outlined v-on:click="bShow = !bShow, hShow=false">Busqueda de Info</v-btn>
          <v-btn color='#a83cde' dark outlined v-on:click="hShow = !hShow, bShow= false">Busquedas Realizadas</v-btn>
     </v-col>
      <buscar v-show="bShow"/>
      <historico v-show="hShow"/>
    </v-container>
  </div>
</template>

<script>
// @ is an alias to /src
import buscar from "@/components/buscar.vue"
import historico from "@/components/historico.vue"
export default {
  data(){
    return{
      bShow: false,
      hShow: false,
    }
  },
  name: 'Home',
  components:{
    buscar,
    historico,
  }

}
</script>

<style scoped>
.contenedor{
  margin-top: 80px;
}
.titulo{
  margin-top: 20px;
}
.subtitulo{
  scroll-margin-bottom: 20px;
}
</style>