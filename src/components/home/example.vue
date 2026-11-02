<template>
<script lang="js">
    export default{ 
        name: "Example",
        data() {
            return {
                items:[],
                loading: false,
                error: null,
                formData: {
                    name: "",
                    email: "",
                    age: "",},
                    isSubmitting: false,
                    submitError: null,
            };
        },
        mounted() {
            this.fetchData();
        },
        methods: {
            async fetchData() {
                this.loading = true;
                this.error = null;
                try {
                    const response = await fetch("https://jsonplaceholder.typicode.com/posts");
                    if (!response.ok) {
                        throw new Error("Network response was not ok");
                    }
                    const data = await response.json();
                    this.items = data;
                } catch (error) {
                    this.error = error.message;
                } finally {
                    this.loading = false;
                }
            },
            async submitForm() {
                this.isSubmitting = true;
                this.submitError = null;
                try {
                    const response = await fetch("https://jsonplaceholder.typicode.com/posts", {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify(this.formData),
                    });
                    if (!response.ok) {
                        throw new Error("Network response was not ok");
                    }
                    const result = await response.json();
                    console.log("Form submitted successfully:", result);
                } catch (error) {
                    this.submitError = error.message;
                } finally {
                    this.isSubmitting = false;
                }
            },
            resetForm() {
                 this.formData = { name: '', email: '', age: '' }
                         }
        },
    }

</script>
</template>