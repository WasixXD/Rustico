<template>
    <div 
        class="border-4 border-[#eaeaea] p-8 flex flex-col md:flex-row justify-evenly items-center md:items-center m-4 w-full md:w-3/4"
    >
        <div class="flex flex-col md:flex-row justify-evenly items-center md:items-center gap-4 md:gap-8 w-full">
            <p>{{ data.id }}.</p>
            <p class="font-semibold">{{ data.project_name }}</p>
            <p class="w-full truncate md:w-64">{{ data.description }}</p>
            <p class="text-sm text-gray-400">{{ data.created_at }}</p>
        </div>

        <div class="flex items-center gap-4 mt-4 md:mt-0">
            <Icon 
                name="line-md:link" 
                size="2rem" 
                class="cursor-pointer" 
                @click="navigateTo(data?.github_url, { external: true, open: { target: data.github_url } })"
            />
            <Icon 
                name="material-symbols:delete" 
                size="2rem" 
                class="cursor-pointer" 
                @click="removeProject"
            />
        </div>
    </div>
</template>


<script setup>

    const { data } = defineProps(["data"])
    const emit = defineEmits("reload")

    const { server_url } = useRuntimeConfig().public

    async function removeProject() {
        await $fetch(`${server_url}/delete?id=${data.id}`)
        emit("reload")
        
    }



</script>