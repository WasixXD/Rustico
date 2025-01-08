<template>
    <div class="flex flex-row w-full h-full inter">
        <div class="h-screen w-1/3 flex justify-center items-start">
            <NuxtLink to="/">

                <Icon name="material-symbols:arrow-right-alt-rounded" class="cursor-pointer rotate-180 mt-16 bg-[#eaeaea]" size="4rem"/>
            </NuxtLink>
        </div>
        <div class="w-full h-full flex justify-center items-start">
            <form class="flex flex-col justify-start items-center h-full w-full mt-16" @submit.prevent="handleSubmit" >
                <input v-model="name" type="text" placeholder="Project Name" name="project_name" required
                    class="border-2 border-[#eaeaea] w-2/4 h-16 bg-[#282828] text-white placeholder:text-[#eaeaea] p-4 my-4"
                >
                <textarea v-model="desc" name="description" placeholder="Project Description"
                    class="border-2 border-[#eaeaea] w-2/4 bg-[#282828] text-white placeholder:text-[#eaeaea] p-4 my-4" rows="20" cols="80"
                ></textarea>
                <button type="submit" class="bg-white w-2/4 p-4 font-black text-[#282828]">ADD</button>
            </form>
        </div>
    </div>
</template>


<style scoped>

    .inter {
        font-family: Inter, sans-serif;
    }

    input, textarea {
        outline: none;
    }
</style>


<script setup>
    const name = ref("")
    const desc = ref("")


    async function handleSubmit() {
        const url = "/api/create_project"

        try {
            const _ = await $fetch(url, {
                    method: "POST",
                    body: {
                        project_name: name.value,
                        description: desc.value
                    },
                },
            )
            navigateTo("/")

        } catch(e) {
            alert(e)
        }


    }
</script>
