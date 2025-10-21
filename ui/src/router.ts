import { createRouter, createWebHashHistory } from "vue-router";

import { MainPage, ListLapsPage, LapPage, ConfigPage } from "./pages";

const routes = [
  { path: "/", component: MainPage },
  { path: "/config", component: ConfigPage },
  { path: "/:trackId", component: ListLapsPage },
  { path: "/:trackId/:lap", component: LapPage },
];

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
