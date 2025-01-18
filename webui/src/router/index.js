import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import UserProfileView from '../views/UserprofileView.vue'
import MessageView from '../views/MessageView.vue'
import ChatsView from '../views/ChatsView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: LoginView},
		{path: '/user/:userId', component: UserProfileView},
		{path: '/message', component: MessageView},
		{path: '/chats', component: ChatsView},
	]
})

export default router
