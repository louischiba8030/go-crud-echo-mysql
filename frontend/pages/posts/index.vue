<template>
	<div>
		<h2>Table</h2>
		<v-data-table :headers="headers" :items="posts">
			<template v-slot:[`item.edit`]="{ item }">
				<v-btn color="success" @click="editItem(item.ID)">Edit</v-btn>
			</template>
			<template v-slot:[`item.delete`]="{ item }">
				<v-btn color="danger" @click="deleteItem(item.ID)">Delete</v-btn>
			</template>
		</v-data-table>
	</div>
</template>

<script>
export default {
	data () {
		return {
			headers: [
				{ text: 'Id', value: 'ID' },
				{ text: 'Name', value: 'name' },
				{ text: 'Age', value: 'age' },
				{ text: 'Bloodtype', value: 'bloodtype' },
				{ text: 'Origin', value: 'origin'},
				{ text: 'Edit', value: 'edit' },
				{ text: 'Delete', value: 'delete'},
			],
		};
	},
	computed: {
		posts() {
			return this.$store.state.posts.data;
		}
	},
	async fetch() {
		this.$store.commit(
			"posts/storeData",
			(await this.$axios.get("/api/members/")).data
		);
	},
	methods: {
	}
};
</script>

<style>

</style>