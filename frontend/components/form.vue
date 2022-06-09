<template>
	<div>
		<br />
		<h2>Form</h2>
		<v-form ref="form" @submit.prevent>
			<v-text-field v-model="id" label="Id"></v-text-field>
			<v-text-field v-model="name" label="Name"></v-text-field>
			<v-text-field v-model="age" label="Age"></v-text-field>
			<v-text-field v-model="bloodtype" label="Bloodtype"></v-text-field>
			<v-text-field v-model="origin" label="Origin"></v-text-field>
			<v-btn @click="submit({ id, name, age, bloodtype, origin })">{{
				id ? "Edit" : "Submit"
				}}</v-btn>
		</v-form>
	</div>
</template>

<script>
export default {
	computed: {
		id: {
			get() {
				return this.$store.state.member.id;
			},
			set (value) {
				this.$store.commit("member/setId", value)
			},
		},
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
	methods:{
		async submit(member) {
			if(member.id) {
				await this.$axios.put("http://localhost:8003/api/members/" + member.id, {
					name: member.name,
					age: member.age,
					bloodtype: member.bloodtype,
					origin: member.origin
				})
			} else {
				console.log("POST mode: ");
				await this.$axios.post("/api/members/create", {
					name: member.name,
					age: member.age,
					bloodtype: member.bloodtype,
					origin: member.origin
				});
			}
			this.resetMember({ id: 0, name: "", age: "", bloodtype: 0, origin: "" })
			this.$store.commit(
				"members/storeData",
				(await this.$axios.get("/api/members/")).data
			);
		},
		resetMember(member) {
			this.$store.commit("member/setId", member.id)
			this.$store.commit("member/setName", member.name)
			this.$store.commit("member/setAge", member.age)
			this.$store.commit("member/setBloodtype", member.bloodtype)
			this.$store.commit("member/setOrigin", member.origin)
		}
	},
};
</script>

<style>

</style>