<template>
    <div class="flex flex-col w-full h-full text-[#eaeaea]">

        <div class="first flex flex-col xl:flex-row w-full h-auto xl:h-1/4 text-4xl inter font-black">
            <div class="w-full h-auto xl:h-full flex items-center justify-center xl:justify-start">
                <NuxtLink to="/completed" class="underline xl:ml-36">Completed</NuxtLink>
            </div>

            <div class="w-full h-auto xl:h-full"></div>

            <div class="w-full h-auto xl:h-full flex items-center justify-center xl:justify-end mt-4 xl:mt-0">
                <NuxtLink to="/create/project" class="px-6 py-4 border-2 border-[#eaeaea] xl:mr-36">Add</NuxtLink>
            </div>
        </div>

        <div class="second flex flex-col w-full h-auto xl:h-3/5 inter">
            <div class="title pt-8 xl:pt-16 pb-6 xl:pb-12">
                <h1 class="font-black text-4xl xl:text-7xl ml-8 xl:ml-40">{{ title }}</h1>
            </div>
            <div class="texto ml-8 xl:ml-80 w-[90%] xl:w-3/4 text-justify border-[#eaeaea] border-l-0 xl:border-l-2 h-fit">
                <p class="text-lg xl:text-xl pl-0 xl:pl-4">
                    {{ description }}
                </p>
            </div>
        </div>

        <div class="w-full h-auto flex flex-row justify-center mt-8 xl:mt-0">
            <div class="icons flex flex-row justify-center items-center w-full xl:w-52">
                <Icon name="tabler:dice-5-filled" size="3rem" class="cursor-pointer mx-2 xl:mx-4" @click="randomPost" />
                <Icon name="pajamas:task-done" size="3rem" class="cursor-pointer mx-2 xl:mx-4" @click="toggleModal" />
            </div>
        </div> 
    </div>
    
    <dialog :class="{ 'modal-open': showModal }" class="modal inter text-[#eaeaea]">
        <div class="modal-box w-11/12 max-w-2xl bg-[#282828]">
            <h3 class="text-lg font-bold">Complete Project?</h3>
            <input class="border-2 border-[#eaeaea] w-full xl:w-2/4 h-12 xl:h-16 bg-[#282828] text-white placeholder:text-[#eaeaea] p-4 my-4" required type="text" placeholder="http://your.url.here" v-model="githubUrl">
            <div class="modal-action flex flex-col sm:flex-row gap-4">
                <button class="btn border-none bg-[#282828] hover:bg-[#282828] hover:border-[#eaeaea]" @click="toggleModal">Close</button>
                <button class="btn border-none hover:bg-[#eaeaea] bg-[#eaeaea] text-[#282828]" @click="projectDone">Complete</button>
            </div>
        </div>
    </dialog>
</template>


<style scoped>

    .inter {
        font-family: Inter, sans-serif;
    }
</style>


<script setup>

    const title = ref("")
    const description = ref("")
    const id = ref(-1)
    const showModal = ref(false)
    const githubUrl = ref("")


    const { server_url } = useRuntimeConfig().public

    async function randomPost() {
        try {
            const { id: _id, project_name, description: desc } = await $fetch(`${server_url}/random`)
            title.value = project_name
            description.value = desc
            id.value = _id

        } catch(e) {
            navigateTo("/create/project")
        }
    }

    function toggleModal() {
        showModal.value = !showModal.value
    }

    async function projectDone() {
        if(!githubUrl.value) return
        toggleModal()

        try {
            const _ = await $fetch(`${server_url}/complete?id=${id.value}&url=${githubUrl.value}`)
            randomPost()
        } catch(e) {
            console.log(e)
        }
    }


    onMounted(async () => {
        await randomPost()
    })
</script>