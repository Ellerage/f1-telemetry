import { createRouter, createWebHashHistory } from "vue-router";

import { MainPage, ListLapsPage, LapPage } from "./pages";

const routes = [
  { path: "/", component: MainPage },
  { path: "/:trackId", component: ListLapsPage },
  { path: "/:trackId/:lap", component: LapPage },
];

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
