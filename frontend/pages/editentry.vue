<template>
	<div>
		<h2>Edit an item</h2>
		<v-form ref="form" @submit.prevent>
			<v-text-field v-model="ID" label="ID"></v-text-field>
			<v-text-field v-model="name" label="Name"></v-text-field>
			<v-text-field v-model="age" label="Age"></v-text-field>
			<v-text-field v-model="bloodtype" label="Bloodtype"></v-text-field>
			<v-text-field v-model="origin" label="Origin"></v-text-field>
			<v-btn @click="submit({ name, age, bloodtype, origin })">{{
				"Submit"
				}}</v-btn>
		</v-form>
	</div>
</template>

<script>
export default {
	computed: {
		name: {
			get() {
				return this.$store.state.member.name;
			},
			set (value) {
				this.$store.commit("member/setName", value)
			},
		},
		age: {
			get() {
				return this.$store.state.member.age;
			},
			set (value) {
				this.$store.commit("member/setAge", value)
			},
		},
		bloodtype: {
			get() {
				return this.$store.state.member.bloodtype;
			},
			set (value) {
				this.$store.commit("member/setBloodtype", value)
			},
		},
		origin: {
			get() {
				return this.$store.state.member.origin;
			},
			set (value) {
				this.$store.commit("member/setOrigin", value)
			},
		},
	},
	methods: {
		async submit(member) {
			if(member) {
				// Update member
				console.log("### DEBUG ###");
				console.log("update data to: ", member);
				await this.$axios.put("/api/members/:id", {
					name: member.name,
					age: member.age,
					bloodtype: member.bloodtype,
					origin: member.origin
				});
			}

		  this.resetmember({ ID: 0, name: "", age: "", bloodtype: 0, origin: "" })
		  this.$store.commit(
			  "members/storeData",
			  (await this.$axios.get("/api/members/")).data
  		);
		},
		resetmember(member) {
			this.$store.commit("member/setId", member.ID)
			this.$store.commit("member/setName", member.name)
			this.$store.commit("member/setAge", member.age)
			this.$store.commit("member/setBloodtype", member.bloodtype)
			this.$store.commit("member/setOrigin", member.origin)
		},
	},
};
</script>

<style>

</style>
