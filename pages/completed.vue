<template>
   <div class="w-full h-full inter text-[#eaeaea]">
      <div class="header w-full h-fit flex flex-col md:flex-row items-center p-8">
         <NuxtLink to="/" class="self-center md:self-start w-2/6">
            <Icon name="material-symbols:arrow-right-alt-rounded" class="cursor-pointer rotate-180 bg-[#eaeaea] mt-4 md:mt-0" size="4rem"/>
         </NuxtLink>
         <h1 class="font-bold text-4xl md:text-7xl self-center md:text-left">COMPLETED</h1>
      </div>
 
      <div class="w-full h-fit flex flex-col items-center gap-4 px-4 md:px-0">
         <Project 
            v-for="project in projects" 
            :data="project" 
            @reload="getAll" 
         />
      </div>

   </div>

</template>


<style scoped>

    .inter {
        font-family: Inter, sans-serif;
    }

</style>


<script setup>
   import Project from './components/Project.vue';

   const { server_url } = useRuntimeConfig().public
   const projects = ref([])

   async function getAll() {
      const result = await $fetch(`${server_url}/all`)

      projects.value = result
   }

   onMounted(async () => {

      await getAll()


   })
</script>