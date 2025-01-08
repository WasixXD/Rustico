<template>
    <div class="flex flex-col w-full h-full text-[#eaeaea]">

        <div class="first flex flex-row w-100 h-1/4 text-4xl inter font-black">
            <div class="w-full h-full flex items-center">
                <NuxtLink to="/completed" class="underline ml-36">Completed</NuxtLink>
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
                <Icon name="pajamas:task-done" size="4rem" class="cursor-pointer mx-4" @click="toggleModal" />
            </div>
        </div> 
    </div>
    
    <dialog :class="{ 'modal-open': showModal }" class="modal inter text-[#eaeaea]">
        <div class="modal-box w-11/12 max-w-2xl bg-[#282828]">
            <h3 class="text-lg font-bold">Complete Project?</h3>
            <input class="border-2 border-[#eaeaea] w-2/4 h-16 bg-[#282828] text-white placeholder:text-[#eaeaea] p-4 my-4" required type="text" placeholder="http://your.url.here" v-model="githubUrl">
            <div class="modal-action">
                <button class="btn border-none bg-[#282828] hover:bg-[#282828] hover:border-1 hover:border-[#eaeaea]" @click="toggleModal">Close</button>
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
        const { id: _id, project_name, description: desc } = await $fetch(`${server_url}/random`)
        title.value = project_name
        description.value = desc
        id.value = _id
    }

    function toggleModal() {
        showModal.value = !showModal.value
    }

    async function projectDone() {
        if(!githubUrl.value) return
        toggleModal()

        try {
            const response = await $fetch(`${server_url}/complete?id=${id.value}&url=${githubUrl.value}`)
            console.log(response)

        } catch(e) {
            console.log(e)
        }
    }


    onMounted(async () => {
        await randomPost()
    })
</script>