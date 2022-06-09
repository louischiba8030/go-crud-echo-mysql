<template>
	<div>
		<v-app>
			<v-main>
				<div>
					<v-app-bar color="gray accent-4" dark>
						<v-app-bar-title>Home</v-app-bar-title>
						<v-btn color="orange" nuxt to="addentry">Add</v-btn>
					</v-app-bar>
				</div>
				<v-container justify="center" align="center">
					<client-only>
						<Nuxt />
					</client-only>
				</v-container>
			</v-main>
			<v-footer>
				<div class="flex-grow-1"></div>
				<div>&copy; 2022 - {{ new Date().getFullYear() }} Demo Project</div>
			</v-footer>
		</v-app>
	</div>
</template>

<script>
export default {
	computed: {

	},
	methods: {
		async addMember () {
			console.log("Add Member: ");
		}
	},
};
</script>

<style>

</style>