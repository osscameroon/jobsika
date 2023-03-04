<template>
    <div class="h-24 flex items-center justify-center" style="background: #d9d9d9;">
        <p class="text-base font-bold">
            {{ name }}
        </p>
    </div>
</template>

<script>
export default {
    name: "PaymentComponent",
    props: {
        name: {
            type: String,
            default: ""
        }
    }
}
</script>
