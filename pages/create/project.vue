<template>
    <div v-show="showAlert">
        Alerta!
    </div>

    <div class="w-full h-full bg-black">
        <form class="flex flex-col" @submit.prevent="handleSubmit" >
            <input v-model="name" type="text" placeholder="Project Name" name="project_name" required>
            <textarea v-model="desc" name="description" placeholder="Project Description"></textarea>
            <button type="submit">Create</button>
        </form>
    </div>
</template>


<script setup>
    const name = ref("")
    const desc = ref("")

    const showAlert = ref(false)

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
