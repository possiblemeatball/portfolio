import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
    {
        path: '/',
        name: 'index',
        component: () => import("../components/HelloWorld.vue"),
        meta: {
            title: ""
        }
    }
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes
});
