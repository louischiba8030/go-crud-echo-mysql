export const state = () => ({
	id: "",
	name: "",
	age: "",
	bloodtype: "",
	origin: ""
})

export const mutations = {
	setId: (state, data) => {
		state.id = data
	},
	setName: (state, data) => {
		state.name = data
	},
	setAge: (state, data) => {
		state.age = data
	},
	setBloodtype: (state, data) => {
		state.bloodtype = data
	},
	setOrigin: (state, data) => {
		state.origin = data
	}
}
