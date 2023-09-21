import axios from 'axios'

const BASE_URL = 'http://localhost:8080/api/v1'

type RegisterData = {
	firstName: string
	identifier: string
	lastName: string
	email: string
	phoneNumber: string
	password: string
}

export const authApi = axios.create({
	baseURL: BASE_URL,
	withCredentials: true,
})

authApi.defaults.headers.common['Content-Type'] = 'application/json'

export const registerAccount = async (newAccount: RegisterData) => {
	return await authApi.post('/auth/register', newAccount).then((res) => res.data)
}
