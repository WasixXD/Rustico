<template>
    <div class="flex flex-col w-full h-full text-[#eaeaea]">

        <div class="first flex flex-row w-100 h-1/4 text-4xl inter font-black">
            <div class="w-full h-full flex items-center">
                <NuxtLink to="/" class="underline ml-36">Completed</NuxtLink>
            </div>

            <div class="w-full h-full">

            </div>

            <div class="w-full h-full flex items-center justify-end">
                <NuxtLink to="/create/project" class="px-6 py-4 border-2 border-[#eaeaea] mr-36">Add</NuxtLink>
            </div>
        </div>

        <div class="second flex flex-col w-100 h-3/5 inter">
            <div class="title pt-16 pb-12">
                <h1 class="font-black text-7xl ml-40">{{ title }}</h1>
            </div>
            <div class="texto ml-80 w-3/4 text-justify border-[#eaeaea] border-left border-l-2 h-fit">
                <p class="text-xl pl-4">
                    {{ description }}
                </p>
            </div>

        </div>

        <div class="w-full h-100 flex flex-row justify-center">
            <div class="icons flex flex-row justify-center items-center w-52">
                <Icon name="tabler:dice-5-filled" size="4rem" class="cursor-pointer mx-4" @click="randomPost" />
                <Icon name="pajamas:task-done" size="4rem" class="cursor-pointer mx-4" @click="projectDone" />
            </div>
        </div> 
    </div>
</template>

<style scoped>

    .inter {
        font-family: Inter, sans-serif;
    }
</style>


<script setup>

    const title = ref("")
    const description = ref("")


    const { server_url } = useRuntimeConfig().public

    async function randomPost() {
        const { project_name, description: desc } = await $fetch(`${server_url}/random`)
        title.value = project_name
        description.value = desc
    }

    async function projectDone() {
        console.log("projectDone")
    }


    onMounted(async () => {
        await randomPost()
    })
</script>