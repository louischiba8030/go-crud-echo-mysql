<template>
	<div>
		<h2>Table</h2>
		<v-data-table :headers="headers" :items="members">
			<template v-slot:[`item.edit`]="{ item }">
				<v-btn color="success" nuxt to="editentry">Edit</v-btn>
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
				{ text: 'ID', value: 'ID' },
				{ text: 'Name', value: 'name' },
				{ text: 'Age', value: 'age' },
				{ text: 'Bloodtype', value: 'bloodtype' },
				{ text: 'Origin', value: 'origin' },
				{ text: 'Edit', value: 'edit' },
				{ text: 'Delete', value: 'delete'},
			],
		};
	},
	computed: {
		members() {
			return this.$store.state.members.data;
		}
	},
	async fetch() {
		this.$store.commit(
			"members/storeData",
			(await this.$axios.get("/api/members/")).data
		);
	},
	methods: {
		async deleteItem(ID) {
			await this.$axios.delete("/api/members/" + ID)
			this.$store.commit(
				"members/storeData",
				(await this.$axios.get("/api/members/")).data
			);
		},
		async editItem(member) { // member
			this.$store.commit("member/setId", member.ID);
			this.$store.commit("member/setName", member.name);
			this.$store.commit("member/setAge", member.age);
			this.$store.commit("member/setBloodtype", member.bloodtype);
			this.$store.commit("member/setOrigin", member.origin);
		}
	}
};
</script>

<style>

</style>