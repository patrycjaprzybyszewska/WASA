import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import UserProfileView from '../views/UserprofileView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: LoginView},
		{path: '/user/:userId', component: UserProfileView},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
